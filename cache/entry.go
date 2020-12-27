/*******************************************************************************
Method: 关联list中的值和map中的键，存放在list中，删除list节点时可快速找到cache中对应的数据
Author: Lemine
Langua: golang v1.14
Modify: 2020-12-12
*******************************************************************************/
package cache

type entry struct {
	key   string //缓存数据的键
	value []byte //缓存数据的值
}
