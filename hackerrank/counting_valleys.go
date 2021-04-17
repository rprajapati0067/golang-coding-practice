package main

func countingValleys(n int32, s string) int32 {
	var valleys, altitude int32 = 0, 0
	for i := 0; int32(i) < n; i++ {
		if s[i] == 'U' {
			altitude++
			if altitude == 0 {
				valleys++
			}
		} else {
			altitude--
		}
	}
	return valleys
}
