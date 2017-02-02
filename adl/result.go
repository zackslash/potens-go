package adl

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Result is a collection of returned query data
type Result struct {
	Items map[string][]ResultItem
}

// KeyValuePair item
type KeyValuePair struct {
	Key   string
	Value string
}

// ResultItem is a single query result
type ResultItem struct {
	Property string
	Value    string
	Type     PropertyType
}

// Get returns property data
func (e *Entity) Get(property string) string {
	return e.get(property, DataType)
}

// GetWithDefault returns property data or default value if nothing is set
func (e *Entity) GetWithDefault(property, defaultValue string) string {
	res := e.get(property, DataType)
	if res == "" {
		res = defaultValue
	}
	return res
}

// GetCounter returns property data
func (e *Entity) GetCounter(property string) int {
	dat := e.get(property, CounterType)
	val := 0
	if dat != "" {
		val, _ = strconv.Atoi(dat)
	}
	return val
}

// GetSet returns property set data
func (e *Entity) GetSet(property string) []string {
	dat := e.getCollection(property, SetType)
	lstOut := []string{}
	for _, v := range dat {
		lst := []string{}
		json.Unmarshal([]byte(v), &lst)
		lstOut = append(lstOut, lst...)
	}
	return lstOut
}

// GetMeta returns property meta data
func (e *Entity) GetMeta(property string) string {
	return e.get(property, MetaType)
}

// GetList returns property list data
func (e *Entity) GetList(listName string) []KeyValuePair {
	dat := e.getCollection(listName, ListType)
	lstOut := []KeyValuePair{}
	for _, v := range dat {
		lst := []KeyValuePair{}
		json.Unmarshal([]byte(v), &lst)
		lstOut = append(lstOut, lst...)
	}
	return lstOut
}

// Get returns property data for type
func (e *Entity) get(property string, ptype PropertyType) string {
	res := ""
	p := e.result.Items[fmt.Sprintf("%s_%d", property, ptype)]
	if p != nil {
		for _, d := range p {
			if d.Type == ptype {
				res = d.Value
			}
		}
	}
	return res
}

// Get returns property collection data for type
func (e *Entity) getCollection(property string, ptype PropertyType) []string {
	res := []string{}
	p := e.result.Items[fmt.Sprintf("%s_%d", property, ptype)]
	if p != nil {
		for _, d := range p {
			if d.Type == ptype {
				res = append(res, d.Value)
			}
		}
	}
	return res
}
