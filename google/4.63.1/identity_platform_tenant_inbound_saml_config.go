// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	identityplatformtenantinboundsamlconfig "github.com/golingon/terraproviders/google/4.63.1/identityplatformtenantinboundsamlconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityPlatformTenantInboundSamlConfig creates a new instance of [IdentityPlatformTenantInboundSamlConfig].
func NewIdentityPlatformTenantInboundSamlConfig(name string, args IdentityPlatformTenantInboundSamlConfigArgs) *IdentityPlatformTenantInboundSamlConfig {
	return &IdentityPlatformTenantInboundSamlConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityPlatformTenantInboundSamlConfig)(nil)

// IdentityPlatformTenantInboundSamlConfig represents the Terraform resource google_identity_platform_tenant_inbound_saml_config.
type IdentityPlatformTenantInboundSamlConfig struct {
	Name      string
	Args      IdentityPlatformTenantInboundSamlConfigArgs
	state     *identityPlatformTenantInboundSamlConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityPlatformTenantInboundSamlConfig].
func (iptisc *IdentityPlatformTenantInboundSamlConfig) Type() string {
	return "google_identity_platform_tenant_inbound_saml_config"
}

// LocalName returns the local name for [IdentityPlatformTenantInboundSamlConfig].
func (iptisc *IdentityPlatformTenantInboundSamlConfig) LocalName() string {
	return iptisc.Name
}

// Configuration returns the configuration (args) for [IdentityPlatformTenantInboundSamlConfig].
func (iptisc *IdentityPlatformTenantInboundSamlConfig) Configuration() interface{} {
	return iptisc.Args
}

// DependOn is used for other resources to depend on [IdentityPlatformTenantInboundSamlConfig].
func (iptisc *IdentityPlatformTenantInboundSamlConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(iptisc)
}

// Dependencies returns the list of resources [IdentityPlatformTenantInboundSamlConfig] depends_on.
func (iptisc *IdentityPlatformTenantInboundSamlConfig) Dependencies() terra.Dependencies {
	return iptisc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityPlatformTenantInboundSamlConfig].
func (iptisc *IdentityPlatformTenantInboundSamlConfig) LifecycleManagement() *terra.Lifecycle {
	return iptisc.Lifecycle
}

// Attributes returns the attributes for [IdentityPlatformTenantInboundSamlConfig].
func (iptisc *IdentityPlatformTenantInboundSamlConfig) Attributes() identityPlatformTenantInboundSamlConfigAttributes {
	return identityPlatformTenantInboundSamlConfigAttributes{ref: terra.ReferenceResource(iptisc)}
}

// ImportState imports the given attribute values into [IdentityPlatformTenantInboundSamlConfig]'s state.
func (iptisc *IdentityPlatformTenantInboundSamlConfig) ImportState(av io.Reader) error {
	iptisc.state = &identityPlatformTenantInboundSamlConfigState{}
	if err := json.NewDecoder(av).Decode(iptisc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iptisc.Type(), iptisc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityPlatformTenantInboundSamlConfig] has state.
func (iptisc *IdentityPlatformTenantInboundSamlConfig) State() (*identityPlatformTenantInboundSamlConfigState, bool) {
	return iptisc.state, iptisc.state != nil
}

// StateMust returns the state for [IdentityPlatformTenantInboundSamlConfig]. Panics if the state is nil.
func (iptisc *IdentityPlatformTenantInboundSamlConfig) StateMust() *identityPlatformTenantInboundSamlConfigState {
	if iptisc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iptisc.Type(), iptisc.LocalName()))
	}
	return iptisc.state
}

// IdentityPlatformTenantInboundSamlConfigArgs contains the configurations for google_identity_platform_tenant_inbound_saml_config.
type IdentityPlatformTenantInboundSamlConfigArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Tenant: string, required
	Tenant terra.StringValue `hcl:"tenant,attr" validate:"required"`
	// IdpConfig: required
	IdpConfig *identityplatformtenantinboundsamlconfig.IdpConfig `hcl:"idp_config,block" validate:"required"`
	// SpConfig: required
	SpConfig *identityplatformtenantinboundsamlconfig.SpConfig `hcl:"sp_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *identityplatformtenantinboundsamlconfig.Timeouts `hcl:"timeouts,block"`
}
type identityPlatformTenantInboundSamlConfigAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_identity_platform_tenant_inbound_saml_config.
func (iptisc identityPlatformTenantInboundSamlConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(iptisc.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of google_identity_platform_tenant_inbound_saml_config.
func (iptisc identityPlatformTenantInboundSamlConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(iptisc.ref.Append("enabled"))
}

// Id returns a reference to field id of google_identity_platform_tenant_inbound_saml_config.
func (iptisc identityPlatformTenantInboundSamlConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iptisc.ref.Append("id"))
}

// Name returns a reference to field name of google_identity_platform_tenant_inbound_saml_config.
func (iptisc identityPlatformTenantInboundSamlConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iptisc.ref.Append("name"))
}

// Project returns a reference to field project of google_identity_platform_tenant_inbound_saml_config.
func (iptisc identityPlatformTenantInboundSamlConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iptisc.ref.Append("project"))
}

// Tenant returns a reference to field tenant of google_identity_platform_tenant_inbound_saml_config.
func (iptisc identityPlatformTenantInboundSamlConfigAttributes) Tenant() terra.StringValue {
	return terra.ReferenceAsString(iptisc.ref.Append("tenant"))
}

func (iptisc identityPlatformTenantInboundSamlConfigAttributes) IdpConfig() terra.ListValue[identityplatformtenantinboundsamlconfig.IdpConfigAttributes] {
	return terra.ReferenceAsList[identityplatformtenantinboundsamlconfig.IdpConfigAttributes](iptisc.ref.Append("idp_config"))
}

func (iptisc identityPlatformTenantInboundSamlConfigAttributes) SpConfig() terra.ListValue[identityplatformtenantinboundsamlconfig.SpConfigAttributes] {
	return terra.ReferenceAsList[identityplatformtenantinboundsamlconfig.SpConfigAttributes](iptisc.ref.Append("sp_config"))
}

func (iptisc identityPlatformTenantInboundSamlConfigAttributes) Timeouts() identityplatformtenantinboundsamlconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[identityplatformtenantinboundsamlconfig.TimeoutsAttributes](iptisc.ref.Append("timeouts"))
}

type identityPlatformTenantInboundSamlConfigState struct {
	DisplayName string                                                   `json:"display_name"`
	Enabled     bool                                                     `json:"enabled"`
	Id          string                                                   `json:"id"`
	Name        string                                                   `json:"name"`
	Project     string                                                   `json:"project"`
	Tenant      string                                                   `json:"tenant"`
	IdpConfig   []identityplatformtenantinboundsamlconfig.IdpConfigState `json:"idp_config"`
	SpConfig    []identityplatformtenantinboundsamlconfig.SpConfigState  `json:"sp_config"`
	Timeouts    *identityplatformtenantinboundsamlconfig.TimeoutsState   `json:"timeouts"`
}
