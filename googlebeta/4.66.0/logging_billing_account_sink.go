// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	loggingbillingaccountsink "github.com/golingon/terraproviders/googlebeta/4.66.0/loggingbillingaccountsink"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingBillingAccountSink creates a new instance of [LoggingBillingAccountSink].
func NewLoggingBillingAccountSink(name string, args LoggingBillingAccountSinkArgs) *LoggingBillingAccountSink {
	return &LoggingBillingAccountSink{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingBillingAccountSink)(nil)

// LoggingBillingAccountSink represents the Terraform resource google_logging_billing_account_sink.
type LoggingBillingAccountSink struct {
	Name      string
	Args      LoggingBillingAccountSinkArgs
	state     *loggingBillingAccountSinkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingBillingAccountSink].
func (lbas *LoggingBillingAccountSink) Type() string {
	return "google_logging_billing_account_sink"
}

// LocalName returns the local name for [LoggingBillingAccountSink].
func (lbas *LoggingBillingAccountSink) LocalName() string {
	return lbas.Name
}

// Configuration returns the configuration (args) for [LoggingBillingAccountSink].
func (lbas *LoggingBillingAccountSink) Configuration() interface{} {
	return lbas.Args
}

// DependOn is used for other resources to depend on [LoggingBillingAccountSink].
func (lbas *LoggingBillingAccountSink) DependOn() terra.Reference {
	return terra.ReferenceResource(lbas)
}

// Dependencies returns the list of resources [LoggingBillingAccountSink] depends_on.
func (lbas *LoggingBillingAccountSink) Dependencies() terra.Dependencies {
	return lbas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingBillingAccountSink].
func (lbas *LoggingBillingAccountSink) LifecycleManagement() *terra.Lifecycle {
	return lbas.Lifecycle
}

// Attributes returns the attributes for [LoggingBillingAccountSink].
func (lbas *LoggingBillingAccountSink) Attributes() loggingBillingAccountSinkAttributes {
	return loggingBillingAccountSinkAttributes{ref: terra.ReferenceResource(lbas)}
}

// ImportState imports the given attribute values into [LoggingBillingAccountSink]'s state.
func (lbas *LoggingBillingAccountSink) ImportState(av io.Reader) error {
	lbas.state = &loggingBillingAccountSinkState{}
	if err := json.NewDecoder(av).Decode(lbas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lbas.Type(), lbas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingBillingAccountSink] has state.
func (lbas *LoggingBillingAccountSink) State() (*loggingBillingAccountSinkState, bool) {
	return lbas.state, lbas.state != nil
}

// StateMust returns the state for [LoggingBillingAccountSink]. Panics if the state is nil.
func (lbas *LoggingBillingAccountSink) StateMust() *loggingBillingAccountSinkState {
	if lbas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lbas.Type(), lbas.LocalName()))
	}
	return lbas.state
}

// LoggingBillingAccountSinkArgs contains the configurations for google_logging_billing_account_sink.
type LoggingBillingAccountSinkArgs struct {
	// BillingAccount: string, required
	BillingAccount terra.StringValue `hcl:"billing_account,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Destination: string, required
	Destination terra.StringValue `hcl:"destination,attr" validate:"required"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// BigqueryOptions: optional
	BigqueryOptions *loggingbillingaccountsink.BigqueryOptions `hcl:"bigquery_options,block"`
	// Exclusions: min=0
	Exclusions []loggingbillingaccountsink.Exclusions `hcl:"exclusions,block" validate:"min=0"`
}
type loggingBillingAccountSinkAttributes struct {
	ref terra.Reference
}

// BillingAccount returns a reference to field billing_account of google_logging_billing_account_sink.
func (lbas loggingBillingAccountSinkAttributes) BillingAccount() terra.StringValue {
	return terra.ReferenceAsString(lbas.ref.Append("billing_account"))
}

// Description returns a reference to field description of google_logging_billing_account_sink.
func (lbas loggingBillingAccountSinkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lbas.ref.Append("description"))
}

// Destination returns a reference to field destination of google_logging_billing_account_sink.
func (lbas loggingBillingAccountSinkAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(lbas.ref.Append("destination"))
}

// Disabled returns a reference to field disabled of google_logging_billing_account_sink.
func (lbas loggingBillingAccountSinkAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(lbas.ref.Append("disabled"))
}

// Filter returns a reference to field filter of google_logging_billing_account_sink.
func (lbas loggingBillingAccountSinkAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(lbas.ref.Append("filter"))
}

// Id returns a reference to field id of google_logging_billing_account_sink.
func (lbas loggingBillingAccountSinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lbas.ref.Append("id"))
}

// Name returns a reference to field name of google_logging_billing_account_sink.
func (lbas loggingBillingAccountSinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lbas.ref.Append("name"))
}

// WriterIdentity returns a reference to field writer_identity of google_logging_billing_account_sink.
func (lbas loggingBillingAccountSinkAttributes) WriterIdentity() terra.StringValue {
	return terra.ReferenceAsString(lbas.ref.Append("writer_identity"))
}

func (lbas loggingBillingAccountSinkAttributes) BigqueryOptions() terra.ListValue[loggingbillingaccountsink.BigqueryOptionsAttributes] {
	return terra.ReferenceAsList[loggingbillingaccountsink.BigqueryOptionsAttributes](lbas.ref.Append("bigquery_options"))
}

func (lbas loggingBillingAccountSinkAttributes) Exclusions() terra.ListValue[loggingbillingaccountsink.ExclusionsAttributes] {
	return terra.ReferenceAsList[loggingbillingaccountsink.ExclusionsAttributes](lbas.ref.Append("exclusions"))
}

type loggingBillingAccountSinkState struct {
	BillingAccount  string                                           `json:"billing_account"`
	Description     string                                           `json:"description"`
	Destination     string                                           `json:"destination"`
	Disabled        bool                                             `json:"disabled"`
	Filter          string                                           `json:"filter"`
	Id              string                                           `json:"id"`
	Name            string                                           `json:"name"`
	WriterIdentity  string                                           `json:"writer_identity"`
	BigqueryOptions []loggingbillingaccountsink.BigqueryOptionsState `json:"bigquery_options"`
	Exclusions      []loggingbillingaccountsink.ExclusionsState      `json:"exclusions"`
}
