// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIamSamlProvider creates a new instance of [DataIamSamlProvider].
func NewDataIamSamlProvider(name string, args DataIamSamlProviderArgs) *DataIamSamlProvider {
	return &DataIamSamlProvider{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamSamlProvider)(nil)

// DataIamSamlProvider represents the Terraform data resource aws_iam_saml_provider.
type DataIamSamlProvider struct {
	Name string
	Args DataIamSamlProviderArgs
}

// DataSource returns the Terraform object type for [DataIamSamlProvider].
func (isp *DataIamSamlProvider) DataSource() string {
	return "aws_iam_saml_provider"
}

// LocalName returns the local name for [DataIamSamlProvider].
func (isp *DataIamSamlProvider) LocalName() string {
	return isp.Name
}

// Configuration returns the configuration (args) for [DataIamSamlProvider].
func (isp *DataIamSamlProvider) Configuration() interface{} {
	return isp.Args
}

// Attributes returns the attributes for [DataIamSamlProvider].
func (isp *DataIamSamlProvider) Attributes() dataIamSamlProviderAttributes {
	return dataIamSamlProviderAttributes{ref: terra.ReferenceDataResource(isp)}
}

// DataIamSamlProviderArgs contains the configurations for aws_iam_saml_provider.
type DataIamSamlProviderArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataIamSamlProviderAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_saml_provider.
func (isp dataIamSamlProviderAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("arn"))
}

// CreateDate returns a reference to field create_date of aws_iam_saml_provider.
func (isp dataIamSamlProviderAttributes) CreateDate() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("create_date"))
}

// Id returns a reference to field id of aws_iam_saml_provider.
func (isp dataIamSamlProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("id"))
}

// Name returns a reference to field name of aws_iam_saml_provider.
func (isp dataIamSamlProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("name"))
}

// SamlMetadataDocument returns a reference to field saml_metadata_document of aws_iam_saml_provider.
func (isp dataIamSamlProviderAttributes) SamlMetadataDocument() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("saml_metadata_document"))
}

// Tags returns a reference to field tags of aws_iam_saml_provider.
func (isp dataIamSamlProviderAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](isp.ref.Append("tags"))
}

// ValidUntil returns a reference to field valid_until of aws_iam_saml_provider.
func (isp dataIamSamlProviderAttributes) ValidUntil() terra.StringValue {
	return terra.ReferenceAsString(isp.ref.Append("valid_until"))
}