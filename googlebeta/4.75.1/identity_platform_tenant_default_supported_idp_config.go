// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	identityplatformtenantdefaultsupportedidpconfig "github.com/golingon/terraproviders/googlebeta/4.75.1/identityplatformtenantdefaultsupportedidpconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityPlatformTenantDefaultSupportedIdpConfig creates a new instance of [IdentityPlatformTenantDefaultSupportedIdpConfig].
func NewIdentityPlatformTenantDefaultSupportedIdpConfig(name string, args IdentityPlatformTenantDefaultSupportedIdpConfigArgs) *IdentityPlatformTenantDefaultSupportedIdpConfig {
	return &IdentityPlatformTenantDefaultSupportedIdpConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityPlatformTenantDefaultSupportedIdpConfig)(nil)

// IdentityPlatformTenantDefaultSupportedIdpConfig represents the Terraform resource google_identity_platform_tenant_default_supported_idp_config.
type IdentityPlatformTenantDefaultSupportedIdpConfig struct {
	Name      string
	Args      IdentityPlatformTenantDefaultSupportedIdpConfigArgs
	state     *identityPlatformTenantDefaultSupportedIdpConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityPlatformTenantDefaultSupportedIdpConfig].
func (iptdsic *IdentityPlatformTenantDefaultSupportedIdpConfig) Type() string {
	return "google_identity_platform_tenant_default_supported_idp_config"
}

// LocalName returns the local name for [IdentityPlatformTenantDefaultSupportedIdpConfig].
func (iptdsic *IdentityPlatformTenantDefaultSupportedIdpConfig) LocalName() string {
	return iptdsic.Name
}

// Configuration returns the configuration (args) for [IdentityPlatformTenantDefaultSupportedIdpConfig].
func (iptdsic *IdentityPlatformTenantDefaultSupportedIdpConfig) Configuration() interface{} {
	return iptdsic.Args
}

// DependOn is used for other resources to depend on [IdentityPlatformTenantDefaultSupportedIdpConfig].
func (iptdsic *IdentityPlatformTenantDefaultSupportedIdpConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(iptdsic)
}

// Dependencies returns the list of resources [IdentityPlatformTenantDefaultSupportedIdpConfig] depends_on.
func (iptdsic *IdentityPlatformTenantDefaultSupportedIdpConfig) Dependencies() terra.Dependencies {
	return iptdsic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityPlatformTenantDefaultSupportedIdpConfig].
func (iptdsic *IdentityPlatformTenantDefaultSupportedIdpConfig) LifecycleManagement() *terra.Lifecycle {
	return iptdsic.Lifecycle
}

// Attributes returns the attributes for [IdentityPlatformTenantDefaultSupportedIdpConfig].
func (iptdsic *IdentityPlatformTenantDefaultSupportedIdpConfig) Attributes() identityPlatformTenantDefaultSupportedIdpConfigAttributes {
	return identityPlatformTenantDefaultSupportedIdpConfigAttributes{ref: terra.ReferenceResource(iptdsic)}
}

// ImportState imports the given attribute values into [IdentityPlatformTenantDefaultSupportedIdpConfig]'s state.
func (iptdsic *IdentityPlatformTenantDefaultSupportedIdpConfig) ImportState(av io.Reader) error {
	iptdsic.state = &identityPlatformTenantDefaultSupportedIdpConfigState{}
	if err := json.NewDecoder(av).Decode(iptdsic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iptdsic.Type(), iptdsic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityPlatformTenantDefaultSupportedIdpConfig] has state.
func (iptdsic *IdentityPlatformTenantDefaultSupportedIdpConfig) State() (*identityPlatformTenantDefaultSupportedIdpConfigState, bool) {
	return iptdsic.state, iptdsic.state != nil
}

// StateMust returns the state for [IdentityPlatformTenantDefaultSupportedIdpConfig]. Panics if the state is nil.
func (iptdsic *IdentityPlatformTenantDefaultSupportedIdpConfig) StateMust() *identityPlatformTenantDefaultSupportedIdpConfigState {
	if iptdsic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iptdsic.Type(), iptdsic.LocalName()))
	}
	return iptdsic.state
}

// IdentityPlatformTenantDefaultSupportedIdpConfigArgs contains the configurations for google_identity_platform_tenant_default_supported_idp_config.
type IdentityPlatformTenantDefaultSupportedIdpConfigArgs struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdpId: string, required
	IdpId terra.StringValue `hcl:"idp_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Tenant: string, required
	Tenant terra.StringValue `hcl:"tenant,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *identityplatformtenantdefaultsupportedidpconfig.Timeouts `hcl:"timeouts,block"`
}
type identityPlatformTenantDefaultSupportedIdpConfigAttributes struct {
	ref terra.Reference
}

// ClientId returns a reference to field client_id of google_identity_platform_tenant_default_supported_idp_config.
func (iptdsic identityPlatformTenantDefaultSupportedIdpConfigAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(iptdsic.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of google_identity_platform_tenant_default_supported_idp_config.
func (iptdsic identityPlatformTenantDefaultSupportedIdpConfigAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(iptdsic.ref.Append("client_secret"))
}

// Enabled returns a reference to field enabled of google_identity_platform_tenant_default_supported_idp_config.
func (iptdsic identityPlatformTenantDefaultSupportedIdpConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(iptdsic.ref.Append("enabled"))
}

// Id returns a reference to field id of google_identity_platform_tenant_default_supported_idp_config.
func (iptdsic identityPlatformTenantDefaultSupportedIdpConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iptdsic.ref.Append("id"))
}

// IdpId returns a reference to field idp_id of google_identity_platform_tenant_default_supported_idp_config.
func (iptdsic identityPlatformTenantDefaultSupportedIdpConfigAttributes) IdpId() terra.StringValue {
	return terra.ReferenceAsString(iptdsic.ref.Append("idp_id"))
}

// Name returns a reference to field name of google_identity_platform_tenant_default_supported_idp_config.
func (iptdsic identityPlatformTenantDefaultSupportedIdpConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iptdsic.ref.Append("name"))
}

// Project returns a reference to field project of google_identity_platform_tenant_default_supported_idp_config.
func (iptdsic identityPlatformTenantDefaultSupportedIdpConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iptdsic.ref.Append("project"))
}

// Tenant returns a reference to field tenant of google_identity_platform_tenant_default_supported_idp_config.
func (iptdsic identityPlatformTenantDefaultSupportedIdpConfigAttributes) Tenant() terra.StringValue {
	return terra.ReferenceAsString(iptdsic.ref.Append("tenant"))
}

func (iptdsic identityPlatformTenantDefaultSupportedIdpConfigAttributes) Timeouts() identityplatformtenantdefaultsupportedidpconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[identityplatformtenantdefaultsupportedidpconfig.TimeoutsAttributes](iptdsic.ref.Append("timeouts"))
}

type identityPlatformTenantDefaultSupportedIdpConfigState struct {
	ClientId     string                                                         `json:"client_id"`
	ClientSecret string                                                         `json:"client_secret"`
	Enabled      bool                                                           `json:"enabled"`
	Id           string                                                         `json:"id"`
	IdpId        string                                                         `json:"idp_id"`
	Name         string                                                         `json:"name"`
	Project      string                                                         `json:"project"`
	Tenant       string                                                         `json:"tenant"`
	Timeouts     *identityplatformtenantdefaultsupportedidpconfig.TimeoutsState `json:"timeouts"`
}
