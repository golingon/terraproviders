// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCloudfunctionsFunctionIamPolicy(name string, args CloudfunctionsFunctionIamPolicyArgs) *CloudfunctionsFunctionIamPolicy {
	return &CloudfunctionsFunctionIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfunctionsFunctionIamPolicy)(nil)

type CloudfunctionsFunctionIamPolicy struct {
	Name  string
	Args  CloudfunctionsFunctionIamPolicyArgs
	state *cloudfunctionsFunctionIamPolicyState
}

func (cfip *CloudfunctionsFunctionIamPolicy) Type() string {
	return "google_cloudfunctions_function_iam_policy"
}

func (cfip *CloudfunctionsFunctionIamPolicy) LocalName() string {
	return cfip.Name
}

func (cfip *CloudfunctionsFunctionIamPolicy) Configuration() interface{} {
	return cfip.Args
}

func (cfip *CloudfunctionsFunctionIamPolicy) Attributes() cloudfunctionsFunctionIamPolicyAttributes {
	return cloudfunctionsFunctionIamPolicyAttributes{ref: terra.ReferenceResource(cfip)}
}

func (cfip *CloudfunctionsFunctionIamPolicy) ImportState(av io.Reader) error {
	cfip.state = &cloudfunctionsFunctionIamPolicyState{}
	if err := json.NewDecoder(av).Decode(cfip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfip.Type(), cfip.LocalName(), err)
	}
	return nil
}

func (cfip *CloudfunctionsFunctionIamPolicy) State() (*cloudfunctionsFunctionIamPolicyState, bool) {
	return cfip.state, cfip.state != nil
}

func (cfip *CloudfunctionsFunctionIamPolicy) StateMust() *cloudfunctionsFunctionIamPolicyState {
	if cfip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfip.Type(), cfip.LocalName()))
	}
	return cfip.state
}

func (cfip *CloudfunctionsFunctionIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cfip)
}

type CloudfunctionsFunctionIamPolicyArgs struct {
	// CloudFunction: string, required
	CloudFunction terra.StringValue `hcl:"cloud_function,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// DependsOn contains resources that CloudfunctionsFunctionIamPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type cloudfunctionsFunctionIamPolicyAttributes struct {
	ref terra.Reference
}

func (cfip cloudfunctionsFunctionIamPolicyAttributes) CloudFunction() terra.StringValue {
	return terra.ReferenceString(cfip.ref.Append("cloud_function"))
}

func (cfip cloudfunctionsFunctionIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(cfip.ref.Append("etag"))
}

func (cfip cloudfunctionsFunctionIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cfip.ref.Append("id"))
}

func (cfip cloudfunctionsFunctionIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceString(cfip.ref.Append("policy_data"))
}

func (cfip cloudfunctionsFunctionIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cfip.ref.Append("project"))
}

func (cfip cloudfunctionsFunctionIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceString(cfip.ref.Append("region"))
}

type cloudfunctionsFunctionIamPolicyState struct {
	CloudFunction string `json:"cloud_function"`
	Etag          string `json:"etag"`
	Id            string `json:"id"`
	PolicyData    string `json:"policy_data"`
	Project       string `json:"project"`
	Region        string `json:"region"`
}
