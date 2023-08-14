// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataconnectuserhierarchystructure "github.com/golingon/terraproviders/aws/5.12.0/dataconnectuserhierarchystructure"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataConnectUserHierarchyStructure creates a new instance of [DataConnectUserHierarchyStructure].
func NewDataConnectUserHierarchyStructure(name string, args DataConnectUserHierarchyStructureArgs) *DataConnectUserHierarchyStructure {
	return &DataConnectUserHierarchyStructure{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectUserHierarchyStructure)(nil)

// DataConnectUserHierarchyStructure represents the Terraform data resource aws_connect_user_hierarchy_structure.
type DataConnectUserHierarchyStructure struct {
	Name string
	Args DataConnectUserHierarchyStructureArgs
}

// DataSource returns the Terraform object type for [DataConnectUserHierarchyStructure].
func (cuhs *DataConnectUserHierarchyStructure) DataSource() string {
	return "aws_connect_user_hierarchy_structure"
}

// LocalName returns the local name for [DataConnectUserHierarchyStructure].
func (cuhs *DataConnectUserHierarchyStructure) LocalName() string {
	return cuhs.Name
}

// Configuration returns the configuration (args) for [DataConnectUserHierarchyStructure].
func (cuhs *DataConnectUserHierarchyStructure) Configuration() interface{} {
	return cuhs.Args
}

// Attributes returns the attributes for [DataConnectUserHierarchyStructure].
func (cuhs *DataConnectUserHierarchyStructure) Attributes() dataConnectUserHierarchyStructureAttributes {
	return dataConnectUserHierarchyStructureAttributes{ref: terra.ReferenceDataResource(cuhs)}
}

// DataConnectUserHierarchyStructureArgs contains the configurations for aws_connect_user_hierarchy_structure.
type DataConnectUserHierarchyStructureArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// HierarchyStructure: min=0
	HierarchyStructure []dataconnectuserhierarchystructure.HierarchyStructure `hcl:"hierarchy_structure,block" validate:"min=0"`
}
type dataConnectUserHierarchyStructureAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_connect_user_hierarchy_structure.
func (cuhs dataConnectUserHierarchyStructureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cuhs.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_user_hierarchy_structure.
func (cuhs dataConnectUserHierarchyStructureAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cuhs.ref.Append("instance_id"))
}

func (cuhs dataConnectUserHierarchyStructureAttributes) HierarchyStructure() terra.ListValue[dataconnectuserhierarchystructure.HierarchyStructureAttributes] {
	return terra.ReferenceAsList[dataconnectuserhierarchystructure.HierarchyStructureAttributes](cuhs.ref.Append("hierarchy_structure"))
}
