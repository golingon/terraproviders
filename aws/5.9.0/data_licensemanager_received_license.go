// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalicensemanagerreceivedlicense "github.com/golingon/terraproviders/aws/5.9.0/datalicensemanagerreceivedlicense"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLicensemanagerReceivedLicense creates a new instance of [DataLicensemanagerReceivedLicense].
func NewDataLicensemanagerReceivedLicense(name string, args DataLicensemanagerReceivedLicenseArgs) *DataLicensemanagerReceivedLicense {
	return &DataLicensemanagerReceivedLicense{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLicensemanagerReceivedLicense)(nil)

// DataLicensemanagerReceivedLicense represents the Terraform data resource aws_licensemanager_received_license.
type DataLicensemanagerReceivedLicense struct {
	Name string
	Args DataLicensemanagerReceivedLicenseArgs
}

// DataSource returns the Terraform object type for [DataLicensemanagerReceivedLicense].
func (lrl *DataLicensemanagerReceivedLicense) DataSource() string {
	return "aws_licensemanager_received_license"
}

// LocalName returns the local name for [DataLicensemanagerReceivedLicense].
func (lrl *DataLicensemanagerReceivedLicense) LocalName() string {
	return lrl.Name
}

// Configuration returns the configuration (args) for [DataLicensemanagerReceivedLicense].
func (lrl *DataLicensemanagerReceivedLicense) Configuration() interface{} {
	return lrl.Args
}

// Attributes returns the attributes for [DataLicensemanagerReceivedLicense].
func (lrl *DataLicensemanagerReceivedLicense) Attributes() dataLicensemanagerReceivedLicenseAttributes {
	return dataLicensemanagerReceivedLicenseAttributes{ref: terra.ReferenceDataResource(lrl)}
}

// DataLicensemanagerReceivedLicenseArgs contains the configurations for aws_licensemanager_received_license.
type DataLicensemanagerReceivedLicenseArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseArn: string, required
	LicenseArn terra.StringValue `hcl:"license_arn,attr" validate:"required"`
	// ConsumptionConfiguration: min=0
	ConsumptionConfiguration []datalicensemanagerreceivedlicense.ConsumptionConfiguration `hcl:"consumption_configuration,block" validate:"min=0"`
	// Entitlements: min=0
	Entitlements []datalicensemanagerreceivedlicense.Entitlements `hcl:"entitlements,block" validate:"min=0"`
	// Issuer: min=0
	Issuer []datalicensemanagerreceivedlicense.Issuer `hcl:"issuer,block" validate:"min=0"`
	// LicenseMetadata: min=0
	LicenseMetadata []datalicensemanagerreceivedlicense.LicenseMetadata `hcl:"license_metadata,block" validate:"min=0"`
	// ReceivedMetadata: min=0
	ReceivedMetadata []datalicensemanagerreceivedlicense.ReceivedMetadata `hcl:"received_metadata,block" validate:"min=0"`
	// Validity: min=0
	Validity []datalicensemanagerreceivedlicense.Validity `hcl:"validity,block" validate:"min=0"`
}
type dataLicensemanagerReceivedLicenseAttributes struct {
	ref terra.Reference
}

// Beneficiary returns a reference to field beneficiary of aws_licensemanager_received_license.
func (lrl dataLicensemanagerReceivedLicenseAttributes) Beneficiary() terra.StringValue {
	return terra.ReferenceAsString(lrl.ref.Append("beneficiary"))
}

// CreateTime returns a reference to field create_time of aws_licensemanager_received_license.
func (lrl dataLicensemanagerReceivedLicenseAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(lrl.ref.Append("create_time"))
}

// HomeRegion returns a reference to field home_region of aws_licensemanager_received_license.
func (lrl dataLicensemanagerReceivedLicenseAttributes) HomeRegion() terra.StringValue {
	return terra.ReferenceAsString(lrl.ref.Append("home_region"))
}

// Id returns a reference to field id of aws_licensemanager_received_license.
func (lrl dataLicensemanagerReceivedLicenseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lrl.ref.Append("id"))
}

// LicenseArn returns a reference to field license_arn of aws_licensemanager_received_license.
func (lrl dataLicensemanagerReceivedLicenseAttributes) LicenseArn() terra.StringValue {
	return terra.ReferenceAsString(lrl.ref.Append("license_arn"))
}

// LicenseName returns a reference to field license_name of aws_licensemanager_received_license.
func (lrl dataLicensemanagerReceivedLicenseAttributes) LicenseName() terra.StringValue {
	return terra.ReferenceAsString(lrl.ref.Append("license_name"))
}

// ProductName returns a reference to field product_name of aws_licensemanager_received_license.
func (lrl dataLicensemanagerReceivedLicenseAttributes) ProductName() terra.StringValue {
	return terra.ReferenceAsString(lrl.ref.Append("product_name"))
}

// ProductSku returns a reference to field product_sku of aws_licensemanager_received_license.
func (lrl dataLicensemanagerReceivedLicenseAttributes) ProductSku() terra.StringValue {
	return terra.ReferenceAsString(lrl.ref.Append("product_sku"))
}

// Status returns a reference to field status of aws_licensemanager_received_license.
func (lrl dataLicensemanagerReceivedLicenseAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(lrl.ref.Append("status"))
}

// Version returns a reference to field version of aws_licensemanager_received_license.
func (lrl dataLicensemanagerReceivedLicenseAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(lrl.ref.Append("version"))
}

func (lrl dataLicensemanagerReceivedLicenseAttributes) ConsumptionConfiguration() terra.ListValue[datalicensemanagerreceivedlicense.ConsumptionConfigurationAttributes] {
	return terra.ReferenceAsList[datalicensemanagerreceivedlicense.ConsumptionConfigurationAttributes](lrl.ref.Append("consumption_configuration"))
}

func (lrl dataLicensemanagerReceivedLicenseAttributes) Entitlements() terra.SetValue[datalicensemanagerreceivedlicense.EntitlementsAttributes] {
	return terra.ReferenceAsSet[datalicensemanagerreceivedlicense.EntitlementsAttributes](lrl.ref.Append("entitlements"))
}

func (lrl dataLicensemanagerReceivedLicenseAttributes) Issuer() terra.ListValue[datalicensemanagerreceivedlicense.IssuerAttributes] {
	return terra.ReferenceAsList[datalicensemanagerreceivedlicense.IssuerAttributes](lrl.ref.Append("issuer"))
}

func (lrl dataLicensemanagerReceivedLicenseAttributes) LicenseMetadata() terra.SetValue[datalicensemanagerreceivedlicense.LicenseMetadataAttributes] {
	return terra.ReferenceAsSet[datalicensemanagerreceivedlicense.LicenseMetadataAttributes](lrl.ref.Append("license_metadata"))
}

func (lrl dataLicensemanagerReceivedLicenseAttributes) ReceivedMetadata() terra.ListValue[datalicensemanagerreceivedlicense.ReceivedMetadataAttributes] {
	return terra.ReferenceAsList[datalicensemanagerreceivedlicense.ReceivedMetadataAttributes](lrl.ref.Append("received_metadata"))
}

func (lrl dataLicensemanagerReceivedLicenseAttributes) Validity() terra.ListValue[datalicensemanagerreceivedlicense.ValidityAttributes] {
	return terra.ReferenceAsList[datalicensemanagerreceivedlicense.ValidityAttributes](lrl.ref.Append("validity"))
}
