// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacdnfrontdoorsecret "github.com/golingon/terraproviders/azurerm/3.49.0/datacdnfrontdoorsecret"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataCdnFrontdoorSecret(name string, args DataCdnFrontdoorSecretArgs) *DataCdnFrontdoorSecret {
	return &DataCdnFrontdoorSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCdnFrontdoorSecret)(nil)

type DataCdnFrontdoorSecret struct {
	Name string
	Args DataCdnFrontdoorSecretArgs
}

func (cfs *DataCdnFrontdoorSecret) DataSource() string {
	return "azurerm_cdn_frontdoor_secret"
}

func (cfs *DataCdnFrontdoorSecret) LocalName() string {
	return cfs.Name
}

func (cfs *DataCdnFrontdoorSecret) Configuration() interface{} {
	return cfs.Args
}

func (cfs *DataCdnFrontdoorSecret) Attributes() dataCdnFrontdoorSecretAttributes {
	return dataCdnFrontdoorSecretAttributes{ref: terra.ReferenceDataResource(cfs)}
}

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

func (cfs dataCdnFrontdoorSecretAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceString(cfs.ref.Append("cdn_frontdoor_profile_id"))
}

func (cfs dataCdnFrontdoorSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cfs.ref.Append("id"))
}

func (cfs dataCdnFrontdoorSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cfs.ref.Append("name"))
}

func (cfs dataCdnFrontdoorSecretAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceString(cfs.ref.Append("profile_name"))
}

func (cfs dataCdnFrontdoorSecretAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(cfs.ref.Append("resource_group_name"))
}

func (cfs dataCdnFrontdoorSecretAttributes) Secret() terra.ListValue[datacdnfrontdoorsecret.SecretAttributes] {
	return terra.ReferenceList[datacdnfrontdoorsecret.SecretAttributes](cfs.ref.Append("secret"))
}

func (cfs dataCdnFrontdoorSecretAttributes) Timeouts() datacdnfrontdoorsecret.TimeoutsAttributes {
	return terra.ReferenceSingle[datacdnfrontdoorsecret.TimeoutsAttributes](cfs.ref.Append("timeouts"))
}
