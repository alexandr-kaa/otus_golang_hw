package hw04_lru_cache //nolint:golint,stylecheck

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		// Write me
	})
}

func TestSeveralPass(t *testing.T) {
	t.Run("add 4 time on capacity 3", func(t *testing.T) {
		c := NewCache(3)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		wasInCache = c.Set("ccc", 300)
		require.False(t, wasInCache)

		wasInCache = c.Set("ddd", 400)
		require.False(t, wasInCache)

		val, ok := c.Get("ddd")
		require.Equal(t, 400, val)

		val, ok = c.Get("aaa")
		require.False(t, ok)
	})
}

func TestFrequencySeveralPass(t *testing.T) {
	t.Run("add 4 time on capacity 3", func(t *testing.T) {
		c := NewCache(3)

		for i := 0; i < 50; i++ {
			_ = c.Set("aaa", 100)
		}

		wasInCache := c.Set("bbb", 200)
		require.False(t, wasInCache)

		for i := 0; i < 35; i++ {
			wasInCache = c.Set("ccc", 300)
		}

		for i := 0; i < 30; i++ {
			_ = c.Set("aaa", 1000)
			_, _ = c.Get("bbb")
		}

		wasInCache = c.Set("ddd", 400)
		require.False(t, wasInCache)

		val, ok := c.Get("ddd")
		require.Equal(t, 400, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
	})
}

func TestMockMoveFront(t *testing.T) {
	mockctl := gomock.NewController(t)
	defer mockctl.Finish()

	listmock := NewMockList(mockctl)

	cache := &lruCache{capacity: 3, queue: listmock, items: make(map[Key]*listItem)}

	listmock.EXPECT().Len().Return(0)
	listmock.EXPECT().PushFront(cacheItem{key: "aaa", value: 100})

	wasInCache := cache.Set("aaa", 100)

	require.False(t, wasInCache)

	listmock.EXPECT().Len().Return(1)
	listmock.EXPECT().PushFront(cacheItem{key: "bbb", value: 200})

	wasInCache = cache.Set("bbb", 200)

	//listmock.EXPECT().MoveToFront(cacheItem{})

	listmock.EXPECT().Back()
	itemList := listmock.Back()

	listmock.EXPECT().MoveToFront(itemList)
	item, _ := cache.Get("aaa")

	require.Equal(t, nil, item)
}

func TestNewCacheWithOptions(t *testing.T) {
	listDefault := NewList()
	first := cacheItem{key: "aaa", value: 100}
	listDefault.PushBack(first)
	listDefault.PushBack(cacheItem{key: "bbb", value: 200})

	_, ok := listDefault.Front().Value.(cacheItem)

	require.True(t, ok)

	cache := NewCache(3, SetList(listDefault))
	item, _ := cache.Get("aaa")
	require.Equal(t, first.value, item)
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove if task with asterisk completed

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
