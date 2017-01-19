package adl

import (
	"github.com/cubex/proto-go/adl"
)

/*
* List modifiers
**/

// AddListItem adds an item to given entity list
func (e *Entity) AddListItem(listName, key, value string) PropertyItem {
	p := PropertyItem{
		Property:     listName,
		Key:          key,
		Value:        value,
		Type:         ListType,
		MutationMode: int32(adl.MutationMode_WRITE),
	}
	e.props = append(e.props, p)
	return p
}

// RemoveListItem removes an existing item to given entity list
func (e *Entity) RemoveListItem(listName, key string) PropertyItem {
	p := PropertyItem{
		Property:     listName,
		Key:          key,
		Type:         ListType,
		MutationMode: int32(adl.MutationMode_REMOVE),
	}
	e.props = append(e.props, p)
	return p
}
