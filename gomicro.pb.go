// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: gomicro.proto

package gomicro

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request for geocoding and reverse geocoding operations.
type GeoRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to RequestType:
	//
	//	*GeoRequest_Query
	//	*GeoRequest_Coords
	RequestType   isGeoRequest_RequestType `protobuf_oneof:"request_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeoRequest) Reset() {
	*x = GeoRequest{}
	mi := &file_gomicro_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoRequest) ProtoMessage() {}

func (x *GeoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomicro_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoRequest.ProtoReflect.Descriptor instead.
func (*GeoRequest) Descriptor() ([]byte, []int) {
	return file_gomicro_proto_rawDescGZIP(), []int{0}
}

func (x *GeoRequest) GetRequestType() isGeoRequest_RequestType {
	if x != nil {
		return x.RequestType
	}
	return nil
}

func (x *GeoRequest) GetQuery() string {
	if x != nil {
		if x, ok := x.RequestType.(*GeoRequest_Query); ok {
			return x.Query
		}
	}
	return ""
}

func (x *GeoRequest) GetCoords() *Coordinates {
	if x != nil {
		if x, ok := x.RequestType.(*GeoRequest_Coords); ok {
			return x.Coords
		}
	}
	return nil
}

type isGeoRequest_RequestType interface {
	isGeoRequest_RequestType()
}

type GeoRequest_Query struct {
	Query string `protobuf:"bytes,1,opt,name=query,proto3,oneof"` // Address query for GetAddress
}

type GeoRequest_Coords struct {
	Coords *Coordinates `protobuf:"bytes,2,opt,name=coords,proto3,oneof"` // Coordinates for GetGeocode
}

func (*GeoRequest_Query) isGeoRequest_RequestType() {}

func (*GeoRequest_Coords) isGeoRequest_RequestType() {}

// Represents geographic coordinates.
type Coordinates struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Latitude      string                 `protobuf:"bytes,1,opt,name=latitude,proto3" json:"latitude,omitempty"`   // Latitude value
	Longitude     string                 `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"` // Longitude value
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	mi := &file_gomicro_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_gomicro_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinates.ProtoReflect.Descriptor instead.
func (*Coordinates) Descriptor() ([]byte, []int) {
	return file_gomicro_proto_rawDescGZIP(), []int{1}
}

func (x *Coordinates) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *Coordinates) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	City          string                 `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`     // Corresponds to the City field in Go
	Street        string                 `protobuf:"bytes,2,opt,name=street,proto3" json:"street,omitempty"` // Corresponds to the Street field in Go
	House         string                 `protobuf:"bytes,3,opt,name=house,proto3" json:"house,omitempty"`   // Corresponds to the House field in Go
	Lat           string                 `protobuf:"bytes,4,opt,name=lat,proto3" json:"lat,omitempty"`       // Corresponds to the Lat field in Go
	Lon           string                 `protobuf:"bytes,5,opt,name=lon,proto3" json:"lon,omitempty"`       // Corresponds to the Lon field in Go
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_gomicro_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_gomicro_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_gomicro_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetHouse() string {
	if x != nil {
		return x.House
	}
	return ""
}

func (x *Address) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Address) GetLon() string {
	if x != nil {
		return x.Lon
	}
	return ""
}

type GeoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addresses     []*Address             `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"` // Address list
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeoResponse) Reset() {
	*x = GeoResponse{}
	mi := &file_gomicro_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoResponse) ProtoMessage() {}

func (x *GeoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gomicro_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoResponse.ProtoReflect.Descriptor instead.
func (*GeoResponse) Descriptor() ([]byte, []int) {
	return file_gomicro_proto_rawDescGZIP(), []int{3}
}

func (x *GeoResponse) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type AuthRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`       // Email to register/authenticate
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // Password to register/authenticate
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	mi := &file_gomicro_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomicro_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_gomicro_proto_rawDescGZIP(), []int{4}
}

func (x *AuthRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`  // Result code of auth operation
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"` // JWT token encrypted
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	mi := &file_gomicro_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gomicro_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_gomicro_proto_rawDescGZIP(), []int{5}
}

func (x *AuthResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AuthResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_gomicro_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomicro_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_gomicro_proto_rawDescGZIP(), []int{6}
}

func (x *UserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *UserData              `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_gomicro_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gomicro_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_gomicro_proto_rawDescGZIP(), []int{7}
}

func (x *UserResponse) GetUser() *UserData {
	if x != nil {
		return x.User
	}
	return nil
}

type ListUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*UserData            `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserResponse) Reset() {
	*x = ListUserResponse{}
	mi := &file_gomicro_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserResponse) ProtoMessage() {}

func (x *ListUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gomicro_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserResponse.ProtoReflect.Descriptor instead.
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return file_gomicro_proto_rawDescGZIP(), []int{8}
}

