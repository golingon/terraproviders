// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	privatecacertificate "github.com/golingon/terraproviders/googlebeta/4.75.1/privatecacertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivatecaCertificate creates a new instance of [PrivatecaCertificate].
func NewPrivatecaCertificate(name string, args PrivatecaCertificateArgs) *PrivatecaCertificate {
	return &PrivatecaCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivatecaCertificate)(nil)

// PrivatecaCertificate represents the Terraform resource google_privateca_certificate.
type PrivatecaCertificate struct {
	Name      string
	Args      PrivatecaCertificateArgs
	state     *privatecaCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivatecaCertificate].
func (pc *PrivatecaCertificate) Type() string {
	return "google_privateca_certificate"
}

// LocalName returns the local name for [PrivatecaCertificate].
func (pc *PrivatecaCertificate) LocalName() string {
	return pc.Name
}

// Configuration returns the configuration (args) for [PrivatecaCertificate].
func (pc *PrivatecaCertificate) Configuration() interface{} {
	return pc.Args
}

// DependOn is used for other resources to depend on [PrivatecaCertificate].
func (pc *PrivatecaCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(pc)
}

// Dependencies returns the list of resources [PrivatecaCertificate] depends_on.
func (pc *PrivatecaCertificate) Dependencies() terra.Dependencies {
	return pc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivatecaCertificate].
func (pc *PrivatecaCertificate) LifecycleManagement() *terra.Lifecycle {
	return pc.Lifecycle
}

// Attributes returns the attributes for [PrivatecaCertificate].
func (pc *PrivatecaCertificate) Attributes() privatecaCertificateAttributes {
	return privatecaCertificateAttributes{ref: terra.ReferenceResource(pc)}
}

// ImportState imports the given attribute values into [PrivatecaCertificate]'s state.
func (pc *PrivatecaCertificate) ImportState(av io.Reader) error {
	pc.state = &privatecaCertificateState{}
	if err := json.NewDecoder(av).Decode(pc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pc.Type(), pc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivatecaCertificate] has state.
func (pc *PrivatecaCertificate) State() (*privatecaCertificateState, bool) {
	return pc.state, pc.state != nil
}

// StateMust returns the state for [PrivatecaCertificate]. Panics if the state is nil.
func (pc *PrivatecaCertificate) StateMust() *privatecaCertificateState {
	if pc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pc.Type(), pc.LocalName()))
	}
	return pc.state
}

