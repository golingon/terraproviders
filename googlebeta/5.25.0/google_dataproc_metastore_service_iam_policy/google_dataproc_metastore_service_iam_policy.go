// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataproc_metastore_service_iam_policy

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

// Resource represents the Terraform resource google_dataproc_metastore_service_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataprocMetastoreServiceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdmsip *Resource) Type() string {
	return "google_dataproc_metastore_service_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gdmsip *Resource) LocalName() string {
	return gdmsip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdmsip *Resource) Configuration() interface{} {
	return gdmsip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdmsip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdmsip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdmsip *Resource) Dependencies() terra.Dependencies {
	return gdmsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdmsip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdmsip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdmsip *Resource) Attributes() googleDataprocMetastoreServiceIamPolicyAttributes {
	return googleDataprocMetastoreServiceIamPolicyAttributes{ref: terra.ReferenceResource(gdmsip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdmsip *Resource) ImportState(state io.Reader) error {
	gdmsip.state = &googleDataprocMetastoreServiceIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gdmsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdmsip.Type(), gdmsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdmsip *Resource) State() (*googleDataprocMetastoreServiceIamPolicyState, bool) {
	return gdmsip.state, gdmsip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdmsip *Resource) StateMust() *googleDataprocMetastoreServiceIamPolicyState {
	if gdmsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdmsip.Type(), gdmsip.LocalName()))
	}
	return gdmsip.state
}

// Args contains the configurations for google_dataproc_metastore_service_iam_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceId: string, required
	ServiceId terra.StringValue `hcl:"service_id,attr" validate:"required"`
}

type googleDataprocMetastoreServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_metastore_service_iam_policy.
func (gdmsip googleDataprocMetastoreServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gdmsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_metastore_service_iam_policy.
func (gdmsip googleDataprocMetastoreServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdmsip.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_metastore_service_iam_policy.
func (gdmsip googleDataprocMetastoreServiceIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gdmsip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataproc_metastore_service_iam_policy.
func (gdmsip googleDataprocMetastoreServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gdmsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataproc_metastore_service_iam_policy.
func (gdmsip googleDataprocMetastoreServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdmsip.ref.Append("project"))
}

// ServiceId returns a reference to field service_id of google_dataproc_metastore_service_iam_policy.
func (gdmsip googleDataprocMetastoreServiceIamPolicyAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(gdmsip.ref.Append("service_id"))
}

type googleDataprocMetastoreServiceIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	ServiceId  string `json:"service_id"`
}
