// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataappservicecertificateorder "github.com/golingon/terraproviders/azurerm/3.49.0/dataappservicecertificateorder"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataAppServiceCertificateOrder(name string, args DataAppServiceCertificateOrderArgs) *DataAppServiceCertificateOrder {
	return &DataAppServiceCertificateOrder{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppServiceCertificateOrder)(nil)

type DataAppServiceCertificateOrder struct {
	Name string
	Args DataAppServiceCertificateOrderArgs
}

func (asco *DataAppServiceCertificateOrder) DataSource() string {
	return "azurerm_app_service_certificate_order"
}

func (asco *DataAppServiceCertificateOrder) LocalName() string {
	return asco.Name
}

func (asco *DataAppServiceCertificateOrder) Configuration() interface{} {
	return asco.Args
}

func (asco *DataAppServiceCertificateOrder) Attributes() dataAppServiceCertificateOrderAttributes {
	return dataAppServiceCertificateOrderAttributes{ref: terra.ReferenceDataResource(asco)}
}

type DataAppServiceCertificateOrderArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Certificates: min=0
	Certificates []dataappservicecertificateorder.Certificates `hcl:"certificates,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataappservicecertificateorder.Timeouts `hcl:"timeouts,block"`
}
type dataAppServiceCertificateOrderAttributes struct {
	ref terra.Reference
}

func (asco dataAppServiceCertificateOrderAttributes) AppServiceCertificateNotRenewableReasons() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](asco.ref.Append("app_service_certificate_not_renewable_reasons"))
}

func (asco dataAppServiceCertificateOrderAttributes) AutoRenew() terra.BoolValue {
	return terra.ReferenceBool(asco.ref.Append("auto_renew"))
}

func (asco dataAppServiceCertificateOrderAttributes) Csr() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("csr"))
}

func (asco dataAppServiceCertificateOrderAttributes) DistinguishedName() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("distinguished_name"))
}

func (asco dataAppServiceCertificateOrderAttributes) DomainVerificationToken() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("domain_verification_token"))
}

func (asco dataAppServiceCertificateOrderAttributes) ExpirationTime() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("expiration_time"))
}

func (asco dataAppServiceCertificateOrderAttributes) Id() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("id"))
}

func (asco dataAppServiceCertificateOrderAttributes) IntermediateThumbprint() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("intermediate_thumbprint"))
}

func (asco dataAppServiceCertificateOrderAttributes) IsPrivateKeyExternal() terra.BoolValue {
	return terra.ReferenceBool(asco.ref.Append("is_private_key_external"))
}

func (asco dataAppServiceCertificateOrderAttributes) KeySize() terra.NumberValue {
	return terra.ReferenceNumber(asco.ref.Append("key_size"))
}

func (asco dataAppServiceCertificateOrderAttributes) Location() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("location"))
}

func (asco dataAppServiceCertificateOrderAttributes) Name() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("name"))
}

func (asco dataAppServiceCertificateOrderAttributes) ProductType() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("product_type"))
}

func (asco dataAppServiceCertificateOrderAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("resource_group_name"))
}

func (asco dataAppServiceCertificateOrderAttributes) RootThumbprint() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("root_thumbprint"))
}

func (asco dataAppServiceCertificateOrderAttributes) SignedCertificateThumbprint() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("signed_certificate_thumbprint"))
}

func (asco dataAppServiceCertificateOrderAttributes) Status() terra.StringValue {
	return terra.ReferenceString(asco.ref.Append("status"))
}

func (asco dataAppServiceCertificateOrderAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](asco.ref.Append("tags"))
}

func (asco dataAppServiceCertificateOrderAttributes) ValidityInYears() terra.NumberValue {
	return terra.ReferenceNumber(asco.ref.Append("validity_in_years"))
}

func (asco dataAppServiceCertificateOrderAttributes) Certificates() terra.ListValue[dataappservicecertificateorder.CertificatesAttributes] {
	return terra.ReferenceList[dataappservicecertificateorder.CertificatesAttributes](asco.ref.Append("certificates"))
}

func (asco dataAppServiceCertificateOrderAttributes) Timeouts() dataappservicecertificateorder.TimeoutsAttributes {
	return terra.ReferenceSingle[dataappservicecertificateorder.TimeoutsAttributes](asco.ref.Append("timeouts"))
}
