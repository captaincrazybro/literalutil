package lu

import (
	"fmt"
)

// Array type of array to add class functions to
type Array []interface{}

// ToA returns the raw array
func (a Array) ToA() []interface{} {
	return a
}

// ForEach loops through the current array and runs the provided function
func (a Array) ForEach(f func(i int, v interface{})) {
	for i, v := range a {
		f(i, v)
	}
}

func (a Array) Len() int {
	return len(a)
}

func (a Array) Join(sep string) String {
	var s string
	for i, v := range a {
		s += fmt.Sprintf("%v", v)
		if i != (a.Len() - 1) { s += sep }
	}
	return String(s)
}

func (a Array) Append(elems ...interface{}) Array {
	return append(a, elems)
}

func (a Array) Filter(f func(i int, elem interface{}) bool) Array {
	arr := Array{}
	for i, v := range a {
		if f(i, v) {
			arr.Append(v)
		}
	}
	return arr
}

func (a Array) Find(f func(i int, elem interface{}) bool) interface{} {
	var val interface{}
	for i, v:= range a {
		if f(i, v) {
			val = v
			break
		}
	}
	return val
}

func (a Array) Contains(elem interface{}) bool {
	b := false
	for _, v := range a {
		if v == elem {
			b = true
			break
		}
	}
	return b
}

func (a Array) IndexOf(elem interface{}) int {
	var index int
	for i, v := range a {
		if v == elem {
			index = i
			break
		}
	}
	return index
}

func (a Array) First() interface{} {
	return a[0]
}

func (a Array) Last() interface{} {
	return a[a.Len() - 1]
}

func (a Array) Remove(index int) Array {
	arr := Array{}
	i := 0
	for i < a.Len() {
		if i != index {
			arr.Append(a[i])
			i++
		}
	}
	return arr
}

// Sort not finished yet
func (arr Array) Sort(f func(a interface{}, b interface{}) int) Array {
	/*for i, v := range arr {

	}*/
	return arr
}

// Swap switches these two elements around
func (a Array) Swap(i1, i2 int) Array {
	elem1 := a[i1]
	elem2 := a[i2]
	a[i1] = elem2
	a[i2] = elem1
	return a
}