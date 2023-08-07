// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	resourceproviderregistration "github.com/golingon/terraproviders/azurerm/3.68.0/resourceproviderregistration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourceProviderRegistration creates a new instance of [ResourceProviderRegistration].
func NewResourceProviderRegistration(name string, args ResourceProviderRegistrationArgs) *ResourceProviderRegistration {
	return &ResourceProviderRegistration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceProviderRegistration)(nil)

// ResourceProviderRegistration represents the Terraform resource azurerm_resource_provider_registration.
type ResourceProviderRegistration struct {
	Name      string
	Args      ResourceProviderRegistrationArgs
	state     *resourceProviderRegistrationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourceProviderRegistration].
func (rpr *ResourceProviderRegistration) Type() string {
	return "azurerm_resource_provider_registration"
}

// LocalName returns the local name for [ResourceProviderRegistration].
func (rpr *ResourceProviderRegistration) LocalName() string {
	return rpr.Name
}

// Configuration returns the configuration (args) for [ResourceProviderRegistration].
func (rpr *ResourceProviderRegistration) Configuration() interface{} {
	return rpr.Args
}

// DependOn is used for other resources to depend on [ResourceProviderRegistration].
func (rpr *ResourceProviderRegistration) DependOn() terra.Reference {
	return terra.ReferenceResource(rpr)
}

// Dependencies returns the list of resources [ResourceProviderRegistration] depends_on.
func (rpr *ResourceProviderRegistration) Dependencies() terra.Dependencies {
	return rpr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourceProviderRegistration].
func (rpr *ResourceProviderRegistration) LifecycleManagement() *terra.Lifecycle {
	return rpr.Lifecycle
}

// Attributes returns the attributes for [ResourceProviderRegistration].
func (rpr *ResourceProviderRegistration) Attributes() resourceProviderRegistrationAttributes {
	return resourceProviderRegistrationAttributes{ref: terra.ReferenceResource(rpr)}
}

// ImportState imports the given attribute values into [ResourceProviderRegistration]'s state.
func (rpr *ResourceProviderRegistration) ImportState(av io.Reader) error {
	rpr.state = &resourceProviderRegistrationState{}
	if err := json.NewDecoder(av).Decode(rpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rpr.Type(), rpr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourceProviderRegistration] has state.
func (rpr *ResourceProviderRegistration) State() (*resourceProviderRegistrationState, bool) {
	return rpr.state, rpr.state != nil
}

// StateMust returns the state for [ResourceProviderRegistration]. Panics if the state is nil.
func (rpr *ResourceProviderRegistration) StateMust() *resourceProviderRegistrationState {
	if rpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rpr.Type(), rpr.LocalName()))
	}
	return rpr.state
}

// ResourceProviderRegistrationArgs contains the configurations for azurerm_resource_provider_registration.
type ResourceProviderRegistrationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Feature: min=0
	Feature []resourceproviderregistration.Feature `hcl:"feature,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *resourceproviderregistration.Timeouts `hcl:"timeouts,block"`
}
type resourceProviderRegistrationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_resource_provider_registration.
func (rpr resourceProviderRegistrationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rpr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_resource_provider_registration.
func (rpr resourceProviderRegistrationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rpr.ref.Append("name"))
}

func (rpr resourceProviderRegistrationAttributes) Feature() terra.SetValue[resourceproviderregistration.FeatureAttributes] {
	return terra.ReferenceAsSet[resourceproviderregistration.FeatureAttributes](rpr.ref.Append("feature"))
}

func (rpr resourceProviderRegistrationAttributes) Timeouts() resourceproviderregistration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourceproviderregistration.TimeoutsAttributes](rpr.ref.Append("timeouts"))
}

type resourceProviderRegistrationState struct {
	Id       string                                      `json:"id"`
	Name     string                                      `json:"name"`
	Feature  []resourceproviderregistration.FeatureState `json:"feature"`
	Timeouts *resourceproviderregistration.TimeoutsState `json:"timeouts"`
}
