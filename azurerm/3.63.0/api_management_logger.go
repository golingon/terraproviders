// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementlogger "github.com/golingon/terraproviders/azurerm/3.63.0/apimanagementlogger"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementLogger creates a new instance of [ApiManagementLogger].
func NewApiManagementLogger(name string, args ApiManagementLoggerArgs) *ApiManagementLogger {
	return &ApiManagementLogger{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementLogger)(nil)

// ApiManagementLogger represents the Terraform resource azurerm_api_management_logger.
type ApiManagementLogger struct {
	Name      string
	Args      ApiManagementLoggerArgs
	state     *apiManagementLoggerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementLogger].
func (aml *ApiManagementLogger) Type() string {
	return "azurerm_api_management_logger"
}

// LocalName returns the local name for [ApiManagementLogger].
func (aml *ApiManagementLogger) LocalName() string {
	return aml.Name
}

// Configuration returns the configuration (args) for [ApiManagementLogger].
func (aml *ApiManagementLogger) Configuration() interface{} {
	return aml.Args
}

// DependOn is used for other resources to depend on [ApiManagementLogger].
func (aml *ApiManagementLogger) DependOn() terra.Reference {
	return terra.ReferenceResource(aml)
}

// Dependencies returns the list of resources [ApiManagementLogger] depends_on.
func (aml *ApiManagementLogger) Dependencies() terra.Dependencies {
	return aml.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementLogger].
func (aml *ApiManagementLogger) LifecycleManagement() *terra.Lifecycle {
	return aml.Lifecycle
}

// Attributes returns the attributes for [ApiManagementLogger].
func (aml *ApiManagementLogger) Attributes() apiManagementLoggerAttributes {
	return apiManagementLoggerAttributes{ref: terra.ReferenceResource(aml)}
}

// ImportState imports the given attribute values into [ApiManagementLogger]'s state.
func (aml *ApiManagementLogger) ImportState(av io.Reader) error {
	aml.state = &apiManagementLoggerState{}
	if err := json.NewDecoder(av).Decode(aml.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aml.Type(), aml.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementLogger] has state.
func (aml *ApiManagementLogger) State() (*apiManagementLoggerState, bool) {
	return aml.state, aml.state != nil
}

// StateMust returns the state for [ApiManagementLogger]. Panics if the state is nil.
func (aml *ApiManagementLogger) StateMust() *apiManagementLoggerState {
	if aml.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aml.Type(), aml.LocalName()))
	}
	return aml.state
}

// ApiManagementLoggerArgs contains the configurations for azurerm_api_management_logger.
type ApiManagementLoggerArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Buffered: bool, optional
	Buffered terra.BoolValue `hcl:"buffered,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ResourceId: string, optional
	ResourceId terra.StringValue `hcl:"resource_id,attr"`
	// ApplicationInsights: optional
	ApplicationInsights *apimanagementlogger.ApplicationInsights `hcl:"application_insights,block"`
	// Eventhub: optional
	Eventhub *apimanagementlogger.Eventhub `hcl:"eventhub,block"`
	// Timeouts: optional
	Timeouts *apimanagementlogger.Timeouts `hcl:"timeouts,block"`
}
type apiManagementLoggerAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_logger.
func (aml apiManagementLoggerAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(aml.ref.Append("api_management_name"))
}

// Buffered returns a reference to field buffered of azurerm_api_management_logger.
func (aml apiManagementLoggerAttributes) Buffered() terra.BoolValue {
	return terra.ReferenceAsBool(aml.ref.Append("buffered"))
}

// Description returns a reference to field description of azurerm_api_management_logger.
func (aml apiManagementLoggerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aml.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_api_management_logger.
func (aml apiManagementLoggerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aml.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_logger.
func (aml apiManagementLoggerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aml.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_logger.
func (aml apiManagementLoggerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aml.ref.Append("resource_group_name"))
}

// ResourceId returns a reference to field resource_id of azurerm_api_management_logger.
func (aml apiManagementLoggerAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(aml.ref.Append("resource_id"))
}

func (aml apiManagementLoggerAttributes) ApplicationInsights() terra.ListValue[apimanagementlogger.ApplicationInsightsAttributes] {
	return terra.ReferenceAsList[apimanagementlogger.ApplicationInsightsAttributes](aml.ref.Append("application_insights"))
}

func (aml apiManagementLoggerAttributes) Eventhub() terra.ListValue[apimanagementlogger.EventhubAttributes] {
	return terra.ReferenceAsList[apimanagementlogger.EventhubAttributes](aml.ref.Append("eventhub"))
}

func (aml apiManagementLoggerAttributes) Timeouts() apimanagementlogger.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementlogger.TimeoutsAttributes](aml.ref.Append("timeouts"))
}

type apiManagementLoggerState struct {
	ApiManagementName   string                                         `json:"api_management_name"`
	Buffered            bool                                           `json:"buffered"`
	Description         string                                         `json:"description"`
	Id                  string                                         `json:"id"`
	Name                string                                         `json:"name"`
	ResourceGroupName   string                                         `json:"resource_group_name"`
	ResourceId          string                                         `json:"resource_id"`
	ApplicationInsights []apimanagementlogger.ApplicationInsightsState `json:"application_insights"`
	Eventhub            []apimanagementlogger.EventhubState            `json:"eventhub"`
	Timeouts            *apimanagementlogger.TimeoutsState             `json:"timeouts"`
}
