// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cdn_frontdoor_secret

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

// Resource represents the Terraform resource azurerm_cdn_frontdoor_secret.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermCdnFrontdoorSecretState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acfs *Resource) Type() string {
	return "azurerm_cdn_frontdoor_secret"
}

// LocalName returns the local name for [Resource].
func (acfs *Resource) LocalName() string {
	return acfs.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acfs *Resource) Configuration() interface{} {
	return acfs.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acfs *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acfs)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acfs *Resource) Dependencies() terra.Dependencies {
	return acfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acfs *Resource) LifecycleManagement() *terra.Lifecycle {
	return acfs.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acfs *Resource) Attributes() azurermCdnFrontdoorSecretAttributes {
	return azurermCdnFrontdoorSecretAttributes{ref: terra.ReferenceResource(acfs)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acfs *Resource) ImportState(state io.Reader) error {
	acfs.state = &azurermCdnFrontdoorSecretState{}
	if err := json.NewDecoder(state).Decode(acfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acfs.Type(), acfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acfs *Resource) State() (*azurermCdnFrontdoorSecretState, bool) {
	return acfs.state, acfs.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acfs *Resource) StateMust() *azurermCdnFrontdoorSecretState {
	if acfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acfs.Type(), acfs.LocalName()))
	}
	return acfs.state
}

// Args contains the configurations for azurerm_cdn_frontdoor_secret.
type Args struct {
	// CdnFrontdoorProfileId: string, required
	CdnFrontdoorProfileId terra.StringValue `hcl:"cdn_frontdoor_profile_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Secret: required
	Secret *Secret `hcl:"secret,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermCdnFrontdoorSecretAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_secret.
func (acfs azurermCdnFrontdoorSecretAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(acfs.ref.Append("cdn_frontdoor_profile_id"))
}

// CdnFrontdoorProfileName returns a reference to field cdn_frontdoor_profile_name of azurerm_cdn_frontdoor_secret.
func (acfs azurermCdnFrontdoorSecretAttributes) CdnFrontdoorProfileName() terra.StringValue {
	return terra.ReferenceAsString(acfs.ref.Append("cdn_frontdoor_profile_name"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_secret.
func (acfs azurermCdnFrontdoorSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acfs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_secret.
func (acfs azurermCdnFrontdoorSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acfs.ref.Append("name"))
}

func (acfs azurermCdnFrontdoorSecretAttributes) Secret() terra.ListValue[SecretAttributes] {
	return terra.ReferenceAsList[SecretAttributes](acfs.ref.Append("secret"))
}

func (acfs azurermCdnFrontdoorSecretAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acfs.ref.Append("timeouts"))
}

type azurermCdnFrontdoorSecretState struct {
	CdnFrontdoorProfileId   string         `json:"cdn_frontdoor_profile_id"`
	CdnFrontdoorProfileName string         `json:"cdn_frontdoor_profile_name"`
	Id                      string         `json:"id"`
	Name                    string         `json:"name"`
	Secret                  []SecretState  `json:"secret"`
	Timeouts                *TimeoutsState `json:"timeouts"`
}
