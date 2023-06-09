// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datassmmaintenancewindows "github.com/golingon/terraproviders/aws/4.60.0/datassmmaintenancewindows"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSsmMaintenanceWindows creates a new instance of [DataSsmMaintenanceWindows].
func NewDataSsmMaintenanceWindows(name string, args DataSsmMaintenanceWindowsArgs) *DataSsmMaintenanceWindows {
	return &DataSsmMaintenanceWindows{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSsmMaintenanceWindows)(nil)

// DataSsmMaintenanceWindows represents the Terraform data resource aws_ssm_maintenance_windows.
type DataSsmMaintenanceWindows struct {
	Name string
	Args DataSsmMaintenanceWindowsArgs
}

// DataSource returns the Terraform object type for [DataSsmMaintenanceWindows].
func (smw *DataSsmMaintenanceWindows) DataSource() string {
	return "aws_ssm_maintenance_windows"
}

// LocalName returns the local name for [DataSsmMaintenanceWindows].
func (smw *DataSsmMaintenanceWindows) LocalName() string {
	return smw.Name
}

// Configuration returns the configuration (args) for [DataSsmMaintenanceWindows].
func (smw *DataSsmMaintenanceWindows) Configuration() interface{} {
	return smw.Args
}

// Attributes returns the attributes for [DataSsmMaintenanceWindows].
func (smw *DataSsmMaintenanceWindows) Attributes() dataSsmMaintenanceWindowsAttributes {
	return dataSsmMaintenanceWindowsAttributes{ref: terra.ReferenceDataResource(smw)}
}

// DataSsmMaintenanceWindowsArgs contains the configurations for aws_ssm_maintenance_windows.
type DataSsmMaintenanceWindowsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filter: min=0
	Filter []datassmmaintenancewindows.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataSsmMaintenanceWindowsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ssm_maintenance_windows.
func (smw dataSsmMaintenanceWindowsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smw.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ssm_maintenance_windows.
func (smw dataSsmMaintenanceWindowsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](smw.ref.Append("ids"))
}

func (smw dataSsmMaintenanceWindowsAttributes) Filter() terra.SetValue[datassmmaintenancewindows.FilterAttributes] {
	return terra.ReferenceAsSet[datassmmaintenancewindows.FilterAttributes](smw.ref.Append("filter"))
}
