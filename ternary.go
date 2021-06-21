package lu

// Ternary single line conditional
func Ternary(cond bool, ifTrue interface{}, ifFalse interface{}) interface{} {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

// STernary single line conditional returning a string
func STernary(cond bool, ifTrue string, ifFalse string) string {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

// ITernary single line conditional returning an integer
func ITernary(cond bool, ifTrue int, ifFalse int) int {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}