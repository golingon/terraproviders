// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPkiSecretBackendSign creates a new instance of [PkiSecretBackendSign].
func NewPkiSecretBackendSign(name string, args PkiSecretBackendSignArgs) *PkiSecretBackendSign {
	return &PkiSecretBackendSign{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PkiSecretBackendSign)(nil)

// PkiSecretBackendSign represents the Terraform resource vault_pki_secret_backend_sign.
type PkiSecretBackendSign struct {
	Name      string
	Args      PkiSecretBackendSignArgs
	state     *pkiSecretBackendSignState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PkiSecretBackendSign].
func (psbs *PkiSecretBackendSign) Type() string {
	return "vault_pki_secret_backend_sign"
}

// LocalName returns the local name for [PkiSecretBackendSign].
func (psbs *PkiSecretBackendSign) LocalName() string {
	return psbs.Name
}

// Configuration returns the configuration (args) for [PkiSecretBackendSign].
func (psbs *PkiSecretBackendSign) Configuration() interface{} {
	return psbs.Args
}

// DependOn is used for other resources to depend on [PkiSecretBackendSign].
func (psbs *PkiSecretBackendSign) DependOn() terra.Reference {
	return terra.ReferenceResource(psbs)
}

// Dependencies returns the list of resources [PkiSecretBackendSign] depends_on.
func (psbs *PkiSecretBackendSign) Dependencies() terra.Dependencies {
	return psbs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PkiSecretBackendSign].
func (psbs *PkiSecretBackendSign) LifecycleManagement() *terra.Lifecycle {
	return psbs.Lifecycle
}

// Attributes returns the attributes for [PkiSecretBackendSign].
func (psbs *PkiSecretBackendSign) Attributes() pkiSecretBackendSignAttributes {
	return pkiSecretBackendSignAttributes{ref: terra.ReferenceResource(psbs)}
}

// ImportState imports the given attribute values into [PkiSecretBackendSign]'s state.
func (psbs *PkiSecretBackendSign) ImportState(av io.Reader) error {
	psbs.state = &pkiSecretBackendSignState{}
	if err := json.NewDecoder(av).Decode(psbs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", psbs.Type(), psbs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PkiSecretBackendSign] has state.
func (psbs *PkiSecretBackendSign) State() (*pkiSecretBackendSignState, bool) {
	return psbs.state, psbs.state != nil
}

// StateMust returns the state for [PkiSecretBackendSign]. Panics if the state is nil.
func (psbs *PkiSecretBackendSign) StateMust() *pkiSecretBackendSignState {
	if psbs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", psbs.Type(), psbs.LocalName()))
	}
	return psbs.state
}

// PkiSecretBackendSignArgs contains the configurations for vault_pki_secret_backend_sign.
type PkiSecretBackendSignArgs struct {
	// AltNames: list of string, optional
	AltNames terra.ListValue[terra.StringValue] `hcl:"alt_names,attr"`
	// AutoRenew: bool, optional
	AutoRenew terra.BoolValue `hcl:"auto_renew,attr"`
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// CommonName: string, required
	CommonName terra.StringValue `hcl:"common_name,attr" validate:"required"`
	// Csr: string, required
	Csr terra.StringValue `hcl:"csr,attr" validate:"required"`
	// ExcludeCnFromSans: bool, optional
	ExcludeCnFromSans terra.BoolValue `hcl:"exclude_cn_from_sans,attr"`
	// Format: string, optional
	Format terra.StringValue `hcl:"format,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpSans: list of string, optional
	IpSans terra.ListValue[terra.StringValue] `hcl:"ip_sans,attr"`
	// IssuerRef: string, optional
	IssuerRef terra.StringValue `hcl:"issuer_ref,attr"`
	// MinSecondsRemaining: number, optional
	MinSecondsRemaining terra.NumberValue `hcl:"min_seconds_remaining,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// OtherSans: list of string, optional
	OtherSans terra.ListValue[terra.StringValue] `hcl:"other_sans,attr"`
	// Ttl: string, optional
	Ttl terra.StringValue `hcl:"ttl,attr"`
	// UriSans: list of string, optional
	UriSans terra.ListValue[terra.StringValue] `hcl:"uri_sans,attr"`
}
type pkiSecretBackendSignAttributes struct {
	ref terra.Reference
}

// AltNames returns a reference to field alt_names of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) AltNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbs.ref.Append("alt_names"))
}

// AutoRenew returns a reference to field auto_renew of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) AutoRenew() terra.BoolValue {
	return terra.ReferenceAsBool(psbs.ref.Append("auto_renew"))
}

