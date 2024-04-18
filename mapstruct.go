package main

type Range struct {
	From int
	To   int
}

func AsciiMap(arr []int) map[int]Range {
	res := map[int]Range{}
	for i := range arr {
		if i+1 < len(arr) {
			res[32+i] = Range{From: arr[i+1] - 8, To: arr[i+1] - 1}
		}
	}
	return res
}
