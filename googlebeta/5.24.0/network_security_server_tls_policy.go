// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	networksecurityservertlspolicy "github.com/golingon/terraproviders/googlebeta/5.24.0/networksecurityservertlspolicy"
	"io"
)

// NewNetworkSecurityServerTlsPolicy creates a new instance of [NetworkSecurityServerTlsPolicy].
func NewNetworkSecurityServerTlsPolicy(name string, args NetworkSecurityServerTlsPolicyArgs) *NetworkSecurityServerTlsPolicy {
	return &NetworkSecurityServerTlsPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkSecurityServerTlsPolicy)(nil)

// NetworkSecurityServerTlsPolicy represents the Terraform resource google_network_security_server_tls_policy.
type NetworkSecurityServerTlsPolicy struct {
	Name      string
	Args      NetworkSecurityServerTlsPolicyArgs
	state     *networkSecurityServerTlsPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkSecurityServerTlsPolicy].
func (nsstp *NetworkSecurityServerTlsPolicy) Type() string {
	return "google_network_security_server_tls_policy"
}

// LocalName returns the local name for [NetworkSecurityServerTlsPolicy].
func (nsstp *NetworkSecurityServerTlsPolicy) LocalName() string {
	return nsstp.Name
}

// Configuration returns the configuration (args) for [NetworkSecurityServerTlsPolicy].
func (nsstp *NetworkSecurityServerTlsPolicy) Configuration() interface{} {
	return nsstp.Args
}

// DependOn is used for other resources to depend on [NetworkSecurityServerTlsPolicy].
func (nsstp *NetworkSecurityServerTlsPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(nsstp)
}

// Dependencies returns the list of resources [NetworkSecurityServerTlsPolicy] depends_on.
func (nsstp *NetworkSecurityServerTlsPolicy) Dependencies() terra.Dependencies {
	return nsstp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkSecurityServerTlsPolicy].
func (nsstp *NetworkSecurityServerTlsPolicy) LifecycleManagement() *terra.Lifecycle {
	return nsstp.Lifecycle
}

// Attributes returns the attributes for [NetworkSecurityServerTlsPolicy].
func (nsstp *NetworkSecurityServerTlsPolicy) Attributes() networkSecurityServerTlsPolicyAttributes {
	return networkSecurityServerTlsPolicyAttributes{ref: terra.ReferenceResource(nsstp)}
}

// ImportState imports the given attribute values into [NetworkSecurityServerTlsPolicy]'s state.
func (nsstp *NetworkSecurityServerTlsPolicy) ImportState(av io.Reader) error {
	nsstp.state = &networkSecurityServerTlsPolicyState{}
	if err := json.NewDecoder(av).Decode(nsstp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsstp.Type(), nsstp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkSecurityServerTlsPolicy] has state.
func (nsstp *NetworkSecurityServerTlsPolicy) State() (*networkSecurityServerTlsPolicyState, bool) {
	return nsstp.state, nsstp.state != nil
}

// StateMust returns the state for [NetworkSecurityServerTlsPolicy]. Panics if the state is nil.
func (nsstp *NetworkSecurityServerTlsPolicy) StateMust() *networkSecurityServerTlsPolicyState {
	if nsstp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsstp.Type(), nsstp.LocalName()))
	}
	return nsstp.state
}

// NetworkSecurityServerTlsPolicyArgs contains the configurations for google_network_security_server_tls_policy.
type NetworkSecurityServerTlsPolicyArgs struct {
	// AllowOpen: bool, optional
	AllowOpen terra.BoolValue `hcl:"allow_open,attr"`
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
	// MtlsPolicy: optional
	MtlsPolicy *networksecurityservertlspolicy.MtlsPolicy `hcl:"mtls_policy,block"`
	// ServerCertificate: optional
	ServerCertificate *networksecurityservertlspolicy.ServerCertificate `hcl:"server_certificate,block"`
	// Timeouts: optional
	Timeouts *networksecurityservertlspolicy.Timeouts `hcl:"timeouts,block"`
}
type networkSecurityServerTlsPolicyAttributes struct {
	ref terra.Reference
}

// AllowOpen returns a reference to field allow_open of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) AllowOpen() terra.BoolValue {
	return terra.ReferenceAsBool(nsstp.ref.Append("allow_open"))
}

// CreateTime returns a reference to field create_time of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nsstp.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsstp.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsstp.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsstp.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsstp.ref.Append("labels"))
}

// Location returns a reference to field location of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nsstp.ref.Append("location"))
}

// Name returns a reference to field name of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsstp.ref.Append("name"))
}

// Project returns a reference to field project of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nsstp.ref.Append("project"))
}

// TerraformLabels returns a reference to field terraform_labels of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsstp.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_network_security_server_tls_policy.
func (nsstp networkSecurityServerTlsPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nsstp.ref.Append("update_time"))
}

func (nsstp networkSecurityServerTlsPolicyAttributes) MtlsPolicy() terra.ListValue[networksecurityservertlspolicy.MtlsPolicyAttributes] {
	return terra.ReferenceAsList[networksecurityservertlspolicy.MtlsPolicyAttributes](nsstp.ref.Append("mtls_policy"))
}

func (nsstp networkSecurityServerTlsPolicyAttributes) ServerCertificate() terra.ListValue[networksecurityservertlspolicy.ServerCertificateAttributes] {
	return terra.ReferenceAsList[networksecurityservertlspolicy.ServerCertificateAttributes](nsstp.ref.Append("server_certificate"))
}

func (nsstp networkSecurityServerTlsPolicyAttributes) Timeouts() networksecurityservertlspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networksecurityservertlspolicy.TimeoutsAttributes](nsstp.ref.Append("timeouts"))
}

type networkSecurityServerTlsPolicyState struct {
	AllowOpen         bool                                                    `json:"allow_open"`
	CreateTime        string                                                  `json:"create_time"`
	Description       string                                                  `json:"description"`
	EffectiveLabels   map[string]string                                       `json:"effective_labels"`
	Id                string                                                  `json:"id"`
	Labels            map[string]string                                       `json:"labels"`
	Location          string                                                  `json:"location"`
	Name              string                                                  `json:"name"`
	Project           string                                                  `json:"project"`
	TerraformLabels   map[string]string                                       `json:"terraform_labels"`
	UpdateTime        string                                                  `json:"update_time"`
	MtlsPolicy        []networksecurityservertlspolicy.MtlsPolicyState        `json:"mtls_policy"`
	ServerCertificate []networksecurityservertlspolicy.ServerCertificateState `json:"server_certificate"`
	Timeouts          *networksecurityservertlspolicy.TimeoutsState           `json:"timeouts"`
}
