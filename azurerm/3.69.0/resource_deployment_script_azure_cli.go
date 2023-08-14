// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	resourcedeploymentscriptazurecli "github.com/golingon/terraproviders/azurerm/3.69.0/resourcedeploymentscriptazurecli"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourceDeploymentScriptAzureCli creates a new instance of [ResourceDeploymentScriptAzureCli].
func NewResourceDeploymentScriptAzureCli(name string, args ResourceDeploymentScriptAzureCliArgs) *ResourceDeploymentScriptAzureCli {
	return &ResourceDeploymentScriptAzureCli{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceDeploymentScriptAzureCli)(nil)

// ResourceDeploymentScriptAzureCli represents the Terraform resource azurerm_resource_deployment_script_azure_cli.
type ResourceDeploymentScriptAzureCli struct {
	Name      string
	Args      ResourceDeploymentScriptAzureCliArgs
	state     *resourceDeploymentScriptAzureCliState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourceDeploymentScriptAzureCli].
func (rdsac *ResourceDeploymentScriptAzureCli) Type() string {
	return "azurerm_resource_deployment_script_azure_cli"
}

// LocalName returns the local name for [ResourceDeploymentScriptAzureCli].
func (rdsac *ResourceDeploymentScriptAzureCli) LocalName() string {
	return rdsac.Name
}

// Configuration returns the configuration (args) for [ResourceDeploymentScriptAzureCli].
func (rdsac *ResourceDeploymentScriptAzureCli) Configuration() interface{} {
	return rdsac.Args
}

// DependOn is used for other resources to depend on [ResourceDeploymentScriptAzureCli].
func (rdsac *ResourceDeploymentScriptAzureCli) DependOn() terra.Reference {
	return terra.ReferenceResource(rdsac)
}

// Dependencies returns the list of resources [ResourceDeploymentScriptAzureCli] depends_on.
func (rdsac *ResourceDeploymentScriptAzureCli) Dependencies() terra.Dependencies {
	return rdsac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourceDeploymentScriptAzureCli].
func (rdsac *ResourceDeploymentScriptAzureCli) LifecycleManagement() *terra.Lifecycle {
	return rdsac.Lifecycle
}

// Attributes returns the attributes for [ResourceDeploymentScriptAzureCli].
func (rdsac *ResourceDeploymentScriptAzureCli) Attributes() resourceDeploymentScriptAzureCliAttributes {
	return resourceDeploymentScriptAzureCliAttributes{ref: terra.ReferenceResource(rdsac)}
}

// ImportState imports the given attribute values into [ResourceDeploymentScriptAzureCli]'s state.
func (rdsac *ResourceDeploymentScriptAzureCli) ImportState(av io.Reader) error {
	rdsac.state = &resourceDeploymentScriptAzureCliState{}
	if err := json.NewDecoder(av).Decode(rdsac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rdsac.Type(), rdsac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourceDeploymentScriptAzureCli] has state.
func (rdsac *ResourceDeploymentScriptAzureCli) State() (*resourceDeploymentScriptAzureCliState, bool) {
	return rdsac.state, rdsac.state != nil
}

// StateMust returns the state for [ResourceDeploymentScriptAzureCli]. Panics if the state is nil.
func (rdsac *ResourceDeploymentScriptAzureCli) StateMust() *resourceDeploymentScriptAzureCliState {
	if rdsac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rdsac.Type(), rdsac.LocalName()))
	}
	return rdsac.state
}

// ResourceDeploymentScriptAzureCliArgs contains the configurations for azurerm_resource_deployment_script_azure_cli.
type ResourceDeploymentScriptAzureCliArgs struct {
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
	Container *resourcedeploymentscriptazurecli.Container `hcl:"container,block"`
	// EnvironmentVariable: min=0
	EnvironmentVariable []resourcedeploymentscriptazurecli.EnvironmentVariable `hcl:"environment_variable,block" validate:"min=0"`
	// Identity: optional
	Identity *resourcedeploymentscriptazurecli.Identity `hcl:"identity,block"`
	// StorageAccount: optional
	StorageAccount *resourcedeploymentscriptazurecli.StorageAccount `hcl:"storage_account,block"`
	// Timeouts: optional
	Timeouts *resourcedeploymentscriptazurecli.Timeouts `hcl:"timeouts,block"`
}
type resourceDeploymentScriptAzureCliAttributes struct {
	ref terra.Reference
}

