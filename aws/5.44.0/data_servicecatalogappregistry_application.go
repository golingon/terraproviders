// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataServicecatalogappregistryApplication creates a new instance of [DataServicecatalogappregistryApplication].
func NewDataServicecatalogappregistryApplication(name string, args DataServicecatalogappregistryApplicationArgs) *DataServicecatalogappregistryApplication {
	return &DataServicecatalogappregistryApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicecatalogappregistryApplication)(nil)

// DataServicecatalogappregistryApplication represents the Terraform data resource aws_servicecatalogappregistry_application.
type DataServicecatalogappregistryApplication struct {
	Name string
	Args DataServicecatalogappregistryApplicationArgs
}

// DataSource returns the Terraform object type for [DataServicecatalogappregistryApplication].
func (sa *DataServicecatalogappregistryApplication) DataSource() string {
	return "aws_servicecatalogappregistry_application"
}

// LocalName returns the local name for [DataServicecatalogappregistryApplication].
func (sa *DataServicecatalogappregistryApplication) LocalName() string {
	return sa.Name
}

// Configuration returns the configuration (args) for [DataServicecatalogappregistryApplication].
func (sa *DataServicecatalogappregistryApplication) Configuration() interface{} {
	return sa.Args
}

// Attributes returns the attributes for [DataServicecatalogappregistryApplication].
func (sa *DataServicecatalogappregistryApplication) Attributes() dataServicecatalogappregistryApplicationAttributes {
	return dataServicecatalogappregistryApplicationAttributes{ref: terra.ReferenceDataResource(sa)}
}

// DataServicecatalogappregistryApplicationArgs contains the configurations for aws_servicecatalogappregistry_application.
type DataServicecatalogappregistryApplicationArgs struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
}
type dataServicecatalogappregistryApplicationAttributes struct {
	ref terra.Reference
}

// ApplicationTag returns a reference to field application_tag of aws_servicecatalogappregistry_application.
func (sa dataServicecatalogappregistryApplicationAttributes) ApplicationTag() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sa.ref.Append("application_tag"))
}

// Arn returns a reference to field arn of aws_servicecatalogappregistry_application.
func (sa dataServicecatalogappregistryApplicationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("arn"))
}

// Description returns a reference to field description of aws_servicecatalogappregistry_application.
func (sa dataServicecatalogappregistryApplicationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("description"))
}

// Id returns a reference to field id of aws_servicecatalogappregistry_application.
func (sa dataServicecatalogappregistryApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

// Name returns a reference to field name of aws_servicecatalogappregistry_application.
func (sa dataServicecatalogappregistryApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}
