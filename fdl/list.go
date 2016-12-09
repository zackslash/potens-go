package fdl

import "github.com/fortifi/proto-go/fdl"

/*
* List modifiers
**/

// AddListItem to property list data
func (e *Entity) AddListItem(property, value string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        value,
		Type:         ListType,
		MutationMode: int32(fdl.MutationMode_APPEND),
		order:        len(e.props),
	}
	e.props = append(e.props, p)
	return p
}

// AddUniqueListItem will add a list item only if it does not already exist
func (e *Entity) AddUniqueListItem(property, value string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        value,
		Type:         UniqueListType,
		MutationMode: int32(fdl.MutationMode_APPEND),
		order:        len(e.props),
	}
	e.props = append(e.props, p)
	return p
}

// RemoveListItem will remove an item from property list data
func (e *Entity) RemoveListItem(property, value string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        value,
		Type:         ListType,
		MutationMode: int32(fdl.MutationMode_REMOVE),
		order:        len(e.props),
	}
	e.props = append(e.props, p)
	return p
}
