func largestAltitude(gain []int) int {
	temp,h := 0,0
	for i, _ := range gain {
		h = h + gain[i]
		if h > temp {
			temp = h
		}
	}
	return temp
}
