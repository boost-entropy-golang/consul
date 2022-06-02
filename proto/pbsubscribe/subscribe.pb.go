//
//Package event provides a service for subscribing to state change events.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-rc.1
// 	protoc        (unknown)
// source: proto/pbsubscribe/subscribe.proto

package pbsubscribe

import (
	pbservice "github.com/hashicorp/consul/proto/pbservice"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Topic enumerates the supported event topics.
type Topic int32

const (
	Topic_Unknown Topic = 0
	// ServiceHealth topic contains events for any changes to service health.
	Topic_ServiceHealth Topic = 1
	// ServiceHealthConnect topic contains events for any changes to service
	// health for connect-enabled services.
	Topic_ServiceHealthConnect Topic = 2
)

// Enum value maps for Topic.
var (
	Topic_name = map[int32]string{
		0: "Unknown",
		1: "ServiceHealth",
		2: "ServiceHealthConnect",
	}
	Topic_value = map[string]int32{
		"Unknown":              0,
		"ServiceHealth":        1,
		"ServiceHealthConnect": 2,
	}
)

func (x Topic) Enum() *Topic {
	p := new(Topic)
	*p = x
	return p
}

func (x Topic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Topic) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_pbsubscribe_subscribe_proto_enumTypes[0].Descriptor()
}

func (Topic) Type() protoreflect.EnumType {
	return &file_proto_pbsubscribe_subscribe_proto_enumTypes[0]
}

func (x Topic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Topic.Descriptor instead.
func (Topic) EnumDescriptor() ([]byte, []int) {
	return file_proto_pbsubscribe_subscribe_proto_rawDescGZIP(), []int{0}
}

type CatalogOp int32

const (
	CatalogOp_Register   CatalogOp = 0
	CatalogOp_Deregister CatalogOp = 1
)

// Enum value maps for CatalogOp.
var (
	CatalogOp_name = map[int32]string{
		0: "Register",
		1: "Deregister",
	}
	CatalogOp_value = map[string]int32{
		"Register":   0,
		"Deregister": 1,
	}
)

func (x CatalogOp) Enum() *CatalogOp {
	p := new(CatalogOp)
	*p = x
	return p
}

func (x CatalogOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CatalogOp) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_pbsubscribe_subscribe_proto_enumTypes[1].Descriptor()
}

func (CatalogOp) Type() protoreflect.EnumType {
	return &file_proto_pbsubscribe_subscribe_proto_enumTypes[1]
}

