package mergesort

// Interface just interface for human-redable code.
type Interface interface{}

// Sort implement `merge sort` algorithm.
// cmp function must return:
//   true if a < b
//   false if something else
// For reverse sort set reverse as true, otherwise set is as false.
func Sort(items []Interface,
	cmp func(a, b Interface) bool, reverse bool) []Interface {
	var num = len(items)

	if num <= 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]Interface, middle)
		right = make([]Interface, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(Sort(left, cmp, reverse), Sort(right, cmp, reverse), cmp, reverse)
}

func merge(left, right []Interface,
	cmp func(a, b Interface) bool, reverse bool) (result []Interface) {
	result = make([]Interface, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if cmp(left[0], right[0]) && !reverse || !cmp(left[0], right[0]) && reverse {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
