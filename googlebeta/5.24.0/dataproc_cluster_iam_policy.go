// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewDataprocClusterIamPolicy creates a new instance of [DataprocClusterIamPolicy].
func NewDataprocClusterIamPolicy(name string, args DataprocClusterIamPolicyArgs) *DataprocClusterIamPolicy {
	return &DataprocClusterIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocClusterIamPolicy)(nil)

// DataprocClusterIamPolicy represents the Terraform resource google_dataproc_cluster_iam_policy.
type DataprocClusterIamPolicy struct {
	Name      string
	Args      DataprocClusterIamPolicyArgs
	state     *dataprocClusterIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocClusterIamPolicy].
func (dcip *DataprocClusterIamPolicy) Type() string {
	return "google_dataproc_cluster_iam_policy"
}

// LocalName returns the local name for [DataprocClusterIamPolicy].
func (dcip *DataprocClusterIamPolicy) LocalName() string {
	return dcip.Name
}

// Configuration returns the configuration (args) for [DataprocClusterIamPolicy].
func (dcip *DataprocClusterIamPolicy) Configuration() interface{} {
	return dcip.Args
}

// DependOn is used for other resources to depend on [DataprocClusterIamPolicy].
func (dcip *DataprocClusterIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dcip)
}

// Dependencies returns the list of resources [DataprocClusterIamPolicy] depends_on.
func (dcip *DataprocClusterIamPolicy) Dependencies() terra.Dependencies {
	return dcip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocClusterIamPolicy].
func (dcip *DataprocClusterIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return dcip.Lifecycle
}

// Attributes returns the attributes for [DataprocClusterIamPolicy].
func (dcip *DataprocClusterIamPolicy) Attributes() dataprocClusterIamPolicyAttributes {
	return dataprocClusterIamPolicyAttributes{ref: terra.ReferenceResource(dcip)}
}

// ImportState imports the given attribute values into [DataprocClusterIamPolicy]'s state.
func (dcip *DataprocClusterIamPolicy) ImportState(av io.Reader) error {
	dcip.state = &dataprocClusterIamPolicyState{}
	if err := json.NewDecoder(av).Decode(dcip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcip.Type(), dcip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocClusterIamPolicy] has state.
func (dcip *DataprocClusterIamPolicy) State() (*dataprocClusterIamPolicyState, bool) {
	return dcip.state, dcip.state != nil
}

// StateMust returns the state for [DataprocClusterIamPolicy]. Panics if the state is nil.
func (dcip *DataprocClusterIamPolicy) StateMust() *dataprocClusterIamPolicyState {
	if dcip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcip.Type(), dcip.LocalName()))
	}
	return dcip.state
}

// DataprocClusterIamPolicyArgs contains the configurations for google_dataproc_cluster_iam_policy.
type DataprocClusterIamPolicyArgs struct {
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataprocClusterIamPolicyAttributes struct {
	ref terra.Reference
}

// Cluster returns a reference to field cluster of google_dataproc_cluster_iam_policy.
func (dcip dataprocClusterIamPolicyAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(dcip.ref.Append("cluster"))
}

// Etag returns a reference to field etag of google_dataproc_cluster_iam_policy.
func (dcip dataprocClusterIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_cluster_iam_policy.
func (dcip dataprocClusterIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_dataproc_cluster_iam_policy.
func (dcip dataprocClusterIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dcip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataproc_cluster_iam_policy.
func (dcip dataprocClusterIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dcip.ref.Append("project"))
}

// Region returns a reference to field region of google_dataproc_cluster_iam_policy.
func (dcip dataprocClusterIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dcip.ref.Append("region"))
}

type dataprocClusterIamPolicyState struct {
	Cluster    string `json:"cluster"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Region     string `json:"region"`
}
