// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_auditmanager_control

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_auditmanager_control.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aac *DataSource) DataSource() string {
	return "aws_auditmanager_control"
}

// LocalName returns the local name for [DataSource].
func (aac *DataSource) LocalName() string {
	return aac.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aac *DataSource) Configuration() interface{} {
	return aac.Args
}

// Attributes returns the attributes for [DataSource].
func (aac *DataSource) Attributes() dataAwsAuditmanagerControlAttributes {
	return dataAwsAuditmanagerControlAttributes{ref: terra.ReferenceDataSource(aac)}
}

// DataArgs contains the configurations for aws_auditmanager_control.
type DataArgs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// ControlMappingSources: min=0
	ControlMappingSources []DataControlMappingSources `hcl:"control_mapping_sources,block" validate:"min=0"`
}

type dataAwsAuditmanagerControlAttributes struct {
	ref terra.Reference
}

// ActionPlanInstructions returns a reference to field action_plan_instructions of aws_auditmanager_control.
func (aac dataAwsAuditmanagerControlAttributes) ActionPlanInstructions() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("action_plan_instructions"))
}

// ActionPlanTitle returns a reference to field action_plan_title of aws_auditmanager_control.
func (aac dataAwsAuditmanagerControlAttributes) ActionPlanTitle() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("action_plan_title"))
}

// Arn returns a reference to field arn of aws_auditmanager_control.
func (aac dataAwsAuditmanagerControlAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("arn"))
}

// Description returns a reference to field description of aws_auditmanager_control.
func (aac dataAwsAuditmanagerControlAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("description"))
}

// Id returns a reference to field id of aws_auditmanager_control.
func (aac dataAwsAuditmanagerControlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("id"))
}

// Name returns a reference to field name of aws_auditmanager_control.
func (aac dataAwsAuditmanagerControlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_auditmanager_control.
func (aac dataAwsAuditmanagerControlAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aac.ref.Append("tags"))
}

// TestingInformation returns a reference to field testing_information of aws_auditmanager_control.
func (aac dataAwsAuditmanagerControlAttributes) TestingInformation() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("testing_information"))
}

// Type returns a reference to field type of aws_auditmanager_control.
func (aac dataAwsAuditmanagerControlAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("type"))
}

func (aac dataAwsAuditmanagerControlAttributes) ControlMappingSources() terra.SetValue[DataControlMappingSourcesAttributes] {
	return terra.ReferenceAsSet[DataControlMappingSourcesAttributes](aac.ref.Append("control_mapping_sources"))
}
