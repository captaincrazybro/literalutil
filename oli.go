package lu

// one line if

// If function for a one line if
func If(cond bool, then func()) {
	if cond {
		then()
	}
}

// Ifelse function for a one line if else statement
func Ifelse(cond bool, then func(), other func()) {
	if cond {
		then()
	} else {
		other()
	}
}