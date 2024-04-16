// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_confidential_ledger

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_confidential_ledger.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermConfidentialLedgerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acl *Resource) Type() string {
	return "azurerm_confidential_ledger"
}

// LocalName returns the local name for [Resource].
func (acl *Resource) LocalName() string {
	return acl.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acl *Resource) Configuration() interface{} {
	return acl.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acl *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acl)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acl *Resource) Dependencies() terra.Dependencies {
	return acl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acl *Resource) LifecycleManagement() *terra.Lifecycle {
	return acl.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acl *Resource) Attributes() azurermConfidentialLedgerAttributes {
	return azurermConfidentialLedgerAttributes{ref: terra.ReferenceResource(acl)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acl *Resource) ImportState(state io.Reader) error {
	acl.state = &azurermConfidentialLedgerState{}
	if err := json.NewDecoder(state).Decode(acl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acl.Type(), acl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acl *Resource) State() (*azurermConfidentialLedgerState, bool) {
	return acl.state, acl.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acl *Resource) StateMust() *azurermConfidentialLedgerState {
	if acl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acl.Type(), acl.LocalName()))
	}
	return acl.state
}

// Args contains the configurations for azurerm_confidential_ledger.
type Args struct {
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
	AzureadBasedServicePrincipal []AzureadBasedServicePrincipal `hcl:"azuread_based_service_principal,block" validate:"min=1"`
	// CertificateBasedSecurityPrincipal: min=0
	CertificateBasedSecurityPrincipal []CertificateBasedSecurityPrincipal `hcl:"certificate_based_security_principal,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermConfidentialLedgerAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_confidential_ledger.
func (acl azurermConfidentialLedgerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acl.ref.Append("id"))
}

// IdentityServiceEndpoint returns a reference to field identity_service_endpoint of azurerm_confidential_ledger.
func (acl azurermConfidentialLedgerAttributes) IdentityServiceEndpoint() terra.StringValue {
	return terra.ReferenceAsString(acl.ref.Append("identity_service_endpoint"))
}

// LedgerEndpoint returns a reference to field ledger_endpoint of azurerm_confidential_ledger.
func (acl azurermConfidentialLedgerAttributes) LedgerEndpoint() terra.StringValue {
	return terra.ReferenceAsString(acl.ref.Append("ledger_endpoint"))
}

// LedgerType returns a reference to field ledger_type of azurerm_confidential_ledger.
func (acl azurermConfidentialLedgerAttributes) LedgerType() terra.StringValue {
	return terra.ReferenceAsString(acl.ref.Append("ledger_type"))
}

// Location returns a reference to field location of azurerm_confidential_ledger.
func (acl azurermConfidentialLedgerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(acl.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_confidential_ledger.
func (acl azurermConfidentialLedgerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acl.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_confidential_ledger.
func (acl azurermConfidentialLedgerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(acl.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_confidential_ledger.
func (acl azurermConfidentialLedgerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acl.ref.Append("tags"))
}

func (acl azurermConfidentialLedgerAttributes) AzureadBasedServicePrincipal() terra.ListValue[AzureadBasedServicePrincipalAttributes] {
	return terra.ReferenceAsList[AzureadBasedServicePrincipalAttributes](acl.ref.Append("azuread_based_service_principal"))
}

func (acl azurermConfidentialLedgerAttributes) CertificateBasedSecurityPrincipal() terra.ListValue[CertificateBasedSecurityPrincipalAttributes] {
	return terra.ReferenceAsList[CertificateBasedSecurityPrincipalAttributes](acl.ref.Append("certificate_based_security_principal"))
}

func (acl azurermConfidentialLedgerAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acl.ref.Append("timeouts"))
}

type azurermConfidentialLedgerState struct {
	Id                                string                                   `json:"id"`
	IdentityServiceEndpoint           string                                   `json:"identity_service_endpoint"`
	LedgerEndpoint                    string                                   `json:"ledger_endpoint"`
	LedgerType                        string                                   `json:"ledger_type"`
	Location                          string                                   `json:"location"`
	Name                              string                                   `json:"name"`
	ResourceGroupName                 string                                   `json:"resource_group_name"`
	Tags                              map[string]string                        `json:"tags"`
	AzureadBasedServicePrincipal      []AzureadBasedServicePrincipalState      `json:"azuread_based_service_principal"`
	CertificateBasedSecurityPrincipal []CertificateBasedSecurityPrincipalState `json:"certificate_based_security_principal"`
	Timeouts                          *TimeoutsState                           `json:"timeouts"`
}
