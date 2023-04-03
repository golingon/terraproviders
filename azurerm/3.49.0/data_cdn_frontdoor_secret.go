// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacdnfrontdoorsecret "github.com/golingon/terraproviders/azurerm/3.49.0/datacdnfrontdoorsecret"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCdnFrontdoorSecret creates a new instance of [DataCdnFrontdoorSecret].
func NewDataCdnFrontdoorSecret(name string, args DataCdnFrontdoorSecretArgs) *DataCdnFrontdoorSecret {
	return &DataCdnFrontdoorSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCdnFrontdoorSecret)(nil)

// DataCdnFrontdoorSecret represents the Terraform data resource azurerm_cdn_frontdoor_secret.
type DataCdnFrontdoorSecret struct {
	Name string
	Args DataCdnFrontdoorSecretArgs
}

// DataSource returns the Terraform object type for [DataCdnFrontdoorSecret].
func (cfs *DataCdnFrontdoorSecret) DataSource() string {
	return "azurerm_cdn_frontdoor_secret"
}

// LocalName returns the local name for [DataCdnFrontdoorSecret].
func (cfs *DataCdnFrontdoorSecret) LocalName() string {
	return cfs.Name
}

// Configuration returns the configuration (args) for [DataCdnFrontdoorSecret].
func (cfs *DataCdnFrontdoorSecret) Configuration() interface{} {
	return cfs.Args
}

// Attributes returns the attributes for [DataCdnFrontdoorSecret].
func (cfs *DataCdnFrontdoorSecret) Attributes() dataCdnFrontdoorSecretAttributes {
	return dataCdnFrontdoorSecretAttributes{ref: terra.ReferenceDataResource(cfs)}
}

// DataCdnFrontdoorSecretArgs contains the configurations for azurerm_cdn_frontdoor_secret.
type DataCdnFrontdoorSecretArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProfileName: string, required
	ProfileName terra.StringValue `hcl:"profile_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Secret: min=0
	Secret []datacdnfrontdoorsecret.Secret `hcl:"secret,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datacdnfrontdoorsecret.Timeouts `hcl:"timeouts,block"`
}
type dataCdnFrontdoorSecretAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_secret.
func (cfs dataCdnFrontdoorSecretAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(cfs.ref.Append("cdn_frontdoor_profile_id"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_secret.
func (cfs dataCdnFrontdoorSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_secret.
func (cfs dataCdnFrontdoorSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfs.ref.Append("name"))
}

// ProfileName returns a reference to field profile_name of azurerm_cdn_frontdoor_secret.
func (cfs dataCdnFrontdoorSecretAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceAsString(cfs.ref.Append("profile_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_frontdoor_secret.
func (cfs dataCdnFrontdoorSecretAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cfs.ref.Append("resource_group_name"))
}

func (cfs dataCdnFrontdoorSecretAttributes) Secret() terra.ListValue[datacdnfrontdoorsecret.SecretAttributes] {
	return terra.ReferenceAsList[datacdnfrontdoorsecret.SecretAttributes](cfs.ref.Append("secret"))
}

func (cfs dataCdnFrontdoorSecretAttributes) Timeouts() datacdnfrontdoorsecret.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacdnfrontdoorsecret.TimeoutsAttributes](cfs.ref.Append("timeouts"))
}
