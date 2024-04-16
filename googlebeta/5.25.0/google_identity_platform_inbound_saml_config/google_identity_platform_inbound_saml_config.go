// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_identity_platform_inbound_saml_config

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_identity_platform_inbound_saml_config.
type Resource struct {
	Name      string
	Args      Args
	state     *googleIdentityPlatformInboundSamlConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gipisc *Resource) Type() string {
	return "google_identity_platform_inbound_saml_config"
}

// LocalName returns the local name for [Resource].
func (gipisc *Resource) LocalName() string {
	return gipisc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gipisc *Resource) Configuration() interface{} {
	return gipisc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gipisc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gipisc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gipisc *Resource) Dependencies() terra.Dependencies {
	return gipisc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gipisc *Resource) LifecycleManagement() *terra.Lifecycle {
	return gipisc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gipisc *Resource) Attributes() googleIdentityPlatformInboundSamlConfigAttributes {
	return googleIdentityPlatformInboundSamlConfigAttributes{ref: terra.ReferenceResource(gipisc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gipisc *Resource) ImportState(state io.Reader) error {
	gipisc.state = &googleIdentityPlatformInboundSamlConfigState{}
	if err := json.NewDecoder(state).Decode(gipisc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gipisc.Type(), gipisc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gipisc *Resource) State() (*googleIdentityPlatformInboundSamlConfigState, bool) {
	return gipisc.state, gipisc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gipisc *Resource) StateMust() *googleIdentityPlatformInboundSamlConfigState {
	if gipisc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gipisc.Type(), gipisc.LocalName()))
	}
	return gipisc.state
}

// Args contains the configurations for google_identity_platform_inbound_saml_config.
type Args struct {
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
	IdpConfig *IdpConfig `hcl:"idp_config,block" validate:"required"`
	// SpConfig: required
	SpConfig *SpConfig `hcl:"sp_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleIdentityPlatformInboundSamlConfigAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_identity_platform_inbound_saml_config.
func (gipisc googleIdentityPlatformInboundSamlConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gipisc.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of google_identity_platform_inbound_saml_config.
func (gipisc googleIdentityPlatformInboundSamlConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(gipisc.ref.Append("enabled"))
}

// Id returns a reference to field id of google_identity_platform_inbound_saml_config.
func (gipisc googleIdentityPlatformInboundSamlConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gipisc.ref.Append("id"))
}

// Name returns a reference to field name of google_identity_platform_inbound_saml_config.
func (gipisc googleIdentityPlatformInboundSamlConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gipisc.ref.Append("name"))
}

// Project returns a reference to field project of google_identity_platform_inbound_saml_config.
func (gipisc googleIdentityPlatformInboundSamlConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gipisc.ref.Append("project"))
}

func (gipisc googleIdentityPlatformInboundSamlConfigAttributes) IdpConfig() terra.ListValue[IdpConfigAttributes] {
	return terra.ReferenceAsList[IdpConfigAttributes](gipisc.ref.Append("idp_config"))
}

func (gipisc googleIdentityPlatformInboundSamlConfigAttributes) SpConfig() terra.ListValue[SpConfigAttributes] {
	return terra.ReferenceAsList[SpConfigAttributes](gipisc.ref.Append("sp_config"))
}

func (gipisc googleIdentityPlatformInboundSamlConfigAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gipisc.ref.Append("timeouts"))
}

type googleIdentityPlatformInboundSamlConfigState struct {
	DisplayName string           `json:"display_name"`
	Enabled     bool             `json:"enabled"`
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	Project     string           `json:"project"`
	IdpConfig   []IdpConfigState `json:"idp_config"`
	SpConfig    []SpConfigState  `json:"sp_config"`
	Timeouts    *TimeoutsState   `json:"timeouts"`
}
