// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_spring_cloud_builder

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

// Resource represents the Terraform resource azurerm_spring_cloud_builder.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSpringCloudBuilderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ascb *Resource) Type() string {
	return "azurerm_spring_cloud_builder"
}

// LocalName returns the local name for [Resource].
func (ascb *Resource) LocalName() string {
	return ascb.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ascb *Resource) Configuration() interface{} {
	return ascb.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ascb *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ascb)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ascb *Resource) Dependencies() terra.Dependencies {
	return ascb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ascb *Resource) LifecycleManagement() *terra.Lifecycle {
	return ascb.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ascb *Resource) Attributes() azurermSpringCloudBuilderAttributes {
	return azurermSpringCloudBuilderAttributes{ref: terra.ReferenceResource(ascb)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ascb *Resource) ImportState(state io.Reader) error {
	ascb.state = &azurermSpringCloudBuilderState{}
	if err := json.NewDecoder(state).Decode(ascb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ascb.Type(), ascb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ascb *Resource) State() (*azurermSpringCloudBuilderState, bool) {
	return ascb.state, ascb.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ascb *Resource) StateMust() *azurermSpringCloudBuilderState {
	if ascb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ascb.Type(), ascb.LocalName()))
	}
	return ascb.state
}

// Args contains the configurations for azurerm_spring_cloud_builder.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudServiceId: string, required
	SpringCloudServiceId terra.StringValue `hcl:"spring_cloud_service_id,attr" validate:"required"`
	// BuildPackGroup: min=1
	BuildPackGroup []BuildPackGroup `hcl:"build_pack_group,block" validate:"min=1"`
	// Stack: required
	Stack *Stack `hcl:"stack,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSpringCloudBuilderAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_spring_cloud_builder.
func (ascb azurermSpringCloudBuilderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ascb.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_builder.
func (ascb azurermSpringCloudBuilderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ascb.ref.Append("name"))
}

// SpringCloudServiceId returns a reference to field spring_cloud_service_id of azurerm_spring_cloud_builder.
func (ascb azurermSpringCloudBuilderAttributes) SpringCloudServiceId() terra.StringValue {
	return terra.ReferenceAsString(ascb.ref.Append("spring_cloud_service_id"))
}

func (ascb azurermSpringCloudBuilderAttributes) BuildPackGroup() terra.SetValue[BuildPackGroupAttributes] {
	return terra.ReferenceAsSet[BuildPackGroupAttributes](ascb.ref.Append("build_pack_group"))
}

func (ascb azurermSpringCloudBuilderAttributes) Stack() terra.ListValue[StackAttributes] {
	return terra.ReferenceAsList[StackAttributes](ascb.ref.Append("stack"))
}

func (ascb azurermSpringCloudBuilderAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ascb.ref.Append("timeouts"))
}

type azurermSpringCloudBuilderState struct {
	Id                   string                `json:"id"`
	Name                 string                `json:"name"`
	SpringCloudServiceId string                `json:"spring_cloud_service_id"`
	BuildPackGroup       []BuildPackGroupState `json:"build_pack_group"`
	Stack                []StackState          `json:"stack"`
	Timeouts             *TimeoutsState        `json:"timeouts"`
}
