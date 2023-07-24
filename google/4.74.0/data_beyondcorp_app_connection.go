// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	databeyondcorpappconnection "github.com/golingon/terraproviders/google/4.74.0/databeyondcorpappconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBeyondcorpAppConnection creates a new instance of [DataBeyondcorpAppConnection].
func NewDataBeyondcorpAppConnection(name string, args DataBeyondcorpAppConnectionArgs) *DataBeyondcorpAppConnection {
	return &DataBeyondcorpAppConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBeyondcorpAppConnection)(nil)

// DataBeyondcorpAppConnection represents the Terraform data resource google_beyondcorp_app_connection.
type DataBeyondcorpAppConnection struct {
	Name string
	Args DataBeyondcorpAppConnectionArgs
}

// DataSource returns the Terraform object type for [DataBeyondcorpAppConnection].
func (bac *DataBeyondcorpAppConnection) DataSource() string {
	return "google_beyondcorp_app_connection"
}

// LocalName returns the local name for [DataBeyondcorpAppConnection].
func (bac *DataBeyondcorpAppConnection) LocalName() string {
	return bac.Name
}

// Configuration returns the configuration (args) for [DataBeyondcorpAppConnection].
func (bac *DataBeyondcorpAppConnection) Configuration() interface{} {
	return bac.Args
}

// Attributes returns the attributes for [DataBeyondcorpAppConnection].
func (bac *DataBeyondcorpAppConnection) Attributes() dataBeyondcorpAppConnectionAttributes {
	return dataBeyondcorpAppConnectionAttributes{ref: terra.ReferenceDataResource(bac)}
}

// DataBeyondcorpAppConnectionArgs contains the configurations for google_beyondcorp_app_connection.
type DataBeyondcorpAppConnectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ApplicationEndpoint: min=0
	ApplicationEndpoint []databeyondcorpappconnection.ApplicationEndpoint `hcl:"application_endpoint,block" validate:"min=0"`
	// Gateway: min=0
	Gateway []databeyondcorpappconnection.Gateway `hcl:"gateway,block" validate:"min=0"`
}
type dataBeyondcorpAppConnectionAttributes struct {
	ref terra.Reference
}

// Connectors returns a reference to field connectors of google_beyondcorp_app_connection.
func (bac dataBeyondcorpAppConnectionAttributes) Connectors() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bac.ref.Append("connectors"))
}

// DisplayName returns a reference to field display_name of google_beyondcorp_app_connection.
func (bac dataBeyondcorpAppConnectionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("display_name"))
}

// Id returns a reference to field id of google_beyondcorp_app_connection.
func (bac dataBeyondcorpAppConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("id"))
}

// Labels returns a reference to field labels of google_beyondcorp_app_connection.
func (bac dataBeyondcorpAppConnectionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bac.ref.Append("labels"))
}

// Name returns a reference to field name of google_beyondcorp_app_connection.
func (bac dataBeyondcorpAppConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("name"))
}

// Project returns a reference to field project of google_beyondcorp_app_connection.
func (bac dataBeyondcorpAppConnectionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("project"))
}

// Region returns a reference to field region of google_beyondcorp_app_connection.
func (bac dataBeyondcorpAppConnectionAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("region"))
}

// Type returns a reference to field type of google_beyondcorp_app_connection.
func (bac dataBeyondcorpAppConnectionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("type"))
}

func (bac dataBeyondcorpAppConnectionAttributes) ApplicationEndpoint() terra.ListValue[databeyondcorpappconnection.ApplicationEndpointAttributes] {
	return terra.ReferenceAsList[databeyondcorpappconnection.ApplicationEndpointAttributes](bac.ref.Append("application_endpoint"))
}

func (bac dataBeyondcorpAppConnectionAttributes) Gateway() terra.ListValue[databeyondcorpappconnection.GatewayAttributes] {
	return terra.ReferenceAsList[databeyondcorpappconnection.GatewayAttributes](bac.ref.Append("gateway"))
}
