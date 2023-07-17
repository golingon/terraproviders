// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataconfidentialledger "github.com/golingon/terraproviders/azurerm/3.65.0/dataconfidentialledger"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataConfidentialLedger creates a new instance of [DataConfidentialLedger].
func NewDataConfidentialLedger(name string, args DataConfidentialLedgerArgs) *DataConfidentialLedger {
	return &DataConfidentialLedger{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConfidentialLedger)(nil)

// DataConfidentialLedger represents the Terraform data resource azurerm_confidential_ledger.
type DataConfidentialLedger struct {
	Name string
	Args DataConfidentialLedgerArgs
}

// DataSource returns the Terraform object type for [DataConfidentialLedger].
func (cl *DataConfidentialLedger) DataSource() string {
	return "azurerm_confidential_ledger"
}

// LocalName returns the local name for [DataConfidentialLedger].
func (cl *DataConfidentialLedger) LocalName() string {
	return cl.Name
}

// Configuration returns the configuration (args) for [DataConfidentialLedger].
func (cl *DataConfidentialLedger) Configuration() interface{} {
	return cl.Args
}

// Attributes returns the attributes for [DataConfidentialLedger].
func (cl *DataConfidentialLedger) Attributes() dataConfidentialLedgerAttributes {
	return dataConfidentialLedgerAttributes{ref: terra.ReferenceDataResource(cl)}
}

// DataConfidentialLedgerArgs contains the configurations for azurerm_confidential_ledger.
type DataConfidentialLedgerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AzureadBasedServicePrincipal: min=0
	AzureadBasedServicePrincipal []dataconfidentialledger.AzureadBasedServicePrincipal `hcl:"azuread_based_service_principal,block" validate:"min=0"`
	// CertificateBasedSecurityPrincipal: min=0
	CertificateBasedSecurityPrincipal []dataconfidentialledger.CertificateBasedSecurityPrincipal `hcl:"certificate_based_security_principal,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataconfidentialledger.Timeouts `hcl:"timeouts,block"`
}
type dataConfidentialLedgerAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_confidential_ledger.
func (cl dataConfidentialLedgerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("id"))
}

// IdentityServiceEndpoint returns a reference to field identity_service_endpoint of azurerm_confidential_ledger.
func (cl dataConfidentialLedgerAttributes) IdentityServiceEndpoint() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("identity_service_endpoint"))
}

// LedgerEndpoint returns a reference to field ledger_endpoint of azurerm_confidential_ledger.
func (cl dataConfidentialLedgerAttributes) LedgerEndpoint() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("ledger_endpoint"))
}

// LedgerType returns a reference to field ledger_type of azurerm_confidential_ledger.
func (cl dataConfidentialLedgerAttributes) LedgerType() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("ledger_type"))
}

// Location returns a reference to field location of azurerm_confidential_ledger.
func (cl dataConfidentialLedgerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_confidential_ledger.
func (cl dataConfidentialLedgerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_confidential_ledger.
func (cl dataConfidentialLedgerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_confidential_ledger.
func (cl dataConfidentialLedgerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cl.ref.Append("tags"))
}

func (cl dataConfidentialLedgerAttributes) AzureadBasedServicePrincipal() terra.ListValue[dataconfidentialledger.AzureadBasedServicePrincipalAttributes] {
	return terra.ReferenceAsList[dataconfidentialledger.AzureadBasedServicePrincipalAttributes](cl.ref.Append("azuread_based_service_principal"))
}

func (cl dataConfidentialLedgerAttributes) CertificateBasedSecurityPrincipal() terra.ListValue[dataconfidentialledger.CertificateBasedSecurityPrincipalAttributes] {
	return terra.ReferenceAsList[dataconfidentialledger.CertificateBasedSecurityPrincipalAttributes](cl.ref.Append("certificate_based_security_principal"))
}

func (cl dataConfidentialLedgerAttributes) Timeouts() dataconfidentialledger.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataconfidentialledger.TimeoutsAttributes](cl.ref.Append("timeouts"))
}
