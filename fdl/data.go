package fdl

import (
	"github.com/fortifi/proto-go/fdl"
)

/*
* Data modifiers
**/

// Write to property data
func (e *Entity) Write(property, value string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        value,
		Type:         DataType,
		MutationMode: int32(fdl.MutationMode_WRITE),
		order:        len(e.props),
	}
	e.props = append(e.props, p)
	return p
}

// Append to property data
func (e *Entity) Append(property, value string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        value,
		Type:         DataType,
		MutationMode: int32(fdl.MutationMode_APPEND),
		order:        len(e.props),
	}
	e.props = append(e.props, p)
	return p
}

// Delete property data
func (e *Entity) Delete(property string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        "",
		Type:         DataType,
		MutationMode: int32(fdl.MutationMode_DELETE),
		order:        len(e.props),
	}
	e.props = append(e.props, p)
	return p
}
