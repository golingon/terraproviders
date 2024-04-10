// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewSyntheticsGroupAssociation creates a new instance of [SyntheticsGroupAssociation].
func NewSyntheticsGroupAssociation(name string, args SyntheticsGroupAssociationArgs) *SyntheticsGroupAssociation {
	return &SyntheticsGroupAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SyntheticsGroupAssociation)(nil)

// SyntheticsGroupAssociation represents the Terraform resource aws_synthetics_group_association.
type SyntheticsGroupAssociation struct {
	Name      string
	Args      SyntheticsGroupAssociationArgs
	state     *syntheticsGroupAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SyntheticsGroupAssociation].
func (sga *SyntheticsGroupAssociation) Type() string {
	return "aws_synthetics_group_association"
}

// LocalName returns the local name for [SyntheticsGroupAssociation].
func (sga *SyntheticsGroupAssociation) LocalName() string {
	return sga.Name
}

// Configuration returns the configuration (args) for [SyntheticsGroupAssociation].
func (sga *SyntheticsGroupAssociation) Configuration() interface{} {
	return sga.Args
}

// DependOn is used for other resources to depend on [SyntheticsGroupAssociation].
func (sga *SyntheticsGroupAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(sga)
}

// Dependencies returns the list of resources [SyntheticsGroupAssociation] depends_on.
func (sga *SyntheticsGroupAssociation) Dependencies() terra.Dependencies {
	return sga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SyntheticsGroupAssociation].
func (sga *SyntheticsGroupAssociation) LifecycleManagement() *terra.Lifecycle {
	return sga.Lifecycle
}

// Attributes returns the attributes for [SyntheticsGroupAssociation].
func (sga *SyntheticsGroupAssociation) Attributes() syntheticsGroupAssociationAttributes {
	return syntheticsGroupAssociationAttributes{ref: terra.ReferenceResource(sga)}
}

// ImportState imports the given attribute values into [SyntheticsGroupAssociation]'s state.
func (sga *SyntheticsGroupAssociation) ImportState(av io.Reader) error {
	sga.state = &syntheticsGroupAssociationState{}
	if err := json.NewDecoder(av).Decode(sga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sga.Type(), sga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SyntheticsGroupAssociation] has state.
func (sga *SyntheticsGroupAssociation) State() (*syntheticsGroupAssociationState, bool) {
	return sga.state, sga.state != nil
}

// StateMust returns the state for [SyntheticsGroupAssociation]. Panics if the state is nil.
func (sga *SyntheticsGroupAssociation) StateMust() *syntheticsGroupAssociationState {
	if sga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sga.Type(), sga.LocalName()))
	}
	return sga.state
}

// SyntheticsGroupAssociationArgs contains the configurations for aws_synthetics_group_association.
type SyntheticsGroupAssociationArgs struct {
	// CanaryArn: string, required
	CanaryArn terra.StringValue `hcl:"canary_arn,attr" validate:"required"`
	// GroupName: string, required
	GroupName terra.StringValue `hcl:"group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type syntheticsGroupAssociationAttributes struct {
	ref terra.Reference
}

// CanaryArn returns a reference to field canary_arn of aws_synthetics_group_association.
func (sga syntheticsGroupAssociationAttributes) CanaryArn() terra.StringValue {
	return terra.ReferenceAsString(sga.ref.Append("canary_arn"))
}

// GroupArn returns a reference to field group_arn of aws_synthetics_group_association.
func (sga syntheticsGroupAssociationAttributes) GroupArn() terra.StringValue {
	return terra.ReferenceAsString(sga.ref.Append("group_arn"))
}

// GroupId returns a reference to field group_id of aws_synthetics_group_association.
func (sga syntheticsGroupAssociationAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(sga.ref.Append("group_id"))
}

// GroupName returns a reference to field group_name of aws_synthetics_group_association.
func (sga syntheticsGroupAssociationAttributes) GroupName() terra.StringValue {
	return terra.ReferenceAsString(sga.ref.Append("group_name"))
}

// Id returns a reference to field id of aws_synthetics_group_association.
func (sga syntheticsGroupAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sga.ref.Append("id"))
}

type syntheticsGroupAssociationState struct {
	CanaryArn string `json:"canary_arn"`
	GroupArn  string `json:"group_arn"`
	GroupId   string `json:"group_id"`
	GroupName string `json:"group_name"`
	Id        string `json:"id"`
}
