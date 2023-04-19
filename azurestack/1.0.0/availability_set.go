// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurestack

import (
	"encoding/json"
	"fmt"
	availabilityset "github.com/golingon/terraproviders/azurestack/1.0.0/availabilityset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAvailabilitySet creates a new instance of [AvailabilitySet].
func NewAvailabilitySet(name string, args AvailabilitySetArgs) *AvailabilitySet {
	return &AvailabilitySet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AvailabilitySet)(nil)

// AvailabilitySet represents the Terraform resource azurestack_availability_set.
type AvailabilitySet struct {
	Name      string
	Args      AvailabilitySetArgs
	state     *availabilitySetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AvailabilitySet].
func (as *AvailabilitySet) Type() string {
	return "azurestack_availability_set"
}

// LocalName returns the local name for [AvailabilitySet].
func (as *AvailabilitySet) LocalName() string {
	return as.Name
}

// Configuration returns the configuration (args) for [AvailabilitySet].
func (as *AvailabilitySet) Configuration() interface{} {
	return as.Args
}

// DependOn is used for other resources to depend on [AvailabilitySet].
func (as *AvailabilitySet) DependOn() terra.Reference {
	return terra.ReferenceResource(as)
}

// Dependencies returns the list of resources [AvailabilitySet] depends_on.
func (as *AvailabilitySet) Dependencies() terra.Dependencies {
	return as.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AvailabilitySet].
func (as *AvailabilitySet) LifecycleManagement() *terra.Lifecycle {
	return as.Lifecycle
}

// Attributes returns the attributes for [AvailabilitySet].
func (as *AvailabilitySet) Attributes() availabilitySetAttributes {
	return availabilitySetAttributes{ref: terra.ReferenceResource(as)}
}

// ImportState imports the given attribute values into [AvailabilitySet]'s state.
func (as *AvailabilitySet) ImportState(av io.Reader) error {
	as.state = &availabilitySetState{}
	if err := json.NewDecoder(av).Decode(as.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", as.Type(), as.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AvailabilitySet] has state.
func (as *AvailabilitySet) State() (*availabilitySetState, bool) {
	return as.state, as.state != nil
}

// StateMust returns the state for [AvailabilitySet]. Panics if the state is nil.
func (as *AvailabilitySet) StateMust() *availabilitySetState {
	if as.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", as.Type(), as.LocalName()))
	}
	return as.state
}

// AvailabilitySetArgs contains the configurations for azurestack_availability_set.
type AvailabilitySetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Managed: bool, optional
	Managed terra.BoolValue `hcl:"managed,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PlatformFaultDomainCount: number, optional
	PlatformFaultDomainCount terra.NumberValue `hcl:"platform_fault_domain_count,attr"`
	// PlatformUpdateDomainCount: number, optional
	PlatformUpdateDomainCount terra.NumberValue `hcl:"platform_update_domain_count,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *availabilityset.Timeouts `hcl:"timeouts,block"`
}
type availabilitySetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurestack_availability_set.
func (as availabilitySetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("id"))
}

// Location returns a reference to field location of azurestack_availability_set.
func (as availabilitySetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("location"))
}

// Managed returns a reference to field managed of azurestack_availability_set.
func (as availabilitySetAttributes) Managed() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("managed"))
}

// Name returns a reference to field name of azurestack_availability_set.
func (as availabilitySetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("name"))
}

// PlatformFaultDomainCount returns a reference to field platform_fault_domain_count of azurestack_availability_set.
func (as availabilitySetAttributes) PlatformFaultDomainCount() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("platform_fault_domain_count"))
}

// PlatformUpdateDomainCount returns a reference to field platform_update_domain_count of azurestack_availability_set.
func (as availabilitySetAttributes) PlatformUpdateDomainCount() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("platform_update_domain_count"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_availability_set.
func (as availabilitySetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurestack_availability_set.
func (as availabilitySetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](as.ref.Append("tags"))
}

func (as availabilitySetAttributes) Timeouts() availabilityset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[availabilityset.TimeoutsAttributes](as.ref.Append("timeouts"))
}

type availabilitySetState struct {
	Id                        string                         `json:"id"`
	Location                  string                         `json:"location"`
	Managed                   bool                           `json:"managed"`
	Name                      string                         `json:"name"`
	PlatformFaultDomainCount  float64                        `json:"platform_fault_domain_count"`
	PlatformUpdateDomainCount float64                        `json:"platform_update_domain_count"`
	ResourceGroupName         string                         `json:"resource_group_name"`
	Tags                      map[string]string              `json:"tags"`
	Timeouts                  *availabilityset.TimeoutsState `json:"timeouts"`
}
