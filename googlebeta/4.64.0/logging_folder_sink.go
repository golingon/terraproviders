// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	loggingfoldersink "github.com/golingon/terraproviders/googlebeta/4.64.0/loggingfoldersink"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingFolderSink creates a new instance of [LoggingFolderSink].
func NewLoggingFolderSink(name string, args LoggingFolderSinkArgs) *LoggingFolderSink {
	return &LoggingFolderSink{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingFolderSink)(nil)

// LoggingFolderSink represents the Terraform resource google_logging_folder_sink.
type LoggingFolderSink struct {
	Name      string
	Args      LoggingFolderSinkArgs
	state     *loggingFolderSinkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingFolderSink].
func (lfs *LoggingFolderSink) Type() string {
	return "google_logging_folder_sink"
}

// LocalName returns the local name for [LoggingFolderSink].
func (lfs *LoggingFolderSink) LocalName() string {
	return lfs.Name
}

// Configuration returns the configuration (args) for [LoggingFolderSink].
func (lfs *LoggingFolderSink) Configuration() interface{} {
	return lfs.Args
}

// DependOn is used for other resources to depend on [LoggingFolderSink].
func (lfs *LoggingFolderSink) DependOn() terra.Reference {
	return terra.ReferenceResource(lfs)
}

// Dependencies returns the list of resources [LoggingFolderSink] depends_on.
func (lfs *LoggingFolderSink) Dependencies() terra.Dependencies {
	return lfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingFolderSink].
func (lfs *LoggingFolderSink) LifecycleManagement() *terra.Lifecycle {
	return lfs.Lifecycle
}

// Attributes returns the attributes for [LoggingFolderSink].
func (lfs *LoggingFolderSink) Attributes() loggingFolderSinkAttributes {
	return loggingFolderSinkAttributes{ref: terra.ReferenceResource(lfs)}
}

// ImportState imports the given attribute values into [LoggingFolderSink]'s state.
func (lfs *LoggingFolderSink) ImportState(av io.Reader) error {
	lfs.state = &loggingFolderSinkState{}
	if err := json.NewDecoder(av).Decode(lfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lfs.Type(), lfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingFolderSink] has state.
func (lfs *LoggingFolderSink) State() (*loggingFolderSinkState, bool) {
	return lfs.state, lfs.state != nil
}

// StateMust returns the state for [LoggingFolderSink]. Panics if the state is nil.
func (lfs *LoggingFolderSink) StateMust() *loggingFolderSinkState {
	if lfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lfs.Type(), lfs.LocalName()))
	}
	return lfs.state
}

// LoggingFolderSinkArgs contains the configurations for google_logging_folder_sink.
type LoggingFolderSinkArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Destination: string, required
	Destination terra.StringValue `hcl:"destination,attr" validate:"required"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// Folder: string, required
	Folder terra.StringValue `hcl:"folder,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludeChildren: bool, optional
	IncludeChildren terra.BoolValue `hcl:"include_children,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// BigqueryOptions: optional
	BigqueryOptions *loggingfoldersink.BigqueryOptions `hcl:"bigquery_options,block"`
	// Exclusions: min=0
	Exclusions []loggingfoldersink.Exclusions `hcl:"exclusions,block" validate:"min=0"`
}
type loggingFolderSinkAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_logging_folder_sink.
func (lfs loggingFolderSinkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lfs.ref.Append("description"))
}

// Destination returns a reference to field destination of google_logging_folder_sink.
func (lfs loggingFolderSinkAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(lfs.ref.Append("destination"))
}

// Disabled returns a reference to field disabled of google_logging_folder_sink.
func (lfs loggingFolderSinkAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfs.ref.Append("disabled"))
}

// Filter returns a reference to field filter of google_logging_folder_sink.
func (lfs loggingFolderSinkAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(lfs.ref.Append("filter"))
}

// Folder returns a reference to field folder of google_logging_folder_sink.
func (lfs loggingFolderSinkAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(lfs.ref.Append("folder"))
}

// Id returns a reference to field id of google_logging_folder_sink.
func (lfs loggingFolderSinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lfs.ref.Append("id"))
}

// IncludeChildren returns a reference to field include_children of google_logging_folder_sink.
func (lfs loggingFolderSinkAttributes) IncludeChildren() terra.BoolValue {
	return terra.ReferenceAsBool(lfs.ref.Append("include_children"))
}

// Name returns a reference to field name of google_logging_folder_sink.
func (lfs loggingFolderSinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lfs.ref.Append("name"))
}

// WriterIdentity returns a reference to field writer_identity of google_logging_folder_sink.
func (lfs loggingFolderSinkAttributes) WriterIdentity() terra.StringValue {
	return terra.ReferenceAsString(lfs.ref.Append("writer_identity"))
}

func (lfs loggingFolderSinkAttributes) BigqueryOptions() terra.ListValue[loggingfoldersink.BigqueryOptionsAttributes] {
	return terra.ReferenceAsList[loggingfoldersink.BigqueryOptionsAttributes](lfs.ref.Append("bigquery_options"))
}

func (lfs loggingFolderSinkAttributes) Exclusions() terra.ListValue[loggingfoldersink.ExclusionsAttributes] {
	return terra.ReferenceAsList[loggingfoldersink.ExclusionsAttributes](lfs.ref.Append("exclusions"))
}

type loggingFolderSinkState struct {
	Description     string                                   `json:"description"`
	Destination     string                                   `json:"destination"`
	Disabled        bool                                     `json:"disabled"`
	Filter          string                                   `json:"filter"`
	Folder          string                                   `json:"folder"`
	Id              string                                   `json:"id"`
	IncludeChildren bool                                     `json:"include_children"`
	Name            string                                   `json:"name"`
	WriterIdentity  string                                   `json:"writer_identity"`
	BigqueryOptions []loggingfoldersink.BigqueryOptionsState `json:"bigquery_options"`
	Exclusions      []loggingfoldersink.ExclusionsState      `json:"exclusions"`
}
