// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudfrontcachepolicy "github.com/golingon/terraproviders/aws/5.11.0/cloudfrontcachepolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfrontCachePolicy creates a new instance of [CloudfrontCachePolicy].
func NewCloudfrontCachePolicy(name string, args CloudfrontCachePolicyArgs) *CloudfrontCachePolicy {
	return &CloudfrontCachePolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontCachePolicy)(nil)

// CloudfrontCachePolicy represents the Terraform resource aws_cloudfront_cache_policy.
type CloudfrontCachePolicy struct {
	Name      string
	Args      CloudfrontCachePolicyArgs
	state     *cloudfrontCachePolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfrontCachePolicy].
func (ccp *CloudfrontCachePolicy) Type() string {
	return "aws_cloudfront_cache_policy"
}

// LocalName returns the local name for [CloudfrontCachePolicy].
func (ccp *CloudfrontCachePolicy) LocalName() string {
	return ccp.Name
}

// Configuration returns the configuration (args) for [CloudfrontCachePolicy].
func (ccp *CloudfrontCachePolicy) Configuration() interface{} {
	return ccp.Args
}

// DependOn is used for other resources to depend on [CloudfrontCachePolicy].
func (ccp *CloudfrontCachePolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ccp)
}

// Dependencies returns the list of resources [CloudfrontCachePolicy] depends_on.
func (ccp *CloudfrontCachePolicy) Dependencies() terra.Dependencies {
	return ccp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfrontCachePolicy].
func (ccp *CloudfrontCachePolicy) LifecycleManagement() *terra.Lifecycle {
	return ccp.Lifecycle
}

// Attributes returns the attributes for [CloudfrontCachePolicy].
func (ccp *CloudfrontCachePolicy) Attributes() cloudfrontCachePolicyAttributes {
	return cloudfrontCachePolicyAttributes{ref: terra.ReferenceResource(ccp)}
}

// ImportState imports the given attribute values into [CloudfrontCachePolicy]'s state.
func (ccp *CloudfrontCachePolicy) ImportState(av io.Reader) error {
	ccp.state = &cloudfrontCachePolicyState{}
	if err := json.NewDecoder(av).Decode(ccp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccp.Type(), ccp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfrontCachePolicy] has state.
func (ccp *CloudfrontCachePolicy) State() (*cloudfrontCachePolicyState, bool) {
	return ccp.state, ccp.state != nil
}

// StateMust returns the state for [CloudfrontCachePolicy]. Panics if the state is nil.
func (ccp *CloudfrontCachePolicy) StateMust() *cloudfrontCachePolicyState {
	if ccp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccp.Type(), ccp.LocalName()))
	}
	return ccp.state
}

// CloudfrontCachePolicyArgs contains the configurations for aws_cloudfront_cache_policy.
type CloudfrontCachePolicyArgs struct {
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// DefaultTtl: number, optional
	DefaultTtl terra.NumberValue `hcl:"default_ttl,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxTtl: number, optional
	MaxTtl terra.NumberValue `hcl:"max_ttl,attr"`
	// MinTtl: number, optional
	MinTtl terra.NumberValue `hcl:"min_ttl,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParametersInCacheKeyAndForwardedToOrigin: required
	ParametersInCacheKeyAndForwardedToOrigin *cloudfrontcachepolicy.ParametersInCacheKeyAndForwardedToOrigin `hcl:"parameters_in_cache_key_and_forwarded_to_origin,block" validate:"required"`
}
type cloudfrontCachePolicyAttributes struct {
	ref terra.Reference
}

// Comment returns a reference to field comment of aws_cloudfront_cache_policy.
func (ccp cloudfrontCachePolicyAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("comment"))
}

// DefaultTtl returns a reference to field default_ttl of aws_cloudfront_cache_policy.
func (ccp cloudfrontCachePolicyAttributes) DefaultTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(ccp.ref.Append("default_ttl"))
}

// Etag returns a reference to field etag of aws_cloudfront_cache_policy.
func (ccp cloudfrontCachePolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_cache_policy.
func (ccp cloudfrontCachePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("id"))
}

// MaxTtl returns a reference to field max_ttl of aws_cloudfront_cache_policy.
func (ccp cloudfrontCachePolicyAttributes) MaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(ccp.ref.Append("max_ttl"))
}

// MinTtl returns a reference to field min_ttl of aws_cloudfront_cache_policy.
func (ccp cloudfrontCachePolicyAttributes) MinTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(ccp.ref.Append("min_ttl"))
}

// Name returns a reference to field name of aws_cloudfront_cache_policy.
func (ccp cloudfrontCachePolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("name"))
}

func (ccp cloudfrontCachePolicyAttributes) ParametersInCacheKeyAndForwardedToOrigin() terra.ListValue[cloudfrontcachepolicy.ParametersInCacheKeyAndForwardedToOriginAttributes] {
	return terra.ReferenceAsList[cloudfrontcachepolicy.ParametersInCacheKeyAndForwardedToOriginAttributes](ccp.ref.Append("parameters_in_cache_key_and_forwarded_to_origin"))
}

type cloudfrontCachePolicyState struct {
	Comment                                  string                                                                `json:"comment"`
	DefaultTtl                               float64                                                               `json:"default_ttl"`
	Etag                                     string                                                                `json:"etag"`
	Id                                       string                                                                `json:"id"`
	MaxTtl                                   float64                                                               `json:"max_ttl"`
	MinTtl                                   float64                                                               `json:"min_ttl"`
	Name                                     string                                                                `json:"name"`
	ParametersInCacheKeyAndForwardedToOrigin []cloudfrontcachepolicy.ParametersInCacheKeyAndForwardedToOriginState `json:"parameters_in_cache_key_and_forwarded_to_origin"`
}
