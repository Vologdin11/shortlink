package inmemory

import (
	"errors"
	avl "github.com/emirpasic/gods/trees/avltree"
	"sync"
)

type Inmemory struct {
	links         *avl.Tree
	shortlinks    *avl.Tree
	mutLinks      sync.RWMutex
	mutShortlinks sync.RWMutex
	wg            sync.WaitGroup
}

func NewInmemory() *Inmemory {
	return &Inmemory{
		links:         avl.NewWithStringComparator(),
		shortlinks:    avl.NewWithStringComparator(),
		mutLinks:      sync.RWMutex{},
		mutShortlinks: sync.RWMutex{},
		wg:            sync.WaitGroup{},
	}
}

func (i *Inmemory) GetLink(shortLink string) (string, error) {
	i.mutLinks.RLock()
	defer i.mutLinks.RUnlock()

	link, check := i.links.Get(shortLink)
	if !check {
		return "", errors.New("link not found")
	}
	return link.(string), nil
}

func (i *Inmemory) GetShortLink(link string) (string, error) {
	i.mutShortlinks.RLock()
	defer i.mutShortlinks.RUnlock()

	shortlink, check := i.shortlinks.Get(link)
	if !check {
		return "", errors.New("shortlink not found")
	}
	return shortlink.(string), nil
}

func (i *Inmemory) AddLink(link, shortlink string) error {
	i.wg.Add(2)

	go i.putLink(shortlink, link)
	go i.putShortlink(link, shortlink)

	i.wg.Wait()

	return nil
}

func (i *Inmemory) putLink(key, value string) {
	i.mutLinks.Lock()
	defer i.wg.Done()
	defer i.mutLinks.Unlock()

	i.links.Put(key, value)
}

func (i *Inmemory) putShortlink(key, value string) {
	i.mutShortlinks.Lock()
	defer i.wg.Done()
	defer i.mutShortlinks.Unlock()

	i.shortlinks.Put(key, value)
}
