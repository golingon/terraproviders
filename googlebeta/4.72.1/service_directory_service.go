// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	servicedirectoryservice "github.com/golingon/terraproviders/googlebeta/4.72.1/servicedirectoryservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceDirectoryService creates a new instance of [ServiceDirectoryService].
func NewServiceDirectoryService(name string, args ServiceDirectoryServiceArgs) *ServiceDirectoryService {
	return &ServiceDirectoryService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceDirectoryService)(nil)

// ServiceDirectoryService represents the Terraform resource google_service_directory_service.
type ServiceDirectoryService struct {
	Name      string
	Args      ServiceDirectoryServiceArgs
	state     *serviceDirectoryServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceDirectoryService].
func (sds *ServiceDirectoryService) Type() string {
	return "google_service_directory_service"
}

// LocalName returns the local name for [ServiceDirectoryService].
func (sds *ServiceDirectoryService) LocalName() string {
	return sds.Name
}

// Configuration returns the configuration (args) for [ServiceDirectoryService].
func (sds *ServiceDirectoryService) Configuration() interface{} {
	return sds.Args
}

// DependOn is used for other resources to depend on [ServiceDirectoryService].
func (sds *ServiceDirectoryService) DependOn() terra.Reference {
	return terra.ReferenceResource(sds)
}

// Dependencies returns the list of resources [ServiceDirectoryService] depends_on.
func (sds *ServiceDirectoryService) Dependencies() terra.Dependencies {
	return sds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceDirectoryService].
func (sds *ServiceDirectoryService) LifecycleManagement() *terra.Lifecycle {
	return sds.Lifecycle
}

// Attributes returns the attributes for [ServiceDirectoryService].
func (sds *ServiceDirectoryService) Attributes() serviceDirectoryServiceAttributes {
	return serviceDirectoryServiceAttributes{ref: terra.ReferenceResource(sds)}
}

// ImportState imports the given attribute values into [ServiceDirectoryService]'s state.
func (sds *ServiceDirectoryService) ImportState(av io.Reader) error {
	sds.state = &serviceDirectoryServiceState{}
	if err := json.NewDecoder(av).Decode(sds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sds.Type(), sds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceDirectoryService] has state.
func (sds *ServiceDirectoryService) State() (*serviceDirectoryServiceState, bool) {
	return sds.state, sds.state != nil
}

// StateMust returns the state for [ServiceDirectoryService]. Panics if the state is nil.
func (sds *ServiceDirectoryService) StateMust() *serviceDirectoryServiceState {
	if sds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sds.Type(), sds.LocalName()))
	}
	return sds.state
}

// ServiceDirectoryServiceArgs contains the configurations for google_service_directory_service.
type ServiceDirectoryServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
	// ServiceId: string, required
	ServiceId terra.StringValue `hcl:"service_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *servicedirectoryservice.Timeouts `hcl:"timeouts,block"`
}
type serviceDirectoryServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_service_directory_service.
func (sds serviceDirectoryServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("id"))
}

// Metadata returns a reference to field metadata of google_service_directory_service.
func (sds serviceDirectoryServiceAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sds.ref.Append("metadata"))
}

// Name returns a reference to field name of google_service_directory_service.
func (sds serviceDirectoryServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("name"))
}

// Namespace returns a reference to field namespace of google_service_directory_service.
func (sds serviceDirectoryServiceAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("namespace"))
}

// ServiceId returns a reference to field service_id of google_service_directory_service.
func (sds serviceDirectoryServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("service_id"))
}

func (sds serviceDirectoryServiceAttributes) Timeouts() servicedirectoryservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicedirectoryservice.TimeoutsAttributes](sds.ref.Append("timeouts"))
}

type serviceDirectoryServiceState struct {
	Id        string                                 `json:"id"`
	Metadata  map[string]string                      `json:"metadata"`
	Name      string                                 `json:"name"`
	Namespace string                                 `json:"namespace"`
	ServiceId string                                 `json:"service_id"`
	Timeouts  *servicedirectoryservice.TimeoutsState `json:"timeouts"`
}
