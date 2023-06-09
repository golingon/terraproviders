// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadatafactory "github.com/golingon/terraproviders/azurerm/3.58.0/datadatafactory"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDataFactory creates a new instance of [DataDataFactory].
func NewDataDataFactory(name string, args DataDataFactoryArgs) *DataDataFactory {
	return &DataDataFactory{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataFactory)(nil)

// DataDataFactory represents the Terraform data resource azurerm_data_factory.
type DataDataFactory struct {
	Name string
	Args DataDataFactoryArgs
}

// DataSource returns the Terraform object type for [DataDataFactory].
func (df *DataDataFactory) DataSource() string {
	return "azurerm_data_factory"
}

// LocalName returns the local name for [DataDataFactory].
func (df *DataDataFactory) LocalName() string {
	return df.Name
}

// Configuration returns the configuration (args) for [DataDataFactory].
func (df *DataDataFactory) Configuration() interface{} {
	return df.Args
}

// Attributes returns the attributes for [DataDataFactory].
func (df *DataDataFactory) Attributes() dataDataFactoryAttributes {
	return dataDataFactoryAttributes{ref: terra.ReferenceDataResource(df)}
}

// DataDataFactoryArgs contains the configurations for azurerm_data_factory.
type DataDataFactoryArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// GithubConfiguration: min=0
	GithubConfiguration []datadatafactory.GithubConfiguration `hcl:"github_configuration,block" validate:"min=0"`
	// Identity: min=0
	Identity []datadatafactory.Identity `hcl:"identity,block" validate:"min=0"`
	// VstsConfiguration: min=0
	VstsConfiguration []datadatafactory.VstsConfiguration `hcl:"vsts_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadatafactory.Timeouts `hcl:"timeouts,block"`
}
type dataDataFactoryAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_data_factory.
func (df dataDataFactoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_data_factory.
func (df dataDataFactoryAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_data_factory.
func (df dataDataFactoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_data_factory.
func (df dataDataFactoryAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_data_factory.
func (df dataDataFactoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](df.ref.Append("tags"))
}

func (df dataDataFactoryAttributes) GithubConfiguration() terra.ListValue[datadatafactory.GithubConfigurationAttributes] {
	return terra.ReferenceAsList[datadatafactory.GithubConfigurationAttributes](df.ref.Append("github_configuration"))
}

func (df dataDataFactoryAttributes) Identity() terra.ListValue[datadatafactory.IdentityAttributes] {
	return terra.ReferenceAsList[datadatafactory.IdentityAttributes](df.ref.Append("identity"))
}

func (df dataDataFactoryAttributes) VstsConfiguration() terra.ListValue[datadatafactory.VstsConfigurationAttributes] {
	return terra.ReferenceAsList[datadatafactory.VstsConfigurationAttributes](df.ref.Append("vsts_configuration"))
}

func (df dataDataFactoryAttributes) Timeouts() datadatafactory.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadatafactory.TimeoutsAttributes](df.ref.Append("timeouts"))
}
