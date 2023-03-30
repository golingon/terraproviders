// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	keyspaceskeyspace "github.com/golingon/terraproviders/aws/4.60.0/keyspaceskeyspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewKeyspacesKeyspace(name string, args KeyspacesKeyspaceArgs) *KeyspacesKeyspace {
	return &KeyspacesKeyspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyspacesKeyspace)(nil)

type KeyspacesKeyspace struct {
	Name  string
	Args  KeyspacesKeyspaceArgs
	state *keyspacesKeyspaceState
}

func (kk *KeyspacesKeyspace) Type() string {
	return "aws_keyspaces_keyspace"
}

func (kk *KeyspacesKeyspace) LocalName() string {
	return kk.Name
}

func (kk *KeyspacesKeyspace) Configuration() interface{} {
	return kk.Args
}

func (kk *KeyspacesKeyspace) Attributes() keyspacesKeyspaceAttributes {
	return keyspacesKeyspaceAttributes{ref: terra.ReferenceResource(kk)}
}

func (kk *KeyspacesKeyspace) ImportState(av io.Reader) error {
	kk.state = &keyspacesKeyspaceState{}
	if err := json.NewDecoder(av).Decode(kk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kk.Type(), kk.LocalName(), err)
	}
	return nil
}

func (kk *KeyspacesKeyspace) State() (*keyspacesKeyspaceState, bool) {
	return kk.state, kk.state != nil
}

func (kk *KeyspacesKeyspace) StateMust() *keyspacesKeyspaceState {
	if kk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kk.Type(), kk.LocalName()))
	}
	return kk.state
}

func (kk *KeyspacesKeyspace) DependOn() terra.Reference {
	return terra.ReferenceResource(kk)
}

type KeyspacesKeyspaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *keyspaceskeyspace.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that KeyspacesKeyspace depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type keyspacesKeyspaceAttributes struct {
	ref terra.Reference
}

func (kk keyspacesKeyspaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(kk.ref.Append("arn"))
}

func (kk keyspacesKeyspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(kk.ref.Append("id"))
}

func (kk keyspacesKeyspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(kk.ref.Append("name"))
}

func (kk keyspacesKeyspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](kk.ref.Append("tags"))
}

func (kk keyspacesKeyspaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](kk.ref.Append("tags_all"))
}

func (kk keyspacesKeyspaceAttributes) Timeouts() keyspaceskeyspace.TimeoutsAttributes {
	return terra.ReferenceSingle[keyspaceskeyspace.TimeoutsAttributes](kk.ref.Append("timeouts"))
}

type keyspacesKeyspaceState struct {
	Arn      string                           `json:"arn"`
	Id       string                           `json:"id"`
	Name     string                           `json:"name"`
	Tags     map[string]string                `json:"tags"`
	TagsAll  map[string]string                `json:"tags_all"`
	Timeouts *keyspaceskeyspace.TimeoutsState `json:"timeouts"`
}
