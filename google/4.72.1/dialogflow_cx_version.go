// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dialogflowcxversion "github.com/golingon/terraproviders/google/4.72.1/dialogflowcxversion"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowCxVersion creates a new instance of [DialogflowCxVersion].
func NewDialogflowCxVersion(name string, args DialogflowCxVersionArgs) *DialogflowCxVersion {
	return &DialogflowCxVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowCxVersion)(nil)

// DialogflowCxVersion represents the Terraform resource google_dialogflow_cx_version.
type DialogflowCxVersion struct {
	Name      string
	Args      DialogflowCxVersionArgs
	state     *dialogflowCxVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowCxVersion].
func (dcv *DialogflowCxVersion) Type() string {
	return "google_dialogflow_cx_version"
}

// LocalName returns the local name for [DialogflowCxVersion].
func (dcv *DialogflowCxVersion) LocalName() string {
	return dcv.Name
}

// Configuration returns the configuration (args) for [DialogflowCxVersion].
func (dcv *DialogflowCxVersion) Configuration() interface{} {
	return dcv.Args
}

// DependOn is used for other resources to depend on [DialogflowCxVersion].
func (dcv *DialogflowCxVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(dcv)
}

// Dependencies returns the list of resources [DialogflowCxVersion] depends_on.
func (dcv *DialogflowCxVersion) Dependencies() terra.Dependencies {
	return dcv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowCxVersion].
func (dcv *DialogflowCxVersion) LifecycleManagement() *terra.Lifecycle {
	return dcv.Lifecycle
}

// Attributes returns the attributes for [DialogflowCxVersion].
func (dcv *DialogflowCxVersion) Attributes() dialogflowCxVersionAttributes {
	return dialogflowCxVersionAttributes{ref: terra.ReferenceResource(dcv)}
}

// ImportState imports the given attribute values into [DialogflowCxVersion]'s state.
func (dcv *DialogflowCxVersion) ImportState(av io.Reader) error {
	dcv.state = &dialogflowCxVersionState{}
	if err := json.NewDecoder(av).Decode(dcv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcv.Type(), dcv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowCxVersion] has state.
func (dcv *DialogflowCxVersion) State() (*dialogflowCxVersionState, bool) {
	return dcv.state, dcv.state != nil
}

// StateMust returns the state for [DialogflowCxVersion]. Panics if the state is nil.
func (dcv *DialogflowCxVersion) StateMust() *dialogflowCxVersionState {
	if dcv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcv.Type(), dcv.LocalName()))
	}
	return dcv.state
}

// DialogflowCxVersionArgs contains the configurations for google_dialogflow_cx_version.
type DialogflowCxVersionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, optional
	Parent terra.StringValue `hcl:"parent,attr"`
	// NluSettings: min=0
	NluSettings []dialogflowcxversion.NluSettings `hcl:"nlu_settings,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dialogflowcxversion.Timeouts `hcl:"timeouts,block"`
}
type dialogflowCxVersionAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_dialogflow_cx_version.
func (dcv dialogflowCxVersionAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dcv.ref.Append("create_time"))
}

// Description returns a reference to field description of google_dialogflow_cx_version.
func (dcv dialogflowCxVersionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dcv.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_dialogflow_cx_version.
func (dcv dialogflowCxVersionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dcv.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dialogflow_cx_version.
func (dcv dialogflowCxVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcv.ref.Append("id"))
}

// Name returns a reference to field name of google_dialogflow_cx_version.
func (dcv dialogflowCxVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcv.ref.Append("name"))
}

// Parent returns a reference to field parent of google_dialogflow_cx_version.
func (dcv dialogflowCxVersionAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dcv.ref.Append("parent"))
}

// State returns a reference to field state of google_dialogflow_cx_version.
func (dcv dialogflowCxVersionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dcv.ref.Append("state"))
}

func (dcv dialogflowCxVersionAttributes) NluSettings() terra.ListValue[dialogflowcxversion.NluSettingsAttributes] {
	return terra.ReferenceAsList[dialogflowcxversion.NluSettingsAttributes](dcv.ref.Append("nlu_settings"))
}

func (dcv dialogflowCxVersionAttributes) Timeouts() dialogflowcxversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowcxversion.TimeoutsAttributes](dcv.ref.Append("timeouts"))
}

type dialogflowCxVersionState struct {
	CreateTime  string                                 `json:"create_time"`
	Description string                                 `json:"description"`
	DisplayName string                                 `json:"display_name"`
	Id          string                                 `json:"id"`
	Name        string                                 `json:"name"`
	Parent      string                                 `json:"parent"`
	State       string                                 `json:"state"`
	NluSettings []dialogflowcxversion.NluSettingsState `json:"nlu_settings"`
	Timeouts    *dialogflowcxversion.TimeoutsState     `json:"timeouts"`
}
