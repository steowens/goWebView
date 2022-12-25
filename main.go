package main

import "C"
import (
	"log"
	"sync"

	"github.com/webview/webview"
)

var wg sync.WaitGroup

const (
	WEBVIEW_HINT_NONE  webview.Hint = 0
	WEBVIEW_HINT_MIN   webview.Hint = 1
	WEBVIEW_HINT_MAX   webview.Hint = 2
	WEBVIEW_HINT_FIXED webview.Hint = 3
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel is a channel which can accept an DataEvent
type DataChannel chan DataEvent

// DataChannelSlice is a slice of DataChannels
type DataChannelSlice []DataChannel

// EventBus stores the information about subscribers interested for // a particular topic
type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

var eb = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		// this is done because the slices refer to same array even though they are passed by value
		// thus we are creating a new slice with our elements thus preserve locking correctly.
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

func init() {

}

func main() {
	log.Println("Hello")
	wg.Add(2)
	go launchWindow("Window 1")
	go launchWindow("Window 2")
	wg.Wait()
}

func launchWindow(title string) webview.WebView {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle(title)
	w.SetSize(480, 320, WEBVIEW_HINT_NONE)
	w.SetHtml("Thanks for using webview!")
	w.Run()
	return w
}