func (x *ListUserResponse) GetList() []*UserData {
	if x != nil {
		return x.List
	}
	return nil
}

type UserData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Registered    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=registered,proto3" json:"registered,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserData) Reset() {
	*x = UserData{}
	mi := &file_gomicro_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_gomicro_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_gomicro_proto_rawDescGZIP(), []int{9}
}

func (x *UserData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserData) GetRegistered() *timestamppb.Timestamp {
	if x != nil {
		return x.Registered
	}
	return nil
}

var File_gomicro_proto protoreflect.FileDescriptor

var file_gomicro_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x0a, 0x47, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x06,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x73, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x0e, 0x0a, 0x0c,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x0b,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x6f, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x0b, 0x47, 0x65, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x38, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x1d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x35, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3a, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x32, 0x7b, 0x0a, 0x03, 0x47, 0x65, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x13, 0x2e, 0x67, 0x6f, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x47, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x47, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6f,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x47,
	0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x47, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0x79, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x2e,
	0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x7d, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x14, 0x2e, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19,
	0x2e, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x61, 0x73, 0x73, 0x65, 0x72,
	0x6d, 0x61, 0x6e, 0x6e, 0x31, 0x2f, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_gomicro_proto_rawDescOnce sync.Once
	file_gomicro_proto_rawDescData []byte
)

func file_gomicro_proto_rawDescGZIP() []byte {
	file_gomicro_proto_rawDescOnce.Do(func() {
		file_gomicro_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_gomicro_proto_rawDesc), len(file_gomicro_proto_rawDesc)))
	})
	return file_gomicro_proto_rawDescData
}

var file_gomicro_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gomicro_proto_goTypes = []any{
	(*GeoRequest)(nil),            // 0: gomicro.GeoRequest
	(*Coordinates)(nil),           // 1: gomicro.Coordinates
	(*Address)(nil),               // 2: gomicro.Address
	(*GeoResponse)(nil),           // 3: gomicro.GeoResponse
	(*AuthRequest)(nil),           // 4: gomicro.AuthRequest
	(*AuthResponse)(nil),          // 5: gomicro.AuthResponse
	(*UserRequest)(nil),           // 6: gomicro.UserRequest
	(*UserResponse)(nil),          // 7: gomicro.UserResponse
	(*ListUserResponse)(nil),      // 8: gomicro.ListUserResponse
	(*UserData)(nil),              // 9: gomicro.UserData
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_gomicro_proto_depIdxs = []int32{
	1,  // 0: gomicro.GeoRequest.coords:type_name -> gomicro.Coordinates
	2,  // 1: gomicro.GeoResponse.addresses:type_name -> gomicro.Address
	9,  // 2: gomicro.UserResponse.user:type_name -> gomicro.UserData
	9,  // 3: gomicro.ListUserResponse.list:type_name -> gomicro.UserData
	10, // 4: gomicro.UserData.registered:type_name -> google.protobuf.Timestamp
	0,  // 5: gomicro.Geo.GetAddress:input_type -> gomicro.GeoRequest
	0,  // 6: gomicro.Geo.GetGeocode:input_type -> gomicro.GeoRequest
	4,  // 7: gomicro.Auth.Register:input_type -> gomicro.AuthRequest
	4,  // 8: gomicro.Auth.Login:input_type -> gomicro.AuthRequest
	6,  // 9: gomicro.User.Profile:input_type -> gomicro.UserRequest
	11, // 10: gomicro.User.List:input_type -> google.protobuf.Empty
	3,  // 11: gomicro.Geo.GetAddress:output_type -> gomicro.GeoResponse
	3,  // 12: gomicro.Geo.GetGeocode:output_type -> gomicro.GeoResponse
	5,  // 13: gomicro.Auth.Register:output_type -> gomicro.AuthResponse
	5,  // 14: gomicro.Auth.Login:output_type -> gomicro.AuthResponse
	7,  // 15: gomicro.User.Profile:output_type -> gomicro.UserResponse
	8,  // 16: gomicro.User.List:output_type -> gomicro.ListUserResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_gomicro_proto_init() }
func file_gomicro_proto_init() {
	if File_gomicro_proto != nil {
		return
	}
	file_gomicro_proto_msgTypes[0].OneofWrappers = []any{
		(*GeoRequest_Query)(nil),
		(*GeoRequest_Coords)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_gomicro_proto_rawDesc), len(file_gomicro_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_gomicro_proto_goTypes,
		DependencyIndexes: file_gomicro_proto_depIdxs,
		MessageInfos:      file_gomicro_proto_msgTypes,
	}.Build()
	File_gomicro_proto = out.File
	file_gomicro_proto_goTypes = nil
	file_gomicro_proto_depIdxs = nil
}
