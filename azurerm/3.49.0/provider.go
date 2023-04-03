// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	provider "github.com/golingon/terraproviders/azurerm/3.49.0/provider"
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
	return "azurerm"
}

// Source returns the provider source for [Provider].
func (p *Provider) Source() string {
	return "hashicorp/azurerm"
}

// Version returns the provider version for [Provider].
func (p *Provider) Version() string {
	return "3.49.0"
}

// Configuration returns the configuration (args) for [Provider].
func (p *Provider) Configuration() interface{} {
	return p.Args
}

// ProviderArgs contains the configurations for provider.
type ProviderArgs struct {
	// AuxiliaryTenantIds: list of string, optional
	AuxiliaryTenantIds terra.ListValue[terra.StringValue] `hcl:"auxiliary_tenant_ids,attr"`
	// ClientCertificate: string, optional
	ClientCertificate terra.StringValue `hcl:"client_certificate,attr"`
	// ClientCertificatePassword: string, optional
	ClientCertificatePassword terra.StringValue `hcl:"client_certificate_password,attr"`
	// ClientCertificatePath: string, optional
	ClientCertificatePath terra.StringValue `hcl:"client_certificate_path,attr"`
	// ClientId: string, optional
	ClientId terra.StringValue `hcl:"client_id,attr"`
	// ClientSecret: string, optional
	ClientSecret terra.StringValue `hcl:"client_secret,attr"`
	// DisableCorrelationRequestId: bool, optional
	DisableCorrelationRequestId terra.BoolValue `hcl:"disable_correlation_request_id,attr"`
	// DisableTerraformPartnerId: bool, optional
	DisableTerraformPartnerId terra.BoolValue `hcl:"disable_terraform_partner_id,attr"`
	// Environment: string, optional
	Environment terra.StringValue `hcl:"environment,attr"`
	// MetadataHost: string, optional
	MetadataHost terra.StringValue `hcl:"metadata_host,attr"`
	// MsiEndpoint: string, optional
	MsiEndpoint terra.StringValue `hcl:"msi_endpoint,attr"`
	// OidcRequestToken: string, optional
	OidcRequestToken terra.StringValue `hcl:"oidc_request_token,attr"`
	// OidcRequestUrl: string, optional
	OidcRequestUrl terra.StringValue `hcl:"oidc_request_url,attr"`
	// OidcToken: string, optional
	OidcToken terra.StringValue `hcl:"oidc_token,attr"`
	// OidcTokenFilePath: string, optional
	OidcTokenFilePath terra.StringValue `hcl:"oidc_token_file_path,attr"`
	// PartnerId: string, optional
	PartnerId terra.StringValue `hcl:"partner_id,attr"`
	// SkipProviderRegistration: bool, optional
	SkipProviderRegistration terra.BoolValue `hcl:"skip_provider_registration,attr"`
	// StorageUseAzuread: bool, optional
	StorageUseAzuread terra.BoolValue `hcl:"storage_use_azuread,attr"`
	// SubscriptionId: string, optional
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// UseCli: bool, optional
	UseCli terra.BoolValue `hcl:"use_cli,attr"`
	// UseMsi: bool, optional
	UseMsi terra.BoolValue `hcl:"use_msi,attr"`
	// UseOidc: bool, optional
	UseOidc terra.BoolValue `hcl:"use_oidc,attr"`
	// Features: required
	Features *provider.Features `hcl:"features,block" validate:"required"`
}
