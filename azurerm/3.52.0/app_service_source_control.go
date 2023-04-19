// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appservicesourcecontrol "github.com/golingon/terraproviders/azurerm/3.52.0/appservicesourcecontrol"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceSourceControl creates a new instance of [AppServiceSourceControl].
func NewAppServiceSourceControl(name string, args AppServiceSourceControlArgs) *AppServiceSourceControl {
	return &AppServiceSourceControl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceSourceControl)(nil)

// AppServiceSourceControl represents the Terraform resource azurerm_app_service_source_control.
type AppServiceSourceControl struct {
	Name      string
	Args      AppServiceSourceControlArgs
	state     *appServiceSourceControlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceSourceControl].
func (assc *AppServiceSourceControl) Type() string {
	return "azurerm_app_service_source_control"
}

// LocalName returns the local name for [AppServiceSourceControl].
func (assc *AppServiceSourceControl) LocalName() string {
	return assc.Name
}

// Configuration returns the configuration (args) for [AppServiceSourceControl].
func (assc *AppServiceSourceControl) Configuration() interface{} {
	return assc.Args
}

// DependOn is used for other resources to depend on [AppServiceSourceControl].
func (assc *AppServiceSourceControl) DependOn() terra.Reference {
	return terra.ReferenceResource(assc)
}

// Dependencies returns the list of resources [AppServiceSourceControl] depends_on.
func (assc *AppServiceSourceControl) Dependencies() terra.Dependencies {
	return assc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceSourceControl].
func (assc *AppServiceSourceControl) LifecycleManagement() *terra.Lifecycle {
	return assc.Lifecycle
}

// Attributes returns the attributes for [AppServiceSourceControl].
func (assc *AppServiceSourceControl) Attributes() appServiceSourceControlAttributes {
	return appServiceSourceControlAttributes{ref: terra.ReferenceResource(assc)}
}

// ImportState imports the given attribute values into [AppServiceSourceControl]'s state.
func (assc *AppServiceSourceControl) ImportState(av io.Reader) error {
	assc.state = &appServiceSourceControlState{}
	if err := json.NewDecoder(av).Decode(assc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", assc.Type(), assc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceSourceControl] has state.
func (assc *AppServiceSourceControl) State() (*appServiceSourceControlState, bool) {
	return assc.state, assc.state != nil
}

// StateMust returns the state for [AppServiceSourceControl]. Panics if the state is nil.
func (assc *AppServiceSourceControl) StateMust() *appServiceSourceControlState {
	if assc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", assc.Type(), assc.LocalName()))
	}
	return assc.state
}

// AppServiceSourceControlArgs contains the configurations for azurerm_app_service_source_control.
type AppServiceSourceControlArgs struct {
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// Branch: string, optional
	Branch terra.StringValue `hcl:"branch,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RepoUrl: string, optional
	RepoUrl terra.StringValue `hcl:"repo_url,attr"`
	// RollbackEnabled: bool, optional
	RollbackEnabled terra.BoolValue `hcl:"rollback_enabled,attr"`
	// UseLocalGit: bool, optional
	UseLocalGit terra.BoolValue `hcl:"use_local_git,attr"`
	// UseManualIntegration: bool, optional
	UseManualIntegration terra.BoolValue `hcl:"use_manual_integration,attr"`
	// UseMercurial: bool, optional
	UseMercurial terra.BoolValue `hcl:"use_mercurial,attr"`
	// GithubActionConfiguration: optional
	GithubActionConfiguration *appservicesourcecontrol.GithubActionConfiguration `hcl:"github_action_configuration,block"`
	// Timeouts: optional
	Timeouts *appservicesourcecontrol.Timeouts `hcl:"timeouts,block"`
}
type appServiceSourceControlAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of azurerm_app_service_source_control.
func (assc appServiceSourceControlAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(assc.ref.Append("app_id"))
}

// Branch returns a reference to field branch of azurerm_app_service_source_control.
func (assc appServiceSourceControlAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(assc.ref.Append("branch"))
}

// Id returns a reference to field id of azurerm_app_service_source_control.
func (assc appServiceSourceControlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(assc.ref.Append("id"))
}

// RepoUrl returns a reference to field repo_url of azurerm_app_service_source_control.
func (assc appServiceSourceControlAttributes) RepoUrl() terra.StringValue {
	return terra.ReferenceAsString(assc.ref.Append("repo_url"))
}

// RollbackEnabled returns a reference to field rollback_enabled of azurerm_app_service_source_control.
func (assc appServiceSourceControlAttributes) RollbackEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(assc.ref.Append("rollback_enabled"))
}

// ScmType returns a reference to field scm_type of azurerm_app_service_source_control.
func (assc appServiceSourceControlAttributes) ScmType() terra.StringValue {
	return terra.ReferenceAsString(assc.ref.Append("scm_type"))
}

// UseLocalGit returns a reference to field use_local_git of azurerm_app_service_source_control.
func (assc appServiceSourceControlAttributes) UseLocalGit() terra.BoolValue {
	return terra.ReferenceAsBool(assc.ref.Append("use_local_git"))
}

// UseManualIntegration returns a reference to field use_manual_integration of azurerm_app_service_source_control.
func (assc appServiceSourceControlAttributes) UseManualIntegration() terra.BoolValue {
	return terra.ReferenceAsBool(assc.ref.Append("use_manual_integration"))
}

// UseMercurial returns a reference to field use_mercurial of azurerm_app_service_source_control.
func (assc appServiceSourceControlAttributes) UseMercurial() terra.BoolValue {
	return terra.ReferenceAsBool(assc.ref.Append("use_mercurial"))
}

// UsesGithubAction returns a reference to field uses_github_action of azurerm_app_service_source_control.
func (assc appServiceSourceControlAttributes) UsesGithubAction() terra.BoolValue {
	return terra.ReferenceAsBool(assc.ref.Append("uses_github_action"))
}

func (assc appServiceSourceControlAttributes) GithubActionConfiguration() terra.ListValue[appservicesourcecontrol.GithubActionConfigurationAttributes] {
	return terra.ReferenceAsList[appservicesourcecontrol.GithubActionConfigurationAttributes](assc.ref.Append("github_action_configuration"))
}

func (assc appServiceSourceControlAttributes) Timeouts() appservicesourcecontrol.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appservicesourcecontrol.TimeoutsAttributes](assc.ref.Append("timeouts"))
}

type appServiceSourceControlState struct {
	AppId                     string                                                   `json:"app_id"`
	Branch                    string                                                   `json:"branch"`
	Id                        string                                                   `json:"id"`
	RepoUrl                   string                                                   `json:"repo_url"`
	RollbackEnabled           bool                                                     `json:"rollback_enabled"`
	ScmType                   string                                                   `json:"scm_type"`
	UseLocalGit               bool                                                     `json:"use_local_git"`
	UseManualIntegration      bool                                                     `json:"use_manual_integration"`
	UseMercurial              bool                                                     `json:"use_mercurial"`
	UsesGithubAction          bool                                                     `json:"uses_github_action"`
	GithubActionConfiguration []appservicesourcecontrol.GithubActionConfigurationState `json:"github_action_configuration"`
	Timeouts                  *appservicesourcecontrol.TimeoutsState                   `json:"timeouts"`
}
