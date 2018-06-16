package goon

import (
	"encoding/json"
)

type Goon struct {
	str   string
	err   []error
	json  map[string]interface{}
	value interface{}
}

func (g *Goon) checkErr(err error) {
	if err != nil {
		g.err = append(g.err, err)
	}
}

func Parse(s string) (g *Goon) {
	j := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &j)
	g = &Goon{str: s, json: j}
	g.checkErr(err)
	return
}

func (g *Goon) Get(key string) *Goon {
	value := g.json[key]
	if v, ok := value.(map[string]interface{}); ok {
		g.json = v
	}
	g.value = value
	return g
}
