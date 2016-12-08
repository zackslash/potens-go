package fdl

import (
	"fmt"

	portcullis "github.com/fortifi/portcullis-go"
	"github.com/fortifi/proto-go/fdl"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// FDLGAID is the FDL Global App ID
	FDLGAID = "fortifi/fdl"

	// Meta proerty type
	Meta PropertyType = 0

	// Data property type
	Data PropertyType = 1

	// List property type
	List PropertyType = 2

	// UniqueList property type
	UniqueList PropertyType = 3

	// Counter property type
	Counter PropertyType = 4
)

var (
	con   *grpc.ClientConn
	ctx   context.Context
	appID string
)

// Entity is the structure for FDL mutation
type Entity struct {
	FID    string
	Props  []PropertyItem
	Client fdl.FdlClient
}

// PropertyItem is FDL's property structure
type PropertyItem struct {
	Property     string
	Value        string
	Type         PropertyType
	MutationMode int32
}

// PropertyType enumeration
type PropertyType int32

// SetContextAppID sets both context and AppID
func SetContextAppID(ctxn context.Context, appid string) {
	ctx = ctxn
	appID = appid
}

// CommitWithUserID writes mutation to FDL with given UserID
func (e *Entity) CommitWithUserID(memberID string) error {
	return commit(e, memberID)
}

// CommitWithContext writes mutation to FDL with ID available in context
func (e *Entity) CommitWithContext(ctx context.Context) error {
	authData := portcullis.FromContext(ctx)
	member := authData.UserID

	if member == "" {
		member = authData.AppID
	}

	if member == "" {
		member = appID
	}

	return commit(e, member)
}

// Commit writes mitation to FDL with application's ID
func (e *Entity) Commit() error {
	return commit(e, appID)
}

// Retrieve starts the process of data retrieval from FDL
func (e *Entity) Retrieve() ([]*fdl.Property, error) {
	return retrieve(e)
}

func retrieve(e *Entity) ([]*fdl.Property, error) {
	props := map[string]*fdl.Property{}

	for _, p := range e.Props {
		props[p.Property] = &fdl.Property{
			Property: p.Property,
		}
	}

	req := fdl.ReadRequest{
		Fid:        e.FID,
		MemberId:   "",
		Properties: props,
	}

	res, err := e.Client.Read(ctx, &req)
	if err != nil {
		return []*fdl.Property{}, err
	}

	ret := []*fdl.Property{}
	for _, r := range res.Properties {
		ret = append(ret, r)
	}

	return ret, nil
}

func commit(e *Entity, memberID string) error {
	props := map[string]*fdl.Property{}

	for n, p := range e.Props {
		props[fmt.Sprintf("%d", n)] = &fdl.Property{
			Property: p.Property,
			Type:     fdl.PropertyType(p.Type),
			Value:    p.Value,
			Mode:     fdl.MutationMode(p.MutationMode),
		}
	}

	req := fdl.MutationRequest{
		Fid:        e.FID,
		MemberId:   memberID,
		Properties: props,
	}

	e.Client.Mutate(ctx, &req)
	return nil
}

// Mutate collects the properties to mutate
func (e *Entity) Mutate(props ...PropertyItem) *Entity {
	e.Props = append(e.Props, props...)
	return e
}

// Read collects the properties to read
func (e *Entity) Read(props ...PropertyItem) *Entity {
	e.Props = append(e.Props, props...)
	return e
}

// Property creates a new property item
func Property(property string) PropertyItem {
	return PropertyItem{Property: property}
}

// Write will create / Replace Property / Replace a list
func Write(property, value string, nType PropertyType) PropertyItem {
	return PropertyItem{
		Property:     property,
		Value:        value,
		Type:         nType,
		MutationMode: int32(fdl.MutationMode_WRITE),
	}
}

// Append will append items to a list / Append string to existing value for data or meta / Increment count for counter
func Append(property, value string, nType PropertyType) PropertyItem {
	return PropertyItem{
		Property:     property,
		Value:        value,
		Type:         nType,
		MutationMode: int32(fdl.MutationMode_APPEND),
	}
}

// Remove will remove items from a list / Decrement count for counter
func Remove(property, value string, nType PropertyType) PropertyItem {
	return PropertyItem{
		Property:     property,
		Value:        value,
		Type:         nType,
		MutationMode: int32(fdl.MutationMode_REMOVE),
	}
}

// Delete will delete a Property
func Delete(property, value string, nType PropertyType) PropertyItem {
	return PropertyItem{
		Property:     property,
		Value:        value,
		Type:         nType,
		MutationMode: int32(fdl.MutationMode_DELETE),
	}
}

// Mutate starts FDL mutation on given FID
func Mutate(fid string, c *fdl.FdlClient) *Entity {
	return &Entity{FID: fid, Props: []PropertyItem{}, Client: *c}
}
