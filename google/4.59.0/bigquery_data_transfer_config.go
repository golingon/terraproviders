// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigquerydatatransferconfig "github.com/golingon/terraproviders/google/4.59.0/bigquerydatatransferconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryDataTransferConfig creates a new instance of [BigqueryDataTransferConfig].
func NewBigqueryDataTransferConfig(name string, args BigqueryDataTransferConfigArgs) *BigqueryDataTransferConfig {
	return &BigqueryDataTransferConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryDataTransferConfig)(nil)

// BigqueryDataTransferConfig represents the Terraform resource google_bigquery_data_transfer_config.
type BigqueryDataTransferConfig struct {
	Name      string
	Args      BigqueryDataTransferConfigArgs
	state     *bigqueryDataTransferConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryDataTransferConfig].
func (bdtc *BigqueryDataTransferConfig) Type() string {
	return "google_bigquery_data_transfer_config"
}

// LocalName returns the local name for [BigqueryDataTransferConfig].
func (bdtc *BigqueryDataTransferConfig) LocalName() string {
	return bdtc.Name
}

// Configuration returns the configuration (args) for [BigqueryDataTransferConfig].
func (bdtc *BigqueryDataTransferConfig) Configuration() interface{} {
	return bdtc.Args
}

// DependOn is used for other resources to depend on [BigqueryDataTransferConfig].
func (bdtc *BigqueryDataTransferConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(bdtc)
}

// Dependencies returns the list of resources [BigqueryDataTransferConfig] depends_on.
func (bdtc *BigqueryDataTransferConfig) Dependencies() terra.Dependencies {
	return bdtc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryDataTransferConfig].
func (bdtc *BigqueryDataTransferConfig) LifecycleManagement() *terra.Lifecycle {
	return bdtc.Lifecycle
}

// Attributes returns the attributes for [BigqueryDataTransferConfig].
func (bdtc *BigqueryDataTransferConfig) Attributes() bigqueryDataTransferConfigAttributes {
	return bigqueryDataTransferConfigAttributes{ref: terra.ReferenceResource(bdtc)}
}

// ImportState imports the given attribute values into [BigqueryDataTransferConfig]'s state.
func (bdtc *BigqueryDataTransferConfig) ImportState(av io.Reader) error {
	bdtc.state = &bigqueryDataTransferConfigState{}
	if err := json.NewDecoder(av).Decode(bdtc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bdtc.Type(), bdtc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryDataTransferConfig] has state.
func (bdtc *BigqueryDataTransferConfig) State() (*bigqueryDataTransferConfigState, bool) {
	return bdtc.state, bdtc.state != nil
}

// StateMust returns the state for [BigqueryDataTransferConfig]. Panics if the state is nil.
func (bdtc *BigqueryDataTransferConfig) StateMust() *bigqueryDataTransferConfigState {
	if bdtc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bdtc.Type(), bdtc.LocalName()))
	}
	return bdtc.state
}

// BigqueryDataTransferConfigArgs contains the configurations for google_bigquery_data_transfer_config.
type BigqueryDataTransferConfigArgs struct {
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
	EmailPreferences *bigquerydatatransferconfig.EmailPreferences `hcl:"email_preferences,block"`
	// ScheduleOptions: optional
	ScheduleOptions *bigquerydatatransferconfig.ScheduleOptions `hcl:"schedule_options,block"`
	// SensitiveParams: optional
	SensitiveParams *bigquerydatatransferconfig.SensitiveParams `hcl:"sensitive_params,block"`
	// Timeouts: optional
	Timeouts *bigquerydatatransferconfig.Timeouts `hcl:"timeouts,block"`
}
type bigqueryDataTransferConfigAttributes struct {
	ref terra.Reference
}

// DataRefreshWindowDays returns a reference to field data_refresh_window_days of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) DataRefreshWindowDays() terra.NumberValue {
	return terra.ReferenceAsNumber(bdtc.ref.Append("data_refresh_window_days"))
}

