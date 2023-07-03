// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregionsslcertificate "github.com/golingon/terraproviders/googlebeta/4.71.0/computeregionsslcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionSslCertificate creates a new instance of [ComputeRegionSslCertificate].
func NewComputeRegionSslCertificate(name string, args ComputeRegionSslCertificateArgs) *ComputeRegionSslCertificate {
	return &ComputeRegionSslCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionSslCertificate)(nil)

// ComputeRegionSslCertificate represents the Terraform resource google_compute_region_ssl_certificate.
type ComputeRegionSslCertificate struct {
	Name      string
	Args      ComputeRegionSslCertificateArgs
	state     *computeRegionSslCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionSslCertificate].
func (crsc *ComputeRegionSslCertificate) Type() string {
	return "google_compute_region_ssl_certificate"
}

// LocalName returns the local name for [ComputeRegionSslCertificate].
func (crsc *ComputeRegionSslCertificate) LocalName() string {
	return crsc.Name
}

// Configuration returns the configuration (args) for [ComputeRegionSslCertificate].
func (crsc *ComputeRegionSslCertificate) Configuration() interface{} {
	return crsc.Args
}

// DependOn is used for other resources to depend on [ComputeRegionSslCertificate].
func (crsc *ComputeRegionSslCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(crsc)
}

// Dependencies returns the list of resources [ComputeRegionSslCertificate] depends_on.
func (crsc *ComputeRegionSslCertificate) Dependencies() terra.Dependencies {
	return crsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionSslCertificate].
func (crsc *ComputeRegionSslCertificate) LifecycleManagement() *terra.Lifecycle {
	return crsc.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionSslCertificate].
func (crsc *ComputeRegionSslCertificate) Attributes() computeRegionSslCertificateAttributes {
	return computeRegionSslCertificateAttributes{ref: terra.ReferenceResource(crsc)}
}

// ImportState imports the given attribute values into [ComputeRegionSslCertificate]'s state.
func (crsc *ComputeRegionSslCertificate) ImportState(av io.Reader) error {
	crsc.state = &computeRegionSslCertificateState{}
	if err := json.NewDecoder(av).Decode(crsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crsc.Type(), crsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionSslCertificate] has state.
func (crsc *ComputeRegionSslCertificate) State() (*computeRegionSslCertificateState, bool) {
	return crsc.state, crsc.state != nil
}

// StateMust returns the state for [ComputeRegionSslCertificate]. Panics if the state is nil.
func (crsc *ComputeRegionSslCertificate) StateMust() *computeRegionSslCertificateState {
	if crsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crsc.Type(), crsc.LocalName()))
	}
	return crsc.state
}

// ComputeRegionSslCertificateArgs contains the configurations for google_compute_region_ssl_certificate.
type ComputeRegionSslCertificateArgs struct {
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
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Timeouts: optional
	Timeouts *computeregionsslcertificate.Timeouts `hcl:"timeouts,block"`
}
type computeRegionSslCertificateAttributes struct {
	ref terra.Reference
}

// Certificate returns a reference to field certificate of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("certificate"))
}

// CertificateId returns a reference to field certificate_id of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) CertificateId() terra.NumberValue {
	return terra.ReferenceAsNumber(crsc.ref.Append("certificate_id"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("description"))
}

// ExpireTime returns a reference to field expire_time of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("expire_time"))
}

// Id returns a reference to field id of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("name_prefix"))
}

// PrivateKey returns a reference to field private_key of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("private_key"))
}

// Project returns a reference to field project of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_ssl_certificate.
func (crsc computeRegionSslCertificateAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crsc.ref.Append("self_link"))
}

func (crsc computeRegionSslCertificateAttributes) Timeouts() computeregionsslcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionsslcertificate.TimeoutsAttributes](crsc.ref.Append("timeouts"))
}

type computeRegionSslCertificateState struct {
	Certificate       string                                     `json:"certificate"`
	CertificateId     float64                                    `json:"certificate_id"`
	CreationTimestamp string                                     `json:"creation_timestamp"`
	Description       string                                     `json:"description"`
	ExpireTime        string                                     `json:"expire_time"`
	Id                string                                     `json:"id"`
	Name              string                                     `json:"name"`
	NamePrefix        string                                     `json:"name_prefix"`
	PrivateKey        string                                     `json:"private_key"`
	Project           string                                     `json:"project"`
	Region            string                                     `json:"region"`
	SelfLink          string                                     `json:"self_link"`
	Timeouts          *computeregionsslcertificate.TimeoutsState `json:"timeouts"`
}
