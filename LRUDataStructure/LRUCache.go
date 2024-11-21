package lrudatastructure

/*
LRU CACHE:
	- An LRU Cache (Least Recently Used Cache) is a data structure designed to store a fixed number of items
	- While ensuring that the least recently accessed (used) item is removed when the cache reaches its capacity.
	- It is widely used in systems where memory is limited and certain data needs to be reused frequently.
*/
type LRUCache struct {
	Head *Node
	Tail *Node
	HT   map[int]*Node
	Cap  int
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{HT: make(map[int]*Node), Cap: capacity}
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.HT[key]
	if ok {
		lru.Remove(node)
		lru.Add(node)
		return node.Val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	node, ok := lru.HT[key]
	if ok {
		node.Val = value
		lru.Remove(node)
		lru.Add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		lru.HT[key] = node
		lru.Add(node)
	}
	if len(lru.HT) > lru.Cap {
		delete(lru.HT, lru.Tail.Key)
		lru.Remove(lru.Tail)
	}
}

func (lru *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = lru.Head
	if lru.Head != nil {
		lru.Head.Prev = node
	}
	lru.Head = node
	if lru.Tail == nil {
		lru.Tail = node
	}
}

func (lru *LRUCache) Remove(node *Node) {
	if node != lru.Head {
		node.Prev.Next = node.Next
	} else {
		lru.Head = node.Next
	}
	if node != lru.Tail {
		node.Next.Prev = node.Prev
	} else {
		lru.Tail = node.Prev
	}
}
