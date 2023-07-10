// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kinesisanalyticsv2applicationsnapshot "github.com/golingon/terraproviders/aws/5.7.0/kinesisanalyticsv2applicationsnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKinesisanalyticsv2ApplicationSnapshot creates a new instance of [Kinesisanalyticsv2ApplicationSnapshot].
func NewKinesisanalyticsv2ApplicationSnapshot(name string, args Kinesisanalyticsv2ApplicationSnapshotArgs) *Kinesisanalyticsv2ApplicationSnapshot {
	return &Kinesisanalyticsv2ApplicationSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Kinesisanalyticsv2ApplicationSnapshot)(nil)

// Kinesisanalyticsv2ApplicationSnapshot represents the Terraform resource aws_kinesisanalyticsv2_application_snapshot.
type Kinesisanalyticsv2ApplicationSnapshot struct {
	Name      string
	Args      Kinesisanalyticsv2ApplicationSnapshotArgs
	state     *kinesisanalyticsv2ApplicationSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Kinesisanalyticsv2ApplicationSnapshot].
func (kas *Kinesisanalyticsv2ApplicationSnapshot) Type() string {
	return "aws_kinesisanalyticsv2_application_snapshot"
}

// LocalName returns the local name for [Kinesisanalyticsv2ApplicationSnapshot].
func (kas *Kinesisanalyticsv2ApplicationSnapshot) LocalName() string {
	return kas.Name
}

// Configuration returns the configuration (args) for [Kinesisanalyticsv2ApplicationSnapshot].
func (kas *Kinesisanalyticsv2ApplicationSnapshot) Configuration() interface{} {
	return kas.Args
}

// DependOn is used for other resources to depend on [Kinesisanalyticsv2ApplicationSnapshot].
func (kas *Kinesisanalyticsv2ApplicationSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(kas)
}

// Dependencies returns the list of resources [Kinesisanalyticsv2ApplicationSnapshot] depends_on.
func (kas *Kinesisanalyticsv2ApplicationSnapshot) Dependencies() terra.Dependencies {
	return kas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Kinesisanalyticsv2ApplicationSnapshot].
func (kas *Kinesisanalyticsv2ApplicationSnapshot) LifecycleManagement() *terra.Lifecycle {
	return kas.Lifecycle
}

// Attributes returns the attributes for [Kinesisanalyticsv2ApplicationSnapshot].
func (kas *Kinesisanalyticsv2ApplicationSnapshot) Attributes() kinesisanalyticsv2ApplicationSnapshotAttributes {
	return kinesisanalyticsv2ApplicationSnapshotAttributes{ref: terra.ReferenceResource(kas)}
}

// ImportState imports the given attribute values into [Kinesisanalyticsv2ApplicationSnapshot]'s state.
func (kas *Kinesisanalyticsv2ApplicationSnapshot) ImportState(av io.Reader) error {
	kas.state = &kinesisanalyticsv2ApplicationSnapshotState{}
	if err := json.NewDecoder(av).Decode(kas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kas.Type(), kas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Kinesisanalyticsv2ApplicationSnapshot] has state.
func (kas *Kinesisanalyticsv2ApplicationSnapshot) State() (*kinesisanalyticsv2ApplicationSnapshotState, bool) {
	return kas.state, kas.state != nil
}

// StateMust returns the state for [Kinesisanalyticsv2ApplicationSnapshot]. Panics if the state is nil.
func (kas *Kinesisanalyticsv2ApplicationSnapshot) StateMust() *kinesisanalyticsv2ApplicationSnapshotState {
	if kas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kas.Type(), kas.LocalName()))
	}
	return kas.state
}

// Kinesisanalyticsv2ApplicationSnapshotArgs contains the configurations for aws_kinesisanalyticsv2_application_snapshot.
type Kinesisanalyticsv2ApplicationSnapshotArgs struct {
	// ApplicationName: string, required
	ApplicationName terra.StringValue `hcl:"application_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SnapshotName: string, required
	SnapshotName terra.StringValue `hcl:"snapshot_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *kinesisanalyticsv2applicationsnapshot.Timeouts `hcl:"timeouts,block"`
}
type kinesisanalyticsv2ApplicationSnapshotAttributes struct {
	ref terra.Reference
}

// ApplicationName returns a reference to field application_name of aws_kinesisanalyticsv2_application_snapshot.
func (kas kinesisanalyticsv2ApplicationSnapshotAttributes) ApplicationName() terra.StringValue {
	return terra.ReferenceAsString(kas.ref.Append("application_name"))
}

// ApplicationVersionId returns a reference to field application_version_id of aws_kinesisanalyticsv2_application_snapshot.
func (kas kinesisanalyticsv2ApplicationSnapshotAttributes) ApplicationVersionId() terra.NumberValue {
	return terra.ReferenceAsNumber(kas.ref.Append("application_version_id"))
}

// Id returns a reference to field id of aws_kinesisanalyticsv2_application_snapshot.
func (kas kinesisanalyticsv2ApplicationSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kas.ref.Append("id"))
}

// SnapshotCreationTimestamp returns a reference to field snapshot_creation_timestamp of aws_kinesisanalyticsv2_application_snapshot.
func (kas kinesisanalyticsv2ApplicationSnapshotAttributes) SnapshotCreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(kas.ref.Append("snapshot_creation_timestamp"))
}

// SnapshotName returns a reference to field snapshot_name of aws_kinesisanalyticsv2_application_snapshot.
func (kas kinesisanalyticsv2ApplicationSnapshotAttributes) SnapshotName() terra.StringValue {
	return terra.ReferenceAsString(kas.ref.Append("snapshot_name"))
}

func (kas kinesisanalyticsv2ApplicationSnapshotAttributes) Timeouts() kinesisanalyticsv2applicationsnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kinesisanalyticsv2applicationsnapshot.TimeoutsAttributes](kas.ref.Append("timeouts"))
}

type kinesisanalyticsv2ApplicationSnapshotState struct {
	ApplicationName           string                                               `json:"application_name"`
	ApplicationVersionId      float64                                              `json:"application_version_id"`
	Id                        string                                               `json:"id"`
	SnapshotCreationTimestamp string                                               `json:"snapshot_creation_timestamp"`
	SnapshotName              string                                               `json:"snapshot_name"`
	Timeouts                  *kinesisanalyticsv2applicationsnapshot.TimeoutsState `json:"timeouts"`
}
