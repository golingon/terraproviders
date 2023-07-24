// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementgateway "github.com/golingon/terraproviders/azurerm/3.66.0/apimanagementgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementGateway creates a new instance of [ApiManagementGateway].
func NewApiManagementGateway(name string, args ApiManagementGatewayArgs) *ApiManagementGateway {
	return &ApiManagementGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementGateway)(nil)

// ApiManagementGateway represents the Terraform resource azurerm_api_management_gateway.
type ApiManagementGateway struct {
	Name      string
	Args      ApiManagementGatewayArgs
	state     *apiManagementGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementGateway].
func (amg *ApiManagementGateway) Type() string {
	return "azurerm_api_management_gateway"
}

// LocalName returns the local name for [ApiManagementGateway].
func (amg *ApiManagementGateway) LocalName() string {
	return amg.Name
}

// Configuration returns the configuration (args) for [ApiManagementGateway].
func (amg *ApiManagementGateway) Configuration() interface{} {
	return amg.Args
}

// DependOn is used for other resources to depend on [ApiManagementGateway].
func (amg *ApiManagementGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(amg)
}

// Dependencies returns the list of resources [ApiManagementGateway] depends_on.
func (amg *ApiManagementGateway) Dependencies() terra.Dependencies {
	return amg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementGateway].
func (amg *ApiManagementGateway) LifecycleManagement() *terra.Lifecycle {
	return amg.Lifecycle
}

// Attributes returns the attributes for [ApiManagementGateway].
func (amg *ApiManagementGateway) Attributes() apiManagementGatewayAttributes {
	return apiManagementGatewayAttributes{ref: terra.ReferenceResource(amg)}
}

// ImportState imports the given attribute values into [ApiManagementGateway]'s state.
func (amg *ApiManagementGateway) ImportState(av io.Reader) error {
	amg.state = &apiManagementGatewayState{}
	if err := json.NewDecoder(av).Decode(amg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amg.Type(), amg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementGateway] has state.
func (amg *ApiManagementGateway) State() (*apiManagementGatewayState, bool) {
	return amg.state, amg.state != nil
}

// StateMust returns the state for [ApiManagementGateway]. Panics if the state is nil.
func (amg *ApiManagementGateway) StateMust() *apiManagementGatewayState {
	if amg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amg.Type(), amg.LocalName()))
	}
	return amg.state
}

// ApiManagementGatewayArgs contains the configurations for azurerm_api_management_gateway.
type ApiManagementGatewayArgs struct {
	// ApiManagementId: string, required
	ApiManagementId terra.StringValue `hcl:"api_management_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// LocationData: required
	LocationData *apimanagementgateway.LocationData `hcl:"location_data,block" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementgateway.Timeouts `hcl:"timeouts,block"`
}
type apiManagementGatewayAttributes struct {
	ref terra.Reference
}

// ApiManagementId returns a reference to field api_management_id of azurerm_api_management_gateway.
func (amg apiManagementGatewayAttributes) ApiManagementId() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("api_management_id"))
}

// Description returns a reference to field description of azurerm_api_management_gateway.
func (amg apiManagementGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_api_management_gateway.
func (amg apiManagementGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_gateway.
func (amg apiManagementGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("name"))
}

func (amg apiManagementGatewayAttributes) LocationData() terra.ListValue[apimanagementgateway.LocationDataAttributes] {
	return terra.ReferenceAsList[apimanagementgateway.LocationDataAttributes](amg.ref.Append("location_data"))
}

func (amg apiManagementGatewayAttributes) Timeouts() apimanagementgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementgateway.TimeoutsAttributes](amg.ref.Append("timeouts"))
}

type apiManagementGatewayState struct {
	ApiManagementId string                                   `json:"api_management_id"`
	Description     string                                   `json:"description"`
	Id              string                                   `json:"id"`
	Name            string                                   `json:"name"`
	LocationData    []apimanagementgateway.LocationDataState `json:"location_data"`
	Timeouts        *apimanagementgateway.TimeoutsState      `json:"timeouts"`
}
