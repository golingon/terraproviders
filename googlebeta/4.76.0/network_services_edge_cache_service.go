// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networkservicesedgecacheservice "github.com/golingon/terraproviders/googlebeta/4.76.0/networkservicesedgecacheservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkServicesEdgeCacheService creates a new instance of [NetworkServicesEdgeCacheService].
func NewNetworkServicesEdgeCacheService(name string, args NetworkServicesEdgeCacheServiceArgs) *NetworkServicesEdgeCacheService {
	return &NetworkServicesEdgeCacheService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkServicesEdgeCacheService)(nil)

// NetworkServicesEdgeCacheService represents the Terraform resource google_network_services_edge_cache_service.
type NetworkServicesEdgeCacheService struct {
	Name      string
	Args      NetworkServicesEdgeCacheServiceArgs
	state     *networkServicesEdgeCacheServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkServicesEdgeCacheService].
func (nsecs *NetworkServicesEdgeCacheService) Type() string {
	return "google_network_services_edge_cache_service"
}

// LocalName returns the local name for [NetworkServicesEdgeCacheService].
func (nsecs *NetworkServicesEdgeCacheService) LocalName() string {
	return nsecs.Name
}

// Configuration returns the configuration (args) for [NetworkServicesEdgeCacheService].
func (nsecs *NetworkServicesEdgeCacheService) Configuration() interface{} {
	return nsecs.Args
}

// DependOn is used for other resources to depend on [NetworkServicesEdgeCacheService].
func (nsecs *NetworkServicesEdgeCacheService) DependOn() terra.Reference {
	return terra.ReferenceResource(nsecs)
}

// Dependencies returns the list of resources [NetworkServicesEdgeCacheService] depends_on.
func (nsecs *NetworkServicesEdgeCacheService) Dependencies() terra.Dependencies {
	return nsecs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkServicesEdgeCacheService].
func (nsecs *NetworkServicesEdgeCacheService) LifecycleManagement() *terra.Lifecycle {
	return nsecs.Lifecycle
}

// Attributes returns the attributes for [NetworkServicesEdgeCacheService].
func (nsecs *NetworkServicesEdgeCacheService) Attributes() networkServicesEdgeCacheServiceAttributes {
	return networkServicesEdgeCacheServiceAttributes{ref: terra.ReferenceResource(nsecs)}
}

// ImportState imports the given attribute values into [NetworkServicesEdgeCacheService]'s state.
func (nsecs *NetworkServicesEdgeCacheService) ImportState(av io.Reader) error {
	nsecs.state = &networkServicesEdgeCacheServiceState{}
	if err := json.NewDecoder(av).Decode(nsecs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsecs.Type(), nsecs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkServicesEdgeCacheService] has state.
func (nsecs *NetworkServicesEdgeCacheService) State() (*networkServicesEdgeCacheServiceState, bool) {
	return nsecs.state, nsecs.state != nil
}

// StateMust returns the state for [NetworkServicesEdgeCacheService]. Panics if the state is nil.
func (nsecs *NetworkServicesEdgeCacheService) StateMust() *networkServicesEdgeCacheServiceState {
	if nsecs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsecs.Type(), nsecs.LocalName()))
	}
	return nsecs.state
}

