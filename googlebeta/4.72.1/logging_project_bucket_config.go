// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	loggingprojectbucketconfig "github.com/golingon/terraproviders/googlebeta/4.72.1/loggingprojectbucketconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingProjectBucketConfig creates a new instance of [LoggingProjectBucketConfig].
func NewLoggingProjectBucketConfig(name string, args LoggingProjectBucketConfigArgs) *LoggingProjectBucketConfig {
	return &LoggingProjectBucketConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingProjectBucketConfig)(nil)

// LoggingProjectBucketConfig represents the Terraform resource google_logging_project_bucket_config.
type LoggingProjectBucketConfig struct {
	Name      string
	Args      LoggingProjectBucketConfigArgs
	state     *loggingProjectBucketConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingProjectBucketConfig].
func (lpbc *LoggingProjectBucketConfig) Type() string {
	return "google_logging_project_bucket_config"
}

// LocalName returns the local name for [LoggingProjectBucketConfig].
func (lpbc *LoggingProjectBucketConfig) LocalName() string {
	return lpbc.Name
}

// Configuration returns the configuration (args) for [LoggingProjectBucketConfig].
func (lpbc *LoggingProjectBucketConfig) Configuration() interface{} {
	return lpbc.Args
}

// DependOn is used for other resources to depend on [LoggingProjectBucketConfig].
func (lpbc *LoggingProjectBucketConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(lpbc)
}

// Dependencies returns the list of resources [LoggingProjectBucketConfig] depends_on.
func (lpbc *LoggingProjectBucketConfig) Dependencies() terra.Dependencies {
	return lpbc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingProjectBucketConfig].
func (lpbc *LoggingProjectBucketConfig) LifecycleManagement() *terra.Lifecycle {
	return lpbc.Lifecycle
}

// Attributes returns the attributes for [LoggingProjectBucketConfig].
func (lpbc *LoggingProjectBucketConfig) Attributes() loggingProjectBucketConfigAttributes {
	return loggingProjectBucketConfigAttributes{ref: terra.ReferenceResource(lpbc)}
}

// ImportState imports the given attribute values into [LoggingProjectBucketConfig]'s state.
func (lpbc *LoggingProjectBucketConfig) ImportState(av io.Reader) error {
	lpbc.state = &loggingProjectBucketConfigState{}
	if err := json.NewDecoder(av).Decode(lpbc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lpbc.Type(), lpbc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingProjectBucketConfig] has state.
func (lpbc *LoggingProjectBucketConfig) State() (*loggingProjectBucketConfigState, bool) {
	return lpbc.state, lpbc.state != nil
}

// StateMust returns the state for [LoggingProjectBucketConfig]. Panics if the state is nil.
func (lpbc *LoggingProjectBucketConfig) StateMust() *loggingProjectBucketConfigState {
	if lpbc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lpbc.Type(), lpbc.LocalName()))
	}
	return lpbc.state
}

// LoggingProjectBucketConfigArgs contains the configurations for google_logging_project_bucket_config.
type LoggingProjectBucketConfigArgs struct {
	// BucketId: string, required
	BucketId terra.StringValue `hcl:"bucket_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnableAnalytics: bool, optional
	EnableAnalytics terra.BoolValue `hcl:"enable_analytics,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Locked: bool, optional
	Locked terra.BoolValue `hcl:"locked,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// RetentionDays: number, optional
	RetentionDays terra.NumberValue `hcl:"retention_days,attr"`
	// CmekSettings: optional
	CmekSettings *loggingprojectbucketconfig.CmekSettings `hcl:"cmek_settings,block"`
}
type loggingProjectBucketConfigAttributes struct {
	ref terra.Reference
}

// BucketId returns a reference to field bucket_id of google_logging_project_bucket_config.
func (lpbc loggingProjectBucketConfigAttributes) BucketId() terra.StringValue {
	return terra.ReferenceAsString(lpbc.ref.Append("bucket_id"))
}

// Description returns a reference to field description of google_logging_project_bucket_config.
func (lpbc loggingProjectBucketConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lpbc.ref.Append("description"))
}

// EnableAnalytics returns a reference to field enable_analytics of google_logging_project_bucket_config.
func (lpbc loggingProjectBucketConfigAttributes) EnableAnalytics() terra.BoolValue {
	return terra.ReferenceAsBool(lpbc.ref.Append("enable_analytics"))
}

// Id returns a reference to field id of google_logging_project_bucket_config.
func (lpbc loggingProjectBucketConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lpbc.ref.Append("id"))
}

// LifecycleState returns a reference to field lifecycle_state of google_logging_project_bucket_config.
func (lpbc loggingProjectBucketConfigAttributes) LifecycleState() terra.StringValue {
	return terra.ReferenceAsString(lpbc.ref.Append("lifecycle_state"))
}

// Location returns a reference to field location of google_logging_project_bucket_config.
func (lpbc loggingProjectBucketConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lpbc.ref.Append("location"))
}

// Locked returns a reference to field locked of google_logging_project_bucket_config.
func (lpbc loggingProjectBucketConfigAttributes) Locked() terra.BoolValue {
	return terra.ReferenceAsBool(lpbc.ref.Append("locked"))
}

// Name returns a reference to field name of google_logging_project_bucket_config.
func (lpbc loggingProjectBucketConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lpbc.ref.Append("name"))
}

// Project returns a reference to field project of google_logging_project_bucket_config.
func (lpbc loggingProjectBucketConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(lpbc.ref.Append("project"))
}

// RetentionDays returns a reference to field retention_days of google_logging_project_bucket_config.
func (lpbc loggingProjectBucketConfigAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(lpbc.ref.Append("retention_days"))
}

func (lpbc loggingProjectBucketConfigAttributes) CmekSettings() terra.ListValue[loggingprojectbucketconfig.CmekSettingsAttributes] {
	return terra.ReferenceAsList[loggingprojectbucketconfig.CmekSettingsAttributes](lpbc.ref.Append("cmek_settings"))
}

type loggingProjectBucketConfigState struct {
	BucketId        string                                         `json:"bucket_id"`
	Description     string                                         `json:"description"`
	EnableAnalytics bool                                           `json:"enable_analytics"`
	Id              string                                         `json:"id"`
	LifecycleState  string                                         `json:"lifecycle_state"`
	Location        string                                         `json:"location"`
	Locked          bool                                           `json:"locked"`
	Name            string                                         `json:"name"`
	Project         string                                         `json:"project"`
	RetentionDays   float64                                        `json:"retention_days"`
	CmekSettings    []loggingprojectbucketconfig.CmekSettingsState `json:"cmek_settings"`
}
