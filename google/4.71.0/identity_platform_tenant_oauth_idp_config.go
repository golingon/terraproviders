// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	identityplatformtenantoauthidpconfig "github.com/golingon/terraproviders/google/4.71.0/identityplatformtenantoauthidpconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityPlatformTenantOauthIdpConfig creates a new instance of [IdentityPlatformTenantOauthIdpConfig].
func NewIdentityPlatformTenantOauthIdpConfig(name string, args IdentityPlatformTenantOauthIdpConfigArgs) *IdentityPlatformTenantOauthIdpConfig {
	return &IdentityPlatformTenantOauthIdpConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityPlatformTenantOauthIdpConfig)(nil)

// IdentityPlatformTenantOauthIdpConfig represents the Terraform resource google_identity_platform_tenant_oauth_idp_config.
type IdentityPlatformTenantOauthIdpConfig struct {
	Name      string
	Args      IdentityPlatformTenantOauthIdpConfigArgs
	state     *identityPlatformTenantOauthIdpConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityPlatformTenantOauthIdpConfig].
func (iptoic *IdentityPlatformTenantOauthIdpConfig) Type() string {
	return "google_identity_platform_tenant_oauth_idp_config"
}

// LocalName returns the local name for [IdentityPlatformTenantOauthIdpConfig].
func (iptoic *IdentityPlatformTenantOauthIdpConfig) LocalName() string {
	return iptoic.Name
}

// Configuration returns the configuration (args) for [IdentityPlatformTenantOauthIdpConfig].
func (iptoic *IdentityPlatformTenantOauthIdpConfig) Configuration() interface{} {
	return iptoic.Args
}

// DependOn is used for other resources to depend on [IdentityPlatformTenantOauthIdpConfig].
func (iptoic *IdentityPlatformTenantOauthIdpConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(iptoic)
}

// Dependencies returns the list of resources [IdentityPlatformTenantOauthIdpConfig] depends_on.
func (iptoic *IdentityPlatformTenantOauthIdpConfig) Dependencies() terra.Dependencies {
	return iptoic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityPlatformTenantOauthIdpConfig].
func (iptoic *IdentityPlatformTenantOauthIdpConfig) LifecycleManagement() *terra.Lifecycle {
	return iptoic.Lifecycle
}

// Attributes returns the attributes for [IdentityPlatformTenantOauthIdpConfig].
func (iptoic *IdentityPlatformTenantOauthIdpConfig) Attributes() identityPlatformTenantOauthIdpConfigAttributes {
	return identityPlatformTenantOauthIdpConfigAttributes{ref: terra.ReferenceResource(iptoic)}
}

// ImportState imports the given attribute values into [IdentityPlatformTenantOauthIdpConfig]'s state.
func (iptoic *IdentityPlatformTenantOauthIdpConfig) ImportState(av io.Reader) error {
	iptoic.state = &identityPlatformTenantOauthIdpConfigState{}
	if err := json.NewDecoder(av).Decode(iptoic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iptoic.Type(), iptoic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityPlatformTenantOauthIdpConfig] has state.
func (iptoic *IdentityPlatformTenantOauthIdpConfig) State() (*identityPlatformTenantOauthIdpConfigState, bool) {
	return iptoic.state, iptoic.state != nil
}

// StateMust returns the state for [IdentityPlatformTenantOauthIdpConfig]. Panics if the state is nil.
func (iptoic *IdentityPlatformTenantOauthIdpConfig) StateMust() *identityPlatformTenantOauthIdpConfigState {
	if iptoic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iptoic.Type(), iptoic.LocalName()))
	}
	return iptoic.state
}

// IdentityPlatformTenantOauthIdpConfigArgs contains the configurations for google_identity_platform_tenant_oauth_idp_config.
type IdentityPlatformTenantOauthIdpConfigArgs struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, optional
	ClientSecret terra.StringValue `hcl:"client_secret,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Issuer: string, required
	Issuer terra.StringValue `hcl:"issuer,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Tenant: string, required
	Tenant terra.StringValue `hcl:"tenant,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *identityplatformtenantoauthidpconfig.Timeouts `hcl:"timeouts,block"`
}
type identityPlatformTenantOauthIdpConfigAttributes struct {
	ref terra.Reference
}

// ClientId returns a reference to field client_id of google_identity_platform_tenant_oauth_idp_config.
func (iptoic identityPlatformTenantOauthIdpConfigAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(iptoic.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of google_identity_platform_tenant_oauth_idp_config.
func (iptoic identityPlatformTenantOauthIdpConfigAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(iptoic.ref.Append("client_secret"))
}

// DisplayName returns a reference to field display_name of google_identity_platform_tenant_oauth_idp_config.
func (iptoic identityPlatformTenantOauthIdpConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(iptoic.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of google_identity_platform_tenant_oauth_idp_config.
func (iptoic identityPlatformTenantOauthIdpConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(iptoic.ref.Append("enabled"))
}

// Id returns a reference to field id of google_identity_platform_tenant_oauth_idp_config.
func (iptoic identityPlatformTenantOauthIdpConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iptoic.ref.Append("id"))
}

// Issuer returns a reference to field issuer of google_identity_platform_tenant_oauth_idp_config.
func (iptoic identityPlatformTenantOauthIdpConfigAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(iptoic.ref.Append("issuer"))
}

// Name returns a reference to field name of google_identity_platform_tenant_oauth_idp_config.
func (iptoic identityPlatformTenantOauthIdpConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iptoic.ref.Append("name"))
}

// Project returns a reference to field project of google_identity_platform_tenant_oauth_idp_config.
func (iptoic identityPlatformTenantOauthIdpConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iptoic.ref.Append("project"))
}

// Tenant returns a reference to field tenant of google_identity_platform_tenant_oauth_idp_config.
func (iptoic identityPlatformTenantOauthIdpConfigAttributes) Tenant() terra.StringValue {
	return terra.ReferenceAsString(iptoic.ref.Append("tenant"))
}

func (iptoic identityPlatformTenantOauthIdpConfigAttributes) Timeouts() identityplatformtenantoauthidpconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[identityplatformtenantoauthidpconfig.TimeoutsAttributes](iptoic.ref.Append("timeouts"))
}

type identityPlatformTenantOauthIdpConfigState struct {
	ClientId     string                                              `json:"client_id"`
	ClientSecret string                                              `json:"client_secret"`
	DisplayName  string                                              `json:"display_name"`
	Enabled      bool                                                `json:"enabled"`
	Id           string                                              `json:"id"`
	Issuer       string                                              `json:"issuer"`
	Name         string                                              `json:"name"`
	Project      string                                              `json:"project"`
	Tenant       string                                              `json:"tenant"`
	Timeouts     *identityplatformtenantoauthidpconfig.TimeoutsState `json:"timeouts"`
}
