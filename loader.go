package main

import (
	"github.com/go-fsnotify/fsnotify"
	"github.com/huichen/sego"
	"log"
)

var segmenter sego.Segmenter

// init
func InitLoader() {
	LoadDict()
	go func() {
		RunWatcher("./dict.txt")
	}()
}

// load dict
func LoadDict() {
	dict := "dict.txt"
	segmenter.LoadDictionary(dict)
}

// api interface
func GetSegmenter() sego.Segmenter {
	return segmenter
}

// watch file change
func RunWatcher(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("file modified:", event.Name)
					// 若 dict.txt 改变 reload
					LoadDict()
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
