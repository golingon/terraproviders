// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataconnectuserhierarchystructure "github.com/golingon/terraproviders/aws/4.60.0/dataconnectuserhierarchystructure"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataConnectUserHierarchyStructure(name string, args DataConnectUserHierarchyStructureArgs) *DataConnectUserHierarchyStructure {
	return &DataConnectUserHierarchyStructure{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectUserHierarchyStructure)(nil)

type DataConnectUserHierarchyStructure struct {
	Name string
	Args DataConnectUserHierarchyStructureArgs
}

func (cuhs *DataConnectUserHierarchyStructure) DataSource() string {
	return "aws_connect_user_hierarchy_structure"
}

func (cuhs *DataConnectUserHierarchyStructure) LocalName() string {
	return cuhs.Name
}

func (cuhs *DataConnectUserHierarchyStructure) Configuration() interface{} {
	return cuhs.Args
}

func (cuhs *DataConnectUserHierarchyStructure) Attributes() dataConnectUserHierarchyStructureAttributes {
	return dataConnectUserHierarchyStructureAttributes{ref: terra.ReferenceDataResource(cuhs)}
}

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

func (cuhs dataConnectUserHierarchyStructureAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cuhs.ref.Append("id"))
}

func (cuhs dataConnectUserHierarchyStructureAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceString(cuhs.ref.Append("instance_id"))
}

func (cuhs dataConnectUserHierarchyStructureAttributes) HierarchyStructure() terra.ListValue[dataconnectuserhierarchystructure.HierarchyStructureAttributes] {
	return terra.ReferenceList[dataconnectuserhierarchystructure.HierarchyStructureAttributes](cuhs.ref.Append("hierarchy_structure"))
}
