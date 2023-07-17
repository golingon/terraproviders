// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	dataprivatecacertificateauthority "github.com/golingon/terraproviders/google/4.73.1/dataprivatecacertificateauthority"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivatecaCertificateAuthority creates a new instance of [DataPrivatecaCertificateAuthority].
func NewDataPrivatecaCertificateAuthority(name string, args DataPrivatecaCertificateAuthorityArgs) *DataPrivatecaCertificateAuthority {
	return &DataPrivatecaCertificateAuthority{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivatecaCertificateAuthority)(nil)

// DataPrivatecaCertificateAuthority represents the Terraform data resource google_privateca_certificate_authority.
type DataPrivatecaCertificateAuthority struct {
	Name string
	Args DataPrivatecaCertificateAuthorityArgs
}

// DataSource returns the Terraform object type for [DataPrivatecaCertificateAuthority].
func (pca *DataPrivatecaCertificateAuthority) DataSource() string {
	return "google_privateca_certificate_authority"
}

// LocalName returns the local name for [DataPrivatecaCertificateAuthority].
func (pca *DataPrivatecaCertificateAuthority) LocalName() string {
	return pca.Name
}

// Configuration returns the configuration (args) for [DataPrivatecaCertificateAuthority].
func (pca *DataPrivatecaCertificateAuthority) Configuration() interface{} {
	return pca.Args
}

// Attributes returns the attributes for [DataPrivatecaCertificateAuthority].
func (pca *DataPrivatecaCertificateAuthority) Attributes() dataPrivatecaCertificateAuthorityAttributes {
	return dataPrivatecaCertificateAuthorityAttributes{ref: terra.ReferenceDataResource(pca)}
}

// DataPrivatecaCertificateAuthorityArgs contains the configurations for google_privateca_certificate_authority.
type DataPrivatecaCertificateAuthorityArgs struct {
	// CertificateAuthorityId: string, optional
	CertificateAuthorityId terra.StringValue `hcl:"certificate_authority_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Pool: string, optional
	Pool terra.StringValue `hcl:"pool,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// AccessUrls: min=0
	AccessUrls []dataprivatecacertificateauthority.AccessUrls `hcl:"access_urls,block" validate:"min=0"`
	// Config: min=0
	Config []dataprivatecacertificateauthority.Config `hcl:"config,block" validate:"min=0"`
	// KeySpec: min=0
	KeySpec []dataprivatecacertificateauthority.KeySpec `hcl:"key_spec,block" validate:"min=0"`
	// SubordinateConfig: min=0
	SubordinateConfig []dataprivatecacertificateauthority.SubordinateConfig `hcl:"subordinate_config,block" validate:"min=0"`
}
type dataPrivatecaCertificateAuthorityAttributes struct {
	ref terra.Reference
}

// CertificateAuthorityId returns a reference to field certificate_authority_id of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) CertificateAuthorityId() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("certificate_authority_id"))
}

// CreateTime returns a reference to field create_time of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("create_time"))
}

// DeletionProtection returns a reference to field deletion_protection of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(pca.ref.Append("deletion_protection"))
}

// DesiredState returns a reference to field desired_state of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) DesiredState() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("desired_state"))
}

// GcsBucket returns a reference to field gcs_bucket of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) GcsBucket() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("gcs_bucket"))
}

// Id returns a reference to field id of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("id"))
}

// IgnoreActiveCertificatesOnDeletion returns a reference to field ignore_active_certificates_on_deletion of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) IgnoreActiveCertificatesOnDeletion() terra.BoolValue {
	return terra.ReferenceAsBool(pca.ref.Append("ignore_active_certificates_on_deletion"))
}

// Labels returns a reference to field labels of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pca.ref.Append("labels"))
}

// Lifetime returns a reference to field lifetime of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) Lifetime() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("lifetime"))
}

// Location returns a reference to field location of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("location"))
}

// Name returns a reference to field name of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("name"))
}

// PemCaCertificate returns a reference to field pem_ca_certificate of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) PemCaCertificate() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("pem_ca_certificate"))
}

// PemCaCertificates returns a reference to field pem_ca_certificates of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) PemCaCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pca.ref.Append("pem_ca_certificates"))
}

// PemCsr returns a reference to field pem_csr of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) PemCsr() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("pem_csr"))
}

// Pool returns a reference to field pool of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) Pool() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("pool"))
}

// Project returns a reference to field project of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("project"))
}

// SkipGracePeriod returns a reference to field skip_grace_period of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) SkipGracePeriod() terra.BoolValue {
	return terra.ReferenceAsBool(pca.ref.Append("skip_grace_period"))
}

// State returns a reference to field state of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("state"))
}

// Type returns a reference to field type of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("type"))
}

// UpdateTime returns a reference to field update_time of google_privateca_certificate_authority.
func (pca dataPrivatecaCertificateAuthorityAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("update_time"))
}

func (pca dataPrivatecaCertificateAuthorityAttributes) AccessUrls() terra.ListValue[dataprivatecacertificateauthority.AccessUrlsAttributes] {
	return terra.ReferenceAsList[dataprivatecacertificateauthority.AccessUrlsAttributes](pca.ref.Append("access_urls"))
}

func (pca dataPrivatecaCertificateAuthorityAttributes) Config() terra.ListValue[dataprivatecacertificateauthority.ConfigAttributes] {
	return terra.ReferenceAsList[dataprivatecacertificateauthority.ConfigAttributes](pca.ref.Append("config"))
}

func (pca dataPrivatecaCertificateAuthorityAttributes) KeySpec() terra.ListValue[dataprivatecacertificateauthority.KeySpecAttributes] {
	return terra.ReferenceAsList[dataprivatecacertificateauthority.KeySpecAttributes](pca.ref.Append("key_spec"))
}

func (pca dataPrivatecaCertificateAuthorityAttributes) SubordinateConfig() terra.ListValue[dataprivatecacertificateauthority.SubordinateConfigAttributes] {
	return terra.ReferenceAsList[dataprivatecacertificateauthority.SubordinateConfigAttributes](pca.ref.Append("subordinate_config"))
}
