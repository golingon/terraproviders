// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	attestationprovider "github.com/golingon/terraproviders/azurerm/3.63.0/attestationprovider"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAttestationProvider creates a new instance of [AttestationProvider].
func NewAttestationProvider(name string, args AttestationProviderArgs) *AttestationProvider {
	return &AttestationProvider{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AttestationProvider)(nil)

// AttestationProvider represents the Terraform resource azurerm_attestation_provider.
type AttestationProvider struct {
	Name      string
	Args      AttestationProviderArgs
	state     *attestationProviderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AttestationProvider].
func (ap *AttestationProvider) Type() string {
	return "azurerm_attestation_provider"
}

// LocalName returns the local name for [AttestationProvider].
func (ap *AttestationProvider) LocalName() string {
	return ap.Name
}

// Configuration returns the configuration (args) for [AttestationProvider].
func (ap *AttestationProvider) Configuration() interface{} {
	return ap.Args
}

// DependOn is used for other resources to depend on [AttestationProvider].
func (ap *AttestationProvider) DependOn() terra.Reference {
	return terra.ReferenceResource(ap)
}

// Dependencies returns the list of resources [AttestationProvider] depends_on.
func (ap *AttestationProvider) Dependencies() terra.Dependencies {
	return ap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AttestationProvider].
func (ap *AttestationProvider) LifecycleManagement() *terra.Lifecycle {
	return ap.Lifecycle
}

// Attributes returns the attributes for [AttestationProvider].
func (ap *AttestationProvider) Attributes() attestationProviderAttributes {
	return attestationProviderAttributes{ref: terra.ReferenceResource(ap)}
}

// ImportState imports the given attribute values into [AttestationProvider]'s state.
func (ap *AttestationProvider) ImportState(av io.Reader) error {
	ap.state = &attestationProviderState{}
	if err := json.NewDecoder(av).Decode(ap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ap.Type(), ap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AttestationProvider] has state.
func (ap *AttestationProvider) State() (*attestationProviderState, bool) {
	return ap.state, ap.state != nil
}

// StateMust returns the state for [AttestationProvider]. Panics if the state is nil.
func (ap *AttestationProvider) StateMust() *attestationProviderState {
	if ap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ap.Type(), ap.LocalName()))
	}
	return ap.state
}

// AttestationProviderArgs contains the configurations for azurerm_attestation_provider.
type AttestationProviderArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OpenEnclavePolicyBase64: string, optional
	OpenEnclavePolicyBase64 terra.StringValue `hcl:"open_enclave_policy_base64,attr"`
	// PolicySigningCertificateData: string, optional
	PolicySigningCertificateData terra.StringValue `hcl:"policy_signing_certificate_data,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SgxEnclavePolicyBase64: string, optional
	SgxEnclavePolicyBase64 terra.StringValue `hcl:"sgx_enclave_policy_base64,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TpmPolicyBase64: string, optional
	TpmPolicyBase64 terra.StringValue `hcl:"tpm_policy_base64,attr"`
	// Policy: min=0
	Policy []attestationprovider.Policy `hcl:"policy,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *attestationprovider.Timeouts `hcl:"timeouts,block"`
}
type attestationProviderAttributes struct {
	ref terra.Reference
}

// AttestationUri returns a reference to field attestation_uri of azurerm_attestation_provider.
func (ap attestationProviderAttributes) AttestationUri() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("attestation_uri"))
}

// Id returns a reference to field id of azurerm_attestation_provider.
func (ap attestationProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_attestation_provider.
func (ap attestationProviderAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_attestation_provider.
func (ap attestationProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("name"))
}

// OpenEnclavePolicyBase64 returns a reference to field open_enclave_policy_base64 of azurerm_attestation_provider.
func (ap attestationProviderAttributes) OpenEnclavePolicyBase64() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("open_enclave_policy_base64"))
}

// PolicySigningCertificateData returns a reference to field policy_signing_certificate_data of azurerm_attestation_provider.
func (ap attestationProviderAttributes) PolicySigningCertificateData() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("policy_signing_certificate_data"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_attestation_provider.
func (ap attestationProviderAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("resource_group_name"))
}

// SgxEnclavePolicyBase64 returns a reference to field sgx_enclave_policy_base64 of azurerm_attestation_provider.
func (ap attestationProviderAttributes) SgxEnclavePolicyBase64() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("sgx_enclave_policy_base64"))
}

// Tags returns a reference to field tags of azurerm_attestation_provider.
func (ap attestationProviderAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ap.ref.Append("tags"))
}

// TpmPolicyBase64 returns a reference to field tpm_policy_base64 of azurerm_attestation_provider.
func (ap attestationProviderAttributes) TpmPolicyBase64() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("tpm_policy_base64"))
}

// TrustModel returns a reference to field trust_model of azurerm_attestation_provider.
func (ap attestationProviderAttributes) TrustModel() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("trust_model"))
}

func (ap attestationProviderAttributes) Policy() terra.ListValue[attestationprovider.PolicyAttributes] {
	return terra.ReferenceAsList[attestationprovider.PolicyAttributes](ap.ref.Append("policy"))
}

func (ap attestationProviderAttributes) Timeouts() attestationprovider.TimeoutsAttributes {
	return terra.ReferenceAsSingle[attestationprovider.TimeoutsAttributes](ap.ref.Append("timeouts"))
}

type attestationProviderState struct {
	AttestationUri               string                             `json:"attestation_uri"`
	Id                           string                             `json:"id"`
	Location                     string                             `json:"location"`
	Name                         string                             `json:"name"`
	OpenEnclavePolicyBase64      string                             `json:"open_enclave_policy_base64"`
	PolicySigningCertificateData string                             `json:"policy_signing_certificate_data"`
	ResourceGroupName            string                             `json:"resource_group_name"`
	SgxEnclavePolicyBase64       string                             `json:"sgx_enclave_policy_base64"`
	Tags                         map[string]string                  `json:"tags"`
	TpmPolicyBase64              string                             `json:"tpm_policy_base64"`
	TrustModel                   string                             `json:"trust_model"`
	Policy                       []attestationprovider.PolicyState  `json:"policy"`
	Timeouts                     *attestationprovider.TimeoutsState `json:"timeouts"`
}
