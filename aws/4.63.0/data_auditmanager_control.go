// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataauditmanagercontrol "github.com/golingon/terraproviders/aws/4.63.0/dataauditmanagercontrol"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAuditmanagerControl creates a new instance of [DataAuditmanagerControl].
func NewDataAuditmanagerControl(name string, args DataAuditmanagerControlArgs) *DataAuditmanagerControl {
	return &DataAuditmanagerControl{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAuditmanagerControl)(nil)

// DataAuditmanagerControl represents the Terraform data resource aws_auditmanager_control.
type DataAuditmanagerControl struct {
	Name string
	Args DataAuditmanagerControlArgs
}

// DataSource returns the Terraform object type for [DataAuditmanagerControl].
func (ac *DataAuditmanagerControl) DataSource() string {
	return "aws_auditmanager_control"
}

// LocalName returns the local name for [DataAuditmanagerControl].
func (ac *DataAuditmanagerControl) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [DataAuditmanagerControl].
func (ac *DataAuditmanagerControl) Configuration() interface{} {
	return ac.Args
}

// Attributes returns the attributes for [DataAuditmanagerControl].
func (ac *DataAuditmanagerControl) Attributes() dataAuditmanagerControlAttributes {
	return dataAuditmanagerControlAttributes{ref: terra.ReferenceDataResource(ac)}
}

// DataAuditmanagerControlArgs contains the configurations for aws_auditmanager_control.
type DataAuditmanagerControlArgs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// ControlMappingSources: min=0
	ControlMappingSources []dataauditmanagercontrol.ControlMappingSources `hcl:"control_mapping_sources,block" validate:"min=0"`
}
type dataAuditmanagerControlAttributes struct {
	ref terra.Reference
}

// ActionPlanInstructions returns a reference to field action_plan_instructions of aws_auditmanager_control.
func (ac dataAuditmanagerControlAttributes) ActionPlanInstructions() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("action_plan_instructions"))
}

// ActionPlanTitle returns a reference to field action_plan_title of aws_auditmanager_control.
func (ac dataAuditmanagerControlAttributes) ActionPlanTitle() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("action_plan_title"))
}

// Arn returns a reference to field arn of aws_auditmanager_control.
func (ac dataAuditmanagerControlAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("arn"))
}

// Description returns a reference to field description of aws_auditmanager_control.
func (ac dataAuditmanagerControlAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("description"))
}

// Id returns a reference to field id of aws_auditmanager_control.
func (ac dataAuditmanagerControlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// Name returns a reference to field name of aws_auditmanager_control.
func (ac dataAuditmanagerControlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_auditmanager_control.
func (ac dataAuditmanagerControlAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("tags"))
}

// TestingInformation returns a reference to field testing_information of aws_auditmanager_control.
func (ac dataAuditmanagerControlAttributes) TestingInformation() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("testing_information"))
}

// Type returns a reference to field type of aws_auditmanager_control.
func (ac dataAuditmanagerControlAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("type"))
}

func (ac dataAuditmanagerControlAttributes) ControlMappingSources() terra.SetValue[dataauditmanagercontrol.ControlMappingSourcesAttributes] {
	return terra.ReferenceAsSet[dataauditmanagercontrol.ControlMappingSourcesAttributes](ac.ref.Append("control_mapping_sources"))
}
