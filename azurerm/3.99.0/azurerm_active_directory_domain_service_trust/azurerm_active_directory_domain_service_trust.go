// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_active_directory_domain_service_trust

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

// Resource represents the Terraform resource azurerm_active_directory_domain_service_trust.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermActiveDirectoryDomainServiceTrustState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aaddst *Resource) Type() string {
	return "azurerm_active_directory_domain_service_trust"
}

// LocalName returns the local name for [Resource].
func (aaddst *Resource) LocalName() string {
	return aaddst.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aaddst *Resource) Configuration() interface{} {
	return aaddst.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aaddst *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aaddst)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aaddst *Resource) Dependencies() terra.Dependencies {
	return aaddst.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aaddst *Resource) LifecycleManagement() *terra.Lifecycle {
	return aaddst.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aaddst *Resource) Attributes() azurermActiveDirectoryDomainServiceTrustAttributes {
	return azurermActiveDirectoryDomainServiceTrustAttributes{ref: terra.ReferenceResource(aaddst)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aaddst *Resource) ImportState(state io.Reader) error {
	aaddst.state = &azurermActiveDirectoryDomainServiceTrustState{}
	if err := json.NewDecoder(state).Decode(aaddst.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aaddst.Type(), aaddst.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aaddst *Resource) State() (*azurermActiveDirectoryDomainServiceTrustState, bool) {
	return aaddst.state, aaddst.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aaddst *Resource) StateMust() *azurermActiveDirectoryDomainServiceTrustState {
	if aaddst.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aaddst.Type(), aaddst.LocalName()))
	}
	return aaddst.state
}

// Args contains the configurations for azurerm_active_directory_domain_service_trust.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermActiveDirectoryDomainServiceTrustAttributes struct {
	ref terra.Reference
}

// DomainServiceId returns a reference to field domain_service_id of azurerm_active_directory_domain_service_trust.
func (aaddst azurermActiveDirectoryDomainServiceTrustAttributes) DomainServiceId() terra.StringValue {
	return terra.ReferenceAsString(aaddst.ref.Append("domain_service_id"))
}

// Id returns a reference to field id of azurerm_active_directory_domain_service_trust.
func (aaddst azurermActiveDirectoryDomainServiceTrustAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aaddst.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_active_directory_domain_service_trust.
func (aaddst azurermActiveDirectoryDomainServiceTrustAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aaddst.ref.Append("name"))
}

// Password returns a reference to field password of azurerm_active_directory_domain_service_trust.
func (aaddst azurermActiveDirectoryDomainServiceTrustAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(aaddst.ref.Append("password"))
}

// TrustedDomainDnsIps returns a reference to field trusted_domain_dns_ips of azurerm_active_directory_domain_service_trust.
func (aaddst azurermActiveDirectoryDomainServiceTrustAttributes) TrustedDomainDnsIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aaddst.ref.Append("trusted_domain_dns_ips"))
}

// TrustedDomainFqdn returns a reference to field trusted_domain_fqdn of azurerm_active_directory_domain_service_trust.
func (aaddst azurermActiveDirectoryDomainServiceTrustAttributes) TrustedDomainFqdn() terra.StringValue {
	return terra.ReferenceAsString(aaddst.ref.Append("trusted_domain_fqdn"))
}

func (aaddst azurermActiveDirectoryDomainServiceTrustAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aaddst.ref.Append("timeouts"))
}

type azurermActiveDirectoryDomainServiceTrustState struct {
	DomainServiceId     string         `json:"domain_service_id"`
	Id                  string         `json:"id"`
	Name                string         `json:"name"`
	Password            string         `json:"password"`
	TrustedDomainDnsIps []string       `json:"trusted_domain_dns_ips"`
	TrustedDomainFqdn   string         `json:"trusted_domain_fqdn"`
	Timeouts            *TimeoutsState `json:"timeouts"`
}
