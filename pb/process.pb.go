// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.2
// source: pb/process.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProcessConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Command          string                              `protobuf:"bytes,2,opt,name=command,proto3" json:"command"`
	Arguments        []string                            `protobuf:"bytes,3,rep,name=arguments,proto3" json:"arguments"`
	Environment      map[string]string                   `protobuf:"bytes,4,rep,name=environment,proto3" json:"environment" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Persist          bool                                `protobuf:"varint,5,opt,name=persist,proto3" json:"persist"`
	RestartPolicy    *ProcessConfiguration_RestartPolicy `protobuf:"bytes,6,opt,name=restart_policy,json=restartPolicy,proto3" json:"restart_policy"`
	WorkingDirectory string                              `protobuf:"bytes,7,opt,name=working_directory,json=workingDirectory,proto3" json:"working_directory"`
}

func (x *ProcessConfiguration) Reset() {
	*x = ProcessConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_process_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessConfiguration) ProtoMessage() {}

func (x *ProcessConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pb_process_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessConfiguration.ProtoReflect.Descriptor instead.
func (*ProcessConfiguration) Descriptor() ([]byte, []int) {
	return file_pb_process_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProcessConfiguration) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ProcessConfiguration) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *ProcessConfiguration) GetEnvironment() map[string]string {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *ProcessConfiguration) GetPersist() bool {
	if x != nil {
		return x.Persist
	}
	return false
}

func (x *ProcessConfiguration) GetRestartPolicy() *ProcessConfiguration_RestartPolicy {
	if x != nil {
		return x.RestartPolicy
	}
	return nil
}

func (x *ProcessConfiguration) GetWorkingDirectory() string {
	if x != nil {
		return x.WorkingDirectory
	}
	return ""
}

type ProcessInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Pid           int64                         `protobuf:"varint,2,opt,name=pid,proto3" json:"pid"`
	Configuration *ProcessConfiguration         `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration"`
	ExitState     *ProcessInformation_ExitState `protobuf:"bytes,4,opt,name=exit_state,json=exitState,proto3,oneof" json:"exit_state"`
	Status        string                        `protobuf:"bytes,5,opt,name=status,proto3" json:"status"`
	Stdout        string                        `protobuf:"bytes,6,opt,name=stdout,proto3" json:"stdout"`
	Restarts      uint32                        `protobuf:"varint,7,opt,name=restarts,proto3" json:"restarts"`
	StartedAt     *timestamppb.Timestamp        `protobuf:"bytes,8,opt,name=started_at,json=startedAt,proto3,oneof" json:"started_at"`
	StoppedAt     *timestamppb.Timestamp        `protobuf:"bytes,9,opt,name=stopped_at,json=stoppedAt,proto3,oneof" json:"stopped_at"`
}

func (x *ProcessInformation) Reset() {
	*x = ProcessInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_process_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessInformation) ProtoMessage() {}

func (x *ProcessInformation) ProtoReflect() protoreflect.Message {
	mi := &file_pb_process_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessInformation.ProtoReflect.Descriptor instead.
func (*ProcessInformation) Descriptor() ([]byte, []int) {
	return file_pb_process_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessInformation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProcessInformation) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ProcessInformation) GetConfiguration() *ProcessConfiguration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *ProcessInformation) GetExitState() *ProcessInformation_ExitState {
	if x != nil {
		return x.ExitState
	}
	return nil
}

func (x *ProcessInformation) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ProcessInformation) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *ProcessInformation) GetRestarts() uint32 {
	if x != nil {
		return x.Restarts
	}
	return 0
}

func (x *ProcessInformation) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *ProcessInformation) GetStoppedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StoppedAt
	}
	return nil
}

type ProcessConfiguration_RestartPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutoRestart bool                 `protobuf:"varint,1,opt,name=auto_restart,json=autoRestart,proto3" json:"auto_restart"`
	Delay       *durationpb.Duration `protobuf:"bytes,2,opt,name=delay,proto3" json:"delay"`
	MaxRetries  uint32               `protobuf:"varint,3,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries"`
}

func (x *ProcessConfiguration_RestartPolicy) Reset() {
	*x = ProcessConfiguration_RestartPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_process_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessConfiguration_RestartPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessConfiguration_RestartPolicy) ProtoMessage() {}

