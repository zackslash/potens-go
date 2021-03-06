package adl

import (
	"github.com/cubex/proto-go/adl"
)

/*
* Counter modifiers
**/

// IncrementCounter will increment properties counter
func (e *Entity) IncrementCounter(property string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        "1",
		Type:         CounterType,
		MutationMode: int32(adl.MutationMode_APPEND),
	}
	e.props = append(e.props, p)
	return p
}

// DecrementCounter will decrement the properties counter
func (e *Entity) DecrementCounter(property string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        "1",
		Type:         CounterType,
		MutationMode: int32(adl.MutationMode_REMOVE),
	}
	e.props = append(e.props, p)
	return p
}

// ResetCounter will reset properties counter
func (e *Entity) ResetCounter(property string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Value:        "0",
		Type:         CounterType,
		MutationMode: int32(adl.MutationMode_WRITE),
	}
	e.props = append(e.props, p)
	return p
}

// RemoveCounter will remove the properties counter
func (e *Entity) RemoveCounter(property string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Type:         CounterType,
		MutationMode: int32(adl.MutationMode_DELETE),
	}
	e.props = append(e.props, p)
	return p
}
