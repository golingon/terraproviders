// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	dataformrepositoryworkflowconfig "github.com/golingon/terraproviders/googlebeta/5.24.0/dataformrepositoryworkflowconfig"
	"io"
)

// NewDataformRepositoryWorkflowConfig creates a new instance of [DataformRepositoryWorkflowConfig].
func NewDataformRepositoryWorkflowConfig(name string, args DataformRepositoryWorkflowConfigArgs) *DataformRepositoryWorkflowConfig {
	return &DataformRepositoryWorkflowConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataformRepositoryWorkflowConfig)(nil)

// DataformRepositoryWorkflowConfig represents the Terraform resource google_dataform_repository_workflow_config.
type DataformRepositoryWorkflowConfig struct {
	Name      string
	Args      DataformRepositoryWorkflowConfigArgs
	state     *dataformRepositoryWorkflowConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataformRepositoryWorkflowConfig].
func (drwc *DataformRepositoryWorkflowConfig) Type() string {
	return "google_dataform_repository_workflow_config"
}

// LocalName returns the local name for [DataformRepositoryWorkflowConfig].
func (drwc *DataformRepositoryWorkflowConfig) LocalName() string {
	return drwc.Name
}

// Configuration returns the configuration (args) for [DataformRepositoryWorkflowConfig].
func (drwc *DataformRepositoryWorkflowConfig) Configuration() interface{} {
	return drwc.Args
}

// DependOn is used for other resources to depend on [DataformRepositoryWorkflowConfig].
func (drwc *DataformRepositoryWorkflowConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(drwc)
}

// Dependencies returns the list of resources [DataformRepositoryWorkflowConfig] depends_on.
func (drwc *DataformRepositoryWorkflowConfig) Dependencies() terra.Dependencies {
	return drwc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataformRepositoryWorkflowConfig].
func (drwc *DataformRepositoryWorkflowConfig) LifecycleManagement() *terra.Lifecycle {
	return drwc.Lifecycle
}

// Attributes returns the attributes for [DataformRepositoryWorkflowConfig].
func (drwc *DataformRepositoryWorkflowConfig) Attributes() dataformRepositoryWorkflowConfigAttributes {
	return dataformRepositoryWorkflowConfigAttributes{ref: terra.ReferenceResource(drwc)}
}

