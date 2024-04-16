// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_directory_service_shared_directory

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_directory_service_shared_directory.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDirectoryServiceSharedDirectoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adssd *Resource) Type() string {
	return "aws_directory_service_shared_directory"
}

// LocalName returns the local name for [Resource].
func (adssd *Resource) LocalName() string {
	return adssd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adssd *Resource) Configuration() interface{} {
	return adssd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adssd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adssd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adssd *Resource) Dependencies() terra.Dependencies {
	return adssd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adssd *Resource) LifecycleManagement() *terra.Lifecycle {
	return adssd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adssd *Resource) Attributes() awsDirectoryServiceSharedDirectoryAttributes {
	return awsDirectoryServiceSharedDirectoryAttributes{ref: terra.ReferenceResource(adssd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adssd *Resource) ImportState(state io.Reader) error {
	adssd.state = &awsDirectoryServiceSharedDirectoryState{}
	if err := json.NewDecoder(state).Decode(adssd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adssd.Type(), adssd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adssd *Resource) State() (*awsDirectoryServiceSharedDirectoryState, bool) {
	return adssd.state, adssd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adssd *Resource) StateMust() *awsDirectoryServiceSharedDirectoryState {
	if adssd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adssd.Type(), adssd.LocalName()))
	}
	return adssd.state
}

// Args contains the configurations for aws_directory_service_shared_directory.
type Args struct {
	// DirectoryId: string, required
	DirectoryId terra.StringValue `hcl:"directory_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Method: string, optional
	Method terra.StringValue `hcl:"method,attr"`
	// Notes: string, optional
	Notes terra.StringValue `hcl:"notes,attr"`
	// Target: required
	Target *Target `hcl:"target,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsDirectoryServiceSharedDirectoryAttributes struct {
	ref terra.Reference
}

// DirectoryId returns a reference to field directory_id of aws_directory_service_shared_directory.
func (adssd awsDirectoryServiceSharedDirectoryAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(adssd.ref.Append("directory_id"))
}

// Id returns a reference to field id of aws_directory_service_shared_directory.
func (adssd awsDirectoryServiceSharedDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adssd.ref.Append("id"))
}

// Method returns a reference to field method of aws_directory_service_shared_directory.
func (adssd awsDirectoryServiceSharedDirectoryAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(adssd.ref.Append("method"))
}

// Notes returns a reference to field notes of aws_directory_service_shared_directory.
func (adssd awsDirectoryServiceSharedDirectoryAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(adssd.ref.Append("notes"))
}

// SharedDirectoryId returns a reference to field shared_directory_id of aws_directory_service_shared_directory.
func (adssd awsDirectoryServiceSharedDirectoryAttributes) SharedDirectoryId() terra.StringValue {
	return terra.ReferenceAsString(adssd.ref.Append("shared_directory_id"))
}

func (adssd awsDirectoryServiceSharedDirectoryAttributes) Target() terra.ListValue[TargetAttributes] {
	return terra.ReferenceAsList[TargetAttributes](adssd.ref.Append("target"))
}

func (adssd awsDirectoryServiceSharedDirectoryAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adssd.ref.Append("timeouts"))
}

type awsDirectoryServiceSharedDirectoryState struct {
	DirectoryId       string         `json:"directory_id"`
	Id                string         `json:"id"`
	Method            string         `json:"method"`
	Notes             string         `json:"notes"`
	SharedDirectoryId string         `json:"shared_directory_id"`
	Target            []TargetState  `json:"target"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
