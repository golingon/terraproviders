// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2TrafficMirrorFilter creates a new instance of [Ec2TrafficMirrorFilter].
func NewEc2TrafficMirrorFilter(name string, args Ec2TrafficMirrorFilterArgs) *Ec2TrafficMirrorFilter {
	return &Ec2TrafficMirrorFilter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TrafficMirrorFilter)(nil)

// Ec2TrafficMirrorFilter represents the Terraform resource aws_ec2_traffic_mirror_filter.
type Ec2TrafficMirrorFilter struct {
	Name      string
	Args      Ec2TrafficMirrorFilterArgs
	state     *ec2TrafficMirrorFilterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TrafficMirrorFilter].
func (etmf *Ec2TrafficMirrorFilter) Type() string {
	return "aws_ec2_traffic_mirror_filter"
}

// LocalName returns the local name for [Ec2TrafficMirrorFilter].
func (etmf *Ec2TrafficMirrorFilter) LocalName() string {
	return etmf.Name
}

// Configuration returns the configuration (args) for [Ec2TrafficMirrorFilter].
func (etmf *Ec2TrafficMirrorFilter) Configuration() interface{} {
	return etmf.Args
}

// DependOn is used for other resources to depend on [Ec2TrafficMirrorFilter].
func (etmf *Ec2TrafficMirrorFilter) DependOn() terra.Reference {
	return terra.ReferenceResource(etmf)
}

// Dependencies returns the list of resources [Ec2TrafficMirrorFilter] depends_on.
func (etmf *Ec2TrafficMirrorFilter) Dependencies() terra.Dependencies {
	return etmf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TrafficMirrorFilter].
func (etmf *Ec2TrafficMirrorFilter) LifecycleManagement() *terra.Lifecycle {
	return etmf.Lifecycle
}

// Attributes returns the attributes for [Ec2TrafficMirrorFilter].
func (etmf *Ec2TrafficMirrorFilter) Attributes() ec2TrafficMirrorFilterAttributes {
	return ec2TrafficMirrorFilterAttributes{ref: terra.ReferenceResource(etmf)}
}

// ImportState imports the given attribute values into [Ec2TrafficMirrorFilter]'s state.
func (etmf *Ec2TrafficMirrorFilter) ImportState(av io.Reader) error {
	etmf.state = &ec2TrafficMirrorFilterState{}
	if err := json.NewDecoder(av).Decode(etmf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etmf.Type(), etmf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TrafficMirrorFilter] has state.
func (etmf *Ec2TrafficMirrorFilter) State() (*ec2TrafficMirrorFilterState, bool) {
	return etmf.state, etmf.state != nil
}

// StateMust returns the state for [Ec2TrafficMirrorFilter]. Panics if the state is nil.
func (etmf *Ec2TrafficMirrorFilter) StateMust() *ec2TrafficMirrorFilterState {
	if etmf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etmf.Type(), etmf.LocalName()))
	}
	return etmf.state
}

// Ec2TrafficMirrorFilterArgs contains the configurations for aws_ec2_traffic_mirror_filter.
type Ec2TrafficMirrorFilterArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkServices: set of string, optional
	NetworkServices terra.SetValue[terra.StringValue] `hcl:"network_services,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type ec2TrafficMirrorFilterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_traffic_mirror_filter.
func (etmf ec2TrafficMirrorFilterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(etmf.ref.Append("arn"))
}

// Description returns a reference to field description of aws_ec2_traffic_mirror_filter.
func (etmf ec2TrafficMirrorFilterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(etmf.ref.Append("description"))
}

// Id returns a reference to field id of aws_ec2_traffic_mirror_filter.
func (etmf ec2TrafficMirrorFilterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etmf.ref.Append("id"))
}

// NetworkServices returns a reference to field network_services of aws_ec2_traffic_mirror_filter.
func (etmf ec2TrafficMirrorFilterAttributes) NetworkServices() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](etmf.ref.Append("network_services"))
}

// Tags returns a reference to field tags of aws_ec2_traffic_mirror_filter.
func (etmf ec2TrafficMirrorFilterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etmf.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_traffic_mirror_filter.
func (etmf ec2TrafficMirrorFilterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etmf.ref.Append("tags_all"))
}

type ec2TrafficMirrorFilterState struct {
	Arn             string            `json:"arn"`
	Description     string            `json:"description"`
	Id              string            `json:"id"`
	NetworkServices []string          `json:"network_services"`
	Tags            map[string]string `json:"tags"`
	TagsAll         map[string]string `json:"tags_all"`
}