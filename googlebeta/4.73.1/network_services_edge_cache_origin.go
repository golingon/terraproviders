// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networkservicesedgecacheorigin "github.com/golingon/terraproviders/googlebeta/4.73.1/networkservicesedgecacheorigin"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkServicesEdgeCacheOrigin creates a new instance of [NetworkServicesEdgeCacheOrigin].
func NewNetworkServicesEdgeCacheOrigin(name string, args NetworkServicesEdgeCacheOriginArgs) *NetworkServicesEdgeCacheOrigin {
	return &NetworkServicesEdgeCacheOrigin{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkServicesEdgeCacheOrigin)(nil)

// NetworkServicesEdgeCacheOrigin represents the Terraform resource google_network_services_edge_cache_origin.
type NetworkServicesEdgeCacheOrigin struct {
	Name      string
	Args      NetworkServicesEdgeCacheOriginArgs
	state     *networkServicesEdgeCacheOriginState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkServicesEdgeCacheOrigin].
func (nseco *NetworkServicesEdgeCacheOrigin) Type() string {
	return "google_network_services_edge_cache_origin"
}

// LocalName returns the local name for [NetworkServicesEdgeCacheOrigin].
func (nseco *NetworkServicesEdgeCacheOrigin) LocalName() string {
	return nseco.Name
}

// Configuration returns the configuration (args) for [NetworkServicesEdgeCacheOrigin].
func (nseco *NetworkServicesEdgeCacheOrigin) Configuration() interface{} {
	return nseco.Args
}

// DependOn is used for other resources to depend on [NetworkServicesEdgeCacheOrigin].
func (nseco *NetworkServicesEdgeCacheOrigin) DependOn() terra.Reference {
	return terra.ReferenceResource(nseco)
}

// Dependencies returns the list of resources [NetworkServicesEdgeCacheOrigin] depends_on.
func (nseco *NetworkServicesEdgeCacheOrigin) Dependencies() terra.Dependencies {
	return nseco.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkServicesEdgeCacheOrigin].
func (nseco *NetworkServicesEdgeCacheOrigin) LifecycleManagement() *terra.Lifecycle {
	return nseco.Lifecycle
}

// Attributes returns the attributes for [NetworkServicesEdgeCacheOrigin].
func (nseco *NetworkServicesEdgeCacheOrigin) Attributes() networkServicesEdgeCacheOriginAttributes {
	return networkServicesEdgeCacheOriginAttributes{ref: terra.ReferenceResource(nseco)}
}

// ImportState imports the given attribute values into [NetworkServicesEdgeCacheOrigin]'s state.
func (nseco *NetworkServicesEdgeCacheOrigin) ImportState(av io.Reader) error {
	nseco.state = &networkServicesEdgeCacheOriginState{}
	if err := json.NewDecoder(av).Decode(nseco.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nseco.Type(), nseco.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkServicesEdgeCacheOrigin] has state.
func (nseco *NetworkServicesEdgeCacheOrigin) State() (*networkServicesEdgeCacheOriginState, bool) {
	return nseco.state, nseco.state != nil
}

// StateMust returns the state for [NetworkServicesEdgeCacheOrigin]. Panics if the state is nil.
func (nseco *NetworkServicesEdgeCacheOrigin) StateMust() *networkServicesEdgeCacheOriginState {
	if nseco.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nseco.Type(), nseco.LocalName()))
	}
	return nseco.state
}

