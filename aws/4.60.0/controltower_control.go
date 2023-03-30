// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	controltowercontrol "github.com/golingon/terraproviders/aws/4.60.0/controltowercontrol"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewControltowerControl(name string, args ControltowerControlArgs) *ControltowerControl {
	return &ControltowerControl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ControltowerControl)(nil)

type ControltowerControl struct {
	Name  string
	Args  ControltowerControlArgs
	state *controltowerControlState
}

func (cc *ControltowerControl) Type() string {
	return "aws_controltower_control"
}

func (cc *ControltowerControl) LocalName() string {
	return cc.Name
}

func (cc *ControltowerControl) Configuration() interface{} {
	return cc.Args
}

func (cc *ControltowerControl) Attributes() controltowerControlAttributes {
	return controltowerControlAttributes{ref: terra.ReferenceResource(cc)}
}

func (cc *ControltowerControl) ImportState(av io.Reader) error {
	cc.state = &controltowerControlState{}
	if err := json.NewDecoder(av).Decode(cc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cc.Type(), cc.LocalName(), err)
	}
	return nil
}

func (cc *ControltowerControl) State() (*controltowerControlState, bool) {
	return cc.state, cc.state != nil
}

func (cc *ControltowerControl) StateMust() *controltowerControlState {
	if cc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cc.Type(), cc.LocalName()))
	}
	return cc.state
}

func (cc *ControltowerControl) DependOn() terra.Reference {
	return terra.ReferenceResource(cc)
}

type ControltowerControlArgs struct {
	// ControlIdentifier: string, required
	ControlIdentifier terra.StringValue `hcl:"control_identifier,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TargetIdentifier: string, required
	TargetIdentifier terra.StringValue `hcl:"target_identifier,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *controltowercontrol.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ControltowerControl depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type controltowerControlAttributes struct {
	ref terra.Reference
}

func (cc controltowerControlAttributes) ControlIdentifier() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("control_identifier"))
}

func (cc controltowerControlAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("id"))
}

func (cc controltowerControlAttributes) TargetIdentifier() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("target_identifier"))
}

func (cc controltowerControlAttributes) Timeouts() controltowercontrol.TimeoutsAttributes {
	return terra.ReferenceSingle[controltowercontrol.TimeoutsAttributes](cc.ref.Append("timeouts"))
}

type controltowerControlState struct {
	ControlIdentifier string                             `json:"control_identifier"`
	Id                string                             `json:"id"`
	TargetIdentifier  string                             `json:"target_identifier"`
	Timeouts          *controltowercontrol.TimeoutsState `json:"timeouts"`
}
