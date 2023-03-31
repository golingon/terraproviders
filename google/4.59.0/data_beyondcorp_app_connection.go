// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	databeyondcorpappconnection "github.com/golingon/terraproviders/google/4.59.0/databeyondcorpappconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataBeyondcorpAppConnection(name string, args DataBeyondcorpAppConnectionArgs) *DataBeyondcorpAppConnection {
	return &DataBeyondcorpAppConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBeyondcorpAppConnection)(nil)

type DataBeyondcorpAppConnection struct {
	Name string
	Args DataBeyondcorpAppConnectionArgs
}

func (bac *DataBeyondcorpAppConnection) DataSource() string {
	return "google_beyondcorp_app_connection"
}

func (bac *DataBeyondcorpAppConnection) LocalName() string {
	return bac.Name
}

func (bac *DataBeyondcorpAppConnection) Configuration() interface{} {
	return bac.Args
}

func (bac *DataBeyondcorpAppConnection) Attributes() dataBeyondcorpAppConnectionAttributes {
	return dataBeyondcorpAppConnectionAttributes{ref: terra.ReferenceDataResource(bac)}
}

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

func (bac dataBeyondcorpAppConnectionAttributes) Connectors() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](bac.ref.Append("connectors"))
}

func (bac dataBeyondcorpAppConnectionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(bac.ref.Append("display_name"))
}

func (bac dataBeyondcorpAppConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bac.ref.Append("id"))
}

func (bac dataBeyondcorpAppConnectionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](bac.ref.Append("labels"))
}

func (bac dataBeyondcorpAppConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(bac.ref.Append("name"))
}

func (bac dataBeyondcorpAppConnectionAttributes) Project() terra.StringValue {
	return terra.ReferenceString(bac.ref.Append("project"))
}

func (bac dataBeyondcorpAppConnectionAttributes) Region() terra.StringValue {
	return terra.ReferenceString(bac.ref.Append("region"))
}

func (bac dataBeyondcorpAppConnectionAttributes) Type() terra.StringValue {
	return terra.ReferenceString(bac.ref.Append("type"))
}

func (bac dataBeyondcorpAppConnectionAttributes) ApplicationEndpoint() terra.ListValue[databeyondcorpappconnection.ApplicationEndpointAttributes] {
	return terra.ReferenceList[databeyondcorpappconnection.ApplicationEndpointAttributes](bac.ref.Append("application_endpoint"))
}

func (bac dataBeyondcorpAppConnectionAttributes) Gateway() terra.ListValue[databeyondcorpappconnection.GatewayAttributes] {
	return terra.ReferenceList[databeyondcorpappconnection.GatewayAttributes](bac.ref.Append("gateway"))
}
