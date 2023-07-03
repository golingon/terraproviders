// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dmsreplicationinstance "github.com/golingon/terraproviders/aws/5.6.2/dmsreplicationinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDmsReplicationInstance creates a new instance of [DmsReplicationInstance].
func NewDmsReplicationInstance(name string, args DmsReplicationInstanceArgs) *DmsReplicationInstance {
	return &DmsReplicationInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DmsReplicationInstance)(nil)

// DmsReplicationInstance represents the Terraform resource aws_dms_replication_instance.
type DmsReplicationInstance struct {
	Name      string
	Args      DmsReplicationInstanceArgs
	state     *dmsReplicationInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DmsReplicationInstance].
func (dri *DmsReplicationInstance) Type() string {
	return "aws_dms_replication_instance"
}

// LocalName returns the local name for [DmsReplicationInstance].
func (dri *DmsReplicationInstance) LocalName() string {
	return dri.Name
}

// Configuration returns the configuration (args) for [DmsReplicationInstance].
func (dri *DmsReplicationInstance) Configuration() interface{} {
	return dri.Args
}

// DependOn is used for other resources to depend on [DmsReplicationInstance].
func (dri *DmsReplicationInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(dri)
}

// Dependencies returns the list of resources [DmsReplicationInstance] depends_on.
func (dri *DmsReplicationInstance) Dependencies() terra.Dependencies {
	return dri.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DmsReplicationInstance].
func (dri *DmsReplicationInstance) LifecycleManagement() *terra.Lifecycle {
	return dri.Lifecycle
}

// Attributes returns the attributes for [DmsReplicationInstance].
func (dri *DmsReplicationInstance) Attributes() dmsReplicationInstanceAttributes {
	return dmsReplicationInstanceAttributes{ref: terra.ReferenceResource(dri)}
}

// ImportState imports the given attribute values into [DmsReplicationInstance]'s state.
func (dri *DmsReplicationInstance) ImportState(av io.Reader) error {
	dri.state = &dmsReplicationInstanceState{}
	if err := json.NewDecoder(av).Decode(dri.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dri.Type(), dri.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DmsReplicationInstance] has state.
func (dri *DmsReplicationInstance) State() (*dmsReplicationInstanceState, bool) {
	return dri.state, dri.state != nil
}

// StateMust returns the state for [DmsReplicationInstance]. Panics if the state is nil.
func (dri *DmsReplicationInstance) StateMust() *dmsReplicationInstanceState {
	if dri.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dri.Type(), dri.LocalName()))
	}
	return dri.state
}

