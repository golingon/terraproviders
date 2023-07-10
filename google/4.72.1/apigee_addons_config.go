// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeeaddonsconfig "github.com/golingon/terraproviders/google/4.72.1/apigeeaddonsconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeAddonsConfig creates a new instance of [ApigeeAddonsConfig].
func NewApigeeAddonsConfig(name string, args ApigeeAddonsConfigArgs) *ApigeeAddonsConfig {
	return &ApigeeAddonsConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeAddonsConfig)(nil)

// ApigeeAddonsConfig represents the Terraform resource google_apigee_addons_config.
type ApigeeAddonsConfig struct {
	Name      string
	Args      ApigeeAddonsConfigArgs
	state     *apigeeAddonsConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeAddonsConfig].
func (aac *ApigeeAddonsConfig) Type() string {
	return "google_apigee_addons_config"
}

// LocalName returns the local name for [ApigeeAddonsConfig].
func (aac *ApigeeAddonsConfig) LocalName() string {
	return aac.Name
}

// Configuration returns the configuration (args) for [ApigeeAddonsConfig].
func (aac *ApigeeAddonsConfig) Configuration() interface{} {
	return aac.Args
}

// DependOn is used for other resources to depend on [ApigeeAddonsConfig].
func (aac *ApigeeAddonsConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(aac)
}

// Dependencies returns the list of resources [ApigeeAddonsConfig] depends_on.
func (aac *ApigeeAddonsConfig) Dependencies() terra.Dependencies {
	return aac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeAddonsConfig].
func (aac *ApigeeAddonsConfig) LifecycleManagement() *terra.Lifecycle {
	return aac.Lifecycle
}

// Attributes returns the attributes for [ApigeeAddonsConfig].
func (aac *ApigeeAddonsConfig) Attributes() apigeeAddonsConfigAttributes {
	return apigeeAddonsConfigAttributes{ref: terra.ReferenceResource(aac)}
}

// ImportState imports the given attribute values into [ApigeeAddonsConfig]'s state.
func (aac *ApigeeAddonsConfig) ImportState(av io.Reader) error {
	aac.state = &apigeeAddonsConfigState{}
	if err := json.NewDecoder(av).Decode(aac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aac.Type(), aac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeAddonsConfig] has state.
func (aac *ApigeeAddonsConfig) State() (*apigeeAddonsConfigState, bool) {
	return aac.state, aac.state != nil
}

// StateMust returns the state for [ApigeeAddonsConfig]. Panics if the state is nil.
func (aac *ApigeeAddonsConfig) StateMust() *apigeeAddonsConfigState {
	if aac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aac.Type(), aac.LocalName()))
	}
	return aac.state
}

// ApigeeAddonsConfigArgs contains the configurations for google_apigee_addons_config.
type ApigeeAddonsConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Org: string, required
	Org terra.StringValue `hcl:"org,attr" validate:"required"`
	// AddonsConfig: optional
	AddonsConfig *apigeeaddonsconfig.AddonsConfig `hcl:"addons_config,block"`
	// Timeouts: optional
	Timeouts *apigeeaddonsconfig.Timeouts `hcl:"timeouts,block"`
}
type apigeeAddonsConfigAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_apigee_addons_config.
func (aac apigeeAddonsConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("id"))
}

// Org returns a reference to field org of google_apigee_addons_config.
func (aac apigeeAddonsConfigAttributes) Org() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("org"))
}

func (aac apigeeAddonsConfigAttributes) AddonsConfig() terra.ListValue[apigeeaddonsconfig.AddonsConfigAttributes] {
	return terra.ReferenceAsList[apigeeaddonsconfig.AddonsConfigAttributes](aac.ref.Append("addons_config"))
}

func (aac apigeeAddonsConfigAttributes) Timeouts() apigeeaddonsconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeeaddonsconfig.TimeoutsAttributes](aac.ref.Append("timeouts"))
}

type apigeeAddonsConfigState struct {
	Id           string                                 `json:"id"`
	Org          string                                 `json:"org"`
	AddonsConfig []apigeeaddonsconfig.AddonsConfigState `json:"addons_config"`
	Timeouts     *apigeeaddonsconfig.TimeoutsState      `json:"timeouts"`
}
