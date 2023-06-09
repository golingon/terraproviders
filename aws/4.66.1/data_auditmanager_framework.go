// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataauditmanagerframework "github.com/golingon/terraproviders/aws/4.66.1/dataauditmanagerframework"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAuditmanagerFramework creates a new instance of [DataAuditmanagerFramework].
func NewDataAuditmanagerFramework(name string, args DataAuditmanagerFrameworkArgs) *DataAuditmanagerFramework {
	return &DataAuditmanagerFramework{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAuditmanagerFramework)(nil)

// DataAuditmanagerFramework represents the Terraform data resource aws_auditmanager_framework.
type DataAuditmanagerFramework struct {
	Name string
	Args DataAuditmanagerFrameworkArgs
}

// DataSource returns the Terraform object type for [DataAuditmanagerFramework].
func (af *DataAuditmanagerFramework) DataSource() string {
	return "aws_auditmanager_framework"
}

// LocalName returns the local name for [DataAuditmanagerFramework].
func (af *DataAuditmanagerFramework) LocalName() string {
	return af.Name
}

// Configuration returns the configuration (args) for [DataAuditmanagerFramework].
func (af *DataAuditmanagerFramework) Configuration() interface{} {
	return af.Args
}

// Attributes returns the attributes for [DataAuditmanagerFramework].
func (af *DataAuditmanagerFramework) Attributes() dataAuditmanagerFrameworkAttributes {
	return dataAuditmanagerFrameworkAttributes{ref: terra.ReferenceDataResource(af)}
}

// DataAuditmanagerFrameworkArgs contains the configurations for aws_auditmanager_framework.
type DataAuditmanagerFrameworkArgs struct {
	// FrameworkType: string, required
	FrameworkType terra.StringValue `hcl:"framework_type,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ControlSets: min=0
	ControlSets []dataauditmanagerframework.ControlSets `hcl:"control_sets,block" validate:"min=0"`
}
type dataAuditmanagerFrameworkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_auditmanager_framework.
func (af dataAuditmanagerFrameworkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("arn"))
}

// ComplianceType returns a reference to field compliance_type of aws_auditmanager_framework.
func (af dataAuditmanagerFrameworkAttributes) ComplianceType() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("compliance_type"))
}

// Description returns a reference to field description of aws_auditmanager_framework.
func (af dataAuditmanagerFrameworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("description"))
}

// FrameworkType returns a reference to field framework_type of aws_auditmanager_framework.
func (af dataAuditmanagerFrameworkAttributes) FrameworkType() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("framework_type"))
}

// Id returns a reference to field id of aws_auditmanager_framework.
func (af dataAuditmanagerFrameworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("id"))
}

// Name returns a reference to field name of aws_auditmanager_framework.
func (af dataAuditmanagerFrameworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_auditmanager_framework.
func (af dataAuditmanagerFrameworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](af.ref.Append("tags"))
}

func (af dataAuditmanagerFrameworkAttributes) ControlSets() terra.SetValue[dataauditmanagerframework.ControlSetsAttributes] {
	return terra.ReferenceAsSet[dataauditmanagerframework.ControlSetsAttributes](af.ref.Append("control_sets"))
}
