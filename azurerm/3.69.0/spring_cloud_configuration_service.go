// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudconfigurationservice "github.com/golingon/terraproviders/azurerm/3.69.0/springcloudconfigurationservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudConfigurationService creates a new instance of [SpringCloudConfigurationService].
func NewSpringCloudConfigurationService(name string, args SpringCloudConfigurationServiceArgs) *SpringCloudConfigurationService {
	return &SpringCloudConfigurationService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudConfigurationService)(nil)

// SpringCloudConfigurationService represents the Terraform resource azurerm_spring_cloud_configuration_service.
type SpringCloudConfigurationService struct {
	Name      string
	Args      SpringCloudConfigurationServiceArgs
	state     *springCloudConfigurationServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudConfigurationService].
func (sccs *SpringCloudConfigurationService) Type() string {
	return "azurerm_spring_cloud_configuration_service"
}

// LocalName returns the local name for [SpringCloudConfigurationService].
func (sccs *SpringCloudConfigurationService) LocalName() string {
	return sccs.Name
}

// Configuration returns the configuration (args) for [SpringCloudConfigurationService].
func (sccs *SpringCloudConfigurationService) Configuration() interface{} {
	return sccs.Args
}

// DependOn is used for other resources to depend on [SpringCloudConfigurationService].
func (sccs *SpringCloudConfigurationService) DependOn() terra.Reference {
	return terra.ReferenceResource(sccs)
}

// Dependencies returns the list of resources [SpringCloudConfigurationService] depends_on.
func (sccs *SpringCloudConfigurationService) Dependencies() terra.Dependencies {
	return sccs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudConfigurationService].
func (sccs *SpringCloudConfigurationService) LifecycleManagement() *terra.Lifecycle {
	return sccs.Lifecycle
}

// Attributes returns the attributes for [SpringCloudConfigurationService].
func (sccs *SpringCloudConfigurationService) Attributes() springCloudConfigurationServiceAttributes {
	return springCloudConfigurationServiceAttributes{ref: terra.ReferenceResource(sccs)}
}

// ImportState imports the given attribute values into [SpringCloudConfigurationService]'s state.
func (sccs *SpringCloudConfigurationService) ImportState(av io.Reader) error {
	sccs.state = &springCloudConfigurationServiceState{}
	if err := json.NewDecoder(av).Decode(sccs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sccs.Type(), sccs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudConfigurationService] has state.
func (sccs *SpringCloudConfigurationService) State() (*springCloudConfigurationServiceState, bool) {
	return sccs.state, sccs.state != nil
}

// StateMust returns the state for [SpringCloudConfigurationService]. Panics if the state is nil.
func (sccs *SpringCloudConfigurationService) StateMust() *springCloudConfigurationServiceState {
	if sccs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sccs.Type(), sccs.LocalName()))
	}
	return sccs.state
}

// SpringCloudConfigurationServiceArgs contains the configurations for azurerm_spring_cloud_configuration_service.
type SpringCloudConfigurationServiceArgs struct {
	// Generation: string, optional
	Generation terra.StringValue `hcl:"generation,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudServiceId: string, required
	SpringCloudServiceId terra.StringValue `hcl:"spring_cloud_service_id,attr" validate:"required"`
	// Repository: min=0
	Repository []springcloudconfigurationservice.Repository `hcl:"repository,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *springcloudconfigurationservice.Timeouts `hcl:"timeouts,block"`
}
type springCloudConfigurationServiceAttributes struct {
	ref terra.Reference
}

// Generation returns a reference to field generation of azurerm_spring_cloud_configuration_service.
func (sccs springCloudConfigurationServiceAttributes) Generation() terra.StringValue {
	return terra.ReferenceAsString(sccs.ref.Append("generation"))
}

// Id returns a reference to field id of azurerm_spring_cloud_configuration_service.
func (sccs springCloudConfigurationServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sccs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_configuration_service.
func (sccs springCloudConfigurationServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sccs.ref.Append("name"))
}

// SpringCloudServiceId returns a reference to field spring_cloud_service_id of azurerm_spring_cloud_configuration_service.
func (sccs springCloudConfigurationServiceAttributes) SpringCloudServiceId() terra.StringValue {
	return terra.ReferenceAsString(sccs.ref.Append("spring_cloud_service_id"))
}

func (sccs springCloudConfigurationServiceAttributes) Repository() terra.ListValue[springcloudconfigurationservice.RepositoryAttributes] {
	return terra.ReferenceAsList[springcloudconfigurationservice.RepositoryAttributes](sccs.ref.Append("repository"))
}

func (sccs springCloudConfigurationServiceAttributes) Timeouts() springcloudconfigurationservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudconfigurationservice.TimeoutsAttributes](sccs.ref.Append("timeouts"))
}

type springCloudConfigurationServiceState struct {
	Generation           string                                            `json:"generation"`
	Id                   string                                            `json:"id"`
	Name                 string                                            `json:"name"`
	SpringCloudServiceId string                                            `json:"spring_cloud_service_id"`
	Repository           []springcloudconfigurationservice.RepositoryState `json:"repository"`
	Timeouts             *springcloudconfigurationservice.TimeoutsState    `json:"timeouts"`
}