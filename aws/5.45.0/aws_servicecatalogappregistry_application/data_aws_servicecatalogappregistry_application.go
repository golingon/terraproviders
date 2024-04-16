// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_servicecatalogappregistry_application

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_servicecatalogappregistry_application.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (asa *DataSource) DataSource() string {
	return "aws_servicecatalogappregistry_application"
}

// LocalName returns the local name for [DataSource].
func (asa *DataSource) LocalName() string {
	return asa.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (asa *DataSource) Configuration() interface{} {
	return asa.Args
}

// Attributes returns the attributes for [DataSource].
func (asa *DataSource) Attributes() dataAwsServicecatalogappregistryApplicationAttributes {
	return dataAwsServicecatalogappregistryApplicationAttributes{ref: terra.ReferenceDataSource(asa)}
}

// DataArgs contains the configurations for aws_servicecatalogappregistry_application.
type DataArgs struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
}

type dataAwsServicecatalogappregistryApplicationAttributes struct {
	ref terra.Reference
}

// ApplicationTag returns a reference to field application_tag of aws_servicecatalogappregistry_application.
func (asa dataAwsServicecatalogappregistryApplicationAttributes) ApplicationTag() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asa.ref.Append("application_tag"))
}

// Arn returns a reference to field arn of aws_servicecatalogappregistry_application.
func (asa dataAwsServicecatalogappregistryApplicationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("arn"))
}

// Description returns a reference to field description of aws_servicecatalogappregistry_application.
func (asa dataAwsServicecatalogappregistryApplicationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("description"))
}

// Id returns a reference to field id of aws_servicecatalogappregistry_application.
func (asa dataAwsServicecatalogappregistryApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("id"))
}

// Name returns a reference to field name of aws_servicecatalogappregistry_application.
func (asa dataAwsServicecatalogappregistryApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("name"))
}
