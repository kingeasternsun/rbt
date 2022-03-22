package rbt

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedBlackTree_Insert(t *testing.T) {

	tr := &RedBlackTree{
		Root: nil,
	}

	for i := 0; i < 10000; i++ {
		t.Run("Insert"+strconv.Itoa(i), func(t *testing.T) {
			tr.Insert(Key{Key: i}, i)
			_, ret := tr.Check()
			assert.Equal(t, ret, true)
		})
	}

	for i := 10000; i >= 0; i-- {
		t.Run("Remove"+strconv.Itoa(i), func(t *testing.T) {
			tr.Remove(Key{Key: i})
			_, ret := tr.Check()
			assert.Equal(t, ret, true)
		})
	}
}
