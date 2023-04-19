// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementdiagnostic "github.com/golingon/terraproviders/azurerm/3.52.0/apimanagementdiagnostic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementDiagnostic creates a new instance of [ApiManagementDiagnostic].
func NewApiManagementDiagnostic(name string, args ApiManagementDiagnosticArgs) *ApiManagementDiagnostic {
	return &ApiManagementDiagnostic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementDiagnostic)(nil)

// ApiManagementDiagnostic represents the Terraform resource azurerm_api_management_diagnostic.
type ApiManagementDiagnostic struct {
	Name      string
	Args      ApiManagementDiagnosticArgs
	state     *apiManagementDiagnosticState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementDiagnostic].
func (amd *ApiManagementDiagnostic) Type() string {
	return "azurerm_api_management_diagnostic"
}

// LocalName returns the local name for [ApiManagementDiagnostic].
func (amd *ApiManagementDiagnostic) LocalName() string {
	return amd.Name
}

// Configuration returns the configuration (args) for [ApiManagementDiagnostic].
func (amd *ApiManagementDiagnostic) Configuration() interface{} {
	return amd.Args
}

// DependOn is used for other resources to depend on [ApiManagementDiagnostic].
func (amd *ApiManagementDiagnostic) DependOn() terra.Reference {
	return terra.ReferenceResource(amd)
}

// Dependencies returns the list of resources [ApiManagementDiagnostic] depends_on.
func (amd *ApiManagementDiagnostic) Dependencies() terra.Dependencies {
	return amd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementDiagnostic].
func (amd *ApiManagementDiagnostic) LifecycleManagement() *terra.Lifecycle {
	return amd.Lifecycle
}

// Attributes returns the attributes for [ApiManagementDiagnostic].
func (amd *ApiManagementDiagnostic) Attributes() apiManagementDiagnosticAttributes {
	return apiManagementDiagnosticAttributes{ref: terra.ReferenceResource(amd)}
}

// ImportState imports the given attribute values into [ApiManagementDiagnostic]'s state.
func (amd *ApiManagementDiagnostic) ImportState(av io.Reader) error {
	amd.state = &apiManagementDiagnosticState{}
	if err := json.NewDecoder(av).Decode(amd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amd.Type(), amd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementDiagnostic] has state.
func (amd *ApiManagementDiagnostic) State() (*apiManagementDiagnosticState, bool) {
	return amd.state, amd.state != nil
}

// StateMust returns the state for [ApiManagementDiagnostic]. Panics if the state is nil.
func (amd *ApiManagementDiagnostic) StateMust() *apiManagementDiagnosticState {
	if amd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amd.Type(), amd.LocalName()))
	}
	return amd.state
}

// ApiManagementDiagnosticArgs contains the configurations for azurerm_api_management_diagnostic.
type ApiManagementDiagnosticArgs struct {
	// AlwaysLogErrors: bool, optional
	AlwaysLogErrors terra.BoolValue `hcl:"always_log_errors,attr"`
	// ApiManagementLoggerId: string, required
	ApiManagementLoggerId terra.StringValue `hcl:"api_management_logger_id,attr" validate:"required"`
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// HttpCorrelationProtocol: string, optional
	HttpCorrelationProtocol terra.StringValue `hcl:"http_correlation_protocol,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Identifier: string, required
	Identifier terra.StringValue `hcl:"identifier,attr" validate:"required"`
	// LogClientIp: bool, optional
	LogClientIp terra.BoolValue `hcl:"log_client_ip,attr"`
	// OperationNameFormat: string, optional
	OperationNameFormat terra.StringValue `hcl:"operation_name_format,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SamplingPercentage: number, optional
	SamplingPercentage terra.NumberValue `hcl:"sampling_percentage,attr"`
	// Verbosity: string, optional
	Verbosity terra.StringValue `hcl:"verbosity,attr"`
	// BackendRequest: optional
	BackendRequest *apimanagementdiagnostic.BackendRequest `hcl:"backend_request,block"`
	// BackendResponse: optional
	BackendResponse *apimanagementdiagnostic.BackendResponse `hcl:"backend_response,block"`
	// FrontendRequest: optional
	FrontendRequest *apimanagementdiagnostic.FrontendRequest `hcl:"frontend_request,block"`
	// FrontendResponse: optional
	FrontendResponse *apimanagementdiagnostic.FrontendResponse `hcl:"frontend_response,block"`
	// Timeouts: optional
	Timeouts *apimanagementdiagnostic.Timeouts `hcl:"timeouts,block"`
}
type apiManagementDiagnosticAttributes struct {
	ref terra.Reference
}

