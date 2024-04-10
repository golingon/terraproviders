// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewRedshiftserverlessSnapshot creates a new instance of [RedshiftserverlessSnapshot].
func NewRedshiftserverlessSnapshot(name string, args RedshiftserverlessSnapshotArgs) *RedshiftserverlessSnapshot {
	return &RedshiftserverlessSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftserverlessSnapshot)(nil)

// RedshiftserverlessSnapshot represents the Terraform resource aws_redshiftserverless_snapshot.
type RedshiftserverlessSnapshot struct {
	Name      string
	Args      RedshiftserverlessSnapshotArgs
	state     *redshiftserverlessSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftserverlessSnapshot].
func (rs *RedshiftserverlessSnapshot) Type() string {
	return "aws_redshiftserverless_snapshot"
}

// LocalName returns the local name for [RedshiftserverlessSnapshot].
func (rs *RedshiftserverlessSnapshot) LocalName() string {
	return rs.Name
}

// Configuration returns the configuration (args) for [RedshiftserverlessSnapshot].
func (rs *RedshiftserverlessSnapshot) Configuration() interface{} {
	return rs.Args
}

// DependOn is used for other resources to depend on [RedshiftserverlessSnapshot].
func (rs *RedshiftserverlessSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(rs)
}

// Dependencies returns the list of resources [RedshiftserverlessSnapshot] depends_on.
func (rs *RedshiftserverlessSnapshot) Dependencies() terra.Dependencies {
	return rs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftserverlessSnapshot].
func (rs *RedshiftserverlessSnapshot) LifecycleManagement() *terra.Lifecycle {
	return rs.Lifecycle
}

// Attributes returns the attributes for [RedshiftserverlessSnapshot].
func (rs *RedshiftserverlessSnapshot) Attributes() redshiftserverlessSnapshotAttributes {
	return redshiftserverlessSnapshotAttributes{ref: terra.ReferenceResource(rs)}
}

// ImportState imports the given attribute values into [RedshiftserverlessSnapshot]'s state.
func (rs *RedshiftserverlessSnapshot) ImportState(av io.Reader) error {
	rs.state = &redshiftserverlessSnapshotState{}
	if err := json.NewDecoder(av).Decode(rs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rs.Type(), rs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftserverlessSnapshot] has state.
func (rs *RedshiftserverlessSnapshot) State() (*redshiftserverlessSnapshotState, bool) {
	return rs.state, rs.state != nil
}

// StateMust returns the state for [RedshiftserverlessSnapshot]. Panics if the state is nil.
func (rs *RedshiftserverlessSnapshot) StateMust() *redshiftserverlessSnapshotState {
	if rs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rs.Type(), rs.LocalName()))
	}
	return rs.state
}

// RedshiftserverlessSnapshotArgs contains the configurations for aws_redshiftserverless_snapshot.
type RedshiftserverlessSnapshotArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// RetentionPeriod: number, optional
	RetentionPeriod terra.NumberValue `hcl:"retention_period,attr"`
	// SnapshotName: string, required
	SnapshotName terra.StringValue `hcl:"snapshot_name,attr" validate:"required"`
}
type redshiftserverlessSnapshotAttributes struct {
	ref terra.Reference
}

// AccountsWithProvisionedRestoreAccess returns a reference to field accounts_with_provisioned_restore_access of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) AccountsWithProvisionedRestoreAccess() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rs.ref.Append("accounts_with_provisioned_restore_access"))
}

// AccountsWithRestoreAccess returns a reference to field accounts_with_restore_access of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) AccountsWithRestoreAccess() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rs.ref.Append("accounts_with_restore_access"))
}

// AdminUsername returns a reference to field admin_username of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("admin_username"))
}

// Arn returns a reference to field arn of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("arn"))
}

// Id returns a reference to field id of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("kms_key_id"))
}

// NamespaceArn returns a reference to field namespace_arn of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) NamespaceArn() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("namespace_arn"))
}

// NamespaceName returns a reference to field namespace_name of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("namespace_name"))
}

// OwnerAccount returns a reference to field owner_account of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) OwnerAccount() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("owner_account"))
}

// RetentionPeriod returns a reference to field retention_period of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) RetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(rs.ref.Append("retention_period"))
}

// SnapshotName returns a reference to field snapshot_name of aws_redshiftserverless_snapshot.
func (rs redshiftserverlessSnapshotAttributes) SnapshotName() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("snapshot_name"))
}

type redshiftserverlessSnapshotState struct {
	AccountsWithProvisionedRestoreAccess []string `json:"accounts_with_provisioned_restore_access"`
	AccountsWithRestoreAccess            []string `json:"accounts_with_restore_access"`
	AdminUsername                        string   `json:"admin_username"`
	Arn                                  string   `json:"arn"`
	Id                                   string   `json:"id"`
	KmsKeyId                             string   `json:"kms_key_id"`
	NamespaceArn                         string   `json:"namespace_arn"`
	NamespaceName                        string   `json:"namespace_name"`
	OwnerAccount                         string   `json:"owner_account"`
	RetentionPeriod                      float64  `json:"retention_period"`
	SnapshotName                         string   `json:"snapshot_name"`
}
