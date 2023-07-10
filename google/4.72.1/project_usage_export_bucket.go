// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	projectusageexportbucket "github.com/golingon/terraproviders/google/4.72.1/projectusageexportbucket"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProjectUsageExportBucket creates a new instance of [ProjectUsageExportBucket].
func NewProjectUsageExportBucket(name string, args ProjectUsageExportBucketArgs) *ProjectUsageExportBucket {
	return &ProjectUsageExportBucket{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProjectUsageExportBucket)(nil)

// ProjectUsageExportBucket represents the Terraform resource google_project_usage_export_bucket.
type ProjectUsageExportBucket struct {
	Name      string
	Args      ProjectUsageExportBucketArgs
	state     *projectUsageExportBucketState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProjectUsageExportBucket].
func (pueb *ProjectUsageExportBucket) Type() string {
	return "google_project_usage_export_bucket"
}

// LocalName returns the local name for [ProjectUsageExportBucket].
func (pueb *ProjectUsageExportBucket) LocalName() string {
	return pueb.Name
}

// Configuration returns the configuration (args) for [ProjectUsageExportBucket].
func (pueb *ProjectUsageExportBucket) Configuration() interface{} {
	return pueb.Args
}

// DependOn is used for other resources to depend on [ProjectUsageExportBucket].
func (pueb *ProjectUsageExportBucket) DependOn() terra.Reference {
	return terra.ReferenceResource(pueb)
}

// Dependencies returns the list of resources [ProjectUsageExportBucket] depends_on.
func (pueb *ProjectUsageExportBucket) Dependencies() terra.Dependencies {
	return pueb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProjectUsageExportBucket].
func (pueb *ProjectUsageExportBucket) LifecycleManagement() *terra.Lifecycle {
	return pueb.Lifecycle
}

// Attributes returns the attributes for [ProjectUsageExportBucket].
func (pueb *ProjectUsageExportBucket) Attributes() projectUsageExportBucketAttributes {
	return projectUsageExportBucketAttributes{ref: terra.ReferenceResource(pueb)}
}

// ImportState imports the given attribute values into [ProjectUsageExportBucket]'s state.
func (pueb *ProjectUsageExportBucket) ImportState(av io.Reader) error {
	pueb.state = &projectUsageExportBucketState{}
	if err := json.NewDecoder(av).Decode(pueb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pueb.Type(), pueb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProjectUsageExportBucket] has state.
func (pueb *ProjectUsageExportBucket) State() (*projectUsageExportBucketState, bool) {
	return pueb.state, pueb.state != nil
}

// StateMust returns the state for [ProjectUsageExportBucket]. Panics if the state is nil.
func (pueb *ProjectUsageExportBucket) StateMust() *projectUsageExportBucketState {
	if pueb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pueb.Type(), pueb.LocalName()))
	}
	return pueb.state
}

// ProjectUsageExportBucketArgs contains the configurations for google_project_usage_export_bucket.
type ProjectUsageExportBucketArgs struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *projectusageexportbucket.Timeouts `hcl:"timeouts,block"`
}
type projectUsageExportBucketAttributes struct {
	ref terra.Reference
}

// BucketName returns a reference to field bucket_name of google_project_usage_export_bucket.
func (pueb projectUsageExportBucketAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(pueb.ref.Append("bucket_name"))
}

// Id returns a reference to field id of google_project_usage_export_bucket.
func (pueb projectUsageExportBucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pueb.ref.Append("id"))
}

// Prefix returns a reference to field prefix of google_project_usage_export_bucket.
func (pueb projectUsageExportBucketAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(pueb.ref.Append("prefix"))
}

// Project returns a reference to field project of google_project_usage_export_bucket.
func (pueb projectUsageExportBucketAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pueb.ref.Append("project"))
}

func (pueb projectUsageExportBucketAttributes) Timeouts() projectusageexportbucket.TimeoutsAttributes {
	return terra.ReferenceAsSingle[projectusageexportbucket.TimeoutsAttributes](pueb.ref.Append("timeouts"))
}

type projectUsageExportBucketState struct {
	BucketName string                                  `json:"bucket_name"`
	Id         string                                  `json:"id"`
	Prefix     string                                  `json:"prefix"`
	Project    string                                  `json:"project"`
	Timeouts   *projectusageexportbucket.TimeoutsState `json:"timeouts"`
}
