// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataOutpostsOutpostInstanceTypes(name string, args DataOutpostsOutpostInstanceTypesArgs) *DataOutpostsOutpostInstanceTypes {
	return &DataOutpostsOutpostInstanceTypes{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOutpostsOutpostInstanceTypes)(nil)

type DataOutpostsOutpostInstanceTypes struct {
	Name string
	Args DataOutpostsOutpostInstanceTypesArgs
}

func (ooit *DataOutpostsOutpostInstanceTypes) DataSource() string {
	return "aws_outposts_outpost_instance_types"
}

func (ooit *DataOutpostsOutpostInstanceTypes) LocalName() string {
	return ooit.Name
}

func (ooit *DataOutpostsOutpostInstanceTypes) Configuration() interface{} {
	return ooit.Args
}

func (ooit *DataOutpostsOutpostInstanceTypes) Attributes() dataOutpostsOutpostInstanceTypesAttributes {
	return dataOutpostsOutpostInstanceTypesAttributes{ref: terra.ReferenceDataResource(ooit)}
}

type DataOutpostsOutpostInstanceTypesArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataOutpostsOutpostInstanceTypesAttributes struct {
	ref terra.Reference
}

func (ooit dataOutpostsOutpostInstanceTypesAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ooit.ref.Append("arn"))
}

func (ooit dataOutpostsOutpostInstanceTypesAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ooit.ref.Append("id"))
}

func (ooit dataOutpostsOutpostInstanceTypesAttributes) InstanceTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ooit.ref.Append("instance_types"))
}
