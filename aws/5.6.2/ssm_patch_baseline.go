// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssmpatchbaseline "github.com/golingon/terraproviders/aws/5.6.2/ssmpatchbaseline"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsmPatchBaseline creates a new instance of [SsmPatchBaseline].
func NewSsmPatchBaseline(name string, args SsmPatchBaselineArgs) *SsmPatchBaseline {
	return &SsmPatchBaseline{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsmPatchBaseline)(nil)

// SsmPatchBaseline represents the Terraform resource aws_ssm_patch_baseline.
type SsmPatchBaseline struct {
	Name      string
	Args      SsmPatchBaselineArgs
	state     *ssmPatchBaselineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsmPatchBaseline].
func (spb *SsmPatchBaseline) Type() string {
	return "aws_ssm_patch_baseline"
}

// LocalName returns the local name for [SsmPatchBaseline].
func (spb *SsmPatchBaseline) LocalName() string {
	return spb.Name
}

// Configuration returns the configuration (args) for [SsmPatchBaseline].
func (spb *SsmPatchBaseline) Configuration() interface{} {
	return spb.Args
}

// DependOn is used for other resources to depend on [SsmPatchBaseline].
func (spb *SsmPatchBaseline) DependOn() terra.Reference {
	return terra.ReferenceResource(spb)
}

// Dependencies returns the list of resources [SsmPatchBaseline] depends_on.
func (spb *SsmPatchBaseline) Dependencies() terra.Dependencies {
	return spb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsmPatchBaseline].
func (spb *SsmPatchBaseline) LifecycleManagement() *terra.Lifecycle {
	return spb.Lifecycle
}

// Attributes returns the attributes for [SsmPatchBaseline].
func (spb *SsmPatchBaseline) Attributes() ssmPatchBaselineAttributes {
	return ssmPatchBaselineAttributes{ref: terra.ReferenceResource(spb)}
}

// ImportState imports the given attribute values into [SsmPatchBaseline]'s state.
func (spb *SsmPatchBaseline) ImportState(av io.Reader) error {
	spb.state = &ssmPatchBaselineState{}
	if err := json.NewDecoder(av).Decode(spb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spb.Type(), spb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsmPatchBaseline] has state.
func (spb *SsmPatchBaseline) State() (*ssmPatchBaselineState, bool) {
	return spb.state, spb.state != nil
}

// StateMust returns the state for [SsmPatchBaseline]. Panics if the state is nil.
func (spb *SsmPatchBaseline) StateMust() *ssmPatchBaselineState {
	if spb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spb.Type(), spb.LocalName()))
	}
	return spb.state
}

// SsmPatchBaselineArgs contains the configurations for aws_ssm_patch_baseline.
type SsmPatchBaselineArgs struct {
	// ApprovedPatches: set of string, optional
	ApprovedPatches terra.SetValue[terra.StringValue] `hcl:"approved_patches,attr"`
	// ApprovedPatchesComplianceLevel: string, optional
	ApprovedPatchesComplianceLevel terra.StringValue `hcl:"approved_patches_compliance_level,attr"`
	// ApprovedPatchesEnableNonSecurity: bool, optional
	ApprovedPatchesEnableNonSecurity terra.BoolValue `hcl:"approved_patches_enable_non_security,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OperatingSystem: string, optional
	OperatingSystem terra.StringValue `hcl:"operating_system,attr"`
	// RejectedPatches: set of string, optional
	RejectedPatches terra.SetValue[terra.StringValue] `hcl:"rejected_patches,attr"`
	// RejectedPatchesAction: string, optional
	RejectedPatchesAction terra.StringValue `hcl:"rejected_patches_action,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ApprovalRule: min=0
	ApprovalRule []ssmpatchbaseline.ApprovalRule `hcl:"approval_rule,block" validate:"min=0"`
	// GlobalFilter: min=0,max=4
	GlobalFilter []ssmpatchbaseline.GlobalFilter `hcl:"global_filter,block" validate:"min=0,max=4"`
	// Source: min=0,max=20
	Source []ssmpatchbaseline.Source `hcl:"source,block" validate:"min=0,max=20"`
}
type ssmPatchBaselineAttributes struct {
	ref terra.Reference
}

// ApprovedPatches returns a reference to field approved_patches of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) ApprovedPatches() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](spb.ref.Append("approved_patches"))
}

// ApprovedPatchesComplianceLevel returns a reference to field approved_patches_compliance_level of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) ApprovedPatchesComplianceLevel() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("approved_patches_compliance_level"))
}

// ApprovedPatchesEnableNonSecurity returns a reference to field approved_patches_enable_non_security of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) ApprovedPatchesEnableNonSecurity() terra.BoolValue {
	return terra.ReferenceAsBool(spb.ref.Append("approved_patches_enable_non_security"))
}

// Arn returns a reference to field arn of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("arn"))
}

// Description returns a reference to field description of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("description"))
}

// Id returns a reference to field id of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("id"))
}

// Name returns a reference to field name of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("name"))
}

// OperatingSystem returns a reference to field operating_system of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) OperatingSystem() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("operating_system"))
}

// RejectedPatches returns a reference to field rejected_patches of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) RejectedPatches() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](spb.ref.Append("rejected_patches"))
}

// RejectedPatchesAction returns a reference to field rejected_patches_action of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) RejectedPatchesAction() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("rejected_patches_action"))
}

// Tags returns a reference to field tags of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](spb.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ssm_patch_baseline.
func (spb ssmPatchBaselineAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](spb.ref.Append("tags_all"))
}

func (spb ssmPatchBaselineAttributes) ApprovalRule() terra.ListValue[ssmpatchbaseline.ApprovalRuleAttributes] {
	return terra.ReferenceAsList[ssmpatchbaseline.ApprovalRuleAttributes](spb.ref.Append("approval_rule"))
}

func (spb ssmPatchBaselineAttributes) GlobalFilter() terra.ListValue[ssmpatchbaseline.GlobalFilterAttributes] {
	return terra.ReferenceAsList[ssmpatchbaseline.GlobalFilterAttributes](spb.ref.Append("global_filter"))
}

func (spb ssmPatchBaselineAttributes) Source() terra.ListValue[ssmpatchbaseline.SourceAttributes] {
	return terra.ReferenceAsList[ssmpatchbaseline.SourceAttributes](spb.ref.Append("source"))
}

type ssmPatchBaselineState struct {
	ApprovedPatches                  []string                             `json:"approved_patches"`
	ApprovedPatchesComplianceLevel   string                               `json:"approved_patches_compliance_level"`
	ApprovedPatchesEnableNonSecurity bool                                 `json:"approved_patches_enable_non_security"`
	Arn                              string                               `json:"arn"`
	Description                      string                               `json:"description"`
	Id                               string                               `json:"id"`
	Name                             string                               `json:"name"`
	OperatingSystem                  string                               `json:"operating_system"`
	RejectedPatches                  []string                             `json:"rejected_patches"`
	RejectedPatchesAction            string                               `json:"rejected_patches_action"`
	Tags                             map[string]string                    `json:"tags"`
	TagsAll                          map[string]string                    `json:"tags_all"`
	ApprovalRule                     []ssmpatchbaseline.ApprovalRuleState `json:"approval_rule"`
	GlobalFilter                     []ssmpatchbaseline.GlobalFilterState `json:"global_filter"`
	Source                           []ssmpatchbaseline.SourceState       `json:"source"`
}