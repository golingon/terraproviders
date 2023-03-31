// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApigeeEnvironmentIamPolicy(name string, args ApigeeEnvironmentIamPolicyArgs) *ApigeeEnvironmentIamPolicy {
	return &ApigeeEnvironmentIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEnvironmentIamPolicy)(nil)

type ApigeeEnvironmentIamPolicy struct {
	Name  string
	Args  ApigeeEnvironmentIamPolicyArgs
	state *apigeeEnvironmentIamPolicyState
}

func (aeip *ApigeeEnvironmentIamPolicy) Type() string {
	return "google_apigee_environment_iam_policy"
}

func (aeip *ApigeeEnvironmentIamPolicy) LocalName() string {
	return aeip.Name
}

func (aeip *ApigeeEnvironmentIamPolicy) Configuration() interface{} {
	return aeip.Args
}

func (aeip *ApigeeEnvironmentIamPolicy) Attributes() apigeeEnvironmentIamPolicyAttributes {
	return apigeeEnvironmentIamPolicyAttributes{ref: terra.ReferenceResource(aeip)}
}

func (aeip *ApigeeEnvironmentIamPolicy) ImportState(av io.Reader) error {
	aeip.state = &apigeeEnvironmentIamPolicyState{}
	if err := json.NewDecoder(av).Decode(aeip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aeip.Type(), aeip.LocalName(), err)
	}
	return nil
}

func (aeip *ApigeeEnvironmentIamPolicy) State() (*apigeeEnvironmentIamPolicyState, bool) {
	return aeip.state, aeip.state != nil
}

func (aeip *ApigeeEnvironmentIamPolicy) StateMust() *apigeeEnvironmentIamPolicyState {
	if aeip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aeip.Type(), aeip.LocalName()))
	}
	return aeip.state
}

func (aeip *ApigeeEnvironmentIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(aeip)
}

type ApigeeEnvironmentIamPolicyArgs struct {
	// EnvId: string, required
	EnvId terra.StringValue `hcl:"env_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// DependsOn contains resources that ApigeeEnvironmentIamPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apigeeEnvironmentIamPolicyAttributes struct {
	ref terra.Reference
}

func (aeip apigeeEnvironmentIamPolicyAttributes) EnvId() terra.StringValue {
	return terra.ReferenceString(aeip.ref.Append("env_id"))
}

func (aeip apigeeEnvironmentIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(aeip.ref.Append("etag"))
}

func (aeip apigeeEnvironmentIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(aeip.ref.Append("id"))
}

func (aeip apigeeEnvironmentIamPolicyAttributes) OrgId() terra.StringValue {
	return terra.ReferenceString(aeip.ref.Append("org_id"))
}

func (aeip apigeeEnvironmentIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceString(aeip.ref.Append("policy_data"))
}

type apigeeEnvironmentIamPolicyState struct {
	EnvId      string `json:"env_id"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	OrgId      string `json:"org_id"`
	PolicyData string `json:"policy_data"`
}
