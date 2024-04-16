// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_beyondcorp_app_connector

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_beyondcorp_app_connector.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gbac *DataSource) DataSource() string {
	return "google_beyondcorp_app_connector"
}

// LocalName returns the local name for [DataSource].
func (gbac *DataSource) LocalName() string {
	return gbac.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gbac *DataSource) Configuration() interface{} {
	return gbac.Args
}

// Attributes returns the attributes for [DataSource].
func (gbac *DataSource) Attributes() dataGoogleBeyondcorpAppConnectorAttributes {
	return dataGoogleBeyondcorpAppConnectorAttributes{ref: terra.ReferenceDataSource(gbac)}
}

// DataArgs contains the configurations for google_beyondcorp_app_connector.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}

type dataGoogleBeyondcorpAppConnectorAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_beyondcorp_app_connector.
func (gbac dataGoogleBeyondcorpAppConnectorAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("display_name"))
}

// EffectiveLabels returns a reference to field effective_labels of google_beyondcorp_app_connector.
func (gbac dataGoogleBeyondcorpAppConnectorAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gbac.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_beyondcorp_app_connector.
func (gbac dataGoogleBeyondcorpAppConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("id"))
}

// Labels returns a reference to field labels of google_beyondcorp_app_connector.
func (gbac dataGoogleBeyondcorpAppConnectorAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gbac.ref.Append("labels"))
}

// Name returns a reference to field name of google_beyondcorp_app_connector.
func (gbac dataGoogleBeyondcorpAppConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("name"))
}

// Project returns a reference to field project of google_beyondcorp_app_connector.
func (gbac dataGoogleBeyondcorpAppConnectorAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("project"))
}

// Region returns a reference to field region of google_beyondcorp_app_connector.
func (gbac dataGoogleBeyondcorpAppConnectorAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("region"))
}

// State returns a reference to field state of google_beyondcorp_app_connector.
func (gbac dataGoogleBeyondcorpAppConnectorAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("state"))
}

// TerraformLabels returns a reference to field terraform_labels of google_beyondcorp_app_connector.
func (gbac dataGoogleBeyondcorpAppConnectorAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gbac.ref.Append("terraform_labels"))
}

func (gbac dataGoogleBeyondcorpAppConnectorAttributes) PrincipalInfo() terra.ListValue[DataPrincipalInfoAttributes] {
	return terra.ReferenceAsList[DataPrincipalInfoAttributes](gbac.ref.Append("principal_info"))
}
