// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	privatecacertificateauthority "github.com/golingon/terraproviders/googlebeta/5.24.0/privatecacertificateauthority"
	"io"
)

// NewPrivatecaCertificateAuthority creates a new instance of [PrivatecaCertificateAuthority].
func NewPrivatecaCertificateAuthority(name string, args PrivatecaCertificateAuthorityArgs) *PrivatecaCertificateAuthority {
	return &PrivatecaCertificateAuthority{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivatecaCertificateAuthority)(nil)

// PrivatecaCertificateAuthority represents the Terraform resource google_privateca_certificate_authority.
type PrivatecaCertificateAuthority struct {
	Name      string
	Args      PrivatecaCertificateAuthorityArgs
	state     *privatecaCertificateAuthorityState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivatecaCertificateAuthority].
func (pca *PrivatecaCertificateAuthority) Type() string {
	return "google_privateca_certificate_authority"
}

// LocalName returns the local name for [PrivatecaCertificateAuthority].
func (pca *PrivatecaCertificateAuthority) LocalName() string {
	return pca.Name
}

// Configuration returns the configuration (args) for [PrivatecaCertificateAuthority].
func (pca *PrivatecaCertificateAuthority) Configuration() interface{} {
	return pca.Args
}

// DependOn is used for other resources to depend on [PrivatecaCertificateAuthority].
func (pca *PrivatecaCertificateAuthority) DependOn() terra.Reference {
	return terra.ReferenceResource(pca)
}

// Dependencies returns the list of resources [PrivatecaCertificateAuthority] depends_on.
func (pca *PrivatecaCertificateAuthority) Dependencies() terra.Dependencies {
	return pca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivatecaCertificateAuthority].
func (pca *PrivatecaCertificateAuthority) LifecycleManagement() *terra.Lifecycle {
	return pca.Lifecycle
}

// Attributes returns the attributes for [PrivatecaCertificateAuthority].
func (pca *PrivatecaCertificateAuthority) Attributes() privatecaCertificateAuthorityAttributes {
	return privatecaCertificateAuthorityAttributes{ref: terra.ReferenceResource(pca)}
}

// ImportState imports the given attribute values into [PrivatecaCertificateAuthority]'s state.
func (pca *PrivatecaCertificateAuthority) ImportState(av io.Reader) error {
	pca.state = &privatecaCertificateAuthorityState{}
	if err := json.NewDecoder(av).Decode(pca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pca.Type(), pca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivatecaCertificateAuthority] has state.
func (pca *PrivatecaCertificateAuthority) State() (*privatecaCertificateAuthorityState, bool) {
	return pca.state, pca.state != nil
}

// StateMust returns the state for [PrivatecaCertificateAuthority]. Panics if the state is nil.
func (pca *PrivatecaCertificateAuthority) StateMust() *privatecaCertificateAuthorityState {
	if pca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pca.Type(), pca.LocalName()))
	}
	return pca.state
}

