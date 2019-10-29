package main

import "sort"

func main() {
	
}

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
//=======================================