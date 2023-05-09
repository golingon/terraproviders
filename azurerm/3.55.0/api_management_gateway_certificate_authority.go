// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementgatewaycertificateauthority "github.com/golingon/terraproviders/azurerm/3.55.0/apimanagementgatewaycertificateauthority"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementGatewayCertificateAuthority creates a new instance of [ApiManagementGatewayCertificateAuthority].
func NewApiManagementGatewayCertificateAuthority(name string, args ApiManagementGatewayCertificateAuthorityArgs) *ApiManagementGatewayCertificateAuthority {
	return &ApiManagementGatewayCertificateAuthority{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementGatewayCertificateAuthority)(nil)

// ApiManagementGatewayCertificateAuthority represents the Terraform resource azurerm_api_management_gateway_certificate_authority.
type ApiManagementGatewayCertificateAuthority struct {
	Name      string
	Args      ApiManagementGatewayCertificateAuthorityArgs
	state     *apiManagementGatewayCertificateAuthorityState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementGatewayCertificateAuthority].
func (amgca *ApiManagementGatewayCertificateAuthority) Type() string {
	return "azurerm_api_management_gateway_certificate_authority"
}

// LocalName returns the local name for [ApiManagementGatewayCertificateAuthority].
func (amgca *ApiManagementGatewayCertificateAuthority) LocalName() string {
	return amgca.Name
}

// Configuration returns the configuration (args) for [ApiManagementGatewayCertificateAuthority].
func (amgca *ApiManagementGatewayCertificateAuthority) Configuration() interface{} {
	return amgca.Args
}

// DependOn is used for other resources to depend on [ApiManagementGatewayCertificateAuthority].
func (amgca *ApiManagementGatewayCertificateAuthority) DependOn() terra.Reference {
	return terra.ReferenceResource(amgca)
}

// Dependencies returns the list of resources [ApiManagementGatewayCertificateAuthority] depends_on.
func (amgca *ApiManagementGatewayCertificateAuthority) Dependencies() terra.Dependencies {
	return amgca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementGatewayCertificateAuthority].
func (amgca *ApiManagementGatewayCertificateAuthority) LifecycleManagement() *terra.Lifecycle {
	return amgca.Lifecycle
}

// Attributes returns the attributes for [ApiManagementGatewayCertificateAuthority].
func (amgca *ApiManagementGatewayCertificateAuthority) Attributes() apiManagementGatewayCertificateAuthorityAttributes {
	return apiManagementGatewayCertificateAuthorityAttributes{ref: terra.ReferenceResource(amgca)}
}

// ImportState imports the given attribute values into [ApiManagementGatewayCertificateAuthority]'s state.
func (amgca *ApiManagementGatewayCertificateAuthority) ImportState(av io.Reader) error {
	amgca.state = &apiManagementGatewayCertificateAuthorityState{}
	if err := json.NewDecoder(av).Decode(amgca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amgca.Type(), amgca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementGatewayCertificateAuthority] has state.
func (amgca *ApiManagementGatewayCertificateAuthority) State() (*apiManagementGatewayCertificateAuthorityState, bool) {
	return amgca.state, amgca.state != nil
}

// StateMust returns the state for [ApiManagementGatewayCertificateAuthority]. Panics if the state is nil.
func (amgca *ApiManagementGatewayCertificateAuthority) StateMust() *apiManagementGatewayCertificateAuthorityState {
	if amgca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amgca.Type(), amgca.LocalName()))
	}
	return amgca.state
}

// ApiManagementGatewayCertificateAuthorityArgs contains the configurations for azurerm_api_management_gateway_certificate_authority.
type ApiManagementGatewayCertificateAuthorityArgs struct {
	// ApiManagementId: string, required
	ApiManagementId terra.StringValue `hcl:"api_management_id,attr" validate:"required"`
	// CertificateName: string, required
	CertificateName terra.StringValue `hcl:"certificate_name,attr" validate:"required"`
	// GatewayName: string, required
	GatewayName terra.StringValue `hcl:"gateway_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsTrusted: bool, optional
	IsTrusted terra.BoolValue `hcl:"is_trusted,attr"`
	// Timeouts: optional
	Timeouts *apimanagementgatewaycertificateauthority.Timeouts `hcl:"timeouts,block"`
}
type apiManagementGatewayCertificateAuthorityAttributes struct {
	ref terra.Reference
}

// ApiManagementId returns a reference to field api_management_id of azurerm_api_management_gateway_certificate_authority.
func (amgca apiManagementGatewayCertificateAuthorityAttributes) ApiManagementId() terra.StringValue {
	return terra.ReferenceAsString(amgca.ref.Append("api_management_id"))
}

// CertificateName returns a reference to field certificate_name of azurerm_api_management_gateway_certificate_authority.
func (amgca apiManagementGatewayCertificateAuthorityAttributes) CertificateName() terra.StringValue {
	return terra.ReferenceAsString(amgca.ref.Append("certificate_name"))
}

// GatewayName returns a reference to field gateway_name of azurerm_api_management_gateway_certificate_authority.
func (amgca apiManagementGatewayCertificateAuthorityAttributes) GatewayName() terra.StringValue {
	return terra.ReferenceAsString(amgca.ref.Append("gateway_name"))
}

// Id returns a reference to field id of azurerm_api_management_gateway_certificate_authority.
func (amgca apiManagementGatewayCertificateAuthorityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amgca.ref.Append("id"))
}

// IsTrusted returns a reference to field is_trusted of azurerm_api_management_gateway_certificate_authority.
func (amgca apiManagementGatewayCertificateAuthorityAttributes) IsTrusted() terra.BoolValue {
	return terra.ReferenceAsBool(amgca.ref.Append("is_trusted"))
}

func (amgca apiManagementGatewayCertificateAuthorityAttributes) Timeouts() apimanagementgatewaycertificateauthority.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementgatewaycertificateauthority.TimeoutsAttributes](amgca.ref.Append("timeouts"))
}

type apiManagementGatewayCertificateAuthorityState struct {
	ApiManagementId string                                                  `json:"api_management_id"`
	CertificateName string                                                  `json:"certificate_name"`
	GatewayName     string                                                  `json:"gateway_name"`
	Id              string                                                  `json:"id"`
	IsTrusted       bool                                                    `json:"is_trusted"`
	Timeouts        *apimanagementgatewaycertificateauthority.TimeoutsState `json:"timeouts"`
}
