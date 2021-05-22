package datastructure

import (
	"errors"
	"fmt"
)

type LinkNode struct {
	prev, next *LinkNode
	Val        interface{}
}

func (l *LinkNode) Next() (nextNode *LinkNode, hasNext bool) {
	if l == nil {
		return
	}

	return l.next, l.next != nil
}
func (l *LinkNode) Append(val interface{}) (newLink *LinkNode, err error) {
	if l.next != nil {
		return newLink, errors.New("next link not nil")
	}
	newLink = &LinkNode{
		prev: l,
		Val:  val,
	}
	l.next = newLink
	return
}

func (l *LinkNode) Insert(val interface{}) (newLink *LinkNode, err error) {
	if l == nil {
		return newLink, errors.New("origin link is nil")
	}
	if l.next == nil {
		return l.Append(val)
	}
	newLink = &LinkNode{
		prev: l,
		next: l.next,
		Val:  val,
	}
	l.next.prev = newLink
	l.next = newLink
	return
}

func (l *LinkNode) Delete() (prev, next *LinkNode, err error) {
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

type SingleLink struct {
	next *SingleLink
	Val  interface{}
}

func (l *SingleLink) Append(val interface{}) (newLink *SingleLink, err error) {
	if l.next != nil {
		return newLink, errors.New("next link not nil")
	}
	l.next = &SingleLink{
		Val: val,
	}
	return l.next, nil
}

func (l *SingleLink) Reverse() *SingleLink {
	if l == nil {
		return l
	}
	var (
		a, b, c *SingleLink
	)

	a = l
	b = a.next
	if b.next != nil {
		c = b.next
	}
	fmt.Println(a,b,c)
	// a -> b -> c

	//      t- > c
	// a <- -> b    c
	l.next = nil
	for b != nil {
		b.next = a
		fmt.Println(a,b,c)
		a = b
		b = c
		if c != nil {
			c = c.next
		}
	}
	return a
}
