// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_opensearch_outbound_connection

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_opensearch_outbound_connection.
type Resource struct {
	Name      string
	Args      Args
	state     *awsOpensearchOutboundConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aooc *Resource) Type() string {
	return "aws_opensearch_outbound_connection"
}

// LocalName returns the local name for [Resource].
func (aooc *Resource) LocalName() string {
	return aooc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aooc *Resource) Configuration() interface{} {
	return aooc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aooc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aooc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aooc *Resource) Dependencies() terra.Dependencies {
	return aooc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aooc *Resource) LifecycleManagement() *terra.Lifecycle {
	return aooc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aooc *Resource) Attributes() awsOpensearchOutboundConnectionAttributes {
	return awsOpensearchOutboundConnectionAttributes{ref: terra.ReferenceResource(aooc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aooc *Resource) ImportState(state io.Reader) error {
	aooc.state = &awsOpensearchOutboundConnectionState{}
	if err := json.NewDecoder(state).Decode(aooc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aooc.Type(), aooc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aooc *Resource) State() (*awsOpensearchOutboundConnectionState, bool) {
	return aooc.state, aooc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aooc *Resource) StateMust() *awsOpensearchOutboundConnectionState {
	if aooc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aooc.Type(), aooc.LocalName()))
	}
	return aooc.state
}

// Args contains the configurations for aws_opensearch_outbound_connection.
type Args struct {
	// AcceptConnection: bool, optional
	AcceptConnection terra.BoolValue `hcl:"accept_connection,attr"`
	// ConnectionAlias: string, required
	ConnectionAlias terra.StringValue `hcl:"connection_alias,attr" validate:"required"`
	// ConnectionMode: string, optional
	ConnectionMode terra.StringValue `hcl:"connection_mode,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ConnectionProperties: optional
	ConnectionProperties *ConnectionProperties `hcl:"connection_properties,block"`
	// LocalDomainInfo: required
	LocalDomainInfo *LocalDomainInfo `hcl:"local_domain_info,block" validate:"required"`
	// RemoteDomainInfo: required
	RemoteDomainInfo *RemoteDomainInfo `hcl:"remote_domain_info,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsOpensearchOutboundConnectionAttributes struct {
	ref terra.Reference
}

// AcceptConnection returns a reference to field accept_connection of aws_opensearch_outbound_connection.
func (aooc awsOpensearchOutboundConnectionAttributes) AcceptConnection() terra.BoolValue {
	return terra.ReferenceAsBool(aooc.ref.Append("accept_connection"))
}

// ConnectionAlias returns a reference to field connection_alias of aws_opensearch_outbound_connection.
func (aooc awsOpensearchOutboundConnectionAttributes) ConnectionAlias() terra.StringValue {
	return terra.ReferenceAsString(aooc.ref.Append("connection_alias"))
}

// ConnectionMode returns a reference to field connection_mode of aws_opensearch_outbound_connection.
func (aooc awsOpensearchOutboundConnectionAttributes) ConnectionMode() terra.StringValue {
	return terra.ReferenceAsString(aooc.ref.Append("connection_mode"))
}

// ConnectionStatus returns a reference to field connection_status of aws_opensearch_outbound_connection.
func (aooc awsOpensearchOutboundConnectionAttributes) ConnectionStatus() terra.StringValue {
	return terra.ReferenceAsString(aooc.ref.Append("connection_status"))
}

// Id returns a reference to field id of aws_opensearch_outbound_connection.
func (aooc awsOpensearchOutboundConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aooc.ref.Append("id"))
}

func (aooc awsOpensearchOutboundConnectionAttributes) ConnectionProperties() terra.ListValue[ConnectionPropertiesAttributes] {
	return terra.ReferenceAsList[ConnectionPropertiesAttributes](aooc.ref.Append("connection_properties"))
}

func (aooc awsOpensearchOutboundConnectionAttributes) LocalDomainInfo() terra.ListValue[LocalDomainInfoAttributes] {
	return terra.ReferenceAsList[LocalDomainInfoAttributes](aooc.ref.Append("local_domain_info"))
}

func (aooc awsOpensearchOutboundConnectionAttributes) RemoteDomainInfo() terra.ListValue[RemoteDomainInfoAttributes] {
	return terra.ReferenceAsList[RemoteDomainInfoAttributes](aooc.ref.Append("remote_domain_info"))
}

func (aooc awsOpensearchOutboundConnectionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aooc.ref.Append("timeouts"))
}

type awsOpensearchOutboundConnectionState struct {
	AcceptConnection     bool                        `json:"accept_connection"`
	ConnectionAlias      string                      `json:"connection_alias"`
	ConnectionMode       string                      `json:"connection_mode"`
	ConnectionStatus     string                      `json:"connection_status"`
	Id                   string                      `json:"id"`
	ConnectionProperties []ConnectionPropertiesState `json:"connection_properties"`
	LocalDomainInfo      []LocalDomainInfoState      `json:"local_domain_info"`
	RemoteDomainInfo     []RemoteDomainInfoState     `json:"remote_domain_info"`
	Timeouts             *TimeoutsState              `json:"timeouts"`
}