// AlwaysLogErrors returns a reference to field always_log_errors of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) AlwaysLogErrors() terra.BoolValue {
	return terra.ReferenceAsBool(amd.ref.Append("always_log_errors"))
}

// ApiManagementLoggerId returns a reference to field api_management_logger_id of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) ApiManagementLoggerId() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("api_management_logger_id"))
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("api_management_name"))
}

// HttpCorrelationProtocol returns a reference to field http_correlation_protocol of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) HttpCorrelationProtocol() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("http_correlation_protocol"))
}

// Id returns a reference to field id of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("id"))
}

// Identifier returns a reference to field identifier of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) Identifier() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("identifier"))
}

// LogClientIp returns a reference to field log_client_ip of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) LogClientIp() terra.BoolValue {
	return terra.ReferenceAsBool(amd.ref.Append("log_client_ip"))
}

// OperationNameFormat returns a reference to field operation_name_format of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) OperationNameFormat() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("operation_name_format"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("resource_group_name"))
}

// SamplingPercentage returns a reference to field sampling_percentage of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) SamplingPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(amd.ref.Append("sampling_percentage"))
}

// Verbosity returns a reference to field verbosity of azurerm_api_management_diagnostic.
func (amd apiManagementDiagnosticAttributes) Verbosity() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("verbosity"))
}

func (amd apiManagementDiagnosticAttributes) BackendRequest() terra.ListValue[apimanagementdiagnostic.BackendRequestAttributes] {
	return terra.ReferenceAsList[apimanagementdiagnostic.BackendRequestAttributes](amd.ref.Append("backend_request"))
}

func (amd apiManagementDiagnosticAttributes) BackendResponse() terra.ListValue[apimanagementdiagnostic.BackendResponseAttributes] {
	return terra.ReferenceAsList[apimanagementdiagnostic.BackendResponseAttributes](amd.ref.Append("backend_response"))
}

func (amd apiManagementDiagnosticAttributes) FrontendRequest() terra.ListValue[apimanagementdiagnostic.FrontendRequestAttributes] {
	return terra.ReferenceAsList[apimanagementdiagnostic.FrontendRequestAttributes](amd.ref.Append("frontend_request"))
}

func (amd apiManagementDiagnosticAttributes) FrontendResponse() terra.ListValue[apimanagementdiagnostic.FrontendResponseAttributes] {
	return terra.ReferenceAsList[apimanagementdiagnostic.FrontendResponseAttributes](amd.ref.Append("frontend_response"))
}

func (amd apiManagementDiagnosticAttributes) Timeouts() apimanagementdiagnostic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementdiagnostic.TimeoutsAttributes](amd.ref.Append("timeouts"))
}

type apiManagementDiagnosticState struct {
	AlwaysLogErrors         bool                                            `json:"always_log_errors"`
	ApiManagementLoggerId   string                                          `json:"api_management_logger_id"`
	ApiManagementName       string                                          `json:"api_management_name"`
	HttpCorrelationProtocol string                                          `json:"http_correlation_protocol"`
	Id                      string                                          `json:"id"`
	Identifier              string                                          `json:"identifier"`
	LogClientIp             bool                                            `json:"log_client_ip"`
	OperationNameFormat     string                                          `json:"operation_name_format"`
	ResourceGroupName       string                                          `json:"resource_group_name"`
	SamplingPercentage      float64                                         `json:"sampling_percentage"`
	Verbosity               string                                          `json:"verbosity"`
	BackendRequest          []apimanagementdiagnostic.BackendRequestState   `json:"backend_request"`
	BackendResponse         []apimanagementdiagnostic.BackendResponseState  `json:"backend_response"`
	FrontendRequest         []apimanagementdiagnostic.FrontendRequestState  `json:"frontend_request"`
	FrontendResponse        []apimanagementdiagnostic.FrontendResponseState `json:"frontend_response"`
	Timeouts                *apimanagementdiagnostic.TimeoutsState          `json:"timeouts"`
}
