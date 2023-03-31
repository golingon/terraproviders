// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	resourcedeploymentscriptazurecli "github.com/golingon/terraproviders/azurerm/3.49.0/resourcedeploymentscriptazurecli"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewResourceDeploymentScriptAzureCli(name string, args ResourceDeploymentScriptAzureCliArgs) *ResourceDeploymentScriptAzureCli {
	return &ResourceDeploymentScriptAzureCli{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceDeploymentScriptAzureCli)(nil)

type ResourceDeploymentScriptAzureCli struct {
	Name  string
	Args  ResourceDeploymentScriptAzureCliArgs
	state *resourceDeploymentScriptAzureCliState
}

func (rdsac *ResourceDeploymentScriptAzureCli) Type() string {
	return "azurerm_resource_deployment_script_azure_cli"
}

func (rdsac *ResourceDeploymentScriptAzureCli) LocalName() string {
	return rdsac.Name
}

func (rdsac *ResourceDeploymentScriptAzureCli) Configuration() interface{} {
	return rdsac.Args
}

func (rdsac *ResourceDeploymentScriptAzureCli) Attributes() resourceDeploymentScriptAzureCliAttributes {
	return resourceDeploymentScriptAzureCliAttributes{ref: terra.ReferenceResource(rdsac)}
}

func (rdsac *ResourceDeploymentScriptAzureCli) ImportState(av io.Reader) error {
	rdsac.state = &resourceDeploymentScriptAzureCliState{}
	if err := json.NewDecoder(av).Decode(rdsac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rdsac.Type(), rdsac.LocalName(), err)
	}
	return nil
}

func (rdsac *ResourceDeploymentScriptAzureCli) State() (*resourceDeploymentScriptAzureCliState, bool) {
	return rdsac.state, rdsac.state != nil
}

func (rdsac *ResourceDeploymentScriptAzureCli) StateMust() *resourceDeploymentScriptAzureCliState {
	if rdsac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rdsac.Type(), rdsac.LocalName()))
	}
	return rdsac.state
}

func (rdsac *ResourceDeploymentScriptAzureCli) DependOn() terra.Reference {
	return terra.ReferenceResource(rdsac)
}

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
	// DependsOn contains resources that ResourceDeploymentScriptAzureCli depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type resourceDeploymentScriptAzureCliAttributes struct {
	ref terra.Reference
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) CleanupPreference() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("cleanup_preference"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) CommandLine() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("command_line"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) ForceUpdateTag() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("force_update_tag"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("id"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Location() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("location"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Name() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("name"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Outputs() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("outputs"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) PrimaryScriptUri() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("primary_script_uri"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("resource_group_name"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) RetentionInterval() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("retention_interval"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) ScriptContent() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("script_content"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) SupportingScriptUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](rdsac.ref.Append("supporting_script_uris"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](rdsac.ref.Append("tags"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Timeout() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("timeout"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Version() terra.StringValue {
	return terra.ReferenceString(rdsac.ref.Append("version"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Container() terra.ListValue[resourcedeploymentscriptazurecli.ContainerAttributes] {
	return terra.ReferenceList[resourcedeploymentscriptazurecli.ContainerAttributes](rdsac.ref.Append("container"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) EnvironmentVariable() terra.SetValue[resourcedeploymentscriptazurecli.EnvironmentVariableAttributes] {
	return terra.ReferenceSet[resourcedeploymentscriptazurecli.EnvironmentVariableAttributes](rdsac.ref.Append("environment_variable"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Identity() terra.ListValue[resourcedeploymentscriptazurecli.IdentityAttributes] {
	return terra.ReferenceList[resourcedeploymentscriptazurecli.IdentityAttributes](rdsac.ref.Append("identity"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) StorageAccount() terra.ListValue[resourcedeploymentscriptazurecli.StorageAccountAttributes] {
	return terra.ReferenceList[resourcedeploymentscriptazurecli.StorageAccountAttributes](rdsac.ref.Append("storage_account"))
}

func (rdsac resourceDeploymentScriptAzureCliAttributes) Timeouts() resourcedeploymentscriptazurecli.TimeoutsAttributes {
	return terra.ReferenceSingle[resourcedeploymentscriptazurecli.TimeoutsAttributes](rdsac.ref.Append("timeouts"))
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
