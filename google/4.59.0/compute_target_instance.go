// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computetargetinstance "github.com/golingon/terraproviders/google/4.59.0/computetargetinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeTargetInstance(name string, args ComputeTargetInstanceArgs) *ComputeTargetInstance {
	return &ComputeTargetInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeTargetInstance)(nil)

type ComputeTargetInstance struct {
	Name  string
	Args  ComputeTargetInstanceArgs
	state *computeTargetInstanceState
}

func (cti *ComputeTargetInstance) Type() string {
	return "google_compute_target_instance"
}

func (cti *ComputeTargetInstance) LocalName() string {
	return cti.Name
}

func (cti *ComputeTargetInstance) Configuration() interface{} {
	return cti.Args
}

func (cti *ComputeTargetInstance) Attributes() computeTargetInstanceAttributes {
	return computeTargetInstanceAttributes{ref: terra.ReferenceResource(cti)}
}

func (cti *ComputeTargetInstance) ImportState(av io.Reader) error {
	cti.state = &computeTargetInstanceState{}
	if err := json.NewDecoder(av).Decode(cti.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cti.Type(), cti.LocalName(), err)
	}
	return nil
}

func (cti *ComputeTargetInstance) State() (*computeTargetInstanceState, bool) {
	return cti.state, cti.state != nil
}

func (cti *ComputeTargetInstance) StateMust() *computeTargetInstanceState {
	if cti.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cti.Type(), cti.LocalName()))
	}
	return cti.state
}

func (cti *ComputeTargetInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(cti)
}

type ComputeTargetInstanceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NatPolicy: string, optional
	NatPolicy terra.StringValue `hcl:"nat_policy,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *computetargetinstance.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ComputeTargetInstance depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeTargetInstanceAttributes struct {
	ref terra.Reference
}

func (cti computeTargetInstanceAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceString(cti.ref.Append("creation_timestamp"))
}

func (cti computeTargetInstanceAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cti.ref.Append("description"))
}

func (cti computeTargetInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cti.ref.Append("id"))
}

func (cti computeTargetInstanceAttributes) Instance() terra.StringValue {
	return terra.ReferenceString(cti.ref.Append("instance"))
}

func (cti computeTargetInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cti.ref.Append("name"))
}

func (cti computeTargetInstanceAttributes) NatPolicy() terra.StringValue {
	return terra.ReferenceString(cti.ref.Append("nat_policy"))
}

func (cti computeTargetInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cti.ref.Append("project"))
}

func (cti computeTargetInstanceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(cti.ref.Append("self_link"))
}

func (cti computeTargetInstanceAttributes) Zone() terra.StringValue {
	return terra.ReferenceString(cti.ref.Append("zone"))
}

func (cti computeTargetInstanceAttributes) Timeouts() computetargetinstance.TimeoutsAttributes {
	return terra.ReferenceSingle[computetargetinstance.TimeoutsAttributes](cti.ref.Append("timeouts"))
}

type computeTargetInstanceState struct {
	CreationTimestamp string                               `json:"creation_timestamp"`
	Description       string                               `json:"description"`
	Id                string                               `json:"id"`
	Instance          string                               `json:"instance"`
	Name              string                               `json:"name"`
	NatPolicy         string                               `json:"nat_policy"`
	Project           string                               `json:"project"`
	SelfLink          string                               `json:"self_link"`
	Zone              string                               `json:"zone"`
	Timeouts          *computetargetinstance.TimeoutsState `json:"timeouts"`
}
