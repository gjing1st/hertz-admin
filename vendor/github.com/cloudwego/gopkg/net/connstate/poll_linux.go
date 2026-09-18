// Copyright 2025 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connstate

import (
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

const _EPOLLET uint32 = 0x80000000

type epoller struct {
	epfd int
}

//go:nocheckptr
func (p *epoller) wait() error {
	events := make([]syscall.EpollEvent, 1024)
	var n int
	var err error
	for {
		// timeout=0 must be set to avoid getting stuck in a blocking syscall,
		// which could occupy a P until runtime.sysmon thread handoff it.
		n, err = syscall.EpollWait(p.epfd, events, 0)
		if err != nil && err != syscall.EINTR {
			return err
		}
		if n <= 0 {
			time.Sleep(10 * time.Millisecond) // avoid busy loop
		} else {
			var cbs []OnRemoteClosed
			for i := 0; i < n; i++ {
				ev := &events[i]
				op := *(**fdOperator)(unsafe.Pointer(&ev.Fd))
				if conn := (*connStater)(atomic.LoadPointer(&op.conn)); conn != nil {
					if ev.Events&(syscall.EPOLLHUP|syscall.EPOLLRDHUP|syscall.EPOLLERR) != 0 {
						if atomic.CompareAndSwapUint32(&conn.state, uint32(StateOK), uint32(StateRemoteClosed)) {
							if conn.onRemoteClosed != nil {
								cbs = append(cbs, conn.onRemoteClosed)
							}
						}
					}
				}
			}
			// execute close callbacks on single goroutine.
			if len(cbs) > 0 {
				go func(cbs []OnRemoteClosed) {
					for _, cb := range cbs {
						cb()
					}
				}(cbs)
			}
		}
		// we can make sure that there is no op remaining if finished handling all events
		pollcache.free()
	}
}

func (p *epoller) control(fd *fdOperator, op op) error {
	if op == opAdd {
		var ev syscall.EpollEvent
		*(**fdOperator)(unsafe.Pointer(&ev.Fd)) = fd
		ev.Events = syscall.EPOLLHUP | syscall.EPOLLRDHUP | syscall.EPOLLERR | _EPOLLET
		return syscall.EpollCtl(p.epfd, syscall.EPOLL_CTL_ADD, fd.fd, &ev)
	} else {
		var ev syscall.EpollEvent
		return syscall.EpollCtl(p.epfd, syscall.EPOLL_CTL_DEL, fd.fd, &ev)
	}
}

func openpoll() (p poller, err error) {
	var epfd int
	epfd, err = syscall.EpollCreate(1)
	if err != nil {
		return nil, err
	}
	return &epoller{epfd: epfd}, nil
}
