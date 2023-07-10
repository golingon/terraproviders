// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dbinstanceautomatedbackupsreplication "github.com/golingon/terraproviders/aws/5.7.0/dbinstanceautomatedbackupsreplication"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDbInstanceAutomatedBackupsReplication creates a new instance of [DbInstanceAutomatedBackupsReplication].
func NewDbInstanceAutomatedBackupsReplication(name string, args DbInstanceAutomatedBackupsReplicationArgs) *DbInstanceAutomatedBackupsReplication {
	return &DbInstanceAutomatedBackupsReplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DbInstanceAutomatedBackupsReplication)(nil)

// DbInstanceAutomatedBackupsReplication represents the Terraform resource aws_db_instance_automated_backups_replication.
type DbInstanceAutomatedBackupsReplication struct {
	Name      string
	Args      DbInstanceAutomatedBackupsReplicationArgs
	state     *dbInstanceAutomatedBackupsReplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DbInstanceAutomatedBackupsReplication].
func (diabr *DbInstanceAutomatedBackupsReplication) Type() string {
	return "aws_db_instance_automated_backups_replication"
}

// LocalName returns the local name for [DbInstanceAutomatedBackupsReplication].
func (diabr *DbInstanceAutomatedBackupsReplication) LocalName() string {
	return diabr.Name
}

// Configuration returns the configuration (args) for [DbInstanceAutomatedBackupsReplication].
func (diabr *DbInstanceAutomatedBackupsReplication) Configuration() interface{} {
	return diabr.Args
}

// DependOn is used for other resources to depend on [DbInstanceAutomatedBackupsReplication].
func (diabr *DbInstanceAutomatedBackupsReplication) DependOn() terra.Reference {
	return terra.ReferenceResource(diabr)
}

// Dependencies returns the list of resources [DbInstanceAutomatedBackupsReplication] depends_on.
func (diabr *DbInstanceAutomatedBackupsReplication) Dependencies() terra.Dependencies {
	return diabr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DbInstanceAutomatedBackupsReplication].
func (diabr *DbInstanceAutomatedBackupsReplication) LifecycleManagement() *terra.Lifecycle {
	return diabr.Lifecycle
}

// Attributes returns the attributes for [DbInstanceAutomatedBackupsReplication].
func (diabr *DbInstanceAutomatedBackupsReplication) Attributes() dbInstanceAutomatedBackupsReplicationAttributes {
	return dbInstanceAutomatedBackupsReplicationAttributes{ref: terra.ReferenceResource(diabr)}
}

// ImportState imports the given attribute values into [DbInstanceAutomatedBackupsReplication]'s state.
func (diabr *DbInstanceAutomatedBackupsReplication) ImportState(av io.Reader) error {
	diabr.state = &dbInstanceAutomatedBackupsReplicationState{}
	if err := json.NewDecoder(av).Decode(diabr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", diabr.Type(), diabr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DbInstanceAutomatedBackupsReplication] has state.
func (diabr *DbInstanceAutomatedBackupsReplication) State() (*dbInstanceAutomatedBackupsReplicationState, bool) {
	return diabr.state, diabr.state != nil
}

// StateMust returns the state for [DbInstanceAutomatedBackupsReplication]. Panics if the state is nil.
func (diabr *DbInstanceAutomatedBackupsReplication) StateMust() *dbInstanceAutomatedBackupsReplicationState {
	if diabr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", diabr.Type(), diabr.LocalName()))
	}
	return diabr.state
}

// DbInstanceAutomatedBackupsReplicationArgs contains the configurations for aws_db_instance_automated_backups_replication.
type DbInstanceAutomatedBackupsReplicationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// PreSignedUrl: string, optional
	PreSignedUrl terra.StringValue `hcl:"pre_signed_url,attr"`
	// RetentionPeriod: number, optional
	RetentionPeriod terra.NumberValue `hcl:"retention_period,attr"`
	// SourceDbInstanceArn: string, required
	SourceDbInstanceArn terra.StringValue `hcl:"source_db_instance_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dbinstanceautomatedbackupsreplication.Timeouts `hcl:"timeouts,block"`
}
type dbInstanceAutomatedBackupsReplicationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_db_instance_automated_backups_replication.
func (diabr dbInstanceAutomatedBackupsReplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(diabr.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_db_instance_automated_backups_replication.
func (diabr dbInstanceAutomatedBackupsReplicationAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(diabr.ref.Append("kms_key_id"))
}

// PreSignedUrl returns a reference to field pre_signed_url of aws_db_instance_automated_backups_replication.
func (diabr dbInstanceAutomatedBackupsReplicationAttributes) PreSignedUrl() terra.StringValue {
	return terra.ReferenceAsString(diabr.ref.Append("pre_signed_url"))
}

// RetentionPeriod returns a reference to field retention_period of aws_db_instance_automated_backups_replication.
func (diabr dbInstanceAutomatedBackupsReplicationAttributes) RetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(diabr.ref.Append("retention_period"))
}

// SourceDbInstanceArn returns a reference to field source_db_instance_arn of aws_db_instance_automated_backups_replication.
func (diabr dbInstanceAutomatedBackupsReplicationAttributes) SourceDbInstanceArn() terra.StringValue {
	return terra.ReferenceAsString(diabr.ref.Append("source_db_instance_arn"))
}

func (diabr dbInstanceAutomatedBackupsReplicationAttributes) Timeouts() dbinstanceautomatedbackupsreplication.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dbinstanceautomatedbackupsreplication.TimeoutsAttributes](diabr.ref.Append("timeouts"))
}

type dbInstanceAutomatedBackupsReplicationState struct {
	Id                  string                                               `json:"id"`
	KmsKeyId            string                                               `json:"kms_key_id"`
	PreSignedUrl        string                                               `json:"pre_signed_url"`
	RetentionPeriod     float64                                              `json:"retention_period"`
	SourceDbInstanceArn string                                               `json:"source_db_instance_arn"`
	Timeouts            *dbinstanceautomatedbackupsreplication.TimeoutsState `json:"timeouts"`
}
