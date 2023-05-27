// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	identityplatformprojectdefaultconfig "github.com/golingon/terraproviders/google/4.66.0/identityplatformprojectdefaultconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityPlatformProjectDefaultConfig creates a new instance of [IdentityPlatformProjectDefaultConfig].
func NewIdentityPlatformProjectDefaultConfig(name string, args IdentityPlatformProjectDefaultConfigArgs) *IdentityPlatformProjectDefaultConfig {
	return &IdentityPlatformProjectDefaultConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityPlatformProjectDefaultConfig)(nil)

// IdentityPlatformProjectDefaultConfig represents the Terraform resource google_identity_platform_project_default_config.
type IdentityPlatformProjectDefaultConfig struct {
	Name      string
	Args      IdentityPlatformProjectDefaultConfigArgs
	state     *identityPlatformProjectDefaultConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityPlatformProjectDefaultConfig].
func (ippdc *IdentityPlatformProjectDefaultConfig) Type() string {
	return "google_identity_platform_project_default_config"
}

// LocalName returns the local name for [IdentityPlatformProjectDefaultConfig].
func (ippdc *IdentityPlatformProjectDefaultConfig) LocalName() string {
	return ippdc.Name
}

// Configuration returns the configuration (args) for [IdentityPlatformProjectDefaultConfig].
func (ippdc *IdentityPlatformProjectDefaultConfig) Configuration() interface{} {
	return ippdc.Args
}

// DependOn is used for other resources to depend on [IdentityPlatformProjectDefaultConfig].
func (ippdc *IdentityPlatformProjectDefaultConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(ippdc)
}

// Dependencies returns the list of resources [IdentityPlatformProjectDefaultConfig] depends_on.
func (ippdc *IdentityPlatformProjectDefaultConfig) Dependencies() terra.Dependencies {
	return ippdc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityPlatformProjectDefaultConfig].
func (ippdc *IdentityPlatformProjectDefaultConfig) LifecycleManagement() *terra.Lifecycle {
	return ippdc.Lifecycle
}

// Attributes returns the attributes for [IdentityPlatformProjectDefaultConfig].
func (ippdc *IdentityPlatformProjectDefaultConfig) Attributes() identityPlatformProjectDefaultConfigAttributes {
	return identityPlatformProjectDefaultConfigAttributes{ref: terra.ReferenceResource(ippdc)}
}

// ImportState imports the given attribute values into [IdentityPlatformProjectDefaultConfig]'s state.
func (ippdc *IdentityPlatformProjectDefaultConfig) ImportState(av io.Reader) error {
	ippdc.state = &identityPlatformProjectDefaultConfigState{}
	if err := json.NewDecoder(av).Decode(ippdc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ippdc.Type(), ippdc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityPlatformProjectDefaultConfig] has state.
func (ippdc *IdentityPlatformProjectDefaultConfig) State() (*identityPlatformProjectDefaultConfigState, bool) {
	return ippdc.state, ippdc.state != nil
}

// StateMust returns the state for [IdentityPlatformProjectDefaultConfig]. Panics if the state is nil.
func (ippdc *IdentityPlatformProjectDefaultConfig) StateMust() *identityPlatformProjectDefaultConfigState {
	if ippdc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ippdc.Type(), ippdc.LocalName()))
	}
	return ippdc.state
}

// IdentityPlatformProjectDefaultConfigArgs contains the configurations for google_identity_platform_project_default_config.
type IdentityPlatformProjectDefaultConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SignIn: optional
	SignIn *identityplatformprojectdefaultconfig.SignIn `hcl:"sign_in,block"`
	// Timeouts: optional
	Timeouts *identityplatformprojectdefaultconfig.Timeouts `hcl:"timeouts,block"`
}
type identityPlatformProjectDefaultConfigAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_identity_platform_project_default_config.
func (ippdc identityPlatformProjectDefaultConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ippdc.ref.Append("id"))
}

// Name returns a reference to field name of google_identity_platform_project_default_config.
func (ippdc identityPlatformProjectDefaultConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ippdc.ref.Append("name"))
}

// Project returns a reference to field project of google_identity_platform_project_default_config.
func (ippdc identityPlatformProjectDefaultConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ippdc.ref.Append("project"))
}

func (ippdc identityPlatformProjectDefaultConfigAttributes) SignIn() terra.ListValue[identityplatformprojectdefaultconfig.SignInAttributes] {
	return terra.ReferenceAsList[identityplatformprojectdefaultconfig.SignInAttributes](ippdc.ref.Append("sign_in"))
}

func (ippdc identityPlatformProjectDefaultConfigAttributes) Timeouts() identityplatformprojectdefaultconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[identityplatformprojectdefaultconfig.TimeoutsAttributes](ippdc.ref.Append("timeouts"))
}

type identityPlatformProjectDefaultConfigState struct {
	Id       string                                              `json:"id"`
	Name     string                                              `json:"name"`
	Project  string                                              `json:"project"`
	SignIn   []identityplatformprojectdefaultconfig.SignInState  `json:"sign_in"`
	Timeouts *identityplatformprojectdefaultconfig.TimeoutsState `json:"timeouts"`
}
