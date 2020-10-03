package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())
		l.Remove(l.Back().Prev)
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]

		l.MoveToFront(l.Back()) // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

}
func TestListAppend(t *testing.T) {
	t.Run("list container", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)
		l.PushFront(9.9)
		l.PushBack(`sigma`)

		require.Equal(t, 3, l.Len())

	})
}

func TestDeleteAllList(t *testing.T) {
	t.Run("remove empty container", func(t *testing.T) {
		l := NewList()
		l.Remove(nil)
		require.Equal(t, 0, l.Len())
	})
}

func TestRemoveHeadTail(t *testing.T) {
	t.Run("remove last item from list", func(t *testing.T) {
		l := NewList()
		for i := 0; i < 10; i++ {
			l.PushBack(i)
		}
		require.True(t, Check(l, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))

		require.Equal(t, l.Back().Value, 9)

		l.Remove(l.Back())

		require.Equal(t, l.Back().Value, 8)
	})

	t.Run("remove head item from list", func(t *testing.T) {
		l := NewList()
		for i := 0; i < 10; i++ {
			l.PushBack(i)
		}

		require.True(t, Check(l, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))

		require.Equal(t, l.Front().Value, 0)

		l.Remove(l.Front())

		require.Equal(t, l.Front().Value, 1)

	})

	t.Run("move to front tail and test back", func(t *testing.T) {
		l := NewList()
		for i := 0; i < 10; i++ {
			l.PushBack(i)
		}

		require.True(t, Check(l, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))

		l.MoveToFront(l.Back())

		require.Equal(t, 9, l.Front().Value)

		require.Equal(t, 8, l.Back().Value)

	})
}
