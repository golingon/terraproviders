// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	privatecacertificatetemplate "github.com/golingon/terraproviders/googlebeta/4.66.0/privatecacertificatetemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivatecaCertificateTemplate creates a new instance of [PrivatecaCertificateTemplate].
func NewPrivatecaCertificateTemplate(name string, args PrivatecaCertificateTemplateArgs) *PrivatecaCertificateTemplate {
	return &PrivatecaCertificateTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivatecaCertificateTemplate)(nil)

// PrivatecaCertificateTemplate represents the Terraform resource google_privateca_certificate_template.
type PrivatecaCertificateTemplate struct {
	Name      string
	Args      PrivatecaCertificateTemplateArgs
	state     *privatecaCertificateTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivatecaCertificateTemplate].
func (pct *PrivatecaCertificateTemplate) Type() string {
	return "google_privateca_certificate_template"
}

// LocalName returns the local name for [PrivatecaCertificateTemplate].
func (pct *PrivatecaCertificateTemplate) LocalName() string {
	return pct.Name
}

// Configuration returns the configuration (args) for [PrivatecaCertificateTemplate].
func (pct *PrivatecaCertificateTemplate) Configuration() interface{} {
	return pct.Args
}

// DependOn is used for other resources to depend on [PrivatecaCertificateTemplate].
func (pct *PrivatecaCertificateTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(pct)
}

// Dependencies returns the list of resources [PrivatecaCertificateTemplate] depends_on.
func (pct *PrivatecaCertificateTemplate) Dependencies() terra.Dependencies {
	return pct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivatecaCertificateTemplate].
func (pct *PrivatecaCertificateTemplate) LifecycleManagement() *terra.Lifecycle {
	return pct.Lifecycle
}

// Attributes returns the attributes for [PrivatecaCertificateTemplate].
func (pct *PrivatecaCertificateTemplate) Attributes() privatecaCertificateTemplateAttributes {
	return privatecaCertificateTemplateAttributes{ref: terra.ReferenceResource(pct)}
}

// ImportState imports the given attribute values into [PrivatecaCertificateTemplate]'s state.
func (pct *PrivatecaCertificateTemplate) ImportState(av io.Reader) error {
	pct.state = &privatecaCertificateTemplateState{}
	if err := json.NewDecoder(av).Decode(pct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pct.Type(), pct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivatecaCertificateTemplate] has state.
func (pct *PrivatecaCertificateTemplate) State() (*privatecaCertificateTemplateState, bool) {
	return pct.state, pct.state != nil
}

// StateMust returns the state for [PrivatecaCertificateTemplate]. Panics if the state is nil.
func (pct *PrivatecaCertificateTemplate) StateMust() *privatecaCertificateTemplateState {
	if pct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pct.Type(), pct.LocalName()))
	}
	return pct.state
}

// PrivatecaCertificateTemplateArgs contains the configurations for google_privateca_certificate_template.
type PrivatecaCertificateTemplateArgs struct {
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
	IdentityConstraints *privatecacertificatetemplate.IdentityConstraints `hcl:"identity_constraints,block"`
	// PassthroughExtensions: optional
	PassthroughExtensions *privatecacertificatetemplate.PassthroughExtensions `hcl:"passthrough_extensions,block"`
	// PredefinedValues: optional
	PredefinedValues *privatecacertificatetemplate.PredefinedValues `hcl:"predefined_values,block"`
	// Timeouts: optional
	Timeouts *privatecacertificatetemplate.Timeouts `hcl:"timeouts,block"`
}
type privatecaCertificateTemplateAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_privateca_certificate_template.
func (pct privatecaCertificateTemplateAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(pct.ref.Append("create_time"))
}

// Description returns a reference to field description of google_privateca_certificate_template.
func (pct privatecaCertificateTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(pct.ref.Append("description"))
}

// Id returns a reference to field id of google_privateca_certificate_template.
func (pct privatecaCertificateTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pct.ref.Append("id"))
}

// Labels returns a reference to field labels of google_privateca_certificate_template.
func (pct privatecaCertificateTemplateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pct.ref.Append("labels"))
}

// Location returns a reference to field location of google_privateca_certificate_template.
func (pct privatecaCertificateTemplateAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pct.ref.Append("location"))
}

// Name returns a reference to field name of google_privateca_certificate_template.
func (pct privatecaCertificateTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pct.ref.Append("name"))
}

// Project returns a reference to field project of google_privateca_certificate_template.
func (pct privatecaCertificateTemplateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pct.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_privateca_certificate_template.
func (pct privatecaCertificateTemplateAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(pct.ref.Append("update_time"))
}

func (pct privatecaCertificateTemplateAttributes) IdentityConstraints() terra.ListValue[privatecacertificatetemplate.IdentityConstraintsAttributes] {
	return terra.ReferenceAsList[privatecacertificatetemplate.IdentityConstraintsAttributes](pct.ref.Append("identity_constraints"))
}

func (pct privatecaCertificateTemplateAttributes) PassthroughExtensions() terra.ListValue[privatecacertificatetemplate.PassthroughExtensionsAttributes] {
	return terra.ReferenceAsList[privatecacertificatetemplate.PassthroughExtensionsAttributes](pct.ref.Append("passthrough_extensions"))
}

func (pct privatecaCertificateTemplateAttributes) PredefinedValues() terra.ListValue[privatecacertificatetemplate.PredefinedValuesAttributes] {
	return terra.ReferenceAsList[privatecacertificatetemplate.PredefinedValuesAttributes](pct.ref.Append("predefined_values"))
}

func (pct privatecaCertificateTemplateAttributes) Timeouts() privatecacertificatetemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatecacertificatetemplate.TimeoutsAttributes](pct.ref.Append("timeouts"))
}

type privatecaCertificateTemplateState struct {
	CreateTime            string                                                    `json:"create_time"`
	Description           string                                                    `json:"description"`
	Id                    string                                                    `json:"id"`
	Labels                map[string]string                                         `json:"labels"`
	Location              string                                                    `json:"location"`
	Name                  string                                                    `json:"name"`
	Project               string                                                    `json:"project"`
	UpdateTime            string                                                    `json:"update_time"`
	IdentityConstraints   []privatecacertificatetemplate.IdentityConstraintsState   `json:"identity_constraints"`
	PassthroughExtensions []privatecacertificatetemplate.PassthroughExtensionsState `json:"passthrough_extensions"`
	PredefinedValues      []privatecacertificatetemplate.PredefinedValuesState      `json:"predefined_values"`
	Timeouts              *privatecacertificatetemplate.TimeoutsState               `json:"timeouts"`
}
