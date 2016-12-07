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
	Props  []Property
	Client fdl.FdlClient
}

// Property is FDL's property structure
type Property struct {
	Property     string
	Value        string
	Type         PropertyType
	MutationMode int32
}

// PropertyType enumeration
type PropertyType int32

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

func commit(e *Entity, memberID string) error {
	props := map[string]*fdl.Property{}
	req := fdl.MutationRequest{
		Fid:        e.FID,
		MemberId:   memberID,
		Properties: props,
	}

	for n, p := range e.Props {
		props[fmt.Sprintf("%d", n)] = &fdl.Property{
			Property: p.Property,
			Type:     fdl.PropertyType(p.Type),
			Value:    p.Value,
			Mode:     fdl.MutationMode(p.MutationMode),
		}
	}

	e.Client.Mutate(ctx, &req)
	return nil
}

// Properties collects the properties to mutate
func (e *Entity) Properties(props ...Property) *Entity {
	e.Props = append(e.Props, props...)
	return e
}

// Write will create / Replace Property / Replace a list
func Write(property, value string, nType PropertyType) Property {
	return Property{
		Property:     property,
		Value:        value,
		Type:         nType,
		MutationMode: int32(fdl.MutationMode_WRITE),
	}
}

// Append will append items to a list / Append string to existing value for data or meta / Increment count for counter
func Append(property, value string, nType PropertyType) Property {
	return Property{
		Property:     property,
		Value:        value,
		Type:         nType,
		MutationMode: int32(fdl.MutationMode_APPEND),
	}
}

// Remove will remove items from a list / Decrement count for counter
func Remove(property, value string, nType PropertyType) Property {
	return Property{
		Property:     property,
		Value:        value,
		Type:         nType,
		MutationMode: int32(fdl.MutationMode_REMOVE),
	}
}

// Delete will delete a Property
func Delete(property, value string, nType PropertyType) Property {
	return Property{
		Property:     property,
		Value:        value,
		Type:         nType,
		MutationMode: int32(fdl.MutationMode_DELETE),
	}
}

// Mutate starts FDL mutation on given FID
func Mutate(fid string, c *fdl.FdlClient) *Entity {
	return &Entity{FID: fid, Props: []Property{}, Client: *c}
}
