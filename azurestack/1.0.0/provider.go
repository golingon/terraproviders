// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack

import "github.com/golingon/lingon/pkg/terra"

var _ terra.Provider = (*Provider)(nil)

// Provider contains the configurations for provider.
type Provider struct {
	// ArmEndpoint: string, optional
	ArmEndpoint terra.StringValue `hcl:"arm_endpoint,attr"`
	// AuxiliaryTenantIds: list of string, optional
	AuxiliaryTenantIds terra.ListValue[terra.StringValue] `hcl:"auxiliary_tenant_ids,attr"`
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
	// Environment: string, optional
	Environment terra.StringValue `hcl:"environment,attr"`
	// MetadataHost: string, optional
	MetadataHost terra.StringValue `hcl:"metadata_host,attr"`
	// MsiEndpoint: string, optional
	MsiEndpoint terra.StringValue `hcl:"msi_endpoint,attr"`
	// SkipProviderRegistration: bool, optional
	SkipProviderRegistration terra.BoolValue `hcl:"skip_provider_registration,attr"`
	// SubscriptionId: string, optional
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// UseMsi: bool, optional
	UseMsi terra.BoolValue `hcl:"use_msi,attr"`
	// Features: required
	Features *Features `hcl:"features,block" validate:"required"`
}

// LocalName returns the provider local name for [Provider].
func (p *Provider) LocalName() string {
	return "azurestack"
}

// Source returns the provider source for [Provider].
func (p *Provider) Source() string {
	return "hashicorp/azurestack"
}

// Version returns the provider version for [Provider].
func (p *Provider) Version() string {
	return "1.0.0"
}

// Configuration returns the provider configuration for [Provider].
func (p *Provider) Configuration() interface{} {
	return p
}
