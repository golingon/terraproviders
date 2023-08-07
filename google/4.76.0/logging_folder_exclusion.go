// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoggingFolderExclusion creates a new instance of [LoggingFolderExclusion].
func NewLoggingFolderExclusion(name string, args LoggingFolderExclusionArgs) *LoggingFolderExclusion {
	return &LoggingFolderExclusion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoggingFolderExclusion)(nil)

// LoggingFolderExclusion represents the Terraform resource google_logging_folder_exclusion.
type LoggingFolderExclusion struct {
	Name      string
	Args      LoggingFolderExclusionArgs
	state     *loggingFolderExclusionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoggingFolderExclusion].
func (lfe *LoggingFolderExclusion) Type() string {
	return "google_logging_folder_exclusion"
}

// LocalName returns the local name for [LoggingFolderExclusion].
func (lfe *LoggingFolderExclusion) LocalName() string {
	return lfe.Name
}

// Configuration returns the configuration (args) for [LoggingFolderExclusion].
func (lfe *LoggingFolderExclusion) Configuration() interface{} {
	return lfe.Args
}

// DependOn is used for other resources to depend on [LoggingFolderExclusion].
func (lfe *LoggingFolderExclusion) DependOn() terra.Reference {
	return terra.ReferenceResource(lfe)
}

// Dependencies returns the list of resources [LoggingFolderExclusion] depends_on.
func (lfe *LoggingFolderExclusion) Dependencies() terra.Dependencies {
	return lfe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoggingFolderExclusion].
func (lfe *LoggingFolderExclusion) LifecycleManagement() *terra.Lifecycle {
	return lfe.Lifecycle
}

// Attributes returns the attributes for [LoggingFolderExclusion].
func (lfe *LoggingFolderExclusion) Attributes() loggingFolderExclusionAttributes {
	return loggingFolderExclusionAttributes{ref: terra.ReferenceResource(lfe)}
}

// ImportState imports the given attribute values into [LoggingFolderExclusion]'s state.
func (lfe *LoggingFolderExclusion) ImportState(av io.Reader) error {
	lfe.state = &loggingFolderExclusionState{}
	if err := json.NewDecoder(av).Decode(lfe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lfe.Type(), lfe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoggingFolderExclusion] has state.
func (lfe *LoggingFolderExclusion) State() (*loggingFolderExclusionState, bool) {
	return lfe.state, lfe.state != nil
}

// StateMust returns the state for [LoggingFolderExclusion]. Panics if the state is nil.
func (lfe *LoggingFolderExclusion) StateMust() *loggingFolderExclusionState {
	if lfe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lfe.Type(), lfe.LocalName()))
	}
	return lfe.state
}

// LoggingFolderExclusionArgs contains the configurations for google_logging_folder_exclusion.
type LoggingFolderExclusionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Filter: string, required
	Filter terra.StringValue `hcl:"filter,attr" validate:"required"`
	// Folder: string, required
	Folder terra.StringValue `hcl:"folder,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type loggingFolderExclusionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_logging_folder_exclusion.
func (lfe loggingFolderExclusionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lfe.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_logging_folder_exclusion.
func (lfe loggingFolderExclusionAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfe.ref.Append("disabled"))
}

// Filter returns a reference to field filter of google_logging_folder_exclusion.
func (lfe loggingFolderExclusionAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(lfe.ref.Append("filter"))
}

// Folder returns a reference to field folder of google_logging_folder_exclusion.
func (lfe loggingFolderExclusionAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(lfe.ref.Append("folder"))
}

// Id returns a reference to field id of google_logging_folder_exclusion.
func (lfe loggingFolderExclusionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lfe.ref.Append("id"))
}

// Name returns a reference to field name of google_logging_folder_exclusion.
func (lfe loggingFolderExclusionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lfe.ref.Append("name"))
}

type loggingFolderExclusionState struct {
	Description string `json:"description"`
	Disabled    bool   `json:"disabled"`
	Filter      string `json:"filter"`
	Folder      string `json:"folder"`
	Id          string `json:"id"`
	Name        string `json:"name"`
}
