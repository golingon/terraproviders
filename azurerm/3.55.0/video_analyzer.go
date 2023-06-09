// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	videoanalyzer "github.com/golingon/terraproviders/azurerm/3.55.0/videoanalyzer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVideoAnalyzer creates a new instance of [VideoAnalyzer].
func NewVideoAnalyzer(name string, args VideoAnalyzerArgs) *VideoAnalyzer {
	return &VideoAnalyzer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VideoAnalyzer)(nil)

// VideoAnalyzer represents the Terraform resource azurerm_video_analyzer.
type VideoAnalyzer struct {
	Name      string
	Args      VideoAnalyzerArgs
	state     *videoAnalyzerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VideoAnalyzer].
func (va *VideoAnalyzer) Type() string {
	return "azurerm_video_analyzer"
}

// LocalName returns the local name for [VideoAnalyzer].
func (va *VideoAnalyzer) LocalName() string {
	return va.Name
}

// Configuration returns the configuration (args) for [VideoAnalyzer].
func (va *VideoAnalyzer) Configuration() interface{} {
	return va.Args
}

// DependOn is used for other resources to depend on [VideoAnalyzer].
func (va *VideoAnalyzer) DependOn() terra.Reference {
	return terra.ReferenceResource(va)
}

// Dependencies returns the list of resources [VideoAnalyzer] depends_on.
func (va *VideoAnalyzer) Dependencies() terra.Dependencies {
	return va.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VideoAnalyzer].
func (va *VideoAnalyzer) LifecycleManagement() *terra.Lifecycle {
	return va.Lifecycle
}

// Attributes returns the attributes for [VideoAnalyzer].
func (va *VideoAnalyzer) Attributes() videoAnalyzerAttributes {
	return videoAnalyzerAttributes{ref: terra.ReferenceResource(va)}
}

// ImportState imports the given attribute values into [VideoAnalyzer]'s state.
func (va *VideoAnalyzer) ImportState(av io.Reader) error {
	va.state = &videoAnalyzerState{}
	if err := json.NewDecoder(av).Decode(va.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", va.Type(), va.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VideoAnalyzer] has state.
func (va *VideoAnalyzer) State() (*videoAnalyzerState, bool) {
	return va.state, va.state != nil
}

// StateMust returns the state for [VideoAnalyzer]. Panics if the state is nil.
func (va *VideoAnalyzer) StateMust() *videoAnalyzerState {
	if va.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", va.Type(), va.LocalName()))
	}
	return va.state
}

// VideoAnalyzerArgs contains the configurations for azurerm_video_analyzer.
type VideoAnalyzerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: required
	Identity *videoanalyzer.Identity `hcl:"identity,block" validate:"required"`
	// StorageAccount: required
	StorageAccount *videoanalyzer.StorageAccount `hcl:"storage_account,block" validate:"required"`
	// Timeouts: optional
	Timeouts *videoanalyzer.Timeouts `hcl:"timeouts,block"`
}
type videoAnalyzerAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_video_analyzer.
func (va videoAnalyzerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(va.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_video_analyzer.
func (va videoAnalyzerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(va.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_video_analyzer.
func (va videoAnalyzerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(va.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_video_analyzer.
func (va videoAnalyzerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(va.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_video_analyzer.
func (va videoAnalyzerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](va.ref.Append("tags"))
}

func (va videoAnalyzerAttributes) Identity() terra.ListValue[videoanalyzer.IdentityAttributes] {
	return terra.ReferenceAsList[videoanalyzer.IdentityAttributes](va.ref.Append("identity"))
}

func (va videoAnalyzerAttributes) StorageAccount() terra.ListValue[videoanalyzer.StorageAccountAttributes] {
	return terra.ReferenceAsList[videoanalyzer.StorageAccountAttributes](va.ref.Append("storage_account"))
}

func (va videoAnalyzerAttributes) Timeouts() videoanalyzer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[videoanalyzer.TimeoutsAttributes](va.ref.Append("timeouts"))
}

type videoAnalyzerState struct {
	Id                string                              `json:"id"`
	Location          string                              `json:"location"`
	Name              string                              `json:"name"`
	ResourceGroupName string                              `json:"resource_group_name"`
	Tags              map[string]string                   `json:"tags"`
	Identity          []videoanalyzer.IdentityState       `json:"identity"`
	StorageAccount    []videoanalyzer.StorageAccountState `json:"storage_account"`
	Timeouts          *videoanalyzer.TimeoutsState        `json:"timeouts"`
}
