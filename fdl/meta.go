package fdl

import (
	"github.com/fortifi/proto-go/fdl"
)

/*
* Data modifiers
**/

// WriteMeta will write to property meta data
func (e *Entity) WriteMeta(property, value string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        value,
		Type:         MetaType,
		MutationMode: int32(fdl.MutationMode_WRITE),
	}
	e.props = append(e.props, p)
	return p
}

// DeleteMeta will delete property meta data
func (e *Entity) DeleteMeta(property, value string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        value,
		Type:         MetaType,
		MutationMode: int32(fdl.MutationMode_DELETE),
	}
	e.props = append(e.props, p)
	return p
}
