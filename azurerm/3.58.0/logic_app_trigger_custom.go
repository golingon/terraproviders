// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicapptriggercustom "github.com/golingon/terraproviders/azurerm/3.58.0/logicapptriggercustom"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogicAppTriggerCustom creates a new instance of [LogicAppTriggerCustom].
func NewLogicAppTriggerCustom(name string, args LogicAppTriggerCustomArgs) *LogicAppTriggerCustom {
	return &LogicAppTriggerCustom{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppTriggerCustom)(nil)

// LogicAppTriggerCustom represents the Terraform resource azurerm_logic_app_trigger_custom.
type LogicAppTriggerCustom struct {
	Name      string
	Args      LogicAppTriggerCustomArgs
	state     *logicAppTriggerCustomState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogicAppTriggerCustom].
func (latc *LogicAppTriggerCustom) Type() string {
	return "azurerm_logic_app_trigger_custom"
}

// LocalName returns the local name for [LogicAppTriggerCustom].
func (latc *LogicAppTriggerCustom) LocalName() string {
	return latc.Name
}

// Configuration returns the configuration (args) for [LogicAppTriggerCustom].
func (latc *LogicAppTriggerCustom) Configuration() interface{} {
	return latc.Args
}

// DependOn is used for other resources to depend on [LogicAppTriggerCustom].
func (latc *LogicAppTriggerCustom) DependOn() terra.Reference {
	return terra.ReferenceResource(latc)
}

// Dependencies returns the list of resources [LogicAppTriggerCustom] depends_on.
func (latc *LogicAppTriggerCustom) Dependencies() terra.Dependencies {
	return latc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogicAppTriggerCustom].
func (latc *LogicAppTriggerCustom) LifecycleManagement() *terra.Lifecycle {
	return latc.Lifecycle
}

// Attributes returns the attributes for [LogicAppTriggerCustom].
func (latc *LogicAppTriggerCustom) Attributes() logicAppTriggerCustomAttributes {
	return logicAppTriggerCustomAttributes{ref: terra.ReferenceResource(latc)}
}

// ImportState imports the given attribute values into [LogicAppTriggerCustom]'s state.
func (latc *LogicAppTriggerCustom) ImportState(av io.Reader) error {
	latc.state = &logicAppTriggerCustomState{}
	if err := json.NewDecoder(av).Decode(latc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", latc.Type(), latc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogicAppTriggerCustom] has state.
func (latc *LogicAppTriggerCustom) State() (*logicAppTriggerCustomState, bool) {
	return latc.state, latc.state != nil
}

// StateMust returns the state for [LogicAppTriggerCustom]. Panics if the state is nil.
func (latc *LogicAppTriggerCustom) StateMust() *logicAppTriggerCustomState {
	if latc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", latc.Type(), latc.LocalName()))
	}
	return latc.state
}

// LogicAppTriggerCustomArgs contains the configurations for azurerm_logic_app_trigger_custom.
type LogicAppTriggerCustomArgs struct {
	// Body: string, required
	Body terra.StringValue `hcl:"body,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogicAppId: string, required
	LogicAppId terra.StringValue `hcl:"logic_app_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *logicapptriggercustom.Timeouts `hcl:"timeouts,block"`
}
type logicAppTriggerCustomAttributes struct {
	ref terra.Reference
}

// Body returns a reference to field body of azurerm_logic_app_trigger_custom.
func (latc logicAppTriggerCustomAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(latc.ref.Append("body"))
}

// Id returns a reference to field id of azurerm_logic_app_trigger_custom.
func (latc logicAppTriggerCustomAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(latc.ref.Append("id"))
}

// LogicAppId returns a reference to field logic_app_id of azurerm_logic_app_trigger_custom.
func (latc logicAppTriggerCustomAttributes) LogicAppId() terra.StringValue {
	return terra.ReferenceAsString(latc.ref.Append("logic_app_id"))
}

// Name returns a reference to field name of azurerm_logic_app_trigger_custom.
func (latc logicAppTriggerCustomAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(latc.ref.Append("name"))
}

func (latc logicAppTriggerCustomAttributes) Timeouts() logicapptriggercustom.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logicapptriggercustom.TimeoutsAttributes](latc.ref.Append("timeouts"))
}

type logicAppTriggerCustomState struct {
	Body       string                               `json:"body"`
	Id         string                               `json:"id"`
	LogicAppId string                               `json:"logic_app_id"`
	Name       string                               `json:"name"`
	Timeouts   *logicapptriggercustom.TimeoutsState `json:"timeouts"`
}
