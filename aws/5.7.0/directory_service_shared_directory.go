// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	directoryserviceshareddirectory "github.com/golingon/terraproviders/aws/5.7.0/directoryserviceshareddirectory"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDirectoryServiceSharedDirectory creates a new instance of [DirectoryServiceSharedDirectory].
func NewDirectoryServiceSharedDirectory(name string, args DirectoryServiceSharedDirectoryArgs) *DirectoryServiceSharedDirectory {
	return &DirectoryServiceSharedDirectory{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DirectoryServiceSharedDirectory)(nil)

// DirectoryServiceSharedDirectory represents the Terraform resource aws_directory_service_shared_directory.
type DirectoryServiceSharedDirectory struct {
	Name      string
	Args      DirectoryServiceSharedDirectoryArgs
	state     *directoryServiceSharedDirectoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DirectoryServiceSharedDirectory].
func (dssd *DirectoryServiceSharedDirectory) Type() string {
	return "aws_directory_service_shared_directory"
}

// LocalName returns the local name for [DirectoryServiceSharedDirectory].
func (dssd *DirectoryServiceSharedDirectory) LocalName() string {
	return dssd.Name
}

// Configuration returns the configuration (args) for [DirectoryServiceSharedDirectory].
func (dssd *DirectoryServiceSharedDirectory) Configuration() interface{} {
	return dssd.Args
}

// DependOn is used for other resources to depend on [DirectoryServiceSharedDirectory].
func (dssd *DirectoryServiceSharedDirectory) DependOn() terra.Reference {
	return terra.ReferenceResource(dssd)
}

// Dependencies returns the list of resources [DirectoryServiceSharedDirectory] depends_on.
func (dssd *DirectoryServiceSharedDirectory) Dependencies() terra.Dependencies {
	return dssd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DirectoryServiceSharedDirectory].
func (dssd *DirectoryServiceSharedDirectory) LifecycleManagement() *terra.Lifecycle {
	return dssd.Lifecycle
}

// Attributes returns the attributes for [DirectoryServiceSharedDirectory].
func (dssd *DirectoryServiceSharedDirectory) Attributes() directoryServiceSharedDirectoryAttributes {
	return directoryServiceSharedDirectoryAttributes{ref: terra.ReferenceResource(dssd)}
}

// ImportState imports the given attribute values into [DirectoryServiceSharedDirectory]'s state.
func (dssd *DirectoryServiceSharedDirectory) ImportState(av io.Reader) error {
	dssd.state = &directoryServiceSharedDirectoryState{}
	if err := json.NewDecoder(av).Decode(dssd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dssd.Type(), dssd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DirectoryServiceSharedDirectory] has state.
func (dssd *DirectoryServiceSharedDirectory) State() (*directoryServiceSharedDirectoryState, bool) {
	return dssd.state, dssd.state != nil
}

// StateMust returns the state for [DirectoryServiceSharedDirectory]. Panics if the state is nil.
func (dssd *DirectoryServiceSharedDirectory) StateMust() *directoryServiceSharedDirectoryState {
	if dssd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dssd.Type(), dssd.LocalName()))
	}
	return dssd.state
}

// DirectoryServiceSharedDirectoryArgs contains the configurations for aws_directory_service_shared_directory.
type DirectoryServiceSharedDirectoryArgs struct {
	// DirectoryId: string, required
	DirectoryId terra.StringValue `hcl:"directory_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Method: string, optional
	Method terra.StringValue `hcl:"method,attr"`
	// Notes: string, optional
	Notes terra.StringValue `hcl:"notes,attr"`
	// Target: required
	Target *directoryserviceshareddirectory.Target `hcl:"target,block" validate:"required"`
	// Timeouts: optional
	Timeouts *directoryserviceshareddirectory.Timeouts `hcl:"timeouts,block"`
}
type directoryServiceSharedDirectoryAttributes struct {
	ref terra.Reference
}

// DirectoryId returns a reference to field directory_id of aws_directory_service_shared_directory.
func (dssd directoryServiceSharedDirectoryAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(dssd.ref.Append("directory_id"))
}

// Id returns a reference to field id of aws_directory_service_shared_directory.
func (dssd directoryServiceSharedDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dssd.ref.Append("id"))
}

// Method returns a reference to field method of aws_directory_service_shared_directory.
func (dssd directoryServiceSharedDirectoryAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(dssd.ref.Append("method"))
}

// Notes returns a reference to field notes of aws_directory_service_shared_directory.
func (dssd directoryServiceSharedDirectoryAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(dssd.ref.Append("notes"))
}

// SharedDirectoryId returns a reference to field shared_directory_id of aws_directory_service_shared_directory.
func (dssd directoryServiceSharedDirectoryAttributes) SharedDirectoryId() terra.StringValue {
	return terra.ReferenceAsString(dssd.ref.Append("shared_directory_id"))
}

func (dssd directoryServiceSharedDirectoryAttributes) Target() terra.ListValue[directoryserviceshareddirectory.TargetAttributes] {
	return terra.ReferenceAsList[directoryserviceshareddirectory.TargetAttributes](dssd.ref.Append("target"))
}

func (dssd directoryServiceSharedDirectoryAttributes) Timeouts() directoryserviceshareddirectory.TimeoutsAttributes {
	return terra.ReferenceAsSingle[directoryserviceshareddirectory.TimeoutsAttributes](dssd.ref.Append("timeouts"))
}

type directoryServiceSharedDirectoryState struct {
	DirectoryId       string                                         `json:"directory_id"`
	Id                string                                         `json:"id"`
	Method            string                                         `json:"method"`
	Notes             string                                         `json:"notes"`
	SharedDirectoryId string                                         `json:"shared_directory_id"`
	Target            []directoryserviceshareddirectory.TargetState  `json:"target"`
	Timeouts          *directoryserviceshareddirectory.TimeoutsState `json:"timeouts"`
}
