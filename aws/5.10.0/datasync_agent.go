// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	datasyncagent "github.com/golingon/terraproviders/aws/5.10.0/datasyncagent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatasyncAgent creates a new instance of [DatasyncAgent].
func NewDatasyncAgent(name string, args DatasyncAgentArgs) *DatasyncAgent {
	return &DatasyncAgent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatasyncAgent)(nil)

// DatasyncAgent represents the Terraform resource aws_datasync_agent.
type DatasyncAgent struct {
	Name      string
	Args      DatasyncAgentArgs
	state     *datasyncAgentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatasyncAgent].
func (da *DatasyncAgent) Type() string {
	return "aws_datasync_agent"
}

// LocalName returns the local name for [DatasyncAgent].
func (da *DatasyncAgent) LocalName() string {
	return da.Name
}

// Configuration returns the configuration (args) for [DatasyncAgent].
func (da *DatasyncAgent) Configuration() interface{} {
	return da.Args
}

// DependOn is used for other resources to depend on [DatasyncAgent].
func (da *DatasyncAgent) DependOn() terra.Reference {
	return terra.ReferenceResource(da)
}

// Dependencies returns the list of resources [DatasyncAgent] depends_on.
func (da *DatasyncAgent) Dependencies() terra.Dependencies {
	return da.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatasyncAgent].
func (da *DatasyncAgent) LifecycleManagement() *terra.Lifecycle {
	return da.Lifecycle
}

// Attributes returns the attributes for [DatasyncAgent].
func (da *DatasyncAgent) Attributes() datasyncAgentAttributes {
	return datasyncAgentAttributes{ref: terra.ReferenceResource(da)}
}

// ImportState imports the given attribute values into [DatasyncAgent]'s state.
func (da *DatasyncAgent) ImportState(av io.Reader) error {
	da.state = &datasyncAgentState{}
	if err := json.NewDecoder(av).Decode(da.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", da.Type(), da.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatasyncAgent] has state.
func (da *DatasyncAgent) State() (*datasyncAgentState, bool) {
	return da.state, da.state != nil
}

// StateMust returns the state for [DatasyncAgent]. Panics if the state is nil.
func (da *DatasyncAgent) StateMust() *datasyncAgentState {
	if da.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", da.Type(), da.LocalName()))
	}
	return da.state
}

// DatasyncAgentArgs contains the configurations for aws_datasync_agent.
type DatasyncAgentArgs struct {
	// ActivationKey: string, optional
	ActivationKey terra.StringValue `hcl:"activation_key,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// PrivateLinkEndpoint: string, optional
	PrivateLinkEndpoint terra.StringValue `hcl:"private_link_endpoint,attr"`
	// SecurityGroupArns: set of string, optional
	SecurityGroupArns terra.SetValue[terra.StringValue] `hcl:"security_group_arns,attr"`
	// SubnetArns: set of string, optional
	SubnetArns terra.SetValue[terra.StringValue] `hcl:"subnet_arns,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcEndpointId: string, optional
	VpcEndpointId terra.StringValue `hcl:"vpc_endpoint_id,attr"`
	// Timeouts: optional
	Timeouts *datasyncagent.Timeouts `hcl:"timeouts,block"`
}
type datasyncAgentAttributes struct {
	ref terra.Reference
}

// ActivationKey returns a reference to field activation_key of aws_datasync_agent.
func (da datasyncAgentAttributes) ActivationKey() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("activation_key"))
}

// Arn returns a reference to field arn of aws_datasync_agent.
func (da datasyncAgentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("arn"))
}

// Id returns a reference to field id of aws_datasync_agent.
func (da datasyncAgentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of aws_datasync_agent.
func (da datasyncAgentAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("ip_address"))
}

// Name returns a reference to field name of aws_datasync_agent.
func (da datasyncAgentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("name"))
}

// PrivateLinkEndpoint returns a reference to field private_link_endpoint of aws_datasync_agent.
func (da datasyncAgentAttributes) PrivateLinkEndpoint() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("private_link_endpoint"))
}

// SecurityGroupArns returns a reference to field security_group_arns of aws_datasync_agent.
func (da datasyncAgentAttributes) SecurityGroupArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](da.ref.Append("security_group_arns"))
}

// SubnetArns returns a reference to field subnet_arns of aws_datasync_agent.
func (da datasyncAgentAttributes) SubnetArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](da.ref.Append("subnet_arns"))
}

// Tags returns a reference to field tags of aws_datasync_agent.
func (da datasyncAgentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](da.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_datasync_agent.
func (da datasyncAgentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](da.ref.Append("tags_all"))
}

// VpcEndpointId returns a reference to field vpc_endpoint_id of aws_datasync_agent.
func (da datasyncAgentAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("vpc_endpoint_id"))
}

func (da datasyncAgentAttributes) Timeouts() datasyncagent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasyncagent.TimeoutsAttributes](da.ref.Append("timeouts"))
}

type datasyncAgentState struct {
	ActivationKey       string                       `json:"activation_key"`
	Arn                 string                       `json:"arn"`
	Id                  string                       `json:"id"`
	IpAddress           string                       `json:"ip_address"`
	Name                string                       `json:"name"`
	PrivateLinkEndpoint string                       `json:"private_link_endpoint"`
	SecurityGroupArns   []string                     `json:"security_group_arns"`
	SubnetArns          []string                     `json:"subnet_arns"`
	Tags                map[string]string            `json:"tags"`
	TagsAll             map[string]string            `json:"tags_all"`
	VpcEndpointId       string                       `json:"vpc_endpoint_id"`
	Timeouts            *datasyncagent.TimeoutsState `json:"timeouts"`
}
