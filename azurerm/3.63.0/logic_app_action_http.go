// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicappactionhttp "github.com/golingon/terraproviders/azurerm/3.63.0/logicappactionhttp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogicAppActionHttp creates a new instance of [LogicAppActionHttp].
func NewLogicAppActionHttp(name string, args LogicAppActionHttpArgs) *LogicAppActionHttp {
	return &LogicAppActionHttp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppActionHttp)(nil)

// LogicAppActionHttp represents the Terraform resource azurerm_logic_app_action_http.
type LogicAppActionHttp struct {
	Name      string
	Args      LogicAppActionHttpArgs
	state     *logicAppActionHttpState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogicAppActionHttp].
func (laah *LogicAppActionHttp) Type() string {
	return "azurerm_logic_app_action_http"
}

// LocalName returns the local name for [LogicAppActionHttp].
func (laah *LogicAppActionHttp) LocalName() string {
	return laah.Name
}

// Configuration returns the configuration (args) for [LogicAppActionHttp].
func (laah *LogicAppActionHttp) Configuration() interface{} {
	return laah.Args
}

// DependOn is used for other resources to depend on [LogicAppActionHttp].
func (laah *LogicAppActionHttp) DependOn() terra.Reference {
	return terra.ReferenceResource(laah)
}

// Dependencies returns the list of resources [LogicAppActionHttp] depends_on.
func (laah *LogicAppActionHttp) Dependencies() terra.Dependencies {
	return laah.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogicAppActionHttp].
func (laah *LogicAppActionHttp) LifecycleManagement() *terra.Lifecycle {
	return laah.Lifecycle
}

// Attributes returns the attributes for [LogicAppActionHttp].
func (laah *LogicAppActionHttp) Attributes() logicAppActionHttpAttributes {
	return logicAppActionHttpAttributes{ref: terra.ReferenceResource(laah)}
}

// ImportState imports the given attribute values into [LogicAppActionHttp]'s state.
func (laah *LogicAppActionHttp) ImportState(av io.Reader) error {
	laah.state = &logicAppActionHttpState{}
	if err := json.NewDecoder(av).Decode(laah.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", laah.Type(), laah.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogicAppActionHttp] has state.
func (laah *LogicAppActionHttp) State() (*logicAppActionHttpState, bool) {
	return laah.state, laah.state != nil
}

// StateMust returns the state for [LogicAppActionHttp]. Panics if the state is nil.
func (laah *LogicAppActionHttp) StateMust() *logicAppActionHttpState {
	if laah.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", laah.Type(), laah.LocalName()))
	}
	return laah.state
}

// LogicAppActionHttpArgs contains the configurations for azurerm_logic_app_action_http.
type LogicAppActionHttpArgs struct {
	// Body: string, optional
	Body terra.StringValue `hcl:"body,attr"`
	// Headers: map of string, optional
	Headers terra.MapValue[terra.StringValue] `hcl:"headers,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogicAppId: string, required
	LogicAppId terra.StringValue `hcl:"logic_app_id,attr" validate:"required"`
	// Method: string, required
	Method terra.StringValue `hcl:"method,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Queries: map of string, optional
	Queries terra.MapValue[terra.StringValue] `hcl:"queries,attr"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
	// RunAfter: min=0
	RunAfter []logicappactionhttp.RunAfter `hcl:"run_after,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *logicappactionhttp.Timeouts `hcl:"timeouts,block"`
}
type logicAppActionHttpAttributes struct {
	ref terra.Reference
}

// Body returns a reference to field body of azurerm_logic_app_action_http.
func (laah logicAppActionHttpAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(laah.ref.Append("body"))
}

// Headers returns a reference to field headers of azurerm_logic_app_action_http.
func (laah logicAppActionHttpAttributes) Headers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](laah.ref.Append("headers"))
}

// Id returns a reference to field id of azurerm_logic_app_action_http.
func (laah logicAppActionHttpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(laah.ref.Append("id"))
}

// LogicAppId returns a reference to field logic_app_id of azurerm_logic_app_action_http.
func (laah logicAppActionHttpAttributes) LogicAppId() terra.StringValue {
	return terra.ReferenceAsString(laah.ref.Append("logic_app_id"))
}

// Method returns a reference to field method of azurerm_logic_app_action_http.
func (laah logicAppActionHttpAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(laah.ref.Append("method"))
}

// Name returns a reference to field name of azurerm_logic_app_action_http.
func (laah logicAppActionHttpAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(laah.ref.Append("name"))
}

// Queries returns a reference to field queries of azurerm_logic_app_action_http.
func (laah logicAppActionHttpAttributes) Queries() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](laah.ref.Append("queries"))
}

// Uri returns a reference to field uri of azurerm_logic_app_action_http.
func (laah logicAppActionHttpAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(laah.ref.Append("uri"))
}

func (laah logicAppActionHttpAttributes) RunAfter() terra.SetValue[logicappactionhttp.RunAfterAttributes] {
	return terra.ReferenceAsSet[logicappactionhttp.RunAfterAttributes](laah.ref.Append("run_after"))
}

func (laah logicAppActionHttpAttributes) Timeouts() logicappactionhttp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logicappactionhttp.TimeoutsAttributes](laah.ref.Append("timeouts"))
}

type logicAppActionHttpState struct {
	Body       string                             `json:"body"`
	Headers    map[string]string                  `json:"headers"`
	Id         string                             `json:"id"`
	LogicAppId string                             `json:"logic_app_id"`
	Method     string                             `json:"method"`
	Name       string                             `json:"name"`
	Queries    map[string]string                  `json:"queries"`
	Uri        string                             `json:"uri"`
	RunAfter   []logicappactionhttp.RunAfterState `json:"run_after"`
	Timeouts   *logicappactionhttp.TimeoutsState  `json:"timeouts"`
}
