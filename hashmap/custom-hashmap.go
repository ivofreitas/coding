package main

type Node struct {
	key   string
	value string
	next  *Node
}

type CustomHashMap struct {
	buckets []*Node
	size    int
}

func New(size int) *CustomHashMap {
	return &CustomHashMap{buckets: make([]*Node, size), size: size}
}

func (hm *CustomHashMap) Put(key, value string) {
	index := hm.hashFunction(key)

	node := &Node{key: key, value: value, next: nil}

	if hm.buckets[index] == nil {
		hm.buckets[index] = node
		return
	}

	current := hm.buckets[index]

	for current != nil {

		if current.key == key {
			current.value = value
			return
		}

		current = current.next
	}

	node.next = hm.buckets[index]
	hm.buckets[index] = node
}

func (hm *CustomHashMap) Get(key string) string {
	index := hm.hashFunction(key)

	current := hm.buckets[index]

	for current != nil {

		if current.key == key {
			return current.value
		}

		current = current.next
	}

	return ""
}

func (hm *CustomHashMap) Delete(key string) {
	index := hm.hashFunction(key)
	current := hm.buckets[index]

	if current == nil {
		return
	}

	if current.key == key {
		current = current.next
		hm.buckets[index] = current
		return
	}

	var prev *Node

	for current != nil {

		if current.key == key {
			prev.next = current.next
			return
		}

		prev = current
		current = current.next
	}

}

func (hm *CustomHashMap) hashFunction(key string) uint {
	return uint(len(key) % hm.size)
}
