// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnendpoint "github.com/golingon/terraproviders/azurerm/3.55.0/cdnendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnEndpoint creates a new instance of [CdnEndpoint].
func NewCdnEndpoint(name string, args CdnEndpointArgs) *CdnEndpoint {
	return &CdnEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnEndpoint)(nil)

// CdnEndpoint represents the Terraform resource azurerm_cdn_endpoint.
type CdnEndpoint struct {
	Name      string
	Args      CdnEndpointArgs
	state     *cdnEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnEndpoint].
func (ce *CdnEndpoint) Type() string {
	return "azurerm_cdn_endpoint"
}

// LocalName returns the local name for [CdnEndpoint].
func (ce *CdnEndpoint) LocalName() string {
	return ce.Name
}

// Configuration returns the configuration (args) for [CdnEndpoint].
func (ce *CdnEndpoint) Configuration() interface{} {
	return ce.Args
}

// DependOn is used for other resources to depend on [CdnEndpoint].
func (ce *CdnEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(ce)
}

// Dependencies returns the list of resources [CdnEndpoint] depends_on.
func (ce *CdnEndpoint) Dependencies() terra.Dependencies {
	return ce.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnEndpoint].
func (ce *CdnEndpoint) LifecycleManagement() *terra.Lifecycle {
	return ce.Lifecycle
}

// Attributes returns the attributes for [CdnEndpoint].
func (ce *CdnEndpoint) Attributes() cdnEndpointAttributes {
	return cdnEndpointAttributes{ref: terra.ReferenceResource(ce)}
}

// ImportState imports the given attribute values into [CdnEndpoint]'s state.
func (ce *CdnEndpoint) ImportState(av io.Reader) error {
	ce.state = &cdnEndpointState{}
	if err := json.NewDecoder(av).Decode(ce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ce.Type(), ce.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnEndpoint] has state.
func (ce *CdnEndpoint) State() (*cdnEndpointState, bool) {
	return ce.state, ce.state != nil
}

// StateMust returns the state for [CdnEndpoint]. Panics if the state is nil.
func (ce *CdnEndpoint) StateMust() *cdnEndpointState {
	if ce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ce.Type(), ce.LocalName()))
	}
	return ce.state
}

// CdnEndpointArgs contains the configurations for azurerm_cdn_endpoint.
type CdnEndpointArgs struct {
	// ContentTypesToCompress: set of string, optional
	ContentTypesToCompress terra.SetValue[terra.StringValue] `hcl:"content_types_to_compress,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsCompressionEnabled: bool, optional
	IsCompressionEnabled terra.BoolValue `hcl:"is_compression_enabled,attr"`
	// IsHttpAllowed: bool, optional
	IsHttpAllowed terra.BoolValue `hcl:"is_http_allowed,attr"`
	// IsHttpsAllowed: bool, optional
	IsHttpsAllowed terra.BoolValue `hcl:"is_https_allowed,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OptimizationType: string, optional
	OptimizationType terra.StringValue `hcl:"optimization_type,attr"`
	// OriginHostHeader: string, optional
	OriginHostHeader terra.StringValue `hcl:"origin_host_header,attr"`
	// OriginPath: string, optional
	OriginPath terra.StringValue `hcl:"origin_path,attr"`
	// ProbePath: string, optional
	ProbePath terra.StringValue `hcl:"probe_path,attr"`
	// ProfileName: string, required
	ProfileName terra.StringValue `hcl:"profile_name,attr" validate:"required"`
	// QuerystringCachingBehaviour: string, optional
	QuerystringCachingBehaviour terra.StringValue `hcl:"querystring_caching_behaviour,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DeliveryRule: min=0
	DeliveryRule []cdnendpoint.DeliveryRule `hcl:"delivery_rule,block" validate:"min=0"`
	// GeoFilter: min=0
	GeoFilter []cdnendpoint.GeoFilter `hcl:"geo_filter,block" validate:"min=0"`
	// GlobalDeliveryRule: optional
	GlobalDeliveryRule *cdnendpoint.GlobalDeliveryRule `hcl:"global_delivery_rule,block"`
	// Origin: min=1
	Origin []cdnendpoint.Origin `hcl:"origin,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *cdnendpoint.Timeouts `hcl:"timeouts,block"`
}
type cdnEndpointAttributes struct {
	ref terra.Reference
}

