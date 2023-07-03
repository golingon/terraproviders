// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicapptriggerhttprequest "github.com/golingon/terraproviders/azurerm/3.63.0/logicapptriggerhttprequest"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogicAppTriggerHttpRequest creates a new instance of [LogicAppTriggerHttpRequest].
func NewLogicAppTriggerHttpRequest(name string, args LogicAppTriggerHttpRequestArgs) *LogicAppTriggerHttpRequest {
	return &LogicAppTriggerHttpRequest{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppTriggerHttpRequest)(nil)

// LogicAppTriggerHttpRequest represents the Terraform resource azurerm_logic_app_trigger_http_request.
type LogicAppTriggerHttpRequest struct {
	Name      string
	Args      LogicAppTriggerHttpRequestArgs
	state     *logicAppTriggerHttpRequestState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogicAppTriggerHttpRequest].
func (lathr *LogicAppTriggerHttpRequest) Type() string {
	return "azurerm_logic_app_trigger_http_request"
}

// LocalName returns the local name for [LogicAppTriggerHttpRequest].
func (lathr *LogicAppTriggerHttpRequest) LocalName() string {
	return lathr.Name
}

// Configuration returns the configuration (args) for [LogicAppTriggerHttpRequest].
func (lathr *LogicAppTriggerHttpRequest) Configuration() interface{} {
	return lathr.Args
}

// DependOn is used for other resources to depend on [LogicAppTriggerHttpRequest].
func (lathr *LogicAppTriggerHttpRequest) DependOn() terra.Reference {
	return terra.ReferenceResource(lathr)
}

// Dependencies returns the list of resources [LogicAppTriggerHttpRequest] depends_on.
func (lathr *LogicAppTriggerHttpRequest) Dependencies() terra.Dependencies {
	return lathr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogicAppTriggerHttpRequest].
func (lathr *LogicAppTriggerHttpRequest) LifecycleManagement() *terra.Lifecycle {
	return lathr.Lifecycle
}

// Attributes returns the attributes for [LogicAppTriggerHttpRequest].
func (lathr *LogicAppTriggerHttpRequest) Attributes() logicAppTriggerHttpRequestAttributes {
	return logicAppTriggerHttpRequestAttributes{ref: terra.ReferenceResource(lathr)}
}

// ImportState imports the given attribute values into [LogicAppTriggerHttpRequest]'s state.
func (lathr *LogicAppTriggerHttpRequest) ImportState(av io.Reader) error {
	lathr.state = &logicAppTriggerHttpRequestState{}
	if err := json.NewDecoder(av).Decode(lathr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lathr.Type(), lathr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogicAppTriggerHttpRequest] has state.
func (lathr *LogicAppTriggerHttpRequest) State() (*logicAppTriggerHttpRequestState, bool) {
	return lathr.state, lathr.state != nil
}

// StateMust returns the state for [LogicAppTriggerHttpRequest]. Panics if the state is nil.
func (lathr *LogicAppTriggerHttpRequest) StateMust() *logicAppTriggerHttpRequestState {
	if lathr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lathr.Type(), lathr.LocalName()))
	}
	return lathr.state
}

// LogicAppTriggerHttpRequestArgs contains the configurations for azurerm_logic_app_trigger_http_request.
type LogicAppTriggerHttpRequestArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogicAppId: string, required
	LogicAppId terra.StringValue `hcl:"logic_app_id,attr" validate:"required"`
	// Method: string, optional
	Method terra.StringValue `hcl:"method,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RelativePath: string, optional
	RelativePath terra.StringValue `hcl:"relative_path,attr"`
	// Schema: string, required
	Schema terra.StringValue `hcl:"schema,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *logicapptriggerhttprequest.Timeouts `hcl:"timeouts,block"`
}
type logicAppTriggerHttpRequestAttributes struct {
	ref terra.Reference
}

// CallbackUrl returns a reference to field callback_url of azurerm_logic_app_trigger_http_request.
func (lathr logicAppTriggerHttpRequestAttributes) CallbackUrl() terra.StringValue {
	return terra.ReferenceAsString(lathr.ref.Append("callback_url"))
}

// Id returns a reference to field id of azurerm_logic_app_trigger_http_request.
func (lathr logicAppTriggerHttpRequestAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lathr.ref.Append("id"))
}

// LogicAppId returns a reference to field logic_app_id of azurerm_logic_app_trigger_http_request.
func (lathr logicAppTriggerHttpRequestAttributes) LogicAppId() terra.StringValue {
	return terra.ReferenceAsString(lathr.ref.Append("logic_app_id"))
}

// Method returns a reference to field method of azurerm_logic_app_trigger_http_request.
func (lathr logicAppTriggerHttpRequestAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(lathr.ref.Append("method"))
}

// Name returns a reference to field name of azurerm_logic_app_trigger_http_request.
func (lathr logicAppTriggerHttpRequestAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lathr.ref.Append("name"))
}

// RelativePath returns a reference to field relative_path of azurerm_logic_app_trigger_http_request.
func (lathr logicAppTriggerHttpRequestAttributes) RelativePath() terra.StringValue {
	return terra.ReferenceAsString(lathr.ref.Append("relative_path"))
}

// Schema returns a reference to field schema of azurerm_logic_app_trigger_http_request.
func (lathr logicAppTriggerHttpRequestAttributes) Schema() terra.StringValue {
	return terra.ReferenceAsString(lathr.ref.Append("schema"))
}

func (lathr logicAppTriggerHttpRequestAttributes) Timeouts() logicapptriggerhttprequest.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logicapptriggerhttprequest.TimeoutsAttributes](lathr.ref.Append("timeouts"))
}

type logicAppTriggerHttpRequestState struct {
	CallbackUrl  string                                    `json:"callback_url"`
	Id           string                                    `json:"id"`
	LogicAppId   string                                    `json:"logic_app_id"`
	Method       string                                    `json:"method"`
	Name         string                                    `json:"name"`
	RelativePath string                                    `json:"relative_path"`
	Schema       string                                    `json:"schema"`
	Timeouts     *logicapptriggerhttprequest.TimeoutsState `json:"timeouts"`
}
