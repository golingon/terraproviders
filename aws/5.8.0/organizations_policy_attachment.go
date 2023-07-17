// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrganizationsPolicyAttachment creates a new instance of [OrganizationsPolicyAttachment].
func NewOrganizationsPolicyAttachment(name string, args OrganizationsPolicyAttachmentArgs) *OrganizationsPolicyAttachment {
	return &OrganizationsPolicyAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrganizationsPolicyAttachment)(nil)

// OrganizationsPolicyAttachment represents the Terraform resource aws_organizations_policy_attachment.
type OrganizationsPolicyAttachment struct {
	Name      string
	Args      OrganizationsPolicyAttachmentArgs
	state     *organizationsPolicyAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrganizationsPolicyAttachment].
func (opa *OrganizationsPolicyAttachment) Type() string {
	return "aws_organizations_policy_attachment"
}

// LocalName returns the local name for [OrganizationsPolicyAttachment].
func (opa *OrganizationsPolicyAttachment) LocalName() string {
	return opa.Name
}

// Configuration returns the configuration (args) for [OrganizationsPolicyAttachment].
func (opa *OrganizationsPolicyAttachment) Configuration() interface{} {
	return opa.Args
}

// DependOn is used for other resources to depend on [OrganizationsPolicyAttachment].
func (opa *OrganizationsPolicyAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(opa)
}

// Dependencies returns the list of resources [OrganizationsPolicyAttachment] depends_on.
func (opa *OrganizationsPolicyAttachment) Dependencies() terra.Dependencies {
	return opa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrganizationsPolicyAttachment].
func (opa *OrganizationsPolicyAttachment) LifecycleManagement() *terra.Lifecycle {
	return opa.Lifecycle
}

// Attributes returns the attributes for [OrganizationsPolicyAttachment].
func (opa *OrganizationsPolicyAttachment) Attributes() organizationsPolicyAttachmentAttributes {
	return organizationsPolicyAttachmentAttributes{ref: terra.ReferenceResource(opa)}
}

// ImportState imports the given attribute values into [OrganizationsPolicyAttachment]'s state.
func (opa *OrganizationsPolicyAttachment) ImportState(av io.Reader) error {
	opa.state = &organizationsPolicyAttachmentState{}
	if err := json.NewDecoder(av).Decode(opa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", opa.Type(), opa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrganizationsPolicyAttachment] has state.
func (opa *OrganizationsPolicyAttachment) State() (*organizationsPolicyAttachmentState, bool) {
	return opa.state, opa.state != nil
}

// StateMust returns the state for [OrganizationsPolicyAttachment]. Panics if the state is nil.
func (opa *OrganizationsPolicyAttachment) StateMust() *organizationsPolicyAttachmentState {
	if opa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", opa.Type(), opa.LocalName()))
	}
	return opa.state
}

// OrganizationsPolicyAttachmentArgs contains the configurations for aws_organizations_policy_attachment.
type OrganizationsPolicyAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// SkipDestroy: bool, optional
	SkipDestroy terra.BoolValue `hcl:"skip_destroy,attr"`
	// TargetId: string, required
	TargetId terra.StringValue `hcl:"target_id,attr" validate:"required"`
}
type organizationsPolicyAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_organizations_policy_attachment.
func (opa organizationsPolicyAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(opa.ref.Append("id"))
}

// PolicyId returns a reference to field policy_id of aws_organizations_policy_attachment.
func (opa organizationsPolicyAttachmentAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(opa.ref.Append("policy_id"))
}

// SkipDestroy returns a reference to field skip_destroy of aws_organizations_policy_attachment.
func (opa organizationsPolicyAttachmentAttributes) SkipDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(opa.ref.Append("skip_destroy"))
}

// TargetId returns a reference to field target_id of aws_organizations_policy_attachment.
func (opa organizationsPolicyAttachmentAttributes) TargetId() terra.StringValue {
	return terra.ReferenceAsString(opa.ref.Append("target_id"))
}

type organizationsPolicyAttachmentState struct {
	Id          string `json:"id"`
	PolicyId    string `json:"policy_id"`
	SkipDestroy bool   `json:"skip_destroy"`
	TargetId    string `json:"target_id"`
}