func (x CatalogOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CatalogOp.Descriptor instead.
func (CatalogOp) EnumDescriptor() ([]byte, []int) {
	return file_proto_pbsubscribe_subscribe_proto_rawDescGZIP(), []int{1}
}

// SubscribeRequest used to subscribe to a topic.
type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Topic identifies the set of events the subscriber is interested in.
	Topic Topic `protobuf:"varint,1,opt,name=Topic,proto3,enum=subscribe.Topic" json:"Topic,omitempty"`
	// Key is a topic-specific identifier that restricts the scope of the
	// subscription to only events pertaining to that identifier. For example,
	// to receive events for a single service, the service's name is specified
	// as the key.
	Key string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	// Token is the ACL token to authenticate the request. The token must have
	// sufficient privileges to read the requested information otherwise events
	// will be filtered, possibly resulting in an empty snapshot and no further
	// updates sent.
	Token string `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	// Index is the raft index the subscriber has already observed up to. This
	// is zero on an initial streaming call, but then can be provided by a
	// client on subsequent re-connections such that the full snapshot doesn't
	// need to be resent if the client is up to date.
	Index uint64 `protobuf:"varint,4,opt,name=Index,proto3" json:"Index,omitempty"`
	// Datacenter specifies the Consul datacenter the request is targeted at.
	// If it's not the local DC the server will forward the request to
	// the remote DC and proxy the results back  to the subscriber. An empty
	// string defaults to the local datacenter.
	Datacenter string `protobuf:"bytes,5,opt,name=Datacenter,proto3" json:"Datacenter,omitempty"`
	// Namespace which contains the resources. If Namespace is not specified the
	// default namespace will be used.
	//
	// Namespace is an enterprise-only feature.
	Namespace string `protobuf:"bytes,6,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	// Partition which contains the resources. If Partition is not specified the
	// default partition will be used.
	//
	// Partition is an enterprise-only feature.
	Partition string `protobuf:"bytes,7,opt,name=Partition,proto3" json:"Partition,omitempty"`
	// PeerName is the name of the peer that the requested service was imported from.
	PeerName string `protobuf:"bytes,8,opt,name=PeerName,proto3" json:"PeerName,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pbsubscribe_subscribe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbsubscribe_subscribe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbsubscribe_subscribe_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetTopic() Topic {
	if x != nil {
		return x.Topic
	}
	return Topic_Unknown
}

func (x *SubscribeRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SubscribeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SubscribeRequest) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SubscribeRequest) GetDatacenter() string {
	if x != nil {
		return x.Datacenter
	}
	return ""
}

func (x *SubscribeRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SubscribeRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *SubscribeRequest) GetPeerName() string {
	if x != nil {
		return x.PeerName
	}
	return ""
}

// Event describes a streaming update on a subscription. Events are used both to
// describe the current "snapshot" of the result as well as ongoing mutations to
// that snapshot.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index is the raft index at which the mutation took place. At the top
	// level of a subscription there will always be at most one Event per index.
	// If multiple events are published to the same topic in a single raft
	// transaction then the batch of events will be encoded inside a single
	// top-level event to ensure they are delivered atomically to clients.
	Index uint64 `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	// Payload is the actual event content.
	//
	// Types that are assignable to Payload:
	//	*Event_EndOfSnapshot
	//	*Event_NewSnapshotToFollow
	//	*Event_EventBatch
	//	*Event_ServiceHealth
	Payload isEvent_Payload `protobuf_oneof:"Payload"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pbsubscribe_subscribe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbsubscribe_subscribe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_proto_pbsubscribe_subscribe_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Event) GetEndOfSnapshot() bool {
	if x, ok := x.GetPayload().(*Event_EndOfSnapshot); ok {
		return x.EndOfSnapshot
	}
	return false
}

func (x *Event) GetNewSnapshotToFollow() bool {
	if x, ok := x.GetPayload().(*Event_NewSnapshotToFollow); ok {
		return x.NewSnapshotToFollow
	}
	return false
}

func (x *Event) GetEventBatch() *EventBatch {
	if x, ok := x.GetPayload().(*Event_EventBatch); ok {
		return x.EventBatch
	}
	return nil
}

func (x *Event) GetServiceHealth() *ServiceHealthUpdate {
	if x, ok := x.GetPayload().(*Event_ServiceHealth); ok {
		return x.ServiceHealth
	}
	return nil
}

type isEvent_Payload interface {
	isEvent_Payload()
}

type Event_EndOfSnapshot struct {
	// EndOfSnapshot indicates the event stream for the initial snapshot has
	// ended. Subsequent Events delivered will be mutations to that result.
	EndOfSnapshot bool `protobuf:"varint,2,opt,name=EndOfSnapshot,proto3,oneof"`
}

type Event_NewSnapshotToFollow struct {
	// NewSnapshotToFollow indicates that the client view is stale. The client
	// must reset its view before handing any more events. Subsequent events
	// in the stream will be for a new snapshot until an EndOfSnapshot event
	// is received.
	NewSnapshotToFollow bool `protobuf:"varint,3,opt,name=NewSnapshotToFollow,proto3,oneof"`
}

type Event_EventBatch struct {
	// EventBatch is a set of events. This is typically used as the payload
	// type where multiple events are emitted in a single topic and raft
	// index (e.g. transactional updates). In this case the Topic and Index
	// values of all events will match and the whole set should be delivered
	// and consumed atomically.
	EventBatch *EventBatch `protobuf:"bytes,4,opt,name=EventBatch,proto3,oneof"`
}

type Event_ServiceHealth struct {
	// ServiceHealth is used for ServiceHealth and ServiceHealthConnect
	// topics.
	ServiceHealth *ServiceHealthUpdate `protobuf:"bytes,10,opt,name=ServiceHealth,proto3,oneof"`
}

func (*Event_EndOfSnapshot) isEvent_Payload() {}

func (*Event_NewSnapshotToFollow) isEvent_Payload() {}

func (*Event_EventBatch) isEvent_Payload() {}

func (*Event_ServiceHealth) isEvent_Payload() {}

type EventBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty"`
}

func (x *EventBatch) Reset() {
	*x = EventBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pbsubscribe_subscribe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventBatch) ProtoMessage() {}

func (x *EventBatch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbsubscribe_subscribe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventBatch.ProtoReflect.Descriptor instead.
func (*EventBatch) Descriptor() ([]byte, []int) {
	return file_proto_pbsubscribe_subscribe_proto_rawDescGZIP(), []int{2}
}

func (x *EventBatch) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type ServiceHealthUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op               CatalogOp                   `protobuf:"varint,1,opt,name=Op,proto3,enum=subscribe.CatalogOp" json:"Op,omitempty"`
	CheckServiceNode *pbservice.CheckServiceNode `protobuf:"bytes,2,opt,name=CheckServiceNode,proto3" json:"CheckServiceNode,omitempty"`
}

func (x *ServiceHealthUpdate) Reset() {
	*x = ServiceHealthUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pbsubscribe_subscribe_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceHealthUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceHealthUpdate) ProtoMessage() {}

