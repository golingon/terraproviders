// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSimpledbDomain creates a new instance of [SimpledbDomain].
func NewSimpledbDomain(name string, args SimpledbDomainArgs) *SimpledbDomain {
	return &SimpledbDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SimpledbDomain)(nil)

// SimpledbDomain represents the Terraform resource aws_simpledb_domain.
type SimpledbDomain struct {
	Name      string
	Args      SimpledbDomainArgs
	state     *simpledbDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SimpledbDomain].
func (sd *SimpledbDomain) Type() string {
	return "aws_simpledb_domain"
}

// LocalName returns the local name for [SimpledbDomain].
func (sd *SimpledbDomain) LocalName() string {
	return sd.Name
}

// Configuration returns the configuration (args) for [SimpledbDomain].
func (sd *SimpledbDomain) Configuration() interface{} {
	return sd.Args
}

// DependOn is used for other resources to depend on [SimpledbDomain].
func (sd *SimpledbDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(sd)
}

// Dependencies returns the list of resources [SimpledbDomain] depends_on.
func (sd *SimpledbDomain) Dependencies() terra.Dependencies {
	return sd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SimpledbDomain].
func (sd *SimpledbDomain) LifecycleManagement() *terra.Lifecycle {
	return sd.Lifecycle
}

// Attributes returns the attributes for [SimpledbDomain].
func (sd *SimpledbDomain) Attributes() simpledbDomainAttributes {
	return simpledbDomainAttributes{ref: terra.ReferenceResource(sd)}
}

// ImportState imports the given attribute values into [SimpledbDomain]'s state.
func (sd *SimpledbDomain) ImportState(av io.Reader) error {
	sd.state = &simpledbDomainState{}
	if err := json.NewDecoder(av).Decode(sd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sd.Type(), sd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SimpledbDomain] has state.
func (sd *SimpledbDomain) State() (*simpledbDomainState, bool) {
	return sd.state, sd.state != nil
}

// StateMust returns the state for [SimpledbDomain]. Panics if the state is nil.
func (sd *SimpledbDomain) StateMust() *simpledbDomainState {
	if sd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sd.Type(), sd.LocalName()))
	}
	return sd.state
}

// SimpledbDomainArgs contains the configurations for aws_simpledb_domain.
type SimpledbDomainArgs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type simpledbDomainAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_simpledb_domain.
func (sd simpledbDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("id"))
}

// Name returns a reference to field name of aws_simpledb_domain.
func (sd simpledbDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("name"))
}

type simpledbDomainState struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
