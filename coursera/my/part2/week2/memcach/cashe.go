package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type CacheItem struct {
	Data json.RawMessage
	Tags map[string]int
}

type CacheItemStore struct {
	Data interface{}
	Tags map[string]int
}

type RebuildFunc func() (interface{}, []string, error)

type TCache struct {
	*memcache.Client
}



func (tc *TCache) TGet (
	mkey string,
	ttl int32,
	in interface{},
	rebuildCb RebuildFunc,
) (err error) {
	inKind := reflect.ValueOf(in).Kind()
	if inKind != reflect.Ptr {
		return fmt.Errorf("in must be ptr, got %s", inKind)
	}

	tc.checkLock(mkey)

	itemRaw, err := tc.Get(mkey)

	if err == memcache.ErrCacheMiss {
		fmt.Println("Record not found in memcache")
		return tc.rebuild(mkey, ttl, in, rebuildCb)
	} else if err != nil {
		return err
	}
}





func main() {
	
}
