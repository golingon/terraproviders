// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexAssetIamPolicy creates a new instance of [DataplexAssetIamPolicy].
func NewDataplexAssetIamPolicy(name string, args DataplexAssetIamPolicyArgs) *DataplexAssetIamPolicy {
	return &DataplexAssetIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexAssetIamPolicy)(nil)

// DataplexAssetIamPolicy represents the Terraform resource google_dataplex_asset_iam_policy.
type DataplexAssetIamPolicy struct {
	Name      string
	Args      DataplexAssetIamPolicyArgs
	state     *dataplexAssetIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexAssetIamPolicy].
func (daip *DataplexAssetIamPolicy) Type() string {
	return "google_dataplex_asset_iam_policy"
}

// LocalName returns the local name for [DataplexAssetIamPolicy].
func (daip *DataplexAssetIamPolicy) LocalName() string {
	return daip.Name
}

// Configuration returns the configuration (args) for [DataplexAssetIamPolicy].
func (daip *DataplexAssetIamPolicy) Configuration() interface{} {
	return daip.Args
}

// DependOn is used for other resources to depend on [DataplexAssetIamPolicy].
func (daip *DataplexAssetIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(daip)
}

// Dependencies returns the list of resources [DataplexAssetIamPolicy] depends_on.
func (daip *DataplexAssetIamPolicy) Dependencies() terra.Dependencies {
	return daip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexAssetIamPolicy].
func (daip *DataplexAssetIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return daip.Lifecycle
}

// Attributes returns the attributes for [DataplexAssetIamPolicy].
func (daip *DataplexAssetIamPolicy) Attributes() dataplexAssetIamPolicyAttributes {
	return dataplexAssetIamPolicyAttributes{ref: terra.ReferenceResource(daip)}
}

// ImportState imports the given attribute values into [DataplexAssetIamPolicy]'s state.
func (daip *DataplexAssetIamPolicy) ImportState(av io.Reader) error {
	daip.state = &dataplexAssetIamPolicyState{}
	if err := json.NewDecoder(av).Decode(daip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", daip.Type(), daip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexAssetIamPolicy] has state.
func (daip *DataplexAssetIamPolicy) State() (*dataplexAssetIamPolicyState, bool) {
	return daip.state, daip.state != nil
}

// StateMust returns the state for [DataplexAssetIamPolicy]. Panics if the state is nil.
func (daip *DataplexAssetIamPolicy) StateMust() *dataplexAssetIamPolicyState {
	if daip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", daip.Type(), daip.LocalName()))
	}
	return daip.state
}

// DataplexAssetIamPolicyArgs contains the configurations for google_dataplex_asset_iam_policy.
type DataplexAssetIamPolicyArgs struct {
	// Asset: string, required
	Asset terra.StringValue `hcl:"asset,attr" validate:"required"`
	// DataplexZone: string, required
	DataplexZone terra.StringValue `hcl:"dataplex_zone,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataplexAssetIamPolicyAttributes struct {
	ref terra.Reference
}

// Asset returns a reference to field asset of google_dataplex_asset_iam_policy.
func (daip dataplexAssetIamPolicyAttributes) Asset() terra.StringValue {
	return terra.ReferenceAsString(daip.ref.Append("asset"))
}

// DataplexZone returns a reference to field dataplex_zone of google_dataplex_asset_iam_policy.
func (daip dataplexAssetIamPolicyAttributes) DataplexZone() terra.StringValue {
	return terra.ReferenceAsString(daip.ref.Append("dataplex_zone"))
}

// Etag returns a reference to field etag of google_dataplex_asset_iam_policy.
func (daip dataplexAssetIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(daip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_asset_iam_policy.
func (daip dataplexAssetIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(daip.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_asset_iam_policy.
func (daip dataplexAssetIamPolicyAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(daip.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_asset_iam_policy.
func (daip dataplexAssetIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(daip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataplex_asset_iam_policy.
func (daip dataplexAssetIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(daip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_dataplex_asset_iam_policy.
func (daip dataplexAssetIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(daip.ref.Append("project"))
}

type dataplexAssetIamPolicyState struct {
	Asset        string `json:"asset"`
	DataplexZone string `json:"dataplex_zone"`
	Etag         string `json:"etag"`
	Id           string `json:"id"`
	Lake         string `json:"lake"`
	Location     string `json:"location"`
	PolicyData   string `json:"policy_data"`
	Project      string `json:"project"`
}