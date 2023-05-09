// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	resourcedeploymentscriptazurepowershell "github.com/golingon/terraproviders/azurerm/3.55.0/resourcedeploymentscriptazurepowershell"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourceDeploymentScriptAzurePowerShell creates a new instance of [ResourceDeploymentScriptAzurePowerShell].
func NewResourceDeploymentScriptAzurePowerShell(name string, args ResourceDeploymentScriptAzurePowerShellArgs) *ResourceDeploymentScriptAzurePowerShell {
	return &ResourceDeploymentScriptAzurePowerShell{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceDeploymentScriptAzurePowerShell)(nil)

// ResourceDeploymentScriptAzurePowerShell represents the Terraform resource azurerm_resource_deployment_script_azure_power_shell.
type ResourceDeploymentScriptAzurePowerShell struct {
	Name      string
	Args      ResourceDeploymentScriptAzurePowerShellArgs
	state     *resourceDeploymentScriptAzurePowerShellState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourceDeploymentScriptAzurePowerShell].
func (rdsaps *ResourceDeploymentScriptAzurePowerShell) Type() string {
	return "azurerm_resource_deployment_script_azure_power_shell"
}

// LocalName returns the local name for [ResourceDeploymentScriptAzurePowerShell].
func (rdsaps *ResourceDeploymentScriptAzurePowerShell) LocalName() string {
	return rdsaps.Name
}

// Configuration returns the configuration (args) for [ResourceDeploymentScriptAzurePowerShell].
func (rdsaps *ResourceDeploymentScriptAzurePowerShell) Configuration() interface{} {
	return rdsaps.Args
}

// DependOn is used for other resources to depend on [ResourceDeploymentScriptAzurePowerShell].
func (rdsaps *ResourceDeploymentScriptAzurePowerShell) DependOn() terra.Reference {
	return terra.ReferenceResource(rdsaps)
}

// Dependencies returns the list of resources [ResourceDeploymentScriptAzurePowerShell] depends_on.
func (rdsaps *ResourceDeploymentScriptAzurePowerShell) Dependencies() terra.Dependencies {
	return rdsaps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourceDeploymentScriptAzurePowerShell].
func (rdsaps *ResourceDeploymentScriptAzurePowerShell) LifecycleManagement() *terra.Lifecycle {
	return rdsaps.Lifecycle
}

// Attributes returns the attributes for [ResourceDeploymentScriptAzurePowerShell].
func (rdsaps *ResourceDeploymentScriptAzurePowerShell) Attributes() resourceDeploymentScriptAzurePowerShellAttributes {
	return resourceDeploymentScriptAzurePowerShellAttributes{ref: terra.ReferenceResource(rdsaps)}
}

// ImportState imports the given attribute values into [ResourceDeploymentScriptAzurePowerShell]'s state.
func (rdsaps *ResourceDeploymentScriptAzurePowerShell) ImportState(av io.Reader) error {
	rdsaps.state = &resourceDeploymentScriptAzurePowerShellState{}
	if err := json.NewDecoder(av).Decode(rdsaps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rdsaps.Type(), rdsaps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourceDeploymentScriptAzurePowerShell] has state.
func (rdsaps *ResourceDeploymentScriptAzurePowerShell) State() (*resourceDeploymentScriptAzurePowerShellState, bool) {
	return rdsaps.state, rdsaps.state != nil
}

// StateMust returns the state for [ResourceDeploymentScriptAzurePowerShell]. Panics if the state is nil.
func (rdsaps *ResourceDeploymentScriptAzurePowerShell) StateMust() *resourceDeploymentScriptAzurePowerShellState {
	if rdsaps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rdsaps.Type(), rdsaps.LocalName()))
	}
	return rdsaps.state
}

// ResourceDeploymentScriptAzurePowerShellArgs contains the configurations for azurerm_resource_deployment_script_azure_power_shell.
type ResourceDeploymentScriptAzurePowerShellArgs struct {
	// CleanupPreference: string, optional
	CleanupPreference terra.StringValue `hcl:"cleanup_preference,attr"`
	// CommandLine: string, optional
	CommandLine terra.StringValue `hcl:"command_line,attr"`
	// ForceUpdateTag: string, optional
	ForceUpdateTag terra.StringValue `hcl:"force_update_tag,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrimaryScriptUri: string, optional
	PrimaryScriptUri terra.StringValue `hcl:"primary_script_uri,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RetentionInterval: string, required
	RetentionInterval terra.StringValue `hcl:"retention_interval,attr" validate:"required"`
	// ScriptContent: string, optional
	ScriptContent terra.StringValue `hcl:"script_content,attr"`
	// SupportingScriptUris: list of string, optional
	SupportingScriptUris terra.ListValue[terra.StringValue] `hcl:"supporting_script_uris,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeout: string, optional
	Timeout terra.StringValue `hcl:"timeout,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Container: optional
	Container *resourcedeploymentscriptazurepowershell.Container `hcl:"container,block"`
	// EnvironmentVariable: min=0
	EnvironmentVariable []resourcedeploymentscriptazurepowershell.EnvironmentVariable `hcl:"environment_variable,block" validate:"min=0"`
	// Identity: optional
	Identity *resourcedeploymentscriptazurepowershell.Identity `hcl:"identity,block"`
	// StorageAccount: optional
	StorageAccount *resourcedeploymentscriptazurepowershell.StorageAccount `hcl:"storage_account,block"`
	// Timeouts: optional
	Timeouts *resourcedeploymentscriptazurepowershell.Timeouts `hcl:"timeouts,block"`
}
type resourceDeploymentScriptAzurePowerShellAttributes struct {
	ref terra.Reference
}

