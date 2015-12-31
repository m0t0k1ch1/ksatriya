package ksatriya

import "net/http"

const (
	RootPathDefault = "/static"
	RootDirDefault  = "static"
)

type Root struct {
	path string
	dir  http.FileSystem
}

func (root *Root) Path() string {
	return root.path
}
func (root *Root) SetPath(val string) {
	root.path = val + "/*filepath"
}

func (root *Root) Dir() http.FileSystem {
	return root.dir
}
func (root *Root) SetDir(val string) {
	root.dir = http.Dir(val)
}

func NewRoot() *Root {
	root := &Root{}
	root.SetPath(RootPathDefault)
	root.SetDir(RootDirDefault)
	return root
}
