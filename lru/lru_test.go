/**
 * @Author: fxl
 * @Description:
 * @File:  lru_test.go
 * @Version: 1.0.0
 * @Date: 2022/9/5 13:49
 */
package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}
func TestCache_RemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))
	if _, ok := lru.Get(k1); ok || lru.Len() != 2 {
		t.Fatalf("removeoldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callBack := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := New(int64(10), callBack)
	lru.Add("key1", String("123456"))
	lru.Add("k2", String("k2"))
	lru.Add("k3", String("k3"))
	lru.Add("k4", String("k4"))

	expect := []string{"key1", "k2"}
	t.Log(keys)
	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("call onEvicted faild,expect keys equals to %s", expect)
	}
}
