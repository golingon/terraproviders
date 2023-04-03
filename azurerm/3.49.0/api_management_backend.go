// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementbackend "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagementbackend"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementBackend creates a new instance of [ApiManagementBackend].
func NewApiManagementBackend(name string, args ApiManagementBackendArgs) *ApiManagementBackend {
	return &ApiManagementBackend{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementBackend)(nil)

// ApiManagementBackend represents the Terraform resource azurerm_api_management_backend.
type ApiManagementBackend struct {
	Name      string
	Args      ApiManagementBackendArgs
	state     *apiManagementBackendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementBackend].
func (amb *ApiManagementBackend) Type() string {
	return "azurerm_api_management_backend"
}

// LocalName returns the local name for [ApiManagementBackend].
func (amb *ApiManagementBackend) LocalName() string {
	return amb.Name
}

// Configuration returns the configuration (args) for [ApiManagementBackend].
func (amb *ApiManagementBackend) Configuration() interface{} {
	return amb.Args
}

// DependOn is used for other resources to depend on [ApiManagementBackend].
func (amb *ApiManagementBackend) DependOn() terra.Reference {
	return terra.ReferenceResource(amb)
}

// Dependencies returns the list of resources [ApiManagementBackend] depends_on.
func (amb *ApiManagementBackend) Dependencies() terra.Dependencies {
	return amb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementBackend].
func (amb *ApiManagementBackend) LifecycleManagement() *terra.Lifecycle {
	return amb.Lifecycle
}

// Attributes returns the attributes for [ApiManagementBackend].
func (amb *ApiManagementBackend) Attributes() apiManagementBackendAttributes {
	return apiManagementBackendAttributes{ref: terra.ReferenceResource(amb)}
}

// ImportState imports the given attribute values into [ApiManagementBackend]'s state.
func (amb *ApiManagementBackend) ImportState(av io.Reader) error {
	amb.state = &apiManagementBackendState{}
	if err := json.NewDecoder(av).Decode(amb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amb.Type(), amb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementBackend] has state.
func (amb *ApiManagementBackend) State() (*apiManagementBackendState, bool) {
	return amb.state, amb.state != nil
}

// StateMust returns the state for [ApiManagementBackend]. Panics if the state is nil.
func (amb *ApiManagementBackend) StateMust() *apiManagementBackendState {
	if amb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amb.Type(), amb.LocalName()))
	}
	return amb.state
}

// ApiManagementBackendArgs contains the configurations for azurerm_api_management_backend.
type ApiManagementBackendArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ResourceId: string, optional
	ResourceId terra.StringValue `hcl:"resource_id,attr"`
	// Title: string, optional
	Title terra.StringValue `hcl:"title,attr"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
	// Credentials: optional
	Credentials *apimanagementbackend.Credentials `hcl:"credentials,block"`
	// Proxy: optional
	Proxy *apimanagementbackend.Proxy `hcl:"proxy,block"`
	// ServiceFabricCluster: optional
	ServiceFabricCluster *apimanagementbackend.ServiceFabricCluster `hcl:"service_fabric_cluster,block"`
	// Timeouts: optional
	Timeouts *apimanagementbackend.Timeouts `hcl:"timeouts,block"`
	// Tls: optional
	Tls *apimanagementbackend.Tls `hcl:"tls,block"`
}
type apiManagementBackendAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_backend.
func (amb apiManagementBackendAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amb.ref.Append("api_management_name"))
}

// Description returns a reference to field description of azurerm_api_management_backend.
func (amb apiManagementBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amb.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_api_management_backend.
func (amb apiManagementBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amb.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_backend.
func (amb apiManagementBackendAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amb.ref.Append("name"))
}

// Protocol returns a reference to field protocol of azurerm_api_management_backend.
func (amb apiManagementBackendAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(amb.ref.Append("protocol"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_backend.
func (amb apiManagementBackendAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amb.ref.Append("resource_group_name"))
}

// ResourceId returns a reference to field resource_id of azurerm_api_management_backend.
func (amb apiManagementBackendAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(amb.ref.Append("resource_id"))
}

// Title returns a reference to field title of azurerm_api_management_backend.
func (amb apiManagementBackendAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(amb.ref.Append("title"))
}

// Url returns a reference to field url of azurerm_api_management_backend.
func (amb apiManagementBackendAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(amb.ref.Append("url"))
}

func (amb apiManagementBackendAttributes) Credentials() terra.ListValue[apimanagementbackend.CredentialsAttributes] {
	return terra.ReferenceAsList[apimanagementbackend.CredentialsAttributes](amb.ref.Append("credentials"))
}

func (amb apiManagementBackendAttributes) Proxy() terra.ListValue[apimanagementbackend.ProxyAttributes] {
	return terra.ReferenceAsList[apimanagementbackend.ProxyAttributes](amb.ref.Append("proxy"))
}

func (amb apiManagementBackendAttributes) ServiceFabricCluster() terra.ListValue[apimanagementbackend.ServiceFabricClusterAttributes] {
	return terra.ReferenceAsList[apimanagementbackend.ServiceFabricClusterAttributes](amb.ref.Append("service_fabric_cluster"))
}

func (amb apiManagementBackendAttributes) Timeouts() apimanagementbackend.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementbackend.TimeoutsAttributes](amb.ref.Append("timeouts"))
}

func (amb apiManagementBackendAttributes) Tls() terra.ListValue[apimanagementbackend.TlsAttributes] {
	return terra.ReferenceAsList[apimanagementbackend.TlsAttributes](amb.ref.Append("tls"))
}

type apiManagementBackendState struct {
	ApiManagementName    string                                           `json:"api_management_name"`
	Description          string                                           `json:"description"`
	Id                   string                                           `json:"id"`
	Name                 string                                           `json:"name"`
	Protocol             string                                           `json:"protocol"`
	ResourceGroupName    string                                           `json:"resource_group_name"`
	ResourceId           string                                           `json:"resource_id"`
	Title                string                                           `json:"title"`
	Url                  string                                           `json:"url"`
	Credentials          []apimanagementbackend.CredentialsState          `json:"credentials"`
	Proxy                []apimanagementbackend.ProxyState                `json:"proxy"`
	ServiceFabricCluster []apimanagementbackend.ServiceFabricClusterState `json:"service_fabric_cluster"`
	Timeouts             *apimanagementbackend.TimeoutsState              `json:"timeouts"`
	Tls                  []apimanagementbackend.TlsState                  `json:"tls"`
}
