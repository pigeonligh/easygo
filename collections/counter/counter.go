package counter

import (
	"fmt"
	"sort"
)

type Counter map[string]int64

func New() Counter {
	return Counter{}
}

func (c Counter) Push(key string) {
	if _, found := c[key]; !found {
		c[key] = 0
	}
	c[key]++
}

func (c Counter) Pushes(key string, times int64) {
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
		dv, found := d[k]
		if !found {
			continue
		}
		if dv < v {
			v = dv
		}
		r.Pushes(k, v)
	}
	return r
}

func (c Counter) Update(d Counter) Counter {
	r := New()
	for k, v := range c {
		r[k] = v
	}
	for k, v := range d {
		r[k] = v
	}
	return r
}

func (c Counter) Get(key string) int64 {
	ret, found := c[key]
	if !found || ret <= 0 {
		return 0
	}
	return ret
}

func (c Counter) List() []string {
	list := []string{}
	for key, value := range c {
		if value <= 0 {
			continue
		}
		list = append(list, key)
	}
	sort.Strings(list)
	return list
}

func (c Counter) String() string {
	ret := ""
	list := c.List()
	for i, key := range list {
		if i > 0 {
			ret += ", "
		}
		ret += fmt.Sprintf("%v: %v", key, c.Get(key))
	}
	return fmt.Sprintf("Counter[%s]", ret)
}
