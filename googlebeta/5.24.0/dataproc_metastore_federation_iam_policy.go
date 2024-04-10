// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewDataprocMetastoreFederationIamPolicy creates a new instance of [DataprocMetastoreFederationIamPolicy].
func NewDataprocMetastoreFederationIamPolicy(name string, args DataprocMetastoreFederationIamPolicyArgs) *DataprocMetastoreFederationIamPolicy {
	return &DataprocMetastoreFederationIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocMetastoreFederationIamPolicy)(nil)

// DataprocMetastoreFederationIamPolicy represents the Terraform resource google_dataproc_metastore_federation_iam_policy.
type DataprocMetastoreFederationIamPolicy struct {
	Name      string
	Args      DataprocMetastoreFederationIamPolicyArgs
	state     *dataprocMetastoreFederationIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocMetastoreFederationIamPolicy].
func (dmfip *DataprocMetastoreFederationIamPolicy) Type() string {
	return "google_dataproc_metastore_federation_iam_policy"
}

// LocalName returns the local name for [DataprocMetastoreFederationIamPolicy].
func (dmfip *DataprocMetastoreFederationIamPolicy) LocalName() string {
	return dmfip.Name
}

// Configuration returns the configuration (args) for [DataprocMetastoreFederationIamPolicy].
func (dmfip *DataprocMetastoreFederationIamPolicy) Configuration() interface{} {
	return dmfip.Args
}

// DependOn is used for other resources to depend on [DataprocMetastoreFederationIamPolicy].
func (dmfip *DataprocMetastoreFederationIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dmfip)
}

// Dependencies returns the list of resources [DataprocMetastoreFederationIamPolicy] depends_on.
func (dmfip *DataprocMetastoreFederationIamPolicy) Dependencies() terra.Dependencies {
	return dmfip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocMetastoreFederationIamPolicy].
func (dmfip *DataprocMetastoreFederationIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return dmfip.Lifecycle
}

// Attributes returns the attributes for [DataprocMetastoreFederationIamPolicy].
func (dmfip *DataprocMetastoreFederationIamPolicy) Attributes() dataprocMetastoreFederationIamPolicyAttributes {
	return dataprocMetastoreFederationIamPolicyAttributes{ref: terra.ReferenceResource(dmfip)}
}

// ImportState imports the given attribute values into [DataprocMetastoreFederationIamPolicy]'s state.
func (dmfip *DataprocMetastoreFederationIamPolicy) ImportState(av io.Reader) error {
	dmfip.state = &dataprocMetastoreFederationIamPolicyState{}
	if err := json.NewDecoder(av).Decode(dmfip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmfip.Type(), dmfip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocMetastoreFederationIamPolicy] has state.
func (dmfip *DataprocMetastoreFederationIamPolicy) State() (*dataprocMetastoreFederationIamPolicyState, bool) {
	return dmfip.state, dmfip.state != nil
}

// StateMust returns the state for [DataprocMetastoreFederationIamPolicy]. Panics if the state is nil.
func (dmfip *DataprocMetastoreFederationIamPolicy) StateMust() *dataprocMetastoreFederationIamPolicyState {
	if dmfip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmfip.Type(), dmfip.LocalName()))
	}
	return dmfip.state
}

// DataprocMetastoreFederationIamPolicyArgs contains the configurations for google_dataproc_metastore_federation_iam_policy.
type DataprocMetastoreFederationIamPolicyArgs struct {
	// FederationId: string, required
	FederationId terra.StringValue `hcl:"federation_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataprocMetastoreFederationIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_metastore_federation_iam_policy.
func (dmfip dataprocMetastoreFederationIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dmfip.ref.Append("etag"))
}

// FederationId returns a reference to field federation_id of google_dataproc_metastore_federation_iam_policy.
func (dmfip dataprocMetastoreFederationIamPolicyAttributes) FederationId() terra.StringValue {
	return terra.ReferenceAsString(dmfip.ref.Append("federation_id"))
}

// Id returns a reference to field id of google_dataproc_metastore_federation_iam_policy.
func (dmfip dataprocMetastoreFederationIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmfip.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_metastore_federation_iam_policy.
func (dmfip dataprocMetastoreFederationIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dmfip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataproc_metastore_federation_iam_policy.
func (dmfip dataprocMetastoreFederationIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dmfip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataproc_metastore_federation_iam_policy.
func (dmfip dataprocMetastoreFederationIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmfip.ref.Append("project"))
}

type dataprocMetastoreFederationIamPolicyState struct {
	Etag         string `json:"etag"`
	FederationId string `json:"federation_id"`
	Id           string `json:"id"`
	Location     string `json:"location"`
	PolicyData   string `json:"policy_data"`
	Project      string `json:"project"`
}
