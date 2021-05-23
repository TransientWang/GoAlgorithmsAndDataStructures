package datastructure

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type HashLinkNode struct {
	prev, next *HashLinkNode
	Key        string
	Val        interface{}
}

func (l *HashLinkNode) Next() (nextNode *HashLinkNode, hasNext bool) {
	if l == nil {
		return
	}

	return l.next, l.next != nil
}
func (l *HashLinkNode) Append(key string, val interface{}) (newLink *HashLinkNode, err error) {
	if l.next != nil {
		return newLink, errors.New("next link not nil")
	}
	newLink = &HashLinkNode{
		prev: l,
		Val:  val,
		Key:  key,
	}
	l.next = newLink
	return
}

func (l *HashLinkNode) Delete() (prev, next *HashLinkNode, err error) {
	if l == nil {
		return nil, nil, errors.New("node is nil")
	}
	prev, next = l.prev, l.next
	if l.prev != nil {
		l.prev.next = l.next
	}
	if l.next.prev != nil {
		l.next.prev = l.prev
	}

	l.prev, l.next = nil, nil
	return
}

type hashTable struct {
	arr, backupArr []*HashLinkNode
	count          int
}

func NewHashTable() *hashTable {
	return &hashTable{
		arr: make([]*HashLinkNode, 8, 8),
	}
}

func (h *hashTable) growth() {
	h.backupArr = make([]*HashLinkNode, 2*len(h.arr), 2*len(h.arr))
	for _, node := range h.arr {
		for node != nil {
			sum := h.sum(&node.Key)
			idx := int(sum) % len(h.backupArr)
			newNode := h.backupArr[idx]
			if newNode == nil {
				h.backupArr[idx] = &HashLinkNode{
					Val: node.Val,
					Key: node.Key,
				}
				continue
			}

			for {
				if newNode.next == nil {
					newNode.Append(node.Key, node.Val)
					break
				} else {
					newNode = newNode.next
				}

			}
			node = node.next

		}

	}
	h.arr = h.backupArr
	h.backupArr = nil
}
func (h *hashTable) sum(key *string) (sum int64) {
	bytes := sha256.Sum256([]byte(*key))
	for _, b := range bytes {
		sum += int64(b)
	}
	return
}
func (h *hashTable) idx(key *string) int64 {
	return h.sum(key) % int64(len(h.arr))
}
func (h *hashTable) Set(key string, val interface{}) {
	defer fmt.Println("---------------------------------------")
	h.count++
	load := float32(h.count) / float32(len(h.arr))
	fmt.Println("key", key, "load", load)
	if load >= 0.75 {
		h.growth()
	}

	idx := h.idx(&key)
	fmt.Println("idx", idx)
	node := h.arr[idx]
	if node == nil {
		h.arr[idx] = &HashLinkNode{
			Key: key,
			Val: val,
		}
		fmt.Println("key ", key, " insert into link head")
		return
	}
	for node.next != nil {
		node = node.next
	}
	_, err := node.Append(key, val)
	if err != nil {
		panic(err)
	}
	fmt.Println("key ", key, " insert into link tail")

}

func (h *hashTable) Get(key string) interface{} {
	idx := h.idx(&key)
	node := h.arr[idx]
	for {
		if node == nil {
			return nil
		} else if node.Key == key {
			return node.Val
		} else {
			node = node.next
		}
	}
}
