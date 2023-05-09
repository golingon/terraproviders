// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementapidiagnostic "github.com/golingon/terraproviders/azurerm/3.55.0/apimanagementapidiagnostic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementApiDiagnostic creates a new instance of [ApiManagementApiDiagnostic].
func NewApiManagementApiDiagnostic(name string, args ApiManagementApiDiagnosticArgs) *ApiManagementApiDiagnostic {
	return &ApiManagementApiDiagnostic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementApiDiagnostic)(nil)

// ApiManagementApiDiagnostic represents the Terraform resource azurerm_api_management_api_diagnostic.
type ApiManagementApiDiagnostic struct {
	Name      string
	Args      ApiManagementApiDiagnosticArgs
	state     *apiManagementApiDiagnosticState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementApiDiagnostic].
func (amad *ApiManagementApiDiagnostic) Type() string {
	return "azurerm_api_management_api_diagnostic"
}

// LocalName returns the local name for [ApiManagementApiDiagnostic].
func (amad *ApiManagementApiDiagnostic) LocalName() string {
	return amad.Name
}

// Configuration returns the configuration (args) for [ApiManagementApiDiagnostic].
func (amad *ApiManagementApiDiagnostic) Configuration() interface{} {
	return amad.Args
}

// DependOn is used for other resources to depend on [ApiManagementApiDiagnostic].
func (amad *ApiManagementApiDiagnostic) DependOn() terra.Reference {
	return terra.ReferenceResource(amad)
}

// Dependencies returns the list of resources [ApiManagementApiDiagnostic] depends_on.
func (amad *ApiManagementApiDiagnostic) Dependencies() terra.Dependencies {
	return amad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementApiDiagnostic].
func (amad *ApiManagementApiDiagnostic) LifecycleManagement() *terra.Lifecycle {
	return amad.Lifecycle
}

// Attributes returns the attributes for [ApiManagementApiDiagnostic].
func (amad *ApiManagementApiDiagnostic) Attributes() apiManagementApiDiagnosticAttributes {
	return apiManagementApiDiagnosticAttributes{ref: terra.ReferenceResource(amad)}
}

// ImportState imports the given attribute values into [ApiManagementApiDiagnostic]'s state.
func (amad *ApiManagementApiDiagnostic) ImportState(av io.Reader) error {
	amad.state = &apiManagementApiDiagnosticState{}
	if err := json.NewDecoder(av).Decode(amad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amad.Type(), amad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementApiDiagnostic] has state.
func (amad *ApiManagementApiDiagnostic) State() (*apiManagementApiDiagnosticState, bool) {
	return amad.state, amad.state != nil
}

// StateMust returns the state for [ApiManagementApiDiagnostic]. Panics if the state is nil.
func (amad *ApiManagementApiDiagnostic) StateMust() *apiManagementApiDiagnosticState {
	if amad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amad.Type(), amad.LocalName()))
	}
	return amad.state
}

// ApiManagementApiDiagnosticArgs contains the configurations for azurerm_api_management_api_diagnostic.
type ApiManagementApiDiagnosticArgs struct {
	// AlwaysLogErrors: bool, optional
	AlwaysLogErrors terra.BoolValue `hcl:"always_log_errors,attr"`
	// ApiManagementLoggerId: string, required
	ApiManagementLoggerId terra.StringValue `hcl:"api_management_logger_id,attr" validate:"required"`
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ApiName: string, required
	ApiName terra.StringValue `hcl:"api_name,attr" validate:"required"`
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
	BackendRequest *apimanagementapidiagnostic.BackendRequest `hcl:"backend_request,block"`
	// BackendResponse: optional
	BackendResponse *apimanagementapidiagnostic.BackendResponse `hcl:"backend_response,block"`
	// FrontendRequest: optional
	FrontendRequest *apimanagementapidiagnostic.FrontendRequest `hcl:"frontend_request,block"`
	// FrontendResponse: optional
	FrontendResponse *apimanagementapidiagnostic.FrontendResponse `hcl:"frontend_response,block"`
	// Timeouts: optional
	Timeouts *apimanagementapidiagnostic.Timeouts `hcl:"timeouts,block"`
}
type apiManagementApiDiagnosticAttributes struct {
	ref terra.Reference
}

// AlwaysLogErrors returns a reference to field always_log_errors of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) AlwaysLogErrors() terra.BoolValue {
	return terra.ReferenceAsBool(amad.ref.Append("always_log_errors"))
}

