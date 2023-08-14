// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dxgateway "github.com/golingon/terraproviders/aws/5.12.0/dxgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDxGateway creates a new instance of [DxGateway].
func NewDxGateway(name string, args DxGatewayArgs) *DxGateway {
	return &DxGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxGateway)(nil)

// DxGateway represents the Terraform resource aws_dx_gateway.
type DxGateway struct {
	Name      string
	Args      DxGatewayArgs
	state     *dxGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DxGateway].
func (dg *DxGateway) Type() string {
	return "aws_dx_gateway"
}

// LocalName returns the local name for [DxGateway].
func (dg *DxGateway) LocalName() string {
	return dg.Name
}

// Configuration returns the configuration (args) for [DxGateway].
func (dg *DxGateway) Configuration() interface{} {
	return dg.Args
}

// DependOn is used for other resources to depend on [DxGateway].
func (dg *DxGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(dg)
}

// Dependencies returns the list of resources [DxGateway] depends_on.
func (dg *DxGateway) Dependencies() terra.Dependencies {
	return dg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DxGateway].
func (dg *DxGateway) LifecycleManagement() *terra.Lifecycle {
	return dg.Lifecycle
}

// Attributes returns the attributes for [DxGateway].
func (dg *DxGateway) Attributes() dxGatewayAttributes {
	return dxGatewayAttributes{ref: terra.ReferenceResource(dg)}
}

// ImportState imports the given attribute values into [DxGateway]'s state.
func (dg *DxGateway) ImportState(av io.Reader) error {
	dg.state = &dxGatewayState{}
	if err := json.NewDecoder(av).Decode(dg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dg.Type(), dg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DxGateway] has state.
func (dg *DxGateway) State() (*dxGatewayState, bool) {
	return dg.state, dg.state != nil
}

// StateMust returns the state for [DxGateway]. Panics if the state is nil.
func (dg *DxGateway) StateMust() *dxGatewayState {
	if dg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dg.Type(), dg.LocalName()))
	}
	return dg.state
}

// DxGatewayArgs contains the configurations for aws_dx_gateway.
type DxGatewayArgs struct {
	// AmazonSideAsn: string, required
	AmazonSideAsn terra.StringValue `hcl:"amazon_side_asn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dxgateway.Timeouts `hcl:"timeouts,block"`
}
type dxGatewayAttributes struct {
	ref terra.Reference
}

// AmazonSideAsn returns a reference to field amazon_side_asn of aws_dx_gateway.
func (dg dxGatewayAttributes) AmazonSideAsn() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("amazon_side_asn"))
}

// Id returns a reference to field id of aws_dx_gateway.
func (dg dxGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("id"))
}

// Name returns a reference to field name of aws_dx_gateway.
func (dg dxGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("name"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_dx_gateway.
func (dg dxGatewayAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dg.ref.Append("owner_account_id"))
}

func (dg dxGatewayAttributes) Timeouts() dxgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dxgateway.TimeoutsAttributes](dg.ref.Append("timeouts"))
}

type dxGatewayState struct {
	AmazonSideAsn  string                   `json:"amazon_side_asn"`
	Id             string                   `json:"id"`
	Name           string                   `json:"name"`
	OwnerAccountId string                   `json:"owner_account_id"`
	Timeouts       *dxgateway.TimeoutsState `json:"timeouts"`
}
