// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataDefaultTags(name string, args DataDefaultTagsArgs) *DataDefaultTags {
	return &DataDefaultTags{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDefaultTags)(nil)

type DataDefaultTags struct {
	Name string
	Args DataDefaultTagsArgs
}

func (dt *DataDefaultTags) DataSource() string {
	return "aws_default_tags"
}

func (dt *DataDefaultTags) LocalName() string {
	return dt.Name
}

func (dt *DataDefaultTags) Configuration() interface{} {
	return dt.Args
}

func (dt *DataDefaultTags) Attributes() dataDefaultTagsAttributes {
	return dataDefaultTagsAttributes{ref: terra.ReferenceDataResource(dt)}
}

type DataDefaultTagsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataDefaultTagsAttributes struct {
	ref terra.Reference
}

func (dt dataDefaultTagsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dt.ref.Append("id"))
}

func (dt dataDefaultTagsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dt.ref.Append("tags"))
}