// ImportState imports the given attribute values into [DataformRepositoryWorkflowConfig]'s state.
func (drwc *DataformRepositoryWorkflowConfig) ImportState(av io.Reader) error {
	drwc.state = &dataformRepositoryWorkflowConfigState{}
	if err := json.NewDecoder(av).Decode(drwc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", drwc.Type(), drwc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataformRepositoryWorkflowConfig] has state.
func (drwc *DataformRepositoryWorkflowConfig) State() (*dataformRepositoryWorkflowConfigState, bool) {
	return drwc.state, drwc.state != nil
}

// StateMust returns the state for [DataformRepositoryWorkflowConfig]. Panics if the state is nil.
func (drwc *DataformRepositoryWorkflowConfig) StateMust() *dataformRepositoryWorkflowConfigState {
	if drwc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", drwc.Type(), drwc.LocalName()))
	}
	return drwc.state
}

// DataformRepositoryWorkflowConfigArgs contains the configurations for google_dataform_repository_workflow_config.
type DataformRepositoryWorkflowConfigArgs struct {
	// CronSchedule: string, optional
	CronSchedule terra.StringValue `hcl:"cron_schedule,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ReleaseConfig: string, required
	ReleaseConfig terra.StringValue `hcl:"release_config,attr" validate:"required"`
	// Repository: string, optional
	Repository terra.StringValue `hcl:"repository,attr"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
	// RecentScheduledExecutionRecords: min=0
	RecentScheduledExecutionRecords []dataformrepositoryworkflowconfig.RecentScheduledExecutionRecords `hcl:"recent_scheduled_execution_records,block" validate:"min=0"`
	// InvocationConfig: optional
	InvocationConfig *dataformrepositoryworkflowconfig.InvocationConfig `hcl:"invocation_config,block"`
	// Timeouts: optional
	Timeouts *dataformrepositoryworkflowconfig.Timeouts `hcl:"timeouts,block"`
}
type dataformRepositoryWorkflowConfigAttributes struct {
	ref terra.Reference
}

// CronSchedule returns a reference to field cron_schedule of google_dataform_repository_workflow_config.
func (drwc dataformRepositoryWorkflowConfigAttributes) CronSchedule() terra.StringValue {
	return terra.ReferenceAsString(drwc.ref.Append("cron_schedule"))
}

// Id returns a reference to field id of google_dataform_repository_workflow_config.
func (drwc dataformRepositoryWorkflowConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(drwc.ref.Append("id"))
}

// Name returns a reference to field name of google_dataform_repository_workflow_config.
func (drwc dataformRepositoryWorkflowConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(drwc.ref.Append("name"))
}

// Project returns a reference to field project of google_dataform_repository_workflow_config.
func (drwc dataformRepositoryWorkflowConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(drwc.ref.Append("project"))
}

// Region returns a reference to field region of google_dataform_repository_workflow_config.
func (drwc dataformRepositoryWorkflowConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(drwc.ref.Append("region"))
}

// ReleaseConfig returns a reference to field release_config of google_dataform_repository_workflow_config.
func (drwc dataformRepositoryWorkflowConfigAttributes) ReleaseConfig() terra.StringValue {
	return terra.ReferenceAsString(drwc.ref.Append("release_config"))
}

// Repository returns a reference to field repository of google_dataform_repository_workflow_config.
func (drwc dataformRepositoryWorkflowConfigAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(drwc.ref.Append("repository"))
}

// TimeZone returns a reference to field time_zone of google_dataform_repository_workflow_config.
func (drwc dataformRepositoryWorkflowConfigAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(drwc.ref.Append("time_zone"))
}

func (drwc dataformRepositoryWorkflowConfigAttributes) RecentScheduledExecutionRecords() terra.ListValue[dataformrepositoryworkflowconfig.RecentScheduledExecutionRecordsAttributes] {
	return terra.ReferenceAsList[dataformrepositoryworkflowconfig.RecentScheduledExecutionRecordsAttributes](drwc.ref.Append("recent_scheduled_execution_records"))
}

func (drwc dataformRepositoryWorkflowConfigAttributes) InvocationConfig() terra.ListValue[dataformrepositoryworkflowconfig.InvocationConfigAttributes] {
	return terra.ReferenceAsList[dataformrepositoryworkflowconfig.InvocationConfigAttributes](drwc.ref.Append("invocation_config"))
}

func (drwc dataformRepositoryWorkflowConfigAttributes) Timeouts() dataformrepositoryworkflowconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataformrepositoryworkflowconfig.TimeoutsAttributes](drwc.ref.Append("timeouts"))
}

type dataformRepositoryWorkflowConfigState struct {
	CronSchedule                    string                                                                  `json:"cron_schedule"`
	Id                              string                                                                  `json:"id"`
	Name                            string                                                                  `json:"name"`
	Project                         string                                                                  `json:"project"`
	Region                          string                                                                  `json:"region"`
	ReleaseConfig                   string                                                                  `json:"release_config"`
	Repository                      string                                                                  `json:"repository"`
	TimeZone                        string                                                                  `json:"time_zone"`
	RecentScheduledExecutionRecords []dataformrepositoryworkflowconfig.RecentScheduledExecutionRecordsState `json:"recent_scheduled_execution_records"`
	InvocationConfig                []dataformrepositoryworkflowconfig.InvocationConfigState                `json:"invocation_config"`
	Timeouts                        *dataformrepositoryworkflowconfig.TimeoutsState                         `json:"timeouts"`
}
