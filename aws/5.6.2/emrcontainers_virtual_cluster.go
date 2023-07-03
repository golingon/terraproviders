// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	emrcontainersvirtualcluster "github.com/golingon/terraproviders/aws/5.6.2/emrcontainersvirtualcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEmrcontainersVirtualCluster creates a new instance of [EmrcontainersVirtualCluster].
func NewEmrcontainersVirtualCluster(name string, args EmrcontainersVirtualClusterArgs) *EmrcontainersVirtualCluster {
	return &EmrcontainersVirtualCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EmrcontainersVirtualCluster)(nil)

// EmrcontainersVirtualCluster represents the Terraform resource aws_emrcontainers_virtual_cluster.
type EmrcontainersVirtualCluster struct {
	Name      string
	Args      EmrcontainersVirtualClusterArgs
	state     *emrcontainersVirtualClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EmrcontainersVirtualCluster].
func (evc *EmrcontainersVirtualCluster) Type() string {
	return "aws_emrcontainers_virtual_cluster"
}

// LocalName returns the local name for [EmrcontainersVirtualCluster].
func (evc *EmrcontainersVirtualCluster) LocalName() string {
	return evc.Name
}

// Configuration returns the configuration (args) for [EmrcontainersVirtualCluster].
func (evc *EmrcontainersVirtualCluster) Configuration() interface{} {
	return evc.Args
}

// DependOn is used for other resources to depend on [EmrcontainersVirtualCluster].
func (evc *EmrcontainersVirtualCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(evc)
}

// Dependencies returns the list of resources [EmrcontainersVirtualCluster] depends_on.
func (evc *EmrcontainersVirtualCluster) Dependencies() terra.Dependencies {
	return evc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EmrcontainersVirtualCluster].
func (evc *EmrcontainersVirtualCluster) LifecycleManagement() *terra.Lifecycle {
	return evc.Lifecycle
}

// Attributes returns the attributes for [EmrcontainersVirtualCluster].
func (evc *EmrcontainersVirtualCluster) Attributes() emrcontainersVirtualClusterAttributes {
	return emrcontainersVirtualClusterAttributes{ref: terra.ReferenceResource(evc)}
}

// ImportState imports the given attribute values into [EmrcontainersVirtualCluster]'s state.
func (evc *EmrcontainersVirtualCluster) ImportState(av io.Reader) error {
	evc.state = &emrcontainersVirtualClusterState{}
	if err := json.NewDecoder(av).Decode(evc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", evc.Type(), evc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EmrcontainersVirtualCluster] has state.
func (evc *EmrcontainersVirtualCluster) State() (*emrcontainersVirtualClusterState, bool) {
	return evc.state, evc.state != nil
}

// StateMust returns the state for [EmrcontainersVirtualCluster]. Panics if the state is nil.
func (evc *EmrcontainersVirtualCluster) StateMust() *emrcontainersVirtualClusterState {
	if evc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", evc.Type(), evc.LocalName()))
	}
	return evc.state
}

// EmrcontainersVirtualClusterArgs contains the configurations for aws_emrcontainers_virtual_cluster.
type EmrcontainersVirtualClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ContainerProvider: required
	ContainerProvider *emrcontainersvirtualcluster.ContainerProvider `hcl:"container_provider,block" validate:"required"`
	// Timeouts: optional
	Timeouts *emrcontainersvirtualcluster.Timeouts `hcl:"timeouts,block"`
}
type emrcontainersVirtualClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_emrcontainers_virtual_cluster.
func (evc emrcontainersVirtualClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("arn"))
}

// Id returns a reference to field id of aws_emrcontainers_virtual_cluster.
func (evc emrcontainersVirtualClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("id"))
}

// Name returns a reference to field name of aws_emrcontainers_virtual_cluster.
func (evc emrcontainersVirtualClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(evc.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_emrcontainers_virtual_cluster.
func (evc emrcontainersVirtualClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](evc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_emrcontainers_virtual_cluster.
func (evc emrcontainersVirtualClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](evc.ref.Append("tags_all"))
}

func (evc emrcontainersVirtualClusterAttributes) ContainerProvider() terra.ListValue[emrcontainersvirtualcluster.ContainerProviderAttributes] {
	return terra.ReferenceAsList[emrcontainersvirtualcluster.ContainerProviderAttributes](evc.ref.Append("container_provider"))
}

func (evc emrcontainersVirtualClusterAttributes) Timeouts() emrcontainersvirtualcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[emrcontainersvirtualcluster.TimeoutsAttributes](evc.ref.Append("timeouts"))
}

type emrcontainersVirtualClusterState struct {
	Arn               string                                               `json:"arn"`
	Id                string                                               `json:"id"`
	Name              string                                               `json:"name"`
	Tags              map[string]string                                    `json:"tags"`
	TagsAll           map[string]string                                    `json:"tags_all"`
	ContainerProvider []emrcontainersvirtualcluster.ContainerProviderState `json:"container_provider"`
	Timeouts          *emrcontainersvirtualcluster.TimeoutsState           `json:"timeouts"`
}
