package datastructure

import (
	"github.com/spf13/cast"
	"testing"
)

func TestName(t *testing.T) {
	hash := NewHashTable()
	for i := 0; i < 100; i++ {
		hash.Set(cast.ToString(i), i)
	}
	for i := 0; i < 100; i++ {
		t.Log(hash.Get(cast.ToString(i)))
	}

}
