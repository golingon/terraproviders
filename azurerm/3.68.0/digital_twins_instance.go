// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	digitaltwinsinstance "github.com/golingon/terraproviders/azurerm/3.68.0/digitaltwinsinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDigitalTwinsInstance creates a new instance of [DigitalTwinsInstance].
func NewDigitalTwinsInstance(name string, args DigitalTwinsInstanceArgs) *DigitalTwinsInstance {
	return &DigitalTwinsInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DigitalTwinsInstance)(nil)

// DigitalTwinsInstance represents the Terraform resource azurerm_digital_twins_instance.
type DigitalTwinsInstance struct {
	Name      string
	Args      DigitalTwinsInstanceArgs
	state     *digitalTwinsInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DigitalTwinsInstance].
func (dti *DigitalTwinsInstance) Type() string {
	return "azurerm_digital_twins_instance"
}

// LocalName returns the local name for [DigitalTwinsInstance].
func (dti *DigitalTwinsInstance) LocalName() string {
	return dti.Name
}

// Configuration returns the configuration (args) for [DigitalTwinsInstance].
func (dti *DigitalTwinsInstance) Configuration() interface{} {
	return dti.Args
}

// DependOn is used for other resources to depend on [DigitalTwinsInstance].
func (dti *DigitalTwinsInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(dti)
}

// Dependencies returns the list of resources [DigitalTwinsInstance] depends_on.
func (dti *DigitalTwinsInstance) Dependencies() terra.Dependencies {
	return dti.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DigitalTwinsInstance].
func (dti *DigitalTwinsInstance) LifecycleManagement() *terra.Lifecycle {
	return dti.Lifecycle
}

// Attributes returns the attributes for [DigitalTwinsInstance].
func (dti *DigitalTwinsInstance) Attributes() digitalTwinsInstanceAttributes {
	return digitalTwinsInstanceAttributes{ref: terra.ReferenceResource(dti)}
}

// ImportState imports the given attribute values into [DigitalTwinsInstance]'s state.
func (dti *DigitalTwinsInstance) ImportState(av io.Reader) error {
	dti.state = &digitalTwinsInstanceState{}
	if err := json.NewDecoder(av).Decode(dti.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dti.Type(), dti.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DigitalTwinsInstance] has state.
func (dti *DigitalTwinsInstance) State() (*digitalTwinsInstanceState, bool) {
	return dti.state, dti.state != nil
}

// StateMust returns the state for [DigitalTwinsInstance]. Panics if the state is nil.
func (dti *DigitalTwinsInstance) StateMust() *digitalTwinsInstanceState {
	if dti.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dti.Type(), dti.LocalName()))
	}
	return dti.state
}

// DigitalTwinsInstanceArgs contains the configurations for azurerm_digital_twins_instance.
type DigitalTwinsInstanceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *digitaltwinsinstance.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *digitaltwinsinstance.Timeouts `hcl:"timeouts,block"`
}
type digitalTwinsInstanceAttributes struct {
	ref terra.Reference
}

// HostName returns a reference to field host_name of azurerm_digital_twins_instance.
func (dti digitalTwinsInstanceAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("host_name"))
}

// Id returns a reference to field id of azurerm_digital_twins_instance.
func (dti digitalTwinsInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_digital_twins_instance.
func (dti digitalTwinsInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_digital_twins_instance.
func (dti digitalTwinsInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_digital_twins_instance.
func (dti digitalTwinsInstanceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_digital_twins_instance.
func (dti digitalTwinsInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dti.ref.Append("tags"))
}

func (dti digitalTwinsInstanceAttributes) Identity() terra.ListValue[digitaltwinsinstance.IdentityAttributes] {
	return terra.ReferenceAsList[digitaltwinsinstance.IdentityAttributes](dti.ref.Append("identity"))
}

func (dti digitalTwinsInstanceAttributes) Timeouts() digitaltwinsinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[digitaltwinsinstance.TimeoutsAttributes](dti.ref.Append("timeouts"))
}

type digitalTwinsInstanceState struct {
	HostName          string                               `json:"host_name"`
	Id                string                               `json:"id"`
	Location          string                               `json:"location"`
	Name              string                               `json:"name"`
	ResourceGroupName string                               `json:"resource_group_name"`
	Tags              map[string]string                    `json:"tags"`
	Identity          []digitaltwinsinstance.IdentityState `json:"identity"`
	Timeouts          *digitaltwinsinstance.TimeoutsState  `json:"timeouts"`
}
