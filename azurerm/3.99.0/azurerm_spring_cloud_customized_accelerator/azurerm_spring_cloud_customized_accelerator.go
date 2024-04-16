// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_spring_cloud_customized_accelerator

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

// Resource represents the Terraform resource azurerm_spring_cloud_customized_accelerator.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSpringCloudCustomizedAcceleratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ascca *Resource) Type() string {
	return "azurerm_spring_cloud_customized_accelerator"
}

// LocalName returns the local name for [Resource].
func (ascca *Resource) LocalName() string {
	return ascca.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ascca *Resource) Configuration() interface{} {
	return ascca.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ascca *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ascca)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ascca *Resource) Dependencies() terra.Dependencies {
	return ascca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ascca *Resource) LifecycleManagement() *terra.Lifecycle {
	return ascca.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ascca *Resource) Attributes() azurermSpringCloudCustomizedAcceleratorAttributes {
	return azurermSpringCloudCustomizedAcceleratorAttributes{ref: terra.ReferenceResource(ascca)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ascca *Resource) ImportState(state io.Reader) error {
	ascca.state = &azurermSpringCloudCustomizedAcceleratorState{}
	if err := json.NewDecoder(state).Decode(ascca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ascca.Type(), ascca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ascca *Resource) State() (*azurermSpringCloudCustomizedAcceleratorState, bool) {
	return ascca.state, ascca.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ascca *Resource) StateMust() *azurermSpringCloudCustomizedAcceleratorState {
	if ascca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ascca.Type(), ascca.LocalName()))
	}
	return ascca.state
}

// Args contains the configurations for azurerm_spring_cloud_customized_accelerator.
type Args struct {
	// AcceleratorTags: list of string, optional
	AcceleratorTags terra.ListValue[terra.StringValue] `hcl:"accelerator_tags,attr"`
	// AcceleratorType: string, optional
	AcceleratorType terra.StringValue `hcl:"accelerator_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// IconUrl: string, optional
	IconUrl terra.StringValue `hcl:"icon_url,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudAcceleratorId: string, required
	SpringCloudAcceleratorId terra.StringValue `hcl:"spring_cloud_accelerator_id,attr" validate:"required"`
	// GitRepository: required
	GitRepository *GitRepository `hcl:"git_repository,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSpringCloudCustomizedAcceleratorAttributes struct {
	ref terra.Reference
}

// AcceleratorTags returns a reference to field accelerator_tags of azurerm_spring_cloud_customized_accelerator.
func (ascca azurermSpringCloudCustomizedAcceleratorAttributes) AcceleratorTags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ascca.ref.Append("accelerator_tags"))
}

// AcceleratorType returns a reference to field accelerator_type of azurerm_spring_cloud_customized_accelerator.
func (ascca azurermSpringCloudCustomizedAcceleratorAttributes) AcceleratorType() terra.StringValue {
	return terra.ReferenceAsString(ascca.ref.Append("accelerator_type"))
}

// Description returns a reference to field description of azurerm_spring_cloud_customized_accelerator.
func (ascca azurermSpringCloudCustomizedAcceleratorAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ascca.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_spring_cloud_customized_accelerator.
func (ascca azurermSpringCloudCustomizedAcceleratorAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ascca.ref.Append("display_name"))
}

// IconUrl returns a reference to field icon_url of azurerm_spring_cloud_customized_accelerator.
func (ascca azurermSpringCloudCustomizedAcceleratorAttributes) IconUrl() terra.StringValue {
	return terra.ReferenceAsString(ascca.ref.Append("icon_url"))
}

// Id returns a reference to field id of azurerm_spring_cloud_customized_accelerator.
func (ascca azurermSpringCloudCustomizedAcceleratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ascca.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_customized_accelerator.
func (ascca azurermSpringCloudCustomizedAcceleratorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ascca.ref.Append("name"))
}

// SpringCloudAcceleratorId returns a reference to field spring_cloud_accelerator_id of azurerm_spring_cloud_customized_accelerator.
func (ascca azurermSpringCloudCustomizedAcceleratorAttributes) SpringCloudAcceleratorId() terra.StringValue {
	return terra.ReferenceAsString(ascca.ref.Append("spring_cloud_accelerator_id"))
}

func (ascca azurermSpringCloudCustomizedAcceleratorAttributes) GitRepository() terra.ListValue[GitRepositoryAttributes] {
	return terra.ReferenceAsList[GitRepositoryAttributes](ascca.ref.Append("git_repository"))
}

func (ascca azurermSpringCloudCustomizedAcceleratorAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ascca.ref.Append("timeouts"))
}

type azurermSpringCloudCustomizedAcceleratorState struct {
	AcceleratorTags          []string             `json:"accelerator_tags"`
	AcceleratorType          string               `json:"accelerator_type"`
	Description              string               `json:"description"`
	DisplayName              string               `json:"display_name"`
	IconUrl                  string               `json:"icon_url"`
	Id                       string               `json:"id"`
	Name                     string               `json:"name"`
	SpringCloudAcceleratorId string               `json:"spring_cloud_accelerator_id"`
	GitRepository            []GitRepositoryState `json:"git_repository"`
	Timeouts                 *TimeoutsState       `json:"timeouts"`
}
