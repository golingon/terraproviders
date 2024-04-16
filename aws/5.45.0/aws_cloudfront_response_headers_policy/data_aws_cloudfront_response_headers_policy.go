// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudfront_response_headers_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_cloudfront_response_headers_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (acrhp *DataSource) DataSource() string {
	return "aws_cloudfront_response_headers_policy"
}

// LocalName returns the local name for [DataSource].
func (acrhp *DataSource) LocalName() string {
	return acrhp.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (acrhp *DataSource) Configuration() interface{} {
	return acrhp.Args
}

// Attributes returns the attributes for [DataSource].
func (acrhp *DataSource) Attributes() dataAwsCloudfrontResponseHeadersPolicyAttributes {
	return dataAwsCloudfrontResponseHeadersPolicyAttributes{ref: terra.ReferenceDataSource(acrhp)}
}

// DataArgs contains the configurations for aws_cloudfront_response_headers_policy.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type dataAwsCloudfrontResponseHeadersPolicyAttributes struct {
	ref terra.Reference
}

// Comment returns a reference to field comment of aws_cloudfront_response_headers_policy.
func (acrhp dataAwsCloudfrontResponseHeadersPolicyAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(acrhp.ref.Append("comment"))
}

// Etag returns a reference to field etag of aws_cloudfront_response_headers_policy.
func (acrhp dataAwsCloudfrontResponseHeadersPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(acrhp.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_response_headers_policy.
func (acrhp dataAwsCloudfrontResponseHeadersPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acrhp.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudfront_response_headers_policy.
func (acrhp dataAwsCloudfrontResponseHeadersPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acrhp.ref.Append("name"))
}

func (acrhp dataAwsCloudfrontResponseHeadersPolicyAttributes) CorsConfig() terra.ListValue[DataCorsConfigAttributes] {
	return terra.ReferenceAsList[DataCorsConfigAttributes](acrhp.ref.Append("cors_config"))
}

func (acrhp dataAwsCloudfrontResponseHeadersPolicyAttributes) CustomHeadersConfig() terra.ListValue[DataCustomHeadersConfigAttributes] {
	return terra.ReferenceAsList[DataCustomHeadersConfigAttributes](acrhp.ref.Append("custom_headers_config"))
}

func (acrhp dataAwsCloudfrontResponseHeadersPolicyAttributes) RemoveHeadersConfig() terra.ListValue[DataRemoveHeadersConfigAttributes] {
	return terra.ReferenceAsList[DataRemoveHeadersConfigAttributes](acrhp.ref.Append("remove_headers_config"))
}

func (acrhp dataAwsCloudfrontResponseHeadersPolicyAttributes) SecurityHeadersConfig() terra.ListValue[DataSecurityHeadersConfigAttributes] {
	return terra.ReferenceAsList[DataSecurityHeadersConfigAttributes](acrhp.ref.Append("security_headers_config"))
}

func (acrhp dataAwsCloudfrontResponseHeadersPolicyAttributes) ServerTimingHeadersConfig() terra.ListValue[DataServerTimingHeadersConfigAttributes] {
	return terra.ReferenceAsList[DataServerTimingHeadersConfigAttributes](acrhp.ref.Append("server_timing_headers_config"))
}
