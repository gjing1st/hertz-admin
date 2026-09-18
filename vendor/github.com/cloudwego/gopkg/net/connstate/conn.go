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
	"errors"
	"net"
	"sync/atomic"
	"syscall"
	"unsafe"
)

type ConnState uint32

const (
	// StateOK means the connection is normal.
	StateOK ConnState = iota
	// StateRemoteClosed means the remote side has closed the connection.
	StateRemoteClosed
	// StateClosed means the connection has been closed by local side.
	StateClosed
)

// ConnStater is the interface to get the ConnState of a connection.
// Must call Close to release it if you're going to close the connection.
type ConnStater interface {
	Close() error
	State() ConnState
}

type Option func(stater *connStater)

// OnRemoteClosed is a callback function type that will be invoked
// when the remote side of the connection is closed.
type OnRemoteClosed func()

// WithOnRemoteClosed sets a callback function that will be called when
// the remote side closes the connection. The callback is invoked only once,
// when the state transitions from StateOK to StateRemoteClosed.
func WithOnRemoteClosed(fn OnRemoteClosed) Option {
	return func(stater *connStater) {
		stater.onRemoteClosed = fn
	}
}

// ListenConnState returns a ConnStater for the given connection.
// Conn must be a syscall.Conn.
func ListenConnState(conn net.Conn, opts ...Option) (ConnStater, error) {
	pollInitOnce.Do(createPoller)
	sysConn, ok := conn.(syscall.Conn)
	if !ok {
		return nil, errors.New("conn is not syscall.Conn")
	}
	rawConn, err := sysConn.SyscallConn()
	if err != nil {
		return nil, err
	}
	var fd *fdOperator
	var opAddErr error
	err = rawConn.Control(func(fileDescriptor uintptr) {
		fd = pollcache.alloc()
		fd.fd = int(fileDescriptor)
		cs := &connStater{fd: unsafe.Pointer(fd)}
		for _, opt := range opts {
			opt(cs)
		}
		atomic.StorePointer(&fd.conn, unsafe.Pointer(cs))
		opAddErr = poll.control(fd, opAdd)
	})
	if fd != nil {
		if err != nil && opAddErr == nil {
			// if rawConn is closed, poller will delete the fd by itself
			_ = rawConn.Control(func(_ uintptr) {
				_ = poll.control(fd, opDel)
			})
		}
		if err != nil || opAddErr != nil {
			atomic.StorePointer(&fd.conn, nil)
			pollcache.freeable(fd)
		}
	}
	if err != nil {
		return nil, err
	}
	if opAddErr != nil {
		return nil, opAddErr
	}
	return (*connStater)(atomic.LoadPointer(&fd.conn)), nil
}

type connStater struct {
	fd    unsafe.Pointer // *fdOperator
	state uint32

	onRemoteClosed OnRemoteClosed
}

func (c *connStater) Close() error {
	fd := (*fdOperator)(atomic.LoadPointer(&c.fd))
	if fd != nil && atomic.CompareAndSwapPointer(&c.fd, unsafe.Pointer(fd), nil) {
		atomic.StoreUint32(&c.state, uint32(StateClosed))
		_ = poll.control(fd, opDel)
		atomic.StorePointer(&fd.conn, nil)
		pollcache.freeable(fd)
	}
	return nil
}

func (c *connStater) State() ConnState {
	return ConnState(atomic.LoadUint32(&c.state))
}
