// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	customipprefix "github.com/golingon/terraproviders/azurerm/3.69.0/customipprefix"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCustomIpPrefix creates a new instance of [CustomIpPrefix].
func NewCustomIpPrefix(name string, args CustomIpPrefixArgs) *CustomIpPrefix {
	return &CustomIpPrefix{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CustomIpPrefix)(nil)

// CustomIpPrefix represents the Terraform resource azurerm_custom_ip_prefix.
type CustomIpPrefix struct {
	Name      string
	Args      CustomIpPrefixArgs
	state     *customIpPrefixState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CustomIpPrefix].
func (cip *CustomIpPrefix) Type() string {
	return "azurerm_custom_ip_prefix"
}

// LocalName returns the local name for [CustomIpPrefix].
func (cip *CustomIpPrefix) LocalName() string {
	return cip.Name
}

// Configuration returns the configuration (args) for [CustomIpPrefix].
func (cip *CustomIpPrefix) Configuration() interface{} {
	return cip.Args
}

// DependOn is used for other resources to depend on [CustomIpPrefix].
func (cip *CustomIpPrefix) DependOn() terra.Reference {
	return terra.ReferenceResource(cip)
}

// Dependencies returns the list of resources [CustomIpPrefix] depends_on.
func (cip *CustomIpPrefix) Dependencies() terra.Dependencies {
	return cip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CustomIpPrefix].
func (cip *CustomIpPrefix) LifecycleManagement() *terra.Lifecycle {
	return cip.Lifecycle
}

// Attributes returns the attributes for [CustomIpPrefix].
func (cip *CustomIpPrefix) Attributes() customIpPrefixAttributes {
	return customIpPrefixAttributes{ref: terra.ReferenceResource(cip)}
}

// ImportState imports the given attribute values into [CustomIpPrefix]'s state.
func (cip *CustomIpPrefix) ImportState(av io.Reader) error {
	cip.state = &customIpPrefixState{}
	if err := json.NewDecoder(av).Decode(cip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cip.Type(), cip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CustomIpPrefix] has state.
func (cip *CustomIpPrefix) State() (*customIpPrefixState, bool) {
	return cip.state, cip.state != nil
}

// StateMust returns the state for [CustomIpPrefix]. Panics if the state is nil.
func (cip *CustomIpPrefix) StateMust() *customIpPrefixState {
	if cip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cip.Type(), cip.LocalName()))
	}
	return cip.state
}

// CustomIpPrefixArgs contains the configurations for azurerm_custom_ip_prefix.
type CustomIpPrefixArgs struct {
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
	Timeouts *customipprefix.Timeouts `hcl:"timeouts,block"`
}
type customIpPrefixAttributes struct {
	ref terra.Reference
}

// Cidr returns a reference to field cidr of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("cidr"))
}

// CommissioningEnabled returns a reference to field commissioning_enabled of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) CommissioningEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cip.ref.Append("commissioning_enabled"))
}

// Id returns a reference to field id of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("id"))
}

// InternetAdvertisingDisabled returns a reference to field internet_advertising_disabled of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) InternetAdvertisingDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(cip.ref.Append("internet_advertising_disabled"))
}

// Location returns a reference to field location of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("name"))
}

// ParentCustomIpPrefixId returns a reference to field parent_custom_ip_prefix_id of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) ParentCustomIpPrefixId() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("parent_custom_ip_prefix_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("resource_group_name"))
}

// RoaValidityEndDate returns a reference to field roa_validity_end_date of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) RoaValidityEndDate() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("roa_validity_end_date"))
}

// Tags returns a reference to field tags of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cip.ref.Append("tags"))
}

// WanValidationSignedMessage returns a reference to field wan_validation_signed_message of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) WanValidationSignedMessage() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("wan_validation_signed_message"))
}

// Zones returns a reference to field zones of azurerm_custom_ip_prefix.
func (cip customIpPrefixAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cip.ref.Append("zones"))
}

func (cip customIpPrefixAttributes) Timeouts() customipprefix.TimeoutsAttributes {
	return terra.ReferenceAsSingle[customipprefix.TimeoutsAttributes](cip.ref.Append("timeouts"))
}

type customIpPrefixState struct {
	Cidr                        string                        `json:"cidr"`
	CommissioningEnabled        bool                          `json:"commissioning_enabled"`
	Id                          string                        `json:"id"`
	InternetAdvertisingDisabled bool                          `json:"internet_advertising_disabled"`
	Location                    string                        `json:"location"`
	Name                        string                        `json:"name"`
	ParentCustomIpPrefixId      string                        `json:"parent_custom_ip_prefix_id"`
	ResourceGroupName           string                        `json:"resource_group_name"`
	RoaValidityEndDate          string                        `json:"roa_validity_end_date"`
	Tags                        map[string]string             `json:"tags"`
	WanValidationSignedMessage  string                        `json:"wan_validation_signed_message"`
	Zones                       []string                      `json:"zones"`
	Timeouts                    *customipprefix.TimeoutsState `json:"timeouts"`
}