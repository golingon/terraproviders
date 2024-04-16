// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_privateca_certificate_template

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

// Resource represents the Terraform resource google_privateca_certificate_template.
type Resource struct {
	Name      string
	Args      Args
	state     *googlePrivatecaCertificateTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gpct *Resource) Type() string {
	return "google_privateca_certificate_template"
}

// LocalName returns the local name for [Resource].
func (gpct *Resource) LocalName() string {
	return gpct.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gpct *Resource) Configuration() interface{} {
	return gpct.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gpct *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gpct)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gpct *Resource) Dependencies() terra.Dependencies {
	return gpct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gpct *Resource) LifecycleManagement() *terra.Lifecycle {
	return gpct.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gpct *Resource) Attributes() googlePrivatecaCertificateTemplateAttributes {
	return googlePrivatecaCertificateTemplateAttributes{ref: terra.ReferenceResource(gpct)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gpct *Resource) ImportState(state io.Reader) error {
	gpct.state = &googlePrivatecaCertificateTemplateState{}
	if err := json.NewDecoder(state).Decode(gpct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gpct.Type(), gpct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gpct *Resource) State() (*googlePrivatecaCertificateTemplateState, bool) {
	return gpct.state, gpct.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gpct *Resource) StateMust() *googlePrivatecaCertificateTemplateState {
	if gpct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gpct.Type(), gpct.LocalName()))
	}
	return gpct.state
}

// Args contains the configurations for google_privateca_certificate_template.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// IdentityConstraints: optional
	IdentityConstraints *IdentityConstraints `hcl:"identity_constraints,block"`
	// PassthroughExtensions: optional
	PassthroughExtensions *PassthroughExtensions `hcl:"passthrough_extensions,block"`
	// PredefinedValues: optional
	PredefinedValues *PredefinedValues `hcl:"predefined_values,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googlePrivatecaCertificateTemplateAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_privateca_certificate_template.
func (gpct googlePrivatecaCertificateTemplateAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gpct.ref.Append("create_time"))
}

// Description returns a reference to field description of google_privateca_certificate_template.
func (gpct googlePrivatecaCertificateTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gpct.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_privateca_certificate_template.
func (gpct googlePrivatecaCertificateTemplateAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gpct.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_privateca_certificate_template.
func (gpct googlePrivatecaCertificateTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gpct.ref.Append("id"))
}

// Labels returns a reference to field labels of google_privateca_certificate_template.
func (gpct googlePrivatecaCertificateTemplateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gpct.ref.Append("labels"))
}

// Location returns a reference to field location of google_privateca_certificate_template.
func (gpct googlePrivatecaCertificateTemplateAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gpct.ref.Append("location"))
}

// Name returns a reference to field name of google_privateca_certificate_template.
func (gpct googlePrivatecaCertificateTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gpct.ref.Append("name"))
}

// Project returns a reference to field project of google_privateca_certificate_template.
func (gpct googlePrivatecaCertificateTemplateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gpct.ref.Append("project"))
}

// TerraformLabels returns a reference to field terraform_labels of google_privateca_certificate_template.
func (gpct googlePrivatecaCertificateTemplateAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gpct.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_privateca_certificate_template.
func (gpct googlePrivatecaCertificateTemplateAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gpct.ref.Append("update_time"))
}

func (gpct googlePrivatecaCertificateTemplateAttributes) IdentityConstraints() terra.ListValue[IdentityConstraintsAttributes] {
	return terra.ReferenceAsList[IdentityConstraintsAttributes](gpct.ref.Append("identity_constraints"))
}

func (gpct googlePrivatecaCertificateTemplateAttributes) PassthroughExtensions() terra.ListValue[PassthroughExtensionsAttributes] {
	return terra.ReferenceAsList[PassthroughExtensionsAttributes](gpct.ref.Append("passthrough_extensions"))
}

func (gpct googlePrivatecaCertificateTemplateAttributes) PredefinedValues() terra.ListValue[PredefinedValuesAttributes] {
	return terra.ReferenceAsList[PredefinedValuesAttributes](gpct.ref.Append("predefined_values"))
}

func (gpct googlePrivatecaCertificateTemplateAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gpct.ref.Append("timeouts"))
}

type googlePrivatecaCertificateTemplateState struct {
	CreateTime            string                       `json:"create_time"`
	Description           string                       `json:"description"`
	EffectiveLabels       map[string]string            `json:"effective_labels"`
	Id                    string                       `json:"id"`
	Labels                map[string]string            `json:"labels"`
	Location              string                       `json:"location"`
	Name                  string                       `json:"name"`
	Project               string                       `json:"project"`
	TerraformLabels       map[string]string            `json:"terraform_labels"`
	UpdateTime            string                       `json:"update_time"`
	IdentityConstraints   []IdentityConstraintsState   `json:"identity_constraints"`
	PassthroughExtensions []PassthroughExtensionsState `json:"passthrough_extensions"`
	PredefinedValues      []PredefinedValuesState      `json:"predefined_values"`
	Timeouts              *TimeoutsState               `json:"timeouts"`
}
