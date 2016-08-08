package main

import (
	"github.com/jcombee/pkg-scan/managers"
)

func main() {
	c := ConfigDefault()
	c.InitArgs()
	fr := FilesDefault(c)
	e := managers.Default(fr)
	e.Run()
}
