// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	storagebucket "github.com/golingon/terraproviders/googlebeta/4.77.0/storagebucket"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageBucket creates a new instance of [StorageBucket].
func NewStorageBucket(name string, args StorageBucketArgs) *StorageBucket {
	return &StorageBucket{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageBucket)(nil)

// StorageBucket represents the Terraform resource google_storage_bucket.
type StorageBucket struct {
	Name      string
	Args      StorageBucketArgs
	state     *storageBucketState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageBucket].
func (sb *StorageBucket) Type() string {
	return "google_storage_bucket"
}

// LocalName returns the local name for [StorageBucket].
func (sb *StorageBucket) LocalName() string {
	return sb.Name
}

// Configuration returns the configuration (args) for [StorageBucket].
func (sb *StorageBucket) Configuration() interface{} {
	return sb.Args
}

// DependOn is used for other resources to depend on [StorageBucket].
func (sb *StorageBucket) DependOn() terra.Reference {
	return terra.ReferenceResource(sb)
}

// Dependencies returns the list of resources [StorageBucket] depends_on.
func (sb *StorageBucket) Dependencies() terra.Dependencies {
	return sb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageBucket].
func (sb *StorageBucket) LifecycleManagement() *terra.Lifecycle {
	return sb.Lifecycle
}

// Attributes returns the attributes for [StorageBucket].
func (sb *StorageBucket) Attributes() storageBucketAttributes {
	return storageBucketAttributes{ref: terra.ReferenceResource(sb)}
}

// ImportState imports the given attribute values into [StorageBucket]'s state.
func (sb *StorageBucket) ImportState(av io.Reader) error {
	sb.state = &storageBucketState{}
	if err := json.NewDecoder(av).Decode(sb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sb.Type(), sb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageBucket] has state.
func (sb *StorageBucket) State() (*storageBucketState, bool) {
	return sb.state, sb.state != nil
}

// StateMust returns the state for [StorageBucket]. Panics if the state is nil.
func (sb *StorageBucket) StateMust() *storageBucketState {
	if sb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sb.Type(), sb.LocalName()))
	}
	return sb.state
}

