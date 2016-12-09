package fdl

import "github.com/gogo/protobuf/proto"

const (
	// MetaType is the meta property type
	MetaType PropertyType = 0

	// DataType is the data property type
	DataType PropertyType = 1

	// ListType is the list property type
	ListType PropertyType = 2

	// UniqueListType is the unique list property type
	UniqueListType PropertyType = 3

	// CounterType is the counter property type
	CounterType PropertyType = 4
)

// PropertyItem is FDL's property structure
type PropertyItem struct {
	order        int
	Property     string
	Value        string
	Type         PropertyType
	MutationMode int32
}

// PropertyItems is a sortable slice of property item
type PropertyItems []PropertyItem

// PropertyType enumeration
type PropertyType int32

var propertyTypename = map[int32]string{
	0: "meta",
	1: "data",
	2: "list",
	3: "ulist",
	4: "counter",
}

// String function for PropertyType enumeration
func (x PropertyType) String() string {
	return proto.EnumName(propertyTypename, int32(x))
}

func (slice PropertyItems) Len() int {
	return len(slice)
}

func (slice PropertyItems) Less(i, j int) bool {
	return slice[i].order < slice[j].order
}

func (slice PropertyItems) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
