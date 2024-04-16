// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_custom_ip_prefix

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

// Resource represents the Terraform resource azurerm_custom_ip_prefix.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermCustomIpPrefixState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acip *Resource) Type() string {
	return "azurerm_custom_ip_prefix"
}

// LocalName returns the local name for [Resource].
func (acip *Resource) LocalName() string {
	return acip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acip *Resource) Configuration() interface{} {
	return acip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acip *Resource) Dependencies() terra.Dependencies {
	return acip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acip *Resource) LifecycleManagement() *terra.Lifecycle {
	return acip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acip *Resource) Attributes() azurermCustomIpPrefixAttributes {
	return azurermCustomIpPrefixAttributes{ref: terra.ReferenceResource(acip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acip *Resource) ImportState(state io.Reader) error {
	acip.state = &azurermCustomIpPrefixState{}
	if err := json.NewDecoder(state).Decode(acip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acip.Type(), acip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acip *Resource) State() (*azurermCustomIpPrefixState, bool) {
	return acip.state, acip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acip *Resource) StateMust() *azurermCustomIpPrefixState {
	if acip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acip.Type(), acip.LocalName()))
	}
	return acip.state
}

// Args contains the configurations for azurerm_custom_ip_prefix.
type Args struct {
	// Cidr: string, required
	Cidr terra.StringValue `hcl:"cidr,attr" validate:"required"`
	// CommissioningEnabled: bool, optional
	CommissioningEnabled terra.BoolValue `hcl:"commissioning_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InternetAdvertisingDisabled: bool, optional
	InternetAdvertisingDisabled terra.BoolValue `hcl:"internet_advertising_disabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentCustomIpPrefixId: string, optional
	ParentCustomIpPrefixId terra.StringValue `hcl:"parent_custom_ip_prefix_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RoaValidityEndDate: string, optional
	RoaValidityEndDate terra.StringValue `hcl:"roa_validity_end_date,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WanValidationSignedMessage: string, optional
	WanValidationSignedMessage terra.StringValue `hcl:"wan_validation_signed_message,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermCustomIpPrefixAttributes struct {
	ref terra.Reference
}

// Cidr returns a reference to field cidr of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(acip.ref.Append("cidr"))
}

// CommissioningEnabled returns a reference to field commissioning_enabled of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) CommissioningEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(acip.ref.Append("commissioning_enabled"))
}

// Id returns a reference to field id of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acip.ref.Append("id"))
}

// InternetAdvertisingDisabled returns a reference to field internet_advertising_disabled of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) InternetAdvertisingDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(acip.ref.Append("internet_advertising_disabled"))
}

// Location returns a reference to field location of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(acip.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acip.ref.Append("name"))
}

// ParentCustomIpPrefixId returns a reference to field parent_custom_ip_prefix_id of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) ParentCustomIpPrefixId() terra.StringValue {
	return terra.ReferenceAsString(acip.ref.Append("parent_custom_ip_prefix_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(acip.ref.Append("resource_group_name"))
}

// RoaValidityEndDate returns a reference to field roa_validity_end_date of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) RoaValidityEndDate() terra.StringValue {
	return terra.ReferenceAsString(acip.ref.Append("roa_validity_end_date"))
}

// Tags returns a reference to field tags of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acip.ref.Append("tags"))
}

// WanValidationSignedMessage returns a reference to field wan_validation_signed_message of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) WanValidationSignedMessage() terra.StringValue {
	return terra.ReferenceAsString(acip.ref.Append("wan_validation_signed_message"))
}

// Zones returns a reference to field zones of azurerm_custom_ip_prefix.
func (acip azurermCustomIpPrefixAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acip.ref.Append("zones"))
}

func (acip azurermCustomIpPrefixAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acip.ref.Append("timeouts"))
}

type azurermCustomIpPrefixState struct {
	Cidr                        string            `json:"cidr"`
	CommissioningEnabled        bool              `json:"commissioning_enabled"`
	Id                          string            `json:"id"`
	InternetAdvertisingDisabled bool              `json:"internet_advertising_disabled"`
	Location                    string            `json:"location"`
	Name                        string            `json:"name"`
	ParentCustomIpPrefixId      string            `json:"parent_custom_ip_prefix_id"`
	ResourceGroupName           string            `json:"resource_group_name"`
	RoaValidityEndDate          string            `json:"roa_validity_end_date"`
	Tags                        map[string]string `json:"tags"`
	WanValidationSignedMessage  string            `json:"wan_validation_signed_message"`
	Zones                       []string          `json:"zones"`
	Timeouts                    *TimeoutsState    `json:"timeouts"`
}