// DataSourceId returns a reference to field data_source_id of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) DataSourceId() terra.StringValue {
	return terra.ReferenceAsString(bdtc.ref.Append("data_source_id"))
}

// DestinationDatasetId returns a reference to field destination_dataset_id of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) DestinationDatasetId() terra.StringValue {
	return terra.ReferenceAsString(bdtc.ref.Append("destination_dataset_id"))
}

// Disabled returns a reference to field disabled of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(bdtc.ref.Append("disabled"))
}

// DisplayName returns a reference to field display_name of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bdtc.ref.Append("display_name"))
}

// Id returns a reference to field id of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bdtc.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bdtc.ref.Append("location"))
}

// Name returns a reference to field name of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bdtc.ref.Append("name"))
}

// NotificationPubsubTopic returns a reference to field notification_pubsub_topic of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) NotificationPubsubTopic() terra.StringValue {
	return terra.ReferenceAsString(bdtc.ref.Append("notification_pubsub_topic"))
}

// Params returns a reference to field params of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) Params() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bdtc.ref.Append("params"))
}

// Project returns a reference to field project of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bdtc.ref.Append("project"))
}

// Schedule returns a reference to field schedule of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(bdtc.ref.Append("schedule"))
}

// ServiceAccountName returns a reference to field service_account_name of google_bigquery_data_transfer_config.
func (bdtc bigqueryDataTransferConfigAttributes) ServiceAccountName() terra.StringValue {
	return terra.ReferenceAsString(bdtc.ref.Append("service_account_name"))
}

func (bdtc bigqueryDataTransferConfigAttributes) EmailPreferences() terra.ListValue[bigquerydatatransferconfig.EmailPreferencesAttributes] {
	return terra.ReferenceAsList[bigquerydatatransferconfig.EmailPreferencesAttributes](bdtc.ref.Append("email_preferences"))
}

func (bdtc bigqueryDataTransferConfigAttributes) ScheduleOptions() terra.ListValue[bigquerydatatransferconfig.ScheduleOptionsAttributes] {
	return terra.ReferenceAsList[bigquerydatatransferconfig.ScheduleOptionsAttributes](bdtc.ref.Append("schedule_options"))
}

func (bdtc bigqueryDataTransferConfigAttributes) SensitiveParams() terra.ListValue[bigquerydatatransferconfig.SensitiveParamsAttributes] {
	return terra.ReferenceAsList[bigquerydatatransferconfig.SensitiveParamsAttributes](bdtc.ref.Append("sensitive_params"))
}

func (bdtc bigqueryDataTransferConfigAttributes) Timeouts() bigquerydatatransferconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigquerydatatransferconfig.TimeoutsAttributes](bdtc.ref.Append("timeouts"))
}

type bigqueryDataTransferConfigState struct {
	DataRefreshWindowDays   float64                                            `json:"data_refresh_window_days"`
	DataSourceId            string                                             `json:"data_source_id"`
	DestinationDatasetId    string                                             `json:"destination_dataset_id"`
	Disabled                bool                                               `json:"disabled"`
	DisplayName             string                                             `json:"display_name"`
	Id                      string                                             `json:"id"`
	Location                string                                             `json:"location"`
	Name                    string                                             `json:"name"`
	NotificationPubsubTopic string                                             `json:"notification_pubsub_topic"`
	Params                  map[string]string                                  `json:"params"`
	Project                 string                                             `json:"project"`
	Schedule                string                                             `json:"schedule"`
	ServiceAccountName      string                                             `json:"service_account_name"`
	EmailPreferences        []bigquerydatatransferconfig.EmailPreferencesState `json:"email_preferences"`
	ScheduleOptions         []bigquerydatatransferconfig.ScheduleOptionsState  `json:"schedule_options"`
	SensitiveParams         []bigquerydatatransferconfig.SensitiveParamsState  `json:"sensitive_params"`
	Timeouts                *bigquerydatatransferconfig.TimeoutsState          `json:"timeouts"`
}
