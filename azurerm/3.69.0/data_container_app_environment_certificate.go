// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacontainerappenvironmentcertificate "github.com/golingon/terraproviders/azurerm/3.69.0/datacontainerappenvironmentcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataContainerAppEnvironmentCertificate creates a new instance of [DataContainerAppEnvironmentCertificate].
func NewDataContainerAppEnvironmentCertificate(name string, args DataContainerAppEnvironmentCertificateArgs) *DataContainerAppEnvironmentCertificate {
	return &DataContainerAppEnvironmentCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerAppEnvironmentCertificate)(nil)

// DataContainerAppEnvironmentCertificate represents the Terraform data resource azurerm_container_app_environment_certificate.
type DataContainerAppEnvironmentCertificate struct {
	Name string
	Args DataContainerAppEnvironmentCertificateArgs
}

// DataSource returns the Terraform object type for [DataContainerAppEnvironmentCertificate].
func (caec *DataContainerAppEnvironmentCertificate) DataSource() string {
	return "azurerm_container_app_environment_certificate"
}

// LocalName returns the local name for [DataContainerAppEnvironmentCertificate].
func (caec *DataContainerAppEnvironmentCertificate) LocalName() string {
	return caec.Name
}

// Configuration returns the configuration (args) for [DataContainerAppEnvironmentCertificate].
func (caec *DataContainerAppEnvironmentCertificate) Configuration() interface{} {
	return caec.Args
}

// Attributes returns the attributes for [DataContainerAppEnvironmentCertificate].
func (caec *DataContainerAppEnvironmentCertificate) Attributes() dataContainerAppEnvironmentCertificateAttributes {
	return dataContainerAppEnvironmentCertificateAttributes{ref: terra.ReferenceDataResource(caec)}
}

// DataContainerAppEnvironmentCertificateArgs contains the configurations for azurerm_container_app_environment_certificate.
type DataContainerAppEnvironmentCertificateArgs struct {
	// ContainerAppEnvironmentId: string, required
	ContainerAppEnvironmentId terra.StringValue `hcl:"container_app_environment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datacontainerappenvironmentcertificate.Timeouts `hcl:"timeouts,block"`
}
type dataContainerAppEnvironmentCertificateAttributes struct {
	ref terra.Reference
}

// ContainerAppEnvironmentId returns a reference to field container_app_environment_id of azurerm_container_app_environment_certificate.
func (caec dataContainerAppEnvironmentCertificateAttributes) ContainerAppEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("container_app_environment_id"))
}

// ExpirationDate returns a reference to field expiration_date of azurerm_container_app_environment_certificate.
func (caec dataContainerAppEnvironmentCertificateAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("expiration_date"))
}

// Id returns a reference to field id of azurerm_container_app_environment_certificate.
func (caec dataContainerAppEnvironmentCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("id"))
}

// IssueDate returns a reference to field issue_date of azurerm_container_app_environment_certificate.
func (caec dataContainerAppEnvironmentCertificateAttributes) IssueDate() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("issue_date"))
}

// Issuer returns a reference to field issuer of azurerm_container_app_environment_certificate.
func (caec dataContainerAppEnvironmentCertificateAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("issuer"))
}

// Name returns a reference to field name of azurerm_container_app_environment_certificate.
func (caec dataContainerAppEnvironmentCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("name"))
}

// SubjectName returns a reference to field subject_name of azurerm_container_app_environment_certificate.
func (caec dataContainerAppEnvironmentCertificateAttributes) SubjectName() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("subject_name"))
}

// Tags returns a reference to field tags of azurerm_container_app_environment_certificate.
func (caec dataContainerAppEnvironmentCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](caec.ref.Append("tags"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_container_app_environment_certificate.
func (caec dataContainerAppEnvironmentCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(caec.ref.Append("thumbprint"))
}

func (caec dataContainerAppEnvironmentCertificateAttributes) Timeouts() datacontainerappenvironmentcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacontainerappenvironmentcertificate.TimeoutsAttributes](caec.ref.Append("timeouts"))
}
