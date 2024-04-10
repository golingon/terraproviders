// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataCloudfrontFunction creates a new instance of [DataCloudfrontFunction].
func NewDataCloudfrontFunction(name string, args DataCloudfrontFunctionArgs) *DataCloudfrontFunction {
	return &DataCloudfrontFunction{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudfrontFunction)(nil)

// DataCloudfrontFunction represents the Terraform data resource aws_cloudfront_function.
type DataCloudfrontFunction struct {
	Name string
	Args DataCloudfrontFunctionArgs
}

// DataSource returns the Terraform object type for [DataCloudfrontFunction].
func (cf *DataCloudfrontFunction) DataSource() string {
	return "aws_cloudfront_function"
}

// LocalName returns the local name for [DataCloudfrontFunction].
func (cf *DataCloudfrontFunction) LocalName() string {
	return cf.Name
}

// Configuration returns the configuration (args) for [DataCloudfrontFunction].
func (cf *DataCloudfrontFunction) Configuration() interface{} {
	return cf.Args
}

// Attributes returns the attributes for [DataCloudfrontFunction].
func (cf *DataCloudfrontFunction) Attributes() dataCloudfrontFunctionAttributes {
	return dataCloudfrontFunctionAttributes{ref: terra.ReferenceDataResource(cf)}
}

// DataCloudfrontFunctionArgs contains the configurations for aws_cloudfront_function.
type DataCloudfrontFunctionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Stage: string, required
	Stage terra.StringValue `hcl:"stage,attr" validate:"required"`
}
type dataCloudfrontFunctionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("arn"))
}

// Code returns a reference to field code of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) Code() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("code"))
}

// Comment returns a reference to field comment of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("comment"))
}

// Etag returns a reference to field etag of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("id"))
}

// KeyValueStoreAssociations returns a reference to field key_value_store_associations of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) KeyValueStoreAssociations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cf.ref.Append("key_value_store_associations"))
}

// LastModifiedTime returns a reference to field last_modified_time of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("last_modified_time"))
}

// Name returns a reference to field name of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("name"))
}

// Runtime returns a reference to field runtime of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) Runtime() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("runtime"))
}

// Stage returns a reference to field stage of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) Stage() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("stage"))
}

// Status returns a reference to field status of aws_cloudfront_function.
func (cf dataCloudfrontFunctionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("status"))
}
