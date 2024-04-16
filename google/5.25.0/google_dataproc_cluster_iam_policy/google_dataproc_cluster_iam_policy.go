// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataproc_cluster_iam_policy

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_dataproc_cluster_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataprocClusterIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdcip *Resource) Type() string {
	return "google_dataproc_cluster_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gdcip *Resource) LocalName() string {
	return gdcip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdcip *Resource) Configuration() interface{} {
	return gdcip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdcip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdcip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdcip *Resource) Dependencies() terra.Dependencies {
	return gdcip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdcip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdcip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdcip *Resource) Attributes() googleDataprocClusterIamPolicyAttributes {
	return googleDataprocClusterIamPolicyAttributes{ref: terra.ReferenceResource(gdcip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdcip *Resource) ImportState(state io.Reader) error {
	gdcip.state = &googleDataprocClusterIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gdcip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdcip.Type(), gdcip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdcip *Resource) State() (*googleDataprocClusterIamPolicyState, bool) {
	return gdcip.state, gdcip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdcip *Resource) StateMust() *googleDataprocClusterIamPolicyState {
	if gdcip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdcip.Type(), gdcip.LocalName()))
	}
	return gdcip.state
}

// Args contains the configurations for google_dataproc_cluster_iam_policy.
type Args struct {
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

type googleDataprocClusterIamPolicyAttributes struct {
	ref terra.Reference
}

// Cluster returns a reference to field cluster of google_dataproc_cluster_iam_policy.
func (gdcip googleDataprocClusterIamPolicyAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(gdcip.ref.Append("cluster"))
}

// Etag returns a reference to field etag of google_dataproc_cluster_iam_policy.
func (gdcip googleDataprocClusterIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gdcip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_cluster_iam_policy.
func (gdcip googleDataprocClusterIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdcip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_dataproc_cluster_iam_policy.
func (gdcip googleDataprocClusterIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gdcip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataproc_cluster_iam_policy.
func (gdcip googleDataprocClusterIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdcip.ref.Append("project"))
}

// Region returns a reference to field region of google_dataproc_cluster_iam_policy.
func (gdcip googleDataprocClusterIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gdcip.ref.Append("region"))
}

type googleDataprocClusterIamPolicyState struct {
	Cluster    string `json:"cluster"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Region     string `json:"region"`
}
