// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dialogflowcxenvironment "github.com/golingon/terraproviders/google/4.63.1/dialogflowcxenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowCxEnvironment creates a new instance of [DialogflowCxEnvironment].
func NewDialogflowCxEnvironment(name string, args DialogflowCxEnvironmentArgs) *DialogflowCxEnvironment {
	return &DialogflowCxEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowCxEnvironment)(nil)

// DialogflowCxEnvironment represents the Terraform resource google_dialogflow_cx_environment.
type DialogflowCxEnvironment struct {
	Name      string
	Args      DialogflowCxEnvironmentArgs
	state     *dialogflowCxEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowCxEnvironment].
func (dce *DialogflowCxEnvironment) Type() string {
	return "google_dialogflow_cx_environment"
}

// LocalName returns the local name for [DialogflowCxEnvironment].
func (dce *DialogflowCxEnvironment) LocalName() string {
	return dce.Name
}

// Configuration returns the configuration (args) for [DialogflowCxEnvironment].
func (dce *DialogflowCxEnvironment) Configuration() interface{} {
	return dce.Args
}

// DependOn is used for other resources to depend on [DialogflowCxEnvironment].
func (dce *DialogflowCxEnvironment) DependOn() terra.Reference {
	return terra.ReferenceResource(dce)
}

// Dependencies returns the list of resources [DialogflowCxEnvironment] depends_on.
func (dce *DialogflowCxEnvironment) Dependencies() terra.Dependencies {
	return dce.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowCxEnvironment].
func (dce *DialogflowCxEnvironment) LifecycleManagement() *terra.Lifecycle {
	return dce.Lifecycle
}

// Attributes returns the attributes for [DialogflowCxEnvironment].
func (dce *DialogflowCxEnvironment) Attributes() dialogflowCxEnvironmentAttributes {
	return dialogflowCxEnvironmentAttributes{ref: terra.ReferenceResource(dce)}
}

// ImportState imports the given attribute values into [DialogflowCxEnvironment]'s state.
func (dce *DialogflowCxEnvironment) ImportState(av io.Reader) error {
	dce.state = &dialogflowCxEnvironmentState{}
	if err := json.NewDecoder(av).Decode(dce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dce.Type(), dce.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowCxEnvironment] has state.
func (dce *DialogflowCxEnvironment) State() (*dialogflowCxEnvironmentState, bool) {
	return dce.state, dce.state != nil
}

// StateMust returns the state for [DialogflowCxEnvironment]. Panics if the state is nil.
func (dce *DialogflowCxEnvironment) StateMust() *dialogflowCxEnvironmentState {
	if dce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dce.Type(), dce.LocalName()))
	}
	return dce.state
}

// DialogflowCxEnvironmentArgs contains the configurations for google_dialogflow_cx_environment.
type DialogflowCxEnvironmentArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, optional
	Parent terra.StringValue `hcl:"parent,attr"`
	// Timeouts: optional
	Timeouts *dialogflowcxenvironment.Timeouts `hcl:"timeouts,block"`
	// VersionConfigs: min=1
	VersionConfigs []dialogflowcxenvironment.VersionConfigs `hcl:"version_configs,block" validate:"min=1"`
}
type dialogflowCxEnvironmentAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_dialogflow_cx_environment.
func (dce dialogflowCxEnvironmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dialogflow_cx_environment.
func (dce dialogflowCxEnvironmentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dialogflow_cx_environment.
func (dce dialogflowCxEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("id"))
}

// Name returns a reference to field name of google_dialogflow_cx_environment.
func (dce dialogflowCxEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("name"))
}

// Parent returns a reference to field parent of google_dialogflow_cx_environment.
func (dce dialogflowCxEnvironmentAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("parent"))
}

// UpdateTime returns a reference to field update_time of google_dialogflow_cx_environment.
func (dce dialogflowCxEnvironmentAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("update_time"))
}

func (dce dialogflowCxEnvironmentAttributes) Timeouts() dialogflowcxenvironment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowcxenvironment.TimeoutsAttributes](dce.ref.Append("timeouts"))
}

func (dce dialogflowCxEnvironmentAttributes) VersionConfigs() terra.ListValue[dialogflowcxenvironment.VersionConfigsAttributes] {
	return terra.ReferenceAsList[dialogflowcxenvironment.VersionConfigsAttributes](dce.ref.Append("version_configs"))
}

type dialogflowCxEnvironmentState struct {
	Description    string                                        `json:"description"`
	DisplayName    string                                        `json:"display_name"`
	Id             string                                        `json:"id"`
	Name           string                                        `json:"name"`
	Parent         string                                        `json:"parent"`
	UpdateTime     string                                        `json:"update_time"`
	Timeouts       *dialogflowcxenvironment.TimeoutsState        `json:"timeouts"`
	VersionConfigs []dialogflowcxenvironment.VersionConfigsState `json:"version_configs"`
}
