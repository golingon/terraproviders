// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package chimesdkvoicesipmediaapplication

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Endpoints struct {
	// LambdaArn: string, required
	LambdaArn terra.StringValue `hcl:"lambda_arn,attr" validate:"required"`
}

type EndpointsAttributes struct {
	ref terra.Reference
}

func (e EndpointsAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EndpointsAttributes) InternalWithRef(ref terra.Reference) EndpointsAttributes {
	return EndpointsAttributes{ref: ref}
}

func (e EndpointsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EndpointsAttributes) LambdaArn() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("lambda_arn"))
}

type EndpointsState struct {
	LambdaArn string `json:"lambda_arn"`
}
