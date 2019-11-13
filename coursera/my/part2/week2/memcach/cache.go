package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
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
	
	item := &CacheItem{}
	err = json.Unmarshal(itemRaw.Value, &item)
	if err != nil {
		return err
	}
	
	tagsValid, err := tc.isTagsValid(item.Tags)
	if err != nil {
		return fmt.Errorf("isTagsValid error %s", err)
	}
	
	if tagsValid {
		err = json.Unmarshal(item.Data, &in)
		return err
	}
	
	return tc.rebuild(mkey, ttl, in, rebuildCb)
}



func (tc *TCache) isTagsValid(itemTags map[string]int) (bool, error) {
	tags := make([]string, 0, len(itemTags))
	
	for tagKey := range itemTags {
		tags = append(tags, tagKey)
	}
	
	curr, err := tc.GetMulti(tags)
	
	if err != nil {
		return false, err
	}
	
	currentTagsMap := make(map[string]int, len(curr))
	
	for tagKey, tagItem := range  curr {
		i, err := strconv.Atoi(string(tagItem.Value))
		
		if err != nil {
			return false, err
		}

		currentTagsMap[tagKey] = i
	}
	return reflect.DeepEqual(itemTags, currentTagsMap), nil	
}



func (tc *TCache) rebuild(
	mkey string,
	ttl int32,
	in interface{},
	rebuildCb RebuildFunc, 
) error {
	tc.lockRebuild(mkey)
	
	result, tags, err := rebuildCb()
	
	// ожидаем и возвращаем одинаковые типы
	if reflect.TypeOf(result) != reflect.TypeOf(in) {
		return fmt.Errorf(
			"data type mismatch, expected %s, got %s", reflect.TypeOf(in), 
			reflect.TypeOf(result),
			)
	}
	
	currTags, err := tc.getCurrentItemTags(tags, ttl)
	
	if err != nil {
		return err
	}
	
	cacheData := CacheItemStore{result, currTags}
	rawData, err := json.Marshal(cacheData)
	if err != nil {
		return err
	}
	
	err = tc.Set(&memcache.Item {
		Key:		mkey,
		Value:		rawData,
		Expiration: int32(ttl),
	})
	
	inVal := reflect.ValueOf(in)
	resultVal := reflect.ValueOf(result)
	rv := reflect.Indirect(inVal)
	rvpresult := reflect.Indirect(resultVal)
	rv.Set(rvpresult)

	return nil
}

func (tc *TCache) checkLock(mkey string) error {
	
	for i := 0; i < 4; i++ {
		_, err := tc.Get("lock_" + mkey)
		
		if err == memcashe.ErrCacheMiss {
			return nil
		}
		if err != nill {
			return err
		}
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}





func main() {
	
}
