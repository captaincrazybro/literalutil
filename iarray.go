package lu

import "fmt"

type IArray []int

func (IA IArray) String() string {
	return fmt.Sprintf("[%s]", IA.Join(", "))
}

// Toia converts IArray to type []int
func (IA IArray) Toia() []int {
	var ia []int
	for _, v := range IA {
		ia = append(ia, v)
	}
	return ia
}

// Fromia converts []string to type SArray
func Fromia(sa []int) IArray {
	SA := IArray{}
	for _, v := range sa {
		SA.Append(v)
	}
	return SA
}

// ForEach loops through the current array and runs the provided function
func (IA IArray) ForEach(f func(i int, v int)) {
	for i, v := range IA {
		f(i, v)
	}
}

func (IA IArray) Len() int {
	return len(IA)
}

func (IA IArray) Join(sep string) String {
	var s string
	for i, v := range IA {
		s += fmt.Sprintf("%v", v)
		if i != (IA.Len() - 1) { s += sep }
	}
	return String(s)
}

func (IA IArray) Append(elems ...int) IArray {
	return append(IA, elems...)
}

func (IA IArray) Filter(f func(i int, elem int) bool) IArray	 {
	arr := IArray{}
	for i, v := range IA {
		if f(i, v) {
			arr.Append(v)
		}
	}
	return arr
}

func (IA IArray) Find(f func(i int, elem int) bool) int {
	var val int
	for i, v:= range IA {
		if f(i, v) {
			val = v
			break
		}
	}
	return val
}

func (IA IArray) Contains(elem int) bool {
	b := false
	for _, v := range IA {
		if v == elem {
			b = true
			break
		}
	}
	return b
}

func (IA IArray) IndexOf(elem int) int {
	var index int
	for i, v := range IA {
		if v == elem {
			index = i
			break
		}
	}
	return index
}

func (IA IArray) First() int {
	return IA[0]
}

func (IA IArray) Last() int {
	return IA[IA.Len() - 1]
}

func (IA IArray) Remove(index int) IArray {
	arr := IArray{}
	i := 0
	for i < IA.Len() {
		if i != index {
			arr.Append(IA[i])
			i++
		}
	}
	return arr
}

// Sort sorts an array
func (IA IArray) Sort(f func(a int, b int) int) IArray {
	for i, elemA := range IA {
		if i != 0 {
			i2 := 0
			for f(elemA, IA[i2]) > 0 && i2 < IA.Len() {
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
				IA.Swap(i3, length)
			}
		}
	}
	return IA
}

// Swap switches these two elements around
func (IA IArray) Swap(i1, i2 int) IArray {
	elem1 := IA[i1]
	elem2 := IA[i2]
	IA[i1] = elem2
	IA[i2] = elem1
	return IA
}