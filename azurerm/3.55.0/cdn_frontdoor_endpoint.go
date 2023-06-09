// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnfrontdoorendpoint "github.com/golingon/terraproviders/azurerm/3.55.0/cdnfrontdoorendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnFrontdoorEndpoint creates a new instance of [CdnFrontdoorEndpoint].
func NewCdnFrontdoorEndpoint(name string, args CdnFrontdoorEndpointArgs) *CdnFrontdoorEndpoint {
	return &CdnFrontdoorEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnFrontdoorEndpoint)(nil)

// CdnFrontdoorEndpoint represents the Terraform resource azurerm_cdn_frontdoor_endpoint.
type CdnFrontdoorEndpoint struct {
	Name      string
	Args      CdnFrontdoorEndpointArgs
	state     *cdnFrontdoorEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnFrontdoorEndpoint].
func (cfe *CdnFrontdoorEndpoint) Type() string {
	return "azurerm_cdn_frontdoor_endpoint"
}

// LocalName returns the local name for [CdnFrontdoorEndpoint].
func (cfe *CdnFrontdoorEndpoint) LocalName() string {
	return cfe.Name
}

// Configuration returns the configuration (args) for [CdnFrontdoorEndpoint].
func (cfe *CdnFrontdoorEndpoint) Configuration() interface{} {
	return cfe.Args
}

// DependOn is used for other resources to depend on [CdnFrontdoorEndpoint].
func (cfe *CdnFrontdoorEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(cfe)
}

// Dependencies returns the list of resources [CdnFrontdoorEndpoint] depends_on.
func (cfe *CdnFrontdoorEndpoint) Dependencies() terra.Dependencies {
	return cfe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnFrontdoorEndpoint].
func (cfe *CdnFrontdoorEndpoint) LifecycleManagement() *terra.Lifecycle {
	return cfe.Lifecycle
}

// Attributes returns the attributes for [CdnFrontdoorEndpoint].
func (cfe *CdnFrontdoorEndpoint) Attributes() cdnFrontdoorEndpointAttributes {
	return cdnFrontdoorEndpointAttributes{ref: terra.ReferenceResource(cfe)}
}

// ImportState imports the given attribute values into [CdnFrontdoorEndpoint]'s state.
func (cfe *CdnFrontdoorEndpoint) ImportState(av io.Reader) error {
	cfe.state = &cdnFrontdoorEndpointState{}
	if err := json.NewDecoder(av).Decode(cfe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfe.Type(), cfe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnFrontdoorEndpoint] has state.
func (cfe *CdnFrontdoorEndpoint) State() (*cdnFrontdoorEndpointState, bool) {
	return cfe.state, cfe.state != nil
}

// StateMust returns the state for [CdnFrontdoorEndpoint]. Panics if the state is nil.
func (cfe *CdnFrontdoorEndpoint) StateMust() *cdnFrontdoorEndpointState {
	if cfe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfe.Type(), cfe.LocalName()))
	}
	return cfe.state
}

// CdnFrontdoorEndpointArgs contains the configurations for azurerm_cdn_frontdoor_endpoint.
type CdnFrontdoorEndpointArgs struct {
	// CdnFrontdoorProfileId: string, required
	CdnFrontdoorProfileId terra.StringValue `hcl:"cdn_frontdoor_profile_id,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *cdnfrontdoorendpoint.Timeouts `hcl:"timeouts,block"`
}
type cdnFrontdoorEndpointAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_endpoint.
func (cfe cdnFrontdoorEndpointAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(cfe.ref.Append("cdn_frontdoor_profile_id"))
}

// Enabled returns a reference to field enabled of azurerm_cdn_frontdoor_endpoint.
func (cfe cdnFrontdoorEndpointAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cfe.ref.Append("enabled"))
}

// HostName returns a reference to field host_name of azurerm_cdn_frontdoor_endpoint.
func (cfe cdnFrontdoorEndpointAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(cfe.ref.Append("host_name"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_endpoint.
func (cfe cdnFrontdoorEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfe.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_endpoint.
func (cfe cdnFrontdoorEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfe.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_cdn_frontdoor_endpoint.
func (cfe cdnFrontdoorEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cfe.ref.Append("tags"))
}

func (cfe cdnFrontdoorEndpointAttributes) Timeouts() cdnfrontdoorendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnfrontdoorendpoint.TimeoutsAttributes](cfe.ref.Append("timeouts"))
}

type cdnFrontdoorEndpointState struct {
	CdnFrontdoorProfileId string                              `json:"cdn_frontdoor_profile_id"`
	Enabled               bool                                `json:"enabled"`
	HostName              string                              `json:"host_name"`
	Id                    string                              `json:"id"`
	Name                  string                              `json:"name"`
	Tags                  map[string]string                   `json:"tags"`
	Timeouts              *cdnfrontdoorendpoint.TimeoutsState `json:"timeouts"`
}
