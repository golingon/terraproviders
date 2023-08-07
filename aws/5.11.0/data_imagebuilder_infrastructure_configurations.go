// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataimagebuilderinfrastructureconfigurations "github.com/golingon/terraproviders/aws/5.11.0/dataimagebuilderinfrastructureconfigurations"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataImagebuilderInfrastructureConfigurations creates a new instance of [DataImagebuilderInfrastructureConfigurations].
func NewDataImagebuilderInfrastructureConfigurations(name string, args DataImagebuilderInfrastructureConfigurationsArgs) *DataImagebuilderInfrastructureConfigurations {
	return &DataImagebuilderInfrastructureConfigurations{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataImagebuilderInfrastructureConfigurations)(nil)

// DataImagebuilderInfrastructureConfigurations represents the Terraform data resource aws_imagebuilder_infrastructure_configurations.
type DataImagebuilderInfrastructureConfigurations struct {
	Name string
	Args DataImagebuilderInfrastructureConfigurationsArgs
}

// DataSource returns the Terraform object type for [DataImagebuilderInfrastructureConfigurations].
func (iic *DataImagebuilderInfrastructureConfigurations) DataSource() string {
	return "aws_imagebuilder_infrastructure_configurations"
}

// LocalName returns the local name for [DataImagebuilderInfrastructureConfigurations].
func (iic *DataImagebuilderInfrastructureConfigurations) LocalName() string {
	return iic.Name
}

// Configuration returns the configuration (args) for [DataImagebuilderInfrastructureConfigurations].
func (iic *DataImagebuilderInfrastructureConfigurations) Configuration() interface{} {
	return iic.Args
}

// Attributes returns the attributes for [DataImagebuilderInfrastructureConfigurations].
func (iic *DataImagebuilderInfrastructureConfigurations) Attributes() dataImagebuilderInfrastructureConfigurationsAttributes {
	return dataImagebuilderInfrastructureConfigurationsAttributes{ref: terra.ReferenceDataResource(iic)}
}

// DataImagebuilderInfrastructureConfigurationsArgs contains the configurations for aws_imagebuilder_infrastructure_configurations.
type DataImagebuilderInfrastructureConfigurationsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filter: min=0
	Filter []dataimagebuilderinfrastructureconfigurations.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataImagebuilderInfrastructureConfigurationsAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_imagebuilder_infrastructure_configurations.
func (iic dataImagebuilderInfrastructureConfigurationsAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iic.ref.Append("arns"))
}

// Id returns a reference to field id of aws_imagebuilder_infrastructure_configurations.
func (iic dataImagebuilderInfrastructureConfigurationsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("id"))
}

// Names returns a reference to field names of aws_imagebuilder_infrastructure_configurations.
func (iic dataImagebuilderInfrastructureConfigurationsAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iic.ref.Append("names"))
}

func (iic dataImagebuilderInfrastructureConfigurationsAttributes) Filter() terra.SetValue[dataimagebuilderinfrastructureconfigurations.FilterAttributes] {
	return terra.ReferenceAsSet[dataimagebuilderinfrastructureconfigurations.FilterAttributes](iic.ref.Append("filter"))
}
