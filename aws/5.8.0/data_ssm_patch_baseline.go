// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datassmpatchbaseline "github.com/golingon/terraproviders/aws/5.8.0/datassmpatchbaseline"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSsmPatchBaseline creates a new instance of [DataSsmPatchBaseline].
func NewDataSsmPatchBaseline(name string, args DataSsmPatchBaselineArgs) *DataSsmPatchBaseline {
	return &DataSsmPatchBaseline{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSsmPatchBaseline)(nil)

// DataSsmPatchBaseline represents the Terraform data resource aws_ssm_patch_baseline.
type DataSsmPatchBaseline struct {
	Name string
	Args DataSsmPatchBaselineArgs
}

// DataSource returns the Terraform object type for [DataSsmPatchBaseline].
func (spb *DataSsmPatchBaseline) DataSource() string {
	return "aws_ssm_patch_baseline"
}

// LocalName returns the local name for [DataSsmPatchBaseline].
func (spb *DataSsmPatchBaseline) LocalName() string {
	return spb.Name
}

// Configuration returns the configuration (args) for [DataSsmPatchBaseline].
func (spb *DataSsmPatchBaseline) Configuration() interface{} {
	return spb.Args
}

// Attributes returns the attributes for [DataSsmPatchBaseline].
func (spb *DataSsmPatchBaseline) Attributes() dataSsmPatchBaselineAttributes {
	return dataSsmPatchBaselineAttributes{ref: terra.ReferenceDataResource(spb)}
}

// DataSsmPatchBaselineArgs contains the configurations for aws_ssm_patch_baseline.
type DataSsmPatchBaselineArgs struct {
	// DefaultBaseline: bool, optional
	DefaultBaseline terra.BoolValue `hcl:"default_baseline,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// OperatingSystem: string, optional
	OperatingSystem terra.StringValue `hcl:"operating_system,attr"`
	// Owner: string, required
	Owner terra.StringValue `hcl:"owner,attr" validate:"required"`
	// ApprovalRule: min=0
	ApprovalRule []datassmpatchbaseline.ApprovalRule `hcl:"approval_rule,block" validate:"min=0"`
	// GlobalFilter: min=0
	GlobalFilter []datassmpatchbaseline.GlobalFilter `hcl:"global_filter,block" validate:"min=0"`
	// Source: min=0
	Source []datassmpatchbaseline.Source `hcl:"source,block" validate:"min=0"`
}
type dataSsmPatchBaselineAttributes struct {
	ref terra.Reference
}

// ApprovedPatches returns a reference to field approved_patches of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) ApprovedPatches() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](spb.ref.Append("approved_patches"))
}

// ApprovedPatchesComplianceLevel returns a reference to field approved_patches_compliance_level of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) ApprovedPatchesComplianceLevel() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("approved_patches_compliance_level"))
}

// ApprovedPatchesEnableNonSecurity returns a reference to field approved_patches_enable_non_security of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) ApprovedPatchesEnableNonSecurity() terra.BoolValue {
	return terra.ReferenceAsBool(spb.ref.Append("approved_patches_enable_non_security"))
}

// DefaultBaseline returns a reference to field default_baseline of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) DefaultBaseline() terra.BoolValue {
	return terra.ReferenceAsBool(spb.ref.Append("default_baseline"))
}

// Description returns a reference to field description of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("description"))
}

// Id returns a reference to field id of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("id"))
}

// Name returns a reference to field name of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("name_prefix"))
}

// OperatingSystem returns a reference to field operating_system of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) OperatingSystem() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("operating_system"))
}

// Owner returns a reference to field owner of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("owner"))
}

// RejectedPatches returns a reference to field rejected_patches of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) RejectedPatches() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](spb.ref.Append("rejected_patches"))
}

// RejectedPatchesAction returns a reference to field rejected_patches_action of aws_ssm_patch_baseline.
func (spb dataSsmPatchBaselineAttributes) RejectedPatchesAction() terra.StringValue {
	return terra.ReferenceAsString(spb.ref.Append("rejected_patches_action"))
}

func (spb dataSsmPatchBaselineAttributes) ApprovalRule() terra.ListValue[datassmpatchbaseline.ApprovalRuleAttributes] {
	return terra.ReferenceAsList[datassmpatchbaseline.ApprovalRuleAttributes](spb.ref.Append("approval_rule"))
}

func (spb dataSsmPatchBaselineAttributes) GlobalFilter() terra.ListValue[datassmpatchbaseline.GlobalFilterAttributes] {
	return terra.ReferenceAsList[datassmpatchbaseline.GlobalFilterAttributes](spb.ref.Append("global_filter"))
}

func (spb dataSsmPatchBaselineAttributes) Source() terra.ListValue[datassmpatchbaseline.SourceAttributes] {
	return terra.ReferenceAsList[datassmpatchbaseline.SourceAttributes](spb.ref.Append("source"))
}
