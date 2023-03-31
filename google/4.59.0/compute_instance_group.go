// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeinstancegroup "github.com/golingon/terraproviders/google/4.59.0/computeinstancegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeInstanceGroup(name string, args ComputeInstanceGroupArgs) *ComputeInstanceGroup {
	return &ComputeInstanceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstanceGroup)(nil)

type ComputeInstanceGroup struct {
	Name  string
	Args  ComputeInstanceGroupArgs
	state *computeInstanceGroupState
}

func (cig *ComputeInstanceGroup) Type() string {
	return "google_compute_instance_group"
}

func (cig *ComputeInstanceGroup) LocalName() string {
	return cig.Name
}

func (cig *ComputeInstanceGroup) Configuration() interface{} {
	return cig.Args
}

func (cig *ComputeInstanceGroup) Attributes() computeInstanceGroupAttributes {
	return computeInstanceGroupAttributes{ref: terra.ReferenceResource(cig)}
}

func (cig *ComputeInstanceGroup) ImportState(av io.Reader) error {
	cig.state = &computeInstanceGroupState{}
	if err := json.NewDecoder(av).Decode(cig.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cig.Type(), cig.LocalName(), err)
	}
	return nil
}

func (cig *ComputeInstanceGroup) State() (*computeInstanceGroupState, bool) {
	return cig.state, cig.state != nil
}

func (cig *ComputeInstanceGroup) StateMust() *computeInstanceGroupState {
	if cig.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cig.Type(), cig.LocalName()))
	}
	return cig.state
}

func (cig *ComputeInstanceGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cig)
}

type ComputeInstanceGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instances: set of string, optional
	Instances terra.SetValue[terra.StringValue] `hcl:"instances,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// NamedPort: min=0
	NamedPort []computeinstancegroup.NamedPort `hcl:"named_port,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeinstancegroup.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ComputeInstanceGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeInstanceGroupAttributes struct {
	ref terra.Reference
}

func (cig computeInstanceGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cig.ref.Append("description"))
}

func (cig computeInstanceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cig.ref.Append("id"))
}

func (cig computeInstanceGroupAttributes) Instances() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cig.ref.Append("instances"))
}

func (cig computeInstanceGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cig.ref.Append("name"))
}

func (cig computeInstanceGroupAttributes) Network() terra.StringValue {
	return terra.ReferenceString(cig.ref.Append("network"))
}

func (cig computeInstanceGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cig.ref.Append("project"))
}

func (cig computeInstanceGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(cig.ref.Append("self_link"))
}

func (cig computeInstanceGroupAttributes) Size() terra.NumberValue {
	return terra.ReferenceNumber(cig.ref.Append("size"))
}

func (cig computeInstanceGroupAttributes) Zone() terra.StringValue {
	return terra.ReferenceString(cig.ref.Append("zone"))
}

func (cig computeInstanceGroupAttributes) NamedPort() terra.ListValue[computeinstancegroup.NamedPortAttributes] {
	return terra.ReferenceList[computeinstancegroup.NamedPortAttributes](cig.ref.Append("named_port"))
}

func (cig computeInstanceGroupAttributes) Timeouts() computeinstancegroup.TimeoutsAttributes {
	return terra.ReferenceSingle[computeinstancegroup.TimeoutsAttributes](cig.ref.Append("timeouts"))
}

type computeInstanceGroupState struct {
	Description string                                `json:"description"`
	Id          string                                `json:"id"`
	Instances   []string                              `json:"instances"`
	Name        string                                `json:"name"`
	Network     string                                `json:"network"`
	Project     string                                `json:"project"`
	SelfLink    string                                `json:"self_link"`
	Size        float64                               `json:"size"`
	Zone        string                                `json:"zone"`
	NamedPort   []computeinstancegroup.NamedPortState `json:"named_port"`
	Timeouts    *computeinstancegroup.TimeoutsState   `json:"timeouts"`
}
