package datastructure

import (
	"github.com/spf13/cast"
	"testing"
)

func TestName(t *testing.T) {
	hash := NewHashTable()
	for i := 0; i < 2; i++ {
		hash.Set(cast.ToString(i), i)
	}
	//for i := 0; i < 100; i++ {
	//	t.Log(hash.Get(cast.ToString(i)))
	//}

	//hash.Set("wfy", "wfy__")
	//hash.Set("lll", "lll__")
	//t.Log(hash.Get("wfy"))
	//t.Log(hash.Get("lll"))
	//var c int
	//for _, node := range hash.arr {
	//	for node != nil {
	//		c ++
	//		//t.Log(node.Val)
	//		node = node.next
	//	}
	//}
	//t.Log(c)
	//for i := 0; i < 100; i++ {
	//	hash.Delete(cast.ToString(i))
	//}

	//for i := 0; i < 100; i++ {
	//	t.Log(hash.Get(cast.ToString(i)))
	//}

	for hasNext, next, getNode := hash.Range(); hasNext(); next() {
		t.Log("get Node", getNode)
	}
}
