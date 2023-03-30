// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	emrcontainersvirtualcluster "github.com/golingon/terraproviders/aws/4.60.0/emrcontainersvirtualcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEmrcontainersVirtualCluster(name string, args EmrcontainersVirtualClusterArgs) *EmrcontainersVirtualCluster {
	return &EmrcontainersVirtualCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EmrcontainersVirtualCluster)(nil)

type EmrcontainersVirtualCluster struct {
	Name  string
	Args  EmrcontainersVirtualClusterArgs
	state *emrcontainersVirtualClusterState
}

func (evc *EmrcontainersVirtualCluster) Type() string {
	return "aws_emrcontainers_virtual_cluster"
}

func (evc *EmrcontainersVirtualCluster) LocalName() string {
	return evc.Name
}

func (evc *EmrcontainersVirtualCluster) Configuration() interface{} {
	return evc.Args
}

func (evc *EmrcontainersVirtualCluster) Attributes() emrcontainersVirtualClusterAttributes {
	return emrcontainersVirtualClusterAttributes{ref: terra.ReferenceResource(evc)}
}

func (evc *EmrcontainersVirtualCluster) ImportState(av io.Reader) error {
	evc.state = &emrcontainersVirtualClusterState{}
	if err := json.NewDecoder(av).Decode(evc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", evc.Type(), evc.LocalName(), err)
	}
	return nil
}

func (evc *EmrcontainersVirtualCluster) State() (*emrcontainersVirtualClusterState, bool) {
	return evc.state, evc.state != nil
}

func (evc *EmrcontainersVirtualCluster) StateMust() *emrcontainersVirtualClusterState {
	if evc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", evc.Type(), evc.LocalName()))
	}
	return evc.state
}

func (evc *EmrcontainersVirtualCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(evc)
}

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
	// DependsOn contains resources that EmrcontainersVirtualCluster depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type emrcontainersVirtualClusterAttributes struct {
	ref terra.Reference
}

func (evc emrcontainersVirtualClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(evc.ref.Append("arn"))
}

func (evc emrcontainersVirtualClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(evc.ref.Append("id"))
}

func (evc emrcontainersVirtualClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceString(evc.ref.Append("name"))
}

func (evc emrcontainersVirtualClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](evc.ref.Append("tags"))
}

func (evc emrcontainersVirtualClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](evc.ref.Append("tags_all"))
}

func (evc emrcontainersVirtualClusterAttributes) ContainerProvider() terra.ListValue[emrcontainersvirtualcluster.ContainerProviderAttributes] {
	return terra.ReferenceList[emrcontainersvirtualcluster.ContainerProviderAttributes](evc.ref.Append("container_provider"))
}

func (evc emrcontainersVirtualClusterAttributes) Timeouts() emrcontainersvirtualcluster.TimeoutsAttributes {
	return terra.ReferenceSingle[emrcontainersvirtualcluster.TimeoutsAttributes](evc.ref.Append("timeouts"))
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
