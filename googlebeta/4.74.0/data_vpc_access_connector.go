// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datavpcaccessconnector "github.com/golingon/terraproviders/googlebeta/4.74.0/datavpcaccessconnector"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcAccessConnector creates a new instance of [DataVpcAccessConnector].
func NewDataVpcAccessConnector(name string, args DataVpcAccessConnectorArgs) *DataVpcAccessConnector {
	return &DataVpcAccessConnector{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcAccessConnector)(nil)

// DataVpcAccessConnector represents the Terraform data resource google_vpc_access_connector.
type DataVpcAccessConnector struct {
	Name string
	Args DataVpcAccessConnectorArgs
}

// DataSource returns the Terraform object type for [DataVpcAccessConnector].
func (vac *DataVpcAccessConnector) DataSource() string {
	return "google_vpc_access_connector"
}

// LocalName returns the local name for [DataVpcAccessConnector].
func (vac *DataVpcAccessConnector) LocalName() string {
	return vac.Name
}

// Configuration returns the configuration (args) for [DataVpcAccessConnector].
func (vac *DataVpcAccessConnector) Configuration() interface{} {
	return vac.Args
}

// Attributes returns the attributes for [DataVpcAccessConnector].
func (vac *DataVpcAccessConnector) Attributes() dataVpcAccessConnectorAttributes {
	return dataVpcAccessConnectorAttributes{ref: terra.ReferenceDataResource(vac)}
}

// DataVpcAccessConnectorArgs contains the configurations for google_vpc_access_connector.
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

// ConnectedProjects returns a reference to field connected_projects of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) ConnectedProjects() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vac.ref.Append("connected_projects"))
}

// Id returns a reference to field id of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("id"))
}

// IpCidrRange returns a reference to field ip_cidr_range of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("ip_cidr_range"))
}

// MachineType returns a reference to field machine_type of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("machine_type"))
}

// MaxInstances returns a reference to field max_instances of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) MaxInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(vac.ref.Append("max_instances"))
}

// MaxThroughput returns a reference to field max_throughput of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) MaxThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(vac.ref.Append("max_throughput"))
}

// MinInstances returns a reference to field min_instances of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) MinInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(vac.ref.Append("min_instances"))
}

// MinThroughput returns a reference to field min_throughput of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) MinThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(vac.ref.Append("min_throughput"))
}

// Name returns a reference to field name of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("name"))
}

// Network returns a reference to field network of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("network"))
}

// Project returns a reference to field project of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("project"))
}

// Region returns a reference to field region of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("self_link"))
}

// State returns a reference to field state of google_vpc_access_connector.
func (vac dataVpcAccessConnectorAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vac.ref.Append("state"))
}

func (vac dataVpcAccessConnectorAttributes) Subnet() terra.ListValue[datavpcaccessconnector.SubnetAttributes] {
	return terra.ReferenceAsList[datavpcaccessconnector.SubnetAttributes](vac.ref.Append("subnet"))
}
