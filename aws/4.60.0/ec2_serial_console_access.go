// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEc2SerialConsoleAccess(name string, args Ec2SerialConsoleAccessArgs) *Ec2SerialConsoleAccess {
	return &Ec2SerialConsoleAccess{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2SerialConsoleAccess)(nil)

type Ec2SerialConsoleAccess struct {
	Name  string
	Args  Ec2SerialConsoleAccessArgs
	state *ec2SerialConsoleAccessState
}

func (esca *Ec2SerialConsoleAccess) Type() string {
	return "aws_ec2_serial_console_access"
}

func (esca *Ec2SerialConsoleAccess) LocalName() string {
	return esca.Name
}

func (esca *Ec2SerialConsoleAccess) Configuration() interface{} {
	return esca.Args
}

func (esca *Ec2SerialConsoleAccess) Attributes() ec2SerialConsoleAccessAttributes {
	return ec2SerialConsoleAccessAttributes{ref: terra.ReferenceResource(esca)}
}

func (esca *Ec2SerialConsoleAccess) ImportState(av io.Reader) error {
	esca.state = &ec2SerialConsoleAccessState{}
	if err := json.NewDecoder(av).Decode(esca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", esca.Type(), esca.LocalName(), err)
	}
	return nil
}

func (esca *Ec2SerialConsoleAccess) State() (*ec2SerialConsoleAccessState, bool) {
	return esca.state, esca.state != nil
}

func (esca *Ec2SerialConsoleAccess) StateMust() *ec2SerialConsoleAccessState {
	if esca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", esca.Type(), esca.LocalName()))
	}
	return esca.state
}

func (esca *Ec2SerialConsoleAccess) DependOn() terra.Reference {
	return terra.ReferenceResource(esca)
}

type Ec2SerialConsoleAccessArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DependsOn contains resources that Ec2SerialConsoleAccess depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type ec2SerialConsoleAccessAttributes struct {
	ref terra.Reference
}

func (esca ec2SerialConsoleAccessAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(esca.ref.Append("enabled"))
}

func (esca ec2SerialConsoleAccessAttributes) Id() terra.StringValue {
	return terra.ReferenceString(esca.ref.Append("id"))
}

type ec2SerialConsoleAccessState struct {
	Enabled bool   `json:"enabled"`
	Id      string `json:"id"`
}
