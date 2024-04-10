// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	quicksightvpcconnection "github.com/golingon/terraproviders/aws/5.44.0/quicksightvpcconnection"
	"io"
)

// NewQuicksightVpcConnection creates a new instance of [QuicksightVpcConnection].
func NewQuicksightVpcConnection(name string, args QuicksightVpcConnectionArgs) *QuicksightVpcConnection {
	return &QuicksightVpcConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*QuicksightVpcConnection)(nil)

// QuicksightVpcConnection represents the Terraform resource aws_quicksight_vpc_connection.
type QuicksightVpcConnection struct {
	Name      string
	Args      QuicksightVpcConnectionArgs
	state     *quicksightVpcConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [QuicksightVpcConnection].
func (qvc *QuicksightVpcConnection) Type() string {
	return "aws_quicksight_vpc_connection"
}

// LocalName returns the local name for [QuicksightVpcConnection].
func (qvc *QuicksightVpcConnection) LocalName() string {
	return qvc.Name
}

// Configuration returns the configuration (args) for [QuicksightVpcConnection].
func (qvc *QuicksightVpcConnection) Configuration() interface{} {
	return qvc.Args
}

// DependOn is used for other resources to depend on [QuicksightVpcConnection].
func (qvc *QuicksightVpcConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(qvc)
}

// Dependencies returns the list of resources [QuicksightVpcConnection] depends_on.
func (qvc *QuicksightVpcConnection) Dependencies() terra.Dependencies {
	return qvc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [QuicksightVpcConnection].
func (qvc *QuicksightVpcConnection) LifecycleManagement() *terra.Lifecycle {
	return qvc.Lifecycle
}

// Attributes returns the attributes for [QuicksightVpcConnection].
func (qvc *QuicksightVpcConnection) Attributes() quicksightVpcConnectionAttributes {
	return quicksightVpcConnectionAttributes{ref: terra.ReferenceResource(qvc)}
}

// ImportState imports the given attribute values into [QuicksightVpcConnection]'s state.
func (qvc *QuicksightVpcConnection) ImportState(av io.Reader) error {
	qvc.state = &quicksightVpcConnectionState{}
	if err := json.NewDecoder(av).Decode(qvc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", qvc.Type(), qvc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [QuicksightVpcConnection] has state.
func (qvc *QuicksightVpcConnection) State() (*quicksightVpcConnectionState, bool) {
	return qvc.state, qvc.state != nil
}

// StateMust returns the state for [QuicksightVpcConnection]. Panics if the state is nil.
func (qvc *QuicksightVpcConnection) StateMust() *quicksightVpcConnectionState {
	if qvc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", qvc.Type(), qvc.LocalName()))
	}
	return qvc.state
}

// QuicksightVpcConnectionArgs contains the configurations for aws_quicksight_vpc_connection.
type QuicksightVpcConnectionArgs struct {
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// DnsResolvers: set of string, optional
	DnsResolvers terra.SetValue[terra.StringValue] `hcl:"dns_resolvers,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// SecurityGroupIds: set of string, required
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr" validate:"required"`
	// SubnetIds: set of string, required
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VpcConnectionId: string, required
	VpcConnectionId terra.StringValue `hcl:"vpc_connection_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *quicksightvpcconnection.Timeouts `hcl:"timeouts,block"`
}
type quicksightVpcConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(qvc.ref.Append("arn"))
}

// AvailabilityStatus returns a reference to field availability_status of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) AvailabilityStatus() terra.StringValue {
	return terra.ReferenceAsString(qvc.ref.Append("availability_status"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(qvc.ref.Append("aws_account_id"))
}

// DnsResolvers returns a reference to field dns_resolvers of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) DnsResolvers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](qvc.ref.Append("dns_resolvers"))
}

// Id returns a reference to field id of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(qvc.ref.Append("id"))
}

// Name returns a reference to field name of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(qvc.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(qvc.ref.Append("role_arn"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](qvc.ref.Append("security_group_ids"))
}

// SubnetIds returns a reference to field subnet_ids of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](qvc.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qvc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qvc.ref.Append("tags_all"))
}

// VpcConnectionId returns a reference to field vpc_connection_id of aws_quicksight_vpc_connection.
func (qvc quicksightVpcConnectionAttributes) VpcConnectionId() terra.StringValue {
	return terra.ReferenceAsString(qvc.ref.Append("vpc_connection_id"))
}

func (qvc quicksightVpcConnectionAttributes) Timeouts() quicksightvpcconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[quicksightvpcconnection.TimeoutsAttributes](qvc.ref.Append("timeouts"))
}

type quicksightVpcConnectionState struct {
	Arn                string                                 `json:"arn"`
	AvailabilityStatus string                                 `json:"availability_status"`
	AwsAccountId       string                                 `json:"aws_account_id"`
	DnsResolvers       []string                               `json:"dns_resolvers"`
	Id                 string                                 `json:"id"`
	Name               string                                 `json:"name"`
	RoleArn            string                                 `json:"role_arn"`
	SecurityGroupIds   []string                               `json:"security_group_ids"`
	SubnetIds          []string                               `json:"subnet_ids"`
	Tags               map[string]string                      `json:"tags"`
	TagsAll            map[string]string                      `json:"tags_all"`
	VpcConnectionId    string                                 `json:"vpc_connection_id"`
	Timeouts           *quicksightvpcconnection.TimeoutsState `json:"timeouts"`
}
