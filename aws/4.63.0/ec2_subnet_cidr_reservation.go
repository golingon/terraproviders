// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2SubnetCidrReservation creates a new instance of [Ec2SubnetCidrReservation].
func NewEc2SubnetCidrReservation(name string, args Ec2SubnetCidrReservationArgs) *Ec2SubnetCidrReservation {
	return &Ec2SubnetCidrReservation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2SubnetCidrReservation)(nil)

// Ec2SubnetCidrReservation represents the Terraform resource aws_ec2_subnet_cidr_reservation.
type Ec2SubnetCidrReservation struct {
	Name      string
	Args      Ec2SubnetCidrReservationArgs
	state     *ec2SubnetCidrReservationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2SubnetCidrReservation].
func (escr *Ec2SubnetCidrReservation) Type() string {
	return "aws_ec2_subnet_cidr_reservation"
}

// LocalName returns the local name for [Ec2SubnetCidrReservation].
func (escr *Ec2SubnetCidrReservation) LocalName() string {
	return escr.Name
}

// Configuration returns the configuration (args) for [Ec2SubnetCidrReservation].
func (escr *Ec2SubnetCidrReservation) Configuration() interface{} {
	return escr.Args
}

// DependOn is used for other resources to depend on [Ec2SubnetCidrReservation].
func (escr *Ec2SubnetCidrReservation) DependOn() terra.Reference {
	return terra.ReferenceResource(escr)
}

// Dependencies returns the list of resources [Ec2SubnetCidrReservation] depends_on.
func (escr *Ec2SubnetCidrReservation) Dependencies() terra.Dependencies {
	return escr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2SubnetCidrReservation].
func (escr *Ec2SubnetCidrReservation) LifecycleManagement() *terra.Lifecycle {
	return escr.Lifecycle
}

// Attributes returns the attributes for [Ec2SubnetCidrReservation].
func (escr *Ec2SubnetCidrReservation) Attributes() ec2SubnetCidrReservationAttributes {
	return ec2SubnetCidrReservationAttributes{ref: terra.ReferenceResource(escr)}
}

// ImportState imports the given attribute values into [Ec2SubnetCidrReservation]'s state.
func (escr *Ec2SubnetCidrReservation) ImportState(av io.Reader) error {
	escr.state = &ec2SubnetCidrReservationState{}
	if err := json.NewDecoder(av).Decode(escr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", escr.Type(), escr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2SubnetCidrReservation] has state.
func (escr *Ec2SubnetCidrReservation) State() (*ec2SubnetCidrReservationState, bool) {
	return escr.state, escr.state != nil
}

// StateMust returns the state for [Ec2SubnetCidrReservation]. Panics if the state is nil.
func (escr *Ec2SubnetCidrReservation) StateMust() *ec2SubnetCidrReservationState {
	if escr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", escr.Type(), escr.LocalName()))
	}
	return escr.state
}

// Ec2SubnetCidrReservationArgs contains the configurations for aws_ec2_subnet_cidr_reservation.
type Ec2SubnetCidrReservationArgs struct {
	// CidrBlock: string, required
	CidrBlock terra.StringValue `hcl:"cidr_block,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReservationType: string, required
	ReservationType terra.StringValue `hcl:"reservation_type,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
}
type ec2SubnetCidrReservationAttributes struct {
	ref terra.Reference
}

// CidrBlock returns a reference to field cidr_block of aws_ec2_subnet_cidr_reservation.
func (escr ec2SubnetCidrReservationAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(escr.ref.Append("cidr_block"))
}

// Description returns a reference to field description of aws_ec2_subnet_cidr_reservation.
func (escr ec2SubnetCidrReservationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(escr.ref.Append("description"))
}

// Id returns a reference to field id of aws_ec2_subnet_cidr_reservation.
func (escr ec2SubnetCidrReservationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(escr.ref.Append("id"))
}

// OwnerId returns a reference to field owner_id of aws_ec2_subnet_cidr_reservation.
func (escr ec2SubnetCidrReservationAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(escr.ref.Append("owner_id"))
}

// ReservationType returns a reference to field reservation_type of aws_ec2_subnet_cidr_reservation.
func (escr ec2SubnetCidrReservationAttributes) ReservationType() terra.StringValue {
	return terra.ReferenceAsString(escr.ref.Append("reservation_type"))
}

// SubnetId returns a reference to field subnet_id of aws_ec2_subnet_cidr_reservation.
func (escr ec2SubnetCidrReservationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(escr.ref.Append("subnet_id"))
}

type ec2SubnetCidrReservationState struct {
	CidrBlock       string `json:"cidr_block"`
	Description     string `json:"description"`
	Id              string `json:"id"`
	OwnerId         string `json:"owner_id"`
	ReservationType string `json:"reservation_type"`
	SubnetId        string `json:"subnet_id"`
}
