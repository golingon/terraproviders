// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datacloudfrontcachepolicy "github.com/golingon/terraproviders/aws/5.8.0/datacloudfrontcachepolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudfrontCachePolicy creates a new instance of [DataCloudfrontCachePolicy].
func NewDataCloudfrontCachePolicy(name string, args DataCloudfrontCachePolicyArgs) *DataCloudfrontCachePolicy {
	return &DataCloudfrontCachePolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudfrontCachePolicy)(nil)

// DataCloudfrontCachePolicy represents the Terraform data resource aws_cloudfront_cache_policy.
type DataCloudfrontCachePolicy struct {
	Name string
	Args DataCloudfrontCachePolicyArgs
}

// DataSource returns the Terraform object type for [DataCloudfrontCachePolicy].
func (ccp *DataCloudfrontCachePolicy) DataSource() string {
	return "aws_cloudfront_cache_policy"
}

// LocalName returns the local name for [DataCloudfrontCachePolicy].
func (ccp *DataCloudfrontCachePolicy) LocalName() string {
	return ccp.Name
}

// Configuration returns the configuration (args) for [DataCloudfrontCachePolicy].
func (ccp *DataCloudfrontCachePolicy) Configuration() interface{} {
	return ccp.Args
}

// Attributes returns the attributes for [DataCloudfrontCachePolicy].
func (ccp *DataCloudfrontCachePolicy) Attributes() dataCloudfrontCachePolicyAttributes {
	return dataCloudfrontCachePolicyAttributes{ref: terra.ReferenceDataResource(ccp)}
}

// DataCloudfrontCachePolicyArgs contains the configurations for aws_cloudfront_cache_policy.
type DataCloudfrontCachePolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ParametersInCacheKeyAndForwardedToOrigin: min=0
	ParametersInCacheKeyAndForwardedToOrigin []datacloudfrontcachepolicy.ParametersInCacheKeyAndForwardedToOrigin `hcl:"parameters_in_cache_key_and_forwarded_to_origin,block" validate:"min=0"`
}
type dataCloudfrontCachePolicyAttributes struct {
	ref terra.Reference
}

// Comment returns a reference to field comment of aws_cloudfront_cache_policy.
func (ccp dataCloudfrontCachePolicyAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("comment"))
}

// DefaultTtl returns a reference to field default_ttl of aws_cloudfront_cache_policy.
func (ccp dataCloudfrontCachePolicyAttributes) DefaultTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(ccp.ref.Append("default_ttl"))
}

// Etag returns a reference to field etag of aws_cloudfront_cache_policy.
func (ccp dataCloudfrontCachePolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_cache_policy.
func (ccp dataCloudfrontCachePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("id"))
}

// MaxTtl returns a reference to field max_ttl of aws_cloudfront_cache_policy.
func (ccp dataCloudfrontCachePolicyAttributes) MaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(ccp.ref.Append("max_ttl"))
}

// MinTtl returns a reference to field min_ttl of aws_cloudfront_cache_policy.
func (ccp dataCloudfrontCachePolicyAttributes) MinTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(ccp.ref.Append("min_ttl"))
}

// Name returns a reference to field name of aws_cloudfront_cache_policy.
func (ccp dataCloudfrontCachePolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("name"))
}

func (ccp dataCloudfrontCachePolicyAttributes) ParametersInCacheKeyAndForwardedToOrigin() terra.ListValue[datacloudfrontcachepolicy.ParametersInCacheKeyAndForwardedToOriginAttributes] {
	return terra.ReferenceAsList[datacloudfrontcachepolicy.ParametersInCacheKeyAndForwardedToOriginAttributes](ccp.ref.Append("parameters_in_cache_key_and_forwarded_to_origin"))
}