// CleanupPreference returns a reference to field cleanup_preference of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) CleanupPreference() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("cleanup_preference"))
}

// CommandLine returns a reference to field command_line of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) CommandLine() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("command_line"))
}

// ForceUpdateTag returns a reference to field force_update_tag of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) ForceUpdateTag() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("force_update_tag"))
}

// Id returns a reference to field id of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("name"))
}

// Outputs returns a reference to field outputs of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) Outputs() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("outputs"))
}

// PrimaryScriptUri returns a reference to field primary_script_uri of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) PrimaryScriptUri() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("primary_script_uri"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("resource_group_name"))
}

// RetentionInterval returns a reference to field retention_interval of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) RetentionInterval() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("retention_interval"))
}

// ScriptContent returns a reference to field script_content of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) ScriptContent() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("script_content"))
}

// SupportingScriptUris returns a reference to field supporting_script_uris of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) SupportingScriptUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rdsac.ref.Append("supporting_script_uris"))
}

// Tags returns a reference to field tags of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rdsac.ref.Append("tags"))
}

// Timeout returns a reference to field timeout of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) Timeout() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("timeout"))
}

// Version returns a reference to field version of azurerm_resource_deployment_script_azure_cli.
func (rdsac resourceDeploymentScriptAzureCliAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(rdsac.ref.Append("version"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Container() terra.ListValue[resourcedeploymentscriptazurecli.ContainerAttributes] {
	return terra.ReferenceAsList[resourcedeploymentscriptazurecli.ContainerAttributes](rdsac.ref.Append("container"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) EnvironmentVariable() terra.SetValue[resourcedeploymentscriptazurecli.EnvironmentVariableAttributes] {
	return terra.ReferenceAsSet[resourcedeploymentscriptazurecli.EnvironmentVariableAttributes](rdsac.ref.Append("environment_variable"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Identity() terra.ListValue[resourcedeploymentscriptazurecli.IdentityAttributes] {
	return terra.ReferenceAsList[resourcedeploymentscriptazurecli.IdentityAttributes](rdsac.ref.Append("identity"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) StorageAccount() terra.ListValue[resourcedeploymentscriptazurecli.StorageAccountAttributes] {
	return terra.ReferenceAsList[resourcedeploymentscriptazurecli.StorageAccountAttributes](rdsac.ref.Append("storage_account"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Timeouts() resourcedeploymentscriptazurecli.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcedeploymentscriptazurecli.TimeoutsAttributes](rdsac.ref.Append("timeouts"))
}

type resourceDeploymentScriptAzureCliState struct {
	CleanupPreference    string                                                      `json:"cleanup_preference"`
	CommandLine          string                                                      `json:"command_line"`
	ForceUpdateTag       string                                                      `json:"force_update_tag"`
	Id                   string                                                      `json:"id"`
	Location             string                                                      `json:"location"`
	Name                 string                                                      `json:"name"`
	Outputs              string                                                      `json:"outputs"`
	PrimaryScriptUri     string                                                      `json:"primary_script_uri"`
	ResourceGroupName    string                                                      `json:"resource_group_name"`
	RetentionInterval    string                                                      `json:"retention_interval"`
	ScriptContent        string                                                      `json:"script_content"`
	SupportingScriptUris []string                                                    `json:"supporting_script_uris"`
	Tags                 map[string]string                                           `json:"tags"`
	Timeout              string                                                      `json:"timeout"`
	Version              string                                                      `json:"version"`
	Container            []resourcedeploymentscriptazurecli.ContainerState           `json:"container"`
	EnvironmentVariable  []resourcedeploymentscriptazurecli.EnvironmentVariableState `json:"environment_variable"`
	Identity             []resourcedeploymentscriptazurecli.IdentityState            `json:"identity"`
	StorageAccount       []resourcedeploymentscriptazurecli.StorageAccountState      `json:"storage_account"`
	Timeouts             *resourcedeploymentscriptazurecli.TimeoutsState             `json:"timeouts"`
}
