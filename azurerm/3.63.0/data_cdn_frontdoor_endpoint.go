// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacdnfrontdoorendpoint "github.com/golingon/terraproviders/azurerm/3.63.0/datacdnfrontdoorendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCdnFrontdoorEndpoint creates a new instance of [DataCdnFrontdoorEndpoint].
func NewDataCdnFrontdoorEndpoint(name string, args DataCdnFrontdoorEndpointArgs) *DataCdnFrontdoorEndpoint {
	return &DataCdnFrontdoorEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCdnFrontdoorEndpoint)(nil)

// DataCdnFrontdoorEndpoint represents the Terraform data resource azurerm_cdn_frontdoor_endpoint.
type DataCdnFrontdoorEndpoint struct {
	Name string
	Args DataCdnFrontdoorEndpointArgs
}

// DataSource returns the Terraform object type for [DataCdnFrontdoorEndpoint].
func (cfe *DataCdnFrontdoorEndpoint) DataSource() string {
	return "azurerm_cdn_frontdoor_endpoint"
}

// LocalName returns the local name for [DataCdnFrontdoorEndpoint].
func (cfe *DataCdnFrontdoorEndpoint) LocalName() string {
	return cfe.Name
}

// Configuration returns the configuration (args) for [DataCdnFrontdoorEndpoint].
func (cfe *DataCdnFrontdoorEndpoint) Configuration() interface{} {
	return cfe.Args
}

// Attributes returns the attributes for [DataCdnFrontdoorEndpoint].
func (cfe *DataCdnFrontdoorEndpoint) Attributes() dataCdnFrontdoorEndpointAttributes {
	return dataCdnFrontdoorEndpointAttributes{ref: terra.ReferenceDataResource(cfe)}
}

// DataCdnFrontdoorEndpointArgs contains the configurations for azurerm_cdn_frontdoor_endpoint.
type DataCdnFrontdoorEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProfileName: string, required
	ProfileName terra.StringValue `hcl:"profile_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datacdnfrontdoorendpoint.Timeouts `hcl:"timeouts,block"`
}
type dataCdnFrontdoorEndpointAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_cdn_frontdoor_endpoint.
func (cfe dataCdnFrontdoorEndpointAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cfe.ref.Append("enabled"))
}

// HostName returns a reference to field host_name of azurerm_cdn_frontdoor_endpoint.
func (cfe dataCdnFrontdoorEndpointAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(cfe.ref.Append("host_name"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_endpoint.
func (cfe dataCdnFrontdoorEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfe.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_endpoint.
func (cfe dataCdnFrontdoorEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfe.ref.Append("name"))
}

// ProfileName returns a reference to field profile_name of azurerm_cdn_frontdoor_endpoint.
func (cfe dataCdnFrontdoorEndpointAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceAsString(cfe.ref.Append("profile_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_frontdoor_endpoint.
func (cfe dataCdnFrontdoorEndpointAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cfe.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_cdn_frontdoor_endpoint.
func (cfe dataCdnFrontdoorEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cfe.ref.Append("tags"))
}

func (cfe dataCdnFrontdoorEndpointAttributes) Timeouts() datacdnfrontdoorendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacdnfrontdoorendpoint.TimeoutsAttributes](cfe.ref.Append("timeouts"))
}
