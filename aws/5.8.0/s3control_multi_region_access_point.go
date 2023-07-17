// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3controlmultiregionaccesspoint "github.com/golingon/terraproviders/aws/5.8.0/s3controlmultiregionaccesspoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3ControlMultiRegionAccessPoint creates a new instance of [S3ControlMultiRegionAccessPoint].
func NewS3ControlMultiRegionAccessPoint(name string, args S3ControlMultiRegionAccessPointArgs) *S3ControlMultiRegionAccessPoint {
	return &S3ControlMultiRegionAccessPoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3ControlMultiRegionAccessPoint)(nil)

// S3ControlMultiRegionAccessPoint represents the Terraform resource aws_s3control_multi_region_access_point.
type S3ControlMultiRegionAccessPoint struct {
	Name      string
	Args      S3ControlMultiRegionAccessPointArgs
	state     *s3ControlMultiRegionAccessPointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3ControlMultiRegionAccessPoint].
func (smrap *S3ControlMultiRegionAccessPoint) Type() string {
	return "aws_s3control_multi_region_access_point"
}

// LocalName returns the local name for [S3ControlMultiRegionAccessPoint].
func (smrap *S3ControlMultiRegionAccessPoint) LocalName() string {
	return smrap.Name
}

// Configuration returns the configuration (args) for [S3ControlMultiRegionAccessPoint].
func (smrap *S3ControlMultiRegionAccessPoint) Configuration() interface{} {
	return smrap.Args
}

// DependOn is used for other resources to depend on [S3ControlMultiRegionAccessPoint].
func (smrap *S3ControlMultiRegionAccessPoint) DependOn() terra.Reference {
	return terra.ReferenceResource(smrap)
}

// Dependencies returns the list of resources [S3ControlMultiRegionAccessPoint] depends_on.
func (smrap *S3ControlMultiRegionAccessPoint) Dependencies() terra.Dependencies {
	return smrap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3ControlMultiRegionAccessPoint].
func (smrap *S3ControlMultiRegionAccessPoint) LifecycleManagement() *terra.Lifecycle {
	return smrap.Lifecycle
}

// Attributes returns the attributes for [S3ControlMultiRegionAccessPoint].
func (smrap *S3ControlMultiRegionAccessPoint) Attributes() s3ControlMultiRegionAccessPointAttributes {
	return s3ControlMultiRegionAccessPointAttributes{ref: terra.ReferenceResource(smrap)}
}

// ImportState imports the given attribute values into [S3ControlMultiRegionAccessPoint]'s state.
func (smrap *S3ControlMultiRegionAccessPoint) ImportState(av io.Reader) error {
	smrap.state = &s3ControlMultiRegionAccessPointState{}
	if err := json.NewDecoder(av).Decode(smrap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smrap.Type(), smrap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3ControlMultiRegionAccessPoint] has state.
func (smrap *S3ControlMultiRegionAccessPoint) State() (*s3ControlMultiRegionAccessPointState, bool) {
	return smrap.state, smrap.state != nil
}

// StateMust returns the state for [S3ControlMultiRegionAccessPoint]. Panics if the state is nil.
func (smrap *S3ControlMultiRegionAccessPoint) StateMust() *s3ControlMultiRegionAccessPointState {
	if smrap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smrap.Type(), smrap.LocalName()))
	}
	return smrap.state
}

// S3ControlMultiRegionAccessPointArgs contains the configurations for aws_s3control_multi_region_access_point.
type S3ControlMultiRegionAccessPointArgs struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Details: required
	Details *s3controlmultiregionaccesspoint.Details `hcl:"details,block" validate:"required"`
	// Timeouts: optional
	Timeouts *s3controlmultiregionaccesspoint.Timeouts `hcl:"timeouts,block"`
}
type s3ControlMultiRegionAccessPointAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_s3control_multi_region_access_point.
func (smrap s3ControlMultiRegionAccessPointAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("account_id"))
}

// Alias returns a reference to field alias of aws_s3control_multi_region_access_point.
func (smrap s3ControlMultiRegionAccessPointAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("alias"))
}

// Arn returns a reference to field arn of aws_s3control_multi_region_access_point.
func (smrap s3ControlMultiRegionAccessPointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("arn"))
}

// DomainName returns a reference to field domain_name of aws_s3control_multi_region_access_point.
func (smrap s3ControlMultiRegionAccessPointAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_s3control_multi_region_access_point.
func (smrap s3ControlMultiRegionAccessPointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("id"))
}

// Status returns a reference to field status of aws_s3control_multi_region_access_point.
func (smrap s3ControlMultiRegionAccessPointAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(smrap.ref.Append("status"))
}

func (smrap s3ControlMultiRegionAccessPointAttributes) Details() terra.ListValue[s3controlmultiregionaccesspoint.DetailsAttributes] {
	return terra.ReferenceAsList[s3controlmultiregionaccesspoint.DetailsAttributes](smrap.ref.Append("details"))
}

func (smrap s3ControlMultiRegionAccessPointAttributes) Timeouts() s3controlmultiregionaccesspoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[s3controlmultiregionaccesspoint.TimeoutsAttributes](smrap.ref.Append("timeouts"))
}

type s3ControlMultiRegionAccessPointState struct {
	AccountId  string                                         `json:"account_id"`
	Alias      string                                         `json:"alias"`
	Arn        string                                         `json:"arn"`
	DomainName string                                         `json:"domain_name"`
	Id         string                                         `json:"id"`
	Status     string                                         `json:"status"`
	Details    []s3controlmultiregionaccesspoint.DetailsState `json:"details"`
	Timeouts   *s3controlmultiregionaccesspoint.TimeoutsState `json:"timeouts"`
}