// PrivatecaCertificateAuthorityArgs contains the configurations for google_privateca_certificate_authority.
type PrivatecaCertificateAuthorityArgs struct {
	// CertificateAuthorityId: string, required
	CertificateAuthorityId terra.StringValue `hcl:"certificate_authority_id,attr" validate:"required"`
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// DesiredState: string, optional
	DesiredState terra.StringValue `hcl:"desired_state,attr"`
	// GcsBucket: string, optional
	GcsBucket terra.StringValue `hcl:"gcs_bucket,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreActiveCertificatesOnDeletion: bool, optional
	IgnoreActiveCertificatesOnDeletion terra.BoolValue `hcl:"ignore_active_certificates_on_deletion,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Lifetime: string, optional
	Lifetime terra.StringValue `hcl:"lifetime,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// PemCaCertificate: string, optional
	PemCaCertificate terra.StringValue `hcl:"pem_ca_certificate,attr"`
	// Pool: string, required
	Pool terra.StringValue `hcl:"pool,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SkipGracePeriod: bool, optional
	SkipGracePeriod terra.BoolValue `hcl:"skip_grace_period,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// AccessUrls: min=0
	AccessUrls []privatecacertificateauthority.AccessUrls `hcl:"access_urls,block" validate:"min=0"`
	// Config: required
	Config *privatecacertificateauthority.Config `hcl:"config,block" validate:"required"`
	// KeySpec: required
	KeySpec *privatecacertificateauthority.KeySpec `hcl:"key_spec,block" validate:"required"`
	// SubordinateConfig: optional
	SubordinateConfig *privatecacertificateauthority.SubordinateConfig `hcl:"subordinate_config,block"`
	// Timeouts: optional
	Timeouts *privatecacertificateauthority.Timeouts `hcl:"timeouts,block"`
}
type privatecaCertificateAuthorityAttributes struct {
	ref terra.Reference
}

// CertificateAuthorityId returns a reference to field certificate_authority_id of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) CertificateAuthorityId() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("certificate_authority_id"))
}

// CreateTime returns a reference to field create_time of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("create_time"))
}

// DeletionProtection returns a reference to field deletion_protection of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(pca.ref.Append("deletion_protection"))
}

// DesiredState returns a reference to field desired_state of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) DesiredState() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("desired_state"))
}

// EffectiveLabels returns a reference to field effective_labels of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pca.ref.Append("effective_labels"))
}

// GcsBucket returns a reference to field gcs_bucket of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) GcsBucket() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("gcs_bucket"))
}

// Id returns a reference to field id of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("id"))
}

// IgnoreActiveCertificatesOnDeletion returns a reference to field ignore_active_certificates_on_deletion of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) IgnoreActiveCertificatesOnDeletion() terra.BoolValue {
	return terra.ReferenceAsBool(pca.ref.Append("ignore_active_certificates_on_deletion"))
}

// Labels returns a reference to field labels of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pca.ref.Append("labels"))
}

// Lifetime returns a reference to field lifetime of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) Lifetime() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("lifetime"))
}

// Location returns a reference to field location of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("location"))
}

// Name returns a reference to field name of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("name"))
}

// PemCaCertificate returns a reference to field pem_ca_certificate of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) PemCaCertificate() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("pem_ca_certificate"))
}

// PemCaCertificates returns a reference to field pem_ca_certificates of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) PemCaCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pca.ref.Append("pem_ca_certificates"))
}

// Pool returns a reference to field pool of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) Pool() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("pool"))
}

// Project returns a reference to field project of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("project"))
}

// SkipGracePeriod returns a reference to field skip_grace_period of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) SkipGracePeriod() terra.BoolValue {
	return terra.ReferenceAsBool(pca.ref.Append("skip_grace_period"))
}

// State returns a reference to field state of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("state"))
}

// TerraformLabels returns a reference to field terraform_labels of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pca.ref.Append("terraform_labels"))
}

// Type returns a reference to field type of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("type"))
}

// UpdateTime returns a reference to field update_time of google_privateca_certificate_authority.
func (pca privatecaCertificateAuthorityAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(pca.ref.Append("update_time"))
}

func (pca privatecaCertificateAuthorityAttributes) AccessUrls() terra.ListValue[privatecacertificateauthority.AccessUrlsAttributes] {
	return terra.ReferenceAsList[privatecacertificateauthority.AccessUrlsAttributes](pca.ref.Append("access_urls"))
}

func (pca privatecaCertificateAuthorityAttributes) Config() terra.ListValue[privatecacertificateauthority.ConfigAttributes] {
	return terra.ReferenceAsList[privatecacertificateauthority.ConfigAttributes](pca.ref.Append("config"))
}

func (pca privatecaCertificateAuthorityAttributes) KeySpec() terra.ListValue[privatecacertificateauthority.KeySpecAttributes] {
	return terra.ReferenceAsList[privatecacertificateauthority.KeySpecAttributes](pca.ref.Append("key_spec"))
}

func (pca privatecaCertificateAuthorityAttributes) SubordinateConfig() terra.ListValue[privatecacertificateauthority.SubordinateConfigAttributes] {
	return terra.ReferenceAsList[privatecacertificateauthority.SubordinateConfigAttributes](pca.ref.Append("subordinate_config"))
}

func (pca privatecaCertificateAuthorityAttributes) Timeouts() privatecacertificateauthority.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatecacertificateauthority.TimeoutsAttributes](pca.ref.Append("timeouts"))
}

type privatecaCertificateAuthorityState struct {
	CertificateAuthorityId             string                                                 `json:"certificate_authority_id"`
	CreateTime                         string                                                 `json:"create_time"`
	DeletionProtection                 bool                                                   `json:"deletion_protection"`
	DesiredState                       string                                                 `json:"desired_state"`
	EffectiveLabels                    map[string]string                                      `json:"effective_labels"`
	GcsBucket                          string                                                 `json:"gcs_bucket"`
	Id                                 string                                                 `json:"id"`
	IgnoreActiveCertificatesOnDeletion bool                                                   `json:"ignore_active_certificates_on_deletion"`
	Labels                             map[string]string                                      `json:"labels"`
	Lifetime                           string                                                 `json:"lifetime"`
	Location                           string                                                 `json:"location"`
	Name                               string                                                 `json:"name"`
	PemCaCertificate                   string                                                 `json:"pem_ca_certificate"`
	PemCaCertificates                  []string                                               `json:"pem_ca_certificates"`
	Pool                               string                                                 `json:"pool"`
	Project                            string                                                 `json:"project"`
	SkipGracePeriod                    bool                                                   `json:"skip_grace_period"`
	State                              string                                                 `json:"state"`
	TerraformLabels                    map[string]string                                      `json:"terraform_labels"`
	Type                               string                                                 `json:"type"`
	UpdateTime                         string                                                 `json:"update_time"`
	AccessUrls                         []privatecacertificateauthority.AccessUrlsState        `json:"access_urls"`
	Config                             []privatecacertificateauthority.ConfigState            `json:"config"`
	KeySpec                            []privatecacertificateauthority.KeySpecState           `json:"key_spec"`
	SubordinateConfig                  []privatecacertificateauthority.SubordinateConfigState `json:"subordinate_config"`
	Timeouts                           *privatecacertificateauthority.TimeoutsState           `json:"timeouts"`
}
