// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	loggingfolderbucketconfig "github.com/golingon/terraproviders/googlebeta/4.76.0/loggingfolderbucketconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingFolderBucketConfig creates a new instance of [LoggingFolderBucketConfig].
func NewLoggingFolderBucketConfig(name string, args LoggingFolderBucketConfigArgs) *LoggingFolderBucketConfig {
	return &LoggingFolderBucketConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingFolderBucketConfig)(nil)

// LoggingFolderBucketConfig represents the Terraform resource google_logging_folder_bucket_config.
type LoggingFolderBucketConfig struct {
	Name      string
	Args      LoggingFolderBucketConfigArgs
	state     *loggingFolderBucketConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingFolderBucketConfig].
func (lfbc *LoggingFolderBucketConfig) Type() string {
	return "google_logging_folder_bucket_config"
}

// LocalName returns the local name for [LoggingFolderBucketConfig].
func (lfbc *LoggingFolderBucketConfig) LocalName() string {
	return lfbc.Name
}

// Configuration returns the configuration (args) for [LoggingFolderBucketConfig].
func (lfbc *LoggingFolderBucketConfig) Configuration() interface{} {
	return lfbc.Args
}

// DependOn is used for other resources to depend on [LoggingFolderBucketConfig].
func (lfbc *LoggingFolderBucketConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(lfbc)
}

// Dependencies returns the list of resources [LoggingFolderBucketConfig] depends_on.
func (lfbc *LoggingFolderBucketConfig) Dependencies() terra.Dependencies {
	return lfbc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingFolderBucketConfig].
func (lfbc *LoggingFolderBucketConfig) LifecycleManagement() *terra.Lifecycle {
	return lfbc.Lifecycle
}

// Attributes returns the attributes for [LoggingFolderBucketConfig].
func (lfbc *LoggingFolderBucketConfig) Attributes() loggingFolderBucketConfigAttributes {
	return loggingFolderBucketConfigAttributes{ref: terra.ReferenceResource(lfbc)}
}

// ImportState imports the given attribute values into [LoggingFolderBucketConfig]'s state.
func (lfbc *LoggingFolderBucketConfig) ImportState(av io.Reader) error {
	lfbc.state = &loggingFolderBucketConfigState{}
	if err := json.NewDecoder(av).Decode(lfbc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lfbc.Type(), lfbc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingFolderBucketConfig] has state.
func (lfbc *LoggingFolderBucketConfig) State() (*loggingFolderBucketConfigState, bool) {
	return lfbc.state, lfbc.state != nil
}

// StateMust returns the state for [LoggingFolderBucketConfig]. Panics if the state is nil.
func (lfbc *LoggingFolderBucketConfig) StateMust() *loggingFolderBucketConfigState {
	if lfbc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lfbc.Type(), lfbc.LocalName()))
	}
	return lfbc.state
}

// LoggingFolderBucketConfigArgs contains the configurations for google_logging_folder_bucket_config.
type LoggingFolderBucketConfigArgs struct {
	// BucketId: string, required
	BucketId terra.StringValue `hcl:"bucket_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Folder: string, required
	Folder terra.StringValue `hcl:"folder,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// RetentionDays: number, optional
	RetentionDays terra.NumberValue `hcl:"retention_days,attr"`
	// CmekSettings: optional
	CmekSettings *loggingfolderbucketconfig.CmekSettings `hcl:"cmek_settings,block"`
}
type loggingFolderBucketConfigAttributes struct {
	ref terra.Reference
}

// BucketId returns a reference to field bucket_id of google_logging_folder_bucket_config.
func (lfbc loggingFolderBucketConfigAttributes) BucketId() terra.StringValue {
	return terra.ReferenceAsString(lfbc.ref.Append("bucket_id"))
}

// Description returns a reference to field description of google_logging_folder_bucket_config.
func (lfbc loggingFolderBucketConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lfbc.ref.Append("description"))
}

// Folder returns a reference to field folder of google_logging_folder_bucket_config.
func (lfbc loggingFolderBucketConfigAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(lfbc.ref.Append("folder"))
}

// Id returns a reference to field id of google_logging_folder_bucket_config.
func (lfbc loggingFolderBucketConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lfbc.ref.Append("id"))
}

// LifecycleState returns a reference to field lifecycle_state of google_logging_folder_bucket_config.
func (lfbc loggingFolderBucketConfigAttributes) LifecycleState() terra.StringValue {
	return terra.ReferenceAsString(lfbc.ref.Append("lifecycle_state"))
}

// Location returns a reference to field location of google_logging_folder_bucket_config.
func (lfbc loggingFolderBucketConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lfbc.ref.Append("location"))
}

// Name returns a reference to field name of google_logging_folder_bucket_config.
func (lfbc loggingFolderBucketConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lfbc.ref.Append("name"))
}

// RetentionDays returns a reference to field retention_days of google_logging_folder_bucket_config.
func (lfbc loggingFolderBucketConfigAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(lfbc.ref.Append("retention_days"))
}

func (lfbc loggingFolderBucketConfigAttributes) CmekSettings() terra.ListValue[loggingfolderbucketconfig.CmekSettingsAttributes] {
	return terra.ReferenceAsList[loggingfolderbucketconfig.CmekSettingsAttributes](lfbc.ref.Append("cmek_settings"))
}

type loggingFolderBucketConfigState struct {
	BucketId       string                                        `json:"bucket_id"`
	Description    string                                        `json:"description"`
	Folder         string                                        `json:"folder"`
	Id             string                                        `json:"id"`
	LifecycleState string                                        `json:"lifecycle_state"`
	Location       string                                        `json:"location"`
	Name           string                                        `json:"name"`
	RetentionDays  float64                                       `json:"retention_days"`
	CmekSettings   []loggingfolderbucketconfig.CmekSettingsState `json:"cmek_settings"`
}
