// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_bigquery_data_transfer_config

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_bigquery_data_transfer_config.
type Resource struct {
	Name      string
	Args      Args
	state     *googleBigqueryDataTransferConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gbdtc *Resource) Type() string {
	return "google_bigquery_data_transfer_config"
}

// LocalName returns the local name for [Resource].
func (gbdtc *Resource) LocalName() string {
	return gbdtc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gbdtc *Resource) Configuration() interface{} {
	return gbdtc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gbdtc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gbdtc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gbdtc *Resource) Dependencies() terra.Dependencies {
	return gbdtc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gbdtc *Resource) LifecycleManagement() *terra.Lifecycle {
	return gbdtc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gbdtc *Resource) Attributes() googleBigqueryDataTransferConfigAttributes {
	return googleBigqueryDataTransferConfigAttributes{ref: terra.ReferenceResource(gbdtc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gbdtc *Resource) ImportState(state io.Reader) error {
	gbdtc.state = &googleBigqueryDataTransferConfigState{}
	if err := json.NewDecoder(state).Decode(gbdtc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gbdtc.Type(), gbdtc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gbdtc *Resource) State() (*googleBigqueryDataTransferConfigState, bool) {
	return gbdtc.state, gbdtc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gbdtc *Resource) StateMust() *googleBigqueryDataTransferConfigState {
	if gbdtc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gbdtc.Type(), gbdtc.LocalName()))
	}
	return gbdtc.state
}

// Args contains the configurations for google_bigquery_data_transfer_config.
type Args struct {
	// DataRefreshWindowDays: number, optional
	DataRefreshWindowDays terra.NumberValue `hcl:"data_refresh_window_days,attr"`
	// DataSourceId: string, required
	DataSourceId terra.StringValue `hcl:"data_source_id,attr" validate:"required"`
	// DestinationDatasetId: string, optional
	DestinationDatasetId terra.StringValue `hcl:"destination_dataset_id,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// NotificationPubsubTopic: string, optional
	NotificationPubsubTopic terra.StringValue `hcl:"notification_pubsub_topic,attr"`
	// Params: map of string, required
	Params terra.MapValue[terra.StringValue] `hcl:"params,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Schedule: string, optional
	Schedule terra.StringValue `hcl:"schedule,attr"`
	// ServiceAccountName: string, optional
	ServiceAccountName terra.StringValue `hcl:"service_account_name,attr"`
	// EmailPreferences: optional
	EmailPreferences *EmailPreferences `hcl:"email_preferences,block"`
	// ScheduleOptions: optional
	ScheduleOptions *ScheduleOptions `hcl:"schedule_options,block"`
	// SensitiveParams: optional
	SensitiveParams *SensitiveParams `hcl:"sensitive_params,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleBigqueryDataTransferConfigAttributes struct {
	ref terra.Reference
}

// DataRefreshWindowDays returns a reference to field data_refresh_window_days of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) DataRefreshWindowDays() terra.NumberValue {
	return terra.ReferenceAsNumber(gbdtc.ref.Append("data_refresh_window_days"))
}

// DataSourceId returns a reference to field data_source_id of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) DataSourceId() terra.StringValue {
	return terra.ReferenceAsString(gbdtc.ref.Append("data_source_id"))
}

// DestinationDatasetId returns a reference to field destination_dataset_id of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) DestinationDatasetId() terra.StringValue {
	return terra.ReferenceAsString(gbdtc.ref.Append("destination_dataset_id"))
}

// Disabled returns a reference to field disabled of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(gbdtc.ref.Append("disabled"))
}

// DisplayName returns a reference to field display_name of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gbdtc.ref.Append("display_name"))
}

// Id returns a reference to field id of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbdtc.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gbdtc.ref.Append("location"))
}

// Name returns a reference to field name of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbdtc.ref.Append("name"))
}

// NotificationPubsubTopic returns a reference to field notification_pubsub_topic of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) NotificationPubsubTopic() terra.StringValue {
	return terra.ReferenceAsString(gbdtc.ref.Append("notification_pubsub_topic"))
}

// Params returns a reference to field params of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) Params() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gbdtc.ref.Append("params"))
}

// Project returns a reference to field project of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbdtc.ref.Append("project"))
}

// Schedule returns a reference to field schedule of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(gbdtc.ref.Append("schedule"))
}

// ServiceAccountName returns a reference to field service_account_name of google_bigquery_data_transfer_config.
func (gbdtc googleBigqueryDataTransferConfigAttributes) ServiceAccountName() terra.StringValue {
	return terra.ReferenceAsString(gbdtc.ref.Append("service_account_name"))
}

func (gbdtc googleBigqueryDataTransferConfigAttributes) EmailPreferences() terra.ListValue[EmailPreferencesAttributes] {
	return terra.ReferenceAsList[EmailPreferencesAttributes](gbdtc.ref.Append("email_preferences"))
}

func (gbdtc googleBigqueryDataTransferConfigAttributes) ScheduleOptions() terra.ListValue[ScheduleOptionsAttributes] {
	return terra.ReferenceAsList[ScheduleOptionsAttributes](gbdtc.ref.Append("schedule_options"))
}

func (gbdtc googleBigqueryDataTransferConfigAttributes) SensitiveParams() terra.ListValue[SensitiveParamsAttributes] {
	return terra.ReferenceAsList[SensitiveParamsAttributes](gbdtc.ref.Append("sensitive_params"))
}

func (gbdtc googleBigqueryDataTransferConfigAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gbdtc.ref.Append("timeouts"))
}

type googleBigqueryDataTransferConfigState struct {
	DataRefreshWindowDays   float64                 `json:"data_refresh_window_days"`
	DataSourceId            string                  `json:"data_source_id"`
	DestinationDatasetId    string                  `json:"destination_dataset_id"`
	Disabled                bool                    `json:"disabled"`
	DisplayName             string                  `json:"display_name"`
	Id                      string                  `json:"id"`
	Location                string                  `json:"location"`
	Name                    string                  `json:"name"`
	NotificationPubsubTopic string                  `json:"notification_pubsub_topic"`
	Params                  map[string]string       `json:"params"`
	Project                 string                  `json:"project"`
	Schedule                string                  `json:"schedule"`
	ServiceAccountName      string                  `json:"service_account_name"`
	EmailPreferences        []EmailPreferencesState `json:"email_preferences"`
	ScheduleOptions         []ScheduleOptionsState  `json:"schedule_options"`
	SensitiveParams         []SensitiveParamsState  `json:"sensitive_params"`
	Timeouts                *TimeoutsState          `json:"timeouts"`
}
