// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudconfigurationservice "github.com/golingon/terraproviders/azurerm/3.49.0/springcloudconfigurationservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSpringCloudConfigurationService(name string, args SpringCloudConfigurationServiceArgs) *SpringCloudConfigurationService {
	return &SpringCloudConfigurationService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudConfigurationService)(nil)

type SpringCloudConfigurationService struct {
	Name  string
	Args  SpringCloudConfigurationServiceArgs
	state *springCloudConfigurationServiceState
}

func (sccs *SpringCloudConfigurationService) Type() string {
	return "azurerm_spring_cloud_configuration_service"
}

func (sccs *SpringCloudConfigurationService) LocalName() string {
	return sccs.Name
}

func (sccs *SpringCloudConfigurationService) Configuration() interface{} {
	return sccs.Args
}

func (sccs *SpringCloudConfigurationService) Attributes() springCloudConfigurationServiceAttributes {
	return springCloudConfigurationServiceAttributes{ref: terra.ReferenceResource(sccs)}
}

func (sccs *SpringCloudConfigurationService) ImportState(av io.Reader) error {
	sccs.state = &springCloudConfigurationServiceState{}
	if err := json.NewDecoder(av).Decode(sccs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sccs.Type(), sccs.LocalName(), err)
	}
	return nil
}

func (sccs *SpringCloudConfigurationService) State() (*springCloudConfigurationServiceState, bool) {
	return sccs.state, sccs.state != nil
}

func (sccs *SpringCloudConfigurationService) StateMust() *springCloudConfigurationServiceState {
	if sccs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sccs.Type(), sccs.LocalName()))
	}
	return sccs.state
}

func (sccs *SpringCloudConfigurationService) DependOn() terra.Reference {
	return terra.ReferenceResource(sccs)
}

type SpringCloudConfigurationServiceArgs struct {
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
	// DependsOn contains resources that SpringCloudConfigurationService depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type springCloudConfigurationServiceAttributes struct {
	ref terra.Reference
}

func (sccs springCloudConfigurationServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sccs.ref.Append("id"))
}

func (sccs springCloudConfigurationServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sccs.ref.Append("name"))
}

func (sccs springCloudConfigurationServiceAttributes) SpringCloudServiceId() terra.StringValue {
	return terra.ReferenceString(sccs.ref.Append("spring_cloud_service_id"))
}

func (sccs springCloudConfigurationServiceAttributes) Repository() terra.ListValue[springcloudconfigurationservice.RepositoryAttributes] {
	return terra.ReferenceList[springcloudconfigurationservice.RepositoryAttributes](sccs.ref.Append("repository"))
}

func (sccs springCloudConfigurationServiceAttributes) Timeouts() springcloudconfigurationservice.TimeoutsAttributes {
	return terra.ReferenceSingle[springcloudconfigurationservice.TimeoutsAttributes](sccs.ref.Append("timeouts"))
}

type springCloudConfigurationServiceState struct {
	Id                   string                                            `json:"id"`
	Name                 string                                            `json:"name"`
	SpringCloudServiceId string                                            `json:"spring_cloud_service_id"`
	Repository           []springcloudconfigurationservice.RepositoryState `json:"repository"`
	Timeouts             *springcloudconfigurationservice.TimeoutsState    `json:"timeouts"`
}
