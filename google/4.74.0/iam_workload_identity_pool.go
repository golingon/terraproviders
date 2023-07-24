// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iamworkloadidentitypool "github.com/golingon/terraproviders/google/4.74.0/iamworkloadidentitypool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIamWorkloadIdentityPool creates a new instance of [IamWorkloadIdentityPool].
func NewIamWorkloadIdentityPool(name string, args IamWorkloadIdentityPoolArgs) *IamWorkloadIdentityPool {
	return &IamWorkloadIdentityPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamWorkloadIdentityPool)(nil)

// IamWorkloadIdentityPool represents the Terraform resource google_iam_workload_identity_pool.
type IamWorkloadIdentityPool struct {
	Name      string
	Args      IamWorkloadIdentityPoolArgs
	state     *iamWorkloadIdentityPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamWorkloadIdentityPool].
func (iwip *IamWorkloadIdentityPool) Type() string {
	return "google_iam_workload_identity_pool"
}

// LocalName returns the local name for [IamWorkloadIdentityPool].
func (iwip *IamWorkloadIdentityPool) LocalName() string {
	return iwip.Name
}

// Configuration returns the configuration (args) for [IamWorkloadIdentityPool].
func (iwip *IamWorkloadIdentityPool) Configuration() interface{} {
	return iwip.Args
}

// DependOn is used for other resources to depend on [IamWorkloadIdentityPool].
func (iwip *IamWorkloadIdentityPool) DependOn() terra.Reference {
	return terra.ReferenceResource(iwip)
}

// Dependencies returns the list of resources [IamWorkloadIdentityPool] depends_on.
func (iwip *IamWorkloadIdentityPool) Dependencies() terra.Dependencies {
	return iwip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamWorkloadIdentityPool].
func (iwip *IamWorkloadIdentityPool) LifecycleManagement() *terra.Lifecycle {
	return iwip.Lifecycle
}

// Attributes returns the attributes for [IamWorkloadIdentityPool].
func (iwip *IamWorkloadIdentityPool) Attributes() iamWorkloadIdentityPoolAttributes {
	return iamWorkloadIdentityPoolAttributes{ref: terra.ReferenceResource(iwip)}
}

// ImportState imports the given attribute values into [IamWorkloadIdentityPool]'s state.
func (iwip *IamWorkloadIdentityPool) ImportState(av io.Reader) error {
	iwip.state = &iamWorkloadIdentityPoolState{}
	if err := json.NewDecoder(av).Decode(iwip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwip.Type(), iwip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamWorkloadIdentityPool] has state.
func (iwip *IamWorkloadIdentityPool) State() (*iamWorkloadIdentityPoolState, bool) {
	return iwip.state, iwip.state != nil
}

// StateMust returns the state for [IamWorkloadIdentityPool]. Panics if the state is nil.
func (iwip *IamWorkloadIdentityPool) StateMust() *iamWorkloadIdentityPoolState {
	if iwip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwip.Type(), iwip.LocalName()))
	}
	return iwip.state
}

// IamWorkloadIdentityPoolArgs contains the configurations for google_iam_workload_identity_pool.
type IamWorkloadIdentityPoolArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// WorkloadIdentityPoolId: string, required
	WorkloadIdentityPoolId terra.StringValue `hcl:"workload_identity_pool_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *iamworkloadidentitypool.Timeouts `hcl:"timeouts,block"`
}
type iamWorkloadIdentityPoolAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_iam_workload_identity_pool.
func (iwip iamWorkloadIdentityPoolAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_iam_workload_identity_pool.
func (iwip iamWorkloadIdentityPoolAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(iwip.ref.Append("disabled"))
}

// DisplayName returns a reference to field display_name of google_iam_workload_identity_pool.
func (iwip iamWorkloadIdentityPoolAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("display_name"))
}

// Id returns a reference to field id of google_iam_workload_identity_pool.
func (iwip iamWorkloadIdentityPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("id"))
}

// Name returns a reference to field name of google_iam_workload_identity_pool.
func (iwip iamWorkloadIdentityPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("name"))
}

// Project returns a reference to field project of google_iam_workload_identity_pool.
func (iwip iamWorkloadIdentityPoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("project"))
}

// State returns a reference to field state of google_iam_workload_identity_pool.
func (iwip iamWorkloadIdentityPoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("state"))
}

// WorkloadIdentityPoolId returns a reference to field workload_identity_pool_id of google_iam_workload_identity_pool.
func (iwip iamWorkloadIdentityPoolAttributes) WorkloadIdentityPoolId() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("workload_identity_pool_id"))
}

func (iwip iamWorkloadIdentityPoolAttributes) Timeouts() iamworkloadidentitypool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iamworkloadidentitypool.TimeoutsAttributes](iwip.ref.Append("timeouts"))
}

type iamWorkloadIdentityPoolState struct {
	Description            string                                 `json:"description"`
	Disabled               bool                                   `json:"disabled"`
	DisplayName            string                                 `json:"display_name"`
	Id                     string                                 `json:"id"`
	Name                   string                                 `json:"name"`
	Project                string                                 `json:"project"`
	State                  string                                 `json:"state"`
	WorkloadIdentityPoolId string                                 `json:"workload_identity_pool_id"`
	Timeouts               *iamworkloadidentitypool.TimeoutsState `json:"timeouts"`
}