// DmsReplicationInstanceArgs contains the configurations for aws_dms_replication_instance.
type DmsReplicationInstanceArgs struct {
	// AllocatedStorage: number, optional
	AllocatedStorage terra.NumberValue `hcl:"allocated_storage,attr"`
	// AllowMajorVersionUpgrade: bool, optional
	AllowMajorVersionUpgrade terra.BoolValue `hcl:"allow_major_version_upgrade,attr"`
	// ApplyImmediately: bool, optional
	ApplyImmediately terra.BoolValue `hcl:"apply_immediately,attr"`
	// AutoMinorVersionUpgrade: bool, optional
	AutoMinorVersionUpgrade terra.BoolValue `hcl:"auto_minor_version_upgrade,attr"`
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// MultiAz: bool, optional
	MultiAz terra.BoolValue `hcl:"multi_az,attr"`
	// PreferredMaintenanceWindow: string, optional
	PreferredMaintenanceWindow terra.StringValue `hcl:"preferred_maintenance_window,attr"`
	// PubliclyAccessible: bool, optional
	PubliclyAccessible terra.BoolValue `hcl:"publicly_accessible,attr"`
	// ReplicationInstanceClass: string, required
	ReplicationInstanceClass terra.StringValue `hcl:"replication_instance_class,attr" validate:"required"`
	// ReplicationInstanceId: string, required
	ReplicationInstanceId terra.StringValue `hcl:"replication_instance_id,attr" validate:"required"`
	// ReplicationSubnetGroupId: string, optional
	ReplicationSubnetGroupId terra.StringValue `hcl:"replication_subnet_group_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// Timeouts: optional
	Timeouts *dmsreplicationinstance.Timeouts `hcl:"timeouts,block"`
}
type dmsReplicationInstanceAttributes struct {
	ref terra.Reference
}

// AllocatedStorage returns a reference to field allocated_storage of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(dri.ref.Append("allocated_storage"))
}

// AllowMajorVersionUpgrade returns a reference to field allow_major_version_upgrade of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) AllowMajorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(dri.ref.Append("allow_major_version_upgrade"))
}

// ApplyImmediately returns a reference to field apply_immediately of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceAsBool(dri.ref.Append("apply_immediately"))
}

// AutoMinorVersionUpgrade returns a reference to field auto_minor_version_upgrade of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) AutoMinorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(dri.ref.Append("auto_minor_version_upgrade"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("availability_zone"))
}

// EngineVersion returns a reference to field engine_version of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("kms_key_arn"))
}

// MultiAz returns a reference to field multi_az of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) MultiAz() terra.BoolValue {
	return terra.ReferenceAsBool(dri.ref.Append("multi_az"))
}

// PreferredMaintenanceWindow returns a reference to field preferred_maintenance_window of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("preferred_maintenance_window"))
}

// PubliclyAccessible returns a reference to field publicly_accessible of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) PubliclyAccessible() terra.BoolValue {
	return terra.ReferenceAsBool(dri.ref.Append("publicly_accessible"))
}

// ReplicationInstanceArn returns a reference to field replication_instance_arn of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) ReplicationInstanceArn() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("replication_instance_arn"))
}

// ReplicationInstanceClass returns a reference to field replication_instance_class of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) ReplicationInstanceClass() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("replication_instance_class"))
}

// ReplicationInstanceId returns a reference to field replication_instance_id of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) ReplicationInstanceId() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("replication_instance_id"))
}

// ReplicationInstancePrivateIps returns a reference to field replication_instance_private_ips of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) ReplicationInstancePrivateIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dri.ref.Append("replication_instance_private_ips"))
}

// ReplicationInstancePublicIps returns a reference to field replication_instance_public_ips of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) ReplicationInstancePublicIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dri.ref.Append("replication_instance_public_ips"))
}

// ReplicationSubnetGroupId returns a reference to field replication_subnet_group_id of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) ReplicationSubnetGroupId() terra.StringValue {
	return terra.ReferenceAsString(dri.ref.Append("replication_subnet_group_id"))
}

// Tags returns a reference to field tags of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dri.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dri.ref.Append("tags_all"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_dms_replication_instance.
func (dri dmsReplicationInstanceAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dri.ref.Append("vpc_security_group_ids"))
}

func (dri dmsReplicationInstanceAttributes) Timeouts() dmsreplicationinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dmsreplicationinstance.TimeoutsAttributes](dri.ref.Append("timeouts"))
}

type dmsReplicationInstanceState struct {
	AllocatedStorage              float64                               `json:"allocated_storage"`
	AllowMajorVersionUpgrade      bool                                  `json:"allow_major_version_upgrade"`
	ApplyImmediately              bool                                  `json:"apply_immediately"`
	AutoMinorVersionUpgrade       bool                                  `json:"auto_minor_version_upgrade"`
	AvailabilityZone              string                                `json:"availability_zone"`
	EngineVersion                 string                                `json:"engine_version"`
	Id                            string                                `json:"id"`
	KmsKeyArn                     string                                `json:"kms_key_arn"`
	MultiAz                       bool                                  `json:"multi_az"`
	PreferredMaintenanceWindow    string                                `json:"preferred_maintenance_window"`
	PubliclyAccessible            bool                                  `json:"publicly_accessible"`
	ReplicationInstanceArn        string                                `json:"replication_instance_arn"`
	ReplicationInstanceClass      string                                `json:"replication_instance_class"`
	ReplicationInstanceId         string                                `json:"replication_instance_id"`
	ReplicationInstancePrivateIps []string                              `json:"replication_instance_private_ips"`
	ReplicationInstancePublicIps  []string                              `json:"replication_instance_public_ips"`
	ReplicationSubnetGroupId      string                                `json:"replication_subnet_group_id"`
	Tags                          map[string]string                     `json:"tags"`
	TagsAll                       map[string]string                     `json:"tags_all"`
	VpcSecurityGroupIds           []string                              `json:"vpc_security_group_ids"`
	Timeouts                      *dmsreplicationinstance.TimeoutsState `json:"timeouts"`
}
