package datastructure

import "testing"

func TestA(t *testing.T) {
	var s = Stack{
		Len: 3,
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	t.Log(s.Push(4))
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
}

func TestLink(t *testing.T) {
	var (
		head = &LinkNode{
			Val: 1,
		}
		err     error
		hasNext bool
		//tail *LinkNode
	)
	h := head
	h, err = h.Append(2)
	h, err = h.Append(4)
	h, err = h.Append(5)
	t.Log(err)
	t.Log(head)
	head, hasNext = head.Next()
	t.Log(head, hasNext)
	head, err = head.Insert(3)
	t.Log(head, err)
	prev, next, err := head.Delete()
	t.Log(prev, next, err)

	//for i := 0; i < 3; i++ {
	//	h, err = h.Append(i + 2)
	//	if err != nil {
	//
	//		break
	//	}
	//
	//}

	//for {
	//	if head != nil {
	//		tail = head
	//		head = head.next
	//
	//		continue
	//	}
	//	break
	//}
	//
	//for {
	//	if tail != nil {
	//		t.Log(tail.Val)
	//		tail = tail.prev
	//		t.Log(tail)
	//		continue
	//	}
	//	break
	//}

}

func TestSingleLink(t *testing.T) {
	var (
		err  error
		head = &SingleLink{
			Val: 1,
		}
		h = head
	)

	head, err = head.Append(2)
	t.Log(head, err)
	head, err = head.Append(3)
	t.Log(head, err)
	t.Log("------------\n")
	head = h.Reverse()
	t.Log(head)
	head = head.next
	t.Log(head)
	head = head.next
	t.Log(head)
	head = head.next

}

func TestSingleLink2(t *testing.T) {
	var (
		err  error
		head = &SingleLink{
			Val: 1,
		}
		h = head
	)

	head, err = head.Append(2)
	t.Log(head, err)
	head, err = head.Append(3)
	t.Log(head, err)
	t.Log("------------\n")
	h = h.Reverse2(nil)
	t.Log("------------\n")
	t.Log(h)
	h = h.next
	t.Log(h)
	h = h.next
	t.Log(h)
	h = h.next

}
