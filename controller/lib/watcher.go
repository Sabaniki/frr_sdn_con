package controller

import (
	"log"
	"regexp"

	"github.com/fsnotify/fsnotify"
)

func Start() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	pathToConfig := "/home/vsix/frr_sdn_con/config/"
	hosts, err := DeserializeHosts(pathToConfig + "hosts.yml")
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write &&
					regexp.MustCompile(`.*\.json$`).Match([]byte(event.Name)) {
					log.Println("modified file:", event.Name)
					if regexp.MustCompile(`.*\/metric\.json$`).Match([]byte(event.Name)) {
						SetMed(event.Name)
					} else if regexp.MustCompile(`.*\/interface\.json$`).Match([]byte(event.Name)) {
						ConfigInterface(event.Name)
					} else {
						log.Println("other!")
						ConfigInterface(event.Name)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	if err != nil {
		log.Fatal(err)
	}
	for _, host := range hosts {
		err = watcher.Add(pathToConfig + host.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	<-done
}
