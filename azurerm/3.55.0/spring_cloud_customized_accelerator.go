// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudcustomizedaccelerator "github.com/golingon/terraproviders/azurerm/3.55.0/springcloudcustomizedaccelerator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudCustomizedAccelerator creates a new instance of [SpringCloudCustomizedAccelerator].
func NewSpringCloudCustomizedAccelerator(name string, args SpringCloudCustomizedAcceleratorArgs) *SpringCloudCustomizedAccelerator {
	return &SpringCloudCustomizedAccelerator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudCustomizedAccelerator)(nil)

// SpringCloudCustomizedAccelerator represents the Terraform resource azurerm_spring_cloud_customized_accelerator.
type SpringCloudCustomizedAccelerator struct {
	Name      string
	Args      SpringCloudCustomizedAcceleratorArgs
	state     *springCloudCustomizedAcceleratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudCustomizedAccelerator].
func (scca *SpringCloudCustomizedAccelerator) Type() string {
	return "azurerm_spring_cloud_customized_accelerator"
}

// LocalName returns the local name for [SpringCloudCustomizedAccelerator].
func (scca *SpringCloudCustomizedAccelerator) LocalName() string {
	return scca.Name
}

// Configuration returns the configuration (args) for [SpringCloudCustomizedAccelerator].
func (scca *SpringCloudCustomizedAccelerator) Configuration() interface{} {
	return scca.Args
}

// DependOn is used for other resources to depend on [SpringCloudCustomizedAccelerator].
func (scca *SpringCloudCustomizedAccelerator) DependOn() terra.Reference {
	return terra.ReferenceResource(scca)
}

// Dependencies returns the list of resources [SpringCloudCustomizedAccelerator] depends_on.
func (scca *SpringCloudCustomizedAccelerator) Dependencies() terra.Dependencies {
	return scca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudCustomizedAccelerator].
func (scca *SpringCloudCustomizedAccelerator) LifecycleManagement() *terra.Lifecycle {
	return scca.Lifecycle
}

// Attributes returns the attributes for [SpringCloudCustomizedAccelerator].
func (scca *SpringCloudCustomizedAccelerator) Attributes() springCloudCustomizedAcceleratorAttributes {
	return springCloudCustomizedAcceleratorAttributes{ref: terra.ReferenceResource(scca)}
}

// ImportState imports the given attribute values into [SpringCloudCustomizedAccelerator]'s state.
func (scca *SpringCloudCustomizedAccelerator) ImportState(av io.Reader) error {
	scca.state = &springCloudCustomizedAcceleratorState{}
	if err := json.NewDecoder(av).Decode(scca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scca.Type(), scca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudCustomizedAccelerator] has state.
func (scca *SpringCloudCustomizedAccelerator) State() (*springCloudCustomizedAcceleratorState, bool) {
	return scca.state, scca.state != nil
}

// StateMust returns the state for [SpringCloudCustomizedAccelerator]. Panics if the state is nil.
func (scca *SpringCloudCustomizedAccelerator) StateMust() *springCloudCustomizedAcceleratorState {
	if scca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scca.Type(), scca.LocalName()))
	}
	return scca.state
}

// SpringCloudCustomizedAcceleratorArgs contains the configurations for azurerm_spring_cloud_customized_accelerator.
type SpringCloudCustomizedAcceleratorArgs struct {
	// AcceleratorTags: list of string, optional
	AcceleratorTags terra.ListValue[terra.StringValue] `hcl:"accelerator_tags,attr"`
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
	GitRepository *springcloudcustomizedaccelerator.GitRepository `hcl:"git_repository,block" validate:"required"`
	// Timeouts: optional
	Timeouts *springcloudcustomizedaccelerator.Timeouts `hcl:"timeouts,block"`
}
type springCloudCustomizedAcceleratorAttributes struct {
	ref terra.Reference
}

// AcceleratorTags returns a reference to field accelerator_tags of azurerm_spring_cloud_customized_accelerator.
func (scca springCloudCustomizedAcceleratorAttributes) AcceleratorTags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](scca.ref.Append("accelerator_tags"))
}

// Description returns a reference to field description of azurerm_spring_cloud_customized_accelerator.
func (scca springCloudCustomizedAcceleratorAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(scca.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_spring_cloud_customized_accelerator.
func (scca springCloudCustomizedAcceleratorAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(scca.ref.Append("display_name"))
}

// IconUrl returns a reference to field icon_url of azurerm_spring_cloud_customized_accelerator.
func (scca springCloudCustomizedAcceleratorAttributes) IconUrl() terra.StringValue {
	return terra.ReferenceAsString(scca.ref.Append("icon_url"))
}

// Id returns a reference to field id of azurerm_spring_cloud_customized_accelerator.
func (scca springCloudCustomizedAcceleratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scca.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_customized_accelerator.
func (scca springCloudCustomizedAcceleratorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scca.ref.Append("name"))
}

// SpringCloudAcceleratorId returns a reference to field spring_cloud_accelerator_id of azurerm_spring_cloud_customized_accelerator.
func (scca springCloudCustomizedAcceleratorAttributes) SpringCloudAcceleratorId() terra.StringValue {
	return terra.ReferenceAsString(scca.ref.Append("spring_cloud_accelerator_id"))
}

func (scca springCloudCustomizedAcceleratorAttributes) GitRepository() terra.ListValue[springcloudcustomizedaccelerator.GitRepositoryAttributes] {
	return terra.ReferenceAsList[springcloudcustomizedaccelerator.GitRepositoryAttributes](scca.ref.Append("git_repository"))
}

func (scca springCloudCustomizedAcceleratorAttributes) Timeouts() springcloudcustomizedaccelerator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudcustomizedaccelerator.TimeoutsAttributes](scca.ref.Append("timeouts"))
}

type springCloudCustomizedAcceleratorState struct {
	AcceleratorTags          []string                                              `json:"accelerator_tags"`
	Description              string                                                `json:"description"`
	DisplayName              string                                                `json:"display_name"`
	IconUrl                  string                                                `json:"icon_url"`
	Id                       string                                                `json:"id"`
	Name                     string                                                `json:"name"`
	SpringCloudAcceleratorId string                                                `json:"spring_cloud_accelerator_id"`
	GitRepository            []springcloudcustomizedaccelerator.GitRepositoryState `json:"git_repository"`
	Timeouts                 *springcloudcustomizedaccelerator.TimeoutsState       `json:"timeouts"`
}
