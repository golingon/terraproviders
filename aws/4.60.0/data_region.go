// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataRegion(name string, args DataRegionArgs) *DataRegion {
	return &DataRegion{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRegion)(nil)

type DataRegion struct {
	Name string
	Args DataRegionArgs
}

func (r *DataRegion) DataSource() string {
	return "aws_region"
}

func (r *DataRegion) LocalName() string {
	return r.Name
}

func (r *DataRegion) Configuration() interface{} {
	return r.Args
}

func (r *DataRegion) Attributes() dataRegionAttributes {
	return dataRegionAttributes{ref: terra.ReferenceDataResource(r)}
}

type DataRegionArgs struct {
	// Endpoint: string, optional
	Endpoint terra.StringValue `hcl:"endpoint,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}
type dataRegionAttributes struct {
	ref terra.Reference
}

func (r dataRegionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("description"))
}

func (r dataRegionAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("endpoint"))
}

func (r dataRegionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("id"))
}

func (r dataRegionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("name"))
}
