// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeSubnetworkIamPolicy(name string, args ComputeSubnetworkIamPolicyArgs) *ComputeSubnetworkIamPolicy {
	return &ComputeSubnetworkIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSubnetworkIamPolicy)(nil)

type ComputeSubnetworkIamPolicy struct {
	Name  string
	Args  ComputeSubnetworkIamPolicyArgs
	state *computeSubnetworkIamPolicyState
}

func (csip *ComputeSubnetworkIamPolicy) Type() string {
	return "google_compute_subnetwork_iam_policy"
}

func (csip *ComputeSubnetworkIamPolicy) LocalName() string {
	return csip.Name
}

func (csip *ComputeSubnetworkIamPolicy) Configuration() interface{} {
	return csip.Args
}

func (csip *ComputeSubnetworkIamPolicy) Attributes() computeSubnetworkIamPolicyAttributes {
	return computeSubnetworkIamPolicyAttributes{ref: terra.ReferenceResource(csip)}
}

func (csip *ComputeSubnetworkIamPolicy) ImportState(av io.Reader) error {
	csip.state = &computeSubnetworkIamPolicyState{}
	if err := json.NewDecoder(av).Decode(csip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csip.Type(), csip.LocalName(), err)
	}
	return nil
}

func (csip *ComputeSubnetworkIamPolicy) State() (*computeSubnetworkIamPolicyState, bool) {
	return csip.state, csip.state != nil
}

func (csip *ComputeSubnetworkIamPolicy) StateMust() *computeSubnetworkIamPolicyState {
	if csip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csip.Type(), csip.LocalName()))
	}
	return csip.state
}

func (csip *ComputeSubnetworkIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(csip)
}

type ComputeSubnetworkIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Subnetwork: string, required
	Subnetwork terra.StringValue `hcl:"subnetwork,attr" validate:"required"`
	// DependsOn contains resources that ComputeSubnetworkIamPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeSubnetworkIamPolicyAttributes struct {
	ref terra.Reference
}

func (csip computeSubnetworkIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(csip.ref.Append("etag"))
}

func (csip computeSubnetworkIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(csip.ref.Append("id"))
}

func (csip computeSubnetworkIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceString(csip.ref.Append("policy_data"))
}

func (csip computeSubnetworkIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceString(csip.ref.Append("project"))
}

func (csip computeSubnetworkIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceString(csip.ref.Append("region"))
}

func (csip computeSubnetworkIamPolicyAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceString(csip.ref.Append("subnetwork"))
}

type computeSubnetworkIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Region     string `json:"region"`
	Subnetwork string `json:"subnetwork"`
}
