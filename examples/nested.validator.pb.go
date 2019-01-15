// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/nested.proto

package validator_examples

import (
	fmt "fmt"
	_ "github.com/georgeciubotaru/go-proto-validators"
	github_com_georgeciubotaru_go_proto_validators "github.com/georgeciubotaru/go-proto-validators"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *InnerMessage) Validate() error {
	if _, ok := InType_name[int32(this.Test)]; !ok {
		return github_com_georgeciubotaru_go_proto_validators.FieldError("Test", fmt.Errorf(`value '%v' must be a valid InType field`, this.Test))
	}
	for _, item := range this.TestArray {
		if _, ok := InType_name[int32(item)]; !ok {
			return github_com_georgeciubotaru_go_proto_validators.FieldError("TestArray", fmt.Errorf(`value '%v' must be a valid InType field`, item))
		}
	}
	return nil
}