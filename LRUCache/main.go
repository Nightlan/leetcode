package main

import (
	"container/list"
	"fmt"
)

type LRUData struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	findMap  map[int]*list.Element
	seqList  *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		findMap:  make(map[int]*list.Element, capacity),
		seqList:  list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	ele, ok := l.findMap[key]
	if !ok {
		return -1
	}
	return l.update(ele)
}

func (l *LRUCache) update(ele *list.Element) int {
	l.seqList.Remove(ele)
	data := ele.Value.(*LRUData)
	newEle := l.seqList.PushFront(data)
	l.findMap[data.key] = newEle
	return data.value
}

func (l *LRUCache) Put(key int, value int) {
	existEle, ok := l.findMap[key]
	if ok {
		data := existEle.Value.(*LRUData)
		data.value = value
		l.update(existEle)
		return
	}
	if len(l.findMap) == l.capacity {
		delEle := l.seqList.Back()
		l.seqList.Remove(delEle)
		delData := delEle.Value.(*LRUData)
		delete(l.findMap, delData.key)
	}
	newData := &LRUData{
		key:   key,
		value: value,
	}
	newEle := l.seqList.PushFront(newData)
	l.findMap[key] = newEle
}

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 返回  1
	cache.Put(3, 3)           // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4)) // 返回  4

}
