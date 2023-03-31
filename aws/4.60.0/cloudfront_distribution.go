// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudfrontdistribution "github.com/golingon/terraproviders/aws/4.60.0/cloudfrontdistribution"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfrontDistribution creates a new instance of [CloudfrontDistribution].
func NewCloudfrontDistribution(name string, args CloudfrontDistributionArgs) *CloudfrontDistribution {
	return &CloudfrontDistribution{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontDistribution)(nil)

// CloudfrontDistribution represents the Terraform resource aws_cloudfront_distribution.
type CloudfrontDistribution struct {
	Name      string
	Args      CloudfrontDistributionArgs
	state     *cloudfrontDistributionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfrontDistribution].
func (cd *CloudfrontDistribution) Type() string {
	return "aws_cloudfront_distribution"
}

// LocalName returns the local name for [CloudfrontDistribution].
func (cd *CloudfrontDistribution) LocalName() string {
	return cd.Name
}

// Configuration returns the configuration (args) for [CloudfrontDistribution].
func (cd *CloudfrontDistribution) Configuration() interface{} {
	return cd.Args
}

// DependOn is used for other resources to depend on [CloudfrontDistribution].
func (cd *CloudfrontDistribution) DependOn() terra.Reference {
	return terra.ReferenceResource(cd)
}

// Dependencies returns the list of resources [CloudfrontDistribution] depends_on.
func (cd *CloudfrontDistribution) Dependencies() terra.Dependencies {
	return cd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfrontDistribution].
func (cd *CloudfrontDistribution) LifecycleManagement() *terra.Lifecycle {
	return cd.Lifecycle
}

// Attributes returns the attributes for [CloudfrontDistribution].
func (cd *CloudfrontDistribution) Attributes() cloudfrontDistributionAttributes {
	return cloudfrontDistributionAttributes{ref: terra.ReferenceResource(cd)}
}

// ImportState imports the given attribute values into [CloudfrontDistribution]'s state.
func (cd *CloudfrontDistribution) ImportState(av io.Reader) error {
	cd.state = &cloudfrontDistributionState{}
	if err := json.NewDecoder(av).Decode(cd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cd.Type(), cd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfrontDistribution] has state.
func (cd *CloudfrontDistribution) State() (*cloudfrontDistributionState, bool) {
	return cd.state, cd.state != nil
}

// StateMust returns the state for [CloudfrontDistribution]. Panics if the state is nil.
func (cd *CloudfrontDistribution) StateMust() *cloudfrontDistributionState {
	if cd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cd.Type(), cd.LocalName()))
	}
	return cd.state
}

