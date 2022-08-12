package main

type Memory struct {
	store       []byte
}
func NewMemory() *Memory {
	return &Memory{}
}