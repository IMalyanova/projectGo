package main

func main() {

}

func unhex(c byte) byte{
	switch{
	case '0' <= c && c<= '9':
		return c - '0'
	case 'a' <= c && c <= 'F':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
//==========================

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}
//==========================

Loop:
	for n := 0; n < len(src); n += size {
		switch {
		case src[n] < sizeOne:
			if validateOnly {
				break
			}
			size = 1
			update(src[n])

		case src[n] < sizeTwo:
			if n+1 > len(src){
				err = errShortInput
				break Loop
			}
			if validateOnly{
				break
			}
			size = 2
			update (src[n] + src[n+1] << shift)
		}
	}
	//=========================================

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b) ; i++  {
		switch {
		case a[i] > b[i]:
			return  1
		case a[i] < b[i]:
			return  -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) > len(b):
		return -1
	}
	return 0
}
	//================================================

var t interface{}
t = functionOfSomeType()
switch t := t.type){
default:
	fmt.Printf("unexpected type %T\n", t)
case bool:
	fmt.PrintF("boolean %t\n", t)
case int:
	fmt.Printf("integer %d\n", t)
case *bool:
	fmt.Printf("pointer to integer %d\n", *t)
case *int:
	fmt.Printf("pointer to integer %d\n", *t)
}
//================================================
	