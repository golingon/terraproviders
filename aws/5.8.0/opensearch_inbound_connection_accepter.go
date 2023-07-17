// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opensearchinboundconnectionaccepter "github.com/golingon/terraproviders/aws/5.8.0/opensearchinboundconnectionaccepter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpensearchInboundConnectionAccepter creates a new instance of [OpensearchInboundConnectionAccepter].
func NewOpensearchInboundConnectionAccepter(name string, args OpensearchInboundConnectionAccepterArgs) *OpensearchInboundConnectionAccepter {
	return &OpensearchInboundConnectionAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpensearchInboundConnectionAccepter)(nil)

// OpensearchInboundConnectionAccepter represents the Terraform resource aws_opensearch_inbound_connection_accepter.
type OpensearchInboundConnectionAccepter struct {
	Name      string
	Args      OpensearchInboundConnectionAccepterArgs
	state     *opensearchInboundConnectionAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpensearchInboundConnectionAccepter].
func (oica *OpensearchInboundConnectionAccepter) Type() string {
	return "aws_opensearch_inbound_connection_accepter"
}

// LocalName returns the local name for [OpensearchInboundConnectionAccepter].
func (oica *OpensearchInboundConnectionAccepter) LocalName() string {
	return oica.Name
}

// Configuration returns the configuration (args) for [OpensearchInboundConnectionAccepter].
func (oica *OpensearchInboundConnectionAccepter) Configuration() interface{} {
	return oica.Args
}

// DependOn is used for other resources to depend on [OpensearchInboundConnectionAccepter].
func (oica *OpensearchInboundConnectionAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(oica)
}

// Dependencies returns the list of resources [OpensearchInboundConnectionAccepter] depends_on.
func (oica *OpensearchInboundConnectionAccepter) Dependencies() terra.Dependencies {
	return oica.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpensearchInboundConnectionAccepter].
func (oica *OpensearchInboundConnectionAccepter) LifecycleManagement() *terra.Lifecycle {
	return oica.Lifecycle
}

// Attributes returns the attributes for [OpensearchInboundConnectionAccepter].
func (oica *OpensearchInboundConnectionAccepter) Attributes() opensearchInboundConnectionAccepterAttributes {
	return opensearchInboundConnectionAccepterAttributes{ref: terra.ReferenceResource(oica)}
}

// ImportState imports the given attribute values into [OpensearchInboundConnectionAccepter]'s state.
func (oica *OpensearchInboundConnectionAccepter) ImportState(av io.Reader) error {
	oica.state = &opensearchInboundConnectionAccepterState{}
	if err := json.NewDecoder(av).Decode(oica.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oica.Type(), oica.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpensearchInboundConnectionAccepter] has state.
func (oica *OpensearchInboundConnectionAccepter) State() (*opensearchInboundConnectionAccepterState, bool) {
	return oica.state, oica.state != nil
}

// StateMust returns the state for [OpensearchInboundConnectionAccepter]. Panics if the state is nil.
func (oica *OpensearchInboundConnectionAccepter) StateMust() *opensearchInboundConnectionAccepterState {
	if oica.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oica.Type(), oica.LocalName()))
	}
	return oica.state
}

// OpensearchInboundConnectionAccepterArgs contains the configurations for aws_opensearch_inbound_connection_accepter.
type OpensearchInboundConnectionAccepterArgs struct {
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *opensearchinboundconnectionaccepter.Timeouts `hcl:"timeouts,block"`
}
type opensearchInboundConnectionAccepterAttributes struct {
	ref terra.Reference
}

// ConnectionId returns a reference to field connection_id of aws_opensearch_inbound_connection_accepter.
func (oica opensearchInboundConnectionAccepterAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(oica.ref.Append("connection_id"))
}

// ConnectionStatus returns a reference to field connection_status of aws_opensearch_inbound_connection_accepter.
func (oica opensearchInboundConnectionAccepterAttributes) ConnectionStatus() terra.StringValue {
	return terra.ReferenceAsString(oica.ref.Append("connection_status"))
}

// Id returns a reference to field id of aws_opensearch_inbound_connection_accepter.
func (oica opensearchInboundConnectionAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oica.ref.Append("id"))
}

func (oica opensearchInboundConnectionAccepterAttributes) Timeouts() opensearchinboundconnectionaccepter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[opensearchinboundconnectionaccepter.TimeoutsAttributes](oica.ref.Append("timeouts"))
}

type opensearchInboundConnectionAccepterState struct {
	ConnectionId     string                                             `json:"connection_id"`
	ConnectionStatus string                                             `json:"connection_status"`
	Id               string                                             `json:"id"`
	Timeouts         *opensearchinboundconnectionaccepter.TimeoutsState `json:"timeouts"`
}
