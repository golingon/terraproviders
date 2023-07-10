// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	arckubernetescluster "github.com/golingon/terraproviders/azurerm/3.64.0/arckubernetescluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewArcKubernetesCluster creates a new instance of [ArcKubernetesCluster].
func NewArcKubernetesCluster(name string, args ArcKubernetesClusterArgs) *ArcKubernetesCluster {
	return &ArcKubernetesCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ArcKubernetesCluster)(nil)

// ArcKubernetesCluster represents the Terraform resource azurerm_arc_kubernetes_cluster.
type ArcKubernetesCluster struct {
	Name      string
	Args      ArcKubernetesClusterArgs
	state     *arcKubernetesClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ArcKubernetesCluster].
func (akc *ArcKubernetesCluster) Type() string {
	return "azurerm_arc_kubernetes_cluster"
}

// LocalName returns the local name for [ArcKubernetesCluster].
func (akc *ArcKubernetesCluster) LocalName() string {
	return akc.Name
}

// Configuration returns the configuration (args) for [ArcKubernetesCluster].
func (akc *ArcKubernetesCluster) Configuration() interface{} {
	return akc.Args
}

// DependOn is used for other resources to depend on [ArcKubernetesCluster].
func (akc *ArcKubernetesCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(akc)
}

// Dependencies returns the list of resources [ArcKubernetesCluster] depends_on.
func (akc *ArcKubernetesCluster) Dependencies() terra.Dependencies {
	return akc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ArcKubernetesCluster].
func (akc *ArcKubernetesCluster) LifecycleManagement() *terra.Lifecycle {
	return akc.Lifecycle
}

// Attributes returns the attributes for [ArcKubernetesCluster].
func (akc *ArcKubernetesCluster) Attributes() arcKubernetesClusterAttributes {
	return arcKubernetesClusterAttributes{ref: terra.ReferenceResource(akc)}
}

// ImportState imports the given attribute values into [ArcKubernetesCluster]'s state.
func (akc *ArcKubernetesCluster) ImportState(av io.Reader) error {
	akc.state = &arcKubernetesClusterState{}
	if err := json.NewDecoder(av).Decode(akc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akc.Type(), akc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ArcKubernetesCluster] has state.
func (akc *ArcKubernetesCluster) State() (*arcKubernetesClusterState, bool) {
	return akc.state, akc.state != nil
}

// StateMust returns the state for [ArcKubernetesCluster]. Panics if the state is nil.
func (akc *ArcKubernetesCluster) StateMust() *arcKubernetesClusterState {
	if akc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akc.Type(), akc.LocalName()))
	}
	return akc.state
}

// ArcKubernetesClusterArgs contains the configurations for azurerm_arc_kubernetes_cluster.
type ArcKubernetesClusterArgs struct {
	// AgentPublicKeyCertificate: string, required
	AgentPublicKeyCertificate terra.StringValue `hcl:"agent_public_key_certificate,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: required
	Identity *arckubernetescluster.Identity `hcl:"identity,block" validate:"required"`
	// Timeouts: optional
	Timeouts *arckubernetescluster.Timeouts `hcl:"timeouts,block"`
}
type arcKubernetesClusterAttributes struct {
	ref terra.Reference
}

// AgentPublicKeyCertificate returns a reference to field agent_public_key_certificate of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) AgentPublicKeyCertificate() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("agent_public_key_certificate"))
}

// AgentVersion returns a reference to field agent_version of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) AgentVersion() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("agent_version"))
}

// Distribution returns a reference to field distribution of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) Distribution() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("distribution"))
}

// Id returns a reference to field id of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("id"))
}

// Infrastructure returns a reference to field infrastructure of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) Infrastructure() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("infrastructure"))
}

// KubernetesVersion returns a reference to field kubernetes_version of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) KubernetesVersion() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("kubernetes_version"))
}

// Location returns a reference to field location of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("name"))
}

// Offering returns a reference to field offering of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) Offering() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("offering"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akc.ref.Append("tags"))
}

// TotalCoreCount returns a reference to field total_core_count of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) TotalCoreCount() terra.NumberValue {
	return terra.ReferenceAsNumber(akc.ref.Append("total_core_count"))
}

// TotalNodeCount returns a reference to field total_node_count of azurerm_arc_kubernetes_cluster.
func (akc arcKubernetesClusterAttributes) TotalNodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(akc.ref.Append("total_node_count"))
}

func (akc arcKubernetesClusterAttributes) Identity() terra.ListValue[arckubernetescluster.IdentityAttributes] {
	return terra.ReferenceAsList[arckubernetescluster.IdentityAttributes](akc.ref.Append("identity"))
}

func (akc arcKubernetesClusterAttributes) Timeouts() arckubernetescluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[arckubernetescluster.TimeoutsAttributes](akc.ref.Append("timeouts"))
}

type arcKubernetesClusterState struct {
	AgentPublicKeyCertificate string                               `json:"agent_public_key_certificate"`
	AgentVersion              string                               `json:"agent_version"`
	Distribution              string                               `json:"distribution"`
	Id                        string                               `json:"id"`
	Infrastructure            string                               `json:"infrastructure"`
	KubernetesVersion         string                               `json:"kubernetes_version"`
	Location                  string                               `json:"location"`
	Name                      string                               `json:"name"`
	Offering                  string                               `json:"offering"`
	ResourceGroupName         string                               `json:"resource_group_name"`
	Tags                      map[string]string                    `json:"tags"`
	TotalCoreCount            float64                              `json:"total_core_count"`
	TotalNodeCount            float64                              `json:"total_node_count"`
	Identity                  []arckubernetescluster.IdentityState `json:"identity"`
	Timeouts                  *arckubernetescluster.TimeoutsState  `json:"timeouts"`
}
