// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacdnprofile "github.com/golingon/terraproviders/azurerm/3.68.0/datacdnprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCdnProfile creates a new instance of [DataCdnProfile].
func NewDataCdnProfile(name string, args DataCdnProfileArgs) *DataCdnProfile {
	return &DataCdnProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCdnProfile)(nil)

// DataCdnProfile represents the Terraform data resource azurerm_cdn_profile.
type DataCdnProfile struct {
	Name string
	Args DataCdnProfileArgs
}

// DataSource returns the Terraform object type for [DataCdnProfile].
func (cp *DataCdnProfile) DataSource() string {
	return "azurerm_cdn_profile"
}

// LocalName returns the local name for [DataCdnProfile].
func (cp *DataCdnProfile) LocalName() string {
	return cp.Name
}

// Configuration returns the configuration (args) for [DataCdnProfile].
func (cp *DataCdnProfile) Configuration() interface{} {
	return cp.Args
}

// Attributes returns the attributes for [DataCdnProfile].
func (cp *DataCdnProfile) Attributes() dataCdnProfileAttributes {
	return dataCdnProfileAttributes{ref: terra.ReferenceDataResource(cp)}
}

// DataCdnProfileArgs contains the configurations for azurerm_cdn_profile.
type DataCdnProfileArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datacdnprofile.Timeouts `hcl:"timeouts,block"`
}
type dataCdnProfileAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_cdn_profile.
func (cp dataCdnProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_cdn_profile.
func (cp dataCdnProfileAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_cdn_profile.
func (cp dataCdnProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_profile.
func (cp dataCdnProfileAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_cdn_profile.
func (cp dataCdnProfileAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_cdn_profile.
func (cp dataCdnProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cp.ref.Append("tags"))
}

func (cp dataCdnProfileAttributes) Timeouts() datacdnprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacdnprofile.TimeoutsAttributes](cp.ref.Append("timeouts"))
}
