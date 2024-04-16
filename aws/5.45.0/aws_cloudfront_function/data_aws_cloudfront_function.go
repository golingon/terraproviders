// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudfront_function

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_cloudfront_function.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (acf *DataSource) DataSource() string {
	return "aws_cloudfront_function"
}

// LocalName returns the local name for [DataSource].
func (acf *DataSource) LocalName() string {
	return acf.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (acf *DataSource) Configuration() interface{} {
	return acf.Args
}

// Attributes returns the attributes for [DataSource].
func (acf *DataSource) Attributes() dataAwsCloudfrontFunctionAttributes {
	return dataAwsCloudfrontFunctionAttributes{ref: terra.ReferenceDataSource(acf)}
}

// DataArgs contains the configurations for aws_cloudfront_function.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Stage: string, required
	Stage terra.StringValue `hcl:"stage,attr" validate:"required"`
}

type dataAwsCloudfrontFunctionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acf.ref.Append("arn"))
}

// Code returns a reference to field code of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) Code() terra.StringValue {
	return terra.ReferenceAsString(acf.ref.Append("code"))
}

// Comment returns a reference to field comment of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(acf.ref.Append("comment"))
}

// Etag returns a reference to field etag of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(acf.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acf.ref.Append("id"))
}

// KeyValueStoreAssociations returns a reference to field key_value_store_associations of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) KeyValueStoreAssociations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acf.ref.Append("key_value_store_associations"))
}

// LastModifiedTime returns a reference to field last_modified_time of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceAsString(acf.ref.Append("last_modified_time"))
}

// Name returns a reference to field name of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acf.ref.Append("name"))
}

// Runtime returns a reference to field runtime of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) Runtime() terra.StringValue {
	return terra.ReferenceAsString(acf.ref.Append("runtime"))
}

// Stage returns a reference to field stage of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) Stage() terra.StringValue {
	return terra.ReferenceAsString(acf.ref.Append("stage"))
}

// Status returns a reference to field status of aws_cloudfront_function.
func (acf dataAwsCloudfrontFunctionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(acf.ref.Append("status"))
}
