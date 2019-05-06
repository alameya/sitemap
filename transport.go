package sitemap

import (
	"context"
	"io/ioutil"
	"net/http"
	"sync"
)

// Transport bus for grab html by url and parse a new urls recursively and parallel.
// Can be stopped by transport.Close().
type Transport struct {
	Close      context.CancelFunc
	ctx        context.Context
	parse      ParserFunc
	maxDepth   uint
	uniqueUrls map[string]uint // url to deep
	urlCh      chan *URLChannelItem
	mx         *sync.RWMutex
	wg         *sync.WaitGroup
	err        chan error
	errHandler ErrorHandler
}

// ErrorHandler interface for function which can to be passed to transport.
type ErrorHandler func(err error)

// URLChannelItem can be passed to Transport.urlCh.
type URLChannelItem struct {
	URL  string
	Deep uint
}

// ParserFunc interface of parser func that can be passed to transport.
type ParserFunc func(body []byte) (urls []string, err error)

// NewTransport creates a new Transport object.
func NewTransport(ctx context.Context, parse ParserFunc, errHandler ErrorHandler, maxDepth, parallel uint) *Transport {
	ctx, cancelFunc := context.WithCancel(ctx)

	transport := &Transport{
		ctx:        ctx,
		Close:      cancelFunc,
		parse:      parse,
		errHandler: errHandler,
		maxDepth:   maxDepth,
		urlCh:      make(chan *URLChannelItem, parallel),
		uniqueUrls: make(map[string]uint),
		err:        make(chan error, 1),
		mx:         &sync.RWMutex{},
		wg:         &sync.WaitGroup{},
	}

	go transport.run()

	return transport
}

func (transport *Transport) run() {
	for {
		select {
		case urlItem := <-transport.urlCh:
			go transport.handle(urlItem.URL, urlItem.Deep)
		case err := <-transport.err:
			transport.errHandler(err)
		case <-transport.ctx.Done():
			return
		}
	}
}

// Fetch collect urls by
func (transport *Transport) Fetch(url string) (urls []string) {
	transport.addToURLChannel(&URLChannelItem{
		URL:  url,
		Deep: 0, // initial value
	})

	transport.wg.Wait()

	for url := range transport.uniqueUrls {
		urls = append(urls, url)
	}

	transport.uniqueUrls = make(map[string]uint)

	return
}
func (transport *Transport) addToURLChannel(item *URLChannelItem) {
	transport.wg.Add(1)
	transport.urlCh <- item
}

func (transport *Transport) handle(url string, deep uint) {
	defer transport.wg.Done()
	if !transport.addToMap(url, deep) {
		return
	}

	if deep == transport.maxDepth {
		return
	}

	newUrls, err := transport.fetch(url)
	if err != nil {
		transport.err <- err
	}

	deep++
	for _, childURL := range newUrls {
		transport.addToURLChannel(&URLChannelItem{
			URL:  childURL,
			Deep: deep,
		})
	}

	return
}

func (transport *Transport) fetch(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []string{}, err
	}

	return transport.parse(body)
}

func (transport *Transport) addToMap(url string, deep uint) bool {
	if url == "" {
		return false
	}

	transport.mx.Lock()
	defer transport.mx.Unlock()

	if _, exists := transport.uniqueUrls[url]; exists {
		return false
	}

	transport.uniqueUrls[url] = deep

	return true
}
