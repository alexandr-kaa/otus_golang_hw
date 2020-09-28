package hw04_lru_cache //nolint:golint,stylecheck

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
		return nil
	}
	return plist.begin
}

func (plist *list) Back() *listItem {
	if plist == nil {
		return nil
	}
	return plist.end
}

func (plist *list) PushFront(v interface{}) *listItem {
	if plist == nil {
		return nil
	}
	item := new(listItem)
	*item = listItem{Prev: nil, Next: plist.begin, Value: v}
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
		return nil
	}
	item := new(listItem)
	*item = listItem{Prev: plist.end, Next: nil, Value: v}
	if plist.count == 0 {
		plist.begin = item
	} else {
		plist.end.Next = item
	}
	plist.end = item
	plist.count++
	return plist.end
}

func (plist *list) Remove(pointer *listItem) {
	if pointer == nil {
		return
	}
	if plist == nil {
		return
	}

	if plist.begin == pointer {
		plist.begin = pointer.Next
	}
	if plist.end == pointer {
		plist.end = pointer.Prev
	}
	if plist.begin != nil && plist.end != nil {
		pointer.Prev.Next = pointer.Next
	}
	plist.count--
}

func (plist *list) MoveToFront(pointer *listItem) {
	if pointer == nil || pointer == plist.Front() {
		return
	}
	pointer.Prev.Next = pointer.Next
	pointer.Next = plist.begin
	plist.begin.Prev = pointer
	pointer.Prev = nil
	plist.begin = pointer
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

	return count == l.Len()-1
}
func NewList() List {
	return &list{}
}
