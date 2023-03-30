// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataAppconfigConfigurationProfiles(name string, args DataAppconfigConfigurationProfilesArgs) *DataAppconfigConfigurationProfiles {
	return &DataAppconfigConfigurationProfiles{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppconfigConfigurationProfiles)(nil)

type DataAppconfigConfigurationProfiles struct {
	Name string
	Args DataAppconfigConfigurationProfilesArgs
}

func (acp *DataAppconfigConfigurationProfiles) DataSource() string {
	return "aws_appconfig_configuration_profiles"
}

func (acp *DataAppconfigConfigurationProfiles) LocalName() string {
	return acp.Name
}

func (acp *DataAppconfigConfigurationProfiles) Configuration() interface{} {
	return acp.Args
}

func (acp *DataAppconfigConfigurationProfiles) Attributes() dataAppconfigConfigurationProfilesAttributes {
	return dataAppconfigConfigurationProfilesAttributes{ref: terra.ReferenceDataResource(acp)}
}

type DataAppconfigConfigurationProfilesArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataAppconfigConfigurationProfilesAttributes struct {
	ref terra.Reference
}

func (acp dataAppconfigConfigurationProfilesAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceString(acp.ref.Append("application_id"))
}

func (acp dataAppconfigConfigurationProfilesAttributes) ConfigurationProfileIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](acp.ref.Append("configuration_profile_ids"))
}

func (acp dataAppconfigConfigurationProfilesAttributes) Id() terra.StringValue {
	return terra.ReferenceString(acp.ref.Append("id"))
}
