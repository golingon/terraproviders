// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_beyondcorp_app_connection

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_beyondcorp_app_connection.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gbac *DataSource) DataSource() string {
	return "google_beyondcorp_app_connection"
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
func (gbac *DataSource) Attributes() dataGoogleBeyondcorpAppConnectionAttributes {
	return dataGoogleBeyondcorpAppConnectionAttributes{ref: terra.ReferenceDataSource(gbac)}
}

// DataArgs contains the configurations for google_beyondcorp_app_connection.
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

type dataGoogleBeyondcorpAppConnectionAttributes struct {
	ref terra.Reference
}

// Connectors returns a reference to field connectors of google_beyondcorp_app_connection.
func (gbac dataGoogleBeyondcorpAppConnectionAttributes) Connectors() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gbac.ref.Append("connectors"))
}

// DisplayName returns a reference to field display_name of google_beyondcorp_app_connection.
func (gbac dataGoogleBeyondcorpAppConnectionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("display_name"))
}

// EffectiveLabels returns a reference to field effective_labels of google_beyondcorp_app_connection.
func (gbac dataGoogleBeyondcorpAppConnectionAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gbac.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_beyondcorp_app_connection.
func (gbac dataGoogleBeyondcorpAppConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("id"))
}

// Labels returns a reference to field labels of google_beyondcorp_app_connection.
func (gbac dataGoogleBeyondcorpAppConnectionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gbac.ref.Append("labels"))
}

// Name returns a reference to field name of google_beyondcorp_app_connection.
func (gbac dataGoogleBeyondcorpAppConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("name"))
}

// Project returns a reference to field project of google_beyondcorp_app_connection.
func (gbac dataGoogleBeyondcorpAppConnectionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("project"))
}

// Region returns a reference to field region of google_beyondcorp_app_connection.
func (gbac dataGoogleBeyondcorpAppConnectionAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("region"))
}

// TerraformLabels returns a reference to field terraform_labels of google_beyondcorp_app_connection.
func (gbac dataGoogleBeyondcorpAppConnectionAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gbac.ref.Append("terraform_labels"))
}

// Type returns a reference to field type of google_beyondcorp_app_connection.
func (gbac dataGoogleBeyondcorpAppConnectionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(gbac.ref.Append("type"))
}

func (gbac dataGoogleBeyondcorpAppConnectionAttributes) ApplicationEndpoint() terra.ListValue[DataApplicationEndpointAttributes] {
	return terra.ReferenceAsList[DataApplicationEndpointAttributes](gbac.ref.Append("application_endpoint"))
}

func (gbac dataGoogleBeyondcorpAppConnectionAttributes) Gateway() terra.ListValue[DataGatewayAttributes] {
	return terra.ReferenceAsList[DataGatewayAttributes](gbac.ref.Append("gateway"))
}
