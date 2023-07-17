// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	mlenginemodel "github.com/golingon/terraproviders/google/4.73.1/mlenginemodel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMlEngineModel creates a new instance of [MlEngineModel].
func NewMlEngineModel(name string, args MlEngineModelArgs) *MlEngineModel {
	return &MlEngineModel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MlEngineModel)(nil)

// MlEngineModel represents the Terraform resource google_ml_engine_model.
type MlEngineModel struct {
	Name      string
	Args      MlEngineModelArgs
	state     *mlEngineModelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MlEngineModel].
func (mem *MlEngineModel) Type() string {
	return "google_ml_engine_model"
}

// LocalName returns the local name for [MlEngineModel].
func (mem *MlEngineModel) LocalName() string {
	return mem.Name
}

// Configuration returns the configuration (args) for [MlEngineModel].
func (mem *MlEngineModel) Configuration() interface{} {
	return mem.Args
}

// DependOn is used for other resources to depend on [MlEngineModel].
func (mem *MlEngineModel) DependOn() terra.Reference {
	return terra.ReferenceResource(mem)
}

// Dependencies returns the list of resources [MlEngineModel] depends_on.
func (mem *MlEngineModel) Dependencies() terra.Dependencies {
	return mem.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MlEngineModel].
func (mem *MlEngineModel) LifecycleManagement() *terra.Lifecycle {
	return mem.Lifecycle
}

// Attributes returns the attributes for [MlEngineModel].
func (mem *MlEngineModel) Attributes() mlEngineModelAttributes {
	return mlEngineModelAttributes{ref: terra.ReferenceResource(mem)}
}

// ImportState imports the given attribute values into [MlEngineModel]'s state.
func (mem *MlEngineModel) ImportState(av io.Reader) error {
	mem.state = &mlEngineModelState{}
	if err := json.NewDecoder(av).Decode(mem.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mem.Type(), mem.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MlEngineModel] has state.
func (mem *MlEngineModel) State() (*mlEngineModelState, bool) {
	return mem.state, mem.state != nil
}

// StateMust returns the state for [MlEngineModel]. Panics if the state is nil.
func (mem *MlEngineModel) StateMust() *mlEngineModelState {
	if mem.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mem.Type(), mem.LocalName()))
	}
	return mem.state
}

// MlEngineModelArgs contains the configurations for google_ml_engine_model.
type MlEngineModelArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OnlinePredictionConsoleLogging: bool, optional
	OnlinePredictionConsoleLogging terra.BoolValue `hcl:"online_prediction_console_logging,attr"`
	// OnlinePredictionLogging: bool, optional
	OnlinePredictionLogging terra.BoolValue `hcl:"online_prediction_logging,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Regions: list of string, optional
	Regions terra.ListValue[terra.StringValue] `hcl:"regions,attr"`
	// DefaultVersion: optional
	DefaultVersion *mlenginemodel.DefaultVersion `hcl:"default_version,block"`
	// Timeouts: optional
	Timeouts *mlenginemodel.Timeouts `hcl:"timeouts,block"`
}
type mlEngineModelAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_ml_engine_model.
func (mem mlEngineModelAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mem.ref.Append("description"))
}

// Id returns a reference to field id of google_ml_engine_model.
func (mem mlEngineModelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mem.ref.Append("id"))
}

// Labels returns a reference to field labels of google_ml_engine_model.
func (mem mlEngineModelAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mem.ref.Append("labels"))
}

// Name returns a reference to field name of google_ml_engine_model.
func (mem mlEngineModelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mem.ref.Append("name"))
}

// OnlinePredictionConsoleLogging returns a reference to field online_prediction_console_logging of google_ml_engine_model.
func (mem mlEngineModelAttributes) OnlinePredictionConsoleLogging() terra.BoolValue {
	return terra.ReferenceAsBool(mem.ref.Append("online_prediction_console_logging"))
}

// OnlinePredictionLogging returns a reference to field online_prediction_logging of google_ml_engine_model.
func (mem mlEngineModelAttributes) OnlinePredictionLogging() terra.BoolValue {
	return terra.ReferenceAsBool(mem.ref.Append("online_prediction_logging"))
}

// Project returns a reference to field project of google_ml_engine_model.
func (mem mlEngineModelAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(mem.ref.Append("project"))
}

// Regions returns a reference to field regions of google_ml_engine_model.
func (mem mlEngineModelAttributes) Regions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mem.ref.Append("regions"))
}

func (mem mlEngineModelAttributes) DefaultVersion() terra.ListValue[mlenginemodel.DefaultVersionAttributes] {
	return terra.ReferenceAsList[mlenginemodel.DefaultVersionAttributes](mem.ref.Append("default_version"))
}

func (mem mlEngineModelAttributes) Timeouts() mlenginemodel.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mlenginemodel.TimeoutsAttributes](mem.ref.Append("timeouts"))
}

type mlEngineModelState struct {
	Description                    string                              `json:"description"`
	Id                             string                              `json:"id"`
	Labels                         map[string]string                   `json:"labels"`
	Name                           string                              `json:"name"`
	OnlinePredictionConsoleLogging bool                                `json:"online_prediction_console_logging"`
	OnlinePredictionLogging        bool                                `json:"online_prediction_logging"`
	Project                        string                              `json:"project"`
	Regions                        []string                            `json:"regions"`
	DefaultVersion                 []mlenginemodel.DefaultVersionState `json:"default_version"`
	Timeouts                       *mlenginemodel.TimeoutsState        `json:"timeouts"`
}
