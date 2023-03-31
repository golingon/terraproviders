// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementapidiagnostic "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagementapidiagnostic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApiManagementApiDiagnostic(name string, args ApiManagementApiDiagnosticArgs) *ApiManagementApiDiagnostic {
	return &ApiManagementApiDiagnostic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementApiDiagnostic)(nil)

type ApiManagementApiDiagnostic struct {
	Name  string
	Args  ApiManagementApiDiagnosticArgs
	state *apiManagementApiDiagnosticState
}

func (amad *ApiManagementApiDiagnostic) Type() string {
	return "azurerm_api_management_api_diagnostic"
}

func (amad *ApiManagementApiDiagnostic) LocalName() string {
	return amad.Name
}

func (amad *ApiManagementApiDiagnostic) Configuration() interface{} {
	return amad.Args
}

func (amad *ApiManagementApiDiagnostic) Attributes() apiManagementApiDiagnosticAttributes {
	return apiManagementApiDiagnosticAttributes{ref: terra.ReferenceResource(amad)}
}

func (amad *ApiManagementApiDiagnostic) ImportState(av io.Reader) error {
	amad.state = &apiManagementApiDiagnosticState{}
	if err := json.NewDecoder(av).Decode(amad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amad.Type(), amad.LocalName(), err)
	}
	return nil
}

func (amad *ApiManagementApiDiagnostic) State() (*apiManagementApiDiagnosticState, bool) {
	return amad.state, amad.state != nil
}

func (amad *ApiManagementApiDiagnostic) StateMust() *apiManagementApiDiagnosticState {
	if amad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amad.Type(), amad.LocalName()))
	}
	return amad.state
}

func (amad *ApiManagementApiDiagnostic) DependOn() terra.Reference {
	return terra.ReferenceResource(amad)
}

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
	// DependsOn contains resources that ApiManagementApiDiagnostic depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apiManagementApiDiagnosticAttributes struct {
	ref terra.Reference
}

func (amad apiManagementApiDiagnosticAttributes) AlwaysLogErrors() terra.BoolValue {
	return terra.ReferenceBool(amad.ref.Append("always_log_errors"))
}

func (amad apiManagementApiDiagnosticAttributes) ApiManagementLoggerId() terra.StringValue {
	return terra.ReferenceString(amad.ref.Append("api_management_logger_id"))
}

func (amad apiManagementApiDiagnosticAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceString(amad.ref.Append("api_management_name"))
}

func (amad apiManagementApiDiagnosticAttributes) ApiName() terra.StringValue {
	return terra.ReferenceString(amad.ref.Append("api_name"))
}

func (amad apiManagementApiDiagnosticAttributes) HttpCorrelationProtocol() terra.StringValue {
	return terra.ReferenceString(amad.ref.Append("http_correlation_protocol"))
}

func (amad apiManagementApiDiagnosticAttributes) Id() terra.StringValue {
	return terra.ReferenceString(amad.ref.Append("id"))
}

func (amad apiManagementApiDiagnosticAttributes) Identifier() terra.StringValue {
	return terra.ReferenceString(amad.ref.Append("identifier"))
}

func (amad apiManagementApiDiagnosticAttributes) LogClientIp() terra.BoolValue {
	return terra.ReferenceBool(amad.ref.Append("log_client_ip"))
}

func (amad apiManagementApiDiagnosticAttributes) OperationNameFormat() terra.StringValue {
	return terra.ReferenceString(amad.ref.Append("operation_name_format"))
}

func (amad apiManagementApiDiagnosticAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(amad.ref.Append("resource_group_name"))
}

func (amad apiManagementApiDiagnosticAttributes) SamplingPercentage() terra.NumberValue {
	return terra.ReferenceNumber(amad.ref.Append("sampling_percentage"))
}

func (amad apiManagementApiDiagnosticAttributes) Verbosity() terra.StringValue {
	return terra.ReferenceString(amad.ref.Append("verbosity"))
}

func (amad apiManagementApiDiagnosticAttributes) BackendRequest() terra.ListValue[apimanagementapidiagnostic.BackendRequestAttributes] {
	return terra.ReferenceList[apimanagementapidiagnostic.BackendRequestAttributes](amad.ref.Append("backend_request"))
}

func (amad apiManagementApiDiagnosticAttributes) BackendResponse() terra.ListValue[apimanagementapidiagnostic.BackendResponseAttributes] {
	return terra.ReferenceList[apimanagementapidiagnostic.BackendResponseAttributes](amad.ref.Append("backend_response"))
}

func (amad apiManagementApiDiagnosticAttributes) FrontendRequest() terra.ListValue[apimanagementapidiagnostic.FrontendRequestAttributes] {
	return terra.ReferenceList[apimanagementapidiagnostic.FrontendRequestAttributes](amad.ref.Append("frontend_request"))
}

func (amad apiManagementApiDiagnosticAttributes) FrontendResponse() terra.ListValue[apimanagementapidiagnostic.FrontendResponseAttributes] {
	return terra.ReferenceList[apimanagementapidiagnostic.FrontendResponseAttributes](amad.ref.Append("frontend_response"))
}

func (amad apiManagementApiDiagnosticAttributes) Timeouts() apimanagementapidiagnostic.TimeoutsAttributes {
	return terra.ReferenceSingle[apimanagementapidiagnostic.TimeoutsAttributes](amad.ref.Append("timeouts"))
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
