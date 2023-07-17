// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	identityplatformoauthidpconfig "github.com/golingon/terraproviders/google/4.73.1/identityplatformoauthidpconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityPlatformOauthIdpConfig creates a new instance of [IdentityPlatformOauthIdpConfig].
func NewIdentityPlatformOauthIdpConfig(name string, args IdentityPlatformOauthIdpConfigArgs) *IdentityPlatformOauthIdpConfig {
	return &IdentityPlatformOauthIdpConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityPlatformOauthIdpConfig)(nil)

// IdentityPlatformOauthIdpConfig represents the Terraform resource google_identity_platform_oauth_idp_config.
type IdentityPlatformOauthIdpConfig struct {
	Name      string
	Args      IdentityPlatformOauthIdpConfigArgs
	state     *identityPlatformOauthIdpConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityPlatformOauthIdpConfig].
func (ipoic *IdentityPlatformOauthIdpConfig) Type() string {
	return "google_identity_platform_oauth_idp_config"
}

// LocalName returns the local name for [IdentityPlatformOauthIdpConfig].
func (ipoic *IdentityPlatformOauthIdpConfig) LocalName() string {
	return ipoic.Name
}

// Configuration returns the configuration (args) for [IdentityPlatformOauthIdpConfig].
func (ipoic *IdentityPlatformOauthIdpConfig) Configuration() interface{} {
	return ipoic.Args
}

// DependOn is used for other resources to depend on [IdentityPlatformOauthIdpConfig].
func (ipoic *IdentityPlatformOauthIdpConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(ipoic)
}

// Dependencies returns the list of resources [IdentityPlatformOauthIdpConfig] depends_on.
func (ipoic *IdentityPlatformOauthIdpConfig) Dependencies() terra.Dependencies {
	return ipoic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityPlatformOauthIdpConfig].
func (ipoic *IdentityPlatformOauthIdpConfig) LifecycleManagement() *terra.Lifecycle {
	return ipoic.Lifecycle
}

// Attributes returns the attributes for [IdentityPlatformOauthIdpConfig].
func (ipoic *IdentityPlatformOauthIdpConfig) Attributes() identityPlatformOauthIdpConfigAttributes {
	return identityPlatformOauthIdpConfigAttributes{ref: terra.ReferenceResource(ipoic)}
}

// ImportState imports the given attribute values into [IdentityPlatformOauthIdpConfig]'s state.
func (ipoic *IdentityPlatformOauthIdpConfig) ImportState(av io.Reader) error {
	ipoic.state = &identityPlatformOauthIdpConfigState{}
	if err := json.NewDecoder(av).Decode(ipoic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ipoic.Type(), ipoic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityPlatformOauthIdpConfig] has state.
func (ipoic *IdentityPlatformOauthIdpConfig) State() (*identityPlatformOauthIdpConfigState, bool) {
	return ipoic.state, ipoic.state != nil
}

// StateMust returns the state for [IdentityPlatformOauthIdpConfig]. Panics if the state is nil.
func (ipoic *IdentityPlatformOauthIdpConfig) StateMust() *identityPlatformOauthIdpConfigState {
	if ipoic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ipoic.Type(), ipoic.LocalName()))
	}
	return ipoic.state
}

// IdentityPlatformOauthIdpConfigArgs contains the configurations for google_identity_platform_oauth_idp_config.
type IdentityPlatformOauthIdpConfigArgs struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, optional
	ClientSecret terra.StringValue `hcl:"client_secret,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
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
	// Timeouts: optional
	Timeouts *identityplatformoauthidpconfig.Timeouts `hcl:"timeouts,block"`
}
type identityPlatformOauthIdpConfigAttributes struct {
	ref terra.Reference
}

// ClientId returns a reference to field client_id of google_identity_platform_oauth_idp_config.
func (ipoic identityPlatformOauthIdpConfigAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(ipoic.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of google_identity_platform_oauth_idp_config.
func (ipoic identityPlatformOauthIdpConfigAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(ipoic.ref.Append("client_secret"))
}

// DisplayName returns a reference to field display_name of google_identity_platform_oauth_idp_config.
func (ipoic identityPlatformOauthIdpConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ipoic.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of google_identity_platform_oauth_idp_config.
func (ipoic identityPlatformOauthIdpConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ipoic.ref.Append("enabled"))
}

// Id returns a reference to field id of google_identity_platform_oauth_idp_config.
func (ipoic identityPlatformOauthIdpConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ipoic.ref.Append("id"))
}

// Issuer returns a reference to field issuer of google_identity_platform_oauth_idp_config.
func (ipoic identityPlatformOauthIdpConfigAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ipoic.ref.Append("issuer"))
}

// Name returns a reference to field name of google_identity_platform_oauth_idp_config.
func (ipoic identityPlatformOauthIdpConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ipoic.ref.Append("name"))
}

// Project returns a reference to field project of google_identity_platform_oauth_idp_config.
func (ipoic identityPlatformOauthIdpConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ipoic.ref.Append("project"))
}

func (ipoic identityPlatformOauthIdpConfigAttributes) Timeouts() identityplatformoauthidpconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[identityplatformoauthidpconfig.TimeoutsAttributes](ipoic.ref.Append("timeouts"))
}

type identityPlatformOauthIdpConfigState struct {
	ClientId     string                                        `json:"client_id"`
	ClientSecret string                                        `json:"client_secret"`
	DisplayName  string                                        `json:"display_name"`
	Enabled      bool                                          `json:"enabled"`
	Id           string                                        `json:"id"`
	Issuer       string                                        `json:"issuer"`
	Name         string                                        `json:"name"`
	Project      string                                        `json:"project"`
	Timeouts     *identityplatformoauthidpconfig.TimeoutsState `json:"timeouts"`
}
