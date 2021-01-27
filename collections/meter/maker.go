package meter

func Make(values ...float64) *Meter {
	m := New()
	m.Adds(values...)
	return m
}

func MakeByInt(values ...int) *Meter {
	m := New()
	for _, v := range values {
		m.Add(float64(v))
	}
	return m
}

func MakeByInt64(values ...int64) *Meter {
	m := New()
	for _, v := range values {
		m.Add(float64(v))
	}
	return m
}
