// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementcertificate "github.com/golingon/terraproviders/azurerm/3.65.0/apimanagementcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementCertificate creates a new instance of [ApiManagementCertificate].
func NewApiManagementCertificate(name string, args ApiManagementCertificateArgs) *ApiManagementCertificate {
	return &ApiManagementCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementCertificate)(nil)

// ApiManagementCertificate represents the Terraform resource azurerm_api_management_certificate.
type ApiManagementCertificate struct {
	Name      string
	Args      ApiManagementCertificateArgs
	state     *apiManagementCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementCertificate].
func (amc *ApiManagementCertificate) Type() string {
	return "azurerm_api_management_certificate"
}

// LocalName returns the local name for [ApiManagementCertificate].
func (amc *ApiManagementCertificate) LocalName() string {
	return amc.Name
}

// Configuration returns the configuration (args) for [ApiManagementCertificate].
func (amc *ApiManagementCertificate) Configuration() interface{} {
	return amc.Args
}

// DependOn is used for other resources to depend on [ApiManagementCertificate].
func (amc *ApiManagementCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(amc)
}

// Dependencies returns the list of resources [ApiManagementCertificate] depends_on.
func (amc *ApiManagementCertificate) Dependencies() terra.Dependencies {
	return amc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementCertificate].
func (amc *ApiManagementCertificate) LifecycleManagement() *terra.Lifecycle {
	return amc.Lifecycle
}

// Attributes returns the attributes for [ApiManagementCertificate].
func (amc *ApiManagementCertificate) Attributes() apiManagementCertificateAttributes {
	return apiManagementCertificateAttributes{ref: terra.ReferenceResource(amc)}
}

// ImportState imports the given attribute values into [ApiManagementCertificate]'s state.
func (amc *ApiManagementCertificate) ImportState(av io.Reader) error {
	amc.state = &apiManagementCertificateState{}
	if err := json.NewDecoder(av).Decode(amc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amc.Type(), amc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementCertificate] has state.
func (amc *ApiManagementCertificate) State() (*apiManagementCertificateState, bool) {
	return amc.state, amc.state != nil
}

// StateMust returns the state for [ApiManagementCertificate]. Panics if the state is nil.
func (amc *ApiManagementCertificate) StateMust() *apiManagementCertificateState {
	if amc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amc.Type(), amc.LocalName()))
	}
	return amc.state
}

// ApiManagementCertificateArgs contains the configurations for azurerm_api_management_certificate.
type ApiManagementCertificateArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Data: string, optional
	Data terra.StringValue `hcl:"data,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultIdentityClientId: string, optional
	KeyVaultIdentityClientId terra.StringValue `hcl:"key_vault_identity_client_id,attr"`
	// KeyVaultSecretId: string, optional
	KeyVaultSecretId terra.StringValue `hcl:"key_vault_secret_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementcertificate.Timeouts `hcl:"timeouts,block"`
}
type apiManagementCertificateAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("api_management_name"))
}

// Data returns a reference to field data of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("data"))
}

// Expiration returns a reference to field expiration of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) Expiration() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("expiration"))
}

// Id returns a reference to field id of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("id"))
}

// KeyVaultIdentityClientId returns a reference to field key_vault_identity_client_id of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) KeyVaultIdentityClientId() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("key_vault_identity_client_id"))
}

// KeyVaultSecretId returns a reference to field key_vault_secret_id of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) KeyVaultSecretId() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("key_vault_secret_id"))
}

// Name returns a reference to field name of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("name"))
}

// Password returns a reference to field password of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("password"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("resource_group_name"))
}

// Subject returns a reference to field subject of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("subject"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_api_management_certificate.
func (amc apiManagementCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("thumbprint"))
}

func (amc apiManagementCertificateAttributes) Timeouts() apimanagementcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementcertificate.TimeoutsAttributes](amc.ref.Append("timeouts"))
}

type apiManagementCertificateState struct {
	ApiManagementName        string                                  `json:"api_management_name"`
	Data                     string                                  `json:"data"`
	Expiration               string                                  `json:"expiration"`
	Id                       string                                  `json:"id"`
	KeyVaultIdentityClientId string                                  `json:"key_vault_identity_client_id"`
	KeyVaultSecretId         string                                  `json:"key_vault_secret_id"`
	Name                     string                                  `json:"name"`
	Password                 string                                  `json:"password"`
	ResourceGroupName        string                                  `json:"resource_group_name"`
	Subject                  string                                  `json:"subject"`
	Thumbprint               string                                  `json:"thumbprint"`
	Timeouts                 *apimanagementcertificate.TimeoutsState `json:"timeouts"`
}
