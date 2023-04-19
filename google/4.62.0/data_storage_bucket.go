// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datastoragebucket "github.com/golingon/terraproviders/google/4.62.0/datastoragebucket"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataStorageBucket creates a new instance of [DataStorageBucket].
func NewDataStorageBucket(name string, args DataStorageBucketArgs) *DataStorageBucket {
	return &DataStorageBucket{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageBucket)(nil)

// DataStorageBucket represents the Terraform data resource google_storage_bucket.
type DataStorageBucket struct {
	Name string
	Args DataStorageBucketArgs
}

// DataSource returns the Terraform object type for [DataStorageBucket].
func (sb *DataStorageBucket) DataSource() string {
	return "google_storage_bucket"
}

// LocalName returns the local name for [DataStorageBucket].
func (sb *DataStorageBucket) LocalName() string {
	return sb.Name
}

// Configuration returns the configuration (args) for [DataStorageBucket].
func (sb *DataStorageBucket) Configuration() interface{} {
	return sb.Args
}

// Attributes returns the attributes for [DataStorageBucket].
func (sb *DataStorageBucket) Attributes() dataStorageBucketAttributes {
	return dataStorageBucketAttributes{ref: terra.ReferenceDataResource(sb)}
}

// DataStorageBucketArgs contains the configurations for google_storage_bucket.
type DataStorageBucketArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Autoclass: min=0
	Autoclass []datastoragebucket.Autoclass `hcl:"autoclass,block" validate:"min=0"`
	// Cors: min=0
	Cors []datastoragebucket.Cors `hcl:"cors,block" validate:"min=0"`
	// CustomPlacementConfig: min=0
	CustomPlacementConfig []datastoragebucket.CustomPlacementConfig `hcl:"custom_placement_config,block" validate:"min=0"`
	// Encryption: min=0
	Encryption []datastoragebucket.Encryption `hcl:"encryption,block" validate:"min=0"`
	// LifecycleRule: min=0
	LifecycleRule []datastoragebucket.LifecycleRule `hcl:"lifecycle_rule,block" validate:"min=0"`
	// Logging: min=0
	Logging []datastoragebucket.Logging `hcl:"logging,block" validate:"min=0"`
	// RetentionPolicy: min=0
	RetentionPolicy []datastoragebucket.RetentionPolicy `hcl:"retention_policy,block" validate:"min=0"`
	// Versioning: min=0
	Versioning []datastoragebucket.Versioning `hcl:"versioning,block" validate:"min=0"`
	// Website: min=0
	Website []datastoragebucket.Website `hcl:"website,block" validate:"min=0"`
}
type dataStorageBucketAttributes struct {
	ref terra.Reference
}

// DefaultEventBasedHold returns a reference to field default_event_based_hold of google_storage_bucket.
func (sb dataStorageBucketAttributes) DefaultEventBasedHold() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("default_event_based_hold"))
}

// ForceDestroy returns a reference to field force_destroy of google_storage_bucket.
func (sb dataStorageBucketAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("force_destroy"))
}

// Id returns a reference to field id of google_storage_bucket.
func (sb dataStorageBucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("id"))
}

// Labels returns a reference to field labels of google_storage_bucket.
func (sb dataStorageBucketAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sb.ref.Append("labels"))
}

// Location returns a reference to field location of google_storage_bucket.
func (sb dataStorageBucketAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("location"))
}

// Name returns a reference to field name of google_storage_bucket.
func (sb dataStorageBucketAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("name"))
}

// Project returns a reference to field project of google_storage_bucket.
func (sb dataStorageBucketAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("project"))
}

// PublicAccessPrevention returns a reference to field public_access_prevention of google_storage_bucket.
func (sb dataStorageBucketAttributes) PublicAccessPrevention() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("public_access_prevention"))
}

// RequesterPays returns a reference to field requester_pays of google_storage_bucket.
func (sb dataStorageBucketAttributes) RequesterPays() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("requester_pays"))
}

// SelfLink returns a reference to field self_link of google_storage_bucket.
func (sb dataStorageBucketAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("self_link"))
}

// StorageClass returns a reference to field storage_class of google_storage_bucket.
func (sb dataStorageBucketAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("storage_class"))
}

// UniformBucketLevelAccess returns a reference to field uniform_bucket_level_access of google_storage_bucket.
func (sb dataStorageBucketAttributes) UniformBucketLevelAccess() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("uniform_bucket_level_access"))
}

// Url returns a reference to field url of google_storage_bucket.
func (sb dataStorageBucketAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("url"))
}

func (sb dataStorageBucketAttributes) Autoclass() terra.ListValue[datastoragebucket.AutoclassAttributes] {
	return terra.ReferenceAsList[datastoragebucket.AutoclassAttributes](sb.ref.Append("autoclass"))
}

func (sb dataStorageBucketAttributes) Cors() terra.ListValue[datastoragebucket.CorsAttributes] {
	return terra.ReferenceAsList[datastoragebucket.CorsAttributes](sb.ref.Append("cors"))
}

func (sb dataStorageBucketAttributes) CustomPlacementConfig() terra.ListValue[datastoragebucket.CustomPlacementConfigAttributes] {
	return terra.ReferenceAsList[datastoragebucket.CustomPlacementConfigAttributes](sb.ref.Append("custom_placement_config"))
}

func (sb dataStorageBucketAttributes) Encryption() terra.ListValue[datastoragebucket.EncryptionAttributes] {
	return terra.ReferenceAsList[datastoragebucket.EncryptionAttributes](sb.ref.Append("encryption"))
}

func (sb dataStorageBucketAttributes) LifecycleRule() terra.ListValue[datastoragebucket.LifecycleRuleAttributes] {
	return terra.ReferenceAsList[datastoragebucket.LifecycleRuleAttributes](sb.ref.Append("lifecycle_rule"))
}

func (sb dataStorageBucketAttributes) Logging() terra.ListValue[datastoragebucket.LoggingAttributes] {
	return terra.ReferenceAsList[datastoragebucket.LoggingAttributes](sb.ref.Append("logging"))
}

func (sb dataStorageBucketAttributes) RetentionPolicy() terra.ListValue[datastoragebucket.RetentionPolicyAttributes] {
	return terra.ReferenceAsList[datastoragebucket.RetentionPolicyAttributes](sb.ref.Append("retention_policy"))
}

func (sb dataStorageBucketAttributes) Versioning() terra.ListValue[datastoragebucket.VersioningAttributes] {
	return terra.ReferenceAsList[datastoragebucket.VersioningAttributes](sb.ref.Append("versioning"))
}

func (sb dataStorageBucketAttributes) Website() terra.ListValue[datastoragebucket.WebsiteAttributes] {
	return terra.ReferenceAsList[datastoragebucket.WebsiteAttributes](sb.ref.Append("website"))
}
