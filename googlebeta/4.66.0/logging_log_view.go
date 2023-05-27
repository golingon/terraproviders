// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	logginglogview "github.com/golingon/terraproviders/googlebeta/4.66.0/logginglogview"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingLogView creates a new instance of [LoggingLogView].
func NewLoggingLogView(name string, args LoggingLogViewArgs) *LoggingLogView {
	return &LoggingLogView{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingLogView)(nil)

// LoggingLogView represents the Terraform resource google_logging_log_view.
type LoggingLogView struct {
	Name      string
	Args      LoggingLogViewArgs
	state     *loggingLogViewState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingLogView].
func (llv *LoggingLogView) Type() string {
	return "google_logging_log_view"
}

// LocalName returns the local name for [LoggingLogView].
func (llv *LoggingLogView) LocalName() string {
	return llv.Name
}

// Configuration returns the configuration (args) for [LoggingLogView].
func (llv *LoggingLogView) Configuration() interface{} {
	return llv.Args
}

// DependOn is used for other resources to depend on [LoggingLogView].
func (llv *LoggingLogView) DependOn() terra.Reference {
	return terra.ReferenceResource(llv)
}

// Dependencies returns the list of resources [LoggingLogView] depends_on.
func (llv *LoggingLogView) Dependencies() terra.Dependencies {
	return llv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingLogView].
func (llv *LoggingLogView) LifecycleManagement() *terra.Lifecycle {
	return llv.Lifecycle
}

// Attributes returns the attributes for [LoggingLogView].
func (llv *LoggingLogView) Attributes() loggingLogViewAttributes {
	return loggingLogViewAttributes{ref: terra.ReferenceResource(llv)}
}

// ImportState imports the given attribute values into [LoggingLogView]'s state.
func (llv *LoggingLogView) ImportState(av io.Reader) error {
	llv.state = &loggingLogViewState{}
	if err := json.NewDecoder(av).Decode(llv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", llv.Type(), llv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingLogView] has state.
func (llv *LoggingLogView) State() (*loggingLogViewState, bool) {
	return llv.state, llv.state != nil
}

// StateMust returns the state for [LoggingLogView]. Panics if the state is nil.
func (llv *LoggingLogView) StateMust() *loggingLogViewState {
	if llv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", llv.Type(), llv.LocalName()))
	}
	return llv.state
}

// LoggingLogViewArgs contains the configurations for google_logging_log_view.
type LoggingLogViewArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, optional
	Parent terra.StringValue `hcl:"parent,attr"`
	// Timeouts: optional
	Timeouts *logginglogview.Timeouts `hcl:"timeouts,block"`
}
type loggingLogViewAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_logging_log_view.
func (llv loggingLogViewAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("bucket"))
}

// CreateTime returns a reference to field create_time of google_logging_log_view.
func (llv loggingLogViewAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("create_time"))
}

// Description returns a reference to field description of google_logging_log_view.
func (llv loggingLogViewAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("description"))
}

// Filter returns a reference to field filter of google_logging_log_view.
func (llv loggingLogViewAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("filter"))
}

// Id returns a reference to field id of google_logging_log_view.
func (llv loggingLogViewAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("id"))
}

// Location returns a reference to field location of google_logging_log_view.
func (llv loggingLogViewAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("location"))
}

// Name returns a reference to field name of google_logging_log_view.
func (llv loggingLogViewAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("name"))
}

// Parent returns a reference to field parent of google_logging_log_view.
func (llv loggingLogViewAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("parent"))
}

// UpdateTime returns a reference to field update_time of google_logging_log_view.
func (llv loggingLogViewAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(llv.ref.Append("update_time"))
}

func (llv loggingLogViewAttributes) Timeouts() logginglogview.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logginglogview.TimeoutsAttributes](llv.ref.Append("timeouts"))
}

type loggingLogViewState struct {
	Bucket      string                        `json:"bucket"`
	CreateTime  string                        `json:"create_time"`
	Description string                        `json:"description"`
	Filter      string                        `json:"filter"`
	Id          string                        `json:"id"`
	Location    string                        `json:"location"`
	Name        string                        `json:"name"`
	Parent      string                        `json:"parent"`
	UpdateTime  string                        `json:"update_time"`
	Timeouts    *logginglogview.TimeoutsState `json:"timeouts"`
}
