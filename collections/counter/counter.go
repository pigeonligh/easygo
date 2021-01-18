package counter

type Counter map[string]uint64

func New() Counter {
	return Counter{}
}

func (c Counter) Push(key string) {
	if _, found := c[key]; !found {
		c[key] = 0
	}
	c[key]++
}

func (c Counter) Pushes(key string, times uint64) {
	if _, found := c[key]; !found {
		c[key] = 0
	}
	c[key] += times
}

func (c Counter) Add(d Counter) Counter {
	r := New()
	for k, v := range c {
		r.Pushes(k, v)
	}
	for k, v := range d {
		r.Pushes(k, v)
	}
	return r
}

func (c Counter) Sub(d Counter) Counter {
	r := New()
	for k, v := range c {
		dv, found := d[k]
		if found {
			if v <= dv {
				continue
			}
			v -= dv
		}
		r.Pushes(k, v)
	}
	return r
}

func (c Counter) Max(d Counter) Counter {
	r := New()
	for k, v := range c {
		r.Pushes(k, v)
	}
	for k, v := range d {
		cv, found := c[k]
		if found && cv > v {
			continue
		}
		r[k] = v
	}
	return r
}

func (c Counter) Min(d Counter) Counter {
	r := New()
	for k, v := range c {
		r.Pushes(k, v)
	}
	for k, v := range d {
		cv, found := c[k]
		if found && cv < v {
			continue
		}
		r[k] = v
	}
	return r
}
