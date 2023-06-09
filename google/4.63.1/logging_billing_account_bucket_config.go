// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	loggingbillingaccountbucketconfig "github.com/golingon/terraproviders/google/4.63.1/loggingbillingaccountbucketconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingBillingAccountBucketConfig creates a new instance of [LoggingBillingAccountBucketConfig].
func NewLoggingBillingAccountBucketConfig(name string, args LoggingBillingAccountBucketConfigArgs) *LoggingBillingAccountBucketConfig {
	return &LoggingBillingAccountBucketConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingBillingAccountBucketConfig)(nil)

// LoggingBillingAccountBucketConfig represents the Terraform resource google_logging_billing_account_bucket_config.
type LoggingBillingAccountBucketConfig struct {
	Name      string
	Args      LoggingBillingAccountBucketConfigArgs
	state     *loggingBillingAccountBucketConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingBillingAccountBucketConfig].
func (lbabc *LoggingBillingAccountBucketConfig) Type() string {
	return "google_logging_billing_account_bucket_config"
}

// LocalName returns the local name for [LoggingBillingAccountBucketConfig].
func (lbabc *LoggingBillingAccountBucketConfig) LocalName() string {
	return lbabc.Name
}

// Configuration returns the configuration (args) for [LoggingBillingAccountBucketConfig].
func (lbabc *LoggingBillingAccountBucketConfig) Configuration() interface{} {
	return lbabc.Args
}

// DependOn is used for other resources to depend on [LoggingBillingAccountBucketConfig].
func (lbabc *LoggingBillingAccountBucketConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(lbabc)
}

// Dependencies returns the list of resources [LoggingBillingAccountBucketConfig] depends_on.
func (lbabc *LoggingBillingAccountBucketConfig) Dependencies() terra.Dependencies {
	return lbabc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingBillingAccountBucketConfig].
func (lbabc *LoggingBillingAccountBucketConfig) LifecycleManagement() *terra.Lifecycle {
	return lbabc.Lifecycle
}

// Attributes returns the attributes for [LoggingBillingAccountBucketConfig].
func (lbabc *LoggingBillingAccountBucketConfig) Attributes() loggingBillingAccountBucketConfigAttributes {
	return loggingBillingAccountBucketConfigAttributes{ref: terra.ReferenceResource(lbabc)}
}

// ImportState imports the given attribute values into [LoggingBillingAccountBucketConfig]'s state.
func (lbabc *LoggingBillingAccountBucketConfig) ImportState(av io.Reader) error {
	lbabc.state = &loggingBillingAccountBucketConfigState{}
	if err := json.NewDecoder(av).Decode(lbabc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lbabc.Type(), lbabc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingBillingAccountBucketConfig] has state.
func (lbabc *LoggingBillingAccountBucketConfig) State() (*loggingBillingAccountBucketConfigState, bool) {
	return lbabc.state, lbabc.state != nil
}

// StateMust returns the state for [LoggingBillingAccountBucketConfig]. Panics if the state is nil.
func (lbabc *LoggingBillingAccountBucketConfig) StateMust() *loggingBillingAccountBucketConfigState {
	if lbabc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lbabc.Type(), lbabc.LocalName()))
	}
	return lbabc.state
}

// LoggingBillingAccountBucketConfigArgs contains the configurations for google_logging_billing_account_bucket_config.
type LoggingBillingAccountBucketConfigArgs struct {
	// BillingAccount: string, required
	BillingAccount terra.StringValue `hcl:"billing_account,attr" validate:"required"`
	// BucketId: string, required
	BucketId terra.StringValue `hcl:"bucket_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// RetentionDays: number, optional
	RetentionDays terra.NumberValue `hcl:"retention_days,attr"`
	// CmekSettings: optional
	CmekSettings *loggingbillingaccountbucketconfig.CmekSettings `hcl:"cmek_settings,block"`
}
type loggingBillingAccountBucketConfigAttributes struct {
	ref terra.Reference
}

// BillingAccount returns a reference to field billing_account of google_logging_billing_account_bucket_config.
func (lbabc loggingBillingAccountBucketConfigAttributes) BillingAccount() terra.StringValue {
	return terra.ReferenceAsString(lbabc.ref.Append("billing_account"))
}

// BucketId returns a reference to field bucket_id of google_logging_billing_account_bucket_config.
func (lbabc loggingBillingAccountBucketConfigAttributes) BucketId() terra.StringValue {
	return terra.ReferenceAsString(lbabc.ref.Append("bucket_id"))
}

// Description returns a reference to field description of google_logging_billing_account_bucket_config.
func (lbabc loggingBillingAccountBucketConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lbabc.ref.Append("description"))
}

// Id returns a reference to field id of google_logging_billing_account_bucket_config.
func (lbabc loggingBillingAccountBucketConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lbabc.ref.Append("id"))
}

// LifecycleState returns a reference to field lifecycle_state of google_logging_billing_account_bucket_config.
func (lbabc loggingBillingAccountBucketConfigAttributes) LifecycleState() terra.StringValue {
	return terra.ReferenceAsString(lbabc.ref.Append("lifecycle_state"))
}

// Location returns a reference to field location of google_logging_billing_account_bucket_config.
func (lbabc loggingBillingAccountBucketConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lbabc.ref.Append("location"))
}

// Name returns a reference to field name of google_logging_billing_account_bucket_config.
func (lbabc loggingBillingAccountBucketConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lbabc.ref.Append("name"))
}

// RetentionDays returns a reference to field retention_days of google_logging_billing_account_bucket_config.
func (lbabc loggingBillingAccountBucketConfigAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(lbabc.ref.Append("retention_days"))
}

func (lbabc loggingBillingAccountBucketConfigAttributes) CmekSettings() terra.ListValue[loggingbillingaccountbucketconfig.CmekSettingsAttributes] {
	return terra.ReferenceAsList[loggingbillingaccountbucketconfig.CmekSettingsAttributes](lbabc.ref.Append("cmek_settings"))
}

type loggingBillingAccountBucketConfigState struct {
	BillingAccount string                                                `json:"billing_account"`
	BucketId       string                                                `json:"bucket_id"`
	Description    string                                                `json:"description"`
	Id             string                                                `json:"id"`
	LifecycleState string                                                `json:"lifecycle_state"`
	Location       string                                                `json:"location"`
	Name           string                                                `json:"name"`
	RetentionDays  float64                                               `json:"retention_days"`
	CmekSettings   []loggingbillingaccountbucketconfig.CmekSettingsState `json:"cmek_settings"`
}
