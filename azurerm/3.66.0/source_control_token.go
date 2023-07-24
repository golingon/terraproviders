// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sourcecontroltoken "github.com/golingon/terraproviders/azurerm/3.66.0/sourcecontroltoken"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSourceControlToken creates a new instance of [SourceControlToken].
func NewSourceControlToken(name string, args SourceControlTokenArgs) *SourceControlToken {
	return &SourceControlToken{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SourceControlToken)(nil)

// SourceControlToken represents the Terraform resource azurerm_source_control_token.
type SourceControlToken struct {
	Name      string
	Args      SourceControlTokenArgs
	state     *sourceControlTokenState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SourceControlToken].
func (sct *SourceControlToken) Type() string {
	return "azurerm_source_control_token"
}

// LocalName returns the local name for [SourceControlToken].
func (sct *SourceControlToken) LocalName() string {
	return sct.Name
}

// Configuration returns the configuration (args) for [SourceControlToken].
func (sct *SourceControlToken) Configuration() interface{} {
	return sct.Args
}

// DependOn is used for other resources to depend on [SourceControlToken].
func (sct *SourceControlToken) DependOn() terra.Reference {
	return terra.ReferenceResource(sct)
}

// Dependencies returns the list of resources [SourceControlToken] depends_on.
func (sct *SourceControlToken) Dependencies() terra.Dependencies {
	return sct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SourceControlToken].
func (sct *SourceControlToken) LifecycleManagement() *terra.Lifecycle {
	return sct.Lifecycle
}

// Attributes returns the attributes for [SourceControlToken].
func (sct *SourceControlToken) Attributes() sourceControlTokenAttributes {
	return sourceControlTokenAttributes{ref: terra.ReferenceResource(sct)}
}

// ImportState imports the given attribute values into [SourceControlToken]'s state.
func (sct *SourceControlToken) ImportState(av io.Reader) error {
	sct.state = &sourceControlTokenState{}
	if err := json.NewDecoder(av).Decode(sct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sct.Type(), sct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SourceControlToken] has state.
func (sct *SourceControlToken) State() (*sourceControlTokenState, bool) {
	return sct.state, sct.state != nil
}

// StateMust returns the state for [SourceControlToken]. Panics if the state is nil.
func (sct *SourceControlToken) StateMust() *sourceControlTokenState {
	if sct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sct.Type(), sct.LocalName()))
	}
	return sct.state
}

// SourceControlTokenArgs contains the configurations for azurerm_source_control_token.
type SourceControlTokenArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Token: string, required
	Token terra.StringValue `hcl:"token,attr" validate:"required"`
	// TokenSecret: string, optional
	TokenSecret terra.StringValue `hcl:"token_secret,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sourcecontroltoken.Timeouts `hcl:"timeouts,block"`
}
type sourceControlTokenAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_source_control_token.
func (sct sourceControlTokenAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sct.ref.Append("id"))
}

// Token returns a reference to field token of azurerm_source_control_token.
func (sct sourceControlTokenAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(sct.ref.Append("token"))
}

// TokenSecret returns a reference to field token_secret of azurerm_source_control_token.
func (sct sourceControlTokenAttributes) TokenSecret() terra.StringValue {
	return terra.ReferenceAsString(sct.ref.Append("token_secret"))
}

// Type returns a reference to field type of azurerm_source_control_token.
func (sct sourceControlTokenAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sct.ref.Append("type"))
}

func (sct sourceControlTokenAttributes) Timeouts() sourcecontroltoken.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sourcecontroltoken.TimeoutsAttributes](sct.ref.Append("timeouts"))
}

type sourceControlTokenState struct {
	Id          string                            `json:"id"`
	Token       string                            `json:"token"`
	TokenSecret string                            `json:"token_secret"`
	Type        string                            `json:"type"`
	Timeouts    *sourcecontroltoken.TimeoutsState `json:"timeouts"`
}
