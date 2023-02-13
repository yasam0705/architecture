package server

import "sync"

var pool = make(map[string]*Client)

type Client struct { // save client to redis
	UserId             string
	Mu                 *sync.Mutex
	ConcurrentRequests map[string]int
}

func NewClient(userId string) *Client {
	if c, ok := pool[userId]; ok {
		return c
	}

	c := &Client{
		UserId:             userId,
		Mu:                 &sync.Mutex{},
		ConcurrentRequests: make(map[string]int),
	}

	pool[userId] = c
	return c
}

func (c *Client) Increment(method string) int {
	c.Mu.Lock()
	c.ConcurrentRequests[method] += 1
	defer c.Mu.Unlock()
	return c.ConcurrentRequests[method]
}

func (c *Client) Decrement(method string) int {
	c.Mu.Lock()
	c.ConcurrentRequests[method] -= 1
	defer c.Mu.Unlock()
	return c.ConcurrentRequests[method]
}

func (c *Client) Remove() {
	delete(pool, c.UserId)
}
