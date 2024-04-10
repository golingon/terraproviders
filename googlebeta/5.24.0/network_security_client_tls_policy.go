// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	networksecurityclienttlspolicy "github.com/golingon/terraproviders/googlebeta/5.24.0/networksecurityclienttlspolicy"
	"io"
)

// NewNetworkSecurityClientTlsPolicy creates a new instance of [NetworkSecurityClientTlsPolicy].
func NewNetworkSecurityClientTlsPolicy(name string, args NetworkSecurityClientTlsPolicyArgs) *NetworkSecurityClientTlsPolicy {
	return &NetworkSecurityClientTlsPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkSecurityClientTlsPolicy)(nil)

// NetworkSecurityClientTlsPolicy represents the Terraform resource google_network_security_client_tls_policy.
type NetworkSecurityClientTlsPolicy struct {
	Name      string
	Args      NetworkSecurityClientTlsPolicyArgs
	state     *networkSecurityClientTlsPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkSecurityClientTlsPolicy].
func (nsctp *NetworkSecurityClientTlsPolicy) Type() string {
	return "google_network_security_client_tls_policy"
}

// LocalName returns the local name for [NetworkSecurityClientTlsPolicy].
func (nsctp *NetworkSecurityClientTlsPolicy) LocalName() string {
	return nsctp.Name
}

// Configuration returns the configuration (args) for [NetworkSecurityClientTlsPolicy].
func (nsctp *NetworkSecurityClientTlsPolicy) Configuration() interface{} {
	return nsctp.Args
}

// DependOn is used for other resources to depend on [NetworkSecurityClientTlsPolicy].
func (nsctp *NetworkSecurityClientTlsPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(nsctp)
}

// Dependencies returns the list of resources [NetworkSecurityClientTlsPolicy] depends_on.
func (nsctp *NetworkSecurityClientTlsPolicy) Dependencies() terra.Dependencies {
	return nsctp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkSecurityClientTlsPolicy].
func (nsctp *NetworkSecurityClientTlsPolicy) LifecycleManagement() *terra.Lifecycle {
	return nsctp.Lifecycle
}

// Attributes returns the attributes for [NetworkSecurityClientTlsPolicy].
func (nsctp *NetworkSecurityClientTlsPolicy) Attributes() networkSecurityClientTlsPolicyAttributes {
	return networkSecurityClientTlsPolicyAttributes{ref: terra.ReferenceResource(nsctp)}
}

// ImportState imports the given attribute values into [NetworkSecurityClientTlsPolicy]'s state.
func (nsctp *NetworkSecurityClientTlsPolicy) ImportState(av io.Reader) error {
	nsctp.state = &networkSecurityClientTlsPolicyState{}
	if err := json.NewDecoder(av).Decode(nsctp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsctp.Type(), nsctp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkSecurityClientTlsPolicy] has state.
func (nsctp *NetworkSecurityClientTlsPolicy) State() (*networkSecurityClientTlsPolicyState, bool) {
	return nsctp.state, nsctp.state != nil
}

// StateMust returns the state for [NetworkSecurityClientTlsPolicy]. Panics if the state is nil.
func (nsctp *NetworkSecurityClientTlsPolicy) StateMust() *networkSecurityClientTlsPolicyState {
	if nsctp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsctp.Type(), nsctp.LocalName()))
	}
	return nsctp.state
}

// NetworkSecurityClientTlsPolicyArgs contains the configurations for google_network_security_client_tls_policy.
type NetworkSecurityClientTlsPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Sni: string, optional
	Sni terra.StringValue `hcl:"sni,attr"`
	// ClientCertificate: optional
	ClientCertificate *networksecurityclienttlspolicy.ClientCertificate `hcl:"client_certificate,block"`
	// ServerValidationCa: min=0
	ServerValidationCa []networksecurityclienttlspolicy.ServerValidationCa `hcl:"server_validation_ca,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *networksecurityclienttlspolicy.Timeouts `hcl:"timeouts,block"`
}
type networkSecurityClientTlsPolicyAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nsctp.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsctp.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsctp.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsctp.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsctp.ref.Append("labels"))
}

// Location returns a reference to field location of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nsctp.ref.Append("location"))
}

// Name returns a reference to field name of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsctp.ref.Append("name"))
}

// Project returns a reference to field project of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nsctp.ref.Append("project"))
}

// Sni returns a reference to field sni of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) Sni() terra.StringValue {
	return terra.ReferenceAsString(nsctp.ref.Append("sni"))
}

// TerraformLabels returns a reference to field terraform_labels of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsctp.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_network_security_client_tls_policy.
func (nsctp networkSecurityClientTlsPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nsctp.ref.Append("update_time"))
}

func (nsctp networkSecurityClientTlsPolicyAttributes) ClientCertificate() terra.ListValue[networksecurityclienttlspolicy.ClientCertificateAttributes] {
	return terra.ReferenceAsList[networksecurityclienttlspolicy.ClientCertificateAttributes](nsctp.ref.Append("client_certificate"))
}

func (nsctp networkSecurityClientTlsPolicyAttributes) ServerValidationCa() terra.ListValue[networksecurityclienttlspolicy.ServerValidationCaAttributes] {
	return terra.ReferenceAsList[networksecurityclienttlspolicy.ServerValidationCaAttributes](nsctp.ref.Append("server_validation_ca"))
}

func (nsctp networkSecurityClientTlsPolicyAttributes) Timeouts() networksecurityclienttlspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networksecurityclienttlspolicy.TimeoutsAttributes](nsctp.ref.Append("timeouts"))
}

type networkSecurityClientTlsPolicyState struct {
	CreateTime         string                                                   `json:"create_time"`
	Description        string                                                   `json:"description"`
	EffectiveLabels    map[string]string                                        `json:"effective_labels"`
	Id                 string                                                   `json:"id"`
	Labels             map[string]string                                        `json:"labels"`
	Location           string                                                   `json:"location"`
	Name               string                                                   `json:"name"`
	Project            string                                                   `json:"project"`
	Sni                string                                                   `json:"sni"`
	TerraformLabels    map[string]string                                        `json:"terraform_labels"`
	UpdateTime         string                                                   `json:"update_time"`
	ClientCertificate  []networksecurityclienttlspolicy.ClientCertificateState  `json:"client_certificate"`
	ServerValidationCa []networksecurityclienttlspolicy.ServerValidationCaState `json:"server_validation_ca"`
	Timeouts           *networksecurityclienttlspolicy.TimeoutsState            `json:"timeouts"`
}
