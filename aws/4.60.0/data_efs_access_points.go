// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataEfsAccessPoints(name string, args DataEfsAccessPointsArgs) *DataEfsAccessPoints {
	return &DataEfsAccessPoints{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEfsAccessPoints)(nil)

type DataEfsAccessPoints struct {
	Name string
	Args DataEfsAccessPointsArgs
}

func (eap *DataEfsAccessPoints) DataSource() string {
	return "aws_efs_access_points"
}

func (eap *DataEfsAccessPoints) LocalName() string {
	return eap.Name
}

func (eap *DataEfsAccessPoints) Configuration() interface{} {
	return eap.Args
}

func (eap *DataEfsAccessPoints) Attributes() dataEfsAccessPointsAttributes {
	return dataEfsAccessPointsAttributes{ref: terra.ReferenceDataResource(eap)}
}

type DataEfsAccessPointsArgs struct {
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataEfsAccessPointsAttributes struct {
	ref terra.Reference
}

func (eap dataEfsAccessPointsAttributes) Arns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](eap.ref.Append("arns"))
}

func (eap dataEfsAccessPointsAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceString(eap.ref.Append("file_system_id"))
}

func (eap dataEfsAccessPointsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(eap.ref.Append("id"))
}

func (eap dataEfsAccessPointsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](eap.ref.Append("ids"))
}
