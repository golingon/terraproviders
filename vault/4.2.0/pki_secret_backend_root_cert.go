// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewPkiSecretBackendRootCert creates a new instance of [PkiSecretBackendRootCert].
func NewPkiSecretBackendRootCert(name string, args PkiSecretBackendRootCertArgs) *PkiSecretBackendRootCert {
	return &PkiSecretBackendRootCert{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PkiSecretBackendRootCert)(nil)

// PkiSecretBackendRootCert represents the Terraform resource vault_pki_secret_backend_root_cert.
type PkiSecretBackendRootCert struct {
	Name      string
	Args      PkiSecretBackendRootCertArgs
	state     *pkiSecretBackendRootCertState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PkiSecretBackendRootCert].
func (psbrc *PkiSecretBackendRootCert) Type() string {
	return "vault_pki_secret_backend_root_cert"
}

// LocalName returns the local name for [PkiSecretBackendRootCert].
func (psbrc *PkiSecretBackendRootCert) LocalName() string {
	return psbrc.Name
}

// Configuration returns the configuration (args) for [PkiSecretBackendRootCert].
func (psbrc *PkiSecretBackendRootCert) Configuration() interface{} {
	return psbrc.Args
}

// DependOn is used for other resources to depend on [PkiSecretBackendRootCert].
func (psbrc *PkiSecretBackendRootCert) DependOn() terra.Reference {
	return terra.ReferenceResource(psbrc)
}

// Dependencies returns the list of resources [PkiSecretBackendRootCert] depends_on.
func (psbrc *PkiSecretBackendRootCert) Dependencies() terra.Dependencies {
	return psbrc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PkiSecretBackendRootCert].
func (psbrc *PkiSecretBackendRootCert) LifecycleManagement() *terra.Lifecycle {
	return psbrc.Lifecycle
}

// Attributes returns the attributes for [PkiSecretBackendRootCert].
func (psbrc *PkiSecretBackendRootCert) Attributes() pkiSecretBackendRootCertAttributes {
	return pkiSecretBackendRootCertAttributes{ref: terra.ReferenceResource(psbrc)}
}

// ImportState imports the given attribute values into [PkiSecretBackendRootCert]'s state.
func (psbrc *PkiSecretBackendRootCert) ImportState(av io.Reader) error {
	psbrc.state = &pkiSecretBackendRootCertState{}
	if err := json.NewDecoder(av).Decode(psbrc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", psbrc.Type(), psbrc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PkiSecretBackendRootCert] has state.
func (psbrc *PkiSecretBackendRootCert) State() (*pkiSecretBackendRootCertState, bool) {
	return psbrc.state, psbrc.state != nil
}

// StateMust returns the state for [PkiSecretBackendRootCert]. Panics if the state is nil.
func (psbrc *PkiSecretBackendRootCert) StateMust() *pkiSecretBackendRootCertState {
	if psbrc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", psbrc.Type(), psbrc.LocalName()))
	}
	return psbrc.state
}

// PkiSecretBackendRootCertArgs contains the configurations for vault_pki_secret_backend_root_cert.
type PkiSecretBackendRootCertArgs struct {
	// AltNames: list of string, optional
	AltNames terra.ListValue[terra.StringValue] `hcl:"alt_names,attr"`
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// CommonName: string, required
	CommonName terra.StringValue `hcl:"common_name,attr" validate:"required"`
	// Country: string, optional
	Country terra.StringValue `hcl:"country,attr"`
	// ExcludeCnFromSans: bool, optional
	ExcludeCnFromSans terra.BoolValue `hcl:"exclude_cn_from_sans,attr"`
	// Format: string, optional
	Format terra.StringValue `hcl:"format,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpSans: list of string, optional
	IpSans terra.ListValue[terra.StringValue] `hcl:"ip_sans,attr"`
	// IssuerName: string, optional
	IssuerName terra.StringValue `hcl:"issuer_name,attr"`
	// KeyBits: number, optional
	KeyBits terra.NumberValue `hcl:"key_bits,attr"`
	// KeyName: string, optional
	KeyName terra.StringValue `hcl:"key_name,attr"`
	// KeyRef: string, optional
	KeyRef terra.StringValue `hcl:"key_ref,attr"`
	// KeyType: string, optional
	KeyType terra.StringValue `hcl:"key_type,attr"`
	// Locality: string, optional
	Locality terra.StringValue `hcl:"locality,attr"`
	// ManagedKeyId: string, optional
	ManagedKeyId terra.StringValue `hcl:"managed_key_id,attr"`
	// ManagedKeyName: string, optional
	ManagedKeyName terra.StringValue `hcl:"managed_key_name,attr"`
	// MaxPathLength: number, optional
	MaxPathLength terra.NumberValue `hcl:"max_path_length,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Organization: string, optional
	Organization terra.StringValue `hcl:"organization,attr"`
	// OtherSans: list of string, optional
	OtherSans terra.ListValue[terra.StringValue] `hcl:"other_sans,attr"`
	// Ou: string, optional
	Ou terra.StringValue `hcl:"ou,attr"`
	// PermittedDnsDomains: list of string, optional
	PermittedDnsDomains terra.ListValue[terra.StringValue] `hcl:"permitted_dns_domains,attr"`
	// PostalCode: string, optional
	PostalCode terra.StringValue `hcl:"postal_code,attr"`
	// PrivateKeyFormat: string, optional
	PrivateKeyFormat terra.StringValue `hcl:"private_key_format,attr"`
	// Province: string, optional
	Province terra.StringValue `hcl:"province,attr"`
	// StreetAddress: string, optional
	StreetAddress terra.StringValue `hcl:"street_address,attr"`
	// Ttl: string, optional
	Ttl terra.StringValue `hcl:"ttl,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// UriSans: list of string, optional
	UriSans terra.ListValue[terra.StringValue] `hcl:"uri_sans,attr"`
}
type pkiSecretBackendRootCertAttributes struct {
	ref terra.Reference
}

// AltNames returns a reference to field alt_names of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) AltNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbrc.ref.Append("alt_names"))
}

// Backend returns a reference to field backend of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("backend"))
}

// Certificate returns a reference to field certificate of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("certificate"))
}

// CommonName returns a reference to field common_name of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) CommonName() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("common_name"))
}

// Country returns a reference to field country of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("country"))
}

// ExcludeCnFromSans returns a reference to field exclude_cn_from_sans of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) ExcludeCnFromSans() terra.BoolValue {
	return terra.ReferenceAsBool(psbrc.ref.Append("exclude_cn_from_sans"))
}

// Format returns a reference to field format of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("format"))
}

// Id returns a reference to field id of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("id"))
}

// IpSans returns a reference to field ip_sans of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) IpSans() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbrc.ref.Append("ip_sans"))
}

// IssuerId returns a reference to field issuer_id of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) IssuerId() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("issuer_id"))
}

// IssuerName returns a reference to field issuer_name of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) IssuerName() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("issuer_name"))
}

// IssuingCa returns a reference to field issuing_ca of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) IssuingCa() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("issuing_ca"))
}

// KeyBits returns a reference to field key_bits of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) KeyBits() terra.NumberValue {
	return terra.ReferenceAsNumber(psbrc.ref.Append("key_bits"))
}

// KeyId returns a reference to field key_id of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("key_id"))
}

// KeyName returns a reference to field key_name of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("key_name"))
}

// KeyRef returns a reference to field key_ref of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) KeyRef() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("key_ref"))
}

// KeyType returns a reference to field key_type of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) KeyType() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("key_type"))
}

// Locality returns a reference to field locality of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Locality() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("locality"))
}

// ManagedKeyId returns a reference to field managed_key_id of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) ManagedKeyId() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("managed_key_id"))
}

// ManagedKeyName returns a reference to field managed_key_name of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) ManagedKeyName() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("managed_key_name"))
}

// MaxPathLength returns a reference to field max_path_length of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) MaxPathLength() terra.NumberValue {
	return terra.ReferenceAsNumber(psbrc.ref.Append("max_path_length"))
}

// Namespace returns a reference to field namespace of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("namespace"))
}

// Organization returns a reference to field organization of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("organization"))
}

// OtherSans returns a reference to field other_sans of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) OtherSans() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbrc.ref.Append("other_sans"))
}

// Ou returns a reference to field ou of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Ou() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("ou"))
}

// PermittedDnsDomains returns a reference to field permitted_dns_domains of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) PermittedDnsDomains() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbrc.ref.Append("permitted_dns_domains"))
}

// PostalCode returns a reference to field postal_code of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("postal_code"))
}

// PrivateKeyFormat returns a reference to field private_key_format of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) PrivateKeyFormat() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("private_key_format"))
}

// Province returns a reference to field province of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Province() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("province"))
}

// SerialNumber returns a reference to field serial_number of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("serial_number"))
}

// StreetAddress returns a reference to field street_address of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) StreetAddress() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("street_address"))
}

// Ttl returns a reference to field ttl of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("ttl"))
}

// Type returns a reference to field type of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(psbrc.ref.Append("type"))
}

// UriSans returns a reference to field uri_sans of vault_pki_secret_backend_root_cert.
func (psbrc pkiSecretBackendRootCertAttributes) UriSans() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbrc.ref.Append("uri_sans"))
}

type pkiSecretBackendRootCertState struct {
	AltNames            []string `json:"alt_names"`
	Backend             string   `json:"backend"`
	Certificate         string   `json:"certificate"`
	CommonName          string   `json:"common_name"`
	Country             string   `json:"country"`
	ExcludeCnFromSans   bool     `json:"exclude_cn_from_sans"`
	Format              string   `json:"format"`
	Id                  string   `json:"id"`
	IpSans              []string `json:"ip_sans"`
	IssuerId            string   `json:"issuer_id"`
	IssuerName          string   `json:"issuer_name"`
	IssuingCa           string   `json:"issuing_ca"`
	KeyBits             float64  `json:"key_bits"`
	KeyId               string   `json:"key_id"`
	KeyName             string   `json:"key_name"`
	KeyRef              string   `json:"key_ref"`
	KeyType             string   `json:"key_type"`
	Locality            string   `json:"locality"`
	ManagedKeyId        string   `json:"managed_key_id"`
	ManagedKeyName      string   `json:"managed_key_name"`
	MaxPathLength       float64  `json:"max_path_length"`
	Namespace           string   `json:"namespace"`
	Organization        string   `json:"organization"`
	OtherSans           []string `json:"other_sans"`
	Ou                  string   `json:"ou"`
	PermittedDnsDomains []string `json:"permitted_dns_domains"`
	PostalCode          string   `json:"postal_code"`
	PrivateKeyFormat    string   `json:"private_key_format"`
	Province            string   `json:"province"`
	SerialNumber        string   `json:"serial_number"`
	StreetAddress       string   `json:"street_address"`
	Ttl                 string   `json:"ttl"`
	Type                string   `json:"type"`
	UriSans             []string `json:"uri_sans"`
}
