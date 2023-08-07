// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2NetworkInsightsPath creates a new instance of [Ec2NetworkInsightsPath].
func NewEc2NetworkInsightsPath(name string, args Ec2NetworkInsightsPathArgs) *Ec2NetworkInsightsPath {
	return &Ec2NetworkInsightsPath{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2NetworkInsightsPath)(nil)

// Ec2NetworkInsightsPath represents the Terraform resource aws_ec2_network_insights_path.
type Ec2NetworkInsightsPath struct {
	Name      string
	Args      Ec2NetworkInsightsPathArgs
	state     *ec2NetworkInsightsPathState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2NetworkInsightsPath].
func (enip *Ec2NetworkInsightsPath) Type() string {
	return "aws_ec2_network_insights_path"
}

// LocalName returns the local name for [Ec2NetworkInsightsPath].
func (enip *Ec2NetworkInsightsPath) LocalName() string {
	return enip.Name
}

// Configuration returns the configuration (args) for [Ec2NetworkInsightsPath].
func (enip *Ec2NetworkInsightsPath) Configuration() interface{} {
	return enip.Args
}

// DependOn is used for other resources to depend on [Ec2NetworkInsightsPath].
func (enip *Ec2NetworkInsightsPath) DependOn() terra.Reference {
	return terra.ReferenceResource(enip)
}

// Dependencies returns the list of resources [Ec2NetworkInsightsPath] depends_on.
func (enip *Ec2NetworkInsightsPath) Dependencies() terra.Dependencies {
	return enip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2NetworkInsightsPath].
func (enip *Ec2NetworkInsightsPath) LifecycleManagement() *terra.Lifecycle {
	return enip.Lifecycle
}

// Attributes returns the attributes for [Ec2NetworkInsightsPath].
func (enip *Ec2NetworkInsightsPath) Attributes() ec2NetworkInsightsPathAttributes {
	return ec2NetworkInsightsPathAttributes{ref: terra.ReferenceResource(enip)}
}

// ImportState imports the given attribute values into [Ec2NetworkInsightsPath]'s state.
func (enip *Ec2NetworkInsightsPath) ImportState(av io.Reader) error {
	enip.state = &ec2NetworkInsightsPathState{}
	if err := json.NewDecoder(av).Decode(enip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", enip.Type(), enip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2NetworkInsightsPath] has state.
func (enip *Ec2NetworkInsightsPath) State() (*ec2NetworkInsightsPathState, bool) {
	return enip.state, enip.state != nil
}

// StateMust returns the state for [Ec2NetworkInsightsPath]. Panics if the state is nil.
func (enip *Ec2NetworkInsightsPath) StateMust() *ec2NetworkInsightsPathState {
	if enip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", enip.Type(), enip.LocalName()))
	}
	return enip.state
}

// Ec2NetworkInsightsPathArgs contains the configurations for aws_ec2_network_insights_path.
type Ec2NetworkInsightsPathArgs struct {
	// Destination: string, required
	Destination terra.StringValue `hcl:"destination,attr" validate:"required"`
	// DestinationIp: string, optional
	DestinationIp terra.StringValue `hcl:"destination_ip,attr"`
	// DestinationPort: number, optional
	DestinationPort terra.NumberValue `hcl:"destination_port,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// SourceIp: string, optional
	SourceIp terra.StringValue `hcl:"source_ip,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type ec2NetworkInsightsPathAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_network_insights_path.
func (enip ec2NetworkInsightsPathAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("arn"))
}

// Destination returns a reference to field destination of aws_ec2_network_insights_path.
func (enip ec2NetworkInsightsPathAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("destination"))
}

// DestinationIp returns a reference to field destination_ip of aws_ec2_network_insights_path.
func (enip ec2NetworkInsightsPathAttributes) DestinationIp() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("destination_ip"))
}

// DestinationPort returns a reference to field destination_port of aws_ec2_network_insights_path.
func (enip ec2NetworkInsightsPathAttributes) DestinationPort() terra.NumberValue {
	return terra.ReferenceAsNumber(enip.ref.Append("destination_port"))
}

// Id returns a reference to field id of aws_ec2_network_insights_path.
func (enip ec2NetworkInsightsPathAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("id"))
}

// Protocol returns a reference to field protocol of aws_ec2_network_insights_path.
func (enip ec2NetworkInsightsPathAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("protocol"))
}

// Source returns a reference to field source of aws_ec2_network_insights_path.
func (enip ec2NetworkInsightsPathAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("source"))
}

// SourceIp returns a reference to field source_ip of aws_ec2_network_insights_path.
func (enip ec2NetworkInsightsPathAttributes) SourceIp() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("source_ip"))
}

// Tags returns a reference to field tags of aws_ec2_network_insights_path.
func (enip ec2NetworkInsightsPathAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](enip.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_network_insights_path.
func (enip ec2NetworkInsightsPathAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](enip.ref.Append("tags_all"))
}

type ec2NetworkInsightsPathState struct {
	Arn             string            `json:"arn"`
	Destination     string            `json:"destination"`
	DestinationIp   string            `json:"destination_ip"`
	DestinationPort float64           `json:"destination_port"`
	Id              string            `json:"id"`
	Protocol        string            `json:"protocol"`
	Source          string            `json:"source"`
	SourceIp        string            `json:"source_ip"`
	Tags            map[string]string `json:"tags"`
	TagsAll         map[string]string `json:"tags_all"`
}
