// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appservicemanagedcertificate "github.com/golingon/terraproviders/azurerm/3.64.0/appservicemanagedcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceManagedCertificate creates a new instance of [AppServiceManagedCertificate].
func NewAppServiceManagedCertificate(name string, args AppServiceManagedCertificateArgs) *AppServiceManagedCertificate {
	return &AppServiceManagedCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceManagedCertificate)(nil)

// AppServiceManagedCertificate represents the Terraform resource azurerm_app_service_managed_certificate.
type AppServiceManagedCertificate struct {
	Name      string
	Args      AppServiceManagedCertificateArgs
	state     *appServiceManagedCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceManagedCertificate].
func (asmc *AppServiceManagedCertificate) Type() string {
	return "azurerm_app_service_managed_certificate"
}

// LocalName returns the local name for [AppServiceManagedCertificate].
func (asmc *AppServiceManagedCertificate) LocalName() string {
	return asmc.Name
}

// Configuration returns the configuration (args) for [AppServiceManagedCertificate].
func (asmc *AppServiceManagedCertificate) Configuration() interface{} {
	return asmc.Args
}

// DependOn is used for other resources to depend on [AppServiceManagedCertificate].
func (asmc *AppServiceManagedCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(asmc)
}

// Dependencies returns the list of resources [AppServiceManagedCertificate] depends_on.
func (asmc *AppServiceManagedCertificate) Dependencies() terra.Dependencies {
	return asmc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceManagedCertificate].
func (asmc *AppServiceManagedCertificate) LifecycleManagement() *terra.Lifecycle {
	return asmc.Lifecycle
}

// Attributes returns the attributes for [AppServiceManagedCertificate].
func (asmc *AppServiceManagedCertificate) Attributes() appServiceManagedCertificateAttributes {
	return appServiceManagedCertificateAttributes{ref: terra.ReferenceResource(asmc)}
}

// ImportState imports the given attribute values into [AppServiceManagedCertificate]'s state.
func (asmc *AppServiceManagedCertificate) ImportState(av io.Reader) error {
	asmc.state = &appServiceManagedCertificateState{}
	if err := json.NewDecoder(av).Decode(asmc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asmc.Type(), asmc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceManagedCertificate] has state.
func (asmc *AppServiceManagedCertificate) State() (*appServiceManagedCertificateState, bool) {
	return asmc.state, asmc.state != nil
}

// StateMust returns the state for [AppServiceManagedCertificate]. Panics if the state is nil.
func (asmc *AppServiceManagedCertificate) StateMust() *appServiceManagedCertificateState {
	if asmc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asmc.Type(), asmc.LocalName()))
	}
	return asmc.state
}

// AppServiceManagedCertificateArgs contains the configurations for azurerm_app_service_managed_certificate.
type AppServiceManagedCertificateArgs struct {
	// CustomHostnameBindingId: string, required
	CustomHostnameBindingId terra.StringValue `hcl:"custom_hostname_binding_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *appservicemanagedcertificate.Timeouts `hcl:"timeouts,block"`
}
type appServiceManagedCertificateAttributes struct {
	ref terra.Reference
}

// CanonicalName returns a reference to field canonical_name of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) CanonicalName() terra.StringValue {
	return terra.ReferenceAsString(asmc.ref.Append("canonical_name"))
}

// CustomHostnameBindingId returns a reference to field custom_hostname_binding_id of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) CustomHostnameBindingId() terra.StringValue {
	return terra.ReferenceAsString(asmc.ref.Append("custom_hostname_binding_id"))
}

// ExpirationDate returns a reference to field expiration_date of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(asmc.ref.Append("expiration_date"))
}

// FriendlyName returns a reference to field friendly_name of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(asmc.ref.Append("friendly_name"))
}

// HostNames returns a reference to field host_names of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) HostNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asmc.ref.Append("host_names"))
}

// Id returns a reference to field id of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asmc.ref.Append("id"))
}

// IssueDate returns a reference to field issue_date of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) IssueDate() terra.StringValue {
	return terra.ReferenceAsString(asmc.ref.Append("issue_date"))
}

// Issuer returns a reference to field issuer of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(asmc.ref.Append("issuer"))
}

// SubjectName returns a reference to field subject_name of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) SubjectName() terra.StringValue {
	return terra.ReferenceAsString(asmc.ref.Append("subject_name"))
}

// Tags returns a reference to field tags of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asmc.ref.Append("tags"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_app_service_managed_certificate.
func (asmc appServiceManagedCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(asmc.ref.Append("thumbprint"))
}

func (asmc appServiceManagedCertificateAttributes) Timeouts() appservicemanagedcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appservicemanagedcertificate.TimeoutsAttributes](asmc.ref.Append("timeouts"))
}

type appServiceManagedCertificateState struct {
	CanonicalName           string                                      `json:"canonical_name"`
	CustomHostnameBindingId string                                      `json:"custom_hostname_binding_id"`
	ExpirationDate          string                                      `json:"expiration_date"`
	FriendlyName            string                                      `json:"friendly_name"`
	HostNames               []string                                    `json:"host_names"`
	Id                      string                                      `json:"id"`
	IssueDate               string                                      `json:"issue_date"`
	Issuer                  string                                      `json:"issuer"`
	SubjectName             string                                      `json:"subject_name"`
	Tags                    map[string]string                           `json:"tags"`
	Thumbprint              string                                      `json:"thumbprint"`
	Timeouts                *appservicemanagedcertificate.TimeoutsState `json:"timeouts"`
}