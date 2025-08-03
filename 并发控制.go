package main

import (
	"os"

	"golang.org/x/tools/godoc/vfs"
	//	"golang.org/x/tools/godoc/vfs/gatefs"
)

type gate chan bool

func (g gate) enter() { g <- true }
func (g gate) leave() { <-g }

type gatefs struct {
	fs vfs.FileSystem
	gate
}

func (fs gatefs) Lstat(p string) (os.FileInfo, error) {
	fs.enter()
	defer fs.leave()
	return fs.fs.Lstat(p)
}
func main() {
	//书中未完成
}
