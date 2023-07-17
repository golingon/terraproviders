// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataimagebuilderdistributionconfigurations "github.com/golingon/terraproviders/aws/5.8.0/dataimagebuilderdistributionconfigurations"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataImagebuilderDistributionConfigurations creates a new instance of [DataImagebuilderDistributionConfigurations].
func NewDataImagebuilderDistributionConfigurations(name string, args DataImagebuilderDistributionConfigurationsArgs) *DataImagebuilderDistributionConfigurations {
	return &DataImagebuilderDistributionConfigurations{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataImagebuilderDistributionConfigurations)(nil)

// DataImagebuilderDistributionConfigurations represents the Terraform data resource aws_imagebuilder_distribution_configurations.
type DataImagebuilderDistributionConfigurations struct {
	Name string
	Args DataImagebuilderDistributionConfigurationsArgs
}

// DataSource returns the Terraform object type for [DataImagebuilderDistributionConfigurations].
func (idc *DataImagebuilderDistributionConfigurations) DataSource() string {
	return "aws_imagebuilder_distribution_configurations"
}

// LocalName returns the local name for [DataImagebuilderDistributionConfigurations].
func (idc *DataImagebuilderDistributionConfigurations) LocalName() string {
	return idc.Name
}

// Configuration returns the configuration (args) for [DataImagebuilderDistributionConfigurations].
func (idc *DataImagebuilderDistributionConfigurations) Configuration() interface{} {
	return idc.Args
}

// Attributes returns the attributes for [DataImagebuilderDistributionConfigurations].
func (idc *DataImagebuilderDistributionConfigurations) Attributes() dataImagebuilderDistributionConfigurationsAttributes {
	return dataImagebuilderDistributionConfigurationsAttributes{ref: terra.ReferenceDataResource(idc)}
}

// DataImagebuilderDistributionConfigurationsArgs contains the configurations for aws_imagebuilder_distribution_configurations.
type DataImagebuilderDistributionConfigurationsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filter: min=0
	Filter []dataimagebuilderdistributionconfigurations.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataImagebuilderDistributionConfigurationsAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_imagebuilder_distribution_configurations.
func (idc dataImagebuilderDistributionConfigurationsAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](idc.ref.Append("arns"))
}

// Id returns a reference to field id of aws_imagebuilder_distribution_configurations.
func (idc dataImagebuilderDistributionConfigurationsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(idc.ref.Append("id"))
}

// Names returns a reference to field names of aws_imagebuilder_distribution_configurations.
func (idc dataImagebuilderDistributionConfigurationsAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](idc.ref.Append("names"))
}

func (idc dataImagebuilderDistributionConfigurationsAttributes) Filter() terra.SetValue[dataimagebuilderdistributionconfigurations.FilterAttributes] {
	return terra.ReferenceAsSet[dataimagebuilderdistributionconfigurations.FilterAttributes](idc.ref.Append("filter"))
}
