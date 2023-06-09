// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudfrontoriginrequestpolicy "github.com/golingon/terraproviders/aws/4.66.1/cloudfrontoriginrequestpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfrontOriginRequestPolicy creates a new instance of [CloudfrontOriginRequestPolicy].
func NewCloudfrontOriginRequestPolicy(name string, args CloudfrontOriginRequestPolicyArgs) *CloudfrontOriginRequestPolicy {
	return &CloudfrontOriginRequestPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontOriginRequestPolicy)(nil)

// CloudfrontOriginRequestPolicy represents the Terraform resource aws_cloudfront_origin_request_policy.
type CloudfrontOriginRequestPolicy struct {
	Name      string
	Args      CloudfrontOriginRequestPolicyArgs
	state     *cloudfrontOriginRequestPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfrontOriginRequestPolicy].
func (corp *CloudfrontOriginRequestPolicy) Type() string {
	return "aws_cloudfront_origin_request_policy"
}

// LocalName returns the local name for [CloudfrontOriginRequestPolicy].
func (corp *CloudfrontOriginRequestPolicy) LocalName() string {
	return corp.Name
}

// Configuration returns the configuration (args) for [CloudfrontOriginRequestPolicy].
func (corp *CloudfrontOriginRequestPolicy) Configuration() interface{} {
	return corp.Args
}

// DependOn is used for other resources to depend on [CloudfrontOriginRequestPolicy].
func (corp *CloudfrontOriginRequestPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(corp)
}

// Dependencies returns the list of resources [CloudfrontOriginRequestPolicy] depends_on.
func (corp *CloudfrontOriginRequestPolicy) Dependencies() terra.Dependencies {
	return corp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfrontOriginRequestPolicy].
func (corp *CloudfrontOriginRequestPolicy) LifecycleManagement() *terra.Lifecycle {
	return corp.Lifecycle
}

// Attributes returns the attributes for [CloudfrontOriginRequestPolicy].
func (corp *CloudfrontOriginRequestPolicy) Attributes() cloudfrontOriginRequestPolicyAttributes {
	return cloudfrontOriginRequestPolicyAttributes{ref: terra.ReferenceResource(corp)}
}

// ImportState imports the given attribute values into [CloudfrontOriginRequestPolicy]'s state.
func (corp *CloudfrontOriginRequestPolicy) ImportState(av io.Reader) error {
	corp.state = &cloudfrontOriginRequestPolicyState{}
	if err := json.NewDecoder(av).Decode(corp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", corp.Type(), corp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfrontOriginRequestPolicy] has state.
func (corp *CloudfrontOriginRequestPolicy) State() (*cloudfrontOriginRequestPolicyState, bool) {
	return corp.state, corp.state != nil
}

// StateMust returns the state for [CloudfrontOriginRequestPolicy]. Panics if the state is nil.
func (corp *CloudfrontOriginRequestPolicy) StateMust() *cloudfrontOriginRequestPolicyState {
	if corp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", corp.Type(), corp.LocalName()))
	}
	return corp.state
}

// CloudfrontOriginRequestPolicyArgs contains the configurations for aws_cloudfront_origin_request_policy.
type CloudfrontOriginRequestPolicyArgs struct {
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// CookiesConfig: required
	CookiesConfig *cloudfrontoriginrequestpolicy.CookiesConfig `hcl:"cookies_config,block" validate:"required"`
	// HeadersConfig: required
	HeadersConfig *cloudfrontoriginrequestpolicy.HeadersConfig `hcl:"headers_config,block" validate:"required"`
	// QueryStringsConfig: required
	QueryStringsConfig *cloudfrontoriginrequestpolicy.QueryStringsConfig `hcl:"query_strings_config,block" validate:"required"`
}
type cloudfrontOriginRequestPolicyAttributes struct {
	ref terra.Reference
}

// Comment returns a reference to field comment of aws_cloudfront_origin_request_policy.
func (corp cloudfrontOriginRequestPolicyAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(corp.ref.Append("comment"))
}

// Etag returns a reference to field etag of aws_cloudfront_origin_request_policy.
func (corp cloudfrontOriginRequestPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(corp.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_origin_request_policy.
func (corp cloudfrontOriginRequestPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(corp.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudfront_origin_request_policy.
func (corp cloudfrontOriginRequestPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(corp.ref.Append("name"))
}

func (corp cloudfrontOriginRequestPolicyAttributes) CookiesConfig() terra.ListValue[cloudfrontoriginrequestpolicy.CookiesConfigAttributes] {
	return terra.ReferenceAsList[cloudfrontoriginrequestpolicy.CookiesConfigAttributes](corp.ref.Append("cookies_config"))
}

func (corp cloudfrontOriginRequestPolicyAttributes) HeadersConfig() terra.ListValue[cloudfrontoriginrequestpolicy.HeadersConfigAttributes] {
	return terra.ReferenceAsList[cloudfrontoriginrequestpolicy.HeadersConfigAttributes](corp.ref.Append("headers_config"))
}

func (corp cloudfrontOriginRequestPolicyAttributes) QueryStringsConfig() terra.ListValue[cloudfrontoriginrequestpolicy.QueryStringsConfigAttributes] {
	return terra.ReferenceAsList[cloudfrontoriginrequestpolicy.QueryStringsConfigAttributes](corp.ref.Append("query_strings_config"))
}

type cloudfrontOriginRequestPolicyState struct {
	Comment            string                                                  `json:"comment"`
	Etag               string                                                  `json:"etag"`
	Id                 string                                                  `json:"id"`
	Name               string                                                  `json:"name"`
	CookiesConfig      []cloudfrontoriginrequestpolicy.CookiesConfigState      `json:"cookies_config"`
	HeadersConfig      []cloudfrontoriginrequestpolicy.HeadersConfigState      `json:"headers_config"`
	QueryStringsConfig []cloudfrontoriginrequestpolicy.QueryStringsConfigState `json:"query_strings_config"`
}
