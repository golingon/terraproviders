// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	identityplatformconfig "github.com/golingon/terraproviders/google/4.74.0/identityplatformconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityPlatformConfig creates a new instance of [IdentityPlatformConfig].
func NewIdentityPlatformConfig(name string, args IdentityPlatformConfigArgs) *IdentityPlatformConfig {
	return &IdentityPlatformConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityPlatformConfig)(nil)

// IdentityPlatformConfig represents the Terraform resource google_identity_platform_config.
type IdentityPlatformConfig struct {
	Name      string
	Args      IdentityPlatformConfigArgs
	state     *identityPlatformConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityPlatformConfig].
func (ipc *IdentityPlatformConfig) Type() string {
	return "google_identity_platform_config"
}

// LocalName returns the local name for [IdentityPlatformConfig].
func (ipc *IdentityPlatformConfig) LocalName() string {
	return ipc.Name
}

// Configuration returns the configuration (args) for [IdentityPlatformConfig].
func (ipc *IdentityPlatformConfig) Configuration() interface{} {
	return ipc.Args
}

// DependOn is used for other resources to depend on [IdentityPlatformConfig].
func (ipc *IdentityPlatformConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(ipc)
}

// Dependencies returns the list of resources [IdentityPlatformConfig] depends_on.
func (ipc *IdentityPlatformConfig) Dependencies() terra.Dependencies {
	return ipc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityPlatformConfig].
func (ipc *IdentityPlatformConfig) LifecycleManagement() *terra.Lifecycle {
	return ipc.Lifecycle
}

// Attributes returns the attributes for [IdentityPlatformConfig].
func (ipc *IdentityPlatformConfig) Attributes() identityPlatformConfigAttributes {
	return identityPlatformConfigAttributes{ref: terra.ReferenceResource(ipc)}
}

// ImportState imports the given attribute values into [IdentityPlatformConfig]'s state.
func (ipc *IdentityPlatformConfig) ImportState(av io.Reader) error {
	ipc.state = &identityPlatformConfigState{}
	if err := json.NewDecoder(av).Decode(ipc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ipc.Type(), ipc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityPlatformConfig] has state.
func (ipc *IdentityPlatformConfig) State() (*identityPlatformConfigState, bool) {
	return ipc.state, ipc.state != nil
}

// StateMust returns the state for [IdentityPlatformConfig]. Panics if the state is nil.
func (ipc *IdentityPlatformConfig) StateMust() *identityPlatformConfigState {
	if ipc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ipc.Type(), ipc.LocalName()))
	}
	return ipc.state
}

// IdentityPlatformConfigArgs contains the configurations for google_identity_platform_config.
type IdentityPlatformConfigArgs struct {
	// AutodeleteAnonymousUsers: bool, optional
	AutodeleteAnonymousUsers terra.BoolValue `hcl:"autodelete_anonymous_users,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *identityplatformconfig.Timeouts `hcl:"timeouts,block"`
}
type identityPlatformConfigAttributes struct {
	ref terra.Reference
}

// AutodeleteAnonymousUsers returns a reference to field autodelete_anonymous_users of google_identity_platform_config.
func (ipc identityPlatformConfigAttributes) AutodeleteAnonymousUsers() terra.BoolValue {
	return terra.ReferenceAsBool(ipc.ref.Append("autodelete_anonymous_users"))
}

// Id returns a reference to field id of google_identity_platform_config.
func (ipc identityPlatformConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ipc.ref.Append("id"))
}

// Name returns a reference to field name of google_identity_platform_config.
func (ipc identityPlatformConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ipc.ref.Append("name"))
}

// Project returns a reference to field project of google_identity_platform_config.
func (ipc identityPlatformConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ipc.ref.Append("project"))
}

func (ipc identityPlatformConfigAttributes) Timeouts() identityplatformconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[identityplatformconfig.TimeoutsAttributes](ipc.ref.Append("timeouts"))
}

type identityPlatformConfigState struct {
	AutodeleteAnonymousUsers bool                                  `json:"autodelete_anonymous_users"`
	Id                       string                                `json:"id"`
	Name                     string                                `json:"name"`
	Project                  string                                `json:"project"`
	Timeouts                 *identityplatformconfig.TimeoutsState `json:"timeouts"`
}
