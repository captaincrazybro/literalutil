package lu

import "fmt"

type SArray []String

func (SA SArray) String() string {
	return fmt.Sprintf("[%s]", SA.Join(", "))
}

// sa converts SArray to type []string
func (SA SArray) sa() []string {
	var sa []string
	for _, v := range SA {
		sa = append(sa, v.s())
	}
	return sa
}

// Fromsa converts []string to type SArray
func Fromsa(sa []string) SArray {
	SA := SArray{}
	for _, v := range sa {
		SA.Append(String(v))
	}
	return SA
}

// ForEach loops through the current array and runs the provided function
func (SA SArray) ForEach(f func(i int, v String)) {
	for i, v := range SA {
		f(i, v)
	}
}

func (SA SArray) Len() int {
	return len(SA)
}

func (SA SArray) Join(sep String) String {
	var s String
	for i, v := range SA {
		s += String(fmt.Sprintf("%v", v))
		if i != (SA.Len() - 1) { s += sep }
	}
	return s
}

func (SA SArray) Append(elems ...String) SArray {
	return append(SA, elems...)
}

func (SA SArray) Filter(f func(i int, elem String) bool) SArray {
	arr := SArray{}
	for i, v := range SA {
		if f(i, v) {
			arr.Append(v)
		}
	}
	return arr
}

func (SA SArray) Find(f func(i int, elem String) bool) String {
	var val String
	for i, v:= range SA {
		if f(i, v) {
			val = v
			break
		}
	}
	return val
}

func (SA SArray) Contains(elem String) bool {
	b := false
	for _, v := range SA {
		if v == elem {
			b = true
			break
		}
	}
	return b
}

func (SA SArray) IndexOf(elem String) int {
	var index int
	for i, v := range SA {
		if v == elem {
			index = i
			break
		}
	}
	return index
}

func (SA SArray) First() String {
	return SA[0]
}

func (SA SArray) Last() String {
	return SA[SA.Len() - 1]
}

func (SA SArray) Remove(index int) SArray {
	arr := SArray{}
	i := 0
	for i < SA.Len() {
		if i != index {
			arr.Append(SA[i])
			i++
		}
	}
	return arr
}

// Sort sorts an array
func (SA SArray) Sort(f func(a String, b String) int) SArray {
	for i, elemA := range SA {
		if i != 0 {
			i2 := 0
			for f(elemA, SA[i2]) > 0 && i2 < SA.Len() {
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
				SA.Swap(i3, length)
			}
		}
	}
	return SA
}

// SortAlpha sorts an array alphabetically
func (SA SArray) SortAlpha() SArray {
	return SA.Sort(func(a, b String) int {
		sa := fmt.Sprintf("%s", a)
		sb := fmt.Sprintf("%s", b)
		if sa > sb {
			return 1
		} else if sb > sa {
			return -1
		} else {
			return 0
		}
	})
}

// Swap switches these two elements around
func (SA SArray) Swap(i1, i2 int) SArray {
	elem1 := SA[i1]
	elem2 := SA[i2]
	SA[i1] = elem2
	SA[i2] = elem1
	return SA
}