// CleanupPreference returns a reference to field cleanup_preference of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) CleanupPreference() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("cleanup_preference"))
}

// CommandLine returns a reference to field command_line of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) CommandLine() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("command_line"))
}

// ForceUpdateTag returns a reference to field force_update_tag of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) ForceUpdateTag() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("force_update_tag"))
}

// Id returns a reference to field id of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("name"))
}

// Outputs returns a reference to field outputs of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) Outputs() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("outputs"))
}

// PrimaryScriptUri returns a reference to field primary_script_uri of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) PrimaryScriptUri() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("primary_script_uri"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("resource_group_name"))
}

// RetentionInterval returns a reference to field retention_interval of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) RetentionInterval() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("retention_interval"))
}

// ScriptContent returns a reference to field script_content of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) ScriptContent() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("script_content"))
}

// SupportingScriptUris returns a reference to field supporting_script_uris of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) SupportingScriptUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rdsaps.ref.Append("supporting_script_uris"))
}

// Tags returns a reference to field tags of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rdsaps.ref.Append("tags"))
}

// Timeout returns a reference to field timeout of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) Timeout() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("timeout"))
}

// Version returns a reference to field version of azurerm_resource_deployment_script_azure_power_shell.
func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(rdsaps.ref.Append("version"))
}

func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) Container() terra.ListValue[resourcedeploymentscriptazurepowershell.ContainerAttributes] {
	return terra.ReferenceAsList[resourcedeploymentscriptazurepowershell.ContainerAttributes](rdsaps.ref.Append("container"))
}

func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) EnvironmentVariable() terra.SetValue[resourcedeploymentscriptazurepowershell.EnvironmentVariableAttributes] {
	return terra.ReferenceAsSet[resourcedeploymentscriptazurepowershell.EnvironmentVariableAttributes](rdsaps.ref.Append("environment_variable"))
}

func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) Identity() terra.ListValue[resourcedeploymentscriptazurepowershell.IdentityAttributes] {
	return terra.ReferenceAsList[resourcedeploymentscriptazurepowershell.IdentityAttributes](rdsaps.ref.Append("identity"))
}

func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) StorageAccount() terra.ListValue[resourcedeploymentscriptazurepowershell.StorageAccountAttributes] {
	return terra.ReferenceAsList[resourcedeploymentscriptazurepowershell.StorageAccountAttributes](rdsaps.ref.Append("storage_account"))
}

func (rdsaps resourceDeploymentScriptAzurePowerShellAttributes) Timeouts() resourcedeploymentscriptazurepowershell.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcedeploymentscriptazurepowershell.TimeoutsAttributes](rdsaps.ref.Append("timeouts"))
}

type resourceDeploymentScriptAzurePowerShellState struct {
	CleanupPreference    string                                                             `json:"cleanup_preference"`
	CommandLine          string                                                             `json:"command_line"`
	ForceUpdateTag       string                                                             `json:"force_update_tag"`
	Id                   string                                                             `json:"id"`
	Location             string                                                             `json:"location"`
	Name                 string                                                             `json:"name"`
	Outputs              string                                                             `json:"outputs"`
	PrimaryScriptUri     string                                                             `json:"primary_script_uri"`
	ResourceGroupName    string                                                             `json:"resource_group_name"`
	RetentionInterval    string                                                             `json:"retention_interval"`
	ScriptContent        string                                                             `json:"script_content"`
	SupportingScriptUris []string                                                           `json:"supporting_script_uris"`
	Tags                 map[string]string                                                  `json:"tags"`
	Timeout              string                                                             `json:"timeout"`
	Version              string                                                             `json:"version"`
	Container            []resourcedeploymentscriptazurepowershell.ContainerState           `json:"container"`
	EnvironmentVariable  []resourcedeploymentscriptazurepowershell.EnvironmentVariableState `json:"environment_variable"`
	Identity             []resourcedeploymentscriptazurepowershell.IdentityState            `json:"identity"`
	StorageAccount       []resourcedeploymentscriptazurepowershell.StorageAccountState      `json:"storage_account"`
	Timeouts             *resourcedeploymentscriptazurepowershell.TimeoutsState             `json:"timeouts"`
}
