package fdl

import (
	"fmt"

	portcullis "github.com/fortifi/portcullis-go"
	"github.com/fortifi/proto-go/fdl"

	"encoding/json"

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
	props := []*fdl.ReadProperty{}
	lst := []PropertyItem{}

	for _, p := range e.rProps {
		if p.Type != ListType {
			props = append(props, &fdl.ReadProperty{
				Property: p.Property,
				Type:     fdl.PropertyType(p.Type),
				IsPrefix: p.IsPrefix,
			})
		} else {
			lst = append(lst, p)
		}
	}

	req := fdl.ReadRequest{
		Fid:        e.fid,
		MemberId:   "",
		Properties: props,
	}

	ret := map[string][]ResultItem{}

	if len(props) > 0 {
		res, err := e.client.Read(ctx, &req)
		if err != nil {
			return Result{}, err
		}

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
	}

	if len(lst) > 0 {
		for _, lp := range lst {
			f := []KeyValuePair{}
			if lp.StartKey == "" && lp.EndKey == "" {
				// list retrieve item
				req := fdl.KeyRequest{
					Fid:      e.fid,
					ListName: lp.Property,
					Key:      lp.Key,
					MemberId: "",
				}

				r, err := e.client.ListRetrieveItem(ctx, &req)
				if err != nil {
					return Result{}, err
				}

				f = append(f, KeyValuePair{
					Key:   r.Key,
					Value: r.Value,
				})
			} else {
				// list retrieve range
				req := fdl.ListRangeRequest{
					Fid:      e.fid,
					ListName: lp.Property,
					StartKey: lp.StartKey,
					EndKey:   lp.EndKey,
					Limit:    lp.Limit,
					MemberId: "",
				}

				r, err := e.client.ListRange(ctx, &req)
				if err != nil {
					return Result{}, err
				}

				for _, l := range r.Items {
					f = append(f, KeyValuePair{
						Key:   l.Key,
						Value: l.Value,
					})
				}
			}

			m, err := json.Marshal(f)
			if err != nil {
				return Result{}, err
			}

			ri := ResultItem{
				Property: lp.Property,
				Value:    string(m),
				Type:     ListType,
			}

			ret[fmt.Sprintf("%s_%d", lp.Property, ListType)] = []ResultItem{ri}
		}
	}

	return Result{Items: ret}, nil
}

func commit(e *Entity, memberID string) error {
	props := []*fdl.Property{}
	lst := []PropertyItem{}
	for _, p := range e.props {
		if p.Type != ListType {
			props = append(props, &fdl.Property{
				Property: p.Property,
				Type:     fdl.PropertyType(p.Type),
				Value:    p.Value,
				Mode:     fdl.MutationMode(p.MutationMode),
			})
		} else {
			lst = append(lst, p)
		}
	}

	if len(props) > 0 {
		req := fdl.MutationRequest{
			Fid:        e.fid,
			MemberId:   memberID,
			Properties: props,
		}
		_, err := e.client.Mutate(ctx, &req)
		if err != nil {
			return err
		}
	}

	if len(lst) > 0 {
		for _, lp := range lst {
			if lp.MutationMode == int32(fdl.MutationMode_WRITE) {
				// list write action
				r := fdl.ListAddRequest{
					Fid:      e.fid,
					ListName: lp.Property,
					Key:      lp.Key,
					Value:    lp.Value,
					MemberId: memberID,
				}
				_, err := e.client.ListAdd(ctx, &r)
				if err != nil {
					return err
				}
			} else if lp.MutationMode == int32(fdl.MutationMode_REMOVE) {
				// list remove action
				r := fdl.KeyRequest{
					Fid:      e.fid,
					ListName: lp.Property,
					Key:      lp.Key,
					MemberId: memberID,
				}
				_, err := e.client.ListRemove(ctx, &r)
				if err != nil {
					return err
				}
			}
		}
	}

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
