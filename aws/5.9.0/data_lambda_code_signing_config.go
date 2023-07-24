// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalambdacodesigningconfig "github.com/golingon/terraproviders/aws/5.9.0/datalambdacodesigningconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLambdaCodeSigningConfig creates a new instance of [DataLambdaCodeSigningConfig].
func NewDataLambdaCodeSigningConfig(name string, args DataLambdaCodeSigningConfigArgs) *DataLambdaCodeSigningConfig {
	return &DataLambdaCodeSigningConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLambdaCodeSigningConfig)(nil)

// DataLambdaCodeSigningConfig represents the Terraform data resource aws_lambda_code_signing_config.
type DataLambdaCodeSigningConfig struct {
	Name string
	Args DataLambdaCodeSigningConfigArgs
}

// DataSource returns the Terraform object type for [DataLambdaCodeSigningConfig].
func (lcsc *DataLambdaCodeSigningConfig) DataSource() string {
	return "aws_lambda_code_signing_config"
}

// LocalName returns the local name for [DataLambdaCodeSigningConfig].
func (lcsc *DataLambdaCodeSigningConfig) LocalName() string {
	return lcsc.Name
}

// Configuration returns the configuration (args) for [DataLambdaCodeSigningConfig].
func (lcsc *DataLambdaCodeSigningConfig) Configuration() interface{} {
	return lcsc.Args
}

// Attributes returns the attributes for [DataLambdaCodeSigningConfig].
func (lcsc *DataLambdaCodeSigningConfig) Attributes() dataLambdaCodeSigningConfigAttributes {
	return dataLambdaCodeSigningConfigAttributes{ref: terra.ReferenceDataResource(lcsc)}
}

// DataLambdaCodeSigningConfigArgs contains the configurations for aws_lambda_code_signing_config.
type DataLambdaCodeSigningConfigArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// AllowedPublishers: min=0
	AllowedPublishers []datalambdacodesigningconfig.AllowedPublishers `hcl:"allowed_publishers,block" validate:"min=0"`
	// Policies: min=0
	Policies []datalambdacodesigningconfig.Policies `hcl:"policies,block" validate:"min=0"`
}
type dataLambdaCodeSigningConfigAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lambda_code_signing_config.
func (lcsc dataLambdaCodeSigningConfigAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(lcsc.ref.Append("arn"))
}

// ConfigId returns a reference to field config_id of aws_lambda_code_signing_config.
func (lcsc dataLambdaCodeSigningConfigAttributes) ConfigId() terra.StringValue {
	return terra.ReferenceAsString(lcsc.ref.Append("config_id"))
}

// Description returns a reference to field description of aws_lambda_code_signing_config.
func (lcsc dataLambdaCodeSigningConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lcsc.ref.Append("description"))
}

// Id returns a reference to field id of aws_lambda_code_signing_config.
func (lcsc dataLambdaCodeSigningConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lcsc.ref.Append("id"))
}

// LastModified returns a reference to field last_modified of aws_lambda_code_signing_config.
func (lcsc dataLambdaCodeSigningConfigAttributes) LastModified() terra.StringValue {
	return terra.ReferenceAsString(lcsc.ref.Append("last_modified"))
}

func (lcsc dataLambdaCodeSigningConfigAttributes) AllowedPublishers() terra.ListValue[datalambdacodesigningconfig.AllowedPublishersAttributes] {
	return terra.ReferenceAsList[datalambdacodesigningconfig.AllowedPublishersAttributes](lcsc.ref.Append("allowed_publishers"))
}

func (lcsc dataLambdaCodeSigningConfigAttributes) Policies() terra.ListValue[datalambdacodesigningconfig.PoliciesAttributes] {
	return terra.ReferenceAsList[datalambdacodesigningconfig.PoliciesAttributes](lcsc.ref.Append("policies"))
}
