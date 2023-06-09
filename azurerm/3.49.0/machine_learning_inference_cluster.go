// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	machinelearninginferencecluster "github.com/golingon/terraproviders/azurerm/3.49.0/machinelearninginferencecluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMachineLearningInferenceCluster creates a new instance of [MachineLearningInferenceCluster].
func NewMachineLearningInferenceCluster(name string, args MachineLearningInferenceClusterArgs) *MachineLearningInferenceCluster {
	return &MachineLearningInferenceCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MachineLearningInferenceCluster)(nil)

// MachineLearningInferenceCluster represents the Terraform resource azurerm_machine_learning_inference_cluster.
type MachineLearningInferenceCluster struct {
	Name      string
	Args      MachineLearningInferenceClusterArgs
	state     *machineLearningInferenceClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MachineLearningInferenceCluster].
func (mlic *MachineLearningInferenceCluster) Type() string {
	return "azurerm_machine_learning_inference_cluster"
}

// LocalName returns the local name for [MachineLearningInferenceCluster].
func (mlic *MachineLearningInferenceCluster) LocalName() string {
	return mlic.Name
}

// Configuration returns the configuration (args) for [MachineLearningInferenceCluster].
func (mlic *MachineLearningInferenceCluster) Configuration() interface{} {
	return mlic.Args
}

// DependOn is used for other resources to depend on [MachineLearningInferenceCluster].
func (mlic *MachineLearningInferenceCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(mlic)
}

// Dependencies returns the list of resources [MachineLearningInferenceCluster] depends_on.
func (mlic *MachineLearningInferenceCluster) Dependencies() terra.Dependencies {
	return mlic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MachineLearningInferenceCluster].
func (mlic *MachineLearningInferenceCluster) LifecycleManagement() *terra.Lifecycle {
	return mlic.Lifecycle
}

// Attributes returns the attributes for [MachineLearningInferenceCluster].
func (mlic *MachineLearningInferenceCluster) Attributes() machineLearningInferenceClusterAttributes {
	return machineLearningInferenceClusterAttributes{ref: terra.ReferenceResource(mlic)}
}

// ImportState imports the given attribute values into [MachineLearningInferenceCluster]'s state.
func (mlic *MachineLearningInferenceCluster) ImportState(av io.Reader) error {
	mlic.state = &machineLearningInferenceClusterState{}
	if err := json.NewDecoder(av).Decode(mlic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mlic.Type(), mlic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MachineLearningInferenceCluster] has state.
func (mlic *MachineLearningInferenceCluster) State() (*machineLearningInferenceClusterState, bool) {
	return mlic.state, mlic.state != nil
}

// StateMust returns the state for [MachineLearningInferenceCluster]. Panics if the state is nil.
func (mlic *MachineLearningInferenceCluster) StateMust() *machineLearningInferenceClusterState {
	if mlic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mlic.Type(), mlic.LocalName()))
	}
	return mlic.state
}

// MachineLearningInferenceClusterArgs contains the configurations for azurerm_machine_learning_inference_cluster.
type MachineLearningInferenceClusterArgs struct {
	// ClusterPurpose: string, optional
	ClusterPurpose terra.StringValue `hcl:"cluster_purpose,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KubernetesClusterId: string, required
	KubernetesClusterId terra.StringValue `hcl:"kubernetes_cluster_id,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MachineLearningWorkspaceId: string, required
	MachineLearningWorkspaceId terra.StringValue `hcl:"machine_learning_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *machinelearninginferencecluster.Identity `hcl:"identity,block"`
	// Ssl: optional
	Ssl *machinelearninginferencecluster.Ssl `hcl:"ssl,block"`
	// Timeouts: optional
	Timeouts *machinelearninginferencecluster.Timeouts `hcl:"timeouts,block"`
}
type machineLearningInferenceClusterAttributes struct {
	ref terra.Reference
}

// ClusterPurpose returns a reference to field cluster_purpose of azurerm_machine_learning_inference_cluster.
func (mlic machineLearningInferenceClusterAttributes) ClusterPurpose() terra.StringValue {
	return terra.ReferenceAsString(mlic.ref.Append("cluster_purpose"))
}

// Description returns a reference to field description of azurerm_machine_learning_inference_cluster.
func (mlic machineLearningInferenceClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mlic.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_machine_learning_inference_cluster.
func (mlic machineLearningInferenceClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mlic.ref.Append("id"))
}

// KubernetesClusterId returns a reference to field kubernetes_cluster_id of azurerm_machine_learning_inference_cluster.
func (mlic machineLearningInferenceClusterAttributes) KubernetesClusterId() terra.StringValue {
	return terra.ReferenceAsString(mlic.ref.Append("kubernetes_cluster_id"))
}

// Location returns a reference to field location of azurerm_machine_learning_inference_cluster.
func (mlic machineLearningInferenceClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mlic.ref.Append("location"))
}

// MachineLearningWorkspaceId returns a reference to field machine_learning_workspace_id of azurerm_machine_learning_inference_cluster.
func (mlic machineLearningInferenceClusterAttributes) MachineLearningWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(mlic.ref.Append("machine_learning_workspace_id"))
}

// Name returns a reference to field name of azurerm_machine_learning_inference_cluster.
func (mlic machineLearningInferenceClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mlic.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_machine_learning_inference_cluster.
func (mlic machineLearningInferenceClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mlic.ref.Append("tags"))
}

func (mlic machineLearningInferenceClusterAttributes) Identity() terra.ListValue[machinelearninginferencecluster.IdentityAttributes] {
	return terra.ReferenceAsList[machinelearninginferencecluster.IdentityAttributes](mlic.ref.Append("identity"))
}

func (mlic machineLearningInferenceClusterAttributes) Ssl() terra.ListValue[machinelearninginferencecluster.SslAttributes] {
	return terra.ReferenceAsList[machinelearninginferencecluster.SslAttributes](mlic.ref.Append("ssl"))
}

func (mlic machineLearningInferenceClusterAttributes) Timeouts() machinelearninginferencecluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[machinelearninginferencecluster.TimeoutsAttributes](mlic.ref.Append("timeouts"))
}

type machineLearningInferenceClusterState struct {
	ClusterPurpose             string                                          `json:"cluster_purpose"`
	Description                string                                          `json:"description"`
	Id                         string                                          `json:"id"`
	KubernetesClusterId        string                                          `json:"kubernetes_cluster_id"`
	Location                   string                                          `json:"location"`
	MachineLearningWorkspaceId string                                          `json:"machine_learning_workspace_id"`
	Name                       string                                          `json:"name"`
	Tags                       map[string]string                               `json:"tags"`
	Identity                   []machinelearninginferencecluster.IdentityState `json:"identity"`
	Ssl                        []machinelearninginferencecluster.SslState      `json:"ssl"`
	Timeouts                   *machinelearninginferencecluster.TimeoutsState  `json:"timeouts"`
}
