package fdl

import "github.com/fortifi/proto-go/fdl"

/*
* Set modifiers
**/

// AddListItem to property list data
func (e *Entity) AddSetItem(property, key, value string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        value,
		Type:         SetType,
		MutationMode: int32(fdl.MutationMode_APPEND),
	}
	e.props = append(e.props, p)
	return p
}

// RemoveListItem will remove an item from property list data
func (e *Entity) RemoveSetItem(property, value string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        value,
		Type:         SetType,
		MutationMode: int32(fdl.MutationMode_REMOVE),
	}
	e.props = append(e.props, p)
	return p
}
