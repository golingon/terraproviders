// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatelinkservice "github.com/golingon/terraproviders/azurerm/3.49.0/privatelinkservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateLinkService creates a new instance of [PrivateLinkService].
func NewPrivateLinkService(name string, args PrivateLinkServiceArgs) *PrivateLinkService {
	return &PrivateLinkService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateLinkService)(nil)

// PrivateLinkService represents the Terraform resource azurerm_private_link_service.
type PrivateLinkService struct {
	Name      string
	Args      PrivateLinkServiceArgs
	state     *privateLinkServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateLinkService].
func (pls *PrivateLinkService) Type() string {
	return "azurerm_private_link_service"
}

// LocalName returns the local name for [PrivateLinkService].
func (pls *PrivateLinkService) LocalName() string {
	return pls.Name
}

// Configuration returns the configuration (args) for [PrivateLinkService].
func (pls *PrivateLinkService) Configuration() interface{} {
	return pls.Args
}

// DependOn is used for other resources to depend on [PrivateLinkService].
func (pls *PrivateLinkService) DependOn() terra.Reference {
	return terra.ReferenceResource(pls)
}

// Dependencies returns the list of resources [PrivateLinkService] depends_on.
func (pls *PrivateLinkService) Dependencies() terra.Dependencies {
	return pls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateLinkService].
func (pls *PrivateLinkService) LifecycleManagement() *terra.Lifecycle {
	return pls.Lifecycle
}

// Attributes returns the attributes for [PrivateLinkService].
func (pls *PrivateLinkService) Attributes() privateLinkServiceAttributes {
	return privateLinkServiceAttributes{ref: terra.ReferenceResource(pls)}
}

// ImportState imports the given attribute values into [PrivateLinkService]'s state.
func (pls *PrivateLinkService) ImportState(av io.Reader) error {
	pls.state = &privateLinkServiceState{}
	if err := json.NewDecoder(av).Decode(pls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pls.Type(), pls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateLinkService] has state.
func (pls *PrivateLinkService) State() (*privateLinkServiceState, bool) {
	return pls.state, pls.state != nil
}

// StateMust returns the state for [PrivateLinkService]. Panics if the state is nil.
func (pls *PrivateLinkService) StateMust() *privateLinkServiceState {
	if pls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pls.Type(), pls.LocalName()))
	}
	return pls.state
}

// PrivateLinkServiceArgs contains the configurations for azurerm_private_link_service.
type PrivateLinkServiceArgs struct {
	// AutoApprovalSubscriptionIds: set of string, optional
	AutoApprovalSubscriptionIds terra.SetValue[terra.StringValue] `hcl:"auto_approval_subscription_ids,attr"`
	// EnableProxyProtocol: bool, optional
	EnableProxyProtocol terra.BoolValue `hcl:"enable_proxy_protocol,attr"`
	// Fqdns: list of string, optional
	Fqdns terra.ListValue[terra.StringValue] `hcl:"fqdns,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoadBalancerFrontendIpConfigurationIds: set of string, required
	LoadBalancerFrontendIpConfigurationIds terra.SetValue[terra.StringValue] `hcl:"load_balancer_frontend_ip_configuration_ids,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VisibilitySubscriptionIds: set of string, optional
	VisibilitySubscriptionIds terra.SetValue[terra.StringValue] `hcl:"visibility_subscription_ids,attr"`
	// NatIpConfiguration: min=1,max=8
	NatIpConfiguration []privatelinkservice.NatIpConfiguration `hcl:"nat_ip_configuration,block" validate:"min=1,max=8"`
	// Timeouts: optional
	Timeouts *privatelinkservice.Timeouts `hcl:"timeouts,block"`
}
type privateLinkServiceAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("alias"))
}

// AutoApprovalSubscriptionIds returns a reference to field auto_approval_subscription_ids of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) AutoApprovalSubscriptionIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pls.ref.Append("auto_approval_subscription_ids"))
}

// EnableProxyProtocol returns a reference to field enable_proxy_protocol of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) EnableProxyProtocol() terra.BoolValue {
	return terra.ReferenceAsBool(pls.ref.Append("enable_proxy_protocol"))
}

// Fqdns returns a reference to field fqdns of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) Fqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pls.ref.Append("fqdns"))
}

// Id returns a reference to field id of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("id"))
}

// LoadBalancerFrontendIpConfigurationIds returns a reference to field load_balancer_frontend_ip_configuration_ids of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) LoadBalancerFrontendIpConfigurationIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pls.ref.Append("load_balancer_frontend_ip_configuration_ids"))
}

// Location returns a reference to field location of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pls.ref.Append("tags"))
}

// VisibilitySubscriptionIds returns a reference to field visibility_subscription_ids of azurerm_private_link_service.
func (pls privateLinkServiceAttributes) VisibilitySubscriptionIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pls.ref.Append("visibility_subscription_ids"))
}

func (pls privateLinkServiceAttributes) NatIpConfiguration() terra.ListValue[privatelinkservice.NatIpConfigurationAttributes] {
	return terra.ReferenceAsList[privatelinkservice.NatIpConfigurationAttributes](pls.ref.Append("nat_ip_configuration"))
}

func (pls privateLinkServiceAttributes) Timeouts() privatelinkservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatelinkservice.TimeoutsAttributes](pls.ref.Append("timeouts"))
}

type privateLinkServiceState struct {
	Alias                                  string                                       `json:"alias"`
	AutoApprovalSubscriptionIds            []string                                     `json:"auto_approval_subscription_ids"`
	EnableProxyProtocol                    bool                                         `json:"enable_proxy_protocol"`
	Fqdns                                  []string                                     `json:"fqdns"`
	Id                                     string                                       `json:"id"`
	LoadBalancerFrontendIpConfigurationIds []string                                     `json:"load_balancer_frontend_ip_configuration_ids"`
	Location                               string                                       `json:"location"`
	Name                                   string                                       `json:"name"`
	ResourceGroupName                      string                                       `json:"resource_group_name"`
	Tags                                   map[string]string                            `json:"tags"`
	VisibilitySubscriptionIds              []string                                     `json:"visibility_subscription_ids"`
	NatIpConfiguration                     []privatelinkservice.NatIpConfigurationState `json:"nat_ip_configuration"`
	Timeouts                               *privatelinkservice.TimeoutsState            `json:"timeouts"`
}