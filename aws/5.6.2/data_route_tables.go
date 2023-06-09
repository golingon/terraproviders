// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataroutetables "github.com/golingon/terraproviders/aws/5.6.2/dataroutetables"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRouteTables creates a new instance of [DataRouteTables].
func NewDataRouteTables(name string, args DataRouteTablesArgs) *DataRouteTables {
	return &DataRouteTables{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRouteTables)(nil)

// DataRouteTables represents the Terraform data resource aws_route_tables.
type DataRouteTables struct {
	Name string
	Args DataRouteTablesArgs
}

// DataSource returns the Terraform object type for [DataRouteTables].
func (rt *DataRouteTables) DataSource() string {
	return "aws_route_tables"
}

// LocalName returns the local name for [DataRouteTables].
func (rt *DataRouteTables) LocalName() string {
	return rt.Name
}

// Configuration returns the configuration (args) for [DataRouteTables].
func (rt *DataRouteTables) Configuration() interface{} {
	return rt.Args
}

// Attributes returns the attributes for [DataRouteTables].
func (rt *DataRouteTables) Attributes() dataRouteTablesAttributes {
	return dataRouteTablesAttributes{ref: terra.ReferenceDataResource(rt)}
}

// DataRouteTablesArgs contains the configurations for aws_route_tables.
type DataRouteTablesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// Filter: min=0
	Filter []dataroutetables.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataroutetables.Timeouts `hcl:"timeouts,block"`
}
type dataRouteTablesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_route_tables.
func (rt dataRouteTablesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_route_tables.
func (rt dataRouteTablesAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rt.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_route_tables.
func (rt dataRouteTablesAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rt.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_route_tables.
func (rt dataRouteTablesAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("vpc_id"))
}

func (rt dataRouteTablesAttributes) Filter() terra.SetValue[dataroutetables.FilterAttributes] {
	return terra.ReferenceAsSet[dataroutetables.FilterAttributes](rt.ref.Append("filter"))
}

func (rt dataRouteTablesAttributes) Timeouts() dataroutetables.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataroutetables.TimeoutsAttributes](rt.ref.Append("timeouts"))
}
