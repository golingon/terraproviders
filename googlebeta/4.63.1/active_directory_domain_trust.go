// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	activedirectorydomaintrust "github.com/golingon/terraproviders/googlebeta/4.63.1/activedirectorydomaintrust"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewActiveDirectoryDomainTrust creates a new instance of [ActiveDirectoryDomainTrust].
func NewActiveDirectoryDomainTrust(name string, args ActiveDirectoryDomainTrustArgs) *ActiveDirectoryDomainTrust {
	return &ActiveDirectoryDomainTrust{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ActiveDirectoryDomainTrust)(nil)

// ActiveDirectoryDomainTrust represents the Terraform resource google_active_directory_domain_trust.
type ActiveDirectoryDomainTrust struct {
	Name      string
	Args      ActiveDirectoryDomainTrustArgs
	state     *activeDirectoryDomainTrustState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ActiveDirectoryDomainTrust].
func (addt *ActiveDirectoryDomainTrust) Type() string {
	return "google_active_directory_domain_trust"
}

// LocalName returns the local name for [ActiveDirectoryDomainTrust].
func (addt *ActiveDirectoryDomainTrust) LocalName() string {
	return addt.Name
}

// Configuration returns the configuration (args) for [ActiveDirectoryDomainTrust].
func (addt *ActiveDirectoryDomainTrust) Configuration() interface{} {
	return addt.Args
}

// DependOn is used for other resources to depend on [ActiveDirectoryDomainTrust].
func (addt *ActiveDirectoryDomainTrust) DependOn() terra.Reference {
	return terra.ReferenceResource(addt)
}

// Dependencies returns the list of resources [ActiveDirectoryDomainTrust] depends_on.
func (addt *ActiveDirectoryDomainTrust) Dependencies() terra.Dependencies {
	return addt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ActiveDirectoryDomainTrust].
func (addt *ActiveDirectoryDomainTrust) LifecycleManagement() *terra.Lifecycle {
	return addt.Lifecycle
}

// Attributes returns the attributes for [ActiveDirectoryDomainTrust].
func (addt *ActiveDirectoryDomainTrust) Attributes() activeDirectoryDomainTrustAttributes {
	return activeDirectoryDomainTrustAttributes{ref: terra.ReferenceResource(addt)}
}

// ImportState imports the given attribute values into [ActiveDirectoryDomainTrust]'s state.
func (addt *ActiveDirectoryDomainTrust) ImportState(av io.Reader) error {
	addt.state = &activeDirectoryDomainTrustState{}
	if err := json.NewDecoder(av).Decode(addt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", addt.Type(), addt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ActiveDirectoryDomainTrust] has state.
func (addt *ActiveDirectoryDomainTrust) State() (*activeDirectoryDomainTrustState, bool) {
	return addt.state, addt.state != nil
}

// StateMust returns the state for [ActiveDirectoryDomainTrust]. Panics if the state is nil.
func (addt *ActiveDirectoryDomainTrust) StateMust() *activeDirectoryDomainTrustState {
	if addt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", addt.Type(), addt.LocalName()))
	}
	return addt.state
}

// ActiveDirectoryDomainTrustArgs contains the configurations for google_active_directory_domain_trust.
type ActiveDirectoryDomainTrustArgs struct {
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SelectiveAuthentication: bool, optional
	SelectiveAuthentication terra.BoolValue `hcl:"selective_authentication,attr"`
	// TargetDnsIpAddresses: set of string, required
	TargetDnsIpAddresses terra.SetValue[terra.StringValue] `hcl:"target_dns_ip_addresses,attr" validate:"required"`
	// TargetDomainName: string, required
	TargetDomainName terra.StringValue `hcl:"target_domain_name,attr" validate:"required"`
	// TrustDirection: string, required
	TrustDirection terra.StringValue `hcl:"trust_direction,attr" validate:"required"`
	// TrustHandshakeSecret: string, required
	TrustHandshakeSecret terra.StringValue `hcl:"trust_handshake_secret,attr" validate:"required"`
	// TrustType: string, required
	TrustType terra.StringValue `hcl:"trust_type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *activedirectorydomaintrust.Timeouts `hcl:"timeouts,block"`
}
type activeDirectoryDomainTrustAttributes struct {
	ref terra.Reference
}

// Domain returns a reference to field domain of google_active_directory_domain_trust.
func (addt activeDirectoryDomainTrustAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(addt.ref.Append("domain"))
}

// Id returns a reference to field id of google_active_directory_domain_trust.
func (addt activeDirectoryDomainTrustAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(addt.ref.Append("id"))
}

// Project returns a reference to field project of google_active_directory_domain_trust.
func (addt activeDirectoryDomainTrustAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(addt.ref.Append("project"))
}

// SelectiveAuthentication returns a reference to field selective_authentication of google_active_directory_domain_trust.
func (addt activeDirectoryDomainTrustAttributes) SelectiveAuthentication() terra.BoolValue {
	return terra.ReferenceAsBool(addt.ref.Append("selective_authentication"))
}

// TargetDnsIpAddresses returns a reference to field target_dns_ip_addresses of google_active_directory_domain_trust.
func (addt activeDirectoryDomainTrustAttributes) TargetDnsIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](addt.ref.Append("target_dns_ip_addresses"))
}

// TargetDomainName returns a reference to field target_domain_name of google_active_directory_domain_trust.
func (addt activeDirectoryDomainTrustAttributes) TargetDomainName() terra.StringValue {
	return terra.ReferenceAsString(addt.ref.Append("target_domain_name"))
}

// TrustDirection returns a reference to field trust_direction of google_active_directory_domain_trust.
func (addt activeDirectoryDomainTrustAttributes) TrustDirection() terra.StringValue {
	return terra.ReferenceAsString(addt.ref.Append("trust_direction"))
}

// TrustHandshakeSecret returns a reference to field trust_handshake_secret of google_active_directory_domain_trust.
func (addt activeDirectoryDomainTrustAttributes) TrustHandshakeSecret() terra.StringValue {
	return terra.ReferenceAsString(addt.ref.Append("trust_handshake_secret"))
}

// TrustType returns a reference to field trust_type of google_active_directory_domain_trust.
func (addt activeDirectoryDomainTrustAttributes) TrustType() terra.StringValue {
	return terra.ReferenceAsString(addt.ref.Append("trust_type"))
}

func (addt activeDirectoryDomainTrustAttributes) Timeouts() activedirectorydomaintrust.TimeoutsAttributes {
	return terra.ReferenceAsSingle[activedirectorydomaintrust.TimeoutsAttributes](addt.ref.Append("timeouts"))
}

type activeDirectoryDomainTrustState struct {
	Domain                  string                                    `json:"domain"`
	Id                      string                                    `json:"id"`
	Project                 string                                    `json:"project"`
	SelectiveAuthentication bool                                      `json:"selective_authentication"`
	TargetDnsIpAddresses    []string                                  `json:"target_dns_ip_addresses"`
	TargetDomainName        string                                    `json:"target_domain_name"`
	TrustDirection          string                                    `json:"trust_direction"`
	TrustHandshakeSecret    string                                    `json:"trust_handshake_secret"`
	TrustType               string                                    `json:"trust_type"`
	Timeouts                *activedirectorydomaintrust.TimeoutsState `json:"timeouts"`
}