package enum

func ReduceAMulInt(argname_1 []int, argname_2 func(int, int) int, argname_3 int) int {

	lenSlice := len(argname_1)
	switch lenSlice {
	case 0:
		return 0
	case 1:
		return argname_1[1]
	}
	out := argname_2(argname_3, argname_1[0])
	next := argname_1[1]
	for i := 1; i < lenSlice; i++ {
		next = argname_1[i]
		out = argname_2(out, next)
	}
	return out

}
