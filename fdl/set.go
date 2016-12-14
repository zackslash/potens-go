package fdl

import "github.com/fortifi/proto-go/fdl"

/*
* Set modifiers
**/

// AddSetItem adds property set key
func (e *Entity) AddSetItem(property, key string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        key,
		Type:         SetType,
		MutationMode: int32(fdl.MutationMode_WRITE),
	}
	e.props = append(e.props, p)
	return p
}

// RemoveSetItem will remove an item from property list data
func (e *Entity) RemoveSetItem(property, key string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        key,
		Type:         SetType,
		MutationMode: int32(fdl.MutationMode_REMOVE),
	}
	e.props = append(e.props, p)
	return p
}
