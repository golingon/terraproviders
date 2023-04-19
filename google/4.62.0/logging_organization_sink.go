// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	loggingorganizationsink "github.com/golingon/terraproviders/google/4.62.0/loggingorganizationsink"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingOrganizationSink creates a new instance of [LoggingOrganizationSink].
func NewLoggingOrganizationSink(name string, args LoggingOrganizationSinkArgs) *LoggingOrganizationSink {
	return &LoggingOrganizationSink{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingOrganizationSink)(nil)

// LoggingOrganizationSink represents the Terraform resource google_logging_organization_sink.
type LoggingOrganizationSink struct {
	Name      string
	Args      LoggingOrganizationSinkArgs
	state     *loggingOrganizationSinkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingOrganizationSink].
func (los *LoggingOrganizationSink) Type() string {
	return "google_logging_organization_sink"
}

// LocalName returns the local name for [LoggingOrganizationSink].
func (los *LoggingOrganizationSink) LocalName() string {
	return los.Name
}

// Configuration returns the configuration (args) for [LoggingOrganizationSink].
func (los *LoggingOrganizationSink) Configuration() interface{} {
	return los.Args
}

// DependOn is used for other resources to depend on [LoggingOrganizationSink].
func (los *LoggingOrganizationSink) DependOn() terra.Reference {
	return terra.ReferenceResource(los)
}

// Dependencies returns the list of resources [LoggingOrganizationSink] depends_on.
func (los *LoggingOrganizationSink) Dependencies() terra.Dependencies {
	return los.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingOrganizationSink].
func (los *LoggingOrganizationSink) LifecycleManagement() *terra.Lifecycle {
	return los.Lifecycle
}

// Attributes returns the attributes for [LoggingOrganizationSink].
func (los *LoggingOrganizationSink) Attributes() loggingOrganizationSinkAttributes {
	return loggingOrganizationSinkAttributes{ref: terra.ReferenceResource(los)}
}

// ImportState imports the given attribute values into [LoggingOrganizationSink]'s state.
func (los *LoggingOrganizationSink) ImportState(av io.Reader) error {
	los.state = &loggingOrganizationSinkState{}
	if err := json.NewDecoder(av).Decode(los.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", los.Type(), los.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingOrganizationSink] has state.
func (los *LoggingOrganizationSink) State() (*loggingOrganizationSinkState, bool) {
	return los.state, los.state != nil
}

// StateMust returns the state for [LoggingOrganizationSink]. Panics if the state is nil.
func (los *LoggingOrganizationSink) StateMust() *loggingOrganizationSinkState {
	if los.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", los.Type(), los.LocalName()))
	}
	return los.state
}

// LoggingOrganizationSinkArgs contains the configurations for google_logging_organization_sink.
type LoggingOrganizationSinkArgs struct {
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
	// IncludeChildren: bool, optional
	IncludeChildren terra.BoolValue `hcl:"include_children,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// BigqueryOptions: optional
	BigqueryOptions *loggingorganizationsink.BigqueryOptions `hcl:"bigquery_options,block"`
	// Exclusions: min=0
	Exclusions []loggingorganizationsink.Exclusions `hcl:"exclusions,block" validate:"min=0"`
}
type loggingOrganizationSinkAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_logging_organization_sink.
func (los loggingOrganizationSinkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(los.ref.Append("description"))
}

// Destination returns a reference to field destination of google_logging_organization_sink.
func (los loggingOrganizationSinkAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(los.ref.Append("destination"))
}

// Disabled returns a reference to field disabled of google_logging_organization_sink.
func (los loggingOrganizationSinkAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(los.ref.Append("disabled"))
}

// Filter returns a reference to field filter of google_logging_organization_sink.
func (los loggingOrganizationSinkAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(los.ref.Append("filter"))
}

// Id returns a reference to field id of google_logging_organization_sink.
func (los loggingOrganizationSinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(los.ref.Append("id"))
}

// IncludeChildren returns a reference to field include_children of google_logging_organization_sink.
func (los loggingOrganizationSinkAttributes) IncludeChildren() terra.BoolValue {
	return terra.ReferenceAsBool(los.ref.Append("include_children"))
}

// Name returns a reference to field name of google_logging_organization_sink.
func (los loggingOrganizationSinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(los.ref.Append("name"))
}

// OrgId returns a reference to field org_id of google_logging_organization_sink.
func (los loggingOrganizationSinkAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(los.ref.Append("org_id"))
}

// WriterIdentity returns a reference to field writer_identity of google_logging_organization_sink.
func (los loggingOrganizationSinkAttributes) WriterIdentity() terra.StringValue {
	return terra.ReferenceAsString(los.ref.Append("writer_identity"))
}

func (los loggingOrganizationSinkAttributes) BigqueryOptions() terra.ListValue[loggingorganizationsink.BigqueryOptionsAttributes] {
	return terra.ReferenceAsList[loggingorganizationsink.BigqueryOptionsAttributes](los.ref.Append("bigquery_options"))
}

func (los loggingOrganizationSinkAttributes) Exclusions() terra.ListValue[loggingorganizationsink.ExclusionsAttributes] {
	return terra.ReferenceAsList[loggingorganizationsink.ExclusionsAttributes](los.ref.Append("exclusions"))
}

type loggingOrganizationSinkState struct {
	Description     string                                         `json:"description"`
	Destination     string                                         `json:"destination"`
	Disabled        bool                                           `json:"disabled"`
	Filter          string                                         `json:"filter"`
	Id              string                                         `json:"id"`
	IncludeChildren bool                                           `json:"include_children"`
	Name            string                                         `json:"name"`
	OrgId           string                                         `json:"org_id"`
	WriterIdentity  string                                         `json:"writer_identity"`
	BigqueryOptions []loggingorganizationsink.BigqueryOptionsState `json:"bigquery_options"`
	Exclusions      []loggingorganizationsink.ExclusionsState      `json:"exclusions"`
}
