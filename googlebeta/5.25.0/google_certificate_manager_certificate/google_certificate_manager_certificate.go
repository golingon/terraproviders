// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_certificate_manager_certificate

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

// Resource represents the Terraform resource google_certificate_manager_certificate.
type Resource struct {
	Name      string
	Args      Args
	state     *googleCertificateManagerCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcmc *Resource) Type() string {
	return "google_certificate_manager_certificate"
}

// LocalName returns the local name for [Resource].
func (gcmc *Resource) LocalName() string {
	return gcmc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcmc *Resource) Configuration() interface{} {
	return gcmc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcmc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcmc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcmc *Resource) Dependencies() terra.Dependencies {
	return gcmc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcmc *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcmc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcmc *Resource) Attributes() googleCertificateManagerCertificateAttributes {
	return googleCertificateManagerCertificateAttributes{ref: terra.ReferenceResource(gcmc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcmc *Resource) ImportState(state io.Reader) error {
	gcmc.state = &googleCertificateManagerCertificateState{}
	if err := json.NewDecoder(state).Decode(gcmc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcmc.Type(), gcmc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcmc *Resource) State() (*googleCertificateManagerCertificateState, bool) {
	return gcmc.state, gcmc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcmc *Resource) StateMust() *googleCertificateManagerCertificateState {
	if gcmc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcmc.Type(), gcmc.LocalName()))
	}
	return gcmc.state
}

// Args contains the configurations for google_certificate_manager_certificate.
type Args struct {
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
	Managed *Managed `hcl:"managed,block"`
	// SelfManaged: optional
	SelfManaged *SelfManaged `hcl:"self_managed,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleCertificateManagerCertificateAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_certificate_manager_certificate.
func (gcmc googleCertificateManagerCertificateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcmc.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_certificate_manager_certificate.
func (gcmc googleCertificateManagerCertificateAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcmc.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_certificate_manager_certificate.
func (gcmc googleCertificateManagerCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcmc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_certificate_manager_certificate.
func (gcmc googleCertificateManagerCertificateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcmc.ref.Append("labels"))
}

// Location returns a reference to field location of google_certificate_manager_certificate.
func (gcmc googleCertificateManagerCertificateAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gcmc.ref.Append("location"))
}

// Name returns a reference to field name of google_certificate_manager_certificate.
func (gcmc googleCertificateManagerCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcmc.ref.Append("name"))
}

// Project returns a reference to field project of google_certificate_manager_certificate.
func (gcmc googleCertificateManagerCertificateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcmc.ref.Append("project"))
}

// Scope returns a reference to field scope of google_certificate_manager_certificate.
func (gcmc googleCertificateManagerCertificateAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(gcmc.ref.Append("scope"))
}

// TerraformLabels returns a reference to field terraform_labels of google_certificate_manager_certificate.
func (gcmc googleCertificateManagerCertificateAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcmc.ref.Append("terraform_labels"))
}

func (gcmc googleCertificateManagerCertificateAttributes) Managed() terra.ListValue[ManagedAttributes] {
	return terra.ReferenceAsList[ManagedAttributes](gcmc.ref.Append("managed"))
}

func (gcmc googleCertificateManagerCertificateAttributes) SelfManaged() terra.ListValue[SelfManagedAttributes] {
	return terra.ReferenceAsList[SelfManagedAttributes](gcmc.ref.Append("self_managed"))
}

func (gcmc googleCertificateManagerCertificateAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcmc.ref.Append("timeouts"))
}

type googleCertificateManagerCertificateState struct {
	Description     string             `json:"description"`
	EffectiveLabels map[string]string  `json:"effective_labels"`
	Id              string             `json:"id"`
	Labels          map[string]string  `json:"labels"`
	Location        string             `json:"location"`
	Name            string             `json:"name"`
	Project         string             `json:"project"`
	Scope           string             `json:"scope"`
	TerraformLabels map[string]string  `json:"terraform_labels"`
	Managed         []ManagedState     `json:"managed"`
	SelfManaged     []SelfManagedState `json:"self_managed"`
	Timeouts        *TimeoutsState     `json:"timeouts"`
}
