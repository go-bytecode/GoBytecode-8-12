package main

import (
	"github.com/holiman/uint256"
	"sync"
)

var stackPool = sync.Pool{
	New: func() interface{} {
		return &Stack{data: make([]uint256.Int, 0, 16)}
	},
}

type Stack struct {
	data []uint256.Int
}
func newstack() *Stack {
	return stackPool.Get().(*Stack)
}

func (st *Stack) pop() (ret uint256.Int) {
	ret = st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return
}

func (st *Stack) peek() *uint256.Int {
	return &st.data[st.len()-1]
}
func (st *Stack) len() int {
	return len(st.data)
}

func (st *Stack) push(d *uint256.Int) {
	// NOTE push limit (1024) is checked in baseCheck
	st.data = append(st.data, *d)
}