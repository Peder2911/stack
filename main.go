/*
Stack is a small service that lets me save things for later.

I use it to put a pin on links, documents and ideas that I
want to rediscover at a later stage.
*/
package main

import (
	"os"
	"fmt"
)

type Stack string

const ChunkSize int64 = 128

func (s *Stack) size() (int64, error) {
	i, err := os.Stat(string(*s))
	if err != nil {
		return 0, nil
	}
	nbytes := i.Size()
	aligned := nbytes % ChunkSize
	if aligned != 0 {
		fmt.Errorf("Stack not aligned: %v\n", aligned)
	}
	nchunks := nbytes / ChunkSize
	return nchunks, nil
}

func (s *Stack) Pop() (string, error) {
	f, err := os.OpenFile(string(*s), os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	size, err := s.size()
	if err != nil {
		return "", err
	}
	if size == 0 {
		return "", nil
	}
	item := make([]byte, ChunkSize)
	f.Seek(-ChunkSize, 2)
	_, err = f.Read(item)
	if err != nil {
		return "", err
	}
	f.Seek(0, 0)
	err = f.Truncate((size - 1) * ChunkSize)
	if err != nil {
		panic(err)
	}
	return string(item), nil
}
func (s *Stack) Push(item string) error {
	f, err := os.OpenFile(string(*s), os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	buf := make([]byte, ChunkSize)
	for i, b := range []byte(item) {
		buf[i] = b
		if i > int(ChunkSize) {
			break
		}
	}
	f.Seek(0, 2)
	f.Write(buf)
	return nil
}

func main() {
	stackPath := os.Args[1]
	s := Stack(stackPath)
	var err error
	err = s.Push("A")
	err = s.Push("A")
	err = s.Push("A")
	err = s.Push("B")
	err = s.Push("B")
	item, err := s.Pop()
	if err != nil {
		panic(err)
	}
	println(item)
}
