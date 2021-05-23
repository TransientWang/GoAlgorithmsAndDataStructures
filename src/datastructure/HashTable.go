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
	fmt.Print("delete ", l.Val, func() interface{} {
		if next != nil {
			if next.next != nil {
				fmt.Println(" next.next ", next.next.Val)
			}
			return next.Val
		}

		return nil
	}())
	fmt.Print("\n")
	if prev != nil {
		l.prev.next = next
	}
	if next != nil {
		l.next.prev = prev
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
	add := func(node *HashLinkNode) {
		sum := h.sum(&node.Key)
		idx := int(sum) % len(h.backupArr)
		newNode := h.backupArr[idx]
		if newNode == nil {
			h.backupArr[idx] = &HashLinkNode{
				Val: node.Val,
				Key: node.Key,
			}
			return
		}

		for {
			if newNode.next == nil {
				newNode.Append(node.Key, node.Val)
				return
			} else {
				newNode = newNode.next
			}

		}
	}
	for _, node := range h.arr {
		for node != nil {
			add(node)
			node = node.next
		}

	}
	//fmt.Println("=================growth=================")
	//fmt.Println(h.arr)
	//fmt.Println(h.backupArr)
	//fmt.Println("=================growth=================")
	h.arr = nil
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
	//defer fmt.Println("---------------------------------------")
	h.count++
	load := float32(h.count) / float32(len(h.arr))
	//fmt.Println("key", key, "load", load)
	if load >= 0.75 {
		h.growth()
	}

	idx := h.idx(&key)
	//fmt.Println("idx", idx)
	node := h.arr[idx]
	if node == nil {
		h.arr[idx] = &HashLinkNode{
			Key: key,
			Val: val,
		}
		//fmt.Println("key ", key, " insert into link head")
		return
	}
	for node.next != nil {
		if node.Key == key {
			return
		}
		node = node.next
	}
	_, err := node.Append(key, val)
	if err != nil {
		panic(err)
	}
	//fmt.Println("key ", key, " insert into link tail")

}
func (l *HashLinkNode) Search(key string) *HashLinkNode {
	if l == nil {
		return nil
	}
	if l.Key == key {
		return l
	}
	return l.next.Search(key)
}
func (h *hashTable) Get(key string) interface{} {
	idx := h.idx(&key)
	node := h.arr[idx].Search(key)
	if node != nil {
		return node.Val
	}
	return nil
}

func (h *hashTable) Delete(key string) {
	idx := h.idx(&key)
	node := h.arr[idx].Search(key)
	if node == nil {
		return
	}
	//fmt.Println("search", node.Val)
	_, next, _ := node.Delete()
	if node == h.arr[idx] {
		//fmt.Println("node ",node.Val, node.prev, node.next)

		h.arr[idx] = next
		//fmt.Println("h.arr[idx] ", h.arr[idx])
	}
}

type RangeNode struct {
	Key string
	Val interface{}
}

//for 1 ；2；3 执行顺序 ：初始化 1 ，判断2 成功则进入执行逻辑，然后执行 3 ，再判断 2 以此类推
// 1 -> 2 -> 3 -> 2 -> 3- >2
func (h *hashTable) Range() (hasNext func() bool, next func(), getNode *RangeNode) {
	var (
		idx, count int
		curNode    *HashLinkNode
	)

	getNode = &RangeNode{}

	hasNext = func() bool {
		return count <= h.count
	}

	next = func() {
		count += 1
		if curNode != nil {
			curNode = curNode.next
		}
		for curNode == nil {
			if idx >= len(h.arr) {
				getNode = nil
				return
			}
			curNode = h.arr[idx]
			idx++
		}
		getNode.Key = curNode.Key
		getNode.Val = curNode.Val
	}

	next() //for 初始化
	return
}
