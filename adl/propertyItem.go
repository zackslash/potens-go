package adl

import "github.com/gogo/protobuf/proto"

const (
	// MetaType is the meta property type
	MetaType PropertyType = 0

	// DataType is the data property type
	DataType PropertyType = 1

	// CounterType is the type for adl counter
	CounterType PropertyType = 2

	// SetType is the type for adl set
	SetType PropertyType = 3

	// ListType is the type for adl list
	ListType PropertyType = 4
)

// PropertyItem is adl's property structure
type PropertyItem struct {
	Property     string
	Key          string
	Value        string
	Type         PropertyType
	MutationMode int32
	IsPrefix     bool
	StartKey     string
	EndKey       string
	Limit        int32
}

// PropertyItems is a sortable slice of property item
type PropertyItems []PropertyItem

// PropertyType enumeration
type PropertyType int32

var propertyTypename = map[int32]string{
	0: "meta",
	1: "data",
	2: "counter",
	3: "set",
	4: "list",
}

// String function for PropertyType enumeration
func (x PropertyType) String() string {
	return proto.EnumName(propertyTypename, int32(x))
}
