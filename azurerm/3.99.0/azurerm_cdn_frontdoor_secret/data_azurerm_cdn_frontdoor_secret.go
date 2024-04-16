// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cdn_frontdoor_secret

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_cdn_frontdoor_secret.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (acfs *DataSource) DataSource() string {
	return "azurerm_cdn_frontdoor_secret"
}

// LocalName returns the local name for [DataSource].
func (acfs *DataSource) LocalName() string {
	return acfs.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (acfs *DataSource) Configuration() interface{} {
	return acfs.Args
}

// Attributes returns the attributes for [DataSource].
func (acfs *DataSource) Attributes() dataAzurermCdnFrontdoorSecretAttributes {
	return dataAzurermCdnFrontdoorSecretAttributes{ref: terra.ReferenceDataSource(acfs)}
}

// DataArgs contains the configurations for azurerm_cdn_frontdoor_secret.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProfileName: string, required
	ProfileName terra.StringValue `hcl:"profile_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermCdnFrontdoorSecretAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_secret.
func (acfs dataAzurermCdnFrontdoorSecretAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(acfs.ref.Append("cdn_frontdoor_profile_id"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_secret.
func (acfs dataAzurermCdnFrontdoorSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acfs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_secret.
func (acfs dataAzurermCdnFrontdoorSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acfs.ref.Append("name"))
}

// ProfileName returns a reference to field profile_name of azurerm_cdn_frontdoor_secret.
func (acfs dataAzurermCdnFrontdoorSecretAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceAsString(acfs.ref.Append("profile_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_frontdoor_secret.
func (acfs dataAzurermCdnFrontdoorSecretAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(acfs.ref.Append("resource_group_name"))
}

func (acfs dataAzurermCdnFrontdoorSecretAttributes) Secret() terra.ListValue[DataSecretAttributes] {
	return terra.ReferenceAsList[DataSecretAttributes](acfs.ref.Append("secret"))
}

func (acfs dataAzurermCdnFrontdoorSecretAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](acfs.ref.Append("timeouts"))
}
