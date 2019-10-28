package main

func main() {
	//for
	for init; condition; post{}

	//while
	for condition{}

	//for ; ;
	for{}
	//================================
	summ: = 0
	for i: = 0; i < 10; i ++ {
		sum + = i
	}
	//	============================
	for key, value := range oldMap {
		newMap[key] = value
	}
	//=================================
	for key := range m {
		if key.expired(){
			delete(m, key)
		}
	}
	//	=========================

	for _, value := range array {
		sum += value
	}
	//	===========================
	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a [j] = a[j], a[i]
	}
	//	===========================


}
