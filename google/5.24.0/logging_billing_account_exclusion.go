// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewLoggingBillingAccountExclusion creates a new instance of [LoggingBillingAccountExclusion].
func NewLoggingBillingAccountExclusion(name string, args LoggingBillingAccountExclusionArgs) *LoggingBillingAccountExclusion {
	return &LoggingBillingAccountExclusion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingBillingAccountExclusion)(nil)

// LoggingBillingAccountExclusion represents the Terraform resource google_logging_billing_account_exclusion.
type LoggingBillingAccountExclusion struct {
	Name      string
	Args      LoggingBillingAccountExclusionArgs
	state     *loggingBillingAccountExclusionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingBillingAccountExclusion].
func (lbae *LoggingBillingAccountExclusion) Type() string {
	return "google_logging_billing_account_exclusion"
}

// LocalName returns the local name for [LoggingBillingAccountExclusion].
func (lbae *LoggingBillingAccountExclusion) LocalName() string {
	return lbae.Name
}

// Configuration returns the configuration (args) for [LoggingBillingAccountExclusion].
func (lbae *LoggingBillingAccountExclusion) Configuration() interface{} {
	return lbae.Args
}

// DependOn is used for other resources to depend on [LoggingBillingAccountExclusion].
func (lbae *LoggingBillingAccountExclusion) DependOn() terra.Reference {
	return terra.ReferenceResource(lbae)
}

// Dependencies returns the list of resources [LoggingBillingAccountExclusion] depends_on.
func (lbae *LoggingBillingAccountExclusion) Dependencies() terra.Dependencies {
	return lbae.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingBillingAccountExclusion].
func (lbae *LoggingBillingAccountExclusion) LifecycleManagement() *terra.Lifecycle {
	return lbae.Lifecycle
}

// Attributes returns the attributes for [LoggingBillingAccountExclusion].
func (lbae *LoggingBillingAccountExclusion) Attributes() loggingBillingAccountExclusionAttributes {
	return loggingBillingAccountExclusionAttributes{ref: terra.ReferenceResource(lbae)}
}

// ImportState imports the given attribute values into [LoggingBillingAccountExclusion]'s state.
func (lbae *LoggingBillingAccountExclusion) ImportState(av io.Reader) error {
	lbae.state = &loggingBillingAccountExclusionState{}
	if err := json.NewDecoder(av).Decode(lbae.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lbae.Type(), lbae.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingBillingAccountExclusion] has state.
func (lbae *LoggingBillingAccountExclusion) State() (*loggingBillingAccountExclusionState, bool) {
	return lbae.state, lbae.state != nil
}

// StateMust returns the state for [LoggingBillingAccountExclusion]. Panics if the state is nil.
func (lbae *LoggingBillingAccountExclusion) StateMust() *loggingBillingAccountExclusionState {
	if lbae.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lbae.Type(), lbae.LocalName()))
	}
	return lbae.state
}

// LoggingBillingAccountExclusionArgs contains the configurations for google_logging_billing_account_exclusion.
type LoggingBillingAccountExclusionArgs struct {
	// BillingAccount: string, required
	BillingAccount terra.StringValue `hcl:"billing_account,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Filter: string, required
	Filter terra.StringValue `hcl:"filter,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type loggingBillingAccountExclusionAttributes struct {
	ref terra.Reference
}

// BillingAccount returns a reference to field billing_account of google_logging_billing_account_exclusion.
func (lbae loggingBillingAccountExclusionAttributes) BillingAccount() terra.StringValue {
	return terra.ReferenceAsString(lbae.ref.Append("billing_account"))
}

// Description returns a reference to field description of google_logging_billing_account_exclusion.
func (lbae loggingBillingAccountExclusionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lbae.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_logging_billing_account_exclusion.
func (lbae loggingBillingAccountExclusionAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(lbae.ref.Append("disabled"))
}

// Filter returns a reference to field filter of google_logging_billing_account_exclusion.
func (lbae loggingBillingAccountExclusionAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(lbae.ref.Append("filter"))
}

// Id returns a reference to field id of google_logging_billing_account_exclusion.
func (lbae loggingBillingAccountExclusionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lbae.ref.Append("id"))
}

// Name returns a reference to field name of google_logging_billing_account_exclusion.
func (lbae loggingBillingAccountExclusionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lbae.ref.Append("name"))
}

type loggingBillingAccountExclusionState struct {
	BillingAccount string `json:"billing_account"`
	Description    string `json:"description"`
	Disabled       bool   `json:"disabled"`
	Filter         string `json:"filter"`
	Id             string `json:"id"`
	Name           string `json:"name"`
}
