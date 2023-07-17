// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigquerydatapolicydatapolicy "github.com/golingon/terraproviders/googlebeta/4.72.1/bigquerydatapolicydatapolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryDatapolicyDataPolicy creates a new instance of [BigqueryDatapolicyDataPolicy].
func NewBigqueryDatapolicyDataPolicy(name string, args BigqueryDatapolicyDataPolicyArgs) *BigqueryDatapolicyDataPolicy {
	return &BigqueryDatapolicyDataPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryDatapolicyDataPolicy)(nil)

// BigqueryDatapolicyDataPolicy represents the Terraform resource google_bigquery_datapolicy_data_policy.
type BigqueryDatapolicyDataPolicy struct {
	Name      string
	Args      BigqueryDatapolicyDataPolicyArgs
	state     *bigqueryDatapolicyDataPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryDatapolicyDataPolicy].
func (bddp *BigqueryDatapolicyDataPolicy) Type() string {
	return "google_bigquery_datapolicy_data_policy"
}

// LocalName returns the local name for [BigqueryDatapolicyDataPolicy].
func (bddp *BigqueryDatapolicyDataPolicy) LocalName() string {
	return bddp.Name
}

// Configuration returns the configuration (args) for [BigqueryDatapolicyDataPolicy].
func (bddp *BigqueryDatapolicyDataPolicy) Configuration() interface{} {
	return bddp.Args
}

// DependOn is used for other resources to depend on [BigqueryDatapolicyDataPolicy].
func (bddp *BigqueryDatapolicyDataPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(bddp)
}

// Dependencies returns the list of resources [BigqueryDatapolicyDataPolicy] depends_on.
func (bddp *BigqueryDatapolicyDataPolicy) Dependencies() terra.Dependencies {
	return bddp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryDatapolicyDataPolicy].
func (bddp *BigqueryDatapolicyDataPolicy) LifecycleManagement() *terra.Lifecycle {
	return bddp.Lifecycle
}

// Attributes returns the attributes for [BigqueryDatapolicyDataPolicy].
func (bddp *BigqueryDatapolicyDataPolicy) Attributes() bigqueryDatapolicyDataPolicyAttributes {
	return bigqueryDatapolicyDataPolicyAttributes{ref: terra.ReferenceResource(bddp)}
}

// ImportState imports the given attribute values into [BigqueryDatapolicyDataPolicy]'s state.
func (bddp *BigqueryDatapolicyDataPolicy) ImportState(av io.Reader) error {
	bddp.state = &bigqueryDatapolicyDataPolicyState{}
	if err := json.NewDecoder(av).Decode(bddp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bddp.Type(), bddp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryDatapolicyDataPolicy] has state.
func (bddp *BigqueryDatapolicyDataPolicy) State() (*bigqueryDatapolicyDataPolicyState, bool) {
	return bddp.state, bddp.state != nil
}

// StateMust returns the state for [BigqueryDatapolicyDataPolicy]. Panics if the state is nil.
func (bddp *BigqueryDatapolicyDataPolicy) StateMust() *bigqueryDatapolicyDataPolicyState {
	if bddp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bddp.Type(), bddp.LocalName()))
	}
	return bddp.state
}

// BigqueryDatapolicyDataPolicyArgs contains the configurations for google_bigquery_datapolicy_data_policy.
type BigqueryDatapolicyDataPolicyArgs struct {
	// DataPolicyId: string, required
	DataPolicyId terra.StringValue `hcl:"data_policy_id,attr" validate:"required"`
	// DataPolicyType: string, required
	DataPolicyType terra.StringValue `hcl:"data_policy_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// PolicyTag: string, required
	PolicyTag terra.StringValue `hcl:"policy_tag,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// DataMaskingPolicy: optional
	DataMaskingPolicy *bigquerydatapolicydatapolicy.DataMaskingPolicy `hcl:"data_masking_policy,block"`
	// Timeouts: optional
	Timeouts *bigquerydatapolicydatapolicy.Timeouts `hcl:"timeouts,block"`
}
type bigqueryDatapolicyDataPolicyAttributes struct {
	ref terra.Reference
}

// DataPolicyId returns a reference to field data_policy_id of google_bigquery_datapolicy_data_policy.
func (bddp bigqueryDatapolicyDataPolicyAttributes) DataPolicyId() terra.StringValue {
	return terra.ReferenceAsString(bddp.ref.Append("data_policy_id"))
}

// DataPolicyType returns a reference to field data_policy_type of google_bigquery_datapolicy_data_policy.
func (bddp bigqueryDatapolicyDataPolicyAttributes) DataPolicyType() terra.StringValue {
	return terra.ReferenceAsString(bddp.ref.Append("data_policy_type"))
}

// Id returns a reference to field id of google_bigquery_datapolicy_data_policy.
func (bddp bigqueryDatapolicyDataPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bddp.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_datapolicy_data_policy.
func (bddp bigqueryDatapolicyDataPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bddp.ref.Append("location"))
}

// Name returns a reference to field name of google_bigquery_datapolicy_data_policy.
func (bddp bigqueryDatapolicyDataPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bddp.ref.Append("name"))
}

// PolicyTag returns a reference to field policy_tag of google_bigquery_datapolicy_data_policy.
func (bddp bigqueryDatapolicyDataPolicyAttributes) PolicyTag() terra.StringValue {
	return terra.ReferenceAsString(bddp.ref.Append("policy_tag"))
}

// Project returns a reference to field project of google_bigquery_datapolicy_data_policy.
func (bddp bigqueryDatapolicyDataPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bddp.ref.Append("project"))
}

func (bddp bigqueryDatapolicyDataPolicyAttributes) DataMaskingPolicy() terra.ListValue[bigquerydatapolicydatapolicy.DataMaskingPolicyAttributes] {
	return terra.ReferenceAsList[bigquerydatapolicydatapolicy.DataMaskingPolicyAttributes](bddp.ref.Append("data_masking_policy"))
}

func (bddp bigqueryDatapolicyDataPolicyAttributes) Timeouts() bigquerydatapolicydatapolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigquerydatapolicydatapolicy.TimeoutsAttributes](bddp.ref.Append("timeouts"))
}

type bigqueryDatapolicyDataPolicyState struct {
	DataPolicyId      string                                                `json:"data_policy_id"`
	DataPolicyType    string                                                `json:"data_policy_type"`
	Id                string                                                `json:"id"`
	Location          string                                                `json:"location"`
	Name              string                                                `json:"name"`
	PolicyTag         string                                                `json:"policy_tag"`
	Project           string                                                `json:"project"`
	DataMaskingPolicy []bigquerydatapolicydatapolicy.DataMaskingPolicyState `json:"data_masking_policy"`
	Timeouts          *bigquerydatapolicydatapolicy.TimeoutsState           `json:"timeouts"`
}