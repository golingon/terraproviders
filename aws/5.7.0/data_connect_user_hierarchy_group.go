// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataconnectuserhierarchygroup "github.com/golingon/terraproviders/aws/5.7.0/dataconnectuserhierarchygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataConnectUserHierarchyGroup creates a new instance of [DataConnectUserHierarchyGroup].
func NewDataConnectUserHierarchyGroup(name string, args DataConnectUserHierarchyGroupArgs) *DataConnectUserHierarchyGroup {
	return &DataConnectUserHierarchyGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectUserHierarchyGroup)(nil)

// DataConnectUserHierarchyGroup represents the Terraform data resource aws_connect_user_hierarchy_group.
type DataConnectUserHierarchyGroup struct {
	Name string
	Args DataConnectUserHierarchyGroupArgs
}

// DataSource returns the Terraform object type for [DataConnectUserHierarchyGroup].
func (cuhg *DataConnectUserHierarchyGroup) DataSource() string {
	return "aws_connect_user_hierarchy_group"
}

// LocalName returns the local name for [DataConnectUserHierarchyGroup].
func (cuhg *DataConnectUserHierarchyGroup) LocalName() string {
	return cuhg.Name
}

// Configuration returns the configuration (args) for [DataConnectUserHierarchyGroup].
func (cuhg *DataConnectUserHierarchyGroup) Configuration() interface{} {
	return cuhg.Args
}

// Attributes returns the attributes for [DataConnectUserHierarchyGroup].
func (cuhg *DataConnectUserHierarchyGroup) Attributes() dataConnectUserHierarchyGroupAttributes {
	return dataConnectUserHierarchyGroupAttributes{ref: terra.ReferenceDataResource(cuhg)}
}

// DataConnectUserHierarchyGroupArgs contains the configurations for aws_connect_user_hierarchy_group.
type DataConnectUserHierarchyGroupArgs struct {
	// HierarchyGroupId: string, optional
	HierarchyGroupId terra.StringValue `hcl:"hierarchy_group_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// HierarchyPath: min=0
	HierarchyPath []dataconnectuserhierarchygroup.HierarchyPath `hcl:"hierarchy_path,block" validate:"min=0"`
}
type dataConnectUserHierarchyGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_user_hierarchy_group.
func (cuhg dataConnectUserHierarchyGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("arn"))
}

// HierarchyGroupId returns a reference to field hierarchy_group_id of aws_connect_user_hierarchy_group.
func (cuhg dataConnectUserHierarchyGroupAttributes) HierarchyGroupId() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("hierarchy_group_id"))
}

// Id returns a reference to field id of aws_connect_user_hierarchy_group.
func (cuhg dataConnectUserHierarchyGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_user_hierarchy_group.
func (cuhg dataConnectUserHierarchyGroupAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("instance_id"))
}

// LevelId returns a reference to field level_id of aws_connect_user_hierarchy_group.
func (cuhg dataConnectUserHierarchyGroupAttributes) LevelId() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("level_id"))
}

// Name returns a reference to field name of aws_connect_user_hierarchy_group.
func (cuhg dataConnectUserHierarchyGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_connect_user_hierarchy_group.
func (cuhg dataConnectUserHierarchyGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cuhg.ref.Append("tags"))
}

func (cuhg dataConnectUserHierarchyGroupAttributes) HierarchyPath() terra.ListValue[dataconnectuserhierarchygroup.HierarchyPathAttributes] {
	return terra.ReferenceAsList[dataconnectuserhierarchygroup.HierarchyPathAttributes](cuhg.ref.Append("hierarchy_path"))
}
