// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	internetgatewayattachment "github.com/golingon/terraproviders/aws/5.6.2/internetgatewayattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewInternetGatewayAttachment creates a new instance of [InternetGatewayAttachment].
func NewInternetGatewayAttachment(name string, args InternetGatewayAttachmentArgs) *InternetGatewayAttachment {
	return &InternetGatewayAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*InternetGatewayAttachment)(nil)

// InternetGatewayAttachment represents the Terraform resource aws_internet_gateway_attachment.
type InternetGatewayAttachment struct {
	Name      string
	Args      InternetGatewayAttachmentArgs
	state     *internetGatewayAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [InternetGatewayAttachment].
func (iga *InternetGatewayAttachment) Type() string {
	return "aws_internet_gateway_attachment"
}

// LocalName returns the local name for [InternetGatewayAttachment].
func (iga *InternetGatewayAttachment) LocalName() string {
	return iga.Name
}

// Configuration returns the configuration (args) for [InternetGatewayAttachment].
func (iga *InternetGatewayAttachment) Configuration() interface{} {
	return iga.Args
}

// DependOn is used for other resources to depend on [InternetGatewayAttachment].
func (iga *InternetGatewayAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(iga)
}

// Dependencies returns the list of resources [InternetGatewayAttachment] depends_on.
func (iga *InternetGatewayAttachment) Dependencies() terra.Dependencies {
	return iga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [InternetGatewayAttachment].
func (iga *InternetGatewayAttachment) LifecycleManagement() *terra.Lifecycle {
	return iga.Lifecycle
}

// Attributes returns the attributes for [InternetGatewayAttachment].
func (iga *InternetGatewayAttachment) Attributes() internetGatewayAttachmentAttributes {
	return internetGatewayAttachmentAttributes{ref: terra.ReferenceResource(iga)}
}

// ImportState imports the given attribute values into [InternetGatewayAttachment]'s state.
func (iga *InternetGatewayAttachment) ImportState(av io.Reader) error {
	iga.state = &internetGatewayAttachmentState{}
	if err := json.NewDecoder(av).Decode(iga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iga.Type(), iga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [InternetGatewayAttachment] has state.
func (iga *InternetGatewayAttachment) State() (*internetGatewayAttachmentState, bool) {
	return iga.state, iga.state != nil
}

// StateMust returns the state for [InternetGatewayAttachment]. Panics if the state is nil.
func (iga *InternetGatewayAttachment) StateMust() *internetGatewayAttachmentState {
	if iga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iga.Type(), iga.LocalName()))
	}
	return iga.state
}

// InternetGatewayAttachmentArgs contains the configurations for aws_internet_gateway_attachment.
type InternetGatewayAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InternetGatewayId: string, required
	InternetGatewayId terra.StringValue `hcl:"internet_gateway_id,attr" validate:"required"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *internetgatewayattachment.Timeouts `hcl:"timeouts,block"`
}
type internetGatewayAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_internet_gateway_attachment.
func (iga internetGatewayAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iga.ref.Append("id"))
}

// InternetGatewayId returns a reference to field internet_gateway_id of aws_internet_gateway_attachment.
func (iga internetGatewayAttachmentAttributes) InternetGatewayId() terra.StringValue {
	return terra.ReferenceAsString(iga.ref.Append("internet_gateway_id"))
}

// VpcId returns a reference to field vpc_id of aws_internet_gateway_attachment.
func (iga internetGatewayAttachmentAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(iga.ref.Append("vpc_id"))
}

func (iga internetGatewayAttachmentAttributes) Timeouts() internetgatewayattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[internetgatewayattachment.TimeoutsAttributes](iga.ref.Append("timeouts"))
}

type internetGatewayAttachmentState struct {
	Id                string                                   `json:"id"`
	InternetGatewayId string                                   `json:"internet_gateway_id"`
	VpcId             string                                   `json:"vpc_id"`
	Timeouts          *internetgatewayattachment.TimeoutsState `json:"timeouts"`
}
