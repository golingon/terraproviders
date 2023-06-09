// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedshiftSnapshotCopyGrant creates a new instance of [RedshiftSnapshotCopyGrant].
func NewRedshiftSnapshotCopyGrant(name string, args RedshiftSnapshotCopyGrantArgs) *RedshiftSnapshotCopyGrant {
	return &RedshiftSnapshotCopyGrant{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftSnapshotCopyGrant)(nil)

// RedshiftSnapshotCopyGrant represents the Terraform resource aws_redshift_snapshot_copy_grant.
type RedshiftSnapshotCopyGrant struct {
	Name      string
	Args      RedshiftSnapshotCopyGrantArgs
	state     *redshiftSnapshotCopyGrantState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftSnapshotCopyGrant].
func (rscg *RedshiftSnapshotCopyGrant) Type() string {
	return "aws_redshift_snapshot_copy_grant"
}

// LocalName returns the local name for [RedshiftSnapshotCopyGrant].
func (rscg *RedshiftSnapshotCopyGrant) LocalName() string {
	return rscg.Name
}

// Configuration returns the configuration (args) for [RedshiftSnapshotCopyGrant].
func (rscg *RedshiftSnapshotCopyGrant) Configuration() interface{} {
	return rscg.Args
}

// DependOn is used for other resources to depend on [RedshiftSnapshotCopyGrant].
func (rscg *RedshiftSnapshotCopyGrant) DependOn() terra.Reference {
	return terra.ReferenceResource(rscg)
}

// Dependencies returns the list of resources [RedshiftSnapshotCopyGrant] depends_on.
func (rscg *RedshiftSnapshotCopyGrant) Dependencies() terra.Dependencies {
	return rscg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftSnapshotCopyGrant].
func (rscg *RedshiftSnapshotCopyGrant) LifecycleManagement() *terra.Lifecycle {
	return rscg.Lifecycle
}

// Attributes returns the attributes for [RedshiftSnapshotCopyGrant].
func (rscg *RedshiftSnapshotCopyGrant) Attributes() redshiftSnapshotCopyGrantAttributes {
	return redshiftSnapshotCopyGrantAttributes{ref: terra.ReferenceResource(rscg)}
}

// ImportState imports the given attribute values into [RedshiftSnapshotCopyGrant]'s state.
func (rscg *RedshiftSnapshotCopyGrant) ImportState(av io.Reader) error {
	rscg.state = &redshiftSnapshotCopyGrantState{}
	if err := json.NewDecoder(av).Decode(rscg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rscg.Type(), rscg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftSnapshotCopyGrant] has state.
func (rscg *RedshiftSnapshotCopyGrant) State() (*redshiftSnapshotCopyGrantState, bool) {
	return rscg.state, rscg.state != nil
}

// StateMust returns the state for [RedshiftSnapshotCopyGrant]. Panics if the state is nil.
func (rscg *RedshiftSnapshotCopyGrant) StateMust() *redshiftSnapshotCopyGrantState {
	if rscg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rscg.Type(), rscg.LocalName()))
	}
	return rscg.state
}

// RedshiftSnapshotCopyGrantArgs contains the configurations for aws_redshift_snapshot_copy_grant.
type RedshiftSnapshotCopyGrantArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// SnapshotCopyGrantName: string, required
	SnapshotCopyGrantName terra.StringValue `hcl:"snapshot_copy_grant_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type redshiftSnapshotCopyGrantAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_redshift_snapshot_copy_grant.
func (rscg redshiftSnapshotCopyGrantAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rscg.ref.Append("arn"))
}

// Id returns a reference to field id of aws_redshift_snapshot_copy_grant.
func (rscg redshiftSnapshotCopyGrantAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rscg.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_redshift_snapshot_copy_grant.
func (rscg redshiftSnapshotCopyGrantAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(rscg.ref.Append("kms_key_id"))
}

// SnapshotCopyGrantName returns a reference to field snapshot_copy_grant_name of aws_redshift_snapshot_copy_grant.
func (rscg redshiftSnapshotCopyGrantAttributes) SnapshotCopyGrantName() terra.StringValue {
	return terra.ReferenceAsString(rscg.ref.Append("snapshot_copy_grant_name"))
}

// Tags returns a reference to field tags of aws_redshift_snapshot_copy_grant.
func (rscg redshiftSnapshotCopyGrantAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rscg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_redshift_snapshot_copy_grant.
func (rscg redshiftSnapshotCopyGrantAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rscg.ref.Append("tags_all"))
}

type redshiftSnapshotCopyGrantState struct {
	Arn                   string            `json:"arn"`
	Id                    string            `json:"id"`
	KmsKeyId              string            `json:"kms_key_id"`
	SnapshotCopyGrantName string            `json:"snapshot_copy_grant_name"`
	Tags                  map[string]string `json:"tags"`
	TagsAll               map[string]string `json:"tags_all"`
}
