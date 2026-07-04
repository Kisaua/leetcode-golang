package main

func angleClock(hour int, minutes int) float64 {
	hourAngle := (float64(hour%12*30) + float64(minutes)/float64(2))

	minutesAngle := float64(minutes * 6)

	result := hourAngle - minutesAngle
	if minutesAngle > hourAngle {
		result = minutesAngle - hourAngle
	}
	if result > 180 {
		result = 360 - result
	}

	return result
}
