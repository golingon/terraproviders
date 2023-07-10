// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssmresourcedatasync "github.com/golingon/terraproviders/aws/5.7.0/ssmresourcedatasync"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsmResourceDataSync creates a new instance of [SsmResourceDataSync].
func NewSsmResourceDataSync(name string, args SsmResourceDataSyncArgs) *SsmResourceDataSync {
	return &SsmResourceDataSync{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsmResourceDataSync)(nil)

// SsmResourceDataSync represents the Terraform resource aws_ssm_resource_data_sync.
type SsmResourceDataSync struct {
	Name      string
	Args      SsmResourceDataSyncArgs
	state     *ssmResourceDataSyncState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsmResourceDataSync].
func (srds *SsmResourceDataSync) Type() string {
	return "aws_ssm_resource_data_sync"
}

// LocalName returns the local name for [SsmResourceDataSync].
func (srds *SsmResourceDataSync) LocalName() string {
	return srds.Name
}

// Configuration returns the configuration (args) for [SsmResourceDataSync].
func (srds *SsmResourceDataSync) Configuration() interface{} {
	return srds.Args
}

// DependOn is used for other resources to depend on [SsmResourceDataSync].
func (srds *SsmResourceDataSync) DependOn() terra.Reference {
	return terra.ReferenceResource(srds)
}

// Dependencies returns the list of resources [SsmResourceDataSync] depends_on.
func (srds *SsmResourceDataSync) Dependencies() terra.Dependencies {
	return srds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsmResourceDataSync].
func (srds *SsmResourceDataSync) LifecycleManagement() *terra.Lifecycle {
	return srds.Lifecycle
}

// Attributes returns the attributes for [SsmResourceDataSync].
func (srds *SsmResourceDataSync) Attributes() ssmResourceDataSyncAttributes {
	return ssmResourceDataSyncAttributes{ref: terra.ReferenceResource(srds)}
}

// ImportState imports the given attribute values into [SsmResourceDataSync]'s state.
func (srds *SsmResourceDataSync) ImportState(av io.Reader) error {
	srds.state = &ssmResourceDataSyncState{}
	if err := json.NewDecoder(av).Decode(srds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srds.Type(), srds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsmResourceDataSync] has state.
func (srds *SsmResourceDataSync) State() (*ssmResourceDataSyncState, bool) {
	return srds.state, srds.state != nil
}

// StateMust returns the state for [SsmResourceDataSync]. Panics if the state is nil.
func (srds *SsmResourceDataSync) StateMust() *ssmResourceDataSyncState {
	if srds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srds.Type(), srds.LocalName()))
	}
	return srds.state
}

// SsmResourceDataSyncArgs contains the configurations for aws_ssm_resource_data_sync.
type SsmResourceDataSyncArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// S3Destination: required
	S3Destination *ssmresourcedatasync.S3Destination `hcl:"s3_destination,block" validate:"required"`
}
type ssmResourceDataSyncAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ssm_resource_data_sync.
func (srds ssmResourceDataSyncAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srds.ref.Append("id"))
}

// Name returns a reference to field name of aws_ssm_resource_data_sync.
func (srds ssmResourceDataSyncAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srds.ref.Append("name"))
}

func (srds ssmResourceDataSyncAttributes) S3Destination() terra.ListValue[ssmresourcedatasync.S3DestinationAttributes] {
	return terra.ReferenceAsList[ssmresourcedatasync.S3DestinationAttributes](srds.ref.Append("s3_destination"))
}

type ssmResourceDataSyncState struct {
	Id            string                                   `json:"id"`
	Name          string                                   `json:"name"`
	S3Destination []ssmresourcedatasync.S3DestinationState `json:"s3_destination"`
}