func (x *ServiceHealthUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbsubscribe_subscribe_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceHealthUpdate.ProtoReflect.Descriptor instead.
func (*ServiceHealthUpdate) Descriptor() ([]byte, []int) {
	return file_proto_pbsubscribe_subscribe_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceHealthUpdate) GetOp() CatalogOp {
	if x != nil {
		return x.Op
	}
	return CatalogOp_Register
}

func (x *ServiceHealthUpdate) GetCheckServiceNode() *pbservice.CheckServiceNode {
	if x != nil {
		return x.CheckServiceNode
	}
	return nil
}

var File_proto_pbsubscribe_subscribe_proto protoreflect.FileDescriptor

var file_proto_pbsubscribe_subscribe_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x1a, 0x1a,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x10, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x65, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x85, 0x02,
	0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a,
	0x0d, 0x45, 0x6e, 0x64, 0x4f, 0x66, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0d, 0x45, 0x6e, 0x64, 0x4f, 0x66, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x32, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x54, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x13, 0x4e, 0x65, 0x77, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x54, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x37, 0x0a, 0x0a, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x48, 0x00, 0x52, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x46, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x36, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x28, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x82, 0x01,
	0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x4f, 0x70, 0x52, 0x02, 0x4f, 0x70, 0x12, 0x45, 0x0a, 0x10, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x2a, 0x41, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x10, 0x02, 0x2a, 0x29, 0x0a, 0x09, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x4f, 0x70, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x10, 0x01,
	0x32, 0x59, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x92, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x0e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x62, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0xa2, 0x02,
	0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0xca, 0x02, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0xe2, 0x02, 0x15, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pbsubscribe_subscribe_proto_rawDescOnce sync.Once
	file_proto_pbsubscribe_subscribe_proto_rawDescData = file_proto_pbsubscribe_subscribe_proto_rawDesc
)

func file_proto_pbsubscribe_subscribe_proto_rawDescGZIP() []byte {
	file_proto_pbsubscribe_subscribe_proto_rawDescOnce.Do(func() {
		file_proto_pbsubscribe_subscribe_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pbsubscribe_subscribe_proto_rawDescData)
	})
	return file_proto_pbsubscribe_subscribe_proto_rawDescData
}

var file_proto_pbsubscribe_subscribe_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_pbsubscribe_subscribe_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_pbsubscribe_subscribe_proto_goTypes = []interface{}{
	(Topic)(0),                         // 0: subscribe.Topic
	(CatalogOp)(0),                     // 1: subscribe.CatalogOp
	(*SubscribeRequest)(nil),           // 2: subscribe.SubscribeRequest
	(*Event)(nil),                      // 3: subscribe.Event
	(*EventBatch)(nil),                 // 4: subscribe.EventBatch
	(*ServiceHealthUpdate)(nil),        // 5: subscribe.ServiceHealthUpdate
	(*pbservice.CheckServiceNode)(nil), // 6: service.CheckServiceNode
}
var file_proto_pbsubscribe_subscribe_proto_depIdxs = []int32{
	0, // 0: subscribe.SubscribeRequest.Topic:type_name -> subscribe.Topic
	4, // 1: subscribe.Event.EventBatch:type_name -> subscribe.EventBatch
	5, // 2: subscribe.Event.ServiceHealth:type_name -> subscribe.ServiceHealthUpdate
	3, // 3: subscribe.EventBatch.Events:type_name -> subscribe.Event
	1, // 4: subscribe.ServiceHealthUpdate.Op:type_name -> subscribe.CatalogOp
	6, // 5: subscribe.ServiceHealthUpdate.CheckServiceNode:type_name -> service.CheckServiceNode
	2, // 6: subscribe.StateChangeSubscription.Subscribe:input_type -> subscribe.SubscribeRequest
	3, // 7: subscribe.StateChangeSubscription.Subscribe:output_type -> subscribe.Event
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_pbsubscribe_subscribe_proto_init() }
func file_proto_pbsubscribe_subscribe_proto_init() {
	if File_proto_pbsubscribe_subscribe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pbsubscribe_subscribe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_pbsubscribe_subscribe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_pbsubscribe_subscribe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventBatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_pbsubscribe_subscribe_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceHealthUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_pbsubscribe_subscribe_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Event_EndOfSnapshot)(nil),
		(*Event_NewSnapshotToFollow)(nil),
		(*Event_EventBatch)(nil),
		(*Event_ServiceHealth)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_pbsubscribe_subscribe_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pbsubscribe_subscribe_proto_goTypes,
		DependencyIndexes: file_proto_pbsubscribe_subscribe_proto_depIdxs,
		EnumInfos:         file_proto_pbsubscribe_subscribe_proto_enumTypes,
		MessageInfos:      file_proto_pbsubscribe_subscribe_proto_msgTypes,
	}.Build()
	File_proto_pbsubscribe_subscribe_proto = out.File
	file_proto_pbsubscribe_subscribe_proto_rawDesc = nil
	file_proto_pbsubscribe_subscribe_proto_goTypes = nil
	file_proto_pbsubscribe_subscribe_proto_depIdxs = nil
}