// StorageBucketArgs contains the configurations for google_storage_bucket.
type StorageBucketArgs struct {
	// DefaultEventBasedHold: bool, optional
	DefaultEventBasedHold terra.BoolValue `hcl:"default_event_based_hold,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// PublicAccessPrevention: string, optional
	PublicAccessPrevention terra.StringValue `hcl:"public_access_prevention,attr"`
	// RequesterPays: bool, optional
	RequesterPays terra.BoolValue `hcl:"requester_pays,attr"`
	// StorageClass: string, optional
	StorageClass terra.StringValue `hcl:"storage_class,attr"`
	// UniformBucketLevelAccess: bool, optional
	UniformBucketLevelAccess terra.BoolValue `hcl:"uniform_bucket_level_access,attr"`
	// Autoclass: optional
	Autoclass *storagebucket.Autoclass `hcl:"autoclass,block"`
	// Cors: min=0
	Cors []storagebucket.Cors `hcl:"cors,block" validate:"min=0"`
	// CustomPlacementConfig: optional
	CustomPlacementConfig *storagebucket.CustomPlacementConfig `hcl:"custom_placement_config,block"`
	// Encryption: optional
	Encryption *storagebucket.Encryption `hcl:"encryption,block"`
	// LifecycleRule: min=0,max=100
	LifecycleRule []storagebucket.LifecycleRule `hcl:"lifecycle_rule,block" validate:"min=0,max=100"`
	// Logging: optional
	Logging *storagebucket.Logging `hcl:"logging,block"`
	// RetentionPolicy: optional
	RetentionPolicy *storagebucket.RetentionPolicy `hcl:"retention_policy,block"`
	// Timeouts: optional
	Timeouts *storagebucket.Timeouts `hcl:"timeouts,block"`
	// Versioning: optional
	Versioning *storagebucket.Versioning `hcl:"versioning,block"`
	// Website: optional
	Website *storagebucket.Website `hcl:"website,block"`
}
type storageBucketAttributes struct {
	ref terra.Reference
}

// DefaultEventBasedHold returns a reference to field default_event_based_hold of google_storage_bucket.
func (sb storageBucketAttributes) DefaultEventBasedHold() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("default_event_based_hold"))
}

// ForceDestroy returns a reference to field force_destroy of google_storage_bucket.
func (sb storageBucketAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("force_destroy"))
}

// Id returns a reference to field id of google_storage_bucket.
func (sb storageBucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("id"))
}

// Labels returns a reference to field labels of google_storage_bucket.
func (sb storageBucketAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sb.ref.Append("labels"))
}

// Location returns a reference to field location of google_storage_bucket.
func (sb storageBucketAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("location"))
}

// Name returns a reference to field name of google_storage_bucket.
func (sb storageBucketAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("name"))
}

// Project returns a reference to field project of google_storage_bucket.
func (sb storageBucketAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("project"))
}

// PublicAccessPrevention returns a reference to field public_access_prevention of google_storage_bucket.
func (sb storageBucketAttributes) PublicAccessPrevention() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("public_access_prevention"))
}

// RequesterPays returns a reference to field requester_pays of google_storage_bucket.
func (sb storageBucketAttributes) RequesterPays() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("requester_pays"))
}

// SelfLink returns a reference to field self_link of google_storage_bucket.
func (sb storageBucketAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("self_link"))
}

// StorageClass returns a reference to field storage_class of google_storage_bucket.
func (sb storageBucketAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("storage_class"))
}

// UniformBucketLevelAccess returns a reference to field uniform_bucket_level_access of google_storage_bucket.
func (sb storageBucketAttributes) UniformBucketLevelAccess() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("uniform_bucket_level_access"))
}

// Url returns a reference to field url of google_storage_bucket.
func (sb storageBucketAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("url"))
}

func (sb storageBucketAttributes) Autoclass() terra.ListValue[storagebucket.AutoclassAttributes] {
	return terra.ReferenceAsList[storagebucket.AutoclassAttributes](sb.ref.Append("autoclass"))
}

func (sb storageBucketAttributes) Cors() terra.ListValue[storagebucket.CorsAttributes] {
	return terra.ReferenceAsList[storagebucket.CorsAttributes](sb.ref.Append("cors"))
}

func (sb storageBucketAttributes) CustomPlacementConfig() terra.ListValue[storagebucket.CustomPlacementConfigAttributes] {
	return terra.ReferenceAsList[storagebucket.CustomPlacementConfigAttributes](sb.ref.Append("custom_placement_config"))
}

func (sb storageBucketAttributes) Encryption() terra.ListValue[storagebucket.EncryptionAttributes] {
	return terra.ReferenceAsList[storagebucket.EncryptionAttributes](sb.ref.Append("encryption"))
}

func (sb storageBucketAttributes) LifecycleRule() terra.ListValue[storagebucket.LifecycleRuleAttributes] {
	return terra.ReferenceAsList[storagebucket.LifecycleRuleAttributes](sb.ref.Append("lifecycle_rule"))
}

func (sb storageBucketAttributes) Logging() terra.ListValue[storagebucket.LoggingAttributes] {
	return terra.ReferenceAsList[storagebucket.LoggingAttributes](sb.ref.Append("logging"))
}

func (sb storageBucketAttributes) RetentionPolicy() terra.ListValue[storagebucket.RetentionPolicyAttributes] {
	return terra.ReferenceAsList[storagebucket.RetentionPolicyAttributes](sb.ref.Append("retention_policy"))
}

func (sb storageBucketAttributes) Timeouts() storagebucket.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagebucket.TimeoutsAttributes](sb.ref.Append("timeouts"))
}

func (sb storageBucketAttributes) Versioning() terra.ListValue[storagebucket.VersioningAttributes] {
	return terra.ReferenceAsList[storagebucket.VersioningAttributes](sb.ref.Append("versioning"))
}

func (sb storageBucketAttributes) Website() terra.ListValue[storagebucket.WebsiteAttributes] {
	return terra.ReferenceAsList[storagebucket.WebsiteAttributes](sb.ref.Append("website"))
}

type storageBucketState struct {
	DefaultEventBasedHold    bool                                       `json:"default_event_based_hold"`
	ForceDestroy             bool                                       `json:"force_destroy"`
	Id                       string                                     `json:"id"`
	Labels                   map[string]string                          `json:"labels"`
	Location                 string                                     `json:"location"`
	Name                     string                                     `json:"name"`
	Project                  string                                     `json:"project"`
	PublicAccessPrevention   string                                     `json:"public_access_prevention"`
	RequesterPays            bool                                       `json:"requester_pays"`
	SelfLink                 string                                     `json:"self_link"`
	StorageClass             string                                     `json:"storage_class"`
	UniformBucketLevelAccess bool                                       `json:"uniform_bucket_level_access"`
	Url                      string                                     `json:"url"`
	Autoclass                []storagebucket.AutoclassState             `json:"autoclass"`
	Cors                     []storagebucket.CorsState                  `json:"cors"`
	CustomPlacementConfig    []storagebucket.CustomPlacementConfigState `json:"custom_placement_config"`
	Encryption               []storagebucket.EncryptionState            `json:"encryption"`
	LifecycleRule            []storagebucket.LifecycleRuleState         `json:"lifecycle_rule"`
	Logging                  []storagebucket.LoggingState               `json:"logging"`
	RetentionPolicy          []storagebucket.RetentionPolicyState       `json:"retention_policy"`
	Timeouts                 *storagebucket.TimeoutsState               `json:"timeouts"`
	Versioning               []storagebucket.VersioningState            `json:"versioning"`
	Website                  []storagebucket.WebsiteState               `json:"website"`
}
