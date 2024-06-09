package stack

import (
   "os"
   "fmt"
   "bytes"
)

type Stack string

const ChunkSize int64 = 128

func (s *Stack) Size() (int64, error) {
	i, err := os.Stat(string(*s))
	if err != nil {
		return 0, err 
	}
	nbytes := i.Size()
	aligned := nbytes % ChunkSize
	if aligned != 0 {
		return 0,fmt.Errorf("Stack not aligned: %v\n", aligned)
	}
	nchunks := nbytes / ChunkSize
	return nchunks, nil
}

func (s *Stack) Pop() (string, error) {
	f, err := os.OpenFile(string(*s), os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	size, err := s.Size()
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
	item = bytes.Trim(item, "\x00")
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

func (s *Stack) Truncate() error {
	f, err := os.OpenFile(string(*s), os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	err = f.Truncate(0)
	if err != nil {
		panic(err)
	}
	return nil
}
