// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementproductgroup "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagementproductgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApiManagementProductGroup(name string, args ApiManagementProductGroupArgs) *ApiManagementProductGroup {
	return &ApiManagementProductGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementProductGroup)(nil)

type ApiManagementProductGroup struct {
	Name  string
	Args  ApiManagementProductGroupArgs
	state *apiManagementProductGroupState
}

func (ampg *ApiManagementProductGroup) Type() string {
	return "azurerm_api_management_product_group"
}

func (ampg *ApiManagementProductGroup) LocalName() string {
	return ampg.Name
}

func (ampg *ApiManagementProductGroup) Configuration() interface{} {
	return ampg.Args
}

func (ampg *ApiManagementProductGroup) Attributes() apiManagementProductGroupAttributes {
	return apiManagementProductGroupAttributes{ref: terra.ReferenceResource(ampg)}
}

func (ampg *ApiManagementProductGroup) ImportState(av io.Reader) error {
	ampg.state = &apiManagementProductGroupState{}
	if err := json.NewDecoder(av).Decode(ampg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ampg.Type(), ampg.LocalName(), err)
	}
	return nil
}

func (ampg *ApiManagementProductGroup) State() (*apiManagementProductGroupState, bool) {
	return ampg.state, ampg.state != nil
}

func (ampg *ApiManagementProductGroup) StateMust() *apiManagementProductGroupState {
	if ampg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ampg.Type(), ampg.LocalName()))
	}
	return ampg.state
}

func (ampg *ApiManagementProductGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ampg)
}

type ApiManagementProductGroupArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// GroupName: string, required
	GroupName terra.StringValue `hcl:"group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProductId: string, required
	ProductId terra.StringValue `hcl:"product_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementproductgroup.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ApiManagementProductGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apiManagementProductGroupAttributes struct {
	ref terra.Reference
}

func (ampg apiManagementProductGroupAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceString(ampg.ref.Append("api_management_name"))
}

func (ampg apiManagementProductGroupAttributes) GroupName() terra.StringValue {
	return terra.ReferenceString(ampg.ref.Append("group_name"))
}

func (ampg apiManagementProductGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ampg.ref.Append("id"))
}

func (ampg apiManagementProductGroupAttributes) ProductId() terra.StringValue {
	return terra.ReferenceString(ampg.ref.Append("product_id"))
}

func (ampg apiManagementProductGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(ampg.ref.Append("resource_group_name"))
}

func (ampg apiManagementProductGroupAttributes) Timeouts() apimanagementproductgroup.TimeoutsAttributes {
	return terra.ReferenceSingle[apimanagementproductgroup.TimeoutsAttributes](ampg.ref.Append("timeouts"))
}

type apiManagementProductGroupState struct {
	ApiManagementName string                                   `json:"api_management_name"`
	GroupName         string                                   `json:"group_name"`
	Id                string                                   `json:"id"`
	ProductId         string                                   `json:"product_id"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	Timeouts          *apimanagementproductgroup.TimeoutsState `json:"timeouts"`
}
