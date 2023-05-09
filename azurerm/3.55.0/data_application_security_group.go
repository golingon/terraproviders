// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapplicationsecuritygroup "github.com/golingon/terraproviders/azurerm/3.55.0/dataapplicationsecuritygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApplicationSecurityGroup creates a new instance of [DataApplicationSecurityGroup].
func NewDataApplicationSecurityGroup(name string, args DataApplicationSecurityGroupArgs) *DataApplicationSecurityGroup {
	return &DataApplicationSecurityGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApplicationSecurityGroup)(nil)

// DataApplicationSecurityGroup represents the Terraform data resource azurerm_application_security_group.
type DataApplicationSecurityGroup struct {
	Name string
	Args DataApplicationSecurityGroupArgs
}

// DataSource returns the Terraform object type for [DataApplicationSecurityGroup].
func (asg *DataApplicationSecurityGroup) DataSource() string {
	return "azurerm_application_security_group"
}

// LocalName returns the local name for [DataApplicationSecurityGroup].
func (asg *DataApplicationSecurityGroup) LocalName() string {
	return asg.Name
}

// Configuration returns the configuration (args) for [DataApplicationSecurityGroup].
func (asg *DataApplicationSecurityGroup) Configuration() interface{} {
	return asg.Args
}

// Attributes returns the attributes for [DataApplicationSecurityGroup].
func (asg *DataApplicationSecurityGroup) Attributes() dataApplicationSecurityGroupAttributes {
	return dataApplicationSecurityGroupAttributes{ref: terra.ReferenceDataResource(asg)}
}

// DataApplicationSecurityGroupArgs contains the configurations for azurerm_application_security_group.
type DataApplicationSecurityGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataapplicationsecuritygroup.Timeouts `hcl:"timeouts,block"`
}
type dataApplicationSecurityGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_application_security_group.
func (asg dataApplicationSecurityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_application_security_group.
func (asg dataApplicationSecurityGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_security_group.
func (asg dataApplicationSecurityGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_security_group.
func (asg dataApplicationSecurityGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_application_security_group.
func (asg dataApplicationSecurityGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asg.ref.Append("tags"))
}

func (asg dataApplicationSecurityGroupAttributes) Timeouts() dataapplicationsecuritygroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapplicationsecuritygroup.TimeoutsAttributes](asg.ref.Append("timeouts"))
}
