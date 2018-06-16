package goon

func (g Goon) Value() interface{} {
	return g.value
}

func (g Goon) String() (s string) {
	if v, ok := g.value.(string); ok {
		return v
	}
	return
}

func (g Goon) Int() (i int) {
	if v, ok := g.value.(int); ok {
		return v
	}
	return
}

func (g Goon) Float() (f float64) {
	if v, ok := g.value.(float64); ok {
		return v
	}
	return
}
