// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opensearchoutboundconnection "github.com/golingon/terraproviders/aws/5.6.2/opensearchoutboundconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpensearchOutboundConnection creates a new instance of [OpensearchOutboundConnection].
func NewOpensearchOutboundConnection(name string, args OpensearchOutboundConnectionArgs) *OpensearchOutboundConnection {
	return &OpensearchOutboundConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpensearchOutboundConnection)(nil)

// OpensearchOutboundConnection represents the Terraform resource aws_opensearch_outbound_connection.
type OpensearchOutboundConnection struct {
	Name      string
	Args      OpensearchOutboundConnectionArgs
	state     *opensearchOutboundConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpensearchOutboundConnection].
func (ooc *OpensearchOutboundConnection) Type() string {
	return "aws_opensearch_outbound_connection"
}

// LocalName returns the local name for [OpensearchOutboundConnection].
func (ooc *OpensearchOutboundConnection) LocalName() string {
	return ooc.Name
}

// Configuration returns the configuration (args) for [OpensearchOutboundConnection].
func (ooc *OpensearchOutboundConnection) Configuration() interface{} {
	return ooc.Args
}

// DependOn is used for other resources to depend on [OpensearchOutboundConnection].
func (ooc *OpensearchOutboundConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(ooc)
}

// Dependencies returns the list of resources [OpensearchOutboundConnection] depends_on.
func (ooc *OpensearchOutboundConnection) Dependencies() terra.Dependencies {
	return ooc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpensearchOutboundConnection].
func (ooc *OpensearchOutboundConnection) LifecycleManagement() *terra.Lifecycle {
	return ooc.Lifecycle
}

// Attributes returns the attributes for [OpensearchOutboundConnection].
func (ooc *OpensearchOutboundConnection) Attributes() opensearchOutboundConnectionAttributes {
	return opensearchOutboundConnectionAttributes{ref: terra.ReferenceResource(ooc)}
}

// ImportState imports the given attribute values into [OpensearchOutboundConnection]'s state.
func (ooc *OpensearchOutboundConnection) ImportState(av io.Reader) error {
	ooc.state = &opensearchOutboundConnectionState{}
	if err := json.NewDecoder(av).Decode(ooc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ooc.Type(), ooc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpensearchOutboundConnection] has state.
func (ooc *OpensearchOutboundConnection) State() (*opensearchOutboundConnectionState, bool) {
	return ooc.state, ooc.state != nil
}

// StateMust returns the state for [OpensearchOutboundConnection]. Panics if the state is nil.
func (ooc *OpensearchOutboundConnection) StateMust() *opensearchOutboundConnectionState {
	if ooc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ooc.Type(), ooc.LocalName()))
	}
	return ooc.state
}

// OpensearchOutboundConnectionArgs contains the configurations for aws_opensearch_outbound_connection.
type OpensearchOutboundConnectionArgs struct {
	// ConnectionAlias: string, required
	ConnectionAlias terra.StringValue `hcl:"connection_alias,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalDomainInfo: required
	LocalDomainInfo *opensearchoutboundconnection.LocalDomainInfo `hcl:"local_domain_info,block" validate:"required"`
	// RemoteDomainInfo: required
	RemoteDomainInfo *opensearchoutboundconnection.RemoteDomainInfo `hcl:"remote_domain_info,block" validate:"required"`
	// Timeouts: optional
	Timeouts *opensearchoutboundconnection.Timeouts `hcl:"timeouts,block"`
}
type opensearchOutboundConnectionAttributes struct {
	ref terra.Reference
}

// ConnectionAlias returns a reference to field connection_alias of aws_opensearch_outbound_connection.
func (ooc opensearchOutboundConnectionAttributes) ConnectionAlias() terra.StringValue {
	return terra.ReferenceAsString(ooc.ref.Append("connection_alias"))
}

// ConnectionStatus returns a reference to field connection_status of aws_opensearch_outbound_connection.
func (ooc opensearchOutboundConnectionAttributes) ConnectionStatus() terra.StringValue {
	return terra.ReferenceAsString(ooc.ref.Append("connection_status"))
}

// Id returns a reference to field id of aws_opensearch_outbound_connection.
func (ooc opensearchOutboundConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ooc.ref.Append("id"))
}

func (ooc opensearchOutboundConnectionAttributes) LocalDomainInfo() terra.ListValue[opensearchoutboundconnection.LocalDomainInfoAttributes] {
	return terra.ReferenceAsList[opensearchoutboundconnection.LocalDomainInfoAttributes](ooc.ref.Append("local_domain_info"))
}

func (ooc opensearchOutboundConnectionAttributes) RemoteDomainInfo() terra.ListValue[opensearchoutboundconnection.RemoteDomainInfoAttributes] {
	return terra.ReferenceAsList[opensearchoutboundconnection.RemoteDomainInfoAttributes](ooc.ref.Append("remote_domain_info"))
}

func (ooc opensearchOutboundConnectionAttributes) Timeouts() opensearchoutboundconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[opensearchoutboundconnection.TimeoutsAttributes](ooc.ref.Append("timeouts"))
}

type opensearchOutboundConnectionState struct {
	ConnectionAlias  string                                               `json:"connection_alias"`
	ConnectionStatus string                                               `json:"connection_status"`
	Id               string                                               `json:"id"`
	LocalDomainInfo  []opensearchoutboundconnection.LocalDomainInfoState  `json:"local_domain_info"`
	RemoteDomainInfo []opensearchoutboundconnection.RemoteDomainInfoState `json:"remote_domain_info"`
	Timeouts         *opensearchoutboundconnection.TimeoutsState          `json:"timeouts"`
}