// Backend returns a reference to field backend of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("backend"))
}

// CaChain returns a reference to field ca_chain of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) CaChain() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbs.ref.Append("ca_chain"))
}

// Certificate returns a reference to field certificate of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("certificate"))
}

// CommonName returns a reference to field common_name of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) CommonName() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("common_name"))
}

// Csr returns a reference to field csr of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) Csr() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("csr"))
}

// ExcludeCnFromSans returns a reference to field exclude_cn_from_sans of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) ExcludeCnFromSans() terra.BoolValue {
	return terra.ReferenceAsBool(psbs.ref.Append("exclude_cn_from_sans"))
}

// Expiration returns a reference to field expiration of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) Expiration() terra.NumberValue {
	return terra.ReferenceAsNumber(psbs.ref.Append("expiration"))
}

// Format returns a reference to field format of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("format"))
}

// Id returns a reference to field id of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("id"))
}

// IpSans returns a reference to field ip_sans of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) IpSans() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbs.ref.Append("ip_sans"))
}

// IssuerRef returns a reference to field issuer_ref of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) IssuerRef() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("issuer_ref"))
}

// IssuingCa returns a reference to field issuing_ca of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) IssuingCa() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("issuing_ca"))
}

// MinSecondsRemaining returns a reference to field min_seconds_remaining of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) MinSecondsRemaining() terra.NumberValue {
	return terra.ReferenceAsNumber(psbs.ref.Append("min_seconds_remaining"))
}

// Name returns a reference to field name of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("namespace"))
}

// OtherSans returns a reference to field other_sans of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) OtherSans() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbs.ref.Append("other_sans"))
}

// RenewPending returns a reference to field renew_pending of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) RenewPending() terra.BoolValue {
	return terra.ReferenceAsBool(psbs.ref.Append("renew_pending"))
}

// Serial returns a reference to field serial of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) Serial() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("serial"))
}

// SerialNumber returns a reference to field serial_number of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("serial_number"))
}

// Ttl returns a reference to field ttl of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(psbs.ref.Append("ttl"))
}

// UriSans returns a reference to field uri_sans of vault_pki_secret_backend_sign.
func (psbs pkiSecretBackendSignAttributes) UriSans() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbs.ref.Append("uri_sans"))
}

type pkiSecretBackendSignState struct {
	AltNames            []string `json:"alt_names"`
	AutoRenew           bool     `json:"auto_renew"`
	Backend             string   `json:"backend"`
	CaChain             []string `json:"ca_chain"`
	Certificate         string   `json:"certificate"`
	CommonName          string   `json:"common_name"`
	Csr                 string   `json:"csr"`
	ExcludeCnFromSans   bool     `json:"exclude_cn_from_sans"`
	Expiration          float64  `json:"expiration"`
	Format              string   `json:"format"`
	Id                  string   `json:"id"`
	IpSans              []string `json:"ip_sans"`
	IssuerRef           string   `json:"issuer_ref"`
	IssuingCa           string   `json:"issuing_ca"`
	MinSecondsRemaining float64  `json:"min_seconds_remaining"`
	Name                string   `json:"name"`
	Namespace           string   `json:"namespace"`
	OtherSans           []string `json:"other_sans"`
	RenewPending        bool     `json:"renew_pending"`
	Serial              string   `json:"serial"`
	SerialNumber        string   `json:"serial_number"`
	Ttl                 string   `json:"ttl"`
	UriSans             []string `json:"uri_sans"`
}