// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opensearchinboundconnectionaccepter "github.com/golingon/terraproviders/aws/4.60.0/opensearchinboundconnectionaccepter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewOpensearchInboundConnectionAccepter(name string, args OpensearchInboundConnectionAccepterArgs) *OpensearchInboundConnectionAccepter {
	return &OpensearchInboundConnectionAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpensearchInboundConnectionAccepter)(nil)

type OpensearchInboundConnectionAccepter struct {
	Name  string
	Args  OpensearchInboundConnectionAccepterArgs
	state *opensearchInboundConnectionAccepterState
}

func (oica *OpensearchInboundConnectionAccepter) Type() string {
	return "aws_opensearch_inbound_connection_accepter"
}

func (oica *OpensearchInboundConnectionAccepter) LocalName() string {
	return oica.Name
}

func (oica *OpensearchInboundConnectionAccepter) Configuration() interface{} {
	return oica.Args
}

func (oica *OpensearchInboundConnectionAccepter) Attributes() opensearchInboundConnectionAccepterAttributes {
	return opensearchInboundConnectionAccepterAttributes{ref: terra.ReferenceResource(oica)}
}

func (oica *OpensearchInboundConnectionAccepter) ImportState(av io.Reader) error {
	oica.state = &opensearchInboundConnectionAccepterState{}
	if err := json.NewDecoder(av).Decode(oica.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oica.Type(), oica.LocalName(), err)
	}
	return nil
}

func (oica *OpensearchInboundConnectionAccepter) State() (*opensearchInboundConnectionAccepterState, bool) {
	return oica.state, oica.state != nil
}

func (oica *OpensearchInboundConnectionAccepter) StateMust() *opensearchInboundConnectionAccepterState {
	if oica.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oica.Type(), oica.LocalName()))
	}
	return oica.state
}

func (oica *OpensearchInboundConnectionAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(oica)
}

type OpensearchInboundConnectionAccepterArgs struct {
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *opensearchinboundconnectionaccepter.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that OpensearchInboundConnectionAccepter depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type opensearchInboundConnectionAccepterAttributes struct {
	ref terra.Reference
}

func (oica opensearchInboundConnectionAccepterAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceString(oica.ref.Append("connection_id"))
}

func (oica opensearchInboundConnectionAccepterAttributes) ConnectionStatus() terra.StringValue {
	return terra.ReferenceString(oica.ref.Append("connection_status"))
}

func (oica opensearchInboundConnectionAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(oica.ref.Append("id"))
}

func (oica opensearchInboundConnectionAccepterAttributes) Timeouts() opensearchinboundconnectionaccepter.TimeoutsAttributes {
	return terra.ReferenceSingle[opensearchinboundconnectionaccepter.TimeoutsAttributes](oica.ref.Append("timeouts"))
}

type opensearchInboundConnectionAccepterState struct {
	ConnectionId     string                                             `json:"connection_id"`
	ConnectionStatus string                                             `json:"connection_status"`
	Id               string                                             `json:"id"`
	Timeouts         *opensearchinboundconnectionaccepter.TimeoutsState `json:"timeouts"`
}
