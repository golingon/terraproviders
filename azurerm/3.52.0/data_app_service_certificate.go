// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataappservicecertificate "github.com/golingon/terraproviders/azurerm/3.52.0/dataappservicecertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppServiceCertificate creates a new instance of [DataAppServiceCertificate].
func NewDataAppServiceCertificate(name string, args DataAppServiceCertificateArgs) *DataAppServiceCertificate {
	return &DataAppServiceCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppServiceCertificate)(nil)

// DataAppServiceCertificate represents the Terraform data resource azurerm_app_service_certificate.
type DataAppServiceCertificate struct {
	Name string
	Args DataAppServiceCertificateArgs
}

// DataSource returns the Terraform object type for [DataAppServiceCertificate].
func (asc *DataAppServiceCertificate) DataSource() string {
	return "azurerm_app_service_certificate"
}

// LocalName returns the local name for [DataAppServiceCertificate].
func (asc *DataAppServiceCertificate) LocalName() string {
	return asc.Name
}

// Configuration returns the configuration (args) for [DataAppServiceCertificate].
func (asc *DataAppServiceCertificate) Configuration() interface{} {
	return asc.Args
}

// Attributes returns the attributes for [DataAppServiceCertificate].
func (asc *DataAppServiceCertificate) Attributes() dataAppServiceCertificateAttributes {
	return dataAppServiceCertificateAttributes{ref: terra.ReferenceDataResource(asc)}
}

// DataAppServiceCertificateArgs contains the configurations for azurerm_app_service_certificate.
type DataAppServiceCertificateArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *dataappservicecertificate.Timeouts `hcl:"timeouts,block"`
}
type dataAppServiceCertificateAttributes struct {
	ref terra.Reference
}

// ExpirationDate returns a reference to field expiration_date of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("expiration_date"))
}

// FriendlyName returns a reference to field friendly_name of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("friendly_name"))
}

// HostNames returns a reference to field host_names of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) HostNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asc.ref.Append("host_names"))
}

// Id returns a reference to field id of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("id"))
}

// IssueDate returns a reference to field issue_date of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) IssueDate() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("issue_date"))
}

// Issuer returns a reference to field issuer of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("issuer"))
}

// Location returns a reference to field location of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("resource_group_name"))
}

// SubjectName returns a reference to field subject_name of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) SubjectName() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("subject_name"))
}

// Tags returns a reference to field tags of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asc.ref.Append("tags"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_app_service_certificate.
func (asc dataAppServiceCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("thumbprint"))
}

func (asc dataAppServiceCertificateAttributes) Timeouts() dataappservicecertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataappservicecertificate.TimeoutsAttributes](asc.ref.Append("timeouts"))
}
