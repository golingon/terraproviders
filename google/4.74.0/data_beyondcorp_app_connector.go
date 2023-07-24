// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	databeyondcorpappconnector "github.com/golingon/terraproviders/google/4.74.0/databeyondcorpappconnector"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBeyondcorpAppConnector creates a new instance of [DataBeyondcorpAppConnector].
func NewDataBeyondcorpAppConnector(name string, args DataBeyondcorpAppConnectorArgs) *DataBeyondcorpAppConnector {
	return &DataBeyondcorpAppConnector{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBeyondcorpAppConnector)(nil)

// DataBeyondcorpAppConnector represents the Terraform data resource google_beyondcorp_app_connector.
type DataBeyondcorpAppConnector struct {
	Name string
	Args DataBeyondcorpAppConnectorArgs
}

// DataSource returns the Terraform object type for [DataBeyondcorpAppConnector].
func (bac *DataBeyondcorpAppConnector) DataSource() string {
	return "google_beyondcorp_app_connector"
}

// LocalName returns the local name for [DataBeyondcorpAppConnector].
func (bac *DataBeyondcorpAppConnector) LocalName() string {
	return bac.Name
}

// Configuration returns the configuration (args) for [DataBeyondcorpAppConnector].
func (bac *DataBeyondcorpAppConnector) Configuration() interface{} {
	return bac.Args
}

// Attributes returns the attributes for [DataBeyondcorpAppConnector].
func (bac *DataBeyondcorpAppConnector) Attributes() dataBeyondcorpAppConnectorAttributes {
	return dataBeyondcorpAppConnectorAttributes{ref: terra.ReferenceDataResource(bac)}
}

// DataBeyondcorpAppConnectorArgs contains the configurations for google_beyondcorp_app_connector.
type DataBeyondcorpAppConnectorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// PrincipalInfo: min=0
	PrincipalInfo []databeyondcorpappconnector.PrincipalInfo `hcl:"principal_info,block" validate:"min=0"`
}
type dataBeyondcorpAppConnectorAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_beyondcorp_app_connector.
func (bac dataBeyondcorpAppConnectorAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("display_name"))
}

// Id returns a reference to field id of google_beyondcorp_app_connector.
func (bac dataBeyondcorpAppConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("id"))
}

// Labels returns a reference to field labels of google_beyondcorp_app_connector.
func (bac dataBeyondcorpAppConnectorAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bac.ref.Append("labels"))
}

// Name returns a reference to field name of google_beyondcorp_app_connector.
func (bac dataBeyondcorpAppConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("name"))
}

// Project returns a reference to field project of google_beyondcorp_app_connector.
func (bac dataBeyondcorpAppConnectorAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("project"))
}

// Region returns a reference to field region of google_beyondcorp_app_connector.
func (bac dataBeyondcorpAppConnectorAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("region"))
}

// State returns a reference to field state of google_beyondcorp_app_connector.
func (bac dataBeyondcorpAppConnectorAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("state"))
}

func (bac dataBeyondcorpAppConnectorAttributes) PrincipalInfo() terra.ListValue[databeyondcorpappconnector.PrincipalInfoAttributes] {
	return terra.ReferenceAsList[databeyondcorpappconnector.PrincipalInfoAttributes](bac.ref.Append("principal_info"))
}
