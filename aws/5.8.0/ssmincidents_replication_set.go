// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssmincidentsreplicationset "github.com/golingon/terraproviders/aws/5.8.0/ssmincidentsreplicationset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsmincidentsReplicationSet creates a new instance of [SsmincidentsReplicationSet].
func NewSsmincidentsReplicationSet(name string, args SsmincidentsReplicationSetArgs) *SsmincidentsReplicationSet {
	return &SsmincidentsReplicationSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsmincidentsReplicationSet)(nil)

// SsmincidentsReplicationSet represents the Terraform resource aws_ssmincidents_replication_set.
type SsmincidentsReplicationSet struct {
	Name      string
	Args      SsmincidentsReplicationSetArgs
	state     *ssmincidentsReplicationSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsmincidentsReplicationSet].
func (srs *SsmincidentsReplicationSet) Type() string {
	return "aws_ssmincidents_replication_set"
}

// LocalName returns the local name for [SsmincidentsReplicationSet].
func (srs *SsmincidentsReplicationSet) LocalName() string {
	return srs.Name
}

// Configuration returns the configuration (args) for [SsmincidentsReplicationSet].
func (srs *SsmincidentsReplicationSet) Configuration() interface{} {
	return srs.Args
}

// DependOn is used for other resources to depend on [SsmincidentsReplicationSet].
func (srs *SsmincidentsReplicationSet) DependOn() terra.Reference {
	return terra.ReferenceResource(srs)
}

// Dependencies returns the list of resources [SsmincidentsReplicationSet] depends_on.
func (srs *SsmincidentsReplicationSet) Dependencies() terra.Dependencies {
	return srs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsmincidentsReplicationSet].
func (srs *SsmincidentsReplicationSet) LifecycleManagement() *terra.Lifecycle {
	return srs.Lifecycle
}

// Attributes returns the attributes for [SsmincidentsReplicationSet].
func (srs *SsmincidentsReplicationSet) Attributes() ssmincidentsReplicationSetAttributes {
	return ssmincidentsReplicationSetAttributes{ref: terra.ReferenceResource(srs)}
}

// ImportState imports the given attribute values into [SsmincidentsReplicationSet]'s state.
func (srs *SsmincidentsReplicationSet) ImportState(av io.Reader) error {
	srs.state = &ssmincidentsReplicationSetState{}
	if err := json.NewDecoder(av).Decode(srs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srs.Type(), srs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsmincidentsReplicationSet] has state.
func (srs *SsmincidentsReplicationSet) State() (*ssmincidentsReplicationSetState, bool) {
	return srs.state, srs.state != nil
}

// StateMust returns the state for [SsmincidentsReplicationSet]. Panics if the state is nil.
func (srs *SsmincidentsReplicationSet) StateMust() *ssmincidentsReplicationSetState {
	if srs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srs.Type(), srs.LocalName()))
	}
	return srs.state
}

// SsmincidentsReplicationSetArgs contains the configurations for aws_ssmincidents_replication_set.
type SsmincidentsReplicationSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Region: min=1
	Region []ssmincidentsreplicationset.Region `hcl:"region,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *ssmincidentsreplicationset.Timeouts `hcl:"timeouts,block"`
}
type ssmincidentsReplicationSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssmincidents_replication_set.
func (srs ssmincidentsReplicationSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(srs.ref.Append("arn"))
}

// CreatedBy returns a reference to field created_by of aws_ssmincidents_replication_set.
func (srs ssmincidentsReplicationSetAttributes) CreatedBy() terra.StringValue {
	return terra.ReferenceAsString(srs.ref.Append("created_by"))
}

// DeletionProtected returns a reference to field deletion_protected of aws_ssmincidents_replication_set.
func (srs ssmincidentsReplicationSetAttributes) DeletionProtected() terra.BoolValue {
	return terra.ReferenceAsBool(srs.ref.Append("deletion_protected"))
}

// Id returns a reference to field id of aws_ssmincidents_replication_set.
func (srs ssmincidentsReplicationSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srs.ref.Append("id"))
}

// LastModifiedBy returns a reference to field last_modified_by of aws_ssmincidents_replication_set.
func (srs ssmincidentsReplicationSetAttributes) LastModifiedBy() terra.StringValue {
	return terra.ReferenceAsString(srs.ref.Append("last_modified_by"))
}

// Status returns a reference to field status of aws_ssmincidents_replication_set.
func (srs ssmincidentsReplicationSetAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(srs.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_ssmincidents_replication_set.
func (srs ssmincidentsReplicationSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](srs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ssmincidents_replication_set.
func (srs ssmincidentsReplicationSetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](srs.ref.Append("tags_all"))
}

func (srs ssmincidentsReplicationSetAttributes) Region() terra.SetValue[ssmincidentsreplicationset.RegionAttributes] {
	return terra.ReferenceAsSet[ssmincidentsreplicationset.RegionAttributes](srs.ref.Append("region"))
}

func (srs ssmincidentsReplicationSetAttributes) Timeouts() ssmincidentsreplicationset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ssmincidentsreplicationset.TimeoutsAttributes](srs.ref.Append("timeouts"))
}

type ssmincidentsReplicationSetState struct {
	Arn               string                                    `json:"arn"`
	CreatedBy         string                                    `json:"created_by"`
	DeletionProtected bool                                      `json:"deletion_protected"`
	Id                string                                    `json:"id"`
	LastModifiedBy    string                                    `json:"last_modified_by"`
	Status            string                                    `json:"status"`
	Tags              map[string]string                         `json:"tags"`
	TagsAll           map[string]string                         `json:"tags_all"`
	Region            []ssmincidentsreplicationset.RegionState  `json:"region"`
	Timeouts          *ssmincidentsreplicationset.TimeoutsState `json:"timeouts"`
}