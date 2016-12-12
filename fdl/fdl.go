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
)

var (
	con   *grpc.ClientConn
	ctx   context.Context
	appID string
)

// Entity is the structure for FDL mutation
type Entity struct {
	fid    string
	props  PropertyItems
	rProps PropertyItems
	result Result
	client fdl.FdlClient
}

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

// retrieve starts the process of data retrieval from FDL
func retrieve(e *Entity) (Result, error) {
	props := []*fdl.Property{}
	for _, p := range e.rProps {
		props = append(props, &fdl.Property{
			Property: p.Property,
			Type:     fdl.PropertyType(p.Type),
		})
	}

	req := fdl.ReadRequest{
		Fid:        e.fid,
		MemberId:   "",
		Properties: props,
	}

	res, err := e.client.Read(ctx, &req)
	if err != nil {
		return Result{}, err
	}

	ret := map[string][]ResultItem{}
	for _, r := range res.Properties {
		nr := ResultItem{
			Property: r.Property,
			Value:    r.Value,
			Type:     PropertyType(r.Type),
		}

		cur := ret[r.Property]
		if cur == nil {
			cur = []ResultItem{}
		}
		cur = append(cur, nr)
		ret[fmt.Sprintf("%s_%d", r.Property, r.Type)] = cur
	}

	return Result{Items: ret}, nil
}

func commit(e *Entity, memberID string) error {
	props := []*fdl.Property{}
	for _, p := range e.props {
		props = append(props, &fdl.Property{
			Property: p.Property,
			Type:     fdl.PropertyType(p.Type),
			Value:    p.Value,
			Mode:     fdl.MutationMode(p.MutationMode),
		})
	}

	req := fdl.MutationRequest{
		Fid:        e.fid,
		MemberId:   memberID,
		Properties: props,
	}

	e.client.Mutate(ctx, &req)
	e.props = []PropertyItem{}
	return nil
}

// Read collects the properties to read
func (e *Entity) Read(props ...PropertyItem) error {
	e.rProps = append(e.rProps, props...)
	res, err := retrieve(e)
	e.rProps = []PropertyItem{}
	e.result = res
	return err
}

// Remove will remove a property
func Remove(property, value string, nType PropertyType) PropertyItem {
	return PropertyItem{
		Property:     property,
		Value:        value,
		Type:         nType,
		MutationMode: int32(fdl.MutationMode_REMOVE),
	}
}

// Mutate starts FDL mutation on given FID
func Mutate(fid string, c *fdl.FdlClient) *Entity {
	return &Entity{fid: fid, props: []PropertyItem{}, rProps: []PropertyItem{}, client: *c}
}
