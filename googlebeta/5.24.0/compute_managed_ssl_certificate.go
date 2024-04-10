// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computemanagedsslcertificate "github.com/golingon/terraproviders/googlebeta/5.24.0/computemanagedsslcertificate"
	"io"
)

// NewComputeManagedSslCertificate creates a new instance of [ComputeManagedSslCertificate].
func NewComputeManagedSslCertificate(name string, args ComputeManagedSslCertificateArgs) *ComputeManagedSslCertificate {
	return &ComputeManagedSslCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeManagedSslCertificate)(nil)

// ComputeManagedSslCertificate represents the Terraform resource google_compute_managed_ssl_certificate.
type ComputeManagedSslCertificate struct {
	Name      string
	Args      ComputeManagedSslCertificateArgs
	state     *computeManagedSslCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeManagedSslCertificate].
func (cmsc *ComputeManagedSslCertificate) Type() string {
	return "google_compute_managed_ssl_certificate"
}

// LocalName returns the local name for [ComputeManagedSslCertificate].
func (cmsc *ComputeManagedSslCertificate) LocalName() string {
	return cmsc.Name
}

// Configuration returns the configuration (args) for [ComputeManagedSslCertificate].
func (cmsc *ComputeManagedSslCertificate) Configuration() interface{} {
	return cmsc.Args
}

// DependOn is used for other resources to depend on [ComputeManagedSslCertificate].
func (cmsc *ComputeManagedSslCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(cmsc)
}

// Dependencies returns the list of resources [ComputeManagedSslCertificate] depends_on.
func (cmsc *ComputeManagedSslCertificate) Dependencies() terra.Dependencies {
	return cmsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeManagedSslCertificate].
func (cmsc *ComputeManagedSslCertificate) LifecycleManagement() *terra.Lifecycle {
	return cmsc.Lifecycle
}

// Attributes returns the attributes for [ComputeManagedSslCertificate].
func (cmsc *ComputeManagedSslCertificate) Attributes() computeManagedSslCertificateAttributes {
	return computeManagedSslCertificateAttributes{ref: terra.ReferenceResource(cmsc)}
}

// ImportState imports the given attribute values into [ComputeManagedSslCertificate]'s state.
func (cmsc *ComputeManagedSslCertificate) ImportState(av io.Reader) error {
	cmsc.state = &computeManagedSslCertificateState{}
	if err := json.NewDecoder(av).Decode(cmsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmsc.Type(), cmsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeManagedSslCertificate] has state.
func (cmsc *ComputeManagedSslCertificate) State() (*computeManagedSslCertificateState, bool) {
	return cmsc.state, cmsc.state != nil
}

// StateMust returns the state for [ComputeManagedSslCertificate]. Panics if the state is nil.
func (cmsc *ComputeManagedSslCertificate) StateMust() *computeManagedSslCertificateState {
	if cmsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmsc.Type(), cmsc.LocalName()))
	}
	return cmsc.state
}

// ComputeManagedSslCertificateArgs contains the configurations for google_compute_managed_ssl_certificate.
type ComputeManagedSslCertificateArgs struct {
	// CertificateId: number, optional
	CertificateId terra.NumberValue `hcl:"certificate_id,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Managed: optional
	Managed *computemanagedsslcertificate.Managed `hcl:"managed,block"`
	// Timeouts: optional
	Timeouts *computemanagedsslcertificate.Timeouts `hcl:"timeouts,block"`
}
type computeManagedSslCertificateAttributes struct {
	ref terra.Reference
}

// CertificateId returns a reference to field certificate_id of google_compute_managed_ssl_certificate.
func (cmsc computeManagedSslCertificateAttributes) CertificateId() terra.NumberValue {
	return terra.ReferenceAsNumber(cmsc.ref.Append("certificate_id"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_managed_ssl_certificate.
func (cmsc computeManagedSslCertificateAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cmsc.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_managed_ssl_certificate.
func (cmsc computeManagedSslCertificateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cmsc.ref.Append("description"))
}

// ExpireTime returns a reference to field expire_time of google_compute_managed_ssl_certificate.
func (cmsc computeManagedSslCertificateAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(cmsc.ref.Append("expire_time"))
}

// Id returns a reference to field id of google_compute_managed_ssl_certificate.
func (cmsc computeManagedSslCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmsc.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_managed_ssl_certificate.
func (cmsc computeManagedSslCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmsc.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_managed_ssl_certificate.
func (cmsc computeManagedSslCertificateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cmsc.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_managed_ssl_certificate.
func (cmsc computeManagedSslCertificateAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cmsc.ref.Append("self_link"))
}

// SubjectAlternativeNames returns a reference to field subject_alternative_names of google_compute_managed_ssl_certificate.
func (cmsc computeManagedSslCertificateAttributes) SubjectAlternativeNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cmsc.ref.Append("subject_alternative_names"))
}

// Type returns a reference to field type of google_compute_managed_ssl_certificate.
func (cmsc computeManagedSslCertificateAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cmsc.ref.Append("type"))
}

func (cmsc computeManagedSslCertificateAttributes) Managed() terra.ListValue[computemanagedsslcertificate.ManagedAttributes] {
	return terra.ReferenceAsList[computemanagedsslcertificate.ManagedAttributes](cmsc.ref.Append("managed"))
}

func (cmsc computeManagedSslCertificateAttributes) Timeouts() computemanagedsslcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computemanagedsslcertificate.TimeoutsAttributes](cmsc.ref.Append("timeouts"))
}

type computeManagedSslCertificateState struct {
	CertificateId           float64                                     `json:"certificate_id"`
	CreationTimestamp       string                                      `json:"creation_timestamp"`
	Description             string                                      `json:"description"`
	ExpireTime              string                                      `json:"expire_time"`
	Id                      string                                      `json:"id"`
	Name                    string                                      `json:"name"`
	Project                 string                                      `json:"project"`
	SelfLink                string                                      `json:"self_link"`
	SubjectAlternativeNames []string                                    `json:"subject_alternative_names"`
	Type                    string                                      `json:"type"`
	Managed                 []computemanagedsslcertificate.ManagedState `json:"managed"`
	Timeouts                *computemanagedsslcertificate.TimeoutsState `json:"timeouts"`
}
