// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	keyspaceskeyspace "github.com/golingon/terraproviders/aws/5.6.2/keyspaceskeyspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKeyspacesKeyspace creates a new instance of [KeyspacesKeyspace].
func NewKeyspacesKeyspace(name string, args KeyspacesKeyspaceArgs) *KeyspacesKeyspace {
	return &KeyspacesKeyspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyspacesKeyspace)(nil)

// KeyspacesKeyspace represents the Terraform resource aws_keyspaces_keyspace.
type KeyspacesKeyspace struct {
	Name      string
	Args      KeyspacesKeyspaceArgs
	state     *keyspacesKeyspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KeyspacesKeyspace].
func (kk *KeyspacesKeyspace) Type() string {
	return "aws_keyspaces_keyspace"
}

// LocalName returns the local name for [KeyspacesKeyspace].
func (kk *KeyspacesKeyspace) LocalName() string {
	return kk.Name
}

// Configuration returns the configuration (args) for [KeyspacesKeyspace].
func (kk *KeyspacesKeyspace) Configuration() interface{} {
	return kk.Args
}

// DependOn is used for other resources to depend on [KeyspacesKeyspace].
func (kk *KeyspacesKeyspace) DependOn() terra.Reference {
	return terra.ReferenceResource(kk)
}

// Dependencies returns the list of resources [KeyspacesKeyspace] depends_on.
func (kk *KeyspacesKeyspace) Dependencies() terra.Dependencies {
	return kk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KeyspacesKeyspace].
func (kk *KeyspacesKeyspace) LifecycleManagement() *terra.Lifecycle {
	return kk.Lifecycle
}

// Attributes returns the attributes for [KeyspacesKeyspace].
func (kk *KeyspacesKeyspace) Attributes() keyspacesKeyspaceAttributes {
	return keyspacesKeyspaceAttributes{ref: terra.ReferenceResource(kk)}
}

// ImportState imports the given attribute values into [KeyspacesKeyspace]'s state.
func (kk *KeyspacesKeyspace) ImportState(av io.Reader) error {
	kk.state = &keyspacesKeyspaceState{}
	if err := json.NewDecoder(av).Decode(kk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kk.Type(), kk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KeyspacesKeyspace] has state.
func (kk *KeyspacesKeyspace) State() (*keyspacesKeyspaceState, bool) {
	return kk.state, kk.state != nil
}

// StateMust returns the state for [KeyspacesKeyspace]. Panics if the state is nil.
func (kk *KeyspacesKeyspace) StateMust() *keyspacesKeyspaceState {
	if kk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kk.Type(), kk.LocalName()))
	}
	return kk.state
}

// KeyspacesKeyspaceArgs contains the configurations for aws_keyspaces_keyspace.
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
}
type keyspacesKeyspaceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_keyspaces_keyspace.
func (kk keyspacesKeyspaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("arn"))
}

// Id returns a reference to field id of aws_keyspaces_keyspace.
func (kk keyspacesKeyspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("id"))
}

// Name returns a reference to field name of aws_keyspaces_keyspace.
func (kk keyspacesKeyspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_keyspaces_keyspace.
func (kk keyspacesKeyspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kk.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_keyspaces_keyspace.
func (kk keyspacesKeyspaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kk.ref.Append("tags_all"))
}

func (kk keyspacesKeyspaceAttributes) Timeouts() keyspaceskeyspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[keyspaceskeyspace.TimeoutsAttributes](kk.ref.Append("timeouts"))
}

type keyspacesKeyspaceState struct {
	Arn      string                           `json:"arn"`
	Id       string                           `json:"id"`
	Name     string                           `json:"name"`
	Tags     map[string]string                `json:"tags"`
	TagsAll  map[string]string                `json:"tags_all"`
	Timeouts *keyspaceskeyspace.TimeoutsState `json:"timeouts"`
}
