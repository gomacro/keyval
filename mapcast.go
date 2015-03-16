package main

import (
	"unsafe"
)

// to break the type system
func mvetype(dst, src *interface{}) {
	*(*uintptr)(unsafe.Pointer(dst)) = *(*uintptr)(unsafe.Pointer(src))
}


// this could be a generic function, currently just for int
func keystoslice(iset map[int]struct{}) (out []int) {
	for i, _ := range iset {
		out = append(out, i)
	}
	return out
}

// this could be a generic function, currently just for int
func valstoslice(ibag map[struct{}]int) (out []int) {
	for _, v := range ibag {
		out = append(out, v)
	}
	return out
}

// mam dva interfejsy do toho co tam dam dam vlastne tu intset
// a potom vlastne mvetype na to ints
func main() {
	var ints = map[int]int {7: 5, 4:9, 8:3, 6: 3}
	var iset map[int]struct{}
	var ibag map[struct{}]int

	// two interfaces to move a type info
	var l,r,q interface{}

	l = ints
	r = iset
	q = ibag

	// a convenient way to break the type system
	mvetype(&l, &r)
	iset = l.(map[int]struct{})

	l = ints

	// again for a bag
	mvetype(&l, &q)
	ibag = l.(map[struct{}]int)


	// 
	keys := keystoslice(iset)
	_ = keys

	// 
	vals := valstoslice(ibag)
	_ = vals

	for i, v := range keys {
		_ = i
		print(i)
		print("~ ")
		print(v)
		print("\n")
	}

	for i, v := range vals {
		_ = i
		print(i)
		print(": ")
		print(v)
		print("\n")
	}


}
