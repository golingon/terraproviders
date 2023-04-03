// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudaccelerator "github.com/golingon/terraproviders/azurerm/3.49.0/springcloudaccelerator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudAccelerator creates a new instance of [SpringCloudAccelerator].
func NewSpringCloudAccelerator(name string, args SpringCloudAcceleratorArgs) *SpringCloudAccelerator {
	return &SpringCloudAccelerator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudAccelerator)(nil)

// SpringCloudAccelerator represents the Terraform resource azurerm_spring_cloud_accelerator.
type SpringCloudAccelerator struct {
	Name      string
	Args      SpringCloudAcceleratorArgs
	state     *springCloudAcceleratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudAccelerator].
func (sca *SpringCloudAccelerator) Type() string {
	return "azurerm_spring_cloud_accelerator"
}

// LocalName returns the local name for [SpringCloudAccelerator].
func (sca *SpringCloudAccelerator) LocalName() string {
	return sca.Name
}

// Configuration returns the configuration (args) for [SpringCloudAccelerator].
func (sca *SpringCloudAccelerator) Configuration() interface{} {
	return sca.Args
}

// DependOn is used for other resources to depend on [SpringCloudAccelerator].
func (sca *SpringCloudAccelerator) DependOn() terra.Reference {
	return terra.ReferenceResource(sca)
}

// Dependencies returns the list of resources [SpringCloudAccelerator] depends_on.
func (sca *SpringCloudAccelerator) Dependencies() terra.Dependencies {
	return sca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudAccelerator].
func (sca *SpringCloudAccelerator) LifecycleManagement() *terra.Lifecycle {
	return sca.Lifecycle
}

// Attributes returns the attributes for [SpringCloudAccelerator].
func (sca *SpringCloudAccelerator) Attributes() springCloudAcceleratorAttributes {
	return springCloudAcceleratorAttributes{ref: terra.ReferenceResource(sca)}
}

// ImportState imports the given attribute values into [SpringCloudAccelerator]'s state.
func (sca *SpringCloudAccelerator) ImportState(av io.Reader) error {
	sca.state = &springCloudAcceleratorState{}
	if err := json.NewDecoder(av).Decode(sca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sca.Type(), sca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudAccelerator] has state.
func (sca *SpringCloudAccelerator) State() (*springCloudAcceleratorState, bool) {
	return sca.state, sca.state != nil
}

// StateMust returns the state for [SpringCloudAccelerator]. Panics if the state is nil.
func (sca *SpringCloudAccelerator) StateMust() *springCloudAcceleratorState {
	if sca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sca.Type(), sca.LocalName()))
	}
	return sca.state
}

// SpringCloudAcceleratorArgs contains the configurations for azurerm_spring_cloud_accelerator.
type SpringCloudAcceleratorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudServiceId: string, required
	SpringCloudServiceId terra.StringValue `hcl:"spring_cloud_service_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *springcloudaccelerator.Timeouts `hcl:"timeouts,block"`
}
type springCloudAcceleratorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_spring_cloud_accelerator.
func (sca springCloudAcceleratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_accelerator.
func (sca springCloudAcceleratorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("name"))
}

// SpringCloudServiceId returns a reference to field spring_cloud_service_id of azurerm_spring_cloud_accelerator.
func (sca springCloudAcceleratorAttributes) SpringCloudServiceId() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("spring_cloud_service_id"))
}

func (sca springCloudAcceleratorAttributes) Timeouts() springcloudaccelerator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudaccelerator.TimeoutsAttributes](sca.ref.Append("timeouts"))
}

type springCloudAcceleratorState struct {
	Id                   string                                `json:"id"`
	Name                 string                                `json:"name"`
	SpringCloudServiceId string                                `json:"spring_cloud_service_id"`
	Timeouts             *springcloudaccelerator.TimeoutsState `json:"timeouts"`
}
