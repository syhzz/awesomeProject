package test

import (
	"geecache"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	var f geecache.Getter = geecache.GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Errorf("callbacked error")
	}
}
