// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appservicesourcecontrolslot "github.com/golingon/terraproviders/azurerm/3.64.0/appservicesourcecontrolslot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceSourceControlSlot creates a new instance of [AppServiceSourceControlSlot].
func NewAppServiceSourceControlSlot(name string, args AppServiceSourceControlSlotArgs) *AppServiceSourceControlSlot {
	return &AppServiceSourceControlSlot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceSourceControlSlot)(nil)

// AppServiceSourceControlSlot represents the Terraform resource azurerm_app_service_source_control_slot.
type AppServiceSourceControlSlot struct {
	Name      string
	Args      AppServiceSourceControlSlotArgs
	state     *appServiceSourceControlSlotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceSourceControlSlot].
func (asscs *AppServiceSourceControlSlot) Type() string {
	return "azurerm_app_service_source_control_slot"
}

// LocalName returns the local name for [AppServiceSourceControlSlot].
func (asscs *AppServiceSourceControlSlot) LocalName() string {
	return asscs.Name
}

// Configuration returns the configuration (args) for [AppServiceSourceControlSlot].
func (asscs *AppServiceSourceControlSlot) Configuration() interface{} {
	return asscs.Args
}

// DependOn is used for other resources to depend on [AppServiceSourceControlSlot].
func (asscs *AppServiceSourceControlSlot) DependOn() terra.Reference {
	return terra.ReferenceResource(asscs)
}

// Dependencies returns the list of resources [AppServiceSourceControlSlot] depends_on.
func (asscs *AppServiceSourceControlSlot) Dependencies() terra.Dependencies {
	return asscs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceSourceControlSlot].
func (asscs *AppServiceSourceControlSlot) LifecycleManagement() *terra.Lifecycle {
	return asscs.Lifecycle
}

// Attributes returns the attributes for [AppServiceSourceControlSlot].
func (asscs *AppServiceSourceControlSlot) Attributes() appServiceSourceControlSlotAttributes {
	return appServiceSourceControlSlotAttributes{ref: terra.ReferenceResource(asscs)}
}

// ImportState imports the given attribute values into [AppServiceSourceControlSlot]'s state.
func (asscs *AppServiceSourceControlSlot) ImportState(av io.Reader) error {
	asscs.state = &appServiceSourceControlSlotState{}
	if err := json.NewDecoder(av).Decode(asscs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asscs.Type(), asscs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceSourceControlSlot] has state.
func (asscs *AppServiceSourceControlSlot) State() (*appServiceSourceControlSlotState, bool) {
	return asscs.state, asscs.state != nil
}

// StateMust returns the state for [AppServiceSourceControlSlot]. Panics if the state is nil.
func (asscs *AppServiceSourceControlSlot) StateMust() *appServiceSourceControlSlotState {
	if asscs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asscs.Type(), asscs.LocalName()))
	}
	return asscs.state
}

// AppServiceSourceControlSlotArgs contains the configurations for azurerm_app_service_source_control_slot.
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
}
type appServiceSourceControlSlotAttributes struct {
	ref terra.Reference
}

// Branch returns a reference to field branch of azurerm_app_service_source_control_slot.
func (asscs appServiceSourceControlSlotAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(asscs.ref.Append("branch"))
}

// Id returns a reference to field id of azurerm_app_service_source_control_slot.
func (asscs appServiceSourceControlSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asscs.ref.Append("id"))
}

// RepoUrl returns a reference to field repo_url of azurerm_app_service_source_control_slot.
func (asscs appServiceSourceControlSlotAttributes) RepoUrl() terra.StringValue {
	return terra.ReferenceAsString(asscs.ref.Append("repo_url"))
}

// RollbackEnabled returns a reference to field rollback_enabled of azurerm_app_service_source_control_slot.
func (asscs appServiceSourceControlSlotAttributes) RollbackEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(asscs.ref.Append("rollback_enabled"))
}

// ScmType returns a reference to field scm_type of azurerm_app_service_source_control_slot.
func (asscs appServiceSourceControlSlotAttributes) ScmType() terra.StringValue {
	return terra.ReferenceAsString(asscs.ref.Append("scm_type"))
}

// SlotId returns a reference to field slot_id of azurerm_app_service_source_control_slot.
func (asscs appServiceSourceControlSlotAttributes) SlotId() terra.StringValue {
	return terra.ReferenceAsString(asscs.ref.Append("slot_id"))
}

// UseLocalGit returns a reference to field use_local_git of azurerm_app_service_source_control_slot.
func (asscs appServiceSourceControlSlotAttributes) UseLocalGit() terra.BoolValue {
	return terra.ReferenceAsBool(asscs.ref.Append("use_local_git"))
}

// UseManualIntegration returns a reference to field use_manual_integration of azurerm_app_service_source_control_slot.
func (asscs appServiceSourceControlSlotAttributes) UseManualIntegration() terra.BoolValue {
	return terra.ReferenceAsBool(asscs.ref.Append("use_manual_integration"))
}

// UseMercurial returns a reference to field use_mercurial of azurerm_app_service_source_control_slot.
func (asscs appServiceSourceControlSlotAttributes) UseMercurial() terra.BoolValue {
	return terra.ReferenceAsBool(asscs.ref.Append("use_mercurial"))
}

// UsesGithubAction returns a reference to field uses_github_action of azurerm_app_service_source_control_slot.
func (asscs appServiceSourceControlSlotAttributes) UsesGithubAction() terra.BoolValue {
	return terra.ReferenceAsBool(asscs.ref.Append("uses_github_action"))
}

func (asscs appServiceSourceControlSlotAttributes) GithubActionConfiguration() terra.ListValue[appservicesourcecontrolslot.GithubActionConfigurationAttributes] {
	return terra.ReferenceAsList[appservicesourcecontrolslot.GithubActionConfigurationAttributes](asscs.ref.Append("github_action_configuration"))
}

func (asscs appServiceSourceControlSlotAttributes) Timeouts() appservicesourcecontrolslot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appservicesourcecontrolslot.TimeoutsAttributes](asscs.ref.Append("timeouts"))
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
