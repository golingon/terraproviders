// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementgatewayapi "github.com/golingon/terraproviders/azurerm/3.67.0/apimanagementgatewayapi"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementGatewayApi creates a new instance of [ApiManagementGatewayApi].
func NewApiManagementGatewayApi(name string, args ApiManagementGatewayApiArgs) *ApiManagementGatewayApi {
	return &ApiManagementGatewayApi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementGatewayApi)(nil)

// ApiManagementGatewayApi represents the Terraform resource azurerm_api_management_gateway_api.
type ApiManagementGatewayApi struct {
	Name      string
	Args      ApiManagementGatewayApiArgs
	state     *apiManagementGatewayApiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementGatewayApi].
func (amga *ApiManagementGatewayApi) Type() string {
	return "azurerm_api_management_gateway_api"
}

// LocalName returns the local name for [ApiManagementGatewayApi].
func (amga *ApiManagementGatewayApi) LocalName() string {
	return amga.Name
}

// Configuration returns the configuration (args) for [ApiManagementGatewayApi].
func (amga *ApiManagementGatewayApi) Configuration() interface{} {
	return amga.Args
}

// DependOn is used for other resources to depend on [ApiManagementGatewayApi].
func (amga *ApiManagementGatewayApi) DependOn() terra.Reference {
	return terra.ReferenceResource(amga)
}

// Dependencies returns the list of resources [ApiManagementGatewayApi] depends_on.
func (amga *ApiManagementGatewayApi) Dependencies() terra.Dependencies {
	return amga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementGatewayApi].
func (amga *ApiManagementGatewayApi) LifecycleManagement() *terra.Lifecycle {
	return amga.Lifecycle
}

// Attributes returns the attributes for [ApiManagementGatewayApi].
func (amga *ApiManagementGatewayApi) Attributes() apiManagementGatewayApiAttributes {
	return apiManagementGatewayApiAttributes{ref: terra.ReferenceResource(amga)}
}

// ImportState imports the given attribute values into [ApiManagementGatewayApi]'s state.
func (amga *ApiManagementGatewayApi) ImportState(av io.Reader) error {
	amga.state = &apiManagementGatewayApiState{}
	if err := json.NewDecoder(av).Decode(amga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amga.Type(), amga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementGatewayApi] has state.
func (amga *ApiManagementGatewayApi) State() (*apiManagementGatewayApiState, bool) {
	return amga.state, amga.state != nil
}

// StateMust returns the state for [ApiManagementGatewayApi]. Panics if the state is nil.
func (amga *ApiManagementGatewayApi) StateMust() *apiManagementGatewayApiState {
	if amga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amga.Type(), amga.LocalName()))
	}
	return amga.state
}

// ApiManagementGatewayApiArgs contains the configurations for azurerm_api_management_gateway_api.
type ApiManagementGatewayApiArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// GatewayId: string, required
	GatewayId terra.StringValue `hcl:"gateway_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *apimanagementgatewayapi.Timeouts `hcl:"timeouts,block"`
}
type apiManagementGatewayApiAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of azurerm_api_management_gateway_api.
func (amga apiManagementGatewayApiAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(amga.ref.Append("api_id"))
}

// GatewayId returns a reference to field gateway_id of azurerm_api_management_gateway_api.
func (amga apiManagementGatewayApiAttributes) GatewayId() terra.StringValue {
	return terra.ReferenceAsString(amga.ref.Append("gateway_id"))
}

// Id returns a reference to field id of azurerm_api_management_gateway_api.
func (amga apiManagementGatewayApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amga.ref.Append("id"))
}

func (amga apiManagementGatewayApiAttributes) Timeouts() apimanagementgatewayapi.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementgatewayapi.TimeoutsAttributes](amga.ref.Append("timeouts"))
}

type apiManagementGatewayApiState struct {
	ApiId     string                                 `json:"api_id"`
	GatewayId string                                 `json:"gateway_id"`
	Id        string                                 `json:"id"`
	Timeouts  *apimanagementgatewayapi.TimeoutsState `json:"timeouts"`
}
