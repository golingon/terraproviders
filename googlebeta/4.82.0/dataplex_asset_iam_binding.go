// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataplexassetiambinding "github.com/golingon/terraproviders/googlebeta/4.82.0/dataplexassetiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexAssetIamBinding creates a new instance of [DataplexAssetIamBinding].
func NewDataplexAssetIamBinding(name string, args DataplexAssetIamBindingArgs) *DataplexAssetIamBinding {
	return &DataplexAssetIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexAssetIamBinding)(nil)

// DataplexAssetIamBinding represents the Terraform resource google_dataplex_asset_iam_binding.
type DataplexAssetIamBinding struct {
	Name      string
	Args      DataplexAssetIamBindingArgs
	state     *dataplexAssetIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexAssetIamBinding].
func (daib *DataplexAssetIamBinding) Type() string {
	return "google_dataplex_asset_iam_binding"
}

// LocalName returns the local name for [DataplexAssetIamBinding].
func (daib *DataplexAssetIamBinding) LocalName() string {
	return daib.Name
}

// Configuration returns the configuration (args) for [DataplexAssetIamBinding].
func (daib *DataplexAssetIamBinding) Configuration() interface{} {
	return daib.Args
}

// DependOn is used for other resources to depend on [DataplexAssetIamBinding].
func (daib *DataplexAssetIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(daib)
}

// Dependencies returns the list of resources [DataplexAssetIamBinding] depends_on.
func (daib *DataplexAssetIamBinding) Dependencies() terra.Dependencies {
	return daib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexAssetIamBinding].
func (daib *DataplexAssetIamBinding) LifecycleManagement() *terra.Lifecycle {
	return daib.Lifecycle
}

// Attributes returns the attributes for [DataplexAssetIamBinding].
func (daib *DataplexAssetIamBinding) Attributes() dataplexAssetIamBindingAttributes {
	return dataplexAssetIamBindingAttributes{ref: terra.ReferenceResource(daib)}
}

// ImportState imports the given attribute values into [DataplexAssetIamBinding]'s state.
func (daib *DataplexAssetIamBinding) ImportState(av io.Reader) error {
	daib.state = &dataplexAssetIamBindingState{}
	if err := json.NewDecoder(av).Decode(daib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", daib.Type(), daib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexAssetIamBinding] has state.
func (daib *DataplexAssetIamBinding) State() (*dataplexAssetIamBindingState, bool) {
	return daib.state, daib.state != nil
}

// StateMust returns the state for [DataplexAssetIamBinding]. Panics if the state is nil.
func (daib *DataplexAssetIamBinding) StateMust() *dataplexAssetIamBindingState {
	if daib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", daib.Type(), daib.LocalName()))
	}
	return daib.state
}

// DataplexAssetIamBindingArgs contains the configurations for google_dataplex_asset_iam_binding.
type DataplexAssetIamBindingArgs struct {
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
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *dataplexassetiambinding.Condition `hcl:"condition,block"`
}
type dataplexAssetIamBindingAttributes struct {
	ref terra.Reference
}

// Asset returns a reference to field asset of google_dataplex_asset_iam_binding.
func (daib dataplexAssetIamBindingAttributes) Asset() terra.StringValue {
	return terra.ReferenceAsString(daib.ref.Append("asset"))
}

// DataplexZone returns a reference to field dataplex_zone of google_dataplex_asset_iam_binding.
func (daib dataplexAssetIamBindingAttributes) DataplexZone() terra.StringValue {
	return terra.ReferenceAsString(daib.ref.Append("dataplex_zone"))
}

// Etag returns a reference to field etag of google_dataplex_asset_iam_binding.
func (daib dataplexAssetIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(daib.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_asset_iam_binding.
func (daib dataplexAssetIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(daib.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_asset_iam_binding.
func (daib dataplexAssetIamBindingAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(daib.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_asset_iam_binding.
func (daib dataplexAssetIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(daib.ref.Append("location"))
}

// Members returns a reference to field members of google_dataplex_asset_iam_binding.
func (daib dataplexAssetIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](daib.ref.Append("members"))
}

// Project returns a reference to field project of google_dataplex_asset_iam_binding.
func (daib dataplexAssetIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(daib.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_asset_iam_binding.
func (daib dataplexAssetIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(daib.ref.Append("role"))
}

func (daib dataplexAssetIamBindingAttributes) Condition() terra.ListValue[dataplexassetiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[dataplexassetiambinding.ConditionAttributes](daib.ref.Append("condition"))
}

type dataplexAssetIamBindingState struct {
	Asset        string                                   `json:"asset"`
	DataplexZone string                                   `json:"dataplex_zone"`
	Etag         string                                   `json:"etag"`
	Id           string                                   `json:"id"`
	Lake         string                                   `json:"lake"`
	Location     string                                   `json:"location"`
	Members      []string                                 `json:"members"`
	Project      string                                   `json:"project"`
	Role         string                                   `json:"role"`
	Condition    []dataplexassetiambinding.ConditionState `json:"condition"`
}
