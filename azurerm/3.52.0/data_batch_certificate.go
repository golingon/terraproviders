// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	databatchcertificate "github.com/golingon/terraproviders/azurerm/3.52.0/databatchcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBatchCertificate creates a new instance of [DataBatchCertificate].
func NewDataBatchCertificate(name string, args DataBatchCertificateArgs) *DataBatchCertificate {
	return &DataBatchCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBatchCertificate)(nil)

// DataBatchCertificate represents the Terraform data resource azurerm_batch_certificate.
type DataBatchCertificate struct {
	Name string
	Args DataBatchCertificateArgs
}

// DataSource returns the Terraform object type for [DataBatchCertificate].
func (bc *DataBatchCertificate) DataSource() string {
	return "azurerm_batch_certificate"
}

// LocalName returns the local name for [DataBatchCertificate].
func (bc *DataBatchCertificate) LocalName() string {
	return bc.Name
}

// Configuration returns the configuration (args) for [DataBatchCertificate].
func (bc *DataBatchCertificate) Configuration() interface{} {
	return bc.Args
}

// Attributes returns the attributes for [DataBatchCertificate].
func (bc *DataBatchCertificate) Attributes() dataBatchCertificateAttributes {
	return dataBatchCertificateAttributes{ref: terra.ReferenceDataResource(bc)}
}

// DataBatchCertificateArgs contains the configurations for azurerm_batch_certificate.
type DataBatchCertificateArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *databatchcertificate.Timeouts `hcl:"timeouts,block"`
}
type dataBatchCertificateAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_batch_certificate.
func (bc dataBatchCertificateAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("account_name"))
}

// Format returns a reference to field format of azurerm_batch_certificate.
func (bc dataBatchCertificateAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("format"))
}

// Id returns a reference to field id of azurerm_batch_certificate.
func (bc dataBatchCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_batch_certificate.
func (bc dataBatchCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("name"))
}

// PublicData returns a reference to field public_data of azurerm_batch_certificate.
func (bc dataBatchCertificateAttributes) PublicData() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("public_data"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_batch_certificate.
func (bc dataBatchCertificateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("resource_group_name"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_batch_certificate.
func (bc dataBatchCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("thumbprint"))
}

// ThumbprintAlgorithm returns a reference to field thumbprint_algorithm of azurerm_batch_certificate.
func (bc dataBatchCertificateAttributes) ThumbprintAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("thumbprint_algorithm"))
}

func (bc dataBatchCertificateAttributes) Timeouts() databatchcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databatchcertificate.TimeoutsAttributes](bc.ref.Append("timeouts"))
}
