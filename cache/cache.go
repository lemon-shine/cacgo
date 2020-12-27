/*******************************************************************************
Method: 缓存的基础数据类型
Author: Lemine
Langua: golang v1.14
Modify: 2020-12-12
*******************************************************************************/
package cache

//定义缓存接口类型
type Cache interface {
	//获取一条键值对数据
	Get(string) []byte
	//新增或修改一条键值对数据
	Set(string, []byte) bool
	//删除一条键值对数据
	Del(string) ([]byte, bool)
	//获取缓存当前的状态
	GetStatus() Status
	//淘汰一条键值对数据
	//NOTE:使用LRU算法
	EliminateOne()
}

//定义删除一条键值对数据时的回调函数
type OnEvictedFunc func(string, []byte)

func NewCache(ctype string, maxBytes int64, onEvicted OnEvictedFunc) Cache {
	var cache Cache
	switch ctype {
	case "memory":
		cache = NewMemoryCacche(maxBytes, onEvicted)
	default:
		panic("Illegal cache type:" + ctype)
	}
	return cache
}
