// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	synchronizationjob "github.com/golingon/terraproviders/azuread/2.37.1/synchronizationjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynchronizationJob creates a new instance of [SynchronizationJob].
func NewSynchronizationJob(name string, args SynchronizationJobArgs) *SynchronizationJob {
	return &SynchronizationJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynchronizationJob)(nil)

// SynchronizationJob represents the Terraform resource azuread_synchronization_job.
type SynchronizationJob struct {
	Name      string
	Args      SynchronizationJobArgs
	state     *synchronizationJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynchronizationJob].
func (sj *SynchronizationJob) Type() string {
	return "azuread_synchronization_job"
}

// LocalName returns the local name for [SynchronizationJob].
func (sj *SynchronizationJob) LocalName() string {
	return sj.Name
}

// Configuration returns the configuration (args) for [SynchronizationJob].
func (sj *SynchronizationJob) Configuration() interface{} {
	return sj.Args
}

// DependOn is used for other resources to depend on [SynchronizationJob].
func (sj *SynchronizationJob) DependOn() terra.Reference {
	return terra.ReferenceResource(sj)
}

// Dependencies returns the list of resources [SynchronizationJob] depends_on.
func (sj *SynchronizationJob) Dependencies() terra.Dependencies {
	return sj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynchronizationJob].
func (sj *SynchronizationJob) LifecycleManagement() *terra.Lifecycle {
	return sj.Lifecycle
}

// Attributes returns the attributes for [SynchronizationJob].
func (sj *SynchronizationJob) Attributes() synchronizationJobAttributes {
	return synchronizationJobAttributes{ref: terra.ReferenceResource(sj)}
}

// ImportState imports the given attribute values into [SynchronizationJob]'s state.
func (sj *SynchronizationJob) ImportState(av io.Reader) error {
	sj.state = &synchronizationJobState{}
	if err := json.NewDecoder(av).Decode(sj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sj.Type(), sj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynchronizationJob] has state.
func (sj *SynchronizationJob) State() (*synchronizationJobState, bool) {
	return sj.state, sj.state != nil
}

// StateMust returns the state for [SynchronizationJob]. Panics if the state is nil.
func (sj *SynchronizationJob) StateMust() *synchronizationJobState {
	if sj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sj.Type(), sj.LocalName()))
	}
	return sj.state
}

// SynchronizationJobArgs contains the configurations for azuread_synchronization_job.
type SynchronizationJobArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServicePrincipalId: string, required
	ServicePrincipalId terra.StringValue `hcl:"service_principal_id,attr" validate:"required"`
	// TemplateId: string, required
	TemplateId terra.StringValue `hcl:"template_id,attr" validate:"required"`
	// Schedule: min=0
	Schedule []synchronizationjob.Schedule `hcl:"schedule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *synchronizationjob.Timeouts `hcl:"timeouts,block"`
}
type synchronizationJobAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azuread_synchronization_job.
func (sj synchronizationJobAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sj.ref.Append("enabled"))
}

// Id returns a reference to field id of azuread_synchronization_job.
func (sj synchronizationJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sj.ref.Append("id"))
}

// ServicePrincipalId returns a reference to field service_principal_id of azuread_synchronization_job.
func (sj synchronizationJobAttributes) ServicePrincipalId() terra.StringValue {
	return terra.ReferenceAsString(sj.ref.Append("service_principal_id"))
}

// TemplateId returns a reference to field template_id of azuread_synchronization_job.
func (sj synchronizationJobAttributes) TemplateId() terra.StringValue {
	return terra.ReferenceAsString(sj.ref.Append("template_id"))
}

func (sj synchronizationJobAttributes) Schedule() terra.ListValue[synchronizationjob.ScheduleAttributes] {
	return terra.ReferenceAsList[synchronizationjob.ScheduleAttributes](sj.ref.Append("schedule"))
}

func (sj synchronizationJobAttributes) Timeouts() synchronizationjob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synchronizationjob.TimeoutsAttributes](sj.ref.Append("timeouts"))
}

type synchronizationJobState struct {
	Enabled            bool                               `json:"enabled"`
	Id                 string                             `json:"id"`
	ServicePrincipalId string                             `json:"service_principal_id"`
	TemplateId         string                             `json:"template_id"`
	Schedule           []synchronizationjob.ScheduleState `json:"schedule"`
	Timeouts           *synchronizationjob.TimeoutsState  `json:"timeouts"`
}
