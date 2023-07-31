// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	identityplatformdefaultsupportedidpconfig "github.com/golingon/terraproviders/googlebeta/4.75.1/identityplatformdefaultsupportedidpconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityPlatformDefaultSupportedIdpConfig creates a new instance of [IdentityPlatformDefaultSupportedIdpConfig].
func NewIdentityPlatformDefaultSupportedIdpConfig(name string, args IdentityPlatformDefaultSupportedIdpConfigArgs) *IdentityPlatformDefaultSupportedIdpConfig {
	return &IdentityPlatformDefaultSupportedIdpConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityPlatformDefaultSupportedIdpConfig)(nil)

// IdentityPlatformDefaultSupportedIdpConfig represents the Terraform resource google_identity_platform_default_supported_idp_config.
type IdentityPlatformDefaultSupportedIdpConfig struct {
	Name      string
	Args      IdentityPlatformDefaultSupportedIdpConfigArgs
	state     *identityPlatformDefaultSupportedIdpConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityPlatformDefaultSupportedIdpConfig].
func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) Type() string {
	return "google_identity_platform_default_supported_idp_config"
}

// LocalName returns the local name for [IdentityPlatformDefaultSupportedIdpConfig].
func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) LocalName() string {
	return ipdsic.Name
}

// Configuration returns the configuration (args) for [IdentityPlatformDefaultSupportedIdpConfig].
func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) Configuration() interface{} {
	return ipdsic.Args
}

// DependOn is used for other resources to depend on [IdentityPlatformDefaultSupportedIdpConfig].
func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(ipdsic)
}

// Dependencies returns the list of resources [IdentityPlatformDefaultSupportedIdpConfig] depends_on.
func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) Dependencies() terra.Dependencies {
	return ipdsic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityPlatformDefaultSupportedIdpConfig].
func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) LifecycleManagement() *terra.Lifecycle {
	return ipdsic.Lifecycle
}

// Attributes returns the attributes for [IdentityPlatformDefaultSupportedIdpConfig].
func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) Attributes() identityPlatformDefaultSupportedIdpConfigAttributes {
	return identityPlatformDefaultSupportedIdpConfigAttributes{ref: terra.ReferenceResource(ipdsic)}
}

// ImportState imports the given attribute values into [IdentityPlatformDefaultSupportedIdpConfig]'s state.
func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) ImportState(av io.Reader) error {
	ipdsic.state = &identityPlatformDefaultSupportedIdpConfigState{}
	if err := json.NewDecoder(av).Decode(ipdsic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ipdsic.Type(), ipdsic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityPlatformDefaultSupportedIdpConfig] has state.
func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) State() (*identityPlatformDefaultSupportedIdpConfigState, bool) {
	return ipdsic.state, ipdsic.state != nil
}

// StateMust returns the state for [IdentityPlatformDefaultSupportedIdpConfig]. Panics if the state is nil.
func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) StateMust() *identityPlatformDefaultSupportedIdpConfigState {
	if ipdsic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ipdsic.Type(), ipdsic.LocalName()))
	}
	return ipdsic.state
}

// IdentityPlatformDefaultSupportedIdpConfigArgs contains the configurations for google_identity_platform_default_supported_idp_config.
type IdentityPlatformDefaultSupportedIdpConfigArgs struct {
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
	// Timeouts: optional
	Timeouts *identityplatformdefaultsupportedidpconfig.Timeouts `hcl:"timeouts,block"`
}
type identityPlatformDefaultSupportedIdpConfigAttributes struct {
	ref terra.Reference
}

// ClientId returns a reference to field client_id of google_identity_platform_default_supported_idp_config.
func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(ipdsic.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of google_identity_platform_default_supported_idp_config.
func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(ipdsic.ref.Append("client_secret"))
}

// Enabled returns a reference to field enabled of google_identity_platform_default_supported_idp_config.
func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ipdsic.ref.Append("enabled"))
}

// Id returns a reference to field id of google_identity_platform_default_supported_idp_config.
func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ipdsic.ref.Append("id"))
}

// IdpId returns a reference to field idp_id of google_identity_platform_default_supported_idp_config.
func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) IdpId() terra.StringValue {
	return terra.ReferenceAsString(ipdsic.ref.Append("idp_id"))
}

// Name returns a reference to field name of google_identity_platform_default_supported_idp_config.
func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ipdsic.ref.Append("name"))
}

// Project returns a reference to field project of google_identity_platform_default_supported_idp_config.
func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ipdsic.ref.Append("project"))
}

func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) Timeouts() identityplatformdefaultsupportedidpconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[identityplatformdefaultsupportedidpconfig.TimeoutsAttributes](ipdsic.ref.Append("timeouts"))
}

type identityPlatformDefaultSupportedIdpConfigState struct {
	ClientId     string                                                   `json:"client_id"`
	ClientSecret string                                                   `json:"client_secret"`
	Enabled      bool                                                     `json:"enabled"`
	Id           string                                                   `json:"id"`
	IdpId        string                                                   `json:"idp_id"`
	Name         string                                                   `json:"name"`
	Project      string                                                   `json:"project"`
	Timeouts     *identityplatformdefaultsupportedidpconfig.TimeoutsState `json:"timeouts"`
}
