package telebot

import (
	"errors"
	"fmt"
	"sync"
)

var KeyNotFound = errors.New("conversation key not found")

// InMemoryStorage is a thread-safe in-memory implementation of the IStorage interface.
type InMemoryStorage struct {
	// keyStrategy defines how to calculate keys for each conversation.
	keyStrategy KeyStrategy
	// conversations is a map of key -> state, which tracks at which point of each conversation a user/chat is.
	conversations map[string]State
	// lock allows us to ensure synchronous data access.
	lock sync.RWMutex
}

func NewInMemoryStorage(strategy KeyStrategy) *InMemoryStorage {
	return &InMemoryStorage{
		keyStrategy:   strategy,
		lock:          sync.RWMutex{},
		conversations: map[string]State{},
	}
}

func (c *InMemoryStorage) Get(ctx IContext) (*State, error) {
	key := StateKey(ctx, c.keyStrategy)
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.conversations == nil {
		return nil, KeyNotFound
	}

	s, ok := c.conversations[key]
	if !ok {
		return nil, KeyNotFound
	}
	fmt.Println("InMemoryStorage Get key:", key)
	return &s, nil
}

func (c *InMemoryStorage) Set(ctx IContext, state State) error {
	key := StateKey(ctx, c.keyStrategy)

	c.lock.Lock()
	defer c.lock.Unlock()

	if c.conversations == nil {
		c.conversations = map[string]State{}
	}

	c.conversations[key] = state
	return nil
}

func (c *InMemoryStorage) Delete(ctx IContext) error {
	key := StateKey(ctx, c.keyStrategy)

	c.lock.Lock()
	defer c.lock.Unlock()

	if c.conversations == nil {
		return nil
	}

	delete(c.conversations, key)
	return nil
}
