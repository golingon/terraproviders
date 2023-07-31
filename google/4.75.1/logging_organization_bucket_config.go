// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	loggingorganizationbucketconfig "github.com/golingon/terraproviders/google/4.75.1/loggingorganizationbucketconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingOrganizationBucketConfig creates a new instance of [LoggingOrganizationBucketConfig].
func NewLoggingOrganizationBucketConfig(name string, args LoggingOrganizationBucketConfigArgs) *LoggingOrganizationBucketConfig {
	return &LoggingOrganizationBucketConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingOrganizationBucketConfig)(nil)

// LoggingOrganizationBucketConfig represents the Terraform resource google_logging_organization_bucket_config.
type LoggingOrganizationBucketConfig struct {
	Name      string
	Args      LoggingOrganizationBucketConfigArgs
	state     *loggingOrganizationBucketConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingOrganizationBucketConfig].
func (lobc *LoggingOrganizationBucketConfig) Type() string {
	return "google_logging_organization_bucket_config"
}

// LocalName returns the local name for [LoggingOrganizationBucketConfig].
func (lobc *LoggingOrganizationBucketConfig) LocalName() string {
	return lobc.Name
}

// Configuration returns the configuration (args) for [LoggingOrganizationBucketConfig].
func (lobc *LoggingOrganizationBucketConfig) Configuration() interface{} {
	return lobc.Args
}

// DependOn is used for other resources to depend on [LoggingOrganizationBucketConfig].
func (lobc *LoggingOrganizationBucketConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(lobc)
}

// Dependencies returns the list of resources [LoggingOrganizationBucketConfig] depends_on.
func (lobc *LoggingOrganizationBucketConfig) Dependencies() terra.Dependencies {
	return lobc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingOrganizationBucketConfig].
func (lobc *LoggingOrganizationBucketConfig) LifecycleManagement() *terra.Lifecycle {
	return lobc.Lifecycle
}

// Attributes returns the attributes for [LoggingOrganizationBucketConfig].
func (lobc *LoggingOrganizationBucketConfig) Attributes() loggingOrganizationBucketConfigAttributes {
	return loggingOrganizationBucketConfigAttributes{ref: terra.ReferenceResource(lobc)}
}

// ImportState imports the given attribute values into [LoggingOrganizationBucketConfig]'s state.
func (lobc *LoggingOrganizationBucketConfig) ImportState(av io.Reader) error {
	lobc.state = &loggingOrganizationBucketConfigState{}
	if err := json.NewDecoder(av).Decode(lobc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lobc.Type(), lobc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingOrganizationBucketConfig] has state.
func (lobc *LoggingOrganizationBucketConfig) State() (*loggingOrganizationBucketConfigState, bool) {
	return lobc.state, lobc.state != nil
}

// StateMust returns the state for [LoggingOrganizationBucketConfig]. Panics if the state is nil.
func (lobc *LoggingOrganizationBucketConfig) StateMust() *loggingOrganizationBucketConfigState {
	if lobc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lobc.Type(), lobc.LocalName()))
	}
	return lobc.state
}

// LoggingOrganizationBucketConfigArgs contains the configurations for google_logging_organization_bucket_config.
type LoggingOrganizationBucketConfigArgs struct {
	// BucketId: string, required
	BucketId terra.StringValue `hcl:"bucket_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// RetentionDays: number, optional
	RetentionDays terra.NumberValue `hcl:"retention_days,attr"`
	// CmekSettings: optional
	CmekSettings *loggingorganizationbucketconfig.CmekSettings `hcl:"cmek_settings,block"`
}
type loggingOrganizationBucketConfigAttributes struct {
	ref terra.Reference
}

// BucketId returns a reference to field bucket_id of google_logging_organization_bucket_config.
func (lobc loggingOrganizationBucketConfigAttributes) BucketId() terra.StringValue {
	return terra.ReferenceAsString(lobc.ref.Append("bucket_id"))
}

// Description returns a reference to field description of google_logging_organization_bucket_config.
func (lobc loggingOrganizationBucketConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lobc.ref.Append("description"))
}

// Id returns a reference to field id of google_logging_organization_bucket_config.
func (lobc loggingOrganizationBucketConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lobc.ref.Append("id"))
}

// LifecycleState returns a reference to field lifecycle_state of google_logging_organization_bucket_config.
func (lobc loggingOrganizationBucketConfigAttributes) LifecycleState() terra.StringValue {
	return terra.ReferenceAsString(lobc.ref.Append("lifecycle_state"))
}

// Location returns a reference to field location of google_logging_organization_bucket_config.
func (lobc loggingOrganizationBucketConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lobc.ref.Append("location"))
}

// Name returns a reference to field name of google_logging_organization_bucket_config.
func (lobc loggingOrganizationBucketConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lobc.ref.Append("name"))
}

// Organization returns a reference to field organization of google_logging_organization_bucket_config.
func (lobc loggingOrganizationBucketConfigAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(lobc.ref.Append("organization"))
}

// RetentionDays returns a reference to field retention_days of google_logging_organization_bucket_config.
func (lobc loggingOrganizationBucketConfigAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(lobc.ref.Append("retention_days"))
}

func (lobc loggingOrganizationBucketConfigAttributes) CmekSettings() terra.ListValue[loggingorganizationbucketconfig.CmekSettingsAttributes] {
	return terra.ReferenceAsList[loggingorganizationbucketconfig.CmekSettingsAttributes](lobc.ref.Append("cmek_settings"))
}

type loggingOrganizationBucketConfigState struct {
	BucketId       string                                              `json:"bucket_id"`
	Description    string                                              `json:"description"`
	Id             string                                              `json:"id"`
	LifecycleState string                                              `json:"lifecycle_state"`
	Location       string                                              `json:"location"`
	Name           string                                              `json:"name"`
	Organization   string                                              `json:"organization"`
	RetentionDays  float64                                             `json:"retention_days"`
	CmekSettings   []loggingorganizationbucketconfig.CmekSettingsState `json:"cmek_settings"`
}
