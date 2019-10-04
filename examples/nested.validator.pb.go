// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/nested.proto

package validator_examples

import (
	_ "examples/shared"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/holdex/go-proto-validators"
	bitbucket_org_holdex_go_proto_validators "github.com/holdex/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *InnerMessage) Validate() error {
	if nil == this.Test {
		return bitbucket_org_holdex_go_proto_validators.FieldError("Test", fmt.Errorf("message must exist"))
	}
	if this.Test != nil {
		if err := bitbucket_org_holdex_go_proto_validators.CallValidatorIfExists(this.Test); err != nil {
			return bitbucket_org_holdex_go_proto_validators.FieldError("Test", err)
		}
	}
	for _, item := range this.TestArray {
		if !(item > 0) {
			return bitbucket_org_holdex_go_proto_validators.FieldError("TestArray", fmt.Errorf(`value '%v' must be greater than '0'`, item))
		}
		if !(item < 2) {
			return bitbucket_org_holdex_go_proto_validators.FieldError("TestArray", fmt.Errorf(`value '%v' must be less than '2'`, item))
		}
	}
	return nil
}
