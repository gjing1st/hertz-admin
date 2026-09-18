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

//go:build darwin || netbsd || freebsd || openbsd || dragonfly
// +build darwin netbsd freebsd openbsd dragonfly

package connstate

import (
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

type kqueue struct {
	fd int
}

func (p *kqueue) wait() error {
	events := make([]syscall.Kevent_t, 1024)
	timeout := &syscall.Timespec{Sec: 0, Nsec: 0}
	var n int
	var err error
	for {
		// timeout=0 must be set to avoid getting stuck in a blocking syscall,
		// which could occupy a P until runtime.sysmon thread handoff it.
		n, err = syscall.Kevent(p.fd, nil, events, timeout)
		if err != nil && err != syscall.EINTR {
			// exit gracefully
			if err == syscall.EBADF {
				return nil
			}
			return err
		}
		if n <= 0 {
			time.Sleep(10 * time.Millisecond) // avoid busy loop
		} else {
			var cbs []OnRemoteClosed
			for i := 0; i < n; i++ {
				ev := &events[i]
				op := *(**fdOperator)(unsafe.Pointer(&ev.Udata))
				if conn := (*connStater)(atomic.LoadPointer(&op.conn)); conn != nil {
					if ev.Flags&(syscall.EV_EOF) != 0 {
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

func (p *kqueue) control(fd *fdOperator, op op) error {
	evs := make([]syscall.Kevent_t, 1)
	evs[0].Ident = uint64(fd.fd)
	*(**fdOperator)(unsafe.Pointer(&evs[0].Udata)) = fd
	if op == opAdd {
		evs[0].Filter = syscall.EVFILT_READ
		evs[0].Flags = syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_CLEAR
		// prevent ordinary data from triggering
		evs[0].Flags |= syscall.EV_OOBAND
		evs[0].Fflags = syscall.NOTE_LOWAT
		evs[0].Data = 0x7FFFFFFF
		_, err := syscall.Kevent(p.fd, evs, nil, nil)
		return err
	} else {
		evs[0].Filter = syscall.EVFILT_READ
		evs[0].Flags = syscall.EV_DELETE
		_, err := syscall.Kevent(p.fd, evs, nil, nil)
		return err
	}
}

func openpoll() (p poller, err error) {
	fd, err := syscall.Kqueue()
	if err != nil {
		return nil, err
	}
	_, err = syscall.Kevent(fd, []syscall.Kevent_t{{
		Ident:  0,
		Filter: syscall.EVFILT_USER,
		Flags:  syscall.EV_ADD | syscall.EV_CLEAR,
	}}, nil, nil)
	if err != nil {
		syscall.Close(fd)
		return nil, err
	}
	return &kqueue{fd: fd}, nil
}
