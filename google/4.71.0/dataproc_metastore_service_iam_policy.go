// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocMetastoreServiceIamPolicy creates a new instance of [DataprocMetastoreServiceIamPolicy].
func NewDataprocMetastoreServiceIamPolicy(name string, args DataprocMetastoreServiceIamPolicyArgs) *DataprocMetastoreServiceIamPolicy {
	return &DataprocMetastoreServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocMetastoreServiceIamPolicy)(nil)

// DataprocMetastoreServiceIamPolicy represents the Terraform resource google_dataproc_metastore_service_iam_policy.
type DataprocMetastoreServiceIamPolicy struct {
	Name      string
	Args      DataprocMetastoreServiceIamPolicyArgs
	state     *dataprocMetastoreServiceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocMetastoreServiceIamPolicy].
func (dmsip *DataprocMetastoreServiceIamPolicy) Type() string {
	return "google_dataproc_metastore_service_iam_policy"
}

// LocalName returns the local name for [DataprocMetastoreServiceIamPolicy].
func (dmsip *DataprocMetastoreServiceIamPolicy) LocalName() string {
	return dmsip.Name
}

// Configuration returns the configuration (args) for [DataprocMetastoreServiceIamPolicy].
func (dmsip *DataprocMetastoreServiceIamPolicy) Configuration() interface{} {
	return dmsip.Args
}

// DependOn is used for other resources to depend on [DataprocMetastoreServiceIamPolicy].
func (dmsip *DataprocMetastoreServiceIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dmsip)
}

// Dependencies returns the list of resources [DataprocMetastoreServiceIamPolicy] depends_on.
func (dmsip *DataprocMetastoreServiceIamPolicy) Dependencies() terra.Dependencies {
	return dmsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocMetastoreServiceIamPolicy].
func (dmsip *DataprocMetastoreServiceIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return dmsip.Lifecycle
}

// Attributes returns the attributes for [DataprocMetastoreServiceIamPolicy].
func (dmsip *DataprocMetastoreServiceIamPolicy) Attributes() dataprocMetastoreServiceIamPolicyAttributes {
	return dataprocMetastoreServiceIamPolicyAttributes{ref: terra.ReferenceResource(dmsip)}
}

// ImportState imports the given attribute values into [DataprocMetastoreServiceIamPolicy]'s state.
func (dmsip *DataprocMetastoreServiceIamPolicy) ImportState(av io.Reader) error {
	dmsip.state = &dataprocMetastoreServiceIamPolicyState{}
	if err := json.NewDecoder(av).Decode(dmsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmsip.Type(), dmsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocMetastoreServiceIamPolicy] has state.
func (dmsip *DataprocMetastoreServiceIamPolicy) State() (*dataprocMetastoreServiceIamPolicyState, bool) {
	return dmsip.state, dmsip.state != nil
}

// StateMust returns the state for [DataprocMetastoreServiceIamPolicy]. Panics if the state is nil.
func (dmsip *DataprocMetastoreServiceIamPolicy) StateMust() *dataprocMetastoreServiceIamPolicyState {
	if dmsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmsip.Type(), dmsip.LocalName()))
	}
	return dmsip.state
}

// DataprocMetastoreServiceIamPolicyArgs contains the configurations for google_dataproc_metastore_service_iam_policy.
type DataprocMetastoreServiceIamPolicyArgs struct {
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
type dataprocMetastoreServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_metastore_service_iam_policy.
func (dmsip dataprocMetastoreServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_metastore_service_iam_policy.
func (dmsip dataprocMetastoreServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_metastore_service_iam_policy.
func (dmsip dataprocMetastoreServiceIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataproc_metastore_service_iam_policy.
func (dmsip dataprocMetastoreServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataproc_metastore_service_iam_policy.
func (dmsip dataprocMetastoreServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("project"))
}

// ServiceId returns a reference to field service_id of google_dataproc_metastore_service_iam_policy.
func (dmsip dataprocMetastoreServiceIamPolicyAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(dmsip.ref.Append("service_id"))
}

type dataprocMetastoreServiceIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	ServiceId  string `json:"service_id"`
}
