package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {
	lastResult := 2
	newResult := 0
	for i := 0; i < len(arr)-1; i++ {
		if f(arr[i], arr[i+1]) == 1 {
			newResult = 1
		} else if f(arr[i], arr[i+1]) == -1 {
			newResult = -1
		}
		if lastResult == 2 {
			lastResult = newResult
		} else if lastResult == newResult {
			return true
		} else if lastResult != newResult {
			return false
		}
	}
	return false
}

func StrCompare(a, b string) int {
	lenA := len(a)
	lenB := len(b)

	for i := 0; i < lenA && i < lenB; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if lenA < lenB {
		return -1
	} else if lenA > lenB {
		return 1
	}
	return 0
}
