// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datavpcaccessconnector "github.com/golingon/terraproviders/google/4.59.0/datavpcaccessconnector"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataVpcAccessConnector(name string, args DataVpcAccessConnectorArgs) *DataVpcAccessConnector {
	return &DataVpcAccessConnector{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcAccessConnector)(nil)

type DataVpcAccessConnector struct {
	Name string
	Args DataVpcAccessConnectorArgs
}

func (vac *DataVpcAccessConnector) DataSource() string {
	return "google_vpc_access_connector"
}

func (vac *DataVpcAccessConnector) LocalName() string {
	return vac.Name
}

func (vac *DataVpcAccessConnector) Configuration() interface{} {
	return vac.Args
}

func (vac *DataVpcAccessConnector) Attributes() dataVpcAccessConnectorAttributes {
	return dataVpcAccessConnectorAttributes{ref: terra.ReferenceDataResource(vac)}
}

type DataVpcAccessConnectorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Subnet: min=0
	Subnet []datavpcaccessconnector.Subnet `hcl:"subnet,block" validate:"min=0"`
}
type dataVpcAccessConnectorAttributes struct {
	ref terra.Reference
}

func (vac dataVpcAccessConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("id"))
}

func (vac dataVpcAccessConnectorAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("ip_cidr_range"))
}

func (vac dataVpcAccessConnectorAttributes) MachineType() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("machine_type"))
}

func (vac dataVpcAccessConnectorAttributes) MaxInstances() terra.NumberValue {
	return terra.ReferenceNumber(vac.ref.Append("max_instances"))
}

func (vac dataVpcAccessConnectorAttributes) MaxThroughput() terra.NumberValue {
	return terra.ReferenceNumber(vac.ref.Append("max_throughput"))
}

func (vac dataVpcAccessConnectorAttributes) MinInstances() terra.NumberValue {
	return terra.ReferenceNumber(vac.ref.Append("min_instances"))
}

func (vac dataVpcAccessConnectorAttributes) MinThroughput() terra.NumberValue {
	return terra.ReferenceNumber(vac.ref.Append("min_throughput"))
}

func (vac dataVpcAccessConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("name"))
}

func (vac dataVpcAccessConnectorAttributes) Network() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("network"))
}

func (vac dataVpcAccessConnectorAttributes) Project() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("project"))
}

func (vac dataVpcAccessConnectorAttributes) Region() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("region"))
}

func (vac dataVpcAccessConnectorAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("self_link"))
}

func (vac dataVpcAccessConnectorAttributes) State() terra.StringValue {
	return terra.ReferenceString(vac.ref.Append("state"))
}

func (vac dataVpcAccessConnectorAttributes) Subnet() terra.ListValue[datavpcaccessconnector.SubnetAttributes] {
	return terra.ReferenceList[datavpcaccessconnector.SubnetAttributes](vac.ref.Append("subnet"))
}
