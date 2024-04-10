// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataServiceAccountJwt creates a new instance of [DataServiceAccountJwt].
func NewDataServiceAccountJwt(name string, args DataServiceAccountJwtArgs) *DataServiceAccountJwt {
	return &DataServiceAccountJwt{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServiceAccountJwt)(nil)

// DataServiceAccountJwt represents the Terraform data resource google_service_account_jwt.
type DataServiceAccountJwt struct {
	Name string
	Args DataServiceAccountJwtArgs
}

// DataSource returns the Terraform object type for [DataServiceAccountJwt].
func (saj *DataServiceAccountJwt) DataSource() string {
	return "google_service_account_jwt"
}

// LocalName returns the local name for [DataServiceAccountJwt].
func (saj *DataServiceAccountJwt) LocalName() string {
	return saj.Name
}

// Configuration returns the configuration (args) for [DataServiceAccountJwt].
func (saj *DataServiceAccountJwt) Configuration() interface{} {
	return saj.Args
}

// Attributes returns the attributes for [DataServiceAccountJwt].
func (saj *DataServiceAccountJwt) Attributes() dataServiceAccountJwtAttributes {
	return dataServiceAccountJwtAttributes{ref: terra.ReferenceDataResource(saj)}
}

// DataServiceAccountJwtArgs contains the configurations for google_service_account_jwt.
type DataServiceAccountJwtArgs struct {
	// Delegates: set of string, optional
	Delegates terra.SetValue[terra.StringValue] `hcl:"delegates,attr"`
	// ExpiresIn: number, optional
	ExpiresIn terra.NumberValue `hcl:"expires_in,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Payload: string, required
	Payload terra.StringValue `hcl:"payload,attr" validate:"required"`
	// TargetServiceAccount: string, required
	TargetServiceAccount terra.StringValue `hcl:"target_service_account,attr" validate:"required"`
}
type dataServiceAccountJwtAttributes struct {
	ref terra.Reference
}

// Delegates returns a reference to field delegates of google_service_account_jwt.
func (saj dataServiceAccountJwtAttributes) Delegates() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](saj.ref.Append("delegates"))
}

// ExpiresIn returns a reference to field expires_in of google_service_account_jwt.
func (saj dataServiceAccountJwtAttributes) ExpiresIn() terra.NumberValue {
	return terra.ReferenceAsNumber(saj.ref.Append("expires_in"))
}

// Id returns a reference to field id of google_service_account_jwt.
func (saj dataServiceAccountJwtAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("id"))
}

// Jwt returns a reference to field jwt of google_service_account_jwt.
func (saj dataServiceAccountJwtAttributes) Jwt() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("jwt"))
}

// Payload returns a reference to field payload of google_service_account_jwt.
func (saj dataServiceAccountJwtAttributes) Payload() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("payload"))
}

// TargetServiceAccount returns a reference to field target_service_account of google_service_account_jwt.
func (saj dataServiceAccountJwtAttributes) TargetServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("target_service_account"))
}
