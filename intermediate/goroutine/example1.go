package goroutine

import (
	"fmt"
	"sync"
)

func execute() string {
	channelResult := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)

	go GetMessage(channelResult, &wg)

	response := <-channelResult

	wg.Wait()

	return response

}

func GetMessage(channelResult chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	username := "Lucho"

	response := fmt.Sprintf("Hello %s", username)
	channelResult <- response
}
