// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudfrontresponseheaderspolicy "github.com/golingon/terraproviders/aws/4.66.1/cloudfrontresponseheaderspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfrontResponseHeadersPolicy creates a new instance of [CloudfrontResponseHeadersPolicy].
func NewCloudfrontResponseHeadersPolicy(name string, args CloudfrontResponseHeadersPolicyArgs) *CloudfrontResponseHeadersPolicy {
	return &CloudfrontResponseHeadersPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontResponseHeadersPolicy)(nil)

// CloudfrontResponseHeadersPolicy represents the Terraform resource aws_cloudfront_response_headers_policy.
type CloudfrontResponseHeadersPolicy struct {
	Name      string
	Args      CloudfrontResponseHeadersPolicyArgs
	state     *cloudfrontResponseHeadersPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfrontResponseHeadersPolicy].
func (crhp *CloudfrontResponseHeadersPolicy) Type() string {
	return "aws_cloudfront_response_headers_policy"
}

// LocalName returns the local name for [CloudfrontResponseHeadersPolicy].
func (crhp *CloudfrontResponseHeadersPolicy) LocalName() string {
	return crhp.Name
}

// Configuration returns the configuration (args) for [CloudfrontResponseHeadersPolicy].
func (crhp *CloudfrontResponseHeadersPolicy) Configuration() interface{} {
	return crhp.Args
}

// DependOn is used for other resources to depend on [CloudfrontResponseHeadersPolicy].
func (crhp *CloudfrontResponseHeadersPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(crhp)
}

// Dependencies returns the list of resources [CloudfrontResponseHeadersPolicy] depends_on.
func (crhp *CloudfrontResponseHeadersPolicy) Dependencies() terra.Dependencies {
	return crhp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfrontResponseHeadersPolicy].
func (crhp *CloudfrontResponseHeadersPolicy) LifecycleManagement() *terra.Lifecycle {
	return crhp.Lifecycle
}

// Attributes returns the attributes for [CloudfrontResponseHeadersPolicy].
func (crhp *CloudfrontResponseHeadersPolicy) Attributes() cloudfrontResponseHeadersPolicyAttributes {
	return cloudfrontResponseHeadersPolicyAttributes{ref: terra.ReferenceResource(crhp)}
}

// ImportState imports the given attribute values into [CloudfrontResponseHeadersPolicy]'s state.
func (crhp *CloudfrontResponseHeadersPolicy) ImportState(av io.Reader) error {
	crhp.state = &cloudfrontResponseHeadersPolicyState{}
	if err := json.NewDecoder(av).Decode(crhp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crhp.Type(), crhp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfrontResponseHeadersPolicy] has state.
func (crhp *CloudfrontResponseHeadersPolicy) State() (*cloudfrontResponseHeadersPolicyState, bool) {
	return crhp.state, crhp.state != nil
}

// StateMust returns the state for [CloudfrontResponseHeadersPolicy]. Panics if the state is nil.
func (crhp *CloudfrontResponseHeadersPolicy) StateMust() *cloudfrontResponseHeadersPolicyState {
	if crhp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crhp.Type(), crhp.LocalName()))
	}
	return crhp.state
}

