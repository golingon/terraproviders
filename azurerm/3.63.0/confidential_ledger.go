// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	confidentialledger "github.com/golingon/terraproviders/azurerm/3.63.0/confidentialledger"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConfidentialLedger creates a new instance of [ConfidentialLedger].
func NewConfidentialLedger(name string, args ConfidentialLedgerArgs) *ConfidentialLedger {
	return &ConfidentialLedger{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfidentialLedger)(nil)

// ConfidentialLedger represents the Terraform resource azurerm_confidential_ledger.
type ConfidentialLedger struct {
	Name      string
	Args      ConfidentialLedgerArgs
	state     *confidentialLedgerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConfidentialLedger].
func (cl *ConfidentialLedger) Type() string {
	return "azurerm_confidential_ledger"
}

// LocalName returns the local name for [ConfidentialLedger].
func (cl *ConfidentialLedger) LocalName() string {
	return cl.Name
}

// Configuration returns the configuration (args) for [ConfidentialLedger].
func (cl *ConfidentialLedger) Configuration() interface{} {
	return cl.Args
}

// DependOn is used for other resources to depend on [ConfidentialLedger].
func (cl *ConfidentialLedger) DependOn() terra.Reference {
	return terra.ReferenceResource(cl)
}

// Dependencies returns the list of resources [ConfidentialLedger] depends_on.
func (cl *ConfidentialLedger) Dependencies() terra.Dependencies {
	return cl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConfidentialLedger].
func (cl *ConfidentialLedger) LifecycleManagement() *terra.Lifecycle {
	return cl.Lifecycle
}

// Attributes returns the attributes for [ConfidentialLedger].
func (cl *ConfidentialLedger) Attributes() confidentialLedgerAttributes {
	return confidentialLedgerAttributes{ref: terra.ReferenceResource(cl)}
}

// ImportState imports the given attribute values into [ConfidentialLedger]'s state.
func (cl *ConfidentialLedger) ImportState(av io.Reader) error {
	cl.state = &confidentialLedgerState{}
	if err := json.NewDecoder(av).Decode(cl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cl.Type(), cl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConfidentialLedger] has state.
func (cl *ConfidentialLedger) State() (*confidentialLedgerState, bool) {
	return cl.state, cl.state != nil
}

// StateMust returns the state for [ConfidentialLedger]. Panics if the state is nil.
func (cl *ConfidentialLedger) StateMust() *confidentialLedgerState {
	if cl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cl.Type(), cl.LocalName()))
	}
	return cl.state
}

// ConfidentialLedgerArgs contains the configurations for azurerm_confidential_ledger.
type ConfidentialLedgerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LedgerType: string, required
	LedgerType terra.StringValue `hcl:"ledger_type,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AzureadBasedServicePrincipal: min=1
	AzureadBasedServicePrincipal []confidentialledger.AzureadBasedServicePrincipal `hcl:"azuread_based_service_principal,block" validate:"min=1"`
	// CertificateBasedSecurityPrincipal: min=0
	CertificateBasedSecurityPrincipal []confidentialledger.CertificateBasedSecurityPrincipal `hcl:"certificate_based_security_principal,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *confidentialledger.Timeouts `hcl:"timeouts,block"`
}
type confidentialLedgerAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_confidential_ledger.
func (cl confidentialLedgerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("id"))
}

// IdentityServiceEndpoint returns a reference to field identity_service_endpoint of azurerm_confidential_ledger.
func (cl confidentialLedgerAttributes) IdentityServiceEndpoint() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("identity_service_endpoint"))
}

// LedgerEndpoint returns a reference to field ledger_endpoint of azurerm_confidential_ledger.
func (cl confidentialLedgerAttributes) LedgerEndpoint() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("ledger_endpoint"))
}

// LedgerType returns a reference to field ledger_type of azurerm_confidential_ledger.
func (cl confidentialLedgerAttributes) LedgerType() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("ledger_type"))
}

// Location returns a reference to field location of azurerm_confidential_ledger.
func (cl confidentialLedgerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_confidential_ledger.
func (cl confidentialLedgerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_confidential_ledger.
func (cl confidentialLedgerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_confidential_ledger.
func (cl confidentialLedgerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cl.ref.Append("tags"))
}

func (cl confidentialLedgerAttributes) AzureadBasedServicePrincipal() terra.ListValue[confidentialledger.AzureadBasedServicePrincipalAttributes] {
	return terra.ReferenceAsList[confidentialledger.AzureadBasedServicePrincipalAttributes](cl.ref.Append("azuread_based_service_principal"))
}

func (cl confidentialLedgerAttributes) CertificateBasedSecurityPrincipal() terra.ListValue[confidentialledger.CertificateBasedSecurityPrincipalAttributes] {
	return terra.ReferenceAsList[confidentialledger.CertificateBasedSecurityPrincipalAttributes](cl.ref.Append("certificate_based_security_principal"))
}

func (cl confidentialLedgerAttributes) Timeouts() confidentialledger.TimeoutsAttributes {
	return terra.ReferenceAsSingle[confidentialledger.TimeoutsAttributes](cl.ref.Append("timeouts"))
}

type confidentialLedgerState struct {
	Id                                string                                                      `json:"id"`
	IdentityServiceEndpoint           string                                                      `json:"identity_service_endpoint"`
	LedgerEndpoint                    string                                                      `json:"ledger_endpoint"`
	LedgerType                        string                                                      `json:"ledger_type"`
	Location                          string                                                      `json:"location"`
	Name                              string                                                      `json:"name"`
	ResourceGroupName                 string                                                      `json:"resource_group_name"`
	Tags                              map[string]string                                           `json:"tags"`
	AzureadBasedServicePrincipal      []confidentialledger.AzureadBasedServicePrincipalState      `json:"azuread_based_service_principal"`
	CertificateBasedSecurityPrincipal []confidentialledger.CertificateBasedSecurityPrincipalState `json:"certificate_based_security_principal"`
	Timeouts                          *confidentialledger.TimeoutsState                           `json:"timeouts"`
}
