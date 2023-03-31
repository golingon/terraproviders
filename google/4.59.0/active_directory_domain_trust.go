// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	activedirectorydomaintrust "github.com/golingon/terraproviders/google/4.59.0/activedirectorydomaintrust"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewActiveDirectoryDomainTrust(name string, args ActiveDirectoryDomainTrustArgs) *ActiveDirectoryDomainTrust {
	return &ActiveDirectoryDomainTrust{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ActiveDirectoryDomainTrust)(nil)

type ActiveDirectoryDomainTrust struct {
	Name  string
	Args  ActiveDirectoryDomainTrustArgs
	state *activeDirectoryDomainTrustState
}

func (addt *ActiveDirectoryDomainTrust) Type() string {
	return "google_active_directory_domain_trust"
}

func (addt *ActiveDirectoryDomainTrust) LocalName() string {
	return addt.Name
}

func (addt *ActiveDirectoryDomainTrust) Configuration() interface{} {
	return addt.Args
}

func (addt *ActiveDirectoryDomainTrust) Attributes() activeDirectoryDomainTrustAttributes {
	return activeDirectoryDomainTrustAttributes{ref: terra.ReferenceResource(addt)}
}

func (addt *ActiveDirectoryDomainTrust) ImportState(av io.Reader) error {
	addt.state = &activeDirectoryDomainTrustState{}
	if err := json.NewDecoder(av).Decode(addt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", addt.Type(), addt.LocalName(), err)
	}
	return nil
}

func (addt *ActiveDirectoryDomainTrust) State() (*activeDirectoryDomainTrustState, bool) {
	return addt.state, addt.state != nil
}

func (addt *ActiveDirectoryDomainTrust) StateMust() *activeDirectoryDomainTrustState {
	if addt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", addt.Type(), addt.LocalName()))
	}
	return addt.state
}

func (addt *ActiveDirectoryDomainTrust) DependOn() terra.Reference {
	return terra.ReferenceResource(addt)
}

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
	// DependsOn contains resources that ActiveDirectoryDomainTrust depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type activeDirectoryDomainTrustAttributes struct {
	ref terra.Reference
}

func (addt activeDirectoryDomainTrustAttributes) Domain() terra.StringValue {
	return terra.ReferenceString(addt.ref.Append("domain"))
}

func (addt activeDirectoryDomainTrustAttributes) Id() terra.StringValue {
	return terra.ReferenceString(addt.ref.Append("id"))
}

func (addt activeDirectoryDomainTrustAttributes) Project() terra.StringValue {
	return terra.ReferenceString(addt.ref.Append("project"))
}

func (addt activeDirectoryDomainTrustAttributes) SelectiveAuthentication() terra.BoolValue {
	return terra.ReferenceBool(addt.ref.Append("selective_authentication"))
}

func (addt activeDirectoryDomainTrustAttributes) TargetDnsIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](addt.ref.Append("target_dns_ip_addresses"))
}

func (addt activeDirectoryDomainTrustAttributes) TargetDomainName() terra.StringValue {
	return terra.ReferenceString(addt.ref.Append("target_domain_name"))
}

func (addt activeDirectoryDomainTrustAttributes) TrustDirection() terra.StringValue {
	return terra.ReferenceString(addt.ref.Append("trust_direction"))
}

func (addt activeDirectoryDomainTrustAttributes) TrustHandshakeSecret() terra.StringValue {
	return terra.ReferenceString(addt.ref.Append("trust_handshake_secret"))
}

func (addt activeDirectoryDomainTrustAttributes) TrustType() terra.StringValue {
	return terra.ReferenceString(addt.ref.Append("trust_type"))
}

func (addt activeDirectoryDomainTrustAttributes) Timeouts() activedirectorydomaintrust.TimeoutsAttributes {
	return terra.ReferenceSingle[activedirectorydomaintrust.TimeoutsAttributes](addt.ref.Append("timeouts"))
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
