package fdl

import (
	"github.com/fortifi/proto-go/fdl"
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
		MutationMode: int32(fdl.MutationMode_APPEND),
		order:        len(e.props),
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
		MutationMode: int32(fdl.MutationMode_REMOVE),
		order:        len(e.props),
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
		MutationMode: int32(fdl.MutationMode_WRITE),
		order:        len(e.props),
	}
	e.props = append(e.props, p)
	return p
}

// RemoveCounter will remove the properties counter
func (e *Entity) RemoveCounter(property string) PropertyItem {
	p := PropertyItem{
		Property:     property,
		Type:         CounterType,
		MutationMode: int32(fdl.MutationMode_DELETE),
		order:        len(e.props),
	}
	e.props = append(e.props, p)
	return p
}
