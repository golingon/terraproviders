// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	inspector2enabler "github.com/golingon/terraproviders/aws/4.63.0/inspector2enabler"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewInspector2Enabler creates a new instance of [Inspector2Enabler].
func NewInspector2Enabler(name string, args Inspector2EnablerArgs) *Inspector2Enabler {
	return &Inspector2Enabler{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Inspector2Enabler)(nil)

// Inspector2Enabler represents the Terraform resource aws_inspector2_enabler.
type Inspector2Enabler struct {
	Name      string
	Args      Inspector2EnablerArgs
	state     *inspector2EnablerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Inspector2Enabler].
func (ie *Inspector2Enabler) Type() string {
	return "aws_inspector2_enabler"
}

// LocalName returns the local name for [Inspector2Enabler].
func (ie *Inspector2Enabler) LocalName() string {
	return ie.Name
}

// Configuration returns the configuration (args) for [Inspector2Enabler].
func (ie *Inspector2Enabler) Configuration() interface{} {
	return ie.Args
}

// DependOn is used for other resources to depend on [Inspector2Enabler].
func (ie *Inspector2Enabler) DependOn() terra.Reference {
	return terra.ReferenceResource(ie)
}

// Dependencies returns the list of resources [Inspector2Enabler] depends_on.
func (ie *Inspector2Enabler) Dependencies() terra.Dependencies {
	return ie.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Inspector2Enabler].
func (ie *Inspector2Enabler) LifecycleManagement() *terra.Lifecycle {
	return ie.Lifecycle
}

// Attributes returns the attributes for [Inspector2Enabler].
func (ie *Inspector2Enabler) Attributes() inspector2EnablerAttributes {
	return inspector2EnablerAttributes{ref: terra.ReferenceResource(ie)}
}

// ImportState imports the given attribute values into [Inspector2Enabler]'s state.
func (ie *Inspector2Enabler) ImportState(av io.Reader) error {
	ie.state = &inspector2EnablerState{}
	if err := json.NewDecoder(av).Decode(ie.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ie.Type(), ie.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Inspector2Enabler] has state.
func (ie *Inspector2Enabler) State() (*inspector2EnablerState, bool) {
	return ie.state, ie.state != nil
}

// StateMust returns the state for [Inspector2Enabler]. Panics if the state is nil.
func (ie *Inspector2Enabler) StateMust() *inspector2EnablerState {
	if ie.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ie.Type(), ie.LocalName()))
	}
	return ie.state
}

// Inspector2EnablerArgs contains the configurations for aws_inspector2_enabler.
type Inspector2EnablerArgs struct {
	// AccountIds: set of string, required
	AccountIds terra.SetValue[terra.StringValue] `hcl:"account_ids,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceTypes: set of string, required
	ResourceTypes terra.SetValue[terra.StringValue] `hcl:"resource_types,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *inspector2enabler.Timeouts `hcl:"timeouts,block"`
}
type inspector2EnablerAttributes struct {
	ref terra.Reference
}

// AccountIds returns a reference to field account_ids of aws_inspector2_enabler.
func (ie inspector2EnablerAttributes) AccountIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ie.ref.Append("account_ids"))
}

// Id returns a reference to field id of aws_inspector2_enabler.
func (ie inspector2EnablerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ie.ref.Append("id"))
}

// ResourceTypes returns a reference to field resource_types of aws_inspector2_enabler.
func (ie inspector2EnablerAttributes) ResourceTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ie.ref.Append("resource_types"))
}

func (ie inspector2EnablerAttributes) Timeouts() inspector2enabler.TimeoutsAttributes {
	return terra.ReferenceAsSingle[inspector2enabler.TimeoutsAttributes](ie.ref.Append("timeouts"))
}

type inspector2EnablerState struct {
	AccountIds    []string                         `json:"account_ids"`
	Id            string                           `json:"id"`
	ResourceTypes []string                         `json:"resource_types"`
	Timeouts      *inspector2enabler.TimeoutsState `json:"timeouts"`
}
