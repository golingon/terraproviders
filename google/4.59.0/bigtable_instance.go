// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigtableinstance "github.com/golingon/terraproviders/google/4.59.0/bigtableinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewBigtableInstance(name string, args BigtableInstanceArgs) *BigtableInstance {
	return &BigtableInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableInstance)(nil)

type BigtableInstance struct {
	Name  string
	Args  BigtableInstanceArgs
	state *bigtableInstanceState
}

func (bi *BigtableInstance) Type() string {
	return "google_bigtable_instance"
}

func (bi *BigtableInstance) LocalName() string {
	return bi.Name
}

func (bi *BigtableInstance) Configuration() interface{} {
	return bi.Args
}

func (bi *BigtableInstance) Attributes() bigtableInstanceAttributes {
	return bigtableInstanceAttributes{ref: terra.ReferenceResource(bi)}
}

func (bi *BigtableInstance) ImportState(av io.Reader) error {
	bi.state = &bigtableInstanceState{}
	if err := json.NewDecoder(av).Decode(bi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bi.Type(), bi.LocalName(), err)
	}
	return nil
}

func (bi *BigtableInstance) State() (*bigtableInstanceState, bool) {
	return bi.state, bi.state != nil
}

func (bi *BigtableInstance) StateMust() *bigtableInstanceState {
	if bi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bi.Type(), bi.LocalName()))
	}
	return bi.state
}

func (bi *BigtableInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(bi)
}

type BigtableInstanceArgs struct {
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Cluster: min=0
	Cluster []bigtableinstance.Cluster `hcl:"cluster,block" validate:"min=0"`
	// DependsOn contains resources that BigtableInstance depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type bigtableInstanceAttributes struct {
	ref terra.Reference
}

func (bi bigtableInstanceAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceBool(bi.ref.Append("deletion_protection"))
}

func (bi bigtableInstanceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(bi.ref.Append("display_name"))
}

func (bi bigtableInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bi.ref.Append("id"))
}

func (bi bigtableInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceString(bi.ref.Append("instance_type"))
}

func (bi bigtableInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](bi.ref.Append("labels"))
}

func (bi bigtableInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(bi.ref.Append("name"))
}

func (bi bigtableInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceString(bi.ref.Append("project"))
}

func (bi bigtableInstanceAttributes) Cluster() terra.ListValue[bigtableinstance.ClusterAttributes] {
	return terra.ReferenceList[bigtableinstance.ClusterAttributes](bi.ref.Append("cluster"))
}

type bigtableInstanceState struct {
	DeletionProtection bool                            `json:"deletion_protection"`
	DisplayName        string                          `json:"display_name"`
	Id                 string                          `json:"id"`
	InstanceType       string                          `json:"instance_type"`
	Labels             map[string]string               `json:"labels"`
	Name               string                          `json:"name"`
	Project            string                          `json:"project"`
	Cluster            []bigtableinstance.ClusterState `json:"cluster"`
}
