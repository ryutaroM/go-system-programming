package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	emitter := make(chan interface{})
	source := observable.Observable(emitter)

	//watching the events
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			line := item.(string)
			if strings.HasPrefix(line, "func ") {
				fmt.Println(line)
			}
		},
		ErrHandler: func(err error) {
			fmt.Printf("Encounterd error: %v\n", err)
		},
		DoneHandler: func() {
			fmt.Println("Done!!")
		},
	}

	// connecting the observable to the observer
	sub := source.Subscribe(watcher)

	// input value to the observable
	go func() {
		content, err := os.ReadFile("./src/observer/main.go")
		if err != nil {
			emitter <- err
		} else {
			for _, line := range strings.Split(string(content), "\n") {
				emitter <- line
			}
		}
		close(emitter)
	}()

	<-sub

}
