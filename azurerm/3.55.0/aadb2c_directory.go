// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	aadb2cdirectory "github.com/golingon/terraproviders/azurerm/3.55.0/aadb2cdirectory"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAadb2CDirectory creates a new instance of [Aadb2CDirectory].
func NewAadb2CDirectory(name string, args Aadb2CDirectoryArgs) *Aadb2CDirectory {
	return &Aadb2CDirectory{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Aadb2CDirectory)(nil)

// Aadb2CDirectory represents the Terraform resource azurerm_aadb2c_directory.
type Aadb2CDirectory struct {
	Name      string
	Args      Aadb2CDirectoryArgs
	state     *aadb2CDirectoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Aadb2CDirectory].
func (ad *Aadb2CDirectory) Type() string {
	return "azurerm_aadb2c_directory"
}

// LocalName returns the local name for [Aadb2CDirectory].
func (ad *Aadb2CDirectory) LocalName() string {
	return ad.Name
}

// Configuration returns the configuration (args) for [Aadb2CDirectory].
func (ad *Aadb2CDirectory) Configuration() interface{} {
	return ad.Args
}

// DependOn is used for other resources to depend on [Aadb2CDirectory].
func (ad *Aadb2CDirectory) DependOn() terra.Reference {
	return terra.ReferenceResource(ad)
}

// Dependencies returns the list of resources [Aadb2CDirectory] depends_on.
func (ad *Aadb2CDirectory) Dependencies() terra.Dependencies {
	return ad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Aadb2CDirectory].
func (ad *Aadb2CDirectory) LifecycleManagement() *terra.Lifecycle {
	return ad.Lifecycle
}

// Attributes returns the attributes for [Aadb2CDirectory].
func (ad *Aadb2CDirectory) Attributes() aadb2CDirectoryAttributes {
	return aadb2CDirectoryAttributes{ref: terra.ReferenceResource(ad)}
}

// ImportState imports the given attribute values into [Aadb2CDirectory]'s state.
func (ad *Aadb2CDirectory) ImportState(av io.Reader) error {
	ad.state = &aadb2CDirectoryState{}
	if err := json.NewDecoder(av).Decode(ad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ad.Type(), ad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Aadb2CDirectory] has state.
func (ad *Aadb2CDirectory) State() (*aadb2CDirectoryState, bool) {
	return ad.state, ad.state != nil
}

// StateMust returns the state for [Aadb2CDirectory]. Panics if the state is nil.
func (ad *Aadb2CDirectory) StateMust() *aadb2CDirectoryState {
	if ad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ad.Type(), ad.LocalName()))
	}
	return ad.state
}

// Aadb2CDirectoryArgs contains the configurations for azurerm_aadb2c_directory.
type Aadb2CDirectoryArgs struct {
	// CountryCode: string, optional
	CountryCode terra.StringValue `hcl:"country_code,attr"`
	// DataResidencyLocation: string, required
	DataResidencyLocation terra.StringValue `hcl:"data_residency_location,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *aadb2cdirectory.Timeouts `hcl:"timeouts,block"`
}
type aadb2CDirectoryAttributes struct {
	ref terra.Reference
}

// BillingType returns a reference to field billing_type of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) BillingType() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("billing_type"))
}

// CountryCode returns a reference to field country_code of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("country_code"))
}

// DataResidencyLocation returns a reference to field data_residency_location of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) DataResidencyLocation() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("data_residency_location"))
}

// DisplayName returns a reference to field display_name of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("display_name"))
}

// DomainName returns a reference to field domain_name of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("domain_name"))
}

// EffectiveStartDate returns a reference to field effective_start_date of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) EffectiveStartDate() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("effective_start_date"))
}

// Id returns a reference to field id of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ad.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_aadb2c_directory.
func (ad aadb2CDirectoryAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("tenant_id"))
}

func (ad aadb2CDirectoryAttributes) Timeouts() aadb2cdirectory.TimeoutsAttributes {
	return terra.ReferenceAsSingle[aadb2cdirectory.TimeoutsAttributes](ad.ref.Append("timeouts"))
}

type aadb2CDirectoryState struct {
	BillingType           string                         `json:"billing_type"`
	CountryCode           string                         `json:"country_code"`
	DataResidencyLocation string                         `json:"data_residency_location"`
	DisplayName           string                         `json:"display_name"`
	DomainName            string                         `json:"domain_name"`
	EffectiveStartDate    string                         `json:"effective_start_date"`
	Id                    string                         `json:"id"`
	ResourceGroupName     string                         `json:"resource_group_name"`
	SkuName               string                         `json:"sku_name"`
	Tags                  map[string]string              `json:"tags"`
	TenantId              string                         `json:"tenant_id"`
	Timeouts              *aadb2cdirectory.TimeoutsState `json:"timeouts"`
}
