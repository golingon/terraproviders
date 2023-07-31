// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	containerappenvironmentcertificate "github.com/golingon/terraproviders/azurerm/3.67.0/containerappenvironmentcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerAppEnvironmentCertificate creates a new instance of [ContainerAppEnvironmentCertificate].
func NewContainerAppEnvironmentCertificate(name string, args ContainerAppEnvironmentCertificateArgs) *ContainerAppEnvironmentCertificate {
	return &ContainerAppEnvironmentCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAppEnvironmentCertificate)(nil)

// ContainerAppEnvironmentCertificate represents the Terraform resource azurerm_container_app_environment_certificate.
type ContainerAppEnvironmentCertificate struct {
	Name      string
	Args      ContainerAppEnvironmentCertificateArgs
	state     *containerAppEnvironmentCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAppEnvironmentCertificate].
func (caec *ContainerAppEnvironmentCertificate) Type() string {
	return "azurerm_container_app_environment_certificate"
}

// LocalName returns the local name for [ContainerAppEnvironmentCertificate].
func (caec *ContainerAppEnvironmentCertificate) LocalName() string {
	return caec.Name
}

// Configuration returns the configuration (args) for [ContainerAppEnvironmentCertificate].
func (caec *ContainerAppEnvironmentCertificate) Configuration() interface{} {
	return caec.Args
}

// DependOn is used for other resources to depend on [ContainerAppEnvironmentCertificate].
func (caec *ContainerAppEnvironmentCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(caec)
}

// Dependencies returns the list of resources [ContainerAppEnvironmentCertificate] depends_on.
func (caec *ContainerAppEnvironmentCertificate) Dependencies() terra.Dependencies {
	return caec.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAppEnvironmentCertificate].
func (caec *ContainerAppEnvironmentCertificate) LifecycleManagement() *terra.Lifecycle {
	return caec.Lifecycle
}

// Attributes returns the attributes for [ContainerAppEnvironmentCertificate].
func (caec *ContainerAppEnvironmentCertificate) Attributes() containerAppEnvironmentCertificateAttributes {
	return containerAppEnvironmentCertificateAttributes{ref: terra.ReferenceResource(caec)}
}

// ImportState imports the given attribute values into [ContainerAppEnvironmentCertificate]'s state.
func (caec *ContainerAppEnvironmentCertificate) ImportState(av io.Reader) error {
	caec.state = &containerAppEnvironmentCertificateState{}
	if err := json.NewDecoder(av).Decode(caec.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", caec.Type(), caec.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAppEnvironmentCertificate] has state.
func (caec *ContainerAppEnvironmentCertificate) State() (*containerAppEnvironmentCertificateState, bool) {
	return caec.state, caec.state != nil
}

// StateMust returns the state for [ContainerAppEnvironmentCertificate]. Panics if the state is nil.
func (caec *ContainerAppEnvironmentCertificate) StateMust() *containerAppEnvironmentCertificateState {
	if caec.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", caec.Type(), caec.LocalName()))
	}
	return caec.state
}

// ContainerAppEnvironmentCertificateArgs contains the configurations for azurerm_container_app_environment_certificate.
type ContainerAppEnvironmentCertificateArgs struct {
	// CertificateBlobBase64: string, required
	CertificateBlobBase64 terra.StringValue `hcl:"certificate_blob_base64,attr" validate:"required"`
	// CertificatePassword: string, required
	CertificatePassword terra.StringValue `hcl:"certificate_password,attr" validate:"required"`
	// ContainerAppEnvironmentId: string, required
	ContainerAppEnvironmentId terra.StringValue `hcl:"container_app_environment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *containerappenvironmentcertificate.Timeouts `hcl:"timeouts,block"`
}
type containerAppEnvironmentCertificateAttributes struct {
	ref terra.Reference
}

// CertificateBlobBase64 returns a reference to field certificate_blob_base64 of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) CertificateBlobBase64() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("certificate_blob_base64"))
}

// CertificatePassword returns a reference to field certificate_password of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) CertificatePassword() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("certificate_password"))
}

// ContainerAppEnvironmentId returns a reference to field container_app_environment_id of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) ContainerAppEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("container_app_environment_id"))
}

// ExpirationDate returns a reference to field expiration_date of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("expiration_date"))
}

// Id returns a reference to field id of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("id"))
}

// IssueDate returns a reference to field issue_date of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) IssueDate() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("issue_date"))
}

// Issuer returns a reference to field issuer of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("issuer"))
}

// Name returns a reference to field name of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("name"))
}

// SubjectName returns a reference to field subject_name of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) SubjectName() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("subject_name"))
}

// Tags returns a reference to field tags of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](caec.ref.Append("tags"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_container_app_environment_certificate.
func (caec containerAppEnvironmentCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("thumbprint"))
}

func (caec containerAppEnvironmentCertificateAttributes) Timeouts() containerappenvironmentcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerappenvironmentcertificate.TimeoutsAttributes](caec.ref.Append("timeouts"))
}

type containerAppEnvironmentCertificateState struct {
	CertificateBlobBase64     string                                            `json:"certificate_blob_base64"`
	CertificatePassword       string                                            `json:"certificate_password"`
	ContainerAppEnvironmentId string                                            `json:"container_app_environment_id"`
	ExpirationDate            string                                            `json:"expiration_date"`
	Id                        string                                            `json:"id"`
	IssueDate                 string                                            `json:"issue_date"`
	Issuer                    string                                            `json:"issuer"`
	Name                      string                                            `json:"name"`
	SubjectName               string                                            `json:"subject_name"`
	Tags                      map[string]string                                 `json:"tags"`
	Thumbprint                string                                            `json:"thumbprint"`
	Timeouts                  *containerappenvironmentcertificate.TimeoutsState `json:"timeouts"`
}
