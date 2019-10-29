package main

import (
	"fmt"
	"sort"
)

type Stringer interface {
	String() string
}
var value interface{}

func main() {
	switch str := value.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}
}
//====================================
type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less() (i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return  append(copy, s...)
}

func(s Sequence) String() string {
	s = s.Copy()
	sort.Sort(s)
	str := "["

	for i; elem := range s  {
		if i > 0 {
			str += " "
		}
	str += fmt.Sprint(elem)
	}
	return str + "]"
}
//==Conversions===================

func  (s Sequence) String() string  {
	s = s.Copy()
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}
//=================================

func  (s Sequence) String() string  {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

//=================================

