// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sccmuteconfig "github.com/golingon/terraproviders/google/4.72.1/sccmuteconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSccMuteConfig creates a new instance of [SccMuteConfig].
func NewSccMuteConfig(name string, args SccMuteConfigArgs) *SccMuteConfig {
	return &SccMuteConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SccMuteConfig)(nil)

// SccMuteConfig represents the Terraform resource google_scc_mute_config.
type SccMuteConfig struct {
	Name      string
	Args      SccMuteConfigArgs
	state     *sccMuteConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SccMuteConfig].
func (smc *SccMuteConfig) Type() string {
	return "google_scc_mute_config"
}

// LocalName returns the local name for [SccMuteConfig].
func (smc *SccMuteConfig) LocalName() string {
	return smc.Name
}

// Configuration returns the configuration (args) for [SccMuteConfig].
func (smc *SccMuteConfig) Configuration() interface{} {
	return smc.Args
}

// DependOn is used for other resources to depend on [SccMuteConfig].
func (smc *SccMuteConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(smc)
}

// Dependencies returns the list of resources [SccMuteConfig] depends_on.
func (smc *SccMuteConfig) Dependencies() terra.Dependencies {
	return smc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SccMuteConfig].
func (smc *SccMuteConfig) LifecycleManagement() *terra.Lifecycle {
	return smc.Lifecycle
}

// Attributes returns the attributes for [SccMuteConfig].
func (smc *SccMuteConfig) Attributes() sccMuteConfigAttributes {
	return sccMuteConfigAttributes{ref: terra.ReferenceResource(smc)}
}

// ImportState imports the given attribute values into [SccMuteConfig]'s state.
func (smc *SccMuteConfig) ImportState(av io.Reader) error {
	smc.state = &sccMuteConfigState{}
	if err := json.NewDecoder(av).Decode(smc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smc.Type(), smc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SccMuteConfig] has state.
func (smc *SccMuteConfig) State() (*sccMuteConfigState, bool) {
	return smc.state, smc.state != nil
}

// StateMust returns the state for [SccMuteConfig]. Panics if the state is nil.
func (smc *SccMuteConfig) StateMust() *sccMuteConfigState {
	if smc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smc.Type(), smc.LocalName()))
	}
	return smc.state
}

// SccMuteConfigArgs contains the configurations for google_scc_mute_config.
type SccMuteConfigArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Filter: string, required
	Filter terra.StringValue `hcl:"filter,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MuteConfigId: string, required
	MuteConfigId terra.StringValue `hcl:"mute_config_id,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sccmuteconfig.Timeouts `hcl:"timeouts,block"`
}
type sccMuteConfigAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_scc_mute_config.
func (smc sccMuteConfigAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(smc.ref.Append("create_time"))
}

// Description returns a reference to field description of google_scc_mute_config.
func (smc sccMuteConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(smc.ref.Append("description"))
}

// Filter returns a reference to field filter of google_scc_mute_config.
func (smc sccMuteConfigAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(smc.ref.Append("filter"))
}

// Id returns a reference to field id of google_scc_mute_config.
func (smc sccMuteConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smc.ref.Append("id"))
}

// MostRecentEditor returns a reference to field most_recent_editor of google_scc_mute_config.
func (smc sccMuteConfigAttributes) MostRecentEditor() terra.StringValue {
	return terra.ReferenceAsString(smc.ref.Append("most_recent_editor"))
}

// MuteConfigId returns a reference to field mute_config_id of google_scc_mute_config.
func (smc sccMuteConfigAttributes) MuteConfigId() terra.StringValue {
	return terra.ReferenceAsString(smc.ref.Append("mute_config_id"))
}

// Name returns a reference to field name of google_scc_mute_config.
func (smc sccMuteConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smc.ref.Append("name"))
}

// Parent returns a reference to field parent of google_scc_mute_config.
func (smc sccMuteConfigAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(smc.ref.Append("parent"))
}

// UpdateTime returns a reference to field update_time of google_scc_mute_config.
func (smc sccMuteConfigAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(smc.ref.Append("update_time"))
}

func (smc sccMuteConfigAttributes) Timeouts() sccmuteconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sccmuteconfig.TimeoutsAttributes](smc.ref.Append("timeouts"))
}

type sccMuteConfigState struct {
	CreateTime       string                       `json:"create_time"`
	Description      string                       `json:"description"`
	Filter           string                       `json:"filter"`
	Id               string                       `json:"id"`
	MostRecentEditor string                       `json:"most_recent_editor"`
	MuteConfigId     string                       `json:"mute_config_id"`
	Name             string                       `json:"name"`
	Parent           string                       `json:"parent"`
	UpdateTime       string                       `json:"update_time"`
	Timeouts         *sccmuteconfig.TimeoutsState `json:"timeouts"`
}
