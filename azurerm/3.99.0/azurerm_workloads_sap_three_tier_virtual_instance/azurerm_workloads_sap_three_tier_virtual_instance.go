// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_workloads_sap_three_tier_virtual_instance

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

// Resource represents the Terraform resource azurerm_workloads_sap_three_tier_virtual_instance.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermWorkloadsSapThreeTierVirtualInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (awsttvi *Resource) Type() string {
	return "azurerm_workloads_sap_three_tier_virtual_instance"
}

// LocalName returns the local name for [Resource].
func (awsttvi *Resource) LocalName() string {
	return awsttvi.Name
}

// Configuration returns the configuration (args) for [Resource].
func (awsttvi *Resource) Configuration() interface{} {
	return awsttvi.Args
}

// DependOn is used for other resources to depend on [Resource].
func (awsttvi *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(awsttvi)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (awsttvi *Resource) Dependencies() terra.Dependencies {
	return awsttvi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (awsttvi *Resource) LifecycleManagement() *terra.Lifecycle {
	return awsttvi.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (awsttvi *Resource) Attributes() azurermWorkloadsSapThreeTierVirtualInstanceAttributes {
	return azurermWorkloadsSapThreeTierVirtualInstanceAttributes{ref: terra.ReferenceResource(awsttvi)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (awsttvi *Resource) ImportState(state io.Reader) error {
	awsttvi.state = &azurermWorkloadsSapThreeTierVirtualInstanceState{}
	if err := json.NewDecoder(state).Decode(awsttvi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", awsttvi.Type(), awsttvi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (awsttvi *Resource) State() (*azurermWorkloadsSapThreeTierVirtualInstanceState, bool) {
	return awsttvi.state, awsttvi.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (awsttvi *Resource) StateMust() *azurermWorkloadsSapThreeTierVirtualInstanceState {
	if awsttvi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", awsttvi.Type(), awsttvi.LocalName()))
	}
	return awsttvi.state
}

// Args contains the configurations for azurerm_workloads_sap_three_tier_virtual_instance.
type Args struct {
	// AppLocation: string, required
	AppLocation terra.StringValue `hcl:"app_location,attr" validate:"required"`
	// Environment: string, required
	Environment terra.StringValue `hcl:"environment,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedResourceGroupName: string, optional
	ManagedResourceGroupName terra.StringValue `hcl:"managed_resource_group_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SapFqdn: string, required
	SapFqdn terra.StringValue `hcl:"sap_fqdn,attr" validate:"required"`
	// SapProduct: string, required
	SapProduct terra.StringValue `hcl:"sap_product,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// ThreeTierConfiguration: required
	ThreeTierConfiguration *ThreeTierConfiguration `hcl:"three_tier_configuration,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermWorkloadsSapThreeTierVirtualInstanceAttributes struct {
	ref terra.Reference
}

// AppLocation returns a reference to field app_location of azurerm_workloads_sap_three_tier_virtual_instance.
func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) AppLocation() terra.StringValue {
	return terra.ReferenceAsString(awsttvi.ref.Append("app_location"))
}

// Environment returns a reference to field environment of azurerm_workloads_sap_three_tier_virtual_instance.
func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(awsttvi.ref.Append("environment"))
}

// Id returns a reference to field id of azurerm_workloads_sap_three_tier_virtual_instance.
func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(awsttvi.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_workloads_sap_three_tier_virtual_instance.
func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(awsttvi.ref.Append("location"))
}

// ManagedResourceGroupName returns a reference to field managed_resource_group_name of azurerm_workloads_sap_three_tier_virtual_instance.
func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) ManagedResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(awsttvi.ref.Append("managed_resource_group_name"))
}

// Name returns a reference to field name of azurerm_workloads_sap_three_tier_virtual_instance.
func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(awsttvi.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_workloads_sap_three_tier_virtual_instance.
func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(awsttvi.ref.Append("resource_group_name"))
}

// SapFqdn returns a reference to field sap_fqdn of azurerm_workloads_sap_three_tier_virtual_instance.
func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) SapFqdn() terra.StringValue {
	return terra.ReferenceAsString(awsttvi.ref.Append("sap_fqdn"))
}

// SapProduct returns a reference to field sap_product of azurerm_workloads_sap_three_tier_virtual_instance.
func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) SapProduct() terra.StringValue {
	return terra.ReferenceAsString(awsttvi.ref.Append("sap_product"))
}

// Tags returns a reference to field tags of azurerm_workloads_sap_three_tier_virtual_instance.
func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](awsttvi.ref.Append("tags"))
}

func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](awsttvi.ref.Append("identity"))
}

func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) ThreeTierConfiguration() terra.ListValue[ThreeTierConfigurationAttributes] {
	return terra.ReferenceAsList[ThreeTierConfigurationAttributes](awsttvi.ref.Append("three_tier_configuration"))
}

func (awsttvi azurermWorkloadsSapThreeTierVirtualInstanceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](awsttvi.ref.Append("timeouts"))
}

type azurermWorkloadsSapThreeTierVirtualInstanceState struct {
	AppLocation              string                        `json:"app_location"`
	Environment              string                        `json:"environment"`
	Id                       string                        `json:"id"`
	Location                 string                        `json:"location"`
	ManagedResourceGroupName string                        `json:"managed_resource_group_name"`
	Name                     string                        `json:"name"`
	ResourceGroupName        string                        `json:"resource_group_name"`
	SapFqdn                  string                        `json:"sap_fqdn"`
	SapProduct               string                        `json:"sap_product"`
	Tags                     map[string]string             `json:"tags"`
	Identity                 []IdentityState               `json:"identity"`
	ThreeTierConfiguration   []ThreeTierConfigurationState `json:"three_tier_configuration"`
	Timeouts                 *TimeoutsState                `json:"timeouts"`
}