func (x *ProcessConfiguration_RestartPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_pb_process_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessConfiguration_RestartPolicy.ProtoReflect.Descriptor instead.
func (*ProcessConfiguration_RestartPolicy) Descriptor() ([]byte, []int) {
	return file_pb_process_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProcessConfiguration_RestartPolicy) GetAutoRestart() bool {
	if x != nil {
		return x.AutoRestart
	}
	return false
}

func (x *ProcessConfiguration_RestartPolicy) GetDelay() *durationpb.Duration {
	if x != nil {
		return x.Delay
	}
	return nil
}

func (x *ProcessConfiguration_RestartPolicy) GetMaxRetries() uint32 {
	if x != nil {
		return x.MaxRetries
	}
	return 0
}

type ProcessInformation_ExitState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
}

func (x *ProcessInformation_ExitState) Reset() {
	*x = ProcessInformation_ExitState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_process_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessInformation_ExitState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessInformation_ExitState) ProtoMessage() {}

func (x *ProcessInformation_ExitState) ProtoReflect() protoreflect.Message {
	mi := &file_pb_process_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessInformation_ExitState.ProtoReflect.Descriptor instead.
func (*ProcessInformation_ExitState) Descriptor() ([]byte, []int) {
	return file_pb_process_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ProcessInformation_ExitState) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_pb_process_proto protoreflect.FileDescriptor

var file_pb_process_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x04, 0x0a,
	0x14, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x50, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x12, 0x52, 0x0a,
	0x0e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x2b, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x84,
	0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe0, 0x03, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x43,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0a, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00,
	0x52, 0x09, 0x65, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74,
	0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x73, 0x74,
	0x6f, 0x70, 0x70, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x1a, 0x1f, 0x0a, 0x09, 0x45, 0x78,
	0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x65, 0x78, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74,
	0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x65, 0x78, 0x6e, 0x7a, 0x61, 0x72, 0x6f,
	0x76, 0x2f, 0x67, 0x6f, 0x66, 0x75, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_process_proto_rawDescOnce sync.Once
	file_pb_process_proto_rawDescData = file_pb_process_proto_rawDesc
)

func file_pb_process_proto_rawDescGZIP() []byte {
	file_pb_process_proto_rawDescOnce.Do(func() {
		file_pb_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_process_proto_rawDescData)
	})
	return file_pb_process_proto_rawDescData
}

var file_pb_process_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_process_proto_goTypes = []interface{}{
	(*ProcessConfiguration)(nil),               // 0: process.ProcessConfiguration
	(*ProcessInformation)(nil),                 // 1: process.ProcessInformation
	(*ProcessConfiguration_RestartPolicy)(nil), // 2: process.ProcessConfiguration.RestartPolicy
	nil,                                  // 3: process.ProcessConfiguration.EnvironmentEntry
	(*ProcessInformation_ExitState)(nil), // 4: process.ProcessInformation.ExitState
	(*timestamppb.Timestamp)(nil),        // 5: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),          // 6: google.protobuf.Duration
}
var file_pb_process_proto_depIdxs = []int32{
	3, // 0: process.ProcessConfiguration.environment:type_name -> process.ProcessConfiguration.EnvironmentEntry
	2, // 1: process.ProcessConfiguration.restart_policy:type_name -> process.ProcessConfiguration.RestartPolicy
	0, // 2: process.ProcessInformation.configuration:type_name -> process.ProcessConfiguration
	4, // 3: process.ProcessInformation.exit_state:type_name -> process.ProcessInformation.ExitState
	5, // 4: process.ProcessInformation.started_at:type_name -> google.protobuf.Timestamp
	5, // 5: process.ProcessInformation.stopped_at:type_name -> google.protobuf.Timestamp
	6, // 6: process.ProcessConfiguration.RestartPolicy.delay:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pb_process_proto_init() }
func file_pb_process_proto_init() {
	if File_pb_process_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_process_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessConfiguration); i {
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
		file_pb_process_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessInformation); i {
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
		file_pb_process_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessConfiguration_RestartPolicy); i {
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
		file_pb_process_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessInformation_ExitState); i {
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
	file_pb_process_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_process_proto_goTypes,
		DependencyIndexes: file_pb_process_proto_depIdxs,
		MessageInfos:      file_pb_process_proto_msgTypes,
	}.Build()
	File_pb_process_proto = out.File
	file_pb_process_proto_rawDesc = nil
	file_pb_process_proto_goTypes = nil
	file_pb_process_proto_depIdxs = nil
}
