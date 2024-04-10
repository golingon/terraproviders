// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewDataFusionInstanceIamPolicy creates a new instance of [DataFusionInstanceIamPolicy].
func NewDataFusionInstanceIamPolicy(name string, args DataFusionInstanceIamPolicyArgs) *DataFusionInstanceIamPolicy {
	return &DataFusionInstanceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFusionInstanceIamPolicy)(nil)

// DataFusionInstanceIamPolicy represents the Terraform resource google_data_fusion_instance_iam_policy.
type DataFusionInstanceIamPolicy struct {
	Name      string
	Args      DataFusionInstanceIamPolicyArgs
	state     *dataFusionInstanceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFusionInstanceIamPolicy].
func (dfiip *DataFusionInstanceIamPolicy) Type() string {
	return "google_data_fusion_instance_iam_policy"
}

// LocalName returns the local name for [DataFusionInstanceIamPolicy].
func (dfiip *DataFusionInstanceIamPolicy) LocalName() string {
	return dfiip.Name
}

// Configuration returns the configuration (args) for [DataFusionInstanceIamPolicy].
func (dfiip *DataFusionInstanceIamPolicy) Configuration() interface{} {
	return dfiip.Args
}

// DependOn is used for other resources to depend on [DataFusionInstanceIamPolicy].
func (dfiip *DataFusionInstanceIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dfiip)
}

// Dependencies returns the list of resources [DataFusionInstanceIamPolicy] depends_on.
func (dfiip *DataFusionInstanceIamPolicy) Dependencies() terra.Dependencies {
	return dfiip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFusionInstanceIamPolicy].
func (dfiip *DataFusionInstanceIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return dfiip.Lifecycle
}

// Attributes returns the attributes for [DataFusionInstanceIamPolicy].
func (dfiip *DataFusionInstanceIamPolicy) Attributes() dataFusionInstanceIamPolicyAttributes {
	return dataFusionInstanceIamPolicyAttributes{ref: terra.ReferenceResource(dfiip)}
}

// ImportState imports the given attribute values into [DataFusionInstanceIamPolicy]'s state.
func (dfiip *DataFusionInstanceIamPolicy) ImportState(av io.Reader) error {
	dfiip.state = &dataFusionInstanceIamPolicyState{}
	if err := json.NewDecoder(av).Decode(dfiip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfiip.Type(), dfiip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFusionInstanceIamPolicy] has state.
func (dfiip *DataFusionInstanceIamPolicy) State() (*dataFusionInstanceIamPolicyState, bool) {
	return dfiip.state, dfiip.state != nil
}

// StateMust returns the state for [DataFusionInstanceIamPolicy]. Panics if the state is nil.
func (dfiip *DataFusionInstanceIamPolicy) StateMust() *dataFusionInstanceIamPolicyState {
	if dfiip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfiip.Type(), dfiip.LocalName()))
	}
	return dfiip.state
}

// DataFusionInstanceIamPolicyArgs contains the configurations for google_data_fusion_instance_iam_policy.
type DataFusionInstanceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataFusionInstanceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_fusion_instance_iam_policy.
func (dfiip dataFusionInstanceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_fusion_instance_iam_policy.
func (dfiip dataFusionInstanceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("id"))
}

// Name returns a reference to field name of google_data_fusion_instance_iam_policy.
func (dfiip dataFusionInstanceIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_data_fusion_instance_iam_policy.
func (dfiip dataFusionInstanceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_data_fusion_instance_iam_policy.
func (dfiip dataFusionInstanceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("project"))
}

// Region returns a reference to field region of google_data_fusion_instance_iam_policy.
func (dfiip dataFusionInstanceIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("region"))
}

type dataFusionInstanceIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Region     string `json:"region"`
}