// CloudfrontResponseHeadersPolicyArgs contains the configurations for aws_cloudfront_response_headers_policy.
type CloudfrontResponseHeadersPolicyArgs struct {
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// Etag: string, optional
	Etag terra.StringValue `hcl:"etag,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// CorsConfig: optional
	CorsConfig *cloudfrontresponseheaderspolicy.CorsConfig `hcl:"cors_config,block"`
	// CustomHeadersConfig: optional
	CustomHeadersConfig *cloudfrontresponseheaderspolicy.CustomHeadersConfig `hcl:"custom_headers_config,block"`
	// RemoveHeadersConfig: optional
	RemoveHeadersConfig *cloudfrontresponseheaderspolicy.RemoveHeadersConfig `hcl:"remove_headers_config,block"`
	// SecurityHeadersConfig: optional
	SecurityHeadersConfig *cloudfrontresponseheaderspolicy.SecurityHeadersConfig `hcl:"security_headers_config,block"`
	// ServerTimingHeadersConfig: optional
	ServerTimingHeadersConfig *cloudfrontresponseheaderspolicy.ServerTimingHeadersConfig `hcl:"server_timing_headers_config,block"`
}
type cloudfrontResponseHeadersPolicyAttributes struct {
	ref terra.Reference
}

// Comment returns a reference to field comment of aws_cloudfront_response_headers_policy.
func (crhp cloudfrontResponseHeadersPolicyAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(crhp.ref.Append("comment"))
}

// Etag returns a reference to field etag of aws_cloudfront_response_headers_policy.
func (crhp cloudfrontResponseHeadersPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crhp.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_response_headers_policy.
func (crhp cloudfrontResponseHeadersPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crhp.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudfront_response_headers_policy.
func (crhp cloudfrontResponseHeadersPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crhp.ref.Append("name"))
}

func (crhp cloudfrontResponseHeadersPolicyAttributes) CorsConfig() terra.ListValue[cloudfrontresponseheaderspolicy.CorsConfigAttributes] {
	return terra.ReferenceAsList[cloudfrontresponseheaderspolicy.CorsConfigAttributes](crhp.ref.Append("cors_config"))
}

func (crhp cloudfrontResponseHeadersPolicyAttributes) CustomHeadersConfig() terra.ListValue[cloudfrontresponseheaderspolicy.CustomHeadersConfigAttributes] {
	return terra.ReferenceAsList[cloudfrontresponseheaderspolicy.CustomHeadersConfigAttributes](crhp.ref.Append("custom_headers_config"))
}

func (crhp cloudfrontResponseHeadersPolicyAttributes) RemoveHeadersConfig() terra.ListValue[cloudfrontresponseheaderspolicy.RemoveHeadersConfigAttributes] {
	return terra.ReferenceAsList[cloudfrontresponseheaderspolicy.RemoveHeadersConfigAttributes](crhp.ref.Append("remove_headers_config"))
}

func (crhp cloudfrontResponseHeadersPolicyAttributes) SecurityHeadersConfig() terra.ListValue[cloudfrontresponseheaderspolicy.SecurityHeadersConfigAttributes] {
	return terra.ReferenceAsList[cloudfrontresponseheaderspolicy.SecurityHeadersConfigAttributes](crhp.ref.Append("security_headers_config"))
}

func (crhp cloudfrontResponseHeadersPolicyAttributes) ServerTimingHeadersConfig() terra.ListValue[cloudfrontresponseheaderspolicy.ServerTimingHeadersConfigAttributes] {
	return terra.ReferenceAsList[cloudfrontresponseheaderspolicy.ServerTimingHeadersConfigAttributes](crhp.ref.Append("server_timing_headers_config"))
}

type cloudfrontResponseHeadersPolicyState struct {
	Comment                   string                                                           `json:"comment"`
	Etag                      string                                                           `json:"etag"`
	Id                        string                                                           `json:"id"`
	Name                      string                                                           `json:"name"`
	CorsConfig                []cloudfrontresponseheaderspolicy.CorsConfigState                `json:"cors_config"`
	CustomHeadersConfig       []cloudfrontresponseheaderspolicy.CustomHeadersConfigState       `json:"custom_headers_config"`
	RemoveHeadersConfig       []cloudfrontresponseheaderspolicy.RemoveHeadersConfigState       `json:"remove_headers_config"`
	SecurityHeadersConfig     []cloudfrontresponseheaderspolicy.SecurityHeadersConfigState     `json:"security_headers_config"`
	ServerTimingHeadersConfig []cloudfrontresponseheaderspolicy.ServerTimingHeadersConfigState `json:"server_timing_headers_config"`
}
