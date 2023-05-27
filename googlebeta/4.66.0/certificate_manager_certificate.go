// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	certificatemanagercertificate "github.com/golingon/terraproviders/googlebeta/4.66.0/certificatemanagercertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCertificateManagerCertificate creates a new instance of [CertificateManagerCertificate].
func NewCertificateManagerCertificate(name string, args CertificateManagerCertificateArgs) *CertificateManagerCertificate {
	return &CertificateManagerCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CertificateManagerCertificate)(nil)

// CertificateManagerCertificate represents the Terraform resource google_certificate_manager_certificate.
type CertificateManagerCertificate struct {
	Name      string
	Args      CertificateManagerCertificateArgs
	state     *certificateManagerCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CertificateManagerCertificate].
func (cmc *CertificateManagerCertificate) Type() string {
	return "google_certificate_manager_certificate"
}

// LocalName returns the local name for [CertificateManagerCertificate].
func (cmc *CertificateManagerCertificate) LocalName() string {
	return cmc.Name
}

// Configuration returns the configuration (args) for [CertificateManagerCertificate].
func (cmc *CertificateManagerCertificate) Configuration() interface{} {
	return cmc.Args
}

// DependOn is used for other resources to depend on [CertificateManagerCertificate].
func (cmc *CertificateManagerCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(cmc)
}

// Dependencies returns the list of resources [CertificateManagerCertificate] depends_on.
func (cmc *CertificateManagerCertificate) Dependencies() terra.Dependencies {
	return cmc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CertificateManagerCertificate].
func (cmc *CertificateManagerCertificate) LifecycleManagement() *terra.Lifecycle {
	return cmc.Lifecycle
}

// Attributes returns the attributes for [CertificateManagerCertificate].
func (cmc *CertificateManagerCertificate) Attributes() certificateManagerCertificateAttributes {
	return certificateManagerCertificateAttributes{ref: terra.ReferenceResource(cmc)}
}

// ImportState imports the given attribute values into [CertificateManagerCertificate]'s state.
func (cmc *CertificateManagerCertificate) ImportState(av io.Reader) error {
	cmc.state = &certificateManagerCertificateState{}
	if err := json.NewDecoder(av).Decode(cmc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmc.Type(), cmc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CertificateManagerCertificate] has state.
func (cmc *CertificateManagerCertificate) State() (*certificateManagerCertificateState, bool) {
	return cmc.state, cmc.state != nil
}

// StateMust returns the state for [CertificateManagerCertificate]. Panics if the state is nil.
func (cmc *CertificateManagerCertificate) StateMust() *certificateManagerCertificateState {
	if cmc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmc.Type(), cmc.LocalName()))
	}
	return cmc.state
}

// CertificateManagerCertificateArgs contains the configurations for google_certificate_manager_certificate.
type CertificateManagerCertificateArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Scope: string, optional
	Scope terra.StringValue `hcl:"scope,attr"`
	// Managed: optional
	Managed *certificatemanagercertificate.Managed `hcl:"managed,block"`
	// SelfManaged: optional
	SelfManaged *certificatemanagercertificate.SelfManaged `hcl:"self_managed,block"`
	// Timeouts: optional
	Timeouts *certificatemanagercertificate.Timeouts `hcl:"timeouts,block"`
}
type certificateManagerCertificateAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_certificate_manager_certificate.
func (cmc certificateManagerCertificateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("description"))
}

// Id returns a reference to field id of google_certificate_manager_certificate.
func (cmc certificateManagerCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_certificate_manager_certificate.
func (cmc certificateManagerCertificateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cmc.ref.Append("labels"))
}

// Location returns a reference to field location of google_certificate_manager_certificate.
func (cmc certificateManagerCertificateAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("location"))
}

// Name returns a reference to field name of google_certificate_manager_certificate.
func (cmc certificateManagerCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("name"))
}

// Project returns a reference to field project of google_certificate_manager_certificate.
func (cmc certificateManagerCertificateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("project"))
}

// Scope returns a reference to field scope of google_certificate_manager_certificate.
func (cmc certificateManagerCertificateAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("scope"))
}

func (cmc certificateManagerCertificateAttributes) Managed() terra.ListValue[certificatemanagercertificate.ManagedAttributes] {
	return terra.ReferenceAsList[certificatemanagercertificate.ManagedAttributes](cmc.ref.Append("managed"))
}

func (cmc certificateManagerCertificateAttributes) SelfManaged() terra.ListValue[certificatemanagercertificate.SelfManagedAttributes] {
	return terra.ReferenceAsList[certificatemanagercertificate.SelfManagedAttributes](cmc.ref.Append("self_managed"))
}

func (cmc certificateManagerCertificateAttributes) Timeouts() certificatemanagercertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[certificatemanagercertificate.TimeoutsAttributes](cmc.ref.Append("timeouts"))
}

type certificateManagerCertificateState struct {
	Description string                                           `json:"description"`
	Id          string                                           `json:"id"`
	Labels      map[string]string                                `json:"labels"`
	Location    string                                           `json:"location"`
	Name        string                                           `json:"name"`
	Project     string                                           `json:"project"`
	Scope       string                                           `json:"scope"`
	Managed     []certificatemanagercertificate.ManagedState     `json:"managed"`
	SelfManaged []certificatemanagercertificate.SelfManagedState `json:"self_managed"`
	Timeouts    *certificatemanagercertificate.TimeoutsState     `json:"timeouts"`
}
