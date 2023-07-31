// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	internetgateway "github.com/golingon/terraproviders/aws/5.10.0/internetgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewInternetGateway creates a new instance of [InternetGateway].
func NewInternetGateway(name string, args InternetGatewayArgs) *InternetGateway {
	return &InternetGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*InternetGateway)(nil)

// InternetGateway represents the Terraform resource aws_internet_gateway.
type InternetGateway struct {
	Name      string
	Args      InternetGatewayArgs
	state     *internetGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [InternetGateway].
func (ig *InternetGateway) Type() string {
	return "aws_internet_gateway"
}

// LocalName returns the local name for [InternetGateway].
func (ig *InternetGateway) LocalName() string {
	return ig.Name
}

// Configuration returns the configuration (args) for [InternetGateway].
func (ig *InternetGateway) Configuration() interface{} {
	return ig.Args
}

// DependOn is used for other resources to depend on [InternetGateway].
func (ig *InternetGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(ig)
}

// Dependencies returns the list of resources [InternetGateway] depends_on.
func (ig *InternetGateway) Dependencies() terra.Dependencies {
	return ig.DependsOn
}

// LifecycleManagement returns the lifecycle block for [InternetGateway].
func (ig *InternetGateway) LifecycleManagement() *terra.Lifecycle {
	return ig.Lifecycle
}

// Attributes returns the attributes for [InternetGateway].
func (ig *InternetGateway) Attributes() internetGatewayAttributes {
	return internetGatewayAttributes{ref: terra.ReferenceResource(ig)}
}

// ImportState imports the given attribute values into [InternetGateway]'s state.
func (ig *InternetGateway) ImportState(av io.Reader) error {
	ig.state = &internetGatewayState{}
	if err := json.NewDecoder(av).Decode(ig.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ig.Type(), ig.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [InternetGateway] has state.
func (ig *InternetGateway) State() (*internetGatewayState, bool) {
	return ig.state, ig.state != nil
}

// StateMust returns the state for [InternetGateway]. Panics if the state is nil.
func (ig *InternetGateway) StateMust() *internetGatewayState {
	if ig.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ig.Type(), ig.LocalName()))
	}
	return ig.state
}

// InternetGatewayArgs contains the configurations for aws_internet_gateway.
type InternetGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// Timeouts: optional
	Timeouts *internetgateway.Timeouts `hcl:"timeouts,block"`
}
type internetGatewayAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_internet_gateway.
func (ig internetGatewayAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("arn"))
}

// Id returns a reference to field id of aws_internet_gateway.
func (ig internetGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("id"))
}

// OwnerId returns a reference to field owner_id of aws_internet_gateway.
func (ig internetGatewayAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_internet_gateway.
func (ig internetGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ig.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_internet_gateway.
func (ig internetGatewayAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ig.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_internet_gateway.
func (ig internetGatewayAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("vpc_id"))
}

func (ig internetGatewayAttributes) Timeouts() internetgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[internetgateway.TimeoutsAttributes](ig.ref.Append("timeouts"))
}

type internetGatewayState struct {
	Arn      string                         `json:"arn"`
	Id       string                         `json:"id"`
	OwnerId  string                         `json:"owner_id"`
	Tags     map[string]string              `json:"tags"`
	TagsAll  map[string]string              `json:"tags_all"`
	VpcId    string                         `json:"vpc_id"`
	Timeouts *internetgateway.TimeoutsState `json:"timeouts"`
}
