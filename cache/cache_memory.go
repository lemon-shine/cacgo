/*******************************************************************************
Method: 内存缓存
Author: Lemine
Langua: golang v1.14
Modify: 2020-12-12
*******************************************************************************/
package cache

import (
	"container/list"
	"sync"
)

//定义内存缓存的数据类型
type MemoryCache struct {
	//主要存放缓存的值-方便数据排序
	list *list.List
	//主要存放缓存的键-方便快速查找
	cache map[string]*list.Element
	//缓存状态
	status *Status
	//移除某条数据时的回调函数
	OnEvicted OnEvictedFunc

	mutex sync.RWMutex
}

func NewMemoryCacche(max int64, onEvicted OnEvictedFunc) *MemoryCache {
	status := &Status{MaxBytes: max}
	return &MemoryCache{
		list:      list.New(),
		cache:     make(map[string]*list.Element),
		status:    status,
		OnEvicted: onEvicted,
	}
}

//添加或更新的一条缓存数据(放队尾)
func (self *MemoryCache) Set(key string, value []byte) bool {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	//更新一条缓存数据
	if elem, ok := self.cache[key]; ok {
		//将缓存数据移动到队尾
		self.list.MoveToBack(elem)
		entry := elem.Value.(*entry)
		entry.value = value
		//更新状态
		self.status.ValBytes += int64(len(value)) - int64(len(entry.value))
		self.status.update()

		//添加一条新的缓存数据
	} else {
		//将新添加的数据放到队尾
		elem := self.list.PushBack(&entry{key, value})
		//将新缓存数据的键与值关联
		self.cache[key] = elem
		//更新状态
		self.status.Count++
		self.status.KeyBytes += int64(len([]byte(key)))
		self.status.ValBytes += int64(len(value))
		self.status.update()
	}
	for self.status.MaxBytes != 0 && self.status.MaxBytes < self.status.NowBytes {
		self.EliminateOne()
	}
	return true
}

//获取一条缓存数据(放队首)
func (self *MemoryCache) Get(key string) []byte {
	self.mutex.RLock()
	defer self.mutex.RUnlock()

	if elem, ok := self.cache[key]; ok {
		//将最近使用的移动到数据链表的队首
		self.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return nil
}

//删除一条缓存数据
func (self *MemoryCache) Del(key string) ([]byte, bool) {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	if elem, ok := self.cache[key]; ok {
		//从list中删除
		entry := self.list.Remove(elem).(*entry)
		//从cache中删除
		delete(self.cache, key)
		//更新状态
		self.status.Count--
		self.status.KeyBytes -= int64(len([]byte(key)))
		self.status.ValBytes -= int64(len(entry.value))
		self.status.update()
		return entry.value, true
	}
	return nil, false
}

func (self *MemoryCache) GetStatus() Status {
	return self.status.Get()
}

//根据淘汰算法移除一条缓存数据
func (self *MemoryCache) EliminateOne() {
	//获取数据链表的队尾节点
	elem := self.list.Back()
	if elem != nil {
		//从数据链表中移除队尾节点
		self.list.Remove(elem)
		entry := elem.Value.(*entry)
		//从数据字典中移除队尾节点
		delete(self.cache, entry.key)
		self.status.Count--
		self.status.KeyBytes -= int64(len([]byte(entry.key)))
		self.status.ValBytes -= int64(len(entry.value))
		self.status.update()
		//处理被移除的队尾节点
		if self.OnEvicted != nil {
			self.OnEvicted(entry.key, entry.value)
		}
	}
}
