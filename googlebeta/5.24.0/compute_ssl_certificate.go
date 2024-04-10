// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computesslcertificate "github.com/golingon/terraproviders/googlebeta/5.24.0/computesslcertificate"
	"io"
)

// NewComputeSslCertificate creates a new instance of [ComputeSslCertificate].
func NewComputeSslCertificate(name string, args ComputeSslCertificateArgs) *ComputeSslCertificate {
	return &ComputeSslCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSslCertificate)(nil)

// ComputeSslCertificate represents the Terraform resource google_compute_ssl_certificate.
type ComputeSslCertificate struct {
	Name      string
	Args      ComputeSslCertificateArgs
	state     *computeSslCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSslCertificate].
func (csc *ComputeSslCertificate) Type() string {
	return "google_compute_ssl_certificate"
}

// LocalName returns the local name for [ComputeSslCertificate].
func (csc *ComputeSslCertificate) LocalName() string {
	return csc.Name
}

// Configuration returns the configuration (args) for [ComputeSslCertificate].
func (csc *ComputeSslCertificate) Configuration() interface{} {
	return csc.Args
}

// DependOn is used for other resources to depend on [ComputeSslCertificate].
func (csc *ComputeSslCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(csc)
}

// Dependencies returns the list of resources [ComputeSslCertificate] depends_on.
func (csc *ComputeSslCertificate) Dependencies() terra.Dependencies {
	return csc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSslCertificate].
func (csc *ComputeSslCertificate) LifecycleManagement() *terra.Lifecycle {
	return csc.Lifecycle
}

// Attributes returns the attributes for [ComputeSslCertificate].
func (csc *ComputeSslCertificate) Attributes() computeSslCertificateAttributes {
	return computeSslCertificateAttributes{ref: terra.ReferenceResource(csc)}
}

// ImportState imports the given attribute values into [ComputeSslCertificate]'s state.
func (csc *ComputeSslCertificate) ImportState(av io.Reader) error {
	csc.state = &computeSslCertificateState{}
	if err := json.NewDecoder(av).Decode(csc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csc.Type(), csc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSslCertificate] has state.
func (csc *ComputeSslCertificate) State() (*computeSslCertificateState, bool) {
	return csc.state, csc.state != nil
}

// StateMust returns the state for [ComputeSslCertificate]. Panics if the state is nil.
func (csc *ComputeSslCertificate) StateMust() *computeSslCertificateState {
	if csc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csc.Type(), csc.LocalName()))
	}
	return csc.state
}

// ComputeSslCertificateArgs contains the configurations for google_compute_ssl_certificate.
type ComputeSslCertificateArgs struct {
	// Certificate: string, required
	Certificate terra.StringValue `hcl:"certificate,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// PrivateKey: string, required
	PrivateKey terra.StringValue `hcl:"private_key,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computesslcertificate.Timeouts `hcl:"timeouts,block"`
}
type computeSslCertificateAttributes struct {
	ref terra.Reference
}

// Certificate returns a reference to field certificate of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("certificate"))
}

// CertificateId returns a reference to field certificate_id of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) CertificateId() terra.NumberValue {
	return terra.ReferenceAsNumber(csc.ref.Append("certificate_id"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("description"))
}

// ExpireTime returns a reference to field expire_time of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("expire_time"))
}

// Id returns a reference to field id of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("name_prefix"))
}

// PrivateKey returns a reference to field private_key of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("private_key"))
}

// Project returns a reference to field project of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_ssl_certificate.
func (csc computeSslCertificateAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("self_link"))
}

func (csc computeSslCertificateAttributes) Timeouts() computesslcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computesslcertificate.TimeoutsAttributes](csc.ref.Append("timeouts"))
}

type computeSslCertificateState struct {
	Certificate       string                               `json:"certificate"`
	CertificateId     float64                              `json:"certificate_id"`
	CreationTimestamp string                               `json:"creation_timestamp"`
	Description       string                               `json:"description"`
	ExpireTime        string                               `json:"expire_time"`
	Id                string                               `json:"id"`
	Name              string                               `json:"name"`
	NamePrefix        string                               `json:"name_prefix"`
	PrivateKey        string                               `json:"private_key"`
	Project           string                               `json:"project"`
	SelfLink          string                               `json:"self_link"`
	Timeouts          *computesslcertificate.TimeoutsState `json:"timeouts"`
}
