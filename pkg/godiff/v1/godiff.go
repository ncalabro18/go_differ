package godiff


func Backtrack(f1, f2 []byte) [][]int {

	//start at the end indices and backtrack
	x := len(f1)
	y := len(f2)

	var kprev int
	retArray := [][]int{}

	t := Trace(f1, f2)
	for d := len(t) - 1; d >= 0; d-- {
		v := t[d]
		k := x - y

		if k == -d || (k != d && v[k-1] < v[k+1]) {
			kprev = k + 1
		} else {
			kprev = k - 1
		}

		xprev := v[kprev]
		yprev := xprev - kprev
		for x > xprev && y > yprev {
			retArray = append(retArray, []int{x - 1, y - 1, x, y})
			x--
			y--
		}

		if d > 0 {
			retArray = append(retArray, []int{xprev, yprev, x, y})
		}
		x = xprev
		y = yprev
	}

	return retArray
}

func Trace(f1, f2 []byte) [][]int {
	var x int
	l1 := len(f1)
	l2 := len(f2)
	max := l1 + l2
	trace := [][]int{}
	v := make([]int, (2*max)+1)
	//v := []int{}

	for d := 0; d <= max; d++ {
		trace = append(trace, v)
		for k := -d; k <= d; k += 2 {
			vkp1 := revValue(v, k+1)
			vkm1 := revValue(v, k-1)
			if k == -d || (k != d && vkm1 < vkp1) {
			//if k == -d || (k != d && v[k-1] < v[k+1]) {
				//x = v[k+1]

				//down
				x = vkp1
			} else {
				//x = v[k-1] + 1

				//right
				x = vkm1 + 1
			}
			y := x - k

			for x < l1 && y < l2 && f1[x] == f2[y] {
				x++
				y++
			}

			//v[k] = x
			revAssignment(v, k, x)
			if x >= l1 && y >= l2 {
				return trace
			}
		}
	}

	return trace
}

func revAssignment(array []int, index int, value int) {
	if index < 0 {
		array[len(array)+index] = value
	} else {
		array[index] = value
	}
}

func revValue(array []int, index int) int {
	if index < 0 {
		return array[len(array)+index]
	} else {
		return array[index]
	}
}

