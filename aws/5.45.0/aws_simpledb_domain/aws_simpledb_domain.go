// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_simpledb_domain

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_simpledb_domain.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSimpledbDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asd *Resource) Type() string {
	return "aws_simpledb_domain"
}

// LocalName returns the local name for [Resource].
func (asd *Resource) LocalName() string {
	return asd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asd *Resource) Configuration() interface{} {
	return asd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asd *Resource) Dependencies() terra.Dependencies {
	return asd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asd *Resource) LifecycleManagement() *terra.Lifecycle {
	return asd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asd *Resource) Attributes() awsSimpledbDomainAttributes {
	return awsSimpledbDomainAttributes{ref: terra.ReferenceResource(asd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asd *Resource) ImportState(state io.Reader) error {
	asd.state = &awsSimpledbDomainState{}
	if err := json.NewDecoder(state).Decode(asd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asd.Type(), asd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asd *Resource) State() (*awsSimpledbDomainState, bool) {
	return asd.state, asd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asd *Resource) StateMust() *awsSimpledbDomainState {
	if asd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asd.Type(), asd.LocalName()))
	}
	return asd.state
}

// Args contains the configurations for aws_simpledb_domain.
type Args struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type awsSimpledbDomainAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_simpledb_domain.
func (asd awsSimpledbDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("id"))
}

// Name returns a reference to field name of aws_simpledb_domain.
func (asd awsSimpledbDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("name"))
}

type awsSimpledbDomainState struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
