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
	"fmt"
	"sync"
)

type op int

const (
	opAdd op = iota
	opDel
)

var (
	pollInitOnce sync.Once
	poll         poller
)

type poller interface {
	wait() error
	control(fd *fdOperator, op op) error
}

func createPoller() {
	var err error
	poll, err = openpoll()
	if err != nil {
		panic(fmt.Sprintf("gopkg.connstate openpoll failed, err: %v", err))
	}
	go func() {
		err := poll.wait()
		fmt.Printf("gopkg.connstate epoll wait exit, err: %v\n", err)
	}()
}
