// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	databeyondcorpappgateway "github.com/golingon/terraproviders/googlebeta/4.77.0/databeyondcorpappgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBeyondcorpAppGateway creates a new instance of [DataBeyondcorpAppGateway].
func NewDataBeyondcorpAppGateway(name string, args DataBeyondcorpAppGatewayArgs) *DataBeyondcorpAppGateway {
	return &DataBeyondcorpAppGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBeyondcorpAppGateway)(nil)

// DataBeyondcorpAppGateway represents the Terraform data resource google_beyondcorp_app_gateway.
type DataBeyondcorpAppGateway struct {
	Name string
	Args DataBeyondcorpAppGatewayArgs
}

// DataSource returns the Terraform object type for [DataBeyondcorpAppGateway].
func (bag *DataBeyondcorpAppGateway) DataSource() string {
	return "google_beyondcorp_app_gateway"
}

// LocalName returns the local name for [DataBeyondcorpAppGateway].
func (bag *DataBeyondcorpAppGateway) LocalName() string {
	return bag.Name
}

// Configuration returns the configuration (args) for [DataBeyondcorpAppGateway].
func (bag *DataBeyondcorpAppGateway) Configuration() interface{} {
	return bag.Args
}

// Attributes returns the attributes for [DataBeyondcorpAppGateway].
func (bag *DataBeyondcorpAppGateway) Attributes() dataBeyondcorpAppGatewayAttributes {
	return dataBeyondcorpAppGatewayAttributes{ref: terra.ReferenceDataResource(bag)}
}

// DataBeyondcorpAppGatewayArgs contains the configurations for google_beyondcorp_app_gateway.
type DataBeyondcorpAppGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// AllocatedConnections: min=0
	AllocatedConnections []databeyondcorpappgateway.AllocatedConnections `hcl:"allocated_connections,block" validate:"min=0"`
}
type dataBeyondcorpAppGatewayAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_beyondcorp_app_gateway.
func (bag dataBeyondcorpAppGatewayAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("display_name"))
}

// HostType returns a reference to field host_type of google_beyondcorp_app_gateway.
func (bag dataBeyondcorpAppGatewayAttributes) HostType() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("host_type"))
}

// Id returns a reference to field id of google_beyondcorp_app_gateway.
func (bag dataBeyondcorpAppGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("id"))
}

// Labels returns a reference to field labels of google_beyondcorp_app_gateway.
func (bag dataBeyondcorpAppGatewayAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bag.ref.Append("labels"))
}

// Name returns a reference to field name of google_beyondcorp_app_gateway.
func (bag dataBeyondcorpAppGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("name"))
}

// Project returns a reference to field project of google_beyondcorp_app_gateway.
func (bag dataBeyondcorpAppGatewayAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("project"))
}

// Region returns a reference to field region of google_beyondcorp_app_gateway.
func (bag dataBeyondcorpAppGatewayAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("region"))
}

// State returns a reference to field state of google_beyondcorp_app_gateway.
func (bag dataBeyondcorpAppGatewayAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("state"))
}

// Type returns a reference to field type of google_beyondcorp_app_gateway.
func (bag dataBeyondcorpAppGatewayAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("type"))
}

// Uri returns a reference to field uri of google_beyondcorp_app_gateway.
func (bag dataBeyondcorpAppGatewayAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("uri"))
}

func (bag dataBeyondcorpAppGatewayAttributes) AllocatedConnections() terra.ListValue[databeyondcorpappgateway.AllocatedConnectionsAttributes] {
	return terra.ReferenceAsList[databeyondcorpappgateway.AllocatedConnectionsAttributes](bag.ref.Append("allocated_connections"))
}
