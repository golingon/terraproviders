// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatelinkserviceendpointconnections "github.com/golingon/terraproviders/azurerm/3.68.0/dataprivatelinkserviceendpointconnections"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateLinkServiceEndpointConnections creates a new instance of [DataPrivateLinkServiceEndpointConnections].
func NewDataPrivateLinkServiceEndpointConnections(name string, args DataPrivateLinkServiceEndpointConnectionsArgs) *DataPrivateLinkServiceEndpointConnections {
	return &DataPrivateLinkServiceEndpointConnections{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateLinkServiceEndpointConnections)(nil)

// DataPrivateLinkServiceEndpointConnections represents the Terraform data resource azurerm_private_link_service_endpoint_connections.
type DataPrivateLinkServiceEndpointConnections struct {
	Name string
	Args DataPrivateLinkServiceEndpointConnectionsArgs
}

// DataSource returns the Terraform object type for [DataPrivateLinkServiceEndpointConnections].
func (plsec *DataPrivateLinkServiceEndpointConnections) DataSource() string {
	return "azurerm_private_link_service_endpoint_connections"
}

// LocalName returns the local name for [DataPrivateLinkServiceEndpointConnections].
func (plsec *DataPrivateLinkServiceEndpointConnections) LocalName() string {
	return plsec.Name
}

// Configuration returns the configuration (args) for [DataPrivateLinkServiceEndpointConnections].
func (plsec *DataPrivateLinkServiceEndpointConnections) Configuration() interface{} {
	return plsec.Args
}

// Attributes returns the attributes for [DataPrivateLinkServiceEndpointConnections].
func (plsec *DataPrivateLinkServiceEndpointConnections) Attributes() dataPrivateLinkServiceEndpointConnectionsAttributes {
	return dataPrivateLinkServiceEndpointConnectionsAttributes{ref: terra.ReferenceDataResource(plsec)}
}

// DataPrivateLinkServiceEndpointConnectionsArgs contains the configurations for azurerm_private_link_service_endpoint_connections.
type DataPrivateLinkServiceEndpointConnectionsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceId: string, required
	ServiceId terra.StringValue `hcl:"service_id,attr" validate:"required"`
	// PrivateEndpointConnections: min=0
	PrivateEndpointConnections []dataprivatelinkserviceendpointconnections.PrivateEndpointConnections `hcl:"private_endpoint_connections,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataprivatelinkserviceendpointconnections.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateLinkServiceEndpointConnectionsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_link_service_endpoint_connections.
func (plsec dataPrivateLinkServiceEndpointConnectionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(plsec.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_private_link_service_endpoint_connections.
func (plsec dataPrivateLinkServiceEndpointConnectionsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(plsec.ref.Append("location"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_link_service_endpoint_connections.
func (plsec dataPrivateLinkServiceEndpointConnectionsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(plsec.ref.Append("resource_group_name"))
}

// ServiceId returns a reference to field service_id of azurerm_private_link_service_endpoint_connections.
func (plsec dataPrivateLinkServiceEndpointConnectionsAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(plsec.ref.Append("service_id"))
}

// ServiceName returns a reference to field service_name of azurerm_private_link_service_endpoint_connections.
func (plsec dataPrivateLinkServiceEndpointConnectionsAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(plsec.ref.Append("service_name"))
}

func (plsec dataPrivateLinkServiceEndpointConnectionsAttributes) PrivateEndpointConnections() terra.ListValue[dataprivatelinkserviceendpointconnections.PrivateEndpointConnectionsAttributes] {
	return terra.ReferenceAsList[dataprivatelinkserviceendpointconnections.PrivateEndpointConnectionsAttributes](plsec.ref.Append("private_endpoint_connections"))
}

func (plsec dataPrivateLinkServiceEndpointConnectionsAttributes) Timeouts() dataprivatelinkserviceendpointconnections.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatelinkserviceendpointconnections.TimeoutsAttributes](plsec.ref.Append("timeouts"))
}