// NetworkServicesEdgeCacheOriginArgs contains the configurations for google_network_services_edge_cache_origin.
type NetworkServicesEdgeCacheOriginArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FailoverOrigin: string, optional
	FailoverOrigin terra.StringValue `hcl:"failover_origin,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MaxAttempts: number, optional
	MaxAttempts terra.NumberValue `hcl:"max_attempts,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OriginAddress: string, required
	OriginAddress terra.StringValue `hcl:"origin_address,attr" validate:"required"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// RetryConditions: list of string, optional
	RetryConditions terra.ListValue[terra.StringValue] `hcl:"retry_conditions,attr"`
	// AwsV4Authentication: optional
	AwsV4Authentication *networkservicesedgecacheorigin.AwsV4Authentication `hcl:"aws_v4_authentication,block"`
	// OriginOverrideAction: optional
	OriginOverrideAction *networkservicesedgecacheorigin.OriginOverrideAction `hcl:"origin_override_action,block"`
	// OriginRedirect: optional
	OriginRedirect *networkservicesedgecacheorigin.OriginRedirect `hcl:"origin_redirect,block"`
	// Timeout: optional
	Timeout *networkservicesedgecacheorigin.Timeout `hcl:"timeout,block"`
	// Timeouts: optional
	Timeouts *networkservicesedgecacheorigin.Timeouts `hcl:"timeouts,block"`
}
type networkServicesEdgeCacheOriginAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nseco.ref.Append("description"))
}

// FailoverOrigin returns a reference to field failover_origin of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) FailoverOrigin() terra.StringValue {
	return terra.ReferenceAsString(nseco.ref.Append("failover_origin"))
}

// Id returns a reference to field id of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nseco.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nseco.ref.Append("labels"))
}

// MaxAttempts returns a reference to field max_attempts of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) MaxAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(nseco.ref.Append("max_attempts"))
}

// Name returns a reference to field name of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nseco.ref.Append("name"))
}

// OriginAddress returns a reference to field origin_address of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) OriginAddress() terra.StringValue {
	return terra.ReferenceAsString(nseco.ref.Append("origin_address"))
}

// Port returns a reference to field port of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(nseco.ref.Append("port"))
}

// Project returns a reference to field project of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nseco.ref.Append("project"))
}

// Protocol returns a reference to field protocol of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(nseco.ref.Append("protocol"))
}

// RetryConditions returns a reference to field retry_conditions of google_network_services_edge_cache_origin.
func (nseco networkServicesEdgeCacheOriginAttributes) RetryConditions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nseco.ref.Append("retry_conditions"))
}

func (nseco networkServicesEdgeCacheOriginAttributes) AwsV4Authentication() terra.ListValue[networkservicesedgecacheorigin.AwsV4AuthenticationAttributes] {
	return terra.ReferenceAsList[networkservicesedgecacheorigin.AwsV4AuthenticationAttributes](nseco.ref.Append("aws_v4_authentication"))
}

func (nseco networkServicesEdgeCacheOriginAttributes) OriginOverrideAction() terra.ListValue[networkservicesedgecacheorigin.OriginOverrideActionAttributes] {
	return terra.ReferenceAsList[networkservicesedgecacheorigin.OriginOverrideActionAttributes](nseco.ref.Append("origin_override_action"))
}

func (nseco networkServicesEdgeCacheOriginAttributes) OriginRedirect() terra.ListValue[networkservicesedgecacheorigin.OriginRedirectAttributes] {
	return terra.ReferenceAsList[networkservicesedgecacheorigin.OriginRedirectAttributes](nseco.ref.Append("origin_redirect"))
}

func (nseco networkServicesEdgeCacheOriginAttributes) Timeout() terra.ListValue[networkservicesedgecacheorigin.TimeoutAttributes] {
	return terra.ReferenceAsList[networkservicesedgecacheorigin.TimeoutAttributes](nseco.ref.Append("timeout"))
}

func (nseco networkServicesEdgeCacheOriginAttributes) Timeouts() networkservicesedgecacheorigin.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkservicesedgecacheorigin.TimeoutsAttributes](nseco.ref.Append("timeouts"))
}

type networkServicesEdgeCacheOriginState struct {
	Description          string                                                     `json:"description"`
	FailoverOrigin       string                                                     `json:"failover_origin"`
	Id                   string                                                     `json:"id"`
	Labels               map[string]string                                          `json:"labels"`
	MaxAttempts          float64                                                    `json:"max_attempts"`
	Name                 string                                                     `json:"name"`
	OriginAddress        string                                                     `json:"origin_address"`
	Port                 float64                                                    `json:"port"`
	Project              string                                                     `json:"project"`
	Protocol             string                                                     `json:"protocol"`
	RetryConditions      []string                                                   `json:"retry_conditions"`
	AwsV4Authentication  []networkservicesedgecacheorigin.AwsV4AuthenticationState  `json:"aws_v4_authentication"`
	OriginOverrideAction []networkservicesedgecacheorigin.OriginOverrideActionState `json:"origin_override_action"`
	OriginRedirect       []networkservicesedgecacheorigin.OriginRedirectState       `json:"origin_redirect"`
	Timeout              []networkservicesedgecacheorigin.TimeoutState              `json:"timeout"`
	Timeouts             *networkservicesedgecacheorigin.TimeoutsState              `json:"timeouts"`
}
