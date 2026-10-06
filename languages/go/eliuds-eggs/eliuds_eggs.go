package eliudseggs

func EggCountRec(displayValue int, acc int) int {
	if displayValue <= 1 {
		return displayValue + acc
	}
	return EggCountRec(displayValue/2, acc+displayValue%2)
}

func EggCount(displayValue int) int {
	return EggCountRec(displayValue, 0)
}
