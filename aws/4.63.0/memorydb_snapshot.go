// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	memorydbsnapshot "github.com/golingon/terraproviders/aws/4.63.0/memorydbsnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMemorydbSnapshot creates a new instance of [MemorydbSnapshot].
func NewMemorydbSnapshot(name string, args MemorydbSnapshotArgs) *MemorydbSnapshot {
	return &MemorydbSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MemorydbSnapshot)(nil)

// MemorydbSnapshot represents the Terraform resource aws_memorydb_snapshot.
type MemorydbSnapshot struct {
	Name      string
	Args      MemorydbSnapshotArgs
	state     *memorydbSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MemorydbSnapshot].
func (ms *MemorydbSnapshot) Type() string {
	return "aws_memorydb_snapshot"
}

// LocalName returns the local name for [MemorydbSnapshot].
func (ms *MemorydbSnapshot) LocalName() string {
	return ms.Name
}

// Configuration returns the configuration (args) for [MemorydbSnapshot].
func (ms *MemorydbSnapshot) Configuration() interface{} {
	return ms.Args
}

// DependOn is used for other resources to depend on [MemorydbSnapshot].
func (ms *MemorydbSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(ms)
}

// Dependencies returns the list of resources [MemorydbSnapshot] depends_on.
func (ms *MemorydbSnapshot) Dependencies() terra.Dependencies {
	return ms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MemorydbSnapshot].
func (ms *MemorydbSnapshot) LifecycleManagement() *terra.Lifecycle {
	return ms.Lifecycle
}

// Attributes returns the attributes for [MemorydbSnapshot].
func (ms *MemorydbSnapshot) Attributes() memorydbSnapshotAttributes {
	return memorydbSnapshotAttributes{ref: terra.ReferenceResource(ms)}
}

// ImportState imports the given attribute values into [MemorydbSnapshot]'s state.
func (ms *MemorydbSnapshot) ImportState(av io.Reader) error {
	ms.state = &memorydbSnapshotState{}
	if err := json.NewDecoder(av).Decode(ms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ms.Type(), ms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MemorydbSnapshot] has state.
func (ms *MemorydbSnapshot) State() (*memorydbSnapshotState, bool) {
	return ms.state, ms.state != nil
}

// StateMust returns the state for [MemorydbSnapshot]. Panics if the state is nil.
func (ms *MemorydbSnapshot) StateMust() *memorydbSnapshotState {
	if ms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ms.Type(), ms.LocalName()))
	}
	return ms.state
}

// MemorydbSnapshotArgs contains the configurations for aws_memorydb_snapshot.
type MemorydbSnapshotArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ClusterConfiguration: min=0
	ClusterConfiguration []memorydbsnapshot.ClusterConfiguration `hcl:"cluster_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *memorydbsnapshot.Timeouts `hcl:"timeouts,block"`
}
type memorydbSnapshotAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_memorydb_snapshot.
func (ms memorydbSnapshotAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("arn"))
}

// ClusterName returns a reference to field cluster_name of aws_memorydb_snapshot.
func (ms memorydbSnapshotAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("cluster_name"))
}

// Id returns a reference to field id of aws_memorydb_snapshot.
func (ms memorydbSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_memorydb_snapshot.
func (ms memorydbSnapshotAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("kms_key_arn"))
}

// Name returns a reference to field name of aws_memorydb_snapshot.
func (ms memorydbSnapshotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_memorydb_snapshot.
func (ms memorydbSnapshotAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("name_prefix"))
}

// Source returns a reference to field source of aws_memorydb_snapshot.
func (ms memorydbSnapshotAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("source"))
}

// Tags returns a reference to field tags of aws_memorydb_snapshot.
func (ms memorydbSnapshotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ms.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_memorydb_snapshot.
func (ms memorydbSnapshotAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ms.ref.Append("tags_all"))
}

func (ms memorydbSnapshotAttributes) ClusterConfiguration() terra.ListValue[memorydbsnapshot.ClusterConfigurationAttributes] {
	return terra.ReferenceAsList[memorydbsnapshot.ClusterConfigurationAttributes](ms.ref.Append("cluster_configuration"))
}

func (ms memorydbSnapshotAttributes) Timeouts() memorydbsnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[memorydbsnapshot.TimeoutsAttributes](ms.ref.Append("timeouts"))
}

type memorydbSnapshotState struct {
	Arn                  string                                       `json:"arn"`
	ClusterName          string                                       `json:"cluster_name"`
	Id                   string                                       `json:"id"`
	KmsKeyArn            string                                       `json:"kms_key_arn"`
	Name                 string                                       `json:"name"`
	NamePrefix           string                                       `json:"name_prefix"`
	Source               string                                       `json:"source"`
	Tags                 map[string]string                            `json:"tags"`
	TagsAll              map[string]string                            `json:"tags_all"`
	ClusterConfiguration []memorydbsnapshot.ClusterConfigurationState `json:"cluster_configuration"`
	Timeouts             *memorydbsnapshot.TimeoutsState              `json:"timeouts"`
}
