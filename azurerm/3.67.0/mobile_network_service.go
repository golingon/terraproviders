// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mobilenetworkservice "github.com/golingon/terraproviders/azurerm/3.67.0/mobilenetworkservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMobileNetworkService creates a new instance of [MobileNetworkService].
func NewMobileNetworkService(name string, args MobileNetworkServiceArgs) *MobileNetworkService {
	return &MobileNetworkService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MobileNetworkService)(nil)

// MobileNetworkService represents the Terraform resource azurerm_mobile_network_service.
type MobileNetworkService struct {
	Name      string
	Args      MobileNetworkServiceArgs
	state     *mobileNetworkServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MobileNetworkService].
func (mns *MobileNetworkService) Type() string {
	return "azurerm_mobile_network_service"
}

// LocalName returns the local name for [MobileNetworkService].
func (mns *MobileNetworkService) LocalName() string {
	return mns.Name
}

// Configuration returns the configuration (args) for [MobileNetworkService].
func (mns *MobileNetworkService) Configuration() interface{} {
	return mns.Args
}

// DependOn is used for other resources to depend on [MobileNetworkService].
func (mns *MobileNetworkService) DependOn() terra.Reference {
	return terra.ReferenceResource(mns)
}

// Dependencies returns the list of resources [MobileNetworkService] depends_on.
func (mns *MobileNetworkService) Dependencies() terra.Dependencies {
	return mns.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MobileNetworkService].
func (mns *MobileNetworkService) LifecycleManagement() *terra.Lifecycle {
	return mns.Lifecycle
}

// Attributes returns the attributes for [MobileNetworkService].
func (mns *MobileNetworkService) Attributes() mobileNetworkServiceAttributes {
	return mobileNetworkServiceAttributes{ref: terra.ReferenceResource(mns)}
}

// ImportState imports the given attribute values into [MobileNetworkService]'s state.
func (mns *MobileNetworkService) ImportState(av io.Reader) error {
	mns.state = &mobileNetworkServiceState{}
	if err := json.NewDecoder(av).Decode(mns.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mns.Type(), mns.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MobileNetworkService] has state.
func (mns *MobileNetworkService) State() (*mobileNetworkServiceState, bool) {
	return mns.state, mns.state != nil
}

// StateMust returns the state for [MobileNetworkService]. Panics if the state is nil.
func (mns *MobileNetworkService) StateMust() *mobileNetworkServiceState {
	if mns.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mns.Type(), mns.LocalName()))
	}
	return mns.state
}

// MobileNetworkServiceArgs contains the configurations for azurerm_mobile_network_service.
type MobileNetworkServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MobileNetworkId: string, required
	MobileNetworkId terra.StringValue `hcl:"mobile_network_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServicePrecedence: number, required
	ServicePrecedence terra.NumberValue `hcl:"service_precedence,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// PccRule: min=1
	PccRule []mobilenetworkservice.PccRule `hcl:"pcc_rule,block" validate:"min=1"`
	// ServiceQosPolicy: optional
	ServiceQosPolicy *mobilenetworkservice.ServiceQosPolicy `hcl:"service_qos_policy,block"`
	// Timeouts: optional
	Timeouts *mobilenetworkservice.Timeouts `hcl:"timeouts,block"`
}
type mobileNetworkServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mobile_network_service.
func (mns mobileNetworkServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_service.
func (mns mobileNetworkServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_service.
func (mns mobileNetworkServiceAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_service.
func (mns mobileNetworkServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("name"))
}

// ServicePrecedence returns a reference to field service_precedence of azurerm_mobile_network_service.
func (mns mobileNetworkServiceAttributes) ServicePrecedence() terra.NumberValue {
	return terra.ReferenceAsNumber(mns.ref.Append("service_precedence"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_service.
func (mns mobileNetworkServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mns.ref.Append("tags"))
}

func (mns mobileNetworkServiceAttributes) PccRule() terra.ListValue[mobilenetworkservice.PccRuleAttributes] {
	return terra.ReferenceAsList[mobilenetworkservice.PccRuleAttributes](mns.ref.Append("pcc_rule"))
}

func (mns mobileNetworkServiceAttributes) ServiceQosPolicy() terra.ListValue[mobilenetworkservice.ServiceQosPolicyAttributes] {
	return terra.ReferenceAsList[mobilenetworkservice.ServiceQosPolicyAttributes](mns.ref.Append("service_qos_policy"))
}

func (mns mobileNetworkServiceAttributes) Timeouts() mobilenetworkservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mobilenetworkservice.TimeoutsAttributes](mns.ref.Append("timeouts"))
}

type mobileNetworkServiceState struct {
	Id                string                                       `json:"id"`
	Location          string                                       `json:"location"`
	MobileNetworkId   string                                       `json:"mobile_network_id"`
	Name              string                                       `json:"name"`
	ServicePrecedence float64                                      `json:"service_precedence"`
	Tags              map[string]string                            `json:"tags"`
	PccRule           []mobilenetworkservice.PccRuleState          `json:"pcc_rule"`
	ServiceQosPolicy  []mobilenetworkservice.ServiceQosPolicyState `json:"service_qos_policy"`
	Timeouts          *mobilenetworkservice.TimeoutsState          `json:"timeouts"`
}
