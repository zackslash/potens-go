package adl

import (
	"github.com/cubex/proto-go/adl"
)

/*
* Meta modifiers
**/

// WriteMeta will write to property meta data
func (e *Entity) WriteMeta(property, value string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        value,
		Type:         MetaType,
		MutationMode: int32(adl.MutationMode_WRITE),
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
		MutationMode: int32(adl.MutationMode_DELETE),
	}
	e.props = append(e.props, p)
	return p
}
