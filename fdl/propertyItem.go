package fdl

import "github.com/gogo/protobuf/proto"

const (
	// MetaType is the meta property type
	MetaType PropertyType = 0

	// DataType is the data property type
	DataType PropertyType = 1

	// CounterType is the type for FDL counter
	CounterType PropertyType = 2

	// SetType is the type for FDL set
	SetType PropertyType = 3
)

// PropertyItem is FDL's property structure
type PropertyItem struct {
	Property     string
	Value        string
	Type         PropertyType
	MutationMode int32
	IsPrefix     bool
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
}

// String function for PropertyType enumeration
func (x PropertyType) String() string {
	return proto.EnumName(propertyTypename, int32(x))
}
