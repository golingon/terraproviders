// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	connectuserhierarchygroup "github.com/golingon/terraproviders/aws/4.66.1/connectuserhierarchygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConnectUserHierarchyGroup creates a new instance of [ConnectUserHierarchyGroup].
func NewConnectUserHierarchyGroup(name string, args ConnectUserHierarchyGroupArgs) *ConnectUserHierarchyGroup {
	return &ConnectUserHierarchyGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectUserHierarchyGroup)(nil)

// ConnectUserHierarchyGroup represents the Terraform resource aws_connect_user_hierarchy_group.
type ConnectUserHierarchyGroup struct {
	Name      string
	Args      ConnectUserHierarchyGroupArgs
	state     *connectUserHierarchyGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectUserHierarchyGroup].
func (cuhg *ConnectUserHierarchyGroup) Type() string {
	return "aws_connect_user_hierarchy_group"
}

// LocalName returns the local name for [ConnectUserHierarchyGroup].
func (cuhg *ConnectUserHierarchyGroup) LocalName() string {
	return cuhg.Name
}

// Configuration returns the configuration (args) for [ConnectUserHierarchyGroup].
func (cuhg *ConnectUserHierarchyGroup) Configuration() interface{} {
	return cuhg.Args
}

// DependOn is used for other resources to depend on [ConnectUserHierarchyGroup].
func (cuhg *ConnectUserHierarchyGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cuhg)
}

// Dependencies returns the list of resources [ConnectUserHierarchyGroup] depends_on.
func (cuhg *ConnectUserHierarchyGroup) Dependencies() terra.Dependencies {
	return cuhg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectUserHierarchyGroup].
func (cuhg *ConnectUserHierarchyGroup) LifecycleManagement() *terra.Lifecycle {
	return cuhg.Lifecycle
}

// Attributes returns the attributes for [ConnectUserHierarchyGroup].
func (cuhg *ConnectUserHierarchyGroup) Attributes() connectUserHierarchyGroupAttributes {
	return connectUserHierarchyGroupAttributes{ref: terra.ReferenceResource(cuhg)}
}

// ImportState imports the given attribute values into [ConnectUserHierarchyGroup]'s state.
func (cuhg *ConnectUserHierarchyGroup) ImportState(av io.Reader) error {
	cuhg.state = &connectUserHierarchyGroupState{}
	if err := json.NewDecoder(av).Decode(cuhg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cuhg.Type(), cuhg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectUserHierarchyGroup] has state.
func (cuhg *ConnectUserHierarchyGroup) State() (*connectUserHierarchyGroupState, bool) {
	return cuhg.state, cuhg.state != nil
}

// StateMust returns the state for [ConnectUserHierarchyGroup]. Panics if the state is nil.
func (cuhg *ConnectUserHierarchyGroup) StateMust() *connectUserHierarchyGroupState {
	if cuhg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cuhg.Type(), cuhg.LocalName()))
	}
	return cuhg.state
}

// ConnectUserHierarchyGroupArgs contains the configurations for aws_connect_user_hierarchy_group.
type ConnectUserHierarchyGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentGroupId: string, optional
	ParentGroupId terra.StringValue `hcl:"parent_group_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// HierarchyPath: min=0
	HierarchyPath []connectuserhierarchygroup.HierarchyPath `hcl:"hierarchy_path,block" validate:"min=0"`
}
type connectUserHierarchyGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_user_hierarchy_group.
func (cuhg connectUserHierarchyGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("arn"))
}

// HierarchyGroupId returns a reference to field hierarchy_group_id of aws_connect_user_hierarchy_group.
func (cuhg connectUserHierarchyGroupAttributes) HierarchyGroupId() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("hierarchy_group_id"))
}

// Id returns a reference to field id of aws_connect_user_hierarchy_group.
func (cuhg connectUserHierarchyGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_user_hierarchy_group.
func (cuhg connectUserHierarchyGroupAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("instance_id"))
}

// LevelId returns a reference to field level_id of aws_connect_user_hierarchy_group.
func (cuhg connectUserHierarchyGroupAttributes) LevelId() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("level_id"))
}

// Name returns a reference to field name of aws_connect_user_hierarchy_group.
func (cuhg connectUserHierarchyGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("name"))
}

// ParentGroupId returns a reference to field parent_group_id of aws_connect_user_hierarchy_group.
func (cuhg connectUserHierarchyGroupAttributes) ParentGroupId() terra.StringValue {
	return terra.ReferenceAsString(cuhg.ref.Append("parent_group_id"))
}

// Tags returns a reference to field tags of aws_connect_user_hierarchy_group.
func (cuhg connectUserHierarchyGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cuhg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_connect_user_hierarchy_group.
func (cuhg connectUserHierarchyGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cuhg.ref.Append("tags_all"))
}

func (cuhg connectUserHierarchyGroupAttributes) HierarchyPath() terra.ListValue[connectuserhierarchygroup.HierarchyPathAttributes] {
	return terra.ReferenceAsList[connectuserhierarchygroup.HierarchyPathAttributes](cuhg.ref.Append("hierarchy_path"))
}

type connectUserHierarchyGroupState struct {
	Arn              string                                         `json:"arn"`
	HierarchyGroupId string                                         `json:"hierarchy_group_id"`
	Id               string                                         `json:"id"`
	InstanceId       string                                         `json:"instance_id"`
	LevelId          string                                         `json:"level_id"`
	Name             string                                         `json:"name"`
	ParentGroupId    string                                         `json:"parent_group_id"`
	Tags             map[string]string                              `json:"tags"`
	TagsAll          map[string]string                              `json:"tags_all"`
	HierarchyPath    []connectuserhierarchygroup.HierarchyPathState `json:"hierarchy_path"`
}
