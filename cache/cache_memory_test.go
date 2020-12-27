package cache

import (
	"fmt"
	"log"
	"testing"
)

func TestNewCache(t *testing.T) {
	_ = NewCache("memory", 10*1024*1024, func(key string, value []byte) {
		log.Printf("delete key: %s; delete value: %v", key, value)
	})
}

func TestMemoryCache(t *testing.T) {
	mc := NewCache("memory", 10*1024*1024, nil)
	for i := 0; i < 10000; i++ {
		mc.Set(fmt.Sprintf("%d", i), []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"))
		if i%1000 == 0 {
			status := mc.GetStatus()
			t.Log(status)
			t.Log(status.Format())
		}
	}
}
