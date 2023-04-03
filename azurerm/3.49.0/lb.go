// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	lb "github.com/golingon/terraproviders/azurerm/3.49.0/lb"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLb creates a new instance of [Lb].
func NewLb(name string, args LbArgs) *Lb {
	return &Lb{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Lb)(nil)

// Lb represents the Terraform resource azurerm_lb.
type Lb struct {
	Name      string
	Args      LbArgs
	state     *lbState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Lb].
func (l *Lb) Type() string {
	return "azurerm_lb"
}

// LocalName returns the local name for [Lb].
func (l *Lb) LocalName() string {
	return l.Name
}

// Configuration returns the configuration (args) for [Lb].
func (l *Lb) Configuration() interface{} {
	return l.Args
}

// DependOn is used for other resources to depend on [Lb].
func (l *Lb) DependOn() terra.Reference {
	return terra.ReferenceResource(l)
}

// Dependencies returns the list of resources [Lb] depends_on.
func (l *Lb) Dependencies() terra.Dependencies {
	return l.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Lb].
func (l *Lb) LifecycleManagement() *terra.Lifecycle {
	return l.Lifecycle
}

// Attributes returns the attributes for [Lb].
func (l *Lb) Attributes() lbAttributes {
	return lbAttributes{ref: terra.ReferenceResource(l)}
}

// ImportState imports the given attribute values into [Lb]'s state.
func (l *Lb) ImportState(av io.Reader) error {
	l.state = &lbState{}
	if err := json.NewDecoder(av).Decode(l.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", l.Type(), l.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Lb] has state.
func (l *Lb) State() (*lbState, bool) {
	return l.state, l.state != nil
}

// StateMust returns the state for [Lb]. Panics if the state is nil.
func (l *Lb) StateMust() *lbState {
	if l.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", l.Type(), l.LocalName()))
	}
	return l.state
}

// LbArgs contains the configurations for azurerm_lb.
type LbArgs struct {
	// EdgeZone: string, optional
	EdgeZone terra.StringValue `hcl:"edge_zone,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// SkuTier: string, optional
	SkuTier terra.StringValue `hcl:"sku_tier,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// FrontendIpConfiguration: min=0
	FrontendIpConfiguration []lb.FrontendIpConfiguration `hcl:"frontend_ip_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *lb.Timeouts `hcl:"timeouts,block"`
}
type lbAttributes struct {
	ref terra.Reference
}

// EdgeZone returns a reference to field edge_zone of azurerm_lb.
func (l lbAttributes) EdgeZone() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("edge_zone"))
}

// Id returns a reference to field id of azurerm_lb.
func (l lbAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_lb.
func (l lbAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_lb.
func (l lbAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("name"))
}

// PrivateIpAddress returns a reference to field private_ip_address of azurerm_lb.
func (l lbAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("private_ip_address"))
}

// PrivateIpAddresses returns a reference to field private_ip_addresses of azurerm_lb.
func (l lbAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](l.ref.Append("private_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_lb.
func (l lbAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_lb.
func (l lbAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("sku"))
}

// SkuTier returns a reference to field sku_tier of azurerm_lb.
func (l lbAttributes) SkuTier() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("sku_tier"))
}

// Tags returns a reference to field tags of azurerm_lb.
func (l lbAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](l.ref.Append("tags"))
}

func (l lbAttributes) FrontendIpConfiguration() terra.ListValue[lb.FrontendIpConfigurationAttributes] {
	return terra.ReferenceAsList[lb.FrontendIpConfigurationAttributes](l.ref.Append("frontend_ip_configuration"))
}

func (l lbAttributes) Timeouts() lb.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lb.TimeoutsAttributes](l.ref.Append("timeouts"))
}

type lbState struct {
	EdgeZone                string                            `json:"edge_zone"`
	Id                      string                            `json:"id"`
	Location                string                            `json:"location"`
	Name                    string                            `json:"name"`
	PrivateIpAddress        string                            `json:"private_ip_address"`
	PrivateIpAddresses      []string                          `json:"private_ip_addresses"`
	ResourceGroupName       string                            `json:"resource_group_name"`
	Sku                     string                            `json:"sku"`
	SkuTier                 string                            `json:"sku_tier"`
	Tags                    map[string]string                 `json:"tags"`
	FrontendIpConfiguration []lb.FrontendIpConfigurationState `json:"frontend_ip_configuration"`
	Timeouts                *lb.TimeoutsState                 `json:"timeouts"`
}