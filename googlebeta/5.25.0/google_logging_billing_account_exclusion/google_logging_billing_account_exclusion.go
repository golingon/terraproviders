// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_logging_billing_account_exclusion

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

// Resource represents the Terraform resource google_logging_billing_account_exclusion.
type Resource struct {
	Name      string
	Args      Args
	state     *googleLoggingBillingAccountExclusionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (glbae *Resource) Type() string {
	return "google_logging_billing_account_exclusion"
}

// LocalName returns the local name for [Resource].
func (glbae *Resource) LocalName() string {
	return glbae.Name
}

// Configuration returns the configuration (args) for [Resource].
func (glbae *Resource) Configuration() interface{} {
	return glbae.Args
}

// DependOn is used for other resources to depend on [Resource].
func (glbae *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(glbae)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (glbae *Resource) Dependencies() terra.Dependencies {
	return glbae.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (glbae *Resource) LifecycleManagement() *terra.Lifecycle {
	return glbae.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (glbae *Resource) Attributes() googleLoggingBillingAccountExclusionAttributes {
	return googleLoggingBillingAccountExclusionAttributes{ref: terra.ReferenceResource(glbae)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (glbae *Resource) ImportState(state io.Reader) error {
	glbae.state = &googleLoggingBillingAccountExclusionState{}
	if err := json.NewDecoder(state).Decode(glbae.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", glbae.Type(), glbae.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (glbae *Resource) State() (*googleLoggingBillingAccountExclusionState, bool) {
	return glbae.state, glbae.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (glbae *Resource) StateMust() *googleLoggingBillingAccountExclusionState {
	if glbae.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", glbae.Type(), glbae.LocalName()))
	}
	return glbae.state
}

// Args contains the configurations for google_logging_billing_account_exclusion.
type Args struct {
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

type googleLoggingBillingAccountExclusionAttributes struct {
	ref terra.Reference
}

// BillingAccount returns a reference to field billing_account of google_logging_billing_account_exclusion.
func (glbae googleLoggingBillingAccountExclusionAttributes) BillingAccount() terra.StringValue {
	return terra.ReferenceAsString(glbae.ref.Append("billing_account"))
}

// Description returns a reference to field description of google_logging_billing_account_exclusion.
func (glbae googleLoggingBillingAccountExclusionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(glbae.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_logging_billing_account_exclusion.
func (glbae googleLoggingBillingAccountExclusionAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(glbae.ref.Append("disabled"))
}

// Filter returns a reference to field filter of google_logging_billing_account_exclusion.
func (glbae googleLoggingBillingAccountExclusionAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(glbae.ref.Append("filter"))
}

// Id returns a reference to field id of google_logging_billing_account_exclusion.
func (glbae googleLoggingBillingAccountExclusionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(glbae.ref.Append("id"))
}

// Name returns a reference to field name of google_logging_billing_account_exclusion.
func (glbae googleLoggingBillingAccountExclusionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(glbae.ref.Append("name"))
}

type googleLoggingBillingAccountExclusionState struct {
	BillingAccount string `json:"billing_account"`
	Description    string `json:"description"`
	Disabled       bool   `json:"disabled"`
	Filter         string `json:"filter"`
	Id             string `json:"id"`
	Name           string `json:"name"`
}
