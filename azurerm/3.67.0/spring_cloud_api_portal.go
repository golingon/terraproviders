// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudapiportal "github.com/golingon/terraproviders/azurerm/3.67.0/springcloudapiportal"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudApiPortal creates a new instance of [SpringCloudApiPortal].
func NewSpringCloudApiPortal(name string, args SpringCloudApiPortalArgs) *SpringCloudApiPortal {
	return &SpringCloudApiPortal{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudApiPortal)(nil)

// SpringCloudApiPortal represents the Terraform resource azurerm_spring_cloud_api_portal.
type SpringCloudApiPortal struct {
	Name      string
	Args      SpringCloudApiPortalArgs
	state     *springCloudApiPortalState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudApiPortal].
func (scap *SpringCloudApiPortal) Type() string {
	return "azurerm_spring_cloud_api_portal"
}

// LocalName returns the local name for [SpringCloudApiPortal].
func (scap *SpringCloudApiPortal) LocalName() string {
	return scap.Name
}

// Configuration returns the configuration (args) for [SpringCloudApiPortal].
func (scap *SpringCloudApiPortal) Configuration() interface{} {
	return scap.Args
}

// DependOn is used for other resources to depend on [SpringCloudApiPortal].
func (scap *SpringCloudApiPortal) DependOn() terra.Reference {
	return terra.ReferenceResource(scap)
}

// Dependencies returns the list of resources [SpringCloudApiPortal] depends_on.
func (scap *SpringCloudApiPortal) Dependencies() terra.Dependencies {
	return scap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudApiPortal].
func (scap *SpringCloudApiPortal) LifecycleManagement() *terra.Lifecycle {
	return scap.Lifecycle
}

// Attributes returns the attributes for [SpringCloudApiPortal].
func (scap *SpringCloudApiPortal) Attributes() springCloudApiPortalAttributes {
	return springCloudApiPortalAttributes{ref: terra.ReferenceResource(scap)}
}

// ImportState imports the given attribute values into [SpringCloudApiPortal]'s state.
func (scap *SpringCloudApiPortal) ImportState(av io.Reader) error {
	scap.state = &springCloudApiPortalState{}
	if err := json.NewDecoder(av).Decode(scap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scap.Type(), scap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudApiPortal] has state.
func (scap *SpringCloudApiPortal) State() (*springCloudApiPortalState, bool) {
	return scap.state, scap.state != nil
}

// StateMust returns the state for [SpringCloudApiPortal]. Panics if the state is nil.
func (scap *SpringCloudApiPortal) StateMust() *springCloudApiPortalState {
	if scap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scap.Type(), scap.LocalName()))
	}
	return scap.state
}

// SpringCloudApiPortalArgs contains the configurations for azurerm_spring_cloud_api_portal.
type SpringCloudApiPortalArgs struct {
	// GatewayIds: set of string, optional
	GatewayIds terra.SetValue[terra.StringValue] `hcl:"gateway_ids,attr"`
	// HttpsOnlyEnabled: bool, optional
	HttpsOnlyEnabled terra.BoolValue `hcl:"https_only_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceCount: number, optional
	InstanceCount terra.NumberValue `hcl:"instance_count,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// SpringCloudServiceId: string, required
	SpringCloudServiceId terra.StringValue `hcl:"spring_cloud_service_id,attr" validate:"required"`
	// Sso: optional
	Sso *springcloudapiportal.Sso `hcl:"sso,block"`
	// Timeouts: optional
	Timeouts *springcloudapiportal.Timeouts `hcl:"timeouts,block"`
}
type springCloudApiPortalAttributes struct {
	ref terra.Reference
}

// GatewayIds returns a reference to field gateway_ids of azurerm_spring_cloud_api_portal.
func (scap springCloudApiPortalAttributes) GatewayIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](scap.ref.Append("gateway_ids"))
}

// HttpsOnlyEnabled returns a reference to field https_only_enabled of azurerm_spring_cloud_api_portal.
func (scap springCloudApiPortalAttributes) HttpsOnlyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(scap.ref.Append("https_only_enabled"))
}

// Id returns a reference to field id of azurerm_spring_cloud_api_portal.
func (scap springCloudApiPortalAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("id"))
}

// InstanceCount returns a reference to field instance_count of azurerm_spring_cloud_api_portal.
func (scap springCloudApiPortalAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(scap.ref.Append("instance_count"))
}

// Name returns a reference to field name of azurerm_spring_cloud_api_portal.
func (scap springCloudApiPortalAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_spring_cloud_api_portal.
func (scap springCloudApiPortalAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(scap.ref.Append("public_network_access_enabled"))
}

// SpringCloudServiceId returns a reference to field spring_cloud_service_id of azurerm_spring_cloud_api_portal.
func (scap springCloudApiPortalAttributes) SpringCloudServiceId() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("spring_cloud_service_id"))
}

// Url returns a reference to field url of azurerm_spring_cloud_api_portal.
func (scap springCloudApiPortalAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("url"))
}

func (scap springCloudApiPortalAttributes) Sso() terra.ListValue[springcloudapiportal.SsoAttributes] {
	return terra.ReferenceAsList[springcloudapiportal.SsoAttributes](scap.ref.Append("sso"))
}

func (scap springCloudApiPortalAttributes) Timeouts() springcloudapiportal.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudapiportal.TimeoutsAttributes](scap.ref.Append("timeouts"))
}

type springCloudApiPortalState struct {
	GatewayIds                 []string                            `json:"gateway_ids"`
	HttpsOnlyEnabled           bool                                `json:"https_only_enabled"`
	Id                         string                              `json:"id"`
	InstanceCount              float64                             `json:"instance_count"`
	Name                       string                              `json:"name"`
	PublicNetworkAccessEnabled bool                                `json:"public_network_access_enabled"`
	SpringCloudServiceId       string                              `json:"spring_cloud_service_id"`
	Url                        string                              `json:"url"`
	Sso                        []springcloudapiportal.SsoState     `json:"sso"`
	Timeouts                   *springcloudapiportal.TimeoutsState `json:"timeouts"`
}
