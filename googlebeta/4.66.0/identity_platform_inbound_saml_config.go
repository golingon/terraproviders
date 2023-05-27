// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	identityplatforminboundsamlconfig "github.com/golingon/terraproviders/googlebeta/4.66.0/identityplatforminboundsamlconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityPlatformInboundSamlConfig creates a new instance of [IdentityPlatformInboundSamlConfig].
func NewIdentityPlatformInboundSamlConfig(name string, args IdentityPlatformInboundSamlConfigArgs) *IdentityPlatformInboundSamlConfig {
	return &IdentityPlatformInboundSamlConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityPlatformInboundSamlConfig)(nil)

// IdentityPlatformInboundSamlConfig represents the Terraform resource google_identity_platform_inbound_saml_config.
type IdentityPlatformInboundSamlConfig struct {
	Name      string
	Args      IdentityPlatformInboundSamlConfigArgs
	state     *identityPlatformInboundSamlConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityPlatformInboundSamlConfig].
func (ipisc *IdentityPlatformInboundSamlConfig) Type() string {
	return "google_identity_platform_inbound_saml_config"
}

// LocalName returns the local name for [IdentityPlatformInboundSamlConfig].
func (ipisc *IdentityPlatformInboundSamlConfig) LocalName() string {
	return ipisc.Name
}

// Configuration returns the configuration (args) for [IdentityPlatformInboundSamlConfig].
func (ipisc *IdentityPlatformInboundSamlConfig) Configuration() interface{} {
	return ipisc.Args
}

// DependOn is used for other resources to depend on [IdentityPlatformInboundSamlConfig].
func (ipisc *IdentityPlatformInboundSamlConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(ipisc)
}

// Dependencies returns the list of resources [IdentityPlatformInboundSamlConfig] depends_on.
func (ipisc *IdentityPlatformInboundSamlConfig) Dependencies() terra.Dependencies {
	return ipisc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityPlatformInboundSamlConfig].
func (ipisc *IdentityPlatformInboundSamlConfig) LifecycleManagement() *terra.Lifecycle {
	return ipisc.Lifecycle
}

// Attributes returns the attributes for [IdentityPlatformInboundSamlConfig].
func (ipisc *IdentityPlatformInboundSamlConfig) Attributes() identityPlatformInboundSamlConfigAttributes {
	return identityPlatformInboundSamlConfigAttributes{ref: terra.ReferenceResource(ipisc)}
}

// ImportState imports the given attribute values into [IdentityPlatformInboundSamlConfig]'s state.
func (ipisc *IdentityPlatformInboundSamlConfig) ImportState(av io.Reader) error {
	ipisc.state = &identityPlatformInboundSamlConfigState{}
	if err := json.NewDecoder(av).Decode(ipisc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ipisc.Type(), ipisc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityPlatformInboundSamlConfig] has state.
func (ipisc *IdentityPlatformInboundSamlConfig) State() (*identityPlatformInboundSamlConfigState, bool) {
	return ipisc.state, ipisc.state != nil
}

// StateMust returns the state for [IdentityPlatformInboundSamlConfig]. Panics if the state is nil.
func (ipisc *IdentityPlatformInboundSamlConfig) StateMust() *identityPlatformInboundSamlConfigState {
	if ipisc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ipisc.Type(), ipisc.LocalName()))
	}
	return ipisc.state
}

// IdentityPlatformInboundSamlConfigArgs contains the configurations for google_identity_platform_inbound_saml_config.
type IdentityPlatformInboundSamlConfigArgs struct {
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
	// IdpConfig: required
	IdpConfig *identityplatforminboundsamlconfig.IdpConfig `hcl:"idp_config,block" validate:"required"`
	// SpConfig: required
	SpConfig *identityplatforminboundsamlconfig.SpConfig `hcl:"sp_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *identityplatforminboundsamlconfig.Timeouts `hcl:"timeouts,block"`
}
type identityPlatformInboundSamlConfigAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_identity_platform_inbound_saml_config.
func (ipisc identityPlatformInboundSamlConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ipisc.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of google_identity_platform_inbound_saml_config.
func (ipisc identityPlatformInboundSamlConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ipisc.ref.Append("enabled"))
}

// Id returns a reference to field id of google_identity_platform_inbound_saml_config.
func (ipisc identityPlatformInboundSamlConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ipisc.ref.Append("id"))
}

// Name returns a reference to field name of google_identity_platform_inbound_saml_config.
func (ipisc identityPlatformInboundSamlConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ipisc.ref.Append("name"))
}

// Project returns a reference to field project of google_identity_platform_inbound_saml_config.
func (ipisc identityPlatformInboundSamlConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ipisc.ref.Append("project"))
}

func (ipisc identityPlatformInboundSamlConfigAttributes) IdpConfig() terra.ListValue[identityplatforminboundsamlconfig.IdpConfigAttributes] {
	return terra.ReferenceAsList[identityplatforminboundsamlconfig.IdpConfigAttributes](ipisc.ref.Append("idp_config"))
}

func (ipisc identityPlatformInboundSamlConfigAttributes) SpConfig() terra.ListValue[identityplatforminboundsamlconfig.SpConfigAttributes] {
	return terra.ReferenceAsList[identityplatforminboundsamlconfig.SpConfigAttributes](ipisc.ref.Append("sp_config"))
}

func (ipisc identityPlatformInboundSamlConfigAttributes) Timeouts() identityplatforminboundsamlconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[identityplatforminboundsamlconfig.TimeoutsAttributes](ipisc.ref.Append("timeouts"))
}

type identityPlatformInboundSamlConfigState struct {
	DisplayName string                                             `json:"display_name"`
	Enabled     bool                                               `json:"enabled"`
	Id          string                                             `json:"id"`
	Name        string                                             `json:"name"`
	Project     string                                             `json:"project"`
	IdpConfig   []identityplatforminboundsamlconfig.IdpConfigState `json:"idp_config"`
	SpConfig    []identityplatforminboundsamlconfig.SpConfigState  `json:"sp_config"`
	Timeouts    *identityplatforminboundsamlconfig.TimeoutsState   `json:"timeouts"`
}
