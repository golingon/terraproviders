// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	batchjob "github.com/golingon/terraproviders/azurerm/3.58.0/batchjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBatchJob creates a new instance of [BatchJob].
func NewBatchJob(name string, args BatchJobArgs) *BatchJob {
	return &BatchJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BatchJob)(nil)

// BatchJob represents the Terraform resource azurerm_batch_job.
type BatchJob struct {
	Name      string
	Args      BatchJobArgs
	state     *batchJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BatchJob].
func (bj *BatchJob) Type() string {
	return "azurerm_batch_job"
}

// LocalName returns the local name for [BatchJob].
func (bj *BatchJob) LocalName() string {
	return bj.Name
}

// Configuration returns the configuration (args) for [BatchJob].
func (bj *BatchJob) Configuration() interface{} {
	return bj.Args
}

// DependOn is used for other resources to depend on [BatchJob].
func (bj *BatchJob) DependOn() terra.Reference {
	return terra.ReferenceResource(bj)
}

// Dependencies returns the list of resources [BatchJob] depends_on.
func (bj *BatchJob) Dependencies() terra.Dependencies {
	return bj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BatchJob].
func (bj *BatchJob) LifecycleManagement() *terra.Lifecycle {
	return bj.Lifecycle
}

// Attributes returns the attributes for [BatchJob].
func (bj *BatchJob) Attributes() batchJobAttributes {
	return batchJobAttributes{ref: terra.ReferenceResource(bj)}
}

// ImportState imports the given attribute values into [BatchJob]'s state.
func (bj *BatchJob) ImportState(av io.Reader) error {
	bj.state = &batchJobState{}
	if err := json.NewDecoder(av).Decode(bj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bj.Type(), bj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BatchJob] has state.
func (bj *BatchJob) State() (*batchJobState, bool) {
	return bj.state, bj.state != nil
}

// StateMust returns the state for [BatchJob]. Panics if the state is nil.
func (bj *BatchJob) StateMust() *batchJobState {
	if bj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bj.Type(), bj.LocalName()))
	}
	return bj.state
}

// BatchJobArgs contains the configurations for azurerm_batch_job.
type BatchJobArgs struct {
	// BatchPoolId: string, required
	BatchPoolId terra.StringValue `hcl:"batch_pool_id,attr" validate:"required"`
	// CommonEnvironmentProperties: map of string, optional
	CommonEnvironmentProperties terra.MapValue[terra.StringValue] `hcl:"common_environment_properties,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// TaskRetryMaximum: number, optional
	TaskRetryMaximum terra.NumberValue `hcl:"task_retry_maximum,attr"`
	// Timeouts: optional
	Timeouts *batchjob.Timeouts `hcl:"timeouts,block"`
}
type batchJobAttributes struct {
	ref terra.Reference
}

// BatchPoolId returns a reference to field batch_pool_id of azurerm_batch_job.
func (bj batchJobAttributes) BatchPoolId() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("batch_pool_id"))
}

// CommonEnvironmentProperties returns a reference to field common_environment_properties of azurerm_batch_job.
func (bj batchJobAttributes) CommonEnvironmentProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bj.ref.Append("common_environment_properties"))
}

// DisplayName returns a reference to field display_name of azurerm_batch_job.
func (bj batchJobAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_batch_job.
func (bj batchJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_batch_job.
func (bj batchJobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_batch_job.
func (bj batchJobAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(bj.ref.Append("priority"))
}

// TaskRetryMaximum returns a reference to field task_retry_maximum of azurerm_batch_job.
func (bj batchJobAttributes) TaskRetryMaximum() terra.NumberValue {
	return terra.ReferenceAsNumber(bj.ref.Append("task_retry_maximum"))
}

func (bj batchJobAttributes) Timeouts() batchjob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[batchjob.TimeoutsAttributes](bj.ref.Append("timeouts"))
}

type batchJobState struct {
	BatchPoolId                 string                  `json:"batch_pool_id"`
	CommonEnvironmentProperties map[string]string       `json:"common_environment_properties"`
	DisplayName                 string                  `json:"display_name"`
	Id                          string                  `json:"id"`
	Name                        string                  `json:"name"`
	Priority                    float64                 `json:"priority"`
	TaskRetryMaximum            float64                 `json:"task_retry_maximum"`
	Timeouts                    *batchjob.TimeoutsState `json:"timeouts"`
}
