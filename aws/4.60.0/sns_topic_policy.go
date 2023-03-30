// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSnsTopicPolicy(name string, args SnsTopicPolicyArgs) *SnsTopicPolicy {
	return &SnsTopicPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SnsTopicPolicy)(nil)

type SnsTopicPolicy struct {
	Name  string
	Args  SnsTopicPolicyArgs
	state *snsTopicPolicyState
}

func (stp *SnsTopicPolicy) Type() string {
	return "aws_sns_topic_policy"
}

func (stp *SnsTopicPolicy) LocalName() string {
	return stp.Name
}

func (stp *SnsTopicPolicy) Configuration() interface{} {
	return stp.Args
}

func (stp *SnsTopicPolicy) Attributes() snsTopicPolicyAttributes {
	return snsTopicPolicyAttributes{ref: terra.ReferenceResource(stp)}
}

func (stp *SnsTopicPolicy) ImportState(av io.Reader) error {
	stp.state = &snsTopicPolicyState{}
	if err := json.NewDecoder(av).Decode(stp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", stp.Type(), stp.LocalName(), err)
	}
	return nil
}

func (stp *SnsTopicPolicy) State() (*snsTopicPolicyState, bool) {
	return stp.state, stp.state != nil
}

func (stp *SnsTopicPolicy) StateMust() *snsTopicPolicyState {
	if stp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", stp.Type(), stp.LocalName()))
	}
	return stp.state
}

func (stp *SnsTopicPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(stp)
}

type SnsTopicPolicyArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// DependsOn contains resources that SnsTopicPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type snsTopicPolicyAttributes struct {
	ref terra.Reference
}

func (stp snsTopicPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(stp.ref.Append("arn"))
}

func (stp snsTopicPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(stp.ref.Append("id"))
}

func (stp snsTopicPolicyAttributes) Owner() terra.StringValue {
	return terra.ReferenceString(stp.ref.Append("owner"))
}

func (stp snsTopicPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceString(stp.ref.Append("policy"))
}

type snsTopicPolicyState struct {
	Arn    string `json:"arn"`
	Id     string `json:"id"`
	Owner  string `json:"owner"`
	Policy string `json:"policy"`
}
