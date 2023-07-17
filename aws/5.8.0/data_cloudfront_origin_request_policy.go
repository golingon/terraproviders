// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datacloudfrontoriginrequestpolicy "github.com/golingon/terraproviders/aws/5.8.0/datacloudfrontoriginrequestpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudfrontOriginRequestPolicy creates a new instance of [DataCloudfrontOriginRequestPolicy].
func NewDataCloudfrontOriginRequestPolicy(name string, args DataCloudfrontOriginRequestPolicyArgs) *DataCloudfrontOriginRequestPolicy {
	return &DataCloudfrontOriginRequestPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudfrontOriginRequestPolicy)(nil)

// DataCloudfrontOriginRequestPolicy represents the Terraform data resource aws_cloudfront_origin_request_policy.
type DataCloudfrontOriginRequestPolicy struct {
	Name string
	Args DataCloudfrontOriginRequestPolicyArgs
}

// DataSource returns the Terraform object type for [DataCloudfrontOriginRequestPolicy].
func (corp *DataCloudfrontOriginRequestPolicy) DataSource() string {
	return "aws_cloudfront_origin_request_policy"
}

// LocalName returns the local name for [DataCloudfrontOriginRequestPolicy].
func (corp *DataCloudfrontOriginRequestPolicy) LocalName() string {
	return corp.Name
}

// Configuration returns the configuration (args) for [DataCloudfrontOriginRequestPolicy].
func (corp *DataCloudfrontOriginRequestPolicy) Configuration() interface{} {
	return corp.Args
}

// Attributes returns the attributes for [DataCloudfrontOriginRequestPolicy].
func (corp *DataCloudfrontOriginRequestPolicy) Attributes() dataCloudfrontOriginRequestPolicyAttributes {
	return dataCloudfrontOriginRequestPolicyAttributes{ref: terra.ReferenceDataResource(corp)}
}

// DataCloudfrontOriginRequestPolicyArgs contains the configurations for aws_cloudfront_origin_request_policy.
type DataCloudfrontOriginRequestPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// CookiesConfig: min=0
	CookiesConfig []datacloudfrontoriginrequestpolicy.CookiesConfig `hcl:"cookies_config,block" validate:"min=0"`
	// HeadersConfig: min=0
	HeadersConfig []datacloudfrontoriginrequestpolicy.HeadersConfig `hcl:"headers_config,block" validate:"min=0"`
	// QueryStringsConfig: min=0
	QueryStringsConfig []datacloudfrontoriginrequestpolicy.QueryStringsConfig `hcl:"query_strings_config,block" validate:"min=0"`
}
type dataCloudfrontOriginRequestPolicyAttributes struct {
	ref terra.Reference
}

// Comment returns a reference to field comment of aws_cloudfront_origin_request_policy.
func (corp dataCloudfrontOriginRequestPolicyAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(corp.ref.Append("comment"))
}

// Etag returns a reference to field etag of aws_cloudfront_origin_request_policy.
func (corp dataCloudfrontOriginRequestPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(corp.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_origin_request_policy.
func (corp dataCloudfrontOriginRequestPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(corp.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudfront_origin_request_policy.
func (corp dataCloudfrontOriginRequestPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(corp.ref.Append("name"))
}

func (corp dataCloudfrontOriginRequestPolicyAttributes) CookiesConfig() terra.ListValue[datacloudfrontoriginrequestpolicy.CookiesConfigAttributes] {
	return terra.ReferenceAsList[datacloudfrontoriginrequestpolicy.CookiesConfigAttributes](corp.ref.Append("cookies_config"))
}

func (corp dataCloudfrontOriginRequestPolicyAttributes) HeadersConfig() terra.ListValue[datacloudfrontoriginrequestpolicy.HeadersConfigAttributes] {
	return terra.ReferenceAsList[datacloudfrontoriginrequestpolicy.HeadersConfigAttributes](corp.ref.Append("headers_config"))
}

func (corp dataCloudfrontOriginRequestPolicyAttributes) QueryStringsConfig() terra.ListValue[datacloudfrontoriginrequestpolicy.QueryStringsConfigAttributes] {
	return terra.ReferenceAsList[datacloudfrontoriginrequestpolicy.QueryStringsConfigAttributes](corp.ref.Append("query_strings_config"))
}