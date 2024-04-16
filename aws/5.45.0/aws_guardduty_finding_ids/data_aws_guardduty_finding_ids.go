// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_guardduty_finding_ids

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_guardduty_finding_ids.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (agfi *DataSource) DataSource() string {
	return "aws_guardduty_finding_ids"
}

// LocalName returns the local name for [DataSource].
func (agfi *DataSource) LocalName() string {
	return agfi.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (agfi *DataSource) Configuration() interface{} {
	return agfi.Args
}

// Attributes returns the attributes for [DataSource].
func (agfi *DataSource) Attributes() dataAwsGuarddutyFindingIdsAttributes {
	return dataAwsGuarddutyFindingIdsAttributes{ref: terra.ReferenceDataSource(agfi)}
}

// DataArgs contains the configurations for aws_guardduty_finding_ids.
type DataArgs struct {
	// DetectorId: string, required
	DetectorId terra.StringValue `hcl:"detector_id,attr" validate:"required"`
}

type dataAwsGuarddutyFindingIdsAttributes struct {
	ref terra.Reference
}

// DetectorId returns a reference to field detector_id of aws_guardduty_finding_ids.
func (agfi dataAwsGuarddutyFindingIdsAttributes) DetectorId() terra.StringValue {
	return terra.ReferenceAsString(agfi.ref.Append("detector_id"))
}

// FindingIds returns a reference to field finding_ids of aws_guardduty_finding_ids.
func (agfi dataAwsGuarddutyFindingIdsAttributes) FindingIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](agfi.ref.Append("finding_ids"))
}

// HasFindings returns a reference to field has_findings of aws_guardduty_finding_ids.
func (agfi dataAwsGuarddutyFindingIdsAttributes) HasFindings() terra.BoolValue {
	return terra.ReferenceAsBool(agfi.ref.Append("has_findings"))
}

// Id returns a reference to field id of aws_guardduty_finding_ids.
func (agfi dataAwsGuarddutyFindingIdsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agfi.ref.Append("id"))
}
