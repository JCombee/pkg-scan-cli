package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jcombee/pkg-scan/managers"
)

type FileReaderHandler struct {
	Config Config
	Files  managers.Files
}

func FilesDefault(c Config) managers.FilesHandler {
	frh := &FileReaderHandler{
		Config: c,
	}
	fh := managers.FilesHandler(frh)
	return fh
}

func (fr *FileReaderHandler) ReadSource() managers.Files {
	return fr.LoopSource(fr.Config.FileRoot)
}

func (fr *FileReaderHandler) LoopSource(path string) managers.Files {
	fia, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range fia {
		if fi.IsDir() {
			fmt.Println("Test")
			fr.LoopSource(path + "/" + fi.Name())
		} else {
			fr.Files = append(fr.Files, managers.File{
				Name: fi.Name(),
				Path: path + "/",
			})
		}
	}
	return fr.Files
}

func (fr *FileReaderHandler) ReadData(files managers.Files) managers.Files {
	var newFiles managers.Files
	for _, file := range files {
		path := file.Path + file.Name
		r, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		f, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		file.Data = string(f)
		newFiles = append(newFiles, file)
	}
	return newFiles
}
