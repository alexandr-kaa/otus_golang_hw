package hw04_lru_cache //nolint:golint,stylecheck
//go:generate mockgen -destination=list_mock.go -package=hw04_lru_cache . List
type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(pointer *listItem)
	MoveToFront(pointer *listItem)
}

type listItem struct {
	Value interface{}

	Next *listItem
	Prev *listItem
	// Place your code here
}

type list struct {
	count int
	begin *listItem
	end   *listItem
}

func (plist *list) Len() int {
	if plist == nil {
		return -1
	}
	return plist.count
}

func (plist *list) Front() *listItem {
	if plist == nil {
		plist = &list{}
	}
	return plist.begin
}

func (plist *list) Back() *listItem {
	if plist == nil {
		plist = &list{}
	}
	return plist.end
}

func (plist *list) PushFront(v interface{}) *listItem {
	if plist == nil {
		plist = &list{}
	}
	item := &listItem{Prev: nil, Next: plist.begin, Value: v}
	if plist.count == 0 {
		plist.end = item
	} else {
		plist.begin.Prev = item
	}
	plist.begin = item
	plist.count++
	return plist.begin
}

func (plist *list) PushBack(v interface{}) *listItem {
	if plist == nil {
		plist = &list{}
	}
	item := &listItem{Prev: plist.end, Next: nil, Value: v}
	if plist.count == 0 {
		plist.begin = item
	} else {
		plist.end.Next = item
	}
	plist.end = item
	plist.count++
	return plist.end
}

func (plist *list) Remove(plistItem *listItem) {
	if plistItem == nil {
		return
	}
	if plist == nil {
		return
	}

	if plist.begin == plistItem {
		plist.begin = plistItem.Next
		plist.begin.Prev = nil
		plist.count--
		return
	}
	if plist.end == plistItem {
		plist.end = plistItem.Prev
		plist.end.Next = nil
		plist.count--
		return
	}
	if plist.begin != nil && plist.end != nil {
		plistItem.Prev.Next = plistItem.Next
	}
	plist.count--
}

func (plist *list) MoveToFront(plistItem *listItem) {
	if plistItem == nil || plistItem == plist.Front() {
		return
	}
	if plist.end == plistItem {
		plist.end = plist.end.Prev
	}

	plistItem.Prev.Next = plistItem.Next
	plistItem.Next = plist.begin
	plist.begin.Prev = plistItem
	plistItem.Prev = nil
	plist.begin = plistItem
}
func Check(l List, r []int) bool {
	if l.Len() != len(r) {
		return false
	}
	count := 0
	for i := l.Front(); i != nil; i = i.Next {
		if i.Value.(int) != r[count] {
			return false
		}
		count++
	}

	return count == l.Len()
}
func NewList() List {
	return &list{}
}
