// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread_user_flow_attribute

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

// Resource represents the Terraform resource azuread_user_flow_attribute.
type Resource struct {
	Name      string
	Args      Args
	state     *azureadUserFlowAttributeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aufa *Resource) Type() string {
	return "azuread_user_flow_attribute"
}

// LocalName returns the local name for [Resource].
func (aufa *Resource) LocalName() string {
	return aufa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aufa *Resource) Configuration() interface{} {
	return aufa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aufa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aufa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aufa *Resource) Dependencies() terra.Dependencies {
	return aufa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aufa *Resource) LifecycleManagement() *terra.Lifecycle {
	return aufa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aufa *Resource) Attributes() azureadUserFlowAttributeAttributes {
	return azureadUserFlowAttributeAttributes{ref: terra.ReferenceResource(aufa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aufa *Resource) ImportState(state io.Reader) error {
	aufa.state = &azureadUserFlowAttributeState{}
	if err := json.NewDecoder(state).Decode(aufa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aufa.Type(), aufa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aufa *Resource) State() (*azureadUserFlowAttributeState, bool) {
	return aufa.state, aufa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aufa *Resource) StateMust() *azureadUserFlowAttributeState {
	if aufa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aufa.Type(), aufa.LocalName()))
	}
	return aufa.state
}

// Args contains the configurations for azuread_user_flow_attribute.
type Args struct {
	// DataType: string, required
	DataType terra.StringValue `hcl:"data_type,attr" validate:"required"`
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azureadUserFlowAttributeAttributes struct {
	ref terra.Reference
}

// AttributeType returns a reference to field attribute_type of azuread_user_flow_attribute.
func (aufa azureadUserFlowAttributeAttributes) AttributeType() terra.StringValue {
	return terra.ReferenceAsString(aufa.ref.Append("attribute_type"))
}

// DataType returns a reference to field data_type of azuread_user_flow_attribute.
func (aufa azureadUserFlowAttributeAttributes) DataType() terra.StringValue {
	return terra.ReferenceAsString(aufa.ref.Append("data_type"))
}

// Description returns a reference to field description of azuread_user_flow_attribute.
func (aufa azureadUserFlowAttributeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aufa.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_user_flow_attribute.
func (aufa azureadUserFlowAttributeAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aufa.ref.Append("display_name"))
}

// Id returns a reference to field id of azuread_user_flow_attribute.
func (aufa azureadUserFlowAttributeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aufa.ref.Append("id"))
}

func (aufa azureadUserFlowAttributeAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aufa.ref.Append("timeouts"))
}

type azureadUserFlowAttributeState struct {
	AttributeType string         `json:"attribute_type"`
	DataType      string         `json:"data_type"`
	Description   string         `json:"description"`
	DisplayName   string         `json:"display_name"`
	Id            string         `json:"id"`
	Timeouts      *TimeoutsState `json:"timeouts"`
}
