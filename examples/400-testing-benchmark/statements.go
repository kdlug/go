package statements

func If(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func Switch(i int) bool {
	switch {
	case i == 1:
		return true
	default:
		return false
	}
}

func SwitchConst(i int) bool {
	switch i {
	case 1:
		return true
	default:
		return false
	}
}
