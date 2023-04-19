// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	nginxdeployment "github.com/golingon/terraproviders/azurerm/3.52.0/nginxdeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNginxDeployment creates a new instance of [NginxDeployment].
func NewNginxDeployment(name string, args NginxDeploymentArgs) *NginxDeployment {
	return &NginxDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NginxDeployment)(nil)

// NginxDeployment represents the Terraform resource azurerm_nginx_deployment.
type NginxDeployment struct {
	Name      string
	Args      NginxDeploymentArgs
	state     *nginxDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NginxDeployment].
func (nd *NginxDeployment) Type() string {
	return "azurerm_nginx_deployment"
}

// LocalName returns the local name for [NginxDeployment].
func (nd *NginxDeployment) LocalName() string {
	return nd.Name
}

// Configuration returns the configuration (args) for [NginxDeployment].
func (nd *NginxDeployment) Configuration() interface{} {
	return nd.Args
}

// DependOn is used for other resources to depend on [NginxDeployment].
func (nd *NginxDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(nd)
}

// Dependencies returns the list of resources [NginxDeployment] depends_on.
func (nd *NginxDeployment) Dependencies() terra.Dependencies {
	return nd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NginxDeployment].
func (nd *NginxDeployment) LifecycleManagement() *terra.Lifecycle {
	return nd.Lifecycle
}

// Attributes returns the attributes for [NginxDeployment].
func (nd *NginxDeployment) Attributes() nginxDeploymentAttributes {
	return nginxDeploymentAttributes{ref: terra.ReferenceResource(nd)}
}

// ImportState imports the given attribute values into [NginxDeployment]'s state.
func (nd *NginxDeployment) ImportState(av io.Reader) error {
	nd.state = &nginxDeploymentState{}
	if err := json.NewDecoder(av).Decode(nd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nd.Type(), nd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NginxDeployment] has state.
func (nd *NginxDeployment) State() (*nginxDeploymentState, bool) {
	return nd.state, nd.state != nil
}

// StateMust returns the state for [NginxDeployment]. Panics if the state is nil.
func (nd *NginxDeployment) StateMust() *nginxDeploymentState {
	if nd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nd.Type(), nd.LocalName()))
	}
	return nd.state
}

// NginxDeploymentArgs contains the configurations for azurerm_nginx_deployment.
type NginxDeploymentArgs struct {
	// DiagnoseSupportEnabled: bool, optional
	DiagnoseSupportEnabled terra.BoolValue `hcl:"diagnose_support_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedResourceGroup: string, optional
	ManagedResourceGroup terra.StringValue `hcl:"managed_resource_group,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// FrontendPrivate: min=0
	FrontendPrivate []nginxdeployment.FrontendPrivate `hcl:"frontend_private,block" validate:"min=0"`
	// FrontendPublic: optional
	FrontendPublic *nginxdeployment.FrontendPublic `hcl:"frontend_public,block"`
	// Identity: optional
	Identity *nginxdeployment.Identity `hcl:"identity,block"`
	// LoggingStorageAccount: min=0
	LoggingStorageAccount []nginxdeployment.LoggingStorageAccount `hcl:"logging_storage_account,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []nginxdeployment.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *nginxdeployment.Timeouts `hcl:"timeouts,block"`
}
type nginxDeploymentAttributes struct {
	ref terra.Reference
}

// DiagnoseSupportEnabled returns a reference to field diagnose_support_enabled of azurerm_nginx_deployment.
func (nd nginxDeploymentAttributes) DiagnoseSupportEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(nd.ref.Append("diagnose_support_enabled"))
}

// Id returns a reference to field id of azurerm_nginx_deployment.
func (nd nginxDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of azurerm_nginx_deployment.
func (nd nginxDeploymentAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("ip_address"))
}

// Location returns a reference to field location of azurerm_nginx_deployment.
func (nd nginxDeploymentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("location"))
}

// ManagedResourceGroup returns a reference to field managed_resource_group of azurerm_nginx_deployment.
func (nd nginxDeploymentAttributes) ManagedResourceGroup() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("managed_resource_group"))
}

// Name returns a reference to field name of azurerm_nginx_deployment.
func (nd nginxDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("name"))
}

// NginxVersion returns a reference to field nginx_version of azurerm_nginx_deployment.
func (nd nginxDeploymentAttributes) NginxVersion() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("nginx_version"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_nginx_deployment.
func (nd nginxDeploymentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_nginx_deployment.
func (nd nginxDeploymentAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_nginx_deployment.
func (nd nginxDeploymentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nd.ref.Append("tags"))
}

func (nd nginxDeploymentAttributes) FrontendPrivate() terra.ListValue[nginxdeployment.FrontendPrivateAttributes] {
	return terra.ReferenceAsList[nginxdeployment.FrontendPrivateAttributes](nd.ref.Append("frontend_private"))
}

func (nd nginxDeploymentAttributes) FrontendPublic() terra.ListValue[nginxdeployment.FrontendPublicAttributes] {
	return terra.ReferenceAsList[nginxdeployment.FrontendPublicAttributes](nd.ref.Append("frontend_public"))
}

func (nd nginxDeploymentAttributes) Identity() terra.ListValue[nginxdeployment.IdentityAttributes] {
	return terra.ReferenceAsList[nginxdeployment.IdentityAttributes](nd.ref.Append("identity"))
}

func (nd nginxDeploymentAttributes) LoggingStorageAccount() terra.ListValue[nginxdeployment.LoggingStorageAccountAttributes] {
	return terra.ReferenceAsList[nginxdeployment.LoggingStorageAccountAttributes](nd.ref.Append("logging_storage_account"))
}

func (nd nginxDeploymentAttributes) NetworkInterface() terra.ListValue[nginxdeployment.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[nginxdeployment.NetworkInterfaceAttributes](nd.ref.Append("network_interface"))
}

func (nd nginxDeploymentAttributes) Timeouts() nginxdeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[nginxdeployment.TimeoutsAttributes](nd.ref.Append("timeouts"))
}

type nginxDeploymentState struct {
	DiagnoseSupportEnabled bool                                         `json:"diagnose_support_enabled"`
	Id                     string                                       `json:"id"`
	IpAddress              string                                       `json:"ip_address"`
	Location               string                                       `json:"location"`
	ManagedResourceGroup   string                                       `json:"managed_resource_group"`
	Name                   string                                       `json:"name"`
	NginxVersion           string                                       `json:"nginx_version"`
	ResourceGroupName      string                                       `json:"resource_group_name"`
	Sku                    string                                       `json:"sku"`
	Tags                   map[string]string                            `json:"tags"`
	FrontendPrivate        []nginxdeployment.FrontendPrivateState       `json:"frontend_private"`
	FrontendPublic         []nginxdeployment.FrontendPublicState        `json:"frontend_public"`
	Identity               []nginxdeployment.IdentityState              `json:"identity"`
	LoggingStorageAccount  []nginxdeployment.LoggingStorageAccountState `json:"logging_storage_account"`
	NetworkInterface       []nginxdeployment.NetworkInterfaceState      `json:"network_interface"`
	Timeouts               *nginxdeployment.TimeoutsState               `json:"timeouts"`
}
