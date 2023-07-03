// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryConnectionIamPolicy creates a new instance of [BigqueryConnectionIamPolicy].
func NewBigqueryConnectionIamPolicy(name string, args BigqueryConnectionIamPolicyArgs) *BigqueryConnectionIamPolicy {
	return &BigqueryConnectionIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryConnectionIamPolicy)(nil)

// BigqueryConnectionIamPolicy represents the Terraform resource google_bigquery_connection_iam_policy.
type BigqueryConnectionIamPolicy struct {
	Name      string
	Args      BigqueryConnectionIamPolicyArgs
	state     *bigqueryConnectionIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryConnectionIamPolicy].
func (bcip *BigqueryConnectionIamPolicy) Type() string {
	return "google_bigquery_connection_iam_policy"
}

// LocalName returns the local name for [BigqueryConnectionIamPolicy].
func (bcip *BigqueryConnectionIamPolicy) LocalName() string {
	return bcip.Name
}

// Configuration returns the configuration (args) for [BigqueryConnectionIamPolicy].
func (bcip *BigqueryConnectionIamPolicy) Configuration() interface{} {
	return bcip.Args
}

// DependOn is used for other resources to depend on [BigqueryConnectionIamPolicy].
func (bcip *BigqueryConnectionIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(bcip)
}

// Dependencies returns the list of resources [BigqueryConnectionIamPolicy] depends_on.
func (bcip *BigqueryConnectionIamPolicy) Dependencies() terra.Dependencies {
	return bcip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryConnectionIamPolicy].
func (bcip *BigqueryConnectionIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return bcip.Lifecycle
}

// Attributes returns the attributes for [BigqueryConnectionIamPolicy].
func (bcip *BigqueryConnectionIamPolicy) Attributes() bigqueryConnectionIamPolicyAttributes {
	return bigqueryConnectionIamPolicyAttributes{ref: terra.ReferenceResource(bcip)}
}

// ImportState imports the given attribute values into [BigqueryConnectionIamPolicy]'s state.
func (bcip *BigqueryConnectionIamPolicy) ImportState(av io.Reader) error {
	bcip.state = &bigqueryConnectionIamPolicyState{}
	if err := json.NewDecoder(av).Decode(bcip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcip.Type(), bcip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryConnectionIamPolicy] has state.
func (bcip *BigqueryConnectionIamPolicy) State() (*bigqueryConnectionIamPolicyState, bool) {
	return bcip.state, bcip.state != nil
}

// StateMust returns the state for [BigqueryConnectionIamPolicy]. Panics if the state is nil.
func (bcip *BigqueryConnectionIamPolicy) StateMust() *bigqueryConnectionIamPolicyState {
	if bcip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcip.Type(), bcip.LocalName()))
	}
	return bcip.state
}

// BigqueryConnectionIamPolicyArgs contains the configurations for google_bigquery_connection_iam_policy.
type BigqueryConnectionIamPolicyArgs struct {
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type bigqueryConnectionIamPolicyAttributes struct {
	ref terra.Reference
}

// ConnectionId returns a reference to field connection_id of google_bigquery_connection_iam_policy.
func (bcip bigqueryConnectionIamPolicyAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("connection_id"))
}

// Etag returns a reference to field etag of google_bigquery_connection_iam_policy.
func (bcip bigqueryConnectionIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_connection_iam_policy.
func (bcip bigqueryConnectionIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_connection_iam_policy.
func (bcip bigqueryConnectionIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_bigquery_connection_iam_policy.
func (bcip bigqueryConnectionIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_bigquery_connection_iam_policy.
func (bcip bigqueryConnectionIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("project"))
}

type bigqueryConnectionIamPolicyState struct {
	ConnectionId string `json:"connection_id"`
	Etag         string `json:"etag"`
	Id           string `json:"id"`
	Location     string `json:"location"`
	PolicyData   string `json:"policy_data"`
	Project      string `json:"project"`
}