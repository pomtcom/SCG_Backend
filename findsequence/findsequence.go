package findsequence

func FindSequnceDigit(index int) int{
	n1 := 3
	n2 := 5

	if index == 0 {
		return n1
	}else if index == 1{
		return n2
	}

	nTarget := 0
	for i := 3; i <= index; i++ {
		nTarget = n1 + n2 + 1
		n1 = n2
		n2 = nTarget
	}
	return nTarget
}