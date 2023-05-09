// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	privatecacertificatetemplateiammember "github.com/golingon/terraproviders/google/4.64.0/privatecacertificatetemplateiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivatecaCertificateTemplateIamMember creates a new instance of [PrivatecaCertificateTemplateIamMember].
func NewPrivatecaCertificateTemplateIamMember(name string, args PrivatecaCertificateTemplateIamMemberArgs) *PrivatecaCertificateTemplateIamMember {
	return &PrivatecaCertificateTemplateIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivatecaCertificateTemplateIamMember)(nil)

// PrivatecaCertificateTemplateIamMember represents the Terraform resource google_privateca_certificate_template_iam_member.
type PrivatecaCertificateTemplateIamMember struct {
	Name      string
	Args      PrivatecaCertificateTemplateIamMemberArgs
	state     *privatecaCertificateTemplateIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivatecaCertificateTemplateIamMember].
func (pctim *PrivatecaCertificateTemplateIamMember) Type() string {
	return "google_privateca_certificate_template_iam_member"
}

// LocalName returns the local name for [PrivatecaCertificateTemplateIamMember].
func (pctim *PrivatecaCertificateTemplateIamMember) LocalName() string {
	return pctim.Name
}

// Configuration returns the configuration (args) for [PrivatecaCertificateTemplateIamMember].
func (pctim *PrivatecaCertificateTemplateIamMember) Configuration() interface{} {
	return pctim.Args
}

// DependOn is used for other resources to depend on [PrivatecaCertificateTemplateIamMember].
func (pctim *PrivatecaCertificateTemplateIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(pctim)
}

// Dependencies returns the list of resources [PrivatecaCertificateTemplateIamMember] depends_on.
func (pctim *PrivatecaCertificateTemplateIamMember) Dependencies() terra.Dependencies {
	return pctim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivatecaCertificateTemplateIamMember].
func (pctim *PrivatecaCertificateTemplateIamMember) LifecycleManagement() *terra.Lifecycle {
	return pctim.Lifecycle
}

// Attributes returns the attributes for [PrivatecaCertificateTemplateIamMember].
func (pctim *PrivatecaCertificateTemplateIamMember) Attributes() privatecaCertificateTemplateIamMemberAttributes {
	return privatecaCertificateTemplateIamMemberAttributes{ref: terra.ReferenceResource(pctim)}
}

// ImportState imports the given attribute values into [PrivatecaCertificateTemplateIamMember]'s state.
func (pctim *PrivatecaCertificateTemplateIamMember) ImportState(av io.Reader) error {
	pctim.state = &privatecaCertificateTemplateIamMemberState{}
	if err := json.NewDecoder(av).Decode(pctim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pctim.Type(), pctim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivatecaCertificateTemplateIamMember] has state.
func (pctim *PrivatecaCertificateTemplateIamMember) State() (*privatecaCertificateTemplateIamMemberState, bool) {
	return pctim.state, pctim.state != nil
}

// StateMust returns the state for [PrivatecaCertificateTemplateIamMember]. Panics if the state is nil.
func (pctim *PrivatecaCertificateTemplateIamMember) StateMust() *privatecaCertificateTemplateIamMemberState {
	if pctim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pctim.Type(), pctim.LocalName()))
	}
	return pctim.state
}

// PrivatecaCertificateTemplateIamMemberArgs contains the configurations for google_privateca_certificate_template_iam_member.
type PrivatecaCertificateTemplateIamMemberArgs struct {
	// CertificateTemplate: string, required
	CertificateTemplate terra.StringValue `hcl:"certificate_template,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *privatecacertificatetemplateiammember.Condition `hcl:"condition,block"`
}
type privatecaCertificateTemplateIamMemberAttributes struct {
	ref terra.Reference
}

// CertificateTemplate returns a reference to field certificate_template of google_privateca_certificate_template_iam_member.
func (pctim privatecaCertificateTemplateIamMemberAttributes) CertificateTemplate() terra.StringValue {
	return terra.ReferenceAsString(pctim.ref.Append("certificate_template"))
}

// Etag returns a reference to field etag of google_privateca_certificate_template_iam_member.
func (pctim privatecaCertificateTemplateIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(pctim.ref.Append("etag"))
}

// Id returns a reference to field id of google_privateca_certificate_template_iam_member.
func (pctim privatecaCertificateTemplateIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pctim.ref.Append("id"))
}

// Location returns a reference to field location of google_privateca_certificate_template_iam_member.
func (pctim privatecaCertificateTemplateIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pctim.ref.Append("location"))
}

// Member returns a reference to field member of google_privateca_certificate_template_iam_member.
func (pctim privatecaCertificateTemplateIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(pctim.ref.Append("member"))
}

// Project returns a reference to field project of google_privateca_certificate_template_iam_member.
func (pctim privatecaCertificateTemplateIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pctim.ref.Append("project"))
}

// Role returns a reference to field role of google_privateca_certificate_template_iam_member.
func (pctim privatecaCertificateTemplateIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(pctim.ref.Append("role"))
}

func (pctim privatecaCertificateTemplateIamMemberAttributes) Condition() terra.ListValue[privatecacertificatetemplateiammember.ConditionAttributes] {
	return terra.ReferenceAsList[privatecacertificatetemplateiammember.ConditionAttributes](pctim.ref.Append("condition"))
}

type privatecaCertificateTemplateIamMemberState struct {
	CertificateTemplate string                                                 `json:"certificate_template"`
	Etag                string                                                 `json:"etag"`
	Id                  string                                                 `json:"id"`
	Location            string                                                 `json:"location"`
	Member              string                                                 `json:"member"`
	Project             string                                                 `json:"project"`
	Role                string                                                 `json:"role"`
	Condition           []privatecacertificatetemplateiammember.ConditionState `json:"condition"`
}
