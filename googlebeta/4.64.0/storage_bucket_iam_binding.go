// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	storagebucketiambinding "github.com/golingon/terraproviders/googlebeta/4.64.0/storagebucketiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageBucketIamBinding creates a new instance of [StorageBucketIamBinding].
func NewStorageBucketIamBinding(name string, args StorageBucketIamBindingArgs) *StorageBucketIamBinding {
	return &StorageBucketIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageBucketIamBinding)(nil)

// StorageBucketIamBinding represents the Terraform resource google_storage_bucket_iam_binding.
type StorageBucketIamBinding struct {
	Name      string
	Args      StorageBucketIamBindingArgs
	state     *storageBucketIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageBucketIamBinding].
func (sbib *StorageBucketIamBinding) Type() string {
	return "google_storage_bucket_iam_binding"
}

// LocalName returns the local name for [StorageBucketIamBinding].
func (sbib *StorageBucketIamBinding) LocalName() string {
	return sbib.Name
}

// Configuration returns the configuration (args) for [StorageBucketIamBinding].
func (sbib *StorageBucketIamBinding) Configuration() interface{} {
	return sbib.Args
}

// DependOn is used for other resources to depend on [StorageBucketIamBinding].
func (sbib *StorageBucketIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(sbib)
}

// Dependencies returns the list of resources [StorageBucketIamBinding] depends_on.
func (sbib *StorageBucketIamBinding) Dependencies() terra.Dependencies {
	return sbib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageBucketIamBinding].
func (sbib *StorageBucketIamBinding) LifecycleManagement() *terra.Lifecycle {
	return sbib.Lifecycle
}

// Attributes returns the attributes for [StorageBucketIamBinding].
func (sbib *StorageBucketIamBinding) Attributes() storageBucketIamBindingAttributes {
	return storageBucketIamBindingAttributes{ref: terra.ReferenceResource(sbib)}
}

// ImportState imports the given attribute values into [StorageBucketIamBinding]'s state.
func (sbib *StorageBucketIamBinding) ImportState(av io.Reader) error {
	sbib.state = &storageBucketIamBindingState{}
	if err := json.NewDecoder(av).Decode(sbib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbib.Type(), sbib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageBucketIamBinding] has state.
func (sbib *StorageBucketIamBinding) State() (*storageBucketIamBindingState, bool) {
	return sbib.state, sbib.state != nil
}

// StateMust returns the state for [StorageBucketIamBinding]. Panics if the state is nil.
func (sbib *StorageBucketIamBinding) StateMust() *storageBucketIamBindingState {
	if sbib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbib.Type(), sbib.LocalName()))
	}
	return sbib.state
}

// StorageBucketIamBindingArgs contains the configurations for google_storage_bucket_iam_binding.
type StorageBucketIamBindingArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *storagebucketiambinding.Condition `hcl:"condition,block"`
}
type storageBucketIamBindingAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_bucket_iam_binding.
func (sbib storageBucketIamBindingAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbib.ref.Append("bucket"))
}

// Etag returns a reference to field etag of google_storage_bucket_iam_binding.
func (sbib storageBucketIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sbib.ref.Append("etag"))
}

// Id returns a reference to field id of google_storage_bucket_iam_binding.
func (sbib storageBucketIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbib.ref.Append("id"))
}

// Members returns a reference to field members of google_storage_bucket_iam_binding.
func (sbib storageBucketIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sbib.ref.Append("members"))
}

// Role returns a reference to field role of google_storage_bucket_iam_binding.
func (sbib storageBucketIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(sbib.ref.Append("role"))
}

func (sbib storageBucketIamBindingAttributes) Condition() terra.ListValue[storagebucketiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[storagebucketiambinding.ConditionAttributes](sbib.ref.Append("condition"))
}

type storageBucketIamBindingState struct {
	Bucket    string                                   `json:"bucket"`
	Etag      string                                   `json:"etag"`
	Id        string                                   `json:"id"`
	Members   []string                                 `json:"members"`
	Role      string                                   `json:"role"`
	Condition []storagebucketiambinding.ConditionState `json:"condition"`
}
