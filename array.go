package lu

import (
	"fmt"
)

// Array type of array to add class functions to
type Array []interface{}

// a converts Array to type []interface{}
func (A Array) a() []interface{} {
	var a []interface{}
	for _, v := range A {
		a = append(a, v)
	}
	return a
}

// Froma converts array to Array
func Froma(a []interface{}) Array {
	A := Array{}
	for _, v := range a {
		A.Append(v)
	}
	return A
}

func (A Array) String() string {
	return fmt.Sprintf("[%s]", A.Join(", "))
}

// ForEach loops through the current array and runs the provided function
func (A Array) ForEach(f func(i int, v interface{})) {
	for i, v := range A {
		f(i, v)
	}
}

func (A Array) Len() int {
	return len(A)
}

func (A Array) Join(sep string) String {
	var s string
	for i, v := range A {
		s += fmt.Sprintf("%v", v)
		if i != (A.Len() - 1) { s += sep }
	}
	return String(s)
}

func (A Array) Append(elems ...interface{}) Array {
	return append(A, elems...)
}

func (A Array) Filter(f func(i int, elem interface{}) bool) Array {
	arr := Array{}
	for i, v := range A {
		if f(i, v) {
			arr.Append(v)
		}
	}
	return arr
}

func (A Array) Find(f func(i int, elem interface{}) bool) interface{} {
	var val interface{}
	for i, v:= range A {
		if f(i, v) {
			val = v
			break
		}
	}
	return val
}

func (A Array) Contains(elem interface{}) bool {
	b := false
	for _, v := range A {
		if v == elem {
			b = true
			break
		}
	}
	return b
}

func (A Array) IndexOf(elem interface{}) int {
	var index int
	for i, v := range A {
		if v == elem {
			index = i
			break
		}
	}
	return index
}

func (A Array) First() interface{} {
	return A[0]
}

func (A Array) Last() interface{} {
	return A[A.Len() - 1]
}

func (A Array) Remove(index int) Array {
	arr := Array{}
	i := 0
	for i < A.Len() {
		if i != index {
			arr.Append(A[i])
			i++
		}
	}
	return arr
}

// Sort sorts an array
func (A Array) Sort(f func(a interface{}, b interface{}) int) Array {
	for i, elemA := range A {
		if i != 0 {
			i2 := 0
			for f(elemA, A[i2]) > 0 && i2 < A.Len() {
				i2++
			}

			var i3, length int
			if i > i2 {
				length = i
				i3 = i2
			} else {
				length = i2
				i3 = i
			}
			for ; i3 < length; i3++ {
				A.Swap(i3, length)
			}
		}
	}
	return A
}

// Swap switches these two elements around
func (A Array) Swap(i1, i2 int) Array {
	elem1 := A[i1]
	elem2 := A[i2]
	A[i1] = elem2
	A[i2] = elem1
	return A
}