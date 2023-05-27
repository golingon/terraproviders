// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudwatcheventconnection "github.com/golingon/terraproviders/aws/5.0.1/cloudwatcheventconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudwatchEventConnection creates a new instance of [CloudwatchEventConnection].
func NewCloudwatchEventConnection(name string, args CloudwatchEventConnectionArgs) *CloudwatchEventConnection {
	return &CloudwatchEventConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudwatchEventConnection)(nil)

// CloudwatchEventConnection represents the Terraform resource aws_cloudwatch_event_connection.
type CloudwatchEventConnection struct {
	Name      string
	Args      CloudwatchEventConnectionArgs
	state     *cloudwatchEventConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudwatchEventConnection].
func (cec *CloudwatchEventConnection) Type() string {
	return "aws_cloudwatch_event_connection"
}

// LocalName returns the local name for [CloudwatchEventConnection].
func (cec *CloudwatchEventConnection) LocalName() string {
	return cec.Name
}

// Configuration returns the configuration (args) for [CloudwatchEventConnection].
func (cec *CloudwatchEventConnection) Configuration() interface{} {
	return cec.Args
}

// DependOn is used for other resources to depend on [CloudwatchEventConnection].
func (cec *CloudwatchEventConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(cec)
}

// Dependencies returns the list of resources [CloudwatchEventConnection] depends_on.
func (cec *CloudwatchEventConnection) Dependencies() terra.Dependencies {
	return cec.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudwatchEventConnection].
func (cec *CloudwatchEventConnection) LifecycleManagement() *terra.Lifecycle {
	return cec.Lifecycle
}

// Attributes returns the attributes for [CloudwatchEventConnection].
func (cec *CloudwatchEventConnection) Attributes() cloudwatchEventConnectionAttributes {
	return cloudwatchEventConnectionAttributes{ref: terra.ReferenceResource(cec)}
}

// ImportState imports the given attribute values into [CloudwatchEventConnection]'s state.
func (cec *CloudwatchEventConnection) ImportState(av io.Reader) error {
	cec.state = &cloudwatchEventConnectionState{}
	if err := json.NewDecoder(av).Decode(cec.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cec.Type(), cec.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudwatchEventConnection] has state.
func (cec *CloudwatchEventConnection) State() (*cloudwatchEventConnectionState, bool) {
	return cec.state, cec.state != nil
}

// StateMust returns the state for [CloudwatchEventConnection]. Panics if the state is nil.
func (cec *CloudwatchEventConnection) StateMust() *cloudwatchEventConnectionState {
	if cec.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cec.Type(), cec.LocalName()))
	}
	return cec.state
}

// CloudwatchEventConnectionArgs contains the configurations for aws_cloudwatch_event_connection.
type CloudwatchEventConnectionArgs struct {
	// AuthorizationType: string, required
	AuthorizationType terra.StringValue `hcl:"authorization_type,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// AuthParameters: required
	AuthParameters *cloudwatcheventconnection.AuthParameters `hcl:"auth_parameters,block" validate:"required"`
}
type cloudwatchEventConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudwatch_event_connection.
func (cec cloudwatchEventConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("arn"))
}

// AuthorizationType returns a reference to field authorization_type of aws_cloudwatch_event_connection.
func (cec cloudwatchEventConnectionAttributes) AuthorizationType() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("authorization_type"))
}

// Description returns a reference to field description of aws_cloudwatch_event_connection.
func (cec cloudwatchEventConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("description"))
}

// Id returns a reference to field id of aws_cloudwatch_event_connection.
func (cec cloudwatchEventConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudwatch_event_connection.
func (cec cloudwatchEventConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("name"))
}

// SecretArn returns a reference to field secret_arn of aws_cloudwatch_event_connection.
func (cec cloudwatchEventConnectionAttributes) SecretArn() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("secret_arn"))
}

func (cec cloudwatchEventConnectionAttributes) AuthParameters() terra.ListValue[cloudwatcheventconnection.AuthParametersAttributes] {
	return terra.ReferenceAsList[cloudwatcheventconnection.AuthParametersAttributes](cec.ref.Append("auth_parameters"))
}

type cloudwatchEventConnectionState struct {
	Arn               string                                          `json:"arn"`
	AuthorizationType string                                          `json:"authorization_type"`
	Description       string                                          `json:"description"`
	Id                string                                          `json:"id"`
	Name              string                                          `json:"name"`
	SecretArn         string                                          `json:"secret_arn"`
	AuthParameters    []cloudwatcheventconnection.AuthParametersState `json:"auth_parameters"`
}
