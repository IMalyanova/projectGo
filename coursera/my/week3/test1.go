package main

import (
	"reflect"
	"testing"
)

func main() {
	
}

var (

Key string
User *User
IsError bool
)

func TestGetUser(t *testing.T) {
	cases := []TestCase {
		TestCase{"ok", &User{ID: 27}, false},
		TestCase{"fail", nil, true},

	}

	for caseNum, item := range cases {
		u, err := GetUser(item.key)

		if item.IsError && err == nil  {
			t.Errorf("[%d] expected error, got nil", caseNum)
		}
		if !item.IsError && err != nil {
			t.Errorf("[%d] unexpected error", caseNum, err)
		}
		if !reflect.DeepEqual(u, item.User) {
			t.Errorf("[%d] wrong results: got %+v, expected %+v",
				caseNum, u, item.User)
		}
	}

}