// NetworkServicesEdgeCacheServiceArgs contains the configurations for google_network_services_edge_cache_service.
type NetworkServicesEdgeCacheServiceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableHttp2: bool, optional
	DisableHttp2 terra.BoolValue `hcl:"disable_http2,attr"`
	// DisableQuic: bool, optional
	DisableQuic terra.BoolValue `hcl:"disable_quic,attr"`
	// EdgeSecurityPolicy: string, optional
	EdgeSecurityPolicy terra.StringValue `hcl:"edge_security_policy,attr"`
	// EdgeSslCertificates: list of string, optional
	EdgeSslCertificates terra.ListValue[terra.StringValue] `hcl:"edge_ssl_certificates,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RequireTls: bool, optional
	RequireTls terra.BoolValue `hcl:"require_tls,attr"`
	// SslPolicy: string, optional
	SslPolicy terra.StringValue `hcl:"ssl_policy,attr"`
	// LogConfig: optional
	LogConfig *networkservicesedgecacheservice.LogConfig `hcl:"log_config,block"`
	// Routing: required
	Routing *networkservicesedgecacheservice.Routing `hcl:"routing,block" validate:"required"`
	// Timeouts: optional
	Timeouts *networkservicesedgecacheservice.Timeouts `hcl:"timeouts,block"`
}
type networkServicesEdgeCacheServiceAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsecs.ref.Append("description"))
}

// DisableHttp2 returns a reference to field disable_http2 of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) DisableHttp2() terra.BoolValue {
	return terra.ReferenceAsBool(nsecs.ref.Append("disable_http2"))
}

// DisableQuic returns a reference to field disable_quic of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) DisableQuic() terra.BoolValue {
	return terra.ReferenceAsBool(nsecs.ref.Append("disable_quic"))
}

// EdgeSecurityPolicy returns a reference to field edge_security_policy of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) EdgeSecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(nsecs.ref.Append("edge_security_policy"))
}

// EdgeSslCertificates returns a reference to field edge_ssl_certificates of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) EdgeSslCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nsecs.ref.Append("edge_ssl_certificates"))
}

// Id returns a reference to field id of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsecs.ref.Append("id"))
}

// Ipv4Addresses returns a reference to field ipv4_addresses of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) Ipv4Addresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nsecs.ref.Append("ipv4_addresses"))
}

// Ipv6Addresses returns a reference to field ipv6_addresses of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) Ipv6Addresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nsecs.ref.Append("ipv6_addresses"))
}

// Labels returns a reference to field labels of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsecs.ref.Append("labels"))
}

// Name returns a reference to field name of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsecs.ref.Append("name"))
}

// Project returns a reference to field project of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nsecs.ref.Append("project"))
}

// RequireTls returns a reference to field require_tls of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) RequireTls() terra.BoolValue {
	return terra.ReferenceAsBool(nsecs.ref.Append("require_tls"))
}

// SslPolicy returns a reference to field ssl_policy of google_network_services_edge_cache_service.
func (nsecs networkServicesEdgeCacheServiceAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceAsString(nsecs.ref.Append("ssl_policy"))
}

func (nsecs networkServicesEdgeCacheServiceAttributes) LogConfig() terra.ListValue[networkservicesedgecacheservice.LogConfigAttributes] {
	return terra.ReferenceAsList[networkservicesedgecacheservice.LogConfigAttributes](nsecs.ref.Append("log_config"))
}

func (nsecs networkServicesEdgeCacheServiceAttributes) Routing() terra.ListValue[networkservicesedgecacheservice.RoutingAttributes] {
	return terra.ReferenceAsList[networkservicesedgecacheservice.RoutingAttributes](nsecs.ref.Append("routing"))
}

func (nsecs networkServicesEdgeCacheServiceAttributes) Timeouts() networkservicesedgecacheservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkservicesedgecacheservice.TimeoutsAttributes](nsecs.ref.Append("timeouts"))
}

type networkServicesEdgeCacheServiceState struct {
	Description         string                                           `json:"description"`
	DisableHttp2        bool                                             `json:"disable_http2"`
	DisableQuic         bool                                             `json:"disable_quic"`
	EdgeSecurityPolicy  string                                           `json:"edge_security_policy"`
	EdgeSslCertificates []string                                         `json:"edge_ssl_certificates"`
	Id                  string                                           `json:"id"`
	Ipv4Addresses       []string                                         `json:"ipv4_addresses"`
	Ipv6Addresses       []string                                         `json:"ipv6_addresses"`
	Labels              map[string]string                                `json:"labels"`
	Name                string                                           `json:"name"`
	Project             string                                           `json:"project"`
	RequireTls          bool                                             `json:"require_tls"`
	SslPolicy           string                                           `json:"ssl_policy"`
	LogConfig           []networkservicesedgecacheservice.LogConfigState `json:"log_config"`
	Routing             []networkservicesedgecacheservice.RoutingState   `json:"routing"`
	Timeouts            *networkservicesedgecacheservice.TimeoutsState   `json:"timeouts"`
}
