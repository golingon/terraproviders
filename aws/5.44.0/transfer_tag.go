// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewTransferTag creates a new instance of [TransferTag].
func NewTransferTag(name string, args TransferTagArgs) *TransferTag {
	return &TransferTag{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TransferTag)(nil)

// TransferTag represents the Terraform resource aws_transfer_tag.
type TransferTag struct {
	Name      string
	Args      TransferTagArgs
	state     *transferTagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TransferTag].
func (tt *TransferTag) Type() string {
	return "aws_transfer_tag"
}

// LocalName returns the local name for [TransferTag].
func (tt *TransferTag) LocalName() string {
	return tt.Name
}

// Configuration returns the configuration (args) for [TransferTag].
func (tt *TransferTag) Configuration() interface{} {
	return tt.Args
}

// DependOn is used for other resources to depend on [TransferTag].
func (tt *TransferTag) DependOn() terra.Reference {
	return terra.ReferenceResource(tt)
}

// Dependencies returns the list of resources [TransferTag] depends_on.
func (tt *TransferTag) Dependencies() terra.Dependencies {
	return tt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TransferTag].
func (tt *TransferTag) LifecycleManagement() *terra.Lifecycle {
	return tt.Lifecycle
}

// Attributes returns the attributes for [TransferTag].
func (tt *TransferTag) Attributes() transferTagAttributes {
	return transferTagAttributes{ref: terra.ReferenceResource(tt)}
}

// ImportState imports the given attribute values into [TransferTag]'s state.
func (tt *TransferTag) ImportState(av io.Reader) error {
	tt.state = &transferTagState{}
	if err := json.NewDecoder(av).Decode(tt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tt.Type(), tt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TransferTag] has state.
func (tt *TransferTag) State() (*transferTagState, bool) {
	return tt.state, tt.state != nil
}

// StateMust returns the state for [TransferTag]. Panics if the state is nil.
func (tt *TransferTag) StateMust() *transferTagState {
	if tt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tt.Type(), tt.LocalName()))
	}
	return tt.state
}

// TransferTagArgs contains the configurations for aws_transfer_tag.
type TransferTagArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}
type transferTagAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_transfer_tag.
func (tt transferTagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("id"))
}

// Key returns a reference to field key of aws_transfer_tag.
func (tt transferTagAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("key"))
}

// ResourceArn returns a reference to field resource_arn of aws_transfer_tag.
func (tt transferTagAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("resource_arn"))
}

// Value returns a reference to field value of aws_transfer_tag.
func (tt transferTagAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("value"))
}

type transferTagState struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	ResourceArn string `json:"resource_arn"`
	Value       string `json:"value"`
}
