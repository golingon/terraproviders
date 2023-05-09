// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamobilenetworksite "github.com/golingon/terraproviders/azurerm/3.55.0/datamobilenetworksite"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMobileNetworkSite creates a new instance of [DataMobileNetworkSite].
func NewDataMobileNetworkSite(name string, args DataMobileNetworkSiteArgs) *DataMobileNetworkSite {
	return &DataMobileNetworkSite{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMobileNetworkSite)(nil)

// DataMobileNetworkSite represents the Terraform data resource azurerm_mobile_network_site.
type DataMobileNetworkSite struct {
	Name string
	Args DataMobileNetworkSiteArgs
}

// DataSource returns the Terraform object type for [DataMobileNetworkSite].
func (mns *DataMobileNetworkSite) DataSource() string {
	return "azurerm_mobile_network_site"
}

// LocalName returns the local name for [DataMobileNetworkSite].
func (mns *DataMobileNetworkSite) LocalName() string {
	return mns.Name
}

// Configuration returns the configuration (args) for [DataMobileNetworkSite].
func (mns *DataMobileNetworkSite) Configuration() interface{} {
	return mns.Args
}

// Attributes returns the attributes for [DataMobileNetworkSite].
func (mns *DataMobileNetworkSite) Attributes() dataMobileNetworkSiteAttributes {
	return dataMobileNetworkSiteAttributes{ref: terra.ReferenceDataResource(mns)}
}

// DataMobileNetworkSiteArgs contains the configurations for azurerm_mobile_network_site.
type DataMobileNetworkSiteArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MobileNetworkId: string, required
	MobileNetworkId terra.StringValue `hcl:"mobile_network_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datamobilenetworksite.Timeouts `hcl:"timeouts,block"`
}
type dataMobileNetworkSiteAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mobile_network_site.
func (mns dataMobileNetworkSiteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_site.
func (mns dataMobileNetworkSiteAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_site.
func (mns dataMobileNetworkSiteAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_site.
func (mns dataMobileNetworkSiteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("name"))
}

// NetworkFunctionIds returns a reference to field network_function_ids of azurerm_mobile_network_site.
func (mns dataMobileNetworkSiteAttributes) NetworkFunctionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mns.ref.Append("network_function_ids"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_site.
func (mns dataMobileNetworkSiteAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mns.ref.Append("tags"))
}

func (mns dataMobileNetworkSiteAttributes) Timeouts() datamobilenetworksite.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamobilenetworksite.TimeoutsAttributes](mns.ref.Append("timeouts"))
}
