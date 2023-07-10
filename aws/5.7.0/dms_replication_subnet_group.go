// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dmsreplicationsubnetgroup "github.com/golingon/terraproviders/aws/5.7.0/dmsreplicationsubnetgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDmsReplicationSubnetGroup creates a new instance of [DmsReplicationSubnetGroup].
func NewDmsReplicationSubnetGroup(name string, args DmsReplicationSubnetGroupArgs) *DmsReplicationSubnetGroup {
	return &DmsReplicationSubnetGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DmsReplicationSubnetGroup)(nil)

// DmsReplicationSubnetGroup represents the Terraform resource aws_dms_replication_subnet_group.
type DmsReplicationSubnetGroup struct {
	Name      string
	Args      DmsReplicationSubnetGroupArgs
	state     *dmsReplicationSubnetGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DmsReplicationSubnetGroup].
func (drsg *DmsReplicationSubnetGroup) Type() string {
	return "aws_dms_replication_subnet_group"
}

// LocalName returns the local name for [DmsReplicationSubnetGroup].
func (drsg *DmsReplicationSubnetGroup) LocalName() string {
	return drsg.Name
}

// Configuration returns the configuration (args) for [DmsReplicationSubnetGroup].
func (drsg *DmsReplicationSubnetGroup) Configuration() interface{} {
	return drsg.Args
}

// DependOn is used for other resources to depend on [DmsReplicationSubnetGroup].
func (drsg *DmsReplicationSubnetGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(drsg)
}

// Dependencies returns the list of resources [DmsReplicationSubnetGroup] depends_on.
func (drsg *DmsReplicationSubnetGroup) Dependencies() terra.Dependencies {
	return drsg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DmsReplicationSubnetGroup].
func (drsg *DmsReplicationSubnetGroup) LifecycleManagement() *terra.Lifecycle {
	return drsg.Lifecycle
}

// Attributes returns the attributes for [DmsReplicationSubnetGroup].
func (drsg *DmsReplicationSubnetGroup) Attributes() dmsReplicationSubnetGroupAttributes {
	return dmsReplicationSubnetGroupAttributes{ref: terra.ReferenceResource(drsg)}
}

// ImportState imports the given attribute values into [DmsReplicationSubnetGroup]'s state.
func (drsg *DmsReplicationSubnetGroup) ImportState(av io.Reader) error {
	drsg.state = &dmsReplicationSubnetGroupState{}
	if err := json.NewDecoder(av).Decode(drsg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", drsg.Type(), drsg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DmsReplicationSubnetGroup] has state.
func (drsg *DmsReplicationSubnetGroup) State() (*dmsReplicationSubnetGroupState, bool) {
	return drsg.state, drsg.state != nil
}

// StateMust returns the state for [DmsReplicationSubnetGroup]. Panics if the state is nil.
func (drsg *DmsReplicationSubnetGroup) StateMust() *dmsReplicationSubnetGroupState {
	if drsg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", drsg.Type(), drsg.LocalName()))
	}
	return drsg.state
}

// DmsReplicationSubnetGroupArgs contains the configurations for aws_dms_replication_subnet_group.
type DmsReplicationSubnetGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReplicationSubnetGroupDescription: string, required
	ReplicationSubnetGroupDescription terra.StringValue `hcl:"replication_subnet_group_description,attr" validate:"required"`
	// ReplicationSubnetGroupId: string, required
	ReplicationSubnetGroupId terra.StringValue `hcl:"replication_subnet_group_id,attr" validate:"required"`
	// SubnetIds: set of string, required
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *dmsreplicationsubnetgroup.Timeouts `hcl:"timeouts,block"`
}
type dmsReplicationSubnetGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_dms_replication_subnet_group.
func (drsg dmsReplicationSubnetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("id"))
}

// ReplicationSubnetGroupArn returns a reference to field replication_subnet_group_arn of aws_dms_replication_subnet_group.
func (drsg dmsReplicationSubnetGroupAttributes) ReplicationSubnetGroupArn() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("replication_subnet_group_arn"))
}

// ReplicationSubnetGroupDescription returns a reference to field replication_subnet_group_description of aws_dms_replication_subnet_group.
func (drsg dmsReplicationSubnetGroupAttributes) ReplicationSubnetGroupDescription() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("replication_subnet_group_description"))
}

// ReplicationSubnetGroupId returns a reference to field replication_subnet_group_id of aws_dms_replication_subnet_group.
func (drsg dmsReplicationSubnetGroupAttributes) ReplicationSubnetGroupId() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("replication_subnet_group_id"))
}

// SubnetIds returns a reference to field subnet_ids of aws_dms_replication_subnet_group.
func (drsg dmsReplicationSubnetGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](drsg.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_dms_replication_subnet_group.
func (drsg dmsReplicationSubnetGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](drsg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dms_replication_subnet_group.
func (drsg dmsReplicationSubnetGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](drsg.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_dms_replication_subnet_group.
func (drsg dmsReplicationSubnetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(drsg.ref.Append("vpc_id"))
}

func (drsg dmsReplicationSubnetGroupAttributes) Timeouts() dmsreplicationsubnetgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dmsreplicationsubnetgroup.TimeoutsAttributes](drsg.ref.Append("timeouts"))
}

type dmsReplicationSubnetGroupState struct {
	Id                                string                                   `json:"id"`
	ReplicationSubnetGroupArn         string                                   `json:"replication_subnet_group_arn"`
	ReplicationSubnetGroupDescription string                                   `json:"replication_subnet_group_description"`
	ReplicationSubnetGroupId          string                                   `json:"replication_subnet_group_id"`
	SubnetIds                         []string                                 `json:"subnet_ids"`
	Tags                              map[string]string                        `json:"tags"`
	TagsAll                           map[string]string                        `json:"tags_all"`
	VpcId                             string                                   `json:"vpc_id"`
	Timeouts                          *dmsreplicationsubnetgroup.TimeoutsState `json:"timeouts"`
}
