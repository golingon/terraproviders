// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vmwareexpressrouteauthorization "github.com/golingon/terraproviders/azurerm/3.49.0/vmwareexpressrouteauthorization"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVmwareExpressRouteAuthorization creates a new instance of [VmwareExpressRouteAuthorization].
func NewVmwareExpressRouteAuthorization(name string, args VmwareExpressRouteAuthorizationArgs) *VmwareExpressRouteAuthorization {
	return &VmwareExpressRouteAuthorization{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VmwareExpressRouteAuthorization)(nil)

// VmwareExpressRouteAuthorization represents the Terraform resource azurerm_vmware_express_route_authorization.
type VmwareExpressRouteAuthorization struct {
	Name      string
	Args      VmwareExpressRouteAuthorizationArgs
	state     *vmwareExpressRouteAuthorizationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VmwareExpressRouteAuthorization].
func (vera *VmwareExpressRouteAuthorization) Type() string {
	return "azurerm_vmware_express_route_authorization"
}

// LocalName returns the local name for [VmwareExpressRouteAuthorization].
func (vera *VmwareExpressRouteAuthorization) LocalName() string {
	return vera.Name
}

// Configuration returns the configuration (args) for [VmwareExpressRouteAuthorization].
func (vera *VmwareExpressRouteAuthorization) Configuration() interface{} {
	return vera.Args
}

// DependOn is used for other resources to depend on [VmwareExpressRouteAuthorization].
func (vera *VmwareExpressRouteAuthorization) DependOn() terra.Reference {
	return terra.ReferenceResource(vera)
}

// Dependencies returns the list of resources [VmwareExpressRouteAuthorization] depends_on.
func (vera *VmwareExpressRouteAuthorization) Dependencies() terra.Dependencies {
	return vera.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VmwareExpressRouteAuthorization].
func (vera *VmwareExpressRouteAuthorization) LifecycleManagement() *terra.Lifecycle {
	return vera.Lifecycle
}

// Attributes returns the attributes for [VmwareExpressRouteAuthorization].
func (vera *VmwareExpressRouteAuthorization) Attributes() vmwareExpressRouteAuthorizationAttributes {
	return vmwareExpressRouteAuthorizationAttributes{ref: terra.ReferenceResource(vera)}
}

// ImportState imports the given attribute values into [VmwareExpressRouteAuthorization]'s state.
func (vera *VmwareExpressRouteAuthorization) ImportState(av io.Reader) error {
	vera.state = &vmwareExpressRouteAuthorizationState{}
	if err := json.NewDecoder(av).Decode(vera.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vera.Type(), vera.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VmwareExpressRouteAuthorization] has state.
func (vera *VmwareExpressRouteAuthorization) State() (*vmwareExpressRouteAuthorizationState, bool) {
	return vera.state, vera.state != nil
}

// StateMust returns the state for [VmwareExpressRouteAuthorization]. Panics if the state is nil.
func (vera *VmwareExpressRouteAuthorization) StateMust() *vmwareExpressRouteAuthorizationState {
	if vera.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vera.Type(), vera.LocalName()))
	}
	return vera.state
}

// VmwareExpressRouteAuthorizationArgs contains the configurations for azurerm_vmware_express_route_authorization.
type VmwareExpressRouteAuthorizationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateCloudId: string, required
	PrivateCloudId terra.StringValue `hcl:"private_cloud_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *vmwareexpressrouteauthorization.Timeouts `hcl:"timeouts,block"`
}
type vmwareExpressRouteAuthorizationAttributes struct {
	ref terra.Reference
}

// ExpressRouteAuthorizationId returns a reference to field express_route_authorization_id of azurerm_vmware_express_route_authorization.
func (vera vmwareExpressRouteAuthorizationAttributes) ExpressRouteAuthorizationId() terra.StringValue {
	return terra.ReferenceAsString(vera.ref.Append("express_route_authorization_id"))
}

// ExpressRouteAuthorizationKey returns a reference to field express_route_authorization_key of azurerm_vmware_express_route_authorization.
func (vera vmwareExpressRouteAuthorizationAttributes) ExpressRouteAuthorizationKey() terra.StringValue {
	return terra.ReferenceAsString(vera.ref.Append("express_route_authorization_key"))
}

// Id returns a reference to field id of azurerm_vmware_express_route_authorization.
func (vera vmwareExpressRouteAuthorizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vera.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_vmware_express_route_authorization.
func (vera vmwareExpressRouteAuthorizationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vera.ref.Append("name"))
}

// PrivateCloudId returns a reference to field private_cloud_id of azurerm_vmware_express_route_authorization.
func (vera vmwareExpressRouteAuthorizationAttributes) PrivateCloudId() terra.StringValue {
	return terra.ReferenceAsString(vera.ref.Append("private_cloud_id"))
}

func (vera vmwareExpressRouteAuthorizationAttributes) Timeouts() vmwareexpressrouteauthorization.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vmwareexpressrouteauthorization.TimeoutsAttributes](vera.ref.Append("timeouts"))
}

type vmwareExpressRouteAuthorizationState struct {
	ExpressRouteAuthorizationId  string                                         `json:"express_route_authorization_id"`
	ExpressRouteAuthorizationKey string                                         `json:"express_route_authorization_key"`
	Id                           string                                         `json:"id"`
	Name                         string                                         `json:"name"`
	PrivateCloudId               string                                         `json:"private_cloud_id"`
	Timeouts                     *vmwareexpressrouteauthorization.TimeoutsState `json:"timeouts"`
}
