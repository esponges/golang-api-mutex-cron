package main

import "sync"

type MyStruct struct {
	globalVariable int
	mutex          sync.Mutex
}

func NewInstance() *MyStruct {
	return &MyStruct{}
}

func (m *MyStruct) SetGlobalVariable(value int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.globalVariable = value
}

func (m *MyStruct) GetGlobalVariable() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.globalVariable
}

func (m *MyStruct) IncrementGlobalVariable() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.globalVariable++
}
