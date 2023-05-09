// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalambdafunctionurl "github.com/golingon/terraproviders/aws/4.66.1/datalambdafunctionurl"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLambdaFunctionUrl creates a new instance of [DataLambdaFunctionUrl].
func NewDataLambdaFunctionUrl(name string, args DataLambdaFunctionUrlArgs) *DataLambdaFunctionUrl {
	return &DataLambdaFunctionUrl{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLambdaFunctionUrl)(nil)

// DataLambdaFunctionUrl represents the Terraform data resource aws_lambda_function_url.
type DataLambdaFunctionUrl struct {
	Name string
	Args DataLambdaFunctionUrlArgs
}

// DataSource returns the Terraform object type for [DataLambdaFunctionUrl].
func (lfu *DataLambdaFunctionUrl) DataSource() string {
	return "aws_lambda_function_url"
}

// LocalName returns the local name for [DataLambdaFunctionUrl].
func (lfu *DataLambdaFunctionUrl) LocalName() string {
	return lfu.Name
}

// Configuration returns the configuration (args) for [DataLambdaFunctionUrl].
func (lfu *DataLambdaFunctionUrl) Configuration() interface{} {
	return lfu.Args
}

// Attributes returns the attributes for [DataLambdaFunctionUrl].
func (lfu *DataLambdaFunctionUrl) Attributes() dataLambdaFunctionUrlAttributes {
	return dataLambdaFunctionUrlAttributes{ref: terra.ReferenceDataResource(lfu)}
}

// DataLambdaFunctionUrlArgs contains the configurations for aws_lambda_function_url.
type DataLambdaFunctionUrlArgs struct {
	// FunctionName: string, required
	FunctionName terra.StringValue `hcl:"function_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Qualifier: string, optional
	Qualifier terra.StringValue `hcl:"qualifier,attr"`
	// Cors: min=0
	Cors []datalambdafunctionurl.Cors `hcl:"cors,block" validate:"min=0"`
}
type dataLambdaFunctionUrlAttributes struct {
	ref terra.Reference
}

// AuthorizationType returns a reference to field authorization_type of aws_lambda_function_url.
func (lfu dataLambdaFunctionUrlAttributes) AuthorizationType() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("authorization_type"))
}

// CreationTime returns a reference to field creation_time of aws_lambda_function_url.
func (lfu dataLambdaFunctionUrlAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("creation_time"))
}

// FunctionArn returns a reference to field function_arn of aws_lambda_function_url.
func (lfu dataLambdaFunctionUrlAttributes) FunctionArn() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("function_arn"))
}

// FunctionName returns a reference to field function_name of aws_lambda_function_url.
func (lfu dataLambdaFunctionUrlAttributes) FunctionName() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("function_name"))
}

// FunctionUrl returns a reference to field function_url of aws_lambda_function_url.
func (lfu dataLambdaFunctionUrlAttributes) FunctionUrl() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("function_url"))
}

// Id returns a reference to field id of aws_lambda_function_url.
func (lfu dataLambdaFunctionUrlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("id"))
}

// InvokeMode returns a reference to field invoke_mode of aws_lambda_function_url.
func (lfu dataLambdaFunctionUrlAttributes) InvokeMode() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("invoke_mode"))
}

// LastModifiedTime returns a reference to field last_modified_time of aws_lambda_function_url.
func (lfu dataLambdaFunctionUrlAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("last_modified_time"))
}

// Qualifier returns a reference to field qualifier of aws_lambda_function_url.
func (lfu dataLambdaFunctionUrlAttributes) Qualifier() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("qualifier"))
}

// UrlId returns a reference to field url_id of aws_lambda_function_url.
func (lfu dataLambdaFunctionUrlAttributes) UrlId() terra.StringValue {
	return terra.ReferenceAsString(lfu.ref.Append("url_id"))
}

func (lfu dataLambdaFunctionUrlAttributes) Cors() terra.ListValue[datalambdafunctionurl.CorsAttributes] {
	return terra.ReferenceAsList[datalambdafunctionurl.CorsAttributes](lfu.ref.Append("cors"))
}
