// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	privatecacertificatetemplateiambinding "github.com/golingon/terraproviders/google/4.64.0/privatecacertificatetemplateiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivatecaCertificateTemplateIamBinding creates a new instance of [PrivatecaCertificateTemplateIamBinding].
func NewPrivatecaCertificateTemplateIamBinding(name string, args PrivatecaCertificateTemplateIamBindingArgs) *PrivatecaCertificateTemplateIamBinding {
	return &PrivatecaCertificateTemplateIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivatecaCertificateTemplateIamBinding)(nil)

// PrivatecaCertificateTemplateIamBinding represents the Terraform resource google_privateca_certificate_template_iam_binding.
type PrivatecaCertificateTemplateIamBinding struct {
	Name      string
	Args      PrivatecaCertificateTemplateIamBindingArgs
	state     *privatecaCertificateTemplateIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivatecaCertificateTemplateIamBinding].
func (pctib *PrivatecaCertificateTemplateIamBinding) Type() string {
	return "google_privateca_certificate_template_iam_binding"
}

// LocalName returns the local name for [PrivatecaCertificateTemplateIamBinding].
func (pctib *PrivatecaCertificateTemplateIamBinding) LocalName() string {
	return pctib.Name
}

// Configuration returns the configuration (args) for [PrivatecaCertificateTemplateIamBinding].
func (pctib *PrivatecaCertificateTemplateIamBinding) Configuration() interface{} {
	return pctib.Args
}

// DependOn is used for other resources to depend on [PrivatecaCertificateTemplateIamBinding].
func (pctib *PrivatecaCertificateTemplateIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(pctib)
}

// Dependencies returns the list of resources [PrivatecaCertificateTemplateIamBinding] depends_on.
func (pctib *PrivatecaCertificateTemplateIamBinding) Dependencies() terra.Dependencies {
	return pctib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivatecaCertificateTemplateIamBinding].
func (pctib *PrivatecaCertificateTemplateIamBinding) LifecycleManagement() *terra.Lifecycle {
	return pctib.Lifecycle
}

// Attributes returns the attributes for [PrivatecaCertificateTemplateIamBinding].
func (pctib *PrivatecaCertificateTemplateIamBinding) Attributes() privatecaCertificateTemplateIamBindingAttributes {
	return privatecaCertificateTemplateIamBindingAttributes{ref: terra.ReferenceResource(pctib)}
}

// ImportState imports the given attribute values into [PrivatecaCertificateTemplateIamBinding]'s state.
func (pctib *PrivatecaCertificateTemplateIamBinding) ImportState(av io.Reader) error {
	pctib.state = &privatecaCertificateTemplateIamBindingState{}
	if err := json.NewDecoder(av).Decode(pctib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pctib.Type(), pctib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivatecaCertificateTemplateIamBinding] has state.
func (pctib *PrivatecaCertificateTemplateIamBinding) State() (*privatecaCertificateTemplateIamBindingState, bool) {
	return pctib.state, pctib.state != nil
}

// StateMust returns the state for [PrivatecaCertificateTemplateIamBinding]. Panics if the state is nil.
func (pctib *PrivatecaCertificateTemplateIamBinding) StateMust() *privatecaCertificateTemplateIamBindingState {
	if pctib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pctib.Type(), pctib.LocalName()))
	}
	return pctib.state
}

// PrivatecaCertificateTemplateIamBindingArgs contains the configurations for google_privateca_certificate_template_iam_binding.
type PrivatecaCertificateTemplateIamBindingArgs struct {
	// CertificateTemplate: string, required
	CertificateTemplate terra.StringValue `hcl:"certificate_template,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *privatecacertificatetemplateiambinding.Condition `hcl:"condition,block"`
}
type privatecaCertificateTemplateIamBindingAttributes struct {
	ref terra.Reference
}

// CertificateTemplate returns a reference to field certificate_template of google_privateca_certificate_template_iam_binding.
func (pctib privatecaCertificateTemplateIamBindingAttributes) CertificateTemplate() terra.StringValue {
	return terra.ReferenceAsString(pctib.ref.Append("certificate_template"))
}

// Etag returns a reference to field etag of google_privateca_certificate_template_iam_binding.
func (pctib privatecaCertificateTemplateIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(pctib.ref.Append("etag"))
}

// Id returns a reference to field id of google_privateca_certificate_template_iam_binding.
func (pctib privatecaCertificateTemplateIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pctib.ref.Append("id"))
}

// Location returns a reference to field location of google_privateca_certificate_template_iam_binding.
func (pctib privatecaCertificateTemplateIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pctib.ref.Append("location"))
}

// Members returns a reference to field members of google_privateca_certificate_template_iam_binding.
func (pctib privatecaCertificateTemplateIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pctib.ref.Append("members"))
}

// Project returns a reference to field project of google_privateca_certificate_template_iam_binding.
func (pctib privatecaCertificateTemplateIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pctib.ref.Append("project"))
}

// Role returns a reference to field role of google_privateca_certificate_template_iam_binding.
func (pctib privatecaCertificateTemplateIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(pctib.ref.Append("role"))
}

func (pctib privatecaCertificateTemplateIamBindingAttributes) Condition() terra.ListValue[privatecacertificatetemplateiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[privatecacertificatetemplateiambinding.ConditionAttributes](pctib.ref.Append("condition"))
}

type privatecaCertificateTemplateIamBindingState struct {
	CertificateTemplate string                                                  `json:"certificate_template"`
	Etag                string                                                  `json:"etag"`
	Id                  string                                                  `json:"id"`
	Location            string                                                  `json:"location"`
	Members             []string                                                `json:"members"`
	Project             string                                                  `json:"project"`
	Role                string                                                  `json:"role"`
	Condition           []privatecacertificatetemplateiambinding.ConditionState `json:"condition"`
}
