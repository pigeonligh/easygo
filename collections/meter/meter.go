package meter

type Meter struct {
	values []float64
	sum    float64
	count  int

	max float64
	min float64
}

func New() *Meter {
	return &Meter{
		values: []float64{},
		sum:    0,
		count:  0,
		max:    0,
		min:    0,
	}
}

func (m *Meter) Adds(values ...float64) {
	for _, v := range values {
		m.Add(v)
	}
}

func (m *Meter) AddMeter(meter *Meter) {
	for _, v := range meter.values {
		m.Add(v)
	}
}

func (m *Meter) Add(v float64) {
	m.values = append(m.values, v)
	m.sum += v
	m.count++

	if m.count == 1 {
		m.max = v
		m.min = v
	} else {
		if v > m.max {
			m.max = v
		}
		if v < m.min {
			m.min = v
		}
	}
}

func (m *Meter) Sum() float64 {
	return m.sum
}

func (m *Meter) Count() int {
	return m.count
}

func (m *Meter) Average() float64 {
	return m.sum / float64(m.count)
}

func (m *Meter) Max() float64 {
	return m.max
}

func (m *Meter) Min() float64 {
	return m.min
}
