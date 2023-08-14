// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapolicyvirtualmachineconfigurationassignment "github.com/golingon/terraproviders/azurerm/3.69.0/datapolicyvirtualmachineconfigurationassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPolicyVirtualMachineConfigurationAssignment creates a new instance of [DataPolicyVirtualMachineConfigurationAssignment].
func NewDataPolicyVirtualMachineConfigurationAssignment(name string, args DataPolicyVirtualMachineConfigurationAssignmentArgs) *DataPolicyVirtualMachineConfigurationAssignment {
	return &DataPolicyVirtualMachineConfigurationAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPolicyVirtualMachineConfigurationAssignment)(nil)

// DataPolicyVirtualMachineConfigurationAssignment represents the Terraform data resource azurerm_policy_virtual_machine_configuration_assignment.
type DataPolicyVirtualMachineConfigurationAssignment struct {
	Name string
	Args DataPolicyVirtualMachineConfigurationAssignmentArgs
}

// DataSource returns the Terraform object type for [DataPolicyVirtualMachineConfigurationAssignment].
func (pvmca *DataPolicyVirtualMachineConfigurationAssignment) DataSource() string {
	return "azurerm_policy_virtual_machine_configuration_assignment"
}

// LocalName returns the local name for [DataPolicyVirtualMachineConfigurationAssignment].
func (pvmca *DataPolicyVirtualMachineConfigurationAssignment) LocalName() string {
	return pvmca.Name
}

// Configuration returns the configuration (args) for [DataPolicyVirtualMachineConfigurationAssignment].
func (pvmca *DataPolicyVirtualMachineConfigurationAssignment) Configuration() interface{} {
	return pvmca.Args
}

// Attributes returns the attributes for [DataPolicyVirtualMachineConfigurationAssignment].
func (pvmca *DataPolicyVirtualMachineConfigurationAssignment) Attributes() dataPolicyVirtualMachineConfigurationAssignmentAttributes {
	return dataPolicyVirtualMachineConfigurationAssignmentAttributes{ref: terra.ReferenceDataResource(pvmca)}
}

// DataPolicyVirtualMachineConfigurationAssignmentArgs contains the configurations for azurerm_policy_virtual_machine_configuration_assignment.
type DataPolicyVirtualMachineConfigurationAssignmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// VirtualMachineName: string, required
	VirtualMachineName terra.StringValue `hcl:"virtual_machine_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datapolicyvirtualmachineconfigurationassignment.Timeouts `hcl:"timeouts,block"`
}
type dataPolicyVirtualMachineConfigurationAssignmentAttributes struct {
	ref terra.Reference
}

// AssignmentHash returns a reference to field assignment_hash of azurerm_policy_virtual_machine_configuration_assignment.
func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) AssignmentHash() terra.StringValue {
	return terra.ReferenceAsString(pvmca.ref.Append("assignment_hash"))
}

// ComplianceStatus returns a reference to field compliance_status of azurerm_policy_virtual_machine_configuration_assignment.
func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) ComplianceStatus() terra.StringValue {
	return terra.ReferenceAsString(pvmca.ref.Append("compliance_status"))
}

// ContentHash returns a reference to field content_hash of azurerm_policy_virtual_machine_configuration_assignment.
func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) ContentHash() terra.StringValue {
	return terra.ReferenceAsString(pvmca.ref.Append("content_hash"))
}

// ContentUri returns a reference to field content_uri of azurerm_policy_virtual_machine_configuration_assignment.
func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) ContentUri() terra.StringValue {
	return terra.ReferenceAsString(pvmca.ref.Append("content_uri"))
}

// Id returns a reference to field id of azurerm_policy_virtual_machine_configuration_assignment.
func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pvmca.ref.Append("id"))
}

// LastComplianceStatusChecked returns a reference to field last_compliance_status_checked of azurerm_policy_virtual_machine_configuration_assignment.
func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) LastComplianceStatusChecked() terra.StringValue {
	return terra.ReferenceAsString(pvmca.ref.Append("last_compliance_status_checked"))
}

// LatestReportId returns a reference to field latest_report_id of azurerm_policy_virtual_machine_configuration_assignment.
func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) LatestReportId() terra.StringValue {
	return terra.ReferenceAsString(pvmca.ref.Append("latest_report_id"))
}

// Name returns a reference to field name of azurerm_policy_virtual_machine_configuration_assignment.
func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pvmca.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_policy_virtual_machine_configuration_assignment.
func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pvmca.ref.Append("resource_group_name"))
}

// VirtualMachineName returns a reference to field virtual_machine_name of azurerm_policy_virtual_machine_configuration_assignment.
func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) VirtualMachineName() terra.StringValue {
	return terra.ReferenceAsString(pvmca.ref.Append("virtual_machine_name"))
}

func (pvmca dataPolicyVirtualMachineConfigurationAssignmentAttributes) Timeouts() datapolicyvirtualmachineconfigurationassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datapolicyvirtualmachineconfigurationassignment.TimeoutsAttributes](pvmca.ref.Append("timeouts"))
}
