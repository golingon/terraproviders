// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mediajob "github.com/golingon/terraproviders/azurerm/3.52.0/mediajob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMediaJob creates a new instance of [MediaJob].
func NewMediaJob(name string, args MediaJobArgs) *MediaJob {
	return &MediaJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaJob)(nil)

// MediaJob represents the Terraform resource azurerm_media_job.
type MediaJob struct {
	Name      string
	Args      MediaJobArgs
	state     *mediaJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MediaJob].
func (mj *MediaJob) Type() string {
	return "azurerm_media_job"
}

// LocalName returns the local name for [MediaJob].
func (mj *MediaJob) LocalName() string {
	return mj.Name
}

// Configuration returns the configuration (args) for [MediaJob].
func (mj *MediaJob) Configuration() interface{} {
	return mj.Args
}

// DependOn is used for other resources to depend on [MediaJob].
func (mj *MediaJob) DependOn() terra.Reference {
	return terra.ReferenceResource(mj)
}

// Dependencies returns the list of resources [MediaJob] depends_on.
func (mj *MediaJob) Dependencies() terra.Dependencies {
	return mj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MediaJob].
func (mj *MediaJob) LifecycleManagement() *terra.Lifecycle {
	return mj.Lifecycle
}

// Attributes returns the attributes for [MediaJob].
func (mj *MediaJob) Attributes() mediaJobAttributes {
	return mediaJobAttributes{ref: terra.ReferenceResource(mj)}
}

// ImportState imports the given attribute values into [MediaJob]'s state.
func (mj *MediaJob) ImportState(av io.Reader) error {
	mj.state = &mediaJobState{}
	if err := json.NewDecoder(av).Decode(mj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mj.Type(), mj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MediaJob] has state.
func (mj *MediaJob) State() (*mediaJobState, bool) {
	return mj.state, mj.state != nil
}

// StateMust returns the state for [MediaJob]. Panics if the state is nil.
func (mj *MediaJob) StateMust() *mediaJobState {
	if mj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mj.Type(), mj.LocalName()))
	}
	return mj.state
}

// MediaJobArgs contains the configurations for azurerm_media_job.
type MediaJobArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MediaServicesAccountName: string, required
	MediaServicesAccountName terra.StringValue `hcl:"media_services_account_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: string, optional
	Priority terra.StringValue `hcl:"priority,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// TransformName: string, required
	TransformName terra.StringValue `hcl:"transform_name,attr" validate:"required"`
	// InputAsset: required
	InputAsset *mediajob.InputAsset `hcl:"input_asset,block" validate:"required"`
	// OutputAsset: min=1
	OutputAsset []mediajob.OutputAsset `hcl:"output_asset,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *mediajob.Timeouts `hcl:"timeouts,block"`
}
type mediaJobAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_media_job.
func (mj mediaJobAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mj.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_media_job.
func (mj mediaJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mj.ref.Append("id"))
}

// MediaServicesAccountName returns a reference to field media_services_account_name of azurerm_media_job.
func (mj mediaJobAttributes) MediaServicesAccountName() terra.StringValue {
	return terra.ReferenceAsString(mj.ref.Append("media_services_account_name"))
}

// Name returns a reference to field name of azurerm_media_job.
func (mj mediaJobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mj.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_media_job.
func (mj mediaJobAttributes) Priority() terra.StringValue {
	return terra.ReferenceAsString(mj.ref.Append("priority"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_media_job.
func (mj mediaJobAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mj.ref.Append("resource_group_name"))
}

// TransformName returns a reference to field transform_name of azurerm_media_job.
func (mj mediaJobAttributes) TransformName() terra.StringValue {
	return terra.ReferenceAsString(mj.ref.Append("transform_name"))
}

func (mj mediaJobAttributes) InputAsset() terra.ListValue[mediajob.InputAssetAttributes] {
	return terra.ReferenceAsList[mediajob.InputAssetAttributes](mj.ref.Append("input_asset"))
}

func (mj mediaJobAttributes) OutputAsset() terra.ListValue[mediajob.OutputAssetAttributes] {
	return terra.ReferenceAsList[mediajob.OutputAssetAttributes](mj.ref.Append("output_asset"))
}

func (mj mediaJobAttributes) Timeouts() mediajob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mediajob.TimeoutsAttributes](mj.ref.Append("timeouts"))
}

type mediaJobState struct {
	Description              string                      `json:"description"`
	Id                       string                      `json:"id"`
	MediaServicesAccountName string                      `json:"media_services_account_name"`
	Name                     string                      `json:"name"`
	Priority                 string                      `json:"priority"`
	ResourceGroupName        string                      `json:"resource_group_name"`
	TransformName            string                      `json:"transform_name"`
	InputAsset               []mediajob.InputAssetState  `json:"input_asset"`
	OutputAsset              []mediajob.OutputAssetState `json:"output_asset"`
	Timeouts                 *mediajob.TimeoutsState     `json:"timeouts"`
}
