// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	endpointsservice "github.com/golingon/terraproviders/googlebeta/4.76.0/endpointsservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEndpointsService creates a new instance of [EndpointsService].
func NewEndpointsService(name string, args EndpointsServiceArgs) *EndpointsService {
	return &EndpointsService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EndpointsService)(nil)

// EndpointsService represents the Terraform resource google_endpoints_service.
type EndpointsService struct {
	Name      string
	Args      EndpointsServiceArgs
	state     *endpointsServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EndpointsService].
func (es *EndpointsService) Type() string {
	return "google_endpoints_service"
}

// LocalName returns the local name for [EndpointsService].
func (es *EndpointsService) LocalName() string {
	return es.Name
}

// Configuration returns the configuration (args) for [EndpointsService].
func (es *EndpointsService) Configuration() interface{} {
	return es.Args
}

// DependOn is used for other resources to depend on [EndpointsService].
func (es *EndpointsService) DependOn() terra.Reference {
	return terra.ReferenceResource(es)
}

// Dependencies returns the list of resources [EndpointsService] depends_on.
func (es *EndpointsService) Dependencies() terra.Dependencies {
	return es.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EndpointsService].
func (es *EndpointsService) LifecycleManagement() *terra.Lifecycle {
	return es.Lifecycle
}

// Attributes returns the attributes for [EndpointsService].
func (es *EndpointsService) Attributes() endpointsServiceAttributes {
	return endpointsServiceAttributes{ref: terra.ReferenceResource(es)}
}

// ImportState imports the given attribute values into [EndpointsService]'s state.
func (es *EndpointsService) ImportState(av io.Reader) error {
	es.state = &endpointsServiceState{}
	if err := json.NewDecoder(av).Decode(es.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", es.Type(), es.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EndpointsService] has state.
func (es *EndpointsService) State() (*endpointsServiceState, bool) {
	return es.state, es.state != nil
}

// StateMust returns the state for [EndpointsService]. Panics if the state is nil.
func (es *EndpointsService) StateMust() *endpointsServiceState {
	if es.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", es.Type(), es.LocalName()))
	}
	return es.state
}

// EndpointsServiceArgs contains the configurations for google_endpoints_service.
type EndpointsServiceArgs struct {
	// GrpcConfig: string, optional
	GrpcConfig terra.StringValue `hcl:"grpc_config,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OpenapiConfig: string, optional
	OpenapiConfig terra.StringValue `hcl:"openapi_config,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ProtocOutputBase64: string, optional
	ProtocOutputBase64 terra.StringValue `hcl:"protoc_output_base64,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Apis: min=0
	Apis []endpointsservice.Apis `hcl:"apis,block" validate:"min=0"`
	// Endpoints: min=0
	Endpoints []endpointsservice.Endpoints `hcl:"endpoints,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *endpointsservice.Timeouts `hcl:"timeouts,block"`
}
type endpointsServiceAttributes struct {
	ref terra.Reference
}

// ConfigId returns a reference to field config_id of google_endpoints_service.
func (es endpointsServiceAttributes) ConfigId() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("config_id"))
}

// DnsAddress returns a reference to field dns_address of google_endpoints_service.
func (es endpointsServiceAttributes) DnsAddress() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("dns_address"))
}

// GrpcConfig returns a reference to field grpc_config of google_endpoints_service.
func (es endpointsServiceAttributes) GrpcConfig() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("grpc_config"))
}

// Id returns a reference to field id of google_endpoints_service.
func (es endpointsServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("id"))
}

// OpenapiConfig returns a reference to field openapi_config of google_endpoints_service.
func (es endpointsServiceAttributes) OpenapiConfig() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("openapi_config"))
}

// Project returns a reference to field project of google_endpoints_service.
func (es endpointsServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("project"))
}

// ProtocOutputBase64 returns a reference to field protoc_output_base64 of google_endpoints_service.
func (es endpointsServiceAttributes) ProtocOutputBase64() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("protoc_output_base64"))
}

// ServiceName returns a reference to field service_name of google_endpoints_service.
func (es endpointsServiceAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("service_name"))
}

func (es endpointsServiceAttributes) Apis() terra.ListValue[endpointsservice.ApisAttributes] {
	return terra.ReferenceAsList[endpointsservice.ApisAttributes](es.ref.Append("apis"))
}

func (es endpointsServiceAttributes) Endpoints() terra.ListValue[endpointsservice.EndpointsAttributes] {
	return terra.ReferenceAsList[endpointsservice.EndpointsAttributes](es.ref.Append("endpoints"))
}

func (es endpointsServiceAttributes) Timeouts() endpointsservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[endpointsservice.TimeoutsAttributes](es.ref.Append("timeouts"))
}

type endpointsServiceState struct {
	ConfigId           string                            `json:"config_id"`
	DnsAddress         string                            `json:"dns_address"`
	GrpcConfig         string                            `json:"grpc_config"`
	Id                 string                            `json:"id"`
	OpenapiConfig      string                            `json:"openapi_config"`
	Project            string                            `json:"project"`
	ProtocOutputBase64 string                            `json:"protoc_output_base64"`
	ServiceName        string                            `json:"service_name"`
	Apis               []endpointsservice.ApisState      `json:"apis"`
	Endpoints          []endpointsservice.EndpointsState `json:"endpoints"`
	Timeouts           *endpointsservice.TimeoutsState   `json:"timeouts"`
}
