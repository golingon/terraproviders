// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	directoryserviceregion "github.com/golingon/terraproviders/aws/4.63.0/directoryserviceregion"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDirectoryServiceRegion creates a new instance of [DirectoryServiceRegion].
func NewDirectoryServiceRegion(name string, args DirectoryServiceRegionArgs) *DirectoryServiceRegion {
	return &DirectoryServiceRegion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DirectoryServiceRegion)(nil)

// DirectoryServiceRegion represents the Terraform resource aws_directory_service_region.
type DirectoryServiceRegion struct {
	Name      string
	Args      DirectoryServiceRegionArgs
	state     *directoryServiceRegionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DirectoryServiceRegion].
func (dsr *DirectoryServiceRegion) Type() string {
	return "aws_directory_service_region"
}

// LocalName returns the local name for [DirectoryServiceRegion].
func (dsr *DirectoryServiceRegion) LocalName() string {
	return dsr.Name
}

// Configuration returns the configuration (args) for [DirectoryServiceRegion].
func (dsr *DirectoryServiceRegion) Configuration() interface{} {
	return dsr.Args
}

// DependOn is used for other resources to depend on [DirectoryServiceRegion].
func (dsr *DirectoryServiceRegion) DependOn() terra.Reference {
	return terra.ReferenceResource(dsr)
}

// Dependencies returns the list of resources [DirectoryServiceRegion] depends_on.
func (dsr *DirectoryServiceRegion) Dependencies() terra.Dependencies {
	return dsr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DirectoryServiceRegion].
func (dsr *DirectoryServiceRegion) LifecycleManagement() *terra.Lifecycle {
	return dsr.Lifecycle
}

// Attributes returns the attributes for [DirectoryServiceRegion].
func (dsr *DirectoryServiceRegion) Attributes() directoryServiceRegionAttributes {
	return directoryServiceRegionAttributes{ref: terra.ReferenceResource(dsr)}
}

// ImportState imports the given attribute values into [DirectoryServiceRegion]'s state.
func (dsr *DirectoryServiceRegion) ImportState(av io.Reader) error {
	dsr.state = &directoryServiceRegionState{}
	if err := json.NewDecoder(av).Decode(dsr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsr.Type(), dsr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DirectoryServiceRegion] has state.
func (dsr *DirectoryServiceRegion) State() (*directoryServiceRegionState, bool) {
	return dsr.state, dsr.state != nil
}

// StateMust returns the state for [DirectoryServiceRegion]. Panics if the state is nil.
func (dsr *DirectoryServiceRegion) StateMust() *directoryServiceRegionState {
	if dsr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsr.Type(), dsr.LocalName()))
	}
	return dsr.state
}

// DirectoryServiceRegionArgs contains the configurations for aws_directory_service_region.
type DirectoryServiceRegionArgs struct {
	// DesiredNumberOfDomainControllers: number, optional
	DesiredNumberOfDomainControllers terra.NumberValue `hcl:"desired_number_of_domain_controllers,attr"`
	// DirectoryId: string, required
	DirectoryId terra.StringValue `hcl:"directory_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RegionName: string, required
	RegionName terra.StringValue `hcl:"region_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *directoryserviceregion.Timeouts `hcl:"timeouts,block"`
	// VpcSettings: required
	VpcSettings *directoryserviceregion.VpcSettings `hcl:"vpc_settings,block" validate:"required"`
}
type directoryServiceRegionAttributes struct {
	ref terra.Reference
}

// DesiredNumberOfDomainControllers returns a reference to field desired_number_of_domain_controllers of aws_directory_service_region.
func (dsr directoryServiceRegionAttributes) DesiredNumberOfDomainControllers() terra.NumberValue {
	return terra.ReferenceAsNumber(dsr.ref.Append("desired_number_of_domain_controllers"))
}

// DirectoryId returns a reference to field directory_id of aws_directory_service_region.
func (dsr directoryServiceRegionAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("directory_id"))
}

// Id returns a reference to field id of aws_directory_service_region.
func (dsr directoryServiceRegionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("id"))
}

// RegionName returns a reference to field region_name of aws_directory_service_region.
func (dsr directoryServiceRegionAttributes) RegionName() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("region_name"))
}

// Tags returns a reference to field tags of aws_directory_service_region.
func (dsr directoryServiceRegionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_directory_service_region.
func (dsr directoryServiceRegionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsr.ref.Append("tags_all"))
}

func (dsr directoryServiceRegionAttributes) Timeouts() directoryserviceregion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[directoryserviceregion.TimeoutsAttributes](dsr.ref.Append("timeouts"))
}

func (dsr directoryServiceRegionAttributes) VpcSettings() terra.ListValue[directoryserviceregion.VpcSettingsAttributes] {
	return terra.ReferenceAsList[directoryserviceregion.VpcSettingsAttributes](dsr.ref.Append("vpc_settings"))
}

type directoryServiceRegionState struct {
	DesiredNumberOfDomainControllers float64                                   `json:"desired_number_of_domain_controllers"`
	DirectoryId                      string                                    `json:"directory_id"`
	Id                               string                                    `json:"id"`
	RegionName                       string                                    `json:"region_name"`
	Tags                             map[string]string                         `json:"tags"`
	TagsAll                          map[string]string                         `json:"tags_all"`
	Timeouts                         *directoryserviceregion.TimeoutsState     `json:"timeouts"`
	VpcSettings                      []directoryserviceregion.VpcSettingsState `json:"vpc_settings"`
}
