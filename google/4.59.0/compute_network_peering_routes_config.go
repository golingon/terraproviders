// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computenetworkpeeringroutesconfig "github.com/golingon/terraproviders/google/4.59.0/computenetworkpeeringroutesconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeNetworkPeeringRoutesConfig creates a new instance of [ComputeNetworkPeeringRoutesConfig].
func NewComputeNetworkPeeringRoutesConfig(name string, args ComputeNetworkPeeringRoutesConfigArgs) *ComputeNetworkPeeringRoutesConfig {
	return &ComputeNetworkPeeringRoutesConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNetworkPeeringRoutesConfig)(nil)

// ComputeNetworkPeeringRoutesConfig represents the Terraform resource google_compute_network_peering_routes_config.
type ComputeNetworkPeeringRoutesConfig struct {
	Name      string
	Args      ComputeNetworkPeeringRoutesConfigArgs
	state     *computeNetworkPeeringRoutesConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNetworkPeeringRoutesConfig].
func (cnprc *ComputeNetworkPeeringRoutesConfig) Type() string {
	return "google_compute_network_peering_routes_config"
}

// LocalName returns the local name for [ComputeNetworkPeeringRoutesConfig].
func (cnprc *ComputeNetworkPeeringRoutesConfig) LocalName() string {
	return cnprc.Name
}

// Configuration returns the configuration (args) for [ComputeNetworkPeeringRoutesConfig].
func (cnprc *ComputeNetworkPeeringRoutesConfig) Configuration() interface{} {
	return cnprc.Args
}

// DependOn is used for other resources to depend on [ComputeNetworkPeeringRoutesConfig].
func (cnprc *ComputeNetworkPeeringRoutesConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(cnprc)
}

// Dependencies returns the list of resources [ComputeNetworkPeeringRoutesConfig] depends_on.
func (cnprc *ComputeNetworkPeeringRoutesConfig) Dependencies() terra.Dependencies {
	return cnprc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNetworkPeeringRoutesConfig].
func (cnprc *ComputeNetworkPeeringRoutesConfig) LifecycleManagement() *terra.Lifecycle {
	return cnprc.Lifecycle
}

// Attributes returns the attributes for [ComputeNetworkPeeringRoutesConfig].
func (cnprc *ComputeNetworkPeeringRoutesConfig) Attributes() computeNetworkPeeringRoutesConfigAttributes {
	return computeNetworkPeeringRoutesConfigAttributes{ref: terra.ReferenceResource(cnprc)}
}

// ImportState imports the given attribute values into [ComputeNetworkPeeringRoutesConfig]'s state.
func (cnprc *ComputeNetworkPeeringRoutesConfig) ImportState(av io.Reader) error {
	cnprc.state = &computeNetworkPeeringRoutesConfigState{}
	if err := json.NewDecoder(av).Decode(cnprc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cnprc.Type(), cnprc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNetworkPeeringRoutesConfig] has state.
func (cnprc *ComputeNetworkPeeringRoutesConfig) State() (*computeNetworkPeeringRoutesConfigState, bool) {
	return cnprc.state, cnprc.state != nil
}

// StateMust returns the state for [ComputeNetworkPeeringRoutesConfig]. Panics if the state is nil.
func (cnprc *ComputeNetworkPeeringRoutesConfig) StateMust() *computeNetworkPeeringRoutesConfigState {
	if cnprc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cnprc.Type(), cnprc.LocalName()))
	}
	return cnprc.state
}

// ComputeNetworkPeeringRoutesConfigArgs contains the configurations for google_compute_network_peering_routes_config.
type ComputeNetworkPeeringRoutesConfigArgs struct {
	// ExportCustomRoutes: bool, required
	ExportCustomRoutes terra.BoolValue `hcl:"export_custom_routes,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImportCustomRoutes: bool, required
	ImportCustomRoutes terra.BoolValue `hcl:"import_custom_routes,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Peering: string, required
	Peering terra.StringValue `hcl:"peering,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computenetworkpeeringroutesconfig.Timeouts `hcl:"timeouts,block"`
}
type computeNetworkPeeringRoutesConfigAttributes struct {
	ref terra.Reference
}

// ExportCustomRoutes returns a reference to field export_custom_routes of google_compute_network_peering_routes_config.
func (cnprc computeNetworkPeeringRoutesConfigAttributes) ExportCustomRoutes() terra.BoolValue {
	return terra.ReferenceAsBool(cnprc.ref.Append("export_custom_routes"))
}

// Id returns a reference to field id of google_compute_network_peering_routes_config.
func (cnprc computeNetworkPeeringRoutesConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cnprc.ref.Append("id"))
}

// ImportCustomRoutes returns a reference to field import_custom_routes of google_compute_network_peering_routes_config.
func (cnprc computeNetworkPeeringRoutesConfigAttributes) ImportCustomRoutes() terra.BoolValue {
	return terra.ReferenceAsBool(cnprc.ref.Append("import_custom_routes"))
}

// Network returns a reference to field network of google_compute_network_peering_routes_config.
func (cnprc computeNetworkPeeringRoutesConfigAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cnprc.ref.Append("network"))
}

// Peering returns a reference to field peering of google_compute_network_peering_routes_config.
func (cnprc computeNetworkPeeringRoutesConfigAttributes) Peering() terra.StringValue {
	return terra.ReferenceAsString(cnprc.ref.Append("peering"))
}

// Project returns a reference to field project of google_compute_network_peering_routes_config.
func (cnprc computeNetworkPeeringRoutesConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cnprc.ref.Append("project"))
}

func (cnprc computeNetworkPeeringRoutesConfigAttributes) Timeouts() computenetworkpeeringroutesconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenetworkpeeringroutesconfig.TimeoutsAttributes](cnprc.ref.Append("timeouts"))
}

type computeNetworkPeeringRoutesConfigState struct {
	ExportCustomRoutes bool                                             `json:"export_custom_routes"`
	Id                 string                                           `json:"id"`
	ImportCustomRoutes bool                                             `json:"import_custom_routes"`
	Network            string                                           `json:"network"`
	Peering            string                                           `json:"peering"`
	Project            string                                           `json:"project"`
	Timeouts           *computenetworkpeeringroutesconfig.TimeoutsState `json:"timeouts"`
}
