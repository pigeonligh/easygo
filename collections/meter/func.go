package meter

func Sum(values ...float64) float64 {
	return Make(values...).Sum()
}

func Average(values ...float64) float64 {
	return Make(values...).Average()
}

func Max(values ...float64) float64 {
	return Make(values...).Max()
}

func Min(values ...float64) float64 {
	return Make(values...).Min()
}
