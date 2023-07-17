// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	provider "github.com/golingon/terraproviders/aws/5.8.0/provider"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewProvider(args ProviderArgs) *Provider {
	return &Provider{Args: args}
}

var _ terra.Provider = (*Provider)(nil)

type Provider struct {
	Args ProviderArgs
}

// LocalName returns the provider local name for [Provider].
func (p *Provider) LocalName() string {
	return "aws"
}

// Source returns the provider source for [Provider].
func (p *Provider) Source() string {
	return "hashicorp/aws"
}

// Version returns the provider version for [Provider].
func (p *Provider) Version() string {
	return "5.8.0"
}

// Configuration returns the configuration (args) for [Provider].
func (p *Provider) Configuration() interface{} {
	return p.Args
}

// ProviderArgs contains the configurations for provider.
type ProviderArgs struct {
	// AccessKey: string, optional
	AccessKey terra.StringValue `hcl:"access_key,attr"`
	// AllowedAccountIds: set of string, optional
	AllowedAccountIds terra.SetValue[terra.StringValue] `hcl:"allowed_account_ids,attr"`
	// CustomCaBundle: string, optional
	CustomCaBundle terra.StringValue `hcl:"custom_ca_bundle,attr"`
	// Ec2MetadataServiceEndpoint: string, optional
	Ec2MetadataServiceEndpoint terra.StringValue `hcl:"ec2_metadata_service_endpoint,attr"`
	// Ec2MetadataServiceEndpointMode: string, optional
	Ec2MetadataServiceEndpointMode terra.StringValue `hcl:"ec2_metadata_service_endpoint_mode,attr"`
	// ForbiddenAccountIds: set of string, optional
	ForbiddenAccountIds terra.SetValue[terra.StringValue] `hcl:"forbidden_account_ids,attr"`
	// HttpProxy: string, optional
	HttpProxy terra.StringValue `hcl:"http_proxy,attr"`
	// Insecure: bool, optional
	Insecure terra.BoolValue `hcl:"insecure,attr"`
	// MaxRetries: number, optional
	MaxRetries terra.NumberValue `hcl:"max_retries,attr"`
	// Profile: string, optional
	Profile terra.StringValue `hcl:"profile,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// RetryMode: string, optional
	RetryMode terra.StringValue `hcl:"retry_mode,attr"`
	// S3UsePathStyle: bool, optional
	S3UsePathStyle terra.BoolValue `hcl:"s3_use_path_style,attr"`
	// SecretKey: string, optional
	SecretKey terra.StringValue `hcl:"secret_key,attr"`
	// SharedConfigFiles: list of string, optional
	SharedConfigFiles terra.ListValue[terra.StringValue] `hcl:"shared_config_files,attr"`
	// SharedCredentialsFiles: list of string, optional
	SharedCredentialsFiles terra.ListValue[terra.StringValue] `hcl:"shared_credentials_files,attr"`
	// SkipCredentialsValidation: bool, optional
	SkipCredentialsValidation terra.BoolValue `hcl:"skip_credentials_validation,attr"`
	// SkipMetadataApiCheck: string, optional
	SkipMetadataApiCheck terra.StringValue `hcl:"skip_metadata_api_check,attr"`
	// SkipRegionValidation: bool, optional
	SkipRegionValidation terra.BoolValue `hcl:"skip_region_validation,attr"`
	// SkipRequestingAccountId: bool, optional
	SkipRequestingAccountId terra.BoolValue `hcl:"skip_requesting_account_id,attr"`
	// StsRegion: string, optional
	StsRegion terra.StringValue `hcl:"sts_region,attr"`
	// Token: string, optional
	Token terra.StringValue `hcl:"token,attr"`
	// UseDualstackEndpoint: bool, optional
	UseDualstackEndpoint terra.BoolValue `hcl:"use_dualstack_endpoint,attr"`
	// UseFipsEndpoint: bool, optional
	UseFipsEndpoint terra.BoolValue `hcl:"use_fips_endpoint,attr"`
	// AssumeRole: min=0
	AssumeRole []provider.AssumeRole `hcl:"assume_role,block" validate:"min=0"`
	// AssumeRoleWithWebIdentity: min=0
	AssumeRoleWithWebIdentity []provider.AssumeRoleWithWebIdentity `hcl:"assume_role_with_web_identity,block" validate:"min=0"`
	// DefaultTags: min=0
	DefaultTags []provider.DefaultTags `hcl:"default_tags,block" validate:"min=0"`
	// Endpoints: min=0
	Endpoints []provider.Endpoints `hcl:"endpoints,block" validate:"min=0"`
	// IgnoreTags: min=0
	IgnoreTags []provider.IgnoreTags `hcl:"ignore_tags,block" validate:"min=0"`
}
