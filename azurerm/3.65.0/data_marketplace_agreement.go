// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamarketplaceagreement "github.com/golingon/terraproviders/azurerm/3.65.0/datamarketplaceagreement"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMarketplaceAgreement creates a new instance of [DataMarketplaceAgreement].
func NewDataMarketplaceAgreement(name string, args DataMarketplaceAgreementArgs) *DataMarketplaceAgreement {
	return &DataMarketplaceAgreement{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMarketplaceAgreement)(nil)

// DataMarketplaceAgreement represents the Terraform data resource azurerm_marketplace_agreement.
type DataMarketplaceAgreement struct {
	Name string
	Args DataMarketplaceAgreementArgs
}

// DataSource returns the Terraform object type for [DataMarketplaceAgreement].
func (ma *DataMarketplaceAgreement) DataSource() string {
	return "azurerm_marketplace_agreement"
}

// LocalName returns the local name for [DataMarketplaceAgreement].
func (ma *DataMarketplaceAgreement) LocalName() string {
	return ma.Name
}

// Configuration returns the configuration (args) for [DataMarketplaceAgreement].
func (ma *DataMarketplaceAgreement) Configuration() interface{} {
	return ma.Args
}

// Attributes returns the attributes for [DataMarketplaceAgreement].
func (ma *DataMarketplaceAgreement) Attributes() dataMarketplaceAgreementAttributes {
	return dataMarketplaceAgreementAttributes{ref: terra.ReferenceDataResource(ma)}
}

// DataMarketplaceAgreementArgs contains the configurations for azurerm_marketplace_agreement.
type DataMarketplaceAgreementArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Offer: string, required
	Offer terra.StringValue `hcl:"offer,attr" validate:"required"`
	// Plan: string, required
	Plan terra.StringValue `hcl:"plan,attr" validate:"required"`
	// Publisher: string, required
	Publisher terra.StringValue `hcl:"publisher,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datamarketplaceagreement.Timeouts `hcl:"timeouts,block"`
}
type dataMarketplaceAgreementAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_marketplace_agreement.
func (ma dataMarketplaceAgreementAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("id"))
}

// LicenseTextLink returns a reference to field license_text_link of azurerm_marketplace_agreement.
func (ma dataMarketplaceAgreementAttributes) LicenseTextLink() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("license_text_link"))
}

// Offer returns a reference to field offer of azurerm_marketplace_agreement.
func (ma dataMarketplaceAgreementAttributes) Offer() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("offer"))
}

// Plan returns a reference to field plan of azurerm_marketplace_agreement.
func (ma dataMarketplaceAgreementAttributes) Plan() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("plan"))
}

// PrivacyPolicyLink returns a reference to field privacy_policy_link of azurerm_marketplace_agreement.
func (ma dataMarketplaceAgreementAttributes) PrivacyPolicyLink() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("privacy_policy_link"))
}

// Publisher returns a reference to field publisher of azurerm_marketplace_agreement.
func (ma dataMarketplaceAgreementAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("publisher"))
}

func (ma dataMarketplaceAgreementAttributes) Timeouts() datamarketplaceagreement.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamarketplaceagreement.TimeoutsAttributes](ma.ref.Append("timeouts"))
}