// CloudfrontDistributionArgs contains the configurations for aws_cloudfront_distribution.
type CloudfrontDistributionArgs struct {
	// Aliases: set of string, optional
	Aliases terra.SetValue[terra.StringValue] `hcl:"aliases,attr"`
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// DefaultRootObject: string, optional
	DefaultRootObject terra.StringValue `hcl:"default_root_object,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// HttpVersion: string, optional
	HttpVersion terra.StringValue `hcl:"http_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsIpv6Enabled: bool, optional
	IsIpv6Enabled terra.BoolValue `hcl:"is_ipv6_enabled,attr"`
	// PriceClass: string, optional
	PriceClass terra.StringValue `hcl:"price_class,attr"`
	// RetainOnDelete: bool, optional
	RetainOnDelete terra.BoolValue `hcl:"retain_on_delete,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// WaitForDeployment: bool, optional
	WaitForDeployment terra.BoolValue `hcl:"wait_for_deployment,attr"`
	// WebAclId: string, optional
	WebAclId terra.StringValue `hcl:"web_acl_id,attr"`
	// TrustedKeyGroups: min=0
	TrustedKeyGroups []cloudfrontdistribution.TrustedKeyGroups `hcl:"trusted_key_groups,block" validate:"min=0"`
	// TrustedSigners: min=0
	TrustedSigners []cloudfrontdistribution.TrustedSigners `hcl:"trusted_signers,block" validate:"min=0"`
	// CustomErrorResponse: min=0
	CustomErrorResponse []cloudfrontdistribution.CustomErrorResponse `hcl:"custom_error_response,block" validate:"min=0"`
	// DefaultCacheBehavior: required
	DefaultCacheBehavior *cloudfrontdistribution.DefaultCacheBehavior `hcl:"default_cache_behavior,block" validate:"required"`
	// LoggingConfig: optional
	LoggingConfig *cloudfrontdistribution.LoggingConfig `hcl:"logging_config,block"`
	// OrderedCacheBehavior: min=0
	OrderedCacheBehavior []cloudfrontdistribution.OrderedCacheBehavior `hcl:"ordered_cache_behavior,block" validate:"min=0"`
	// Origin: min=1
	Origin []cloudfrontdistribution.Origin `hcl:"origin,block" validate:"min=1"`
	// OriginGroup: min=0
	OriginGroup []cloudfrontdistribution.OriginGroup `hcl:"origin_group,block" validate:"min=0"`
	// Restrictions: required
	Restrictions *cloudfrontdistribution.Restrictions `hcl:"restrictions,block" validate:"required"`
	// ViewerCertificate: required
	ViewerCertificate *cloudfrontdistribution.ViewerCertificate `hcl:"viewer_certificate,block" validate:"required"`
}
type cloudfrontDistributionAttributes struct {
	ref terra.Reference
}

// Aliases returns a reference to field aliases of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) Aliases() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cd.ref.Append("aliases"))
}

// Arn returns a reference to field arn of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("arn"))
}

// CallerReference returns a reference to field caller_reference of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) CallerReference() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("caller_reference"))
}

// Comment returns a reference to field comment of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("comment"))
}

// DefaultRootObject returns a reference to field default_root_object of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) DefaultRootObject() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("default_root_object"))
}

// DomainName returns a reference to field domain_name of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("domain_name"))
}

// Enabled returns a reference to field enabled of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cd.ref.Append("enabled"))
}

// Etag returns a reference to field etag of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("etag"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("hosted_zone_id"))
}

// HttpVersion returns a reference to field http_version of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) HttpVersion() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("http_version"))
}

// Id returns a reference to field id of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("id"))
}

// InProgressValidationBatches returns a reference to field in_progress_validation_batches of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) InProgressValidationBatches() terra.NumberValue {
	return terra.ReferenceAsNumber(cd.ref.Append("in_progress_validation_batches"))
}

// IsIpv6Enabled returns a reference to field is_ipv6_enabled of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) IsIpv6Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cd.ref.Append("is_ipv6_enabled"))
}

// LastModifiedTime returns a reference to field last_modified_time of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_modified_time"))
}

// PriceClass returns a reference to field price_class of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) PriceClass() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("price_class"))
}

// RetainOnDelete returns a reference to field retain_on_delete of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) RetainOnDelete() terra.BoolValue {
	return terra.ReferenceAsBool(cd.ref.Append("retain_on_delete"))
}

// Status returns a reference to field status of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cd.ref.Append("tags_all"))
}

// WaitForDeployment returns a reference to field wait_for_deployment of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) WaitForDeployment() terra.BoolValue {
	return terra.ReferenceAsBool(cd.ref.Append("wait_for_deployment"))
}

// WebAclId returns a reference to field web_acl_id of aws_cloudfront_distribution.
func (cd cloudfrontDistributionAttributes) WebAclId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("web_acl_id"))
}

func (cd cloudfrontDistributionAttributes) TrustedKeyGroups() terra.ListValue[cloudfrontdistribution.TrustedKeyGroupsAttributes] {
	return terra.ReferenceAsList[cloudfrontdistribution.TrustedKeyGroupsAttributes](cd.ref.Append("trusted_key_groups"))
}

func (cd cloudfrontDistributionAttributes) TrustedSigners() terra.ListValue[cloudfrontdistribution.TrustedSignersAttributes] {
	return terra.ReferenceAsList[cloudfrontdistribution.TrustedSignersAttributes](cd.ref.Append("trusted_signers"))
}

func (cd cloudfrontDistributionAttributes) CustomErrorResponse() terra.SetValue[cloudfrontdistribution.CustomErrorResponseAttributes] {
	return terra.ReferenceAsSet[cloudfrontdistribution.CustomErrorResponseAttributes](cd.ref.Append("custom_error_response"))
}

func (cd cloudfrontDistributionAttributes) DefaultCacheBehavior() terra.ListValue[cloudfrontdistribution.DefaultCacheBehaviorAttributes] {
	return terra.ReferenceAsList[cloudfrontdistribution.DefaultCacheBehaviorAttributes](cd.ref.Append("default_cache_behavior"))
}

func (cd cloudfrontDistributionAttributes) LoggingConfig() terra.ListValue[cloudfrontdistribution.LoggingConfigAttributes] {
	return terra.ReferenceAsList[cloudfrontdistribution.LoggingConfigAttributes](cd.ref.Append("logging_config"))
}

func (cd cloudfrontDistributionAttributes) OrderedCacheBehavior() terra.ListValue[cloudfrontdistribution.OrderedCacheBehaviorAttributes] {
	return terra.ReferenceAsList[cloudfrontdistribution.OrderedCacheBehaviorAttributes](cd.ref.Append("ordered_cache_behavior"))
}

func (cd cloudfrontDistributionAttributes) Origin() terra.SetValue[cloudfrontdistribution.OriginAttributes] {
	return terra.ReferenceAsSet[cloudfrontdistribution.OriginAttributes](cd.ref.Append("origin"))
}

func (cd cloudfrontDistributionAttributes) OriginGroup() terra.SetValue[cloudfrontdistribution.OriginGroupAttributes] {
	return terra.ReferenceAsSet[cloudfrontdistribution.OriginGroupAttributes](cd.ref.Append("origin_group"))
}

func (cd cloudfrontDistributionAttributes) Restrictions() terra.ListValue[cloudfrontdistribution.RestrictionsAttributes] {
	return terra.ReferenceAsList[cloudfrontdistribution.RestrictionsAttributes](cd.ref.Append("restrictions"))
}

func (cd cloudfrontDistributionAttributes) ViewerCertificate() terra.ListValue[cloudfrontdistribution.ViewerCertificateAttributes] {
	return terra.ReferenceAsList[cloudfrontdistribution.ViewerCertificateAttributes](cd.ref.Append("viewer_certificate"))
}

type cloudfrontDistributionState struct {
	Aliases                     []string                                           `json:"aliases"`
	Arn                         string                                             `json:"arn"`
	CallerReference             string                                             `json:"caller_reference"`
	Comment                     string                                             `json:"comment"`
	DefaultRootObject           string                                             `json:"default_root_object"`
	DomainName                  string                                             `json:"domain_name"`
	Enabled                     bool                                               `json:"enabled"`
	Etag                        string                                             `json:"etag"`
	HostedZoneId                string                                             `json:"hosted_zone_id"`
	HttpVersion                 string                                             `json:"http_version"`
	Id                          string                                             `json:"id"`
	InProgressValidationBatches float64                                            `json:"in_progress_validation_batches"`
	IsIpv6Enabled               bool                                               `json:"is_ipv6_enabled"`
	LastModifiedTime            string                                             `json:"last_modified_time"`
	PriceClass                  string                                             `json:"price_class"`
	RetainOnDelete              bool                                               `json:"retain_on_delete"`
	Status                      string                                             `json:"status"`
	Tags                        map[string]string                                  `json:"tags"`
	TagsAll                     map[string]string                                  `json:"tags_all"`
	WaitForDeployment           bool                                               `json:"wait_for_deployment"`
	WebAclId                    string                                             `json:"web_acl_id"`
	TrustedKeyGroups            []cloudfrontdistribution.TrustedKeyGroupsState     `json:"trusted_key_groups"`
	TrustedSigners              []cloudfrontdistribution.TrustedSignersState       `json:"trusted_signers"`
	CustomErrorResponse         []cloudfrontdistribution.CustomErrorResponseState  `json:"custom_error_response"`
	DefaultCacheBehavior        []cloudfrontdistribution.DefaultCacheBehaviorState `json:"default_cache_behavior"`
	LoggingConfig               []cloudfrontdistribution.LoggingConfigState        `json:"logging_config"`
	OrderedCacheBehavior        []cloudfrontdistribution.OrderedCacheBehaviorState `json:"ordered_cache_behavior"`
	Origin                      []cloudfrontdistribution.OriginState               `json:"origin"`
	OriginGroup                 []cloudfrontdistribution.OriginGroupState          `json:"origin_group"`
	Restrictions                []cloudfrontdistribution.RestrictionsState         `json:"restrictions"`
	ViewerCertificate           []cloudfrontdistribution.ViewerCertificateState    `json:"viewer_certificate"`
}
