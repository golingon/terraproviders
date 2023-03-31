// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appservicesourcecontrolslot "github.com/golingon/terraproviders/azurerm/3.49.0/appservicesourcecontrolslot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAppServiceSourceControlSlot(name string, args AppServiceSourceControlSlotArgs) *AppServiceSourceControlSlot {
	return &AppServiceSourceControlSlot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceSourceControlSlot)(nil)

type AppServiceSourceControlSlot struct {
	Name  string
	Args  AppServiceSourceControlSlotArgs
	state *appServiceSourceControlSlotState
}

func (asscs *AppServiceSourceControlSlot) Type() string {
	return "azurerm_app_service_source_control_slot"
}

func (asscs *AppServiceSourceControlSlot) LocalName() string {
	return asscs.Name
}

func (asscs *AppServiceSourceControlSlot) Configuration() interface{} {
	return asscs.Args
}

func (asscs *AppServiceSourceControlSlot) Attributes() appServiceSourceControlSlotAttributes {
	return appServiceSourceControlSlotAttributes{ref: terra.ReferenceResource(asscs)}
}

func (asscs *AppServiceSourceControlSlot) ImportState(av io.Reader) error {
	asscs.state = &appServiceSourceControlSlotState{}
	if err := json.NewDecoder(av).Decode(asscs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asscs.Type(), asscs.LocalName(), err)
	}
	return nil
}

func (asscs *AppServiceSourceControlSlot) State() (*appServiceSourceControlSlotState, bool) {
	return asscs.state, asscs.state != nil
}

func (asscs *AppServiceSourceControlSlot) StateMust() *appServiceSourceControlSlotState {
	if asscs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asscs.Type(), asscs.LocalName()))
	}
	return asscs.state
}

func (asscs *AppServiceSourceControlSlot) DependOn() terra.Reference {
	return terra.ReferenceResource(asscs)
}

type AppServiceSourceControlSlotArgs struct {
	// Branch: string, optional
	Branch terra.StringValue `hcl:"branch,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RepoUrl: string, optional
	RepoUrl terra.StringValue `hcl:"repo_url,attr"`
	// RollbackEnabled: bool, optional
	RollbackEnabled terra.BoolValue `hcl:"rollback_enabled,attr"`
	// SlotId: string, required
	SlotId terra.StringValue `hcl:"slot_id,attr" validate:"required"`
	// UseLocalGit: bool, optional
	UseLocalGit terra.BoolValue `hcl:"use_local_git,attr"`
	// UseManualIntegration: bool, optional
	UseManualIntegration terra.BoolValue `hcl:"use_manual_integration,attr"`
	// UseMercurial: bool, optional
	UseMercurial terra.BoolValue `hcl:"use_mercurial,attr"`
	// GithubActionConfiguration: optional
	GithubActionConfiguration *appservicesourcecontrolslot.GithubActionConfiguration `hcl:"github_action_configuration,block"`
	// Timeouts: optional
	Timeouts *appservicesourcecontrolslot.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that AppServiceSourceControlSlot depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type appServiceSourceControlSlotAttributes struct {
	ref terra.Reference
}

func (asscs appServiceSourceControlSlotAttributes) Branch() terra.StringValue {
	return terra.ReferenceString(asscs.ref.Append("branch"))
}

func (asscs appServiceSourceControlSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceString(asscs.ref.Append("id"))
}

func (asscs appServiceSourceControlSlotAttributes) RepoUrl() terra.StringValue {
	return terra.ReferenceString(asscs.ref.Append("repo_url"))
}

func (asscs appServiceSourceControlSlotAttributes) RollbackEnabled() terra.BoolValue {
	return terra.ReferenceBool(asscs.ref.Append("rollback_enabled"))
}

func (asscs appServiceSourceControlSlotAttributes) ScmType() terra.StringValue {
	return terra.ReferenceString(asscs.ref.Append("scm_type"))
}

func (asscs appServiceSourceControlSlotAttributes) SlotId() terra.StringValue {
	return terra.ReferenceString(asscs.ref.Append("slot_id"))
}

func (asscs appServiceSourceControlSlotAttributes) UseLocalGit() terra.BoolValue {
	return terra.ReferenceBool(asscs.ref.Append("use_local_git"))
}

func (asscs appServiceSourceControlSlotAttributes) UseManualIntegration() terra.BoolValue {
	return terra.ReferenceBool(asscs.ref.Append("use_manual_integration"))
}

func (asscs appServiceSourceControlSlotAttributes) UseMercurial() terra.BoolValue {
	return terra.ReferenceBool(asscs.ref.Append("use_mercurial"))
}

func (asscs appServiceSourceControlSlotAttributes) UsesGithubAction() terra.BoolValue {
	return terra.ReferenceBool(asscs.ref.Append("uses_github_action"))
}

func (asscs appServiceSourceControlSlotAttributes) GithubActionConfiguration() terra.ListValue[appservicesourcecontrolslot.GithubActionConfigurationAttributes] {
	return terra.ReferenceList[appservicesourcecontrolslot.GithubActionConfigurationAttributes](asscs.ref.Append("github_action_configuration"))
}

func (asscs appServiceSourceControlSlotAttributes) Timeouts() appservicesourcecontrolslot.TimeoutsAttributes {
	return terra.ReferenceSingle[appservicesourcecontrolslot.TimeoutsAttributes](asscs.ref.Append("timeouts"))
}

type appServiceSourceControlSlotState struct {
	Branch                    string                                                       `json:"branch"`
	Id                        string                                                       `json:"id"`
	RepoUrl                   string                                                       `json:"repo_url"`
	RollbackEnabled           bool                                                         `json:"rollback_enabled"`
	ScmType                   string                                                       `json:"scm_type"`
	SlotId                    string                                                       `json:"slot_id"`
	UseLocalGit               bool                                                         `json:"use_local_git"`
	UseManualIntegration      bool                                                         `json:"use_manual_integration"`
	UseMercurial              bool                                                         `json:"use_mercurial"`
	UsesGithubAction          bool                                                         `json:"uses_github_action"`
	GithubActionConfiguration []appservicesourcecontrolslot.GithubActionConfigurationState `json:"github_action_configuration"`
	Timeouts                  *appservicesourcecontrolslot.TimeoutsState                   `json:"timeouts"`
}
