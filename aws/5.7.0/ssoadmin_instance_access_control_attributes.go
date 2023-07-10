// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssoadmininstanceaccesscontrolattributes "github.com/golingon/terraproviders/aws/5.7.0/ssoadmininstanceaccesscontrolattributes"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsoadminInstanceAccessControlAttributes creates a new instance of [SsoadminInstanceAccessControlAttributes].
func NewSsoadminInstanceAccessControlAttributes(name string, args SsoadminInstanceAccessControlAttributesArgs) *SsoadminInstanceAccessControlAttributes {
	return &SsoadminInstanceAccessControlAttributes{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsoadminInstanceAccessControlAttributes)(nil)

// SsoadminInstanceAccessControlAttributes represents the Terraform resource aws_ssoadmin_instance_access_control_attributes.
type SsoadminInstanceAccessControlAttributes struct {
	Name      string
	Args      SsoadminInstanceAccessControlAttributesArgs
	state     *ssoadminInstanceAccessControlAttributesState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsoadminInstanceAccessControlAttributes].
func (siaca *SsoadminInstanceAccessControlAttributes) Type() string {
	return "aws_ssoadmin_instance_access_control_attributes"
}

// LocalName returns the local name for [SsoadminInstanceAccessControlAttributes].
func (siaca *SsoadminInstanceAccessControlAttributes) LocalName() string {
	return siaca.Name
}

// Configuration returns the configuration (args) for [SsoadminInstanceAccessControlAttributes].
func (siaca *SsoadminInstanceAccessControlAttributes) Configuration() interface{} {
	return siaca.Args
}

// DependOn is used for other resources to depend on [SsoadminInstanceAccessControlAttributes].
func (siaca *SsoadminInstanceAccessControlAttributes) DependOn() terra.Reference {
	return terra.ReferenceResource(siaca)
}

// Dependencies returns the list of resources [SsoadminInstanceAccessControlAttributes] depends_on.
func (siaca *SsoadminInstanceAccessControlAttributes) Dependencies() terra.Dependencies {
	return siaca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsoadminInstanceAccessControlAttributes].
func (siaca *SsoadminInstanceAccessControlAttributes) LifecycleManagement() *terra.Lifecycle {
	return siaca.Lifecycle
}

// Attributes returns the attributes for [SsoadminInstanceAccessControlAttributes].
func (siaca *SsoadminInstanceAccessControlAttributes) Attributes() ssoadminInstanceAccessControlAttributesAttributes {
	return ssoadminInstanceAccessControlAttributesAttributes{ref: terra.ReferenceResource(siaca)}
}

// ImportState imports the given attribute values into [SsoadminInstanceAccessControlAttributes]'s state.
func (siaca *SsoadminInstanceAccessControlAttributes) ImportState(av io.Reader) error {
	siaca.state = &ssoadminInstanceAccessControlAttributesState{}
	if err := json.NewDecoder(av).Decode(siaca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", siaca.Type(), siaca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsoadminInstanceAccessControlAttributes] has state.
func (siaca *SsoadminInstanceAccessControlAttributes) State() (*ssoadminInstanceAccessControlAttributesState, bool) {
	return siaca.state, siaca.state != nil
}

// StateMust returns the state for [SsoadminInstanceAccessControlAttributes]. Panics if the state is nil.
func (siaca *SsoadminInstanceAccessControlAttributes) StateMust() *ssoadminInstanceAccessControlAttributesState {
	if siaca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", siaca.Type(), siaca.LocalName()))
	}
	return siaca.state
}

// SsoadminInstanceAccessControlAttributesArgs contains the configurations for aws_ssoadmin_instance_access_control_attributes.
type SsoadminInstanceAccessControlAttributesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceArn: string, required
	InstanceArn terra.StringValue `hcl:"instance_arn,attr" validate:"required"`
	// Attribute: min=1
	Attribute []ssoadmininstanceaccesscontrolattributes.Attribute `hcl:"attribute,block" validate:"min=1"`
}
type ssoadminInstanceAccessControlAttributesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ssoadmin_instance_access_control_attributes.
func (siaca ssoadminInstanceAccessControlAttributesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(siaca.ref.Append("id"))
}

// InstanceArn returns a reference to field instance_arn of aws_ssoadmin_instance_access_control_attributes.
func (siaca ssoadminInstanceAccessControlAttributesAttributes) InstanceArn() terra.StringValue {
	return terra.ReferenceAsString(siaca.ref.Append("instance_arn"))
}

// Status returns a reference to field status of aws_ssoadmin_instance_access_control_attributes.
func (siaca ssoadminInstanceAccessControlAttributesAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(siaca.ref.Append("status"))
}

// StatusReason returns a reference to field status_reason of aws_ssoadmin_instance_access_control_attributes.
func (siaca ssoadminInstanceAccessControlAttributesAttributes) StatusReason() terra.StringValue {
	return terra.ReferenceAsString(siaca.ref.Append("status_reason"))
}

func (siaca ssoadminInstanceAccessControlAttributesAttributes) Attribute() terra.SetValue[ssoadmininstanceaccesscontrolattributes.AttributeAttributes] {
	return terra.ReferenceAsSet[ssoadmininstanceaccesscontrolattributes.AttributeAttributes](siaca.ref.Append("attribute"))
}

type ssoadminInstanceAccessControlAttributesState struct {
	Id           string                                                   `json:"id"`
	InstanceArn  string                                                   `json:"instance_arn"`
	Status       string                                                   `json:"status"`
	StatusReason string                                                   `json:"status_reason"`
	Attribute    []ssoadmininstanceaccesscontrolattributes.AttributeState `json:"attribute"`
}