// PrivatecaCertificateArgs contains the configurations for google_privateca_certificate.
type PrivatecaCertificateArgs struct {
	// CertificateAuthority: string, optional
	CertificateAuthority terra.StringValue `hcl:"certificate_authority,attr"`
	// CertificateTemplate: string, optional
	CertificateTemplate terra.StringValue `hcl:"certificate_template,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Lifetime: string, optional
	Lifetime terra.StringValue `hcl:"lifetime,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PemCsr: string, optional
	PemCsr terra.StringValue `hcl:"pem_csr,attr"`
	// Pool: string, required
	Pool terra.StringValue `hcl:"pool,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// CertificateDescription: min=0
	CertificateDescription []privatecacertificate.CertificateDescription `hcl:"certificate_description,block" validate:"min=0"`
	// RevocationDetails: min=0
	RevocationDetails []privatecacertificate.RevocationDetails `hcl:"revocation_details,block" validate:"min=0"`
	// Config: optional
	Config *privatecacertificate.Config `hcl:"config,block"`
	// Timeouts: optional
	Timeouts *privatecacertificate.Timeouts `hcl:"timeouts,block"`
}
type privatecaCertificateAttributes struct {
	ref terra.Reference
}

// CertificateAuthority returns a reference to field certificate_authority of google_privateca_certificate.
func (pc privatecaCertificateAttributes) CertificateAuthority() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("certificate_authority"))
}

// CertificateTemplate returns a reference to field certificate_template of google_privateca_certificate.
func (pc privatecaCertificateAttributes) CertificateTemplate() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("certificate_template"))
}

// CreateTime returns a reference to field create_time of google_privateca_certificate.
func (pc privatecaCertificateAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("create_time"))
}

// Id returns a reference to field id of google_privateca_certificate.
func (pc privatecaCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("id"))
}

// IssuerCertificateAuthority returns a reference to field issuer_certificate_authority of google_privateca_certificate.
func (pc privatecaCertificateAttributes) IssuerCertificateAuthority() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("issuer_certificate_authority"))
}

// Labels returns a reference to field labels of google_privateca_certificate.
func (pc privatecaCertificateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pc.ref.Append("labels"))
}

// Lifetime returns a reference to field lifetime of google_privateca_certificate.
func (pc privatecaCertificateAttributes) Lifetime() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("lifetime"))
}

// Location returns a reference to field location of google_privateca_certificate.
func (pc privatecaCertificateAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("location"))
}

// Name returns a reference to field name of google_privateca_certificate.
func (pc privatecaCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("name"))
}

// PemCertificate returns a reference to field pem_certificate of google_privateca_certificate.
func (pc privatecaCertificateAttributes) PemCertificate() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("pem_certificate"))
}

// PemCertificateChain returns a reference to field pem_certificate_chain of google_privateca_certificate.
func (pc privatecaCertificateAttributes) PemCertificateChain() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("pem_certificate_chain"))
}

// PemCertificates returns a reference to field pem_certificates of google_privateca_certificate.
func (pc privatecaCertificateAttributes) PemCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("pem_certificates"))
}

// PemCsr returns a reference to field pem_csr of google_privateca_certificate.
func (pc privatecaCertificateAttributes) PemCsr() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("pem_csr"))
}

// Pool returns a reference to field pool of google_privateca_certificate.
func (pc privatecaCertificateAttributes) Pool() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("pool"))
}

// Project returns a reference to field project of google_privateca_certificate.
func (pc privatecaCertificateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_privateca_certificate.
func (pc privatecaCertificateAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("update_time"))
}

func (pc privatecaCertificateAttributes) CertificateDescription() terra.ListValue[privatecacertificate.CertificateDescriptionAttributes] {
	return terra.ReferenceAsList[privatecacertificate.CertificateDescriptionAttributes](pc.ref.Append("certificate_description"))
}

func (pc privatecaCertificateAttributes) RevocationDetails() terra.ListValue[privatecacertificate.RevocationDetailsAttributes] {
	return terra.ReferenceAsList[privatecacertificate.RevocationDetailsAttributes](pc.ref.Append("revocation_details"))
}

func (pc privatecaCertificateAttributes) Config() terra.ListValue[privatecacertificate.ConfigAttributes] {
	return terra.ReferenceAsList[privatecacertificate.ConfigAttributes](pc.ref.Append("config"))
}

func (pc privatecaCertificateAttributes) Timeouts() privatecacertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatecacertificate.TimeoutsAttributes](pc.ref.Append("timeouts"))
}

type privatecaCertificateState struct {
	CertificateAuthority       string                                             `json:"certificate_authority"`
	CertificateTemplate        string                                             `json:"certificate_template"`
	CreateTime                 string                                             `json:"create_time"`
	Id                         string                                             `json:"id"`
	IssuerCertificateAuthority string                                             `json:"issuer_certificate_authority"`
	Labels                     map[string]string                                  `json:"labels"`
	Lifetime                   string                                             `json:"lifetime"`
	Location                   string                                             `json:"location"`
	Name                       string                                             `json:"name"`
	PemCertificate             string                                             `json:"pem_certificate"`
	PemCertificateChain        []string                                           `json:"pem_certificate_chain"`
	PemCertificates            []string                                           `json:"pem_certificates"`
	PemCsr                     string                                             `json:"pem_csr"`
	Pool                       string                                             `json:"pool"`
	Project                    string                                             `json:"project"`
	UpdateTime                 string                                             `json:"update_time"`
	CertificateDescription     []privatecacertificate.CertificateDescriptionState `json:"certificate_description"`
	RevocationDetails          []privatecacertificate.RevocationDetailsState      `json:"revocation_details"`
	Config                     []privatecacertificate.ConfigState                 `json:"config"`
	Timeouts                   *privatecacertificate.TimeoutsState                `json:"timeouts"`
}
