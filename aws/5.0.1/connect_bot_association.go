// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	connectbotassociation "github.com/golingon/terraproviders/aws/5.0.1/connectbotassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConnectBotAssociation creates a new instance of [ConnectBotAssociation].
func NewConnectBotAssociation(name string, args ConnectBotAssociationArgs) *ConnectBotAssociation {
	return &ConnectBotAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectBotAssociation)(nil)

// ConnectBotAssociation represents the Terraform resource aws_connect_bot_association.
type ConnectBotAssociation struct {
	Name      string
	Args      ConnectBotAssociationArgs
	state     *connectBotAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectBotAssociation].
func (cba *ConnectBotAssociation) Type() string {
	return "aws_connect_bot_association"
}

// LocalName returns the local name for [ConnectBotAssociation].
func (cba *ConnectBotAssociation) LocalName() string {
	return cba.Name
}

// Configuration returns the configuration (args) for [ConnectBotAssociation].
func (cba *ConnectBotAssociation) Configuration() interface{} {
	return cba.Args
}

// DependOn is used for other resources to depend on [ConnectBotAssociation].
func (cba *ConnectBotAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(cba)
}

// Dependencies returns the list of resources [ConnectBotAssociation] depends_on.
func (cba *ConnectBotAssociation) Dependencies() terra.Dependencies {
	return cba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectBotAssociation].
func (cba *ConnectBotAssociation) LifecycleManagement() *terra.Lifecycle {
	return cba.Lifecycle
}

// Attributes returns the attributes for [ConnectBotAssociation].
func (cba *ConnectBotAssociation) Attributes() connectBotAssociationAttributes {
	return connectBotAssociationAttributes{ref: terra.ReferenceResource(cba)}
}

// ImportState imports the given attribute values into [ConnectBotAssociation]'s state.
func (cba *ConnectBotAssociation) ImportState(av io.Reader) error {
	cba.state = &connectBotAssociationState{}
	if err := json.NewDecoder(av).Decode(cba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cba.Type(), cba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectBotAssociation] has state.
func (cba *ConnectBotAssociation) State() (*connectBotAssociationState, bool) {
	return cba.state, cba.state != nil
}

// StateMust returns the state for [ConnectBotAssociation]. Panics if the state is nil.
func (cba *ConnectBotAssociation) StateMust() *connectBotAssociationState {
	if cba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cba.Type(), cba.LocalName()))
	}
	return cba.state
}

// ConnectBotAssociationArgs contains the configurations for aws_connect_bot_association.
type ConnectBotAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// LexBot: required
	LexBot *connectbotassociation.LexBot `hcl:"lex_bot,block" validate:"required"`
}
type connectBotAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_connect_bot_association.
func (cba connectBotAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cba.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_bot_association.
func (cba connectBotAssociationAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cba.ref.Append("instance_id"))
}

func (cba connectBotAssociationAttributes) LexBot() terra.ListValue[connectbotassociation.LexBotAttributes] {
	return terra.ReferenceAsList[connectbotassociation.LexBotAttributes](cba.ref.Append("lex_bot"))
}

type connectBotAssociationState struct {
	Id         string                              `json:"id"`
	InstanceId string                              `json:"instance_id"`
	LexBot     []connectbotassociation.LexBotState `json:"lex_bot"`
}