// ContentTypesToCompress returns a reference to field content_types_to_compress of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) ContentTypesToCompress() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ce.ref.Append("content_types_to_compress"))
}

// Fqdn returns a reference to field fqdn of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("id"))
}

// IsCompressionEnabled returns a reference to field is_compression_enabled of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) IsCompressionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ce.ref.Append("is_compression_enabled"))
}

// IsHttpAllowed returns a reference to field is_http_allowed of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) IsHttpAllowed() terra.BoolValue {
	return terra.ReferenceAsBool(ce.ref.Append("is_http_allowed"))
}

// IsHttpsAllowed returns a reference to field is_https_allowed of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) IsHttpsAllowed() terra.BoolValue {
	return terra.ReferenceAsBool(ce.ref.Append("is_https_allowed"))
}

// Location returns a reference to field location of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("name"))
}

// OptimizationType returns a reference to field optimization_type of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) OptimizationType() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("optimization_type"))
}

// OriginHostHeader returns a reference to field origin_host_header of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) OriginHostHeader() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("origin_host_header"))
}

// OriginPath returns a reference to field origin_path of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) OriginPath() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("origin_path"))
}

// ProbePath returns a reference to field probe_path of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) ProbePath() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("probe_path"))
}

// ProfileName returns a reference to field profile_name of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("profile_name"))
}

// QuerystringCachingBehaviour returns a reference to field querystring_caching_behaviour of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) QuerystringCachingBehaviour() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("querystring_caching_behaviour"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_cdn_endpoint.
func (ce cdnEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ce.ref.Append("tags"))
}

func (ce cdnEndpointAttributes) DeliveryRule() terra.ListValue[cdnendpoint.DeliveryRuleAttributes] {
	return terra.ReferenceAsList[cdnendpoint.DeliveryRuleAttributes](ce.ref.Append("delivery_rule"))
}

func (ce cdnEndpointAttributes) GeoFilter() terra.ListValue[cdnendpoint.GeoFilterAttributes] {
	return terra.ReferenceAsList[cdnendpoint.GeoFilterAttributes](ce.ref.Append("geo_filter"))
}

func (ce cdnEndpointAttributes) GlobalDeliveryRule() terra.ListValue[cdnendpoint.GlobalDeliveryRuleAttributes] {
	return terra.ReferenceAsList[cdnendpoint.GlobalDeliveryRuleAttributes](ce.ref.Append("global_delivery_rule"))
}

func (ce cdnEndpointAttributes) Origin() terra.SetValue[cdnendpoint.OriginAttributes] {
	return terra.ReferenceAsSet[cdnendpoint.OriginAttributes](ce.ref.Append("origin"))
}

func (ce cdnEndpointAttributes) Timeouts() cdnendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnendpoint.TimeoutsAttributes](ce.ref.Append("timeouts"))
}

type cdnEndpointState struct {
	ContentTypesToCompress      []string                              `json:"content_types_to_compress"`
	Fqdn                        string                                `json:"fqdn"`
	Id                          string                                `json:"id"`
	IsCompressionEnabled        bool                                  `json:"is_compression_enabled"`
	IsHttpAllowed               bool                                  `json:"is_http_allowed"`
	IsHttpsAllowed              bool                                  `json:"is_https_allowed"`
	Location                    string                                `json:"location"`
	Name                        string                                `json:"name"`
	OptimizationType            string                                `json:"optimization_type"`
	OriginHostHeader            string                                `json:"origin_host_header"`
	OriginPath                  string                                `json:"origin_path"`
	ProbePath                   string                                `json:"probe_path"`
	ProfileName                 string                                `json:"profile_name"`
	QuerystringCachingBehaviour string                                `json:"querystring_caching_behaviour"`
	ResourceGroupName           string                                `json:"resource_group_name"`
	Tags                        map[string]string                     `json:"tags"`
	DeliveryRule                []cdnendpoint.DeliveryRuleState       `json:"delivery_rule"`
	GeoFilter                   []cdnendpoint.GeoFilterState          `json:"geo_filter"`
	GlobalDeliveryRule          []cdnendpoint.GlobalDeliveryRuleState `json:"global_delivery_rule"`
	Origin                      []cdnendpoint.OriginState             `json:"origin"`
	Timeouts                    *cdnendpoint.TimeoutsState            `json:"timeouts"`
}
