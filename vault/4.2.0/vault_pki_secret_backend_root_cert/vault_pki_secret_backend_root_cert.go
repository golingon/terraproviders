// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_pki_secret_backend_root_cert

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

// Resource represents the Terraform resource vault_pki_secret_backend_root_cert.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultPkiSecretBackendRootCertState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vpsbrc *Resource) Type() string {
	return "vault_pki_secret_backend_root_cert"
}

// LocalName returns the local name for [Resource].
func (vpsbrc *Resource) LocalName() string {
	return vpsbrc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vpsbrc *Resource) Configuration() interface{} {
	return vpsbrc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vpsbrc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vpsbrc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vpsbrc *Resource) Dependencies() terra.Dependencies {
	return vpsbrc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vpsbrc *Resource) LifecycleManagement() *terra.Lifecycle {
	return vpsbrc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vpsbrc *Resource) Attributes() vaultPkiSecretBackendRootCertAttributes {
	return vaultPkiSecretBackendRootCertAttributes{ref: terra.ReferenceResource(vpsbrc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vpsbrc *Resource) ImportState(state io.Reader) error {
	vpsbrc.state = &vaultPkiSecretBackendRootCertState{}
	if err := json.NewDecoder(state).Decode(vpsbrc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vpsbrc.Type(), vpsbrc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vpsbrc *Resource) State() (*vaultPkiSecretBackendRootCertState, bool) {
	return vpsbrc.state, vpsbrc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vpsbrc *Resource) StateMust() *vaultPkiSecretBackendRootCertState {
	if vpsbrc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vpsbrc.Type(), vpsbrc.LocalName()))
	}
	return vpsbrc.state
}

// Args contains the configurations for vault_pki_secret_backend_root_cert.
type Args struct {
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

type vaultPkiSecretBackendRootCertAttributes struct {
	ref terra.Reference
}

// AltNames returns a reference to field alt_names of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) AltNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vpsbrc.ref.Append("alt_names"))
}

// Backend returns a reference to field backend of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("backend"))
}

// Certificate returns a reference to field certificate of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("certificate"))
}

// CommonName returns a reference to field common_name of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) CommonName() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("common_name"))
}

// Country returns a reference to field country of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("country"))
}

// ExcludeCnFromSans returns a reference to field exclude_cn_from_sans of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) ExcludeCnFromSans() terra.BoolValue {
	return terra.ReferenceAsBool(vpsbrc.ref.Append("exclude_cn_from_sans"))
}

// Format returns a reference to field format of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("format"))
}

// Id returns a reference to field id of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("id"))
}

// IpSans returns a reference to field ip_sans of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) IpSans() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vpsbrc.ref.Append("ip_sans"))
}

// IssuerId returns a reference to field issuer_id of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) IssuerId() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("issuer_id"))
}

// IssuerName returns a reference to field issuer_name of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) IssuerName() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("issuer_name"))
}

// IssuingCa returns a reference to field issuing_ca of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) IssuingCa() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("issuing_ca"))
}

// KeyBits returns a reference to field key_bits of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) KeyBits() terra.NumberValue {
	return terra.ReferenceAsNumber(vpsbrc.ref.Append("key_bits"))
}

// KeyId returns a reference to field key_id of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("key_id"))
}

// KeyName returns a reference to field key_name of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("key_name"))
}

// KeyRef returns a reference to field key_ref of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) KeyRef() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("key_ref"))
}

// KeyType returns a reference to field key_type of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) KeyType() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("key_type"))
}

// Locality returns a reference to field locality of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Locality() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("locality"))
}

// ManagedKeyId returns a reference to field managed_key_id of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) ManagedKeyId() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("managed_key_id"))
}

// ManagedKeyName returns a reference to field managed_key_name of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) ManagedKeyName() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("managed_key_name"))
}

// MaxPathLength returns a reference to field max_path_length of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) MaxPathLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vpsbrc.ref.Append("max_path_length"))
}

// Namespace returns a reference to field namespace of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("namespace"))
}

// Organization returns a reference to field organization of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("organization"))
}

// OtherSans returns a reference to field other_sans of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) OtherSans() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vpsbrc.ref.Append("other_sans"))
}

// Ou returns a reference to field ou of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Ou() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("ou"))
}

// PermittedDnsDomains returns a reference to field permitted_dns_domains of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) PermittedDnsDomains() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vpsbrc.ref.Append("permitted_dns_domains"))
}

// PostalCode returns a reference to field postal_code of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("postal_code"))
}

// PrivateKeyFormat returns a reference to field private_key_format of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) PrivateKeyFormat() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("private_key_format"))
}

// Province returns a reference to field province of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Province() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("province"))
}

// SerialNumber returns a reference to field serial_number of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("serial_number"))
}

// StreetAddress returns a reference to field street_address of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) StreetAddress() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("street_address"))
}

// Ttl returns a reference to field ttl of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("ttl"))
}

// Type returns a reference to field type of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vpsbrc.ref.Append("type"))
}

// UriSans returns a reference to field uri_sans of vault_pki_secret_backend_root_cert.
func (vpsbrc vaultPkiSecretBackendRootCertAttributes) UriSans() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vpsbrc.ref.Append("uri_sans"))
}

type vaultPkiSecretBackendRootCertState struct {
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
