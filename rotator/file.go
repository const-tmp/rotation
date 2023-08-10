package rotator

import (
	"os"
	"sync"
)

type File struct {
	sync.Mutex
	Path, value string
}

func NewFile(path string) *File {
	return &File{Path: path}
}

func (f *File) Rotate() {
	if data, err := os.ReadFile(f.Path); err == nil {
		f.Lock()
		f.value = string(data)
		f.Unlock()
	}
}

func (f *File) GetValue() string {
	f.Lock()
	defer f.Unlock()
	return f.value
}
