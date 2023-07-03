// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	activedirectorydomainservicetrust "github.com/golingon/terraproviders/azurerm/3.63.0/activedirectorydomainservicetrust"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewActiveDirectoryDomainServiceTrust creates a new instance of [ActiveDirectoryDomainServiceTrust].
func NewActiveDirectoryDomainServiceTrust(name string, args ActiveDirectoryDomainServiceTrustArgs) *ActiveDirectoryDomainServiceTrust {
	return &ActiveDirectoryDomainServiceTrust{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ActiveDirectoryDomainServiceTrust)(nil)

// ActiveDirectoryDomainServiceTrust represents the Terraform resource azurerm_active_directory_domain_service_trust.
type ActiveDirectoryDomainServiceTrust struct {
	Name      string
	Args      ActiveDirectoryDomainServiceTrustArgs
	state     *activeDirectoryDomainServiceTrustState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ActiveDirectoryDomainServiceTrust].
func (addst *ActiveDirectoryDomainServiceTrust) Type() string {
	return "azurerm_active_directory_domain_service_trust"
}

// LocalName returns the local name for [ActiveDirectoryDomainServiceTrust].
func (addst *ActiveDirectoryDomainServiceTrust) LocalName() string {
	return addst.Name
}

// Configuration returns the configuration (args) for [ActiveDirectoryDomainServiceTrust].
func (addst *ActiveDirectoryDomainServiceTrust) Configuration() interface{} {
	return addst.Args
}

// DependOn is used for other resources to depend on [ActiveDirectoryDomainServiceTrust].
func (addst *ActiveDirectoryDomainServiceTrust) DependOn() terra.Reference {
	return terra.ReferenceResource(addst)
}

// Dependencies returns the list of resources [ActiveDirectoryDomainServiceTrust] depends_on.
func (addst *ActiveDirectoryDomainServiceTrust) Dependencies() terra.Dependencies {
	return addst.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ActiveDirectoryDomainServiceTrust].
func (addst *ActiveDirectoryDomainServiceTrust) LifecycleManagement() *terra.Lifecycle {
	return addst.Lifecycle
}

// Attributes returns the attributes for [ActiveDirectoryDomainServiceTrust].
func (addst *ActiveDirectoryDomainServiceTrust) Attributes() activeDirectoryDomainServiceTrustAttributes {
	return activeDirectoryDomainServiceTrustAttributes{ref: terra.ReferenceResource(addst)}
}

// ImportState imports the given attribute values into [ActiveDirectoryDomainServiceTrust]'s state.
func (addst *ActiveDirectoryDomainServiceTrust) ImportState(av io.Reader) error {
	addst.state = &activeDirectoryDomainServiceTrustState{}
	if err := json.NewDecoder(av).Decode(addst.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", addst.Type(), addst.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ActiveDirectoryDomainServiceTrust] has state.
func (addst *ActiveDirectoryDomainServiceTrust) State() (*activeDirectoryDomainServiceTrustState, bool) {
	return addst.state, addst.state != nil
}

// StateMust returns the state for [ActiveDirectoryDomainServiceTrust]. Panics if the state is nil.
func (addst *ActiveDirectoryDomainServiceTrust) StateMust() *activeDirectoryDomainServiceTrustState {
	if addst.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", addst.Type(), addst.LocalName()))
	}
	return addst.state
}

// ActiveDirectoryDomainServiceTrustArgs contains the configurations for azurerm_active_directory_domain_service_trust.
type ActiveDirectoryDomainServiceTrustArgs struct {
	// DomainServiceId: string, required
	DomainServiceId terra.StringValue `hcl:"domain_service_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// TrustedDomainDnsIps: list of string, required
	TrustedDomainDnsIps terra.ListValue[terra.StringValue] `hcl:"trusted_domain_dns_ips,attr" validate:"required"`
	// TrustedDomainFqdn: string, required
	TrustedDomainFqdn terra.StringValue `hcl:"trusted_domain_fqdn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *activedirectorydomainservicetrust.Timeouts `hcl:"timeouts,block"`
}
type activeDirectoryDomainServiceTrustAttributes struct {
	ref terra.Reference
}

// DomainServiceId returns a reference to field domain_service_id of azurerm_active_directory_domain_service_trust.
func (addst activeDirectoryDomainServiceTrustAttributes) DomainServiceId() terra.StringValue {
	return terra.ReferenceAsString(addst.ref.Append("domain_service_id"))
}

// Id returns a reference to field id of azurerm_active_directory_domain_service_trust.
func (addst activeDirectoryDomainServiceTrustAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(addst.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_active_directory_domain_service_trust.
func (addst activeDirectoryDomainServiceTrustAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(addst.ref.Append("name"))
}

// Password returns a reference to field password of azurerm_active_directory_domain_service_trust.
func (addst activeDirectoryDomainServiceTrustAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(addst.ref.Append("password"))
}

// TrustedDomainDnsIps returns a reference to field trusted_domain_dns_ips of azurerm_active_directory_domain_service_trust.
func (addst activeDirectoryDomainServiceTrustAttributes) TrustedDomainDnsIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](addst.ref.Append("trusted_domain_dns_ips"))
}

// TrustedDomainFqdn returns a reference to field trusted_domain_fqdn of azurerm_active_directory_domain_service_trust.
func (addst activeDirectoryDomainServiceTrustAttributes) TrustedDomainFqdn() terra.StringValue {
	return terra.ReferenceAsString(addst.ref.Append("trusted_domain_fqdn"))
}

func (addst activeDirectoryDomainServiceTrustAttributes) Timeouts() activedirectorydomainservicetrust.TimeoutsAttributes {
	return terra.ReferenceAsSingle[activedirectorydomainservicetrust.TimeoutsAttributes](addst.ref.Append("timeouts"))
}

type activeDirectoryDomainServiceTrustState struct {
	DomainServiceId     string                                           `json:"domain_service_id"`
	Id                  string                                           `json:"id"`
	Name                string                                           `json:"name"`
	Password            string                                           `json:"password"`
	TrustedDomainDnsIps []string                                         `json:"trusted_domain_dns_ips"`
	TrustedDomainFqdn   string                                           `json:"trusted_domain_fqdn"`
	Timeouts            *activedirectorydomainservicetrust.TimeoutsState `json:"timeouts"`
}