// ApiManagementLoggerId returns a reference to field api_management_logger_id of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) ApiManagementLoggerId() terra.StringValue {
	return terra.ReferenceAsString(amad.ref.Append("api_management_logger_id"))
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amad.ref.Append("api_management_name"))
}

// ApiName returns a reference to field api_name of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) ApiName() terra.StringValue {
	return terra.ReferenceAsString(amad.ref.Append("api_name"))
}

// HttpCorrelationProtocol returns a reference to field http_correlation_protocol of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) HttpCorrelationProtocol() terra.StringValue {
	return terra.ReferenceAsString(amad.ref.Append("http_correlation_protocol"))
}

// Id returns a reference to field id of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amad.ref.Append("id"))
}

// Identifier returns a reference to field identifier of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) Identifier() terra.StringValue {
	return terra.ReferenceAsString(amad.ref.Append("identifier"))
}

// LogClientIp returns a reference to field log_client_ip of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) LogClientIp() terra.BoolValue {
	return terra.ReferenceAsBool(amad.ref.Append("log_client_ip"))
}

// OperationNameFormat returns a reference to field operation_name_format of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) OperationNameFormat() terra.StringValue {
	return terra.ReferenceAsString(amad.ref.Append("operation_name_format"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amad.ref.Append("resource_group_name"))
}

// SamplingPercentage returns a reference to field sampling_percentage of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) SamplingPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(amad.ref.Append("sampling_percentage"))
}

// Verbosity returns a reference to field verbosity of azurerm_api_management_api_diagnostic.
func (amad apiManagementApiDiagnosticAttributes) Verbosity() terra.StringValue {
	return terra.ReferenceAsString(amad.ref.Append("verbosity"))
}

func (amad apiManagementApiDiagnosticAttributes) BackendRequest() terra.ListValue[apimanagementapidiagnostic.BackendRequestAttributes] {
	return terra.ReferenceAsList[apimanagementapidiagnostic.BackendRequestAttributes](amad.ref.Append("backend_request"))
}

func (amad apiManagementApiDiagnosticAttributes) BackendResponse() terra.ListValue[apimanagementapidiagnostic.BackendResponseAttributes] {
	return terra.ReferenceAsList[apimanagementapidiagnostic.BackendResponseAttributes](amad.ref.Append("backend_response"))
}

func (amad apiManagementApiDiagnosticAttributes) FrontendRequest() terra.ListValue[apimanagementapidiagnostic.FrontendRequestAttributes] {
	return terra.ReferenceAsList[apimanagementapidiagnostic.FrontendRequestAttributes](amad.ref.Append("frontend_request"))
}

func (amad apiManagementApiDiagnosticAttributes) FrontendResponse() terra.ListValue[apimanagementapidiagnostic.FrontendResponseAttributes] {
	return terra.ReferenceAsList[apimanagementapidiagnostic.FrontendResponseAttributes](amad.ref.Append("frontend_response"))
}

func (amad apiManagementApiDiagnosticAttributes) Timeouts() apimanagementapidiagnostic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementapidiagnostic.TimeoutsAttributes](amad.ref.Append("timeouts"))
}

type apiManagementApiDiagnosticState struct {
	AlwaysLogErrors         bool                                               `json:"always_log_errors"`
	ApiManagementLoggerId   string                                             `json:"api_management_logger_id"`
	ApiManagementName       string                                             `json:"api_management_name"`
	ApiName                 string                                             `json:"api_name"`
	HttpCorrelationProtocol string                                             `json:"http_correlation_protocol"`
	Id                      string                                             `json:"id"`
	Identifier              string                                             `json:"identifier"`
	LogClientIp             bool                                               `json:"log_client_ip"`
	OperationNameFormat     string                                             `json:"operation_name_format"`
	ResourceGroupName       string                                             `json:"resource_group_name"`
	SamplingPercentage      float64                                            `json:"sampling_percentage"`
	Verbosity               string                                             `json:"verbosity"`
	BackendRequest          []apimanagementapidiagnostic.BackendRequestState   `json:"backend_request"`
	BackendResponse         []apimanagementapidiagnostic.BackendResponseState  `json:"backend_response"`
	FrontendRequest         []apimanagementapidiagnostic.FrontendRequestState  `json:"frontend_request"`
	FrontendResponse        []apimanagementapidiagnostic.FrontendResponseState `json:"frontend_response"`
	Timeouts                *apimanagementapidiagnostic.TimeoutsState          `json:"timeouts"`
}
