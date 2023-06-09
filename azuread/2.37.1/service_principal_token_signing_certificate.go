// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	serviceprincipaltokensigningcertificate "github.com/golingon/terraproviders/azuread/2.37.1/serviceprincipaltokensigningcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicePrincipalTokenSigningCertificate creates a new instance of [ServicePrincipalTokenSigningCertificate].
func NewServicePrincipalTokenSigningCertificate(name string, args ServicePrincipalTokenSigningCertificateArgs) *ServicePrincipalTokenSigningCertificate {
	return &ServicePrincipalTokenSigningCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicePrincipalTokenSigningCertificate)(nil)

// ServicePrincipalTokenSigningCertificate represents the Terraform resource azuread_service_principal_token_signing_certificate.
type ServicePrincipalTokenSigningCertificate struct {
	Name      string
	Args      ServicePrincipalTokenSigningCertificateArgs
	state     *servicePrincipalTokenSigningCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicePrincipalTokenSigningCertificate].
func (sptsc *ServicePrincipalTokenSigningCertificate) Type() string {
	return "azuread_service_principal_token_signing_certificate"
}

// LocalName returns the local name for [ServicePrincipalTokenSigningCertificate].
func (sptsc *ServicePrincipalTokenSigningCertificate) LocalName() string {
	return sptsc.Name
}

// Configuration returns the configuration (args) for [ServicePrincipalTokenSigningCertificate].
func (sptsc *ServicePrincipalTokenSigningCertificate) Configuration() interface{} {
	return sptsc.Args
}

// DependOn is used for other resources to depend on [ServicePrincipalTokenSigningCertificate].
func (sptsc *ServicePrincipalTokenSigningCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(sptsc)
}

// Dependencies returns the list of resources [ServicePrincipalTokenSigningCertificate] depends_on.
func (sptsc *ServicePrincipalTokenSigningCertificate) Dependencies() terra.Dependencies {
	return sptsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicePrincipalTokenSigningCertificate].
func (sptsc *ServicePrincipalTokenSigningCertificate) LifecycleManagement() *terra.Lifecycle {
	return sptsc.Lifecycle
}

// Attributes returns the attributes for [ServicePrincipalTokenSigningCertificate].
func (sptsc *ServicePrincipalTokenSigningCertificate) Attributes() servicePrincipalTokenSigningCertificateAttributes {
	return servicePrincipalTokenSigningCertificateAttributes{ref: terra.ReferenceResource(sptsc)}
}

// ImportState imports the given attribute values into [ServicePrincipalTokenSigningCertificate]'s state.
func (sptsc *ServicePrincipalTokenSigningCertificate) ImportState(av io.Reader) error {
	sptsc.state = &servicePrincipalTokenSigningCertificateState{}
	if err := json.NewDecoder(av).Decode(sptsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sptsc.Type(), sptsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicePrincipalTokenSigningCertificate] has state.
func (sptsc *ServicePrincipalTokenSigningCertificate) State() (*servicePrincipalTokenSigningCertificateState, bool) {
	return sptsc.state, sptsc.state != nil
}

// StateMust returns the state for [ServicePrincipalTokenSigningCertificate]. Panics if the state is nil.
func (sptsc *ServicePrincipalTokenSigningCertificate) StateMust() *servicePrincipalTokenSigningCertificateState {
	if sptsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sptsc.Type(), sptsc.LocalName()))
	}
	return sptsc.state
}

// ServicePrincipalTokenSigningCertificateArgs contains the configurations for azuread_service_principal_token_signing_certificate.
type ServicePrincipalTokenSigningCertificateArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EndDate: string, optional
	EndDate terra.StringValue `hcl:"end_date,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServicePrincipalId: string, required
	ServicePrincipalId terra.StringValue `hcl:"service_principal_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *serviceprincipaltokensigningcertificate.Timeouts `hcl:"timeouts,block"`
}
type servicePrincipalTokenSigningCertificateAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azuread_service_principal_token_signing_certificate.
func (sptsc servicePrincipalTokenSigningCertificateAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sptsc.ref.Append("display_name"))
}

// EndDate returns a reference to field end_date of azuread_service_principal_token_signing_certificate.
func (sptsc servicePrincipalTokenSigningCertificateAttributes) EndDate() terra.StringValue {
	return terra.ReferenceAsString(sptsc.ref.Append("end_date"))
}

// Id returns a reference to field id of azuread_service_principal_token_signing_certificate.
func (sptsc servicePrincipalTokenSigningCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sptsc.ref.Append("id"))
}

// KeyId returns a reference to field key_id of azuread_service_principal_token_signing_certificate.
func (sptsc servicePrincipalTokenSigningCertificateAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(sptsc.ref.Append("key_id"))
}

// ServicePrincipalId returns a reference to field service_principal_id of azuread_service_principal_token_signing_certificate.
func (sptsc servicePrincipalTokenSigningCertificateAttributes) ServicePrincipalId() terra.StringValue {
	return terra.ReferenceAsString(sptsc.ref.Append("service_principal_id"))
}

// StartDate returns a reference to field start_date of azuread_service_principal_token_signing_certificate.
func (sptsc servicePrincipalTokenSigningCertificateAttributes) StartDate() terra.StringValue {
	return terra.ReferenceAsString(sptsc.ref.Append("start_date"))
}

// Thumbprint returns a reference to field thumbprint of azuread_service_principal_token_signing_certificate.
func (sptsc servicePrincipalTokenSigningCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(sptsc.ref.Append("thumbprint"))
}

// Value returns a reference to field value of azuread_service_principal_token_signing_certificate.
func (sptsc servicePrincipalTokenSigningCertificateAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(sptsc.ref.Append("value"))
}

func (sptsc servicePrincipalTokenSigningCertificateAttributes) Timeouts() serviceprincipaltokensigningcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[serviceprincipaltokensigningcertificate.TimeoutsAttributes](sptsc.ref.Append("timeouts"))
}

type servicePrincipalTokenSigningCertificateState struct {
	DisplayName        string                                                 `json:"display_name"`
	EndDate            string                                                 `json:"end_date"`
	Id                 string                                                 `json:"id"`
	KeyId              string                                                 `json:"key_id"`
	ServicePrincipalId string                                                 `json:"service_principal_id"`
	StartDate          string                                                 `json:"start_date"`
	Thumbprint         string                                                 `json:"thumbprint"`
	Value              string                                                 `json:"value"`
	Timeouts           *serviceprincipaltokensigningcertificate.TimeoutsState `json:"timeouts"`
}
