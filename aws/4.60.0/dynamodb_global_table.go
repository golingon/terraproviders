// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dynamodbglobaltable "github.com/golingon/terraproviders/aws/4.60.0/dynamodbglobaltable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDynamodbGlobalTable(name string, args DynamodbGlobalTableArgs) *DynamodbGlobalTable {
	return &DynamodbGlobalTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DynamodbGlobalTable)(nil)

type DynamodbGlobalTable struct {
	Name  string
	Args  DynamodbGlobalTableArgs
	state *dynamodbGlobalTableState
}

func (dgt *DynamodbGlobalTable) Type() string {
	return "aws_dynamodb_global_table"
}

func (dgt *DynamodbGlobalTable) LocalName() string {
	return dgt.Name
}

func (dgt *DynamodbGlobalTable) Configuration() interface{} {
	return dgt.Args
}

func (dgt *DynamodbGlobalTable) Attributes() dynamodbGlobalTableAttributes {
	return dynamodbGlobalTableAttributes{ref: terra.ReferenceResource(dgt)}
}

func (dgt *DynamodbGlobalTable) ImportState(av io.Reader) error {
	dgt.state = &dynamodbGlobalTableState{}
	if err := json.NewDecoder(av).Decode(dgt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dgt.Type(), dgt.LocalName(), err)
	}
	return nil
}

func (dgt *DynamodbGlobalTable) State() (*dynamodbGlobalTableState, bool) {
	return dgt.state, dgt.state != nil
}

func (dgt *DynamodbGlobalTable) StateMust() *dynamodbGlobalTableState {
	if dgt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dgt.Type(), dgt.LocalName()))
	}
	return dgt.state
}

func (dgt *DynamodbGlobalTable) DependOn() terra.Reference {
	return terra.ReferenceResource(dgt)
}

type DynamodbGlobalTableArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Replica: min=1
	Replica []dynamodbglobaltable.Replica `hcl:"replica,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *dynamodbglobaltable.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that DynamodbGlobalTable depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dynamodbGlobalTableAttributes struct {
	ref terra.Reference
}

func (dgt dynamodbGlobalTableAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(dgt.ref.Append("arn"))
}

func (dgt dynamodbGlobalTableAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dgt.ref.Append("id"))
}

func (dgt dynamodbGlobalTableAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dgt.ref.Append("name"))
}

func (dgt dynamodbGlobalTableAttributes) Replica() terra.SetValue[dynamodbglobaltable.ReplicaAttributes] {
	return terra.ReferenceSet[dynamodbglobaltable.ReplicaAttributes](dgt.ref.Append("replica"))
}

func (dgt dynamodbGlobalTableAttributes) Timeouts() dynamodbglobaltable.TimeoutsAttributes {
	return terra.ReferenceSingle[dynamodbglobaltable.TimeoutsAttributes](dgt.ref.Append("timeouts"))
}

type dynamodbGlobalTableState struct {
	Arn      string                             `json:"arn"`
	Id       string                             `json:"id"`
	Name     string                             `json:"name"`
	Replica  []dynamodbglobaltable.ReplicaState `json:"replica"`
	Timeouts *dynamodbglobaltable.TimeoutsState `json:"timeouts"`
}
