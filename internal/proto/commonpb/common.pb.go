// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package commonpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrorCode int32

const (
	ErrorCode_Success               ErrorCode = 0
	ErrorCode_UnexpectedError       ErrorCode = 1
	ErrorCode_ConnectFailed         ErrorCode = 2
	ErrorCode_PermissionDenied      ErrorCode = 3
	ErrorCode_CollectionNotExists   ErrorCode = 4
	ErrorCode_IllegalArgument       ErrorCode = 5
	ErrorCode_IllegalDimension      ErrorCode = 7
	ErrorCode_IllegalIndexType      ErrorCode = 8
	ErrorCode_IllegalCollectionName ErrorCode = 9
	ErrorCode_IllegalTOPK           ErrorCode = 10
	ErrorCode_IllegalRowRecord      ErrorCode = 11
	ErrorCode_IllegalVectorID       ErrorCode = 12
	ErrorCode_IllegalSearchResult   ErrorCode = 13
	ErrorCode_FileNotFound          ErrorCode = 14
	ErrorCode_MetaFailed            ErrorCode = 15
	ErrorCode_CacheFailed           ErrorCode = 16
	ErrorCode_CannotCreateFolder    ErrorCode = 17
	ErrorCode_CannotCreateFile      ErrorCode = 18
	ErrorCode_CannotDeleteFolder    ErrorCode = 19
	ErrorCode_CannotDeleteFile      ErrorCode = 20
	ErrorCode_BuildIndexError       ErrorCode = 21
	ErrorCode_IllegalNLIST          ErrorCode = 22
	ErrorCode_IllegalMetricType     ErrorCode = 23
	ErrorCode_OutOfMemory           ErrorCode = 24
	ErrorCode_IndexNotExist         ErrorCode = 25
	ErrorCode_EmptyCollection       ErrorCode = 26
	// internal error code.
	ErrorCode_DDRequestRace ErrorCode = 1000
)

var ErrorCode_name = map[int32]string{
	0:    "Success",
	1:    "UnexpectedError",
	2:    "ConnectFailed",
	3:    "PermissionDenied",
	4:    "CollectionNotExists",
	5:    "IllegalArgument",
	7:    "IllegalDimension",
	8:    "IllegalIndexType",
	9:    "IllegalCollectionName",
	10:   "IllegalTOPK",
	11:   "IllegalRowRecord",
	12:   "IllegalVectorID",
	13:   "IllegalSearchResult",
	14:   "FileNotFound",
	15:   "MetaFailed",
	16:   "CacheFailed",
	17:   "CannotCreateFolder",
	18:   "CannotCreateFile",
	19:   "CannotDeleteFolder",
	20:   "CannotDeleteFile",
	21:   "BuildIndexError",
	22:   "IllegalNLIST",
	23:   "IllegalMetricType",
	24:   "OutOfMemory",
	25:   "IndexNotExist",
	26:   "EmptyCollection",
	1000: "DDRequestRace",
}

var ErrorCode_value = map[string]int32{
	"Success":               0,
	"UnexpectedError":       1,
	"ConnectFailed":         2,
	"PermissionDenied":      3,
	"CollectionNotExists":   4,
	"IllegalArgument":       5,
	"IllegalDimension":      7,
	"IllegalIndexType":      8,
	"IllegalCollectionName": 9,
	"IllegalTOPK":           10,
	"IllegalRowRecord":      11,
	"IllegalVectorID":       12,
	"IllegalSearchResult":   13,
	"FileNotFound":          14,
	"MetaFailed":            15,
	"CacheFailed":           16,
	"CannotCreateFolder":    17,
	"CannotCreateFile":      18,
	"CannotDeleteFolder":    19,
	"CannotDeleteFile":      20,
	"BuildIndexError":       21,
	"IllegalNLIST":          22,
	"IllegalMetricType":     23,
	"OutOfMemory":           24,
	"IndexNotExist":         25,
	"EmptyCollection":       26,
	"DDRequestRace":         1000,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type IndexState int32

const (
	IndexState_IndexStateNone IndexState = 0
	IndexState_Unissued       IndexState = 1
	IndexState_InProgress     IndexState = 2
	IndexState_Finished       IndexState = 3
	IndexState_Failed         IndexState = 4
)

var IndexState_name = map[int32]string{
	0: "IndexStateNone",
	1: "Unissued",
	2: "InProgress",
	3: "Finished",
	4: "Failed",
}

var IndexState_value = map[string]int32{
	"IndexStateNone": 0,
	"Unissued":       1,
	"InProgress":     2,
	"Finished":       3,
	"Failed":         4,
}

func (x IndexState) String() string {
	return proto.EnumName(IndexState_name, int32(x))
}

func (IndexState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type SegmentState int32

const (
	SegmentState_SegmentStateNone SegmentState = 0
	SegmentState_NotExist         SegmentState = 1
	SegmentState_Growing          SegmentState = 2
	SegmentState_Sealed           SegmentState = 3
	SegmentState_Flushed          SegmentState = 4
	SegmentState_Flushing         SegmentState = 5
	SegmentState_Dropped          SegmentState = 6
)

var SegmentState_name = map[int32]string{
	0: "SegmentStateNone",
	1: "NotExist",
	2: "Growing",
	3: "Sealed",
	4: "Flushed",
	5: "Flushing",
	6: "Dropped",
}

var SegmentState_value = map[string]int32{
	"SegmentStateNone": 0,
	"NotExist":         1,
	"Growing":          2,
	"Sealed":           3,
	"Flushed":          4,
	"Flushing":         5,
	"Dropped":          6,
}

func (x SegmentState) String() string {
	return proto.EnumName(SegmentState_name, int32(x))
}

func (SegmentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

type MsgType int32

const (
	MsgType_Undefined MsgType = 0
	// DEFINITION REQUESTS: COLLECTION
	MsgType_CreateCollection   MsgType = 100
	MsgType_DropCollection     MsgType = 101
	MsgType_HasCollection      MsgType = 102
	MsgType_DescribeCollection MsgType = 103
	MsgType_ShowCollections    MsgType = 104
	MsgType_GetSystemConfigs   MsgType = 105
	MsgType_LoadCollection     MsgType = 106
	MsgType_ReleaseCollection  MsgType = 107
	MsgType_CreateAlias        MsgType = 108
	MsgType_DropAlias          MsgType = 109
	MsgType_AlterAlias         MsgType = 110
	// DEFINITION REQUESTS: PARTITION
	MsgType_CreatePartition   MsgType = 200
	MsgType_DropPartition     MsgType = 201
	MsgType_HasPartition      MsgType = 202
	MsgType_DescribePartition MsgType = 203
	MsgType_ShowPartitions    MsgType = 204
	MsgType_LoadPartitions    MsgType = 205
	MsgType_ReleasePartitions MsgType = 206
	// DEFINE REQUESTS: SEGMENT
	MsgType_ShowSegments        MsgType = 250
	MsgType_DescribeSegment     MsgType = 251
	MsgType_LoadSegments        MsgType = 252
	MsgType_ReleaseSegments     MsgType = 253
	MsgType_HandoffSegments     MsgType = 254
	MsgType_LoadBalanceSegments MsgType = 255
	// DEFINITION REQUESTS: INDEX
	MsgType_CreateIndex   MsgType = 300
	MsgType_DescribeIndex MsgType = 301
	MsgType_DropIndex     MsgType = 302
	// MANIPULATION REQUESTS
	MsgType_Insert MsgType = 400
	MsgType_Delete MsgType = 401
	MsgType_Flush  MsgType = 402
	// QUERY
	MsgType_Search                   MsgType = 500
	MsgType_SearchResult             MsgType = 501
	MsgType_GetIndexState            MsgType = 502
	MsgType_GetIndexBuildProgress    MsgType = 503
	MsgType_GetCollectionStatistics  MsgType = 504
	MsgType_GetPartitionStatistics   MsgType = 505
	MsgType_Retrieve                 MsgType = 506
	MsgType_RetrieveResult           MsgType = 507
	MsgType_WatchDmChannels          MsgType = 508
	MsgType_RemoveDmChannels         MsgType = 509
	MsgType_WatchQueryChannels       MsgType = 510
	MsgType_RemoveQueryChannels      MsgType = 511
	MsgType_SealedSegmentsChangeInfo MsgType = 512
	MsgType_WatchDeltaChannels       MsgType = 513
	// DATA SERVICE
	MsgType_SegmentInfo     MsgType = 600
	MsgType_SystemInfo      MsgType = 601
	MsgType_GetRecoveryInfo MsgType = 602
	MsgType_GetSegmentState MsgType = 603
	// SYSTEM CONTROL
	MsgType_TimeTick          MsgType = 1200
	MsgType_QueryNodeStats    MsgType = 1201
	MsgType_LoadIndex         MsgType = 1202
	MsgType_RequestID         MsgType = 1203
	MsgType_RequestTSO        MsgType = 1204
	MsgType_AllocateSegment   MsgType = 1205
	MsgType_SegmentStatistics MsgType = 1206
	MsgType_SegmentFlushDone  MsgType = 1207
	MsgType_DataNodeTt        MsgType = 1208
)

var MsgType_name = map[int32]string{
	0:    "Undefined",
	100:  "CreateCollection",
	101:  "DropCollection",
	102:  "HasCollection",
	103:  "DescribeCollection",
	104:  "ShowCollections",
	105:  "GetSystemConfigs",
	106:  "LoadCollection",
	107:  "ReleaseCollection",
	108:  "CreateAlias",
	109:  "DropAlias",
	110:  "AlterAlias",
	200:  "CreatePartition",
	201:  "DropPartition",
	202:  "HasPartition",
	203:  "DescribePartition",
	204:  "ShowPartitions",
	205:  "LoadPartitions",
	206:  "ReleasePartitions",
	250:  "ShowSegments",
	251:  "DescribeSegment",
	252:  "LoadSegments",
	253:  "ReleaseSegments",
	254:  "HandoffSegments",
	255:  "LoadBalanceSegments",
	300:  "CreateIndex",
	301:  "DescribeIndex",
	302:  "DropIndex",
	400:  "Insert",
	401:  "Delete",
	402:  "Flush",
	500:  "Search",
	501:  "SearchResult",
	502:  "GetIndexState",
	503:  "GetIndexBuildProgress",
	504:  "GetCollectionStatistics",
	505:  "GetPartitionStatistics",
	506:  "Retrieve",
	507:  "RetrieveResult",
	508:  "WatchDmChannels",
	509:  "RemoveDmChannels",
	510:  "WatchQueryChannels",
	511:  "RemoveQueryChannels",
	512:  "SealedSegmentsChangeInfo",
	513:  "WatchDeltaChannels",
	600:  "SegmentInfo",
	601:  "SystemInfo",
	602:  "GetRecoveryInfo",
	603:  "GetSegmentState",
	1200: "TimeTick",
	1201: "QueryNodeStats",
	1202: "LoadIndex",
	1203: "RequestID",
	1204: "RequestTSO",
	1205: "AllocateSegment",
	1206: "SegmentStatistics",
	1207: "SegmentFlushDone",
	1208: "DataNodeTt",
}

var MsgType_value = map[string]int32{
	"Undefined":                0,
	"CreateCollection":         100,
	"DropCollection":           101,
	"HasCollection":            102,
	"DescribeCollection":       103,
	"ShowCollections":          104,
	"GetSystemConfigs":         105,
	"LoadCollection":           106,
	"ReleaseCollection":        107,
	"CreateAlias":              108,
	"DropAlias":                109,
	"AlterAlias":               110,
	"CreatePartition":          200,
	"DropPartition":            201,
	"HasPartition":             202,
	"DescribePartition":        203,
	"ShowPartitions":           204,
	"LoadPartitions":           205,
	"ReleasePartitions":        206,
	"ShowSegments":             250,
	"DescribeSegment":          251,
	"LoadSegments":             252,
	"ReleaseSegments":          253,
	"HandoffSegments":          254,
	"LoadBalanceSegments":      255,
	"CreateIndex":              300,
	"DescribeIndex":            301,
	"DropIndex":                302,
	"Insert":                   400,
	"Delete":                   401,
	"Flush":                    402,
	"Search":                   500,
	"SearchResult":             501,
	"GetIndexState":            502,
	"GetIndexBuildProgress":    503,
	"GetCollectionStatistics":  504,
	"GetPartitionStatistics":   505,
	"Retrieve":                 506,
	"RetrieveResult":           507,
	"WatchDmChannels":          508,
	"RemoveDmChannels":         509,
	"WatchQueryChannels":       510,
	"RemoveQueryChannels":      511,
	"SealedSegmentsChangeInfo": 512,
	"WatchDeltaChannels":       513,
	"SegmentInfo":              600,
	"SystemInfo":               601,
	"GetRecoveryInfo":          602,
	"GetSegmentState":          603,
	"TimeTick":                 1200,
	"QueryNodeStats":           1201,
	"LoadIndex":                1202,
	"RequestID":                1203,
	"RequestTSO":               1204,
	"AllocateSegment":          1205,
	"SegmentStatistics":        1206,
	"SegmentFlushDone":         1207,
	"DataNodeTt":               1208,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

type DslType int32

const (
	DslType_Dsl        DslType = 0
	DslType_BoolExprV1 DslType = 1
)

var DslType_name = map[int32]string{
	0: "Dsl",
	1: "BoolExprV1",
}

var DslType_value = map[string]int32{
	"Dsl":        0,
	"BoolExprV1": 1,
}

func (x DslType) String() string {
	return proto.EnumName(DslType_name, int32(x))
}

func (DslType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

type CompactionState int32

const (
	CompactionState_UndefiedState CompactionState = 0
	CompactionState_Executing     CompactionState = 1
	CompactionState_Completed     CompactionState = 2
)

var CompactionState_name = map[int32]string{
	0: "UndefiedState",
	1: "Executing",
	2: "Completed",
}

var CompactionState_value = map[string]int32{
	"UndefiedState": 0,
	"Executing":     1,
	"Completed":     2,
}

func (x CompactionState) String() string {
	return proto.EnumName(CompactionState_name, int32(x))
}

func (CompactionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

type ConsistencyLevel int32

const (
	ConsistencyLevel_Strong     ConsistencyLevel = 0
	ConsistencyLevel_Session    ConsistencyLevel = 1
	ConsistencyLevel_Bounded    ConsistencyLevel = 2
	ConsistencyLevel_Eventually ConsistencyLevel = 3
	ConsistencyLevel_Customized ConsistencyLevel = 4
)

var ConsistencyLevel_name = map[int32]string{
	0: "Strong",
	1: "Session",
	2: "Bounded",
	3: "Eventually",
	4: "Customized",
}

var ConsistencyLevel_value = map[string]int32{
	"Strong":     0,
	"Session":    1,
	"Bounded":    2,
	"Eventually": 3,
	"Customized": 4,
}

func (x ConsistencyLevel) String() string {
	return proto.EnumName(ConsistencyLevel_name, int32(x))
}

func (ConsistencyLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

type Status struct {
	ErrorCode            ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=milvus.proto.common.ErrorCode" json:"error_code,omitempty"`
	Reason               string    `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetErrorCode() ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorCode_Success
}

func (m *Status) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type KeyValuePair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValuePair.Unmarshal(m, b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyValuePair.Size(m)
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KeyDataPair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyDataPair) Reset()         { *m = KeyDataPair{} }
func (m *KeyDataPair) String() string { return proto.CompactTextString(m) }
func (*KeyDataPair) ProtoMessage()    {}
func (*KeyDataPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *KeyDataPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyDataPair.Unmarshal(m, b)
}
func (m *KeyDataPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyDataPair.Marshal(b, m, deterministic)
}
func (m *KeyDataPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyDataPair.Merge(m, src)
}
func (m *KeyDataPair) XXX_Size() int {
	return xxx_messageInfo_KeyDataPair.Size(m)
}
func (m *KeyDataPair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyDataPair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyDataPair proto.InternalMessageInfo

func (m *KeyDataPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyDataPair) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Blob struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blob) Reset()         { *m = Blob{} }
func (m *Blob) String() string { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()    {}
func (*Blob) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *Blob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blob.Unmarshal(m, b)
}
func (m *Blob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blob.Marshal(b, m, deterministic)
}
func (m *Blob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blob.Merge(m, src)
}
func (m *Blob) XXX_Size() int {
	return xxx_messageInfo_Blob.Size(m)
}
func (m *Blob) XXX_DiscardUnknown() {
	xxx_messageInfo_Blob.DiscardUnknown(m)
}

var xxx_messageInfo_Blob proto.InternalMessageInfo

func (m *Blob) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Address struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int64    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Address) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type MsgBase struct {
	MsgType              MsgType  `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=milvus.proto.common.MsgType" json:"msg_type,omitempty"`
	MsgID                int64    `protobuf:"varint,2,opt,name=msgID,proto3" json:"msgID,omitempty"`
	Timestamp            uint64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SourceID             int64    `protobuf:"varint,4,opt,name=sourceID,proto3" json:"sourceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBase) Reset()         { *m = MsgBase{} }
func (m *MsgBase) String() string { return proto.CompactTextString(m) }
func (*MsgBase) ProtoMessage()    {}
func (*MsgBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *MsgBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBase.Unmarshal(m, b)
}
func (m *MsgBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBase.Marshal(b, m, deterministic)
}
func (m *MsgBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBase.Merge(m, src)
}
func (m *MsgBase) XXX_Size() int {
	return xxx_messageInfo_MsgBase.Size(m)
}
func (m *MsgBase) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBase.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBase proto.InternalMessageInfo

func (m *MsgBase) GetMsgType() MsgType {
	if m != nil {
		return m.MsgType
	}
	return MsgType_Undefined
}

func (m *MsgBase) GetMsgID() int64 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

func (m *MsgBase) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MsgBase) GetSourceID() int64 {
	if m != nil {
		return m.SourceID
	}
	return 0
}

// Don't Modify This. @czs
type MsgHeader struct {
	Base                 *MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgHeader) Reset()         { *m = MsgHeader{} }
func (m *MsgHeader) String() string { return proto.CompactTextString(m) }
func (*MsgHeader) ProtoMessage()    {}
func (*MsgHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

func (m *MsgHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgHeader.Unmarshal(m, b)
}
func (m *MsgHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgHeader.Marshal(b, m, deterministic)
}
func (m *MsgHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgHeader.Merge(m, src)
}
func (m *MsgHeader) XXX_Size() int {
	return xxx_messageInfo_MsgHeader.Size(m)
}
func (m *MsgHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MsgHeader proto.InternalMessageInfo

func (m *MsgHeader) GetBase() *MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

// Don't Modify This. @czs
type DMLMsgHeader struct {
	Base                 *MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	ShardName            string   `protobuf:"bytes,2,opt,name=shardName,proto3" json:"shardName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DMLMsgHeader) Reset()         { *m = DMLMsgHeader{} }
func (m *DMLMsgHeader) String() string { return proto.CompactTextString(m) }
func (*DMLMsgHeader) ProtoMessage()    {}
func (*DMLMsgHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{7}
}

func (m *DMLMsgHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DMLMsgHeader.Unmarshal(m, b)
}
func (m *DMLMsgHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DMLMsgHeader.Marshal(b, m, deterministic)
}
func (m *DMLMsgHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DMLMsgHeader.Merge(m, src)
}
func (m *DMLMsgHeader) XXX_Size() int {
	return xxx_messageInfo_DMLMsgHeader.Size(m)
}
func (m *DMLMsgHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DMLMsgHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DMLMsgHeader proto.InternalMessageInfo

func (m *DMLMsgHeader) GetBase() *MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *DMLMsgHeader) GetShardName() string {
	if m != nil {
		return m.ShardName
	}
	return ""
}

func init() {
	proto.RegisterEnum("milvus.proto.common.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("milvus.proto.common.IndexState", IndexState_name, IndexState_value)
	proto.RegisterEnum("milvus.proto.common.SegmentState", SegmentState_name, SegmentState_value)
	proto.RegisterEnum("milvus.proto.common.MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("milvus.proto.common.DslType", DslType_name, DslType_value)
	proto.RegisterEnum("milvus.proto.common.CompactionState", CompactionState_name, CompactionState_value)
	proto.RegisterEnum("milvus.proto.common.ConsistencyLevel", ConsistencyLevel_name, ConsistencyLevel_value)
	proto.RegisterType((*Status)(nil), "milvus.proto.common.Status")
	proto.RegisterType((*KeyValuePair)(nil), "milvus.proto.common.KeyValuePair")
	proto.RegisterType((*KeyDataPair)(nil), "milvus.proto.common.KeyDataPair")
	proto.RegisterType((*Blob)(nil), "milvus.proto.common.Blob")
	proto.RegisterType((*Address)(nil), "milvus.proto.common.Address")
	proto.RegisterType((*MsgBase)(nil), "milvus.proto.common.MsgBase")
	proto.RegisterType((*MsgHeader)(nil), "milvus.proto.common.MsgHeader")
	proto.RegisterType((*DMLMsgHeader)(nil), "milvus.proto.common.DMLMsgHeader")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 1509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5b, 0x73, 0x1b, 0x4b,
	0x11, 0xf6, 0x6a, 0x15, 0xcb, 0x6a, 0xcb, 0xf6, 0x64, 0x7c, 0x89, 0x4f, 0x08, 0x54, 0xca, 0x4f,
	0x29, 0x57, 0x9d, 0x04, 0x48, 0x01, 0x4f, 0xe7, 0xc1, 0xd2, 0xda, 0x8e, 0x2a, 0xb1, 0x63, 0x24,
	0x27, 0x9c, 0xe2, 0x81, 0xd4, 0x78, 0xb7, 0x2d, 0x0d, 0x99, 0x9d, 0x11, 0x33, 0xb3, 0x8e, 0xc5,
	0x13, 0xfc, 0x03, 0x38, 0xff, 0x82, 0x2a, 0xa0, 0xb8, 0x43, 0xf1, 0x0b, 0xb8, 0x3f, 0x73, 0xf9,
	0x03, 0xfc, 0x00, 0xae, 0xe7, 0x4a, 0xf5, 0xec, 0x4a, 0xda, 0x53, 0x75, 0xf2, 0xc4, 0xdb, 0xf6,
	0xd7, 0xdd, 0x5f, 0xf7, 0x74, 0xf7, 0xf4, 0x0e, 0x74, 0x52, 0x93, 0xe7, 0x46, 0xdf, 0x9f, 0x58,
	0xe3, 0x0d, 0xdf, 0xcc, 0xa5, 0xba, 0x2a, 0x5c, 0x29, 0xdd, 0x2f, 0x55, 0x7b, 0x2f, 0x60, 0x79,
	0xe8, 0x85, 0x2f, 0x1c, 0x7f, 0x0b, 0x00, 0xad, 0x35, 0xf6, 0x45, 0x6a, 0x32, 0xdc, 0x8d, 0xee,
	0x46, 0xf7, 0xd6, 0x3f, 0xff, 0x99, 0xfb, 0x9f, 0xe0, 0x73, 0xff, 0x90, 0xcc, 0x7a, 0x26, 0xc3,
	0x41, 0x1b, 0x67, 0x9f, 0x7c, 0x07, 0x96, 0x2d, 0x0a, 0x67, 0xf4, 0x6e, 0xe3, 0x6e, 0x74, 0xaf,
	0x3d, 0xa8, 0xa4, 0xbd, 0x2f, 0x42, 0xe7, 0x31, 0x4e, 0x9f, 0x0b, 0x55, 0xe0, 0x99, 0x90, 0x96,
	0x33, 0x88, 0x5f, 0xe2, 0x34, 0xf0, 0xb7, 0x07, 0xf4, 0xc9, 0xb7, 0xe0, 0xc6, 0x15, 0xa9, 0x2b,
	0xc7, 0x52, 0xd8, 0x7b, 0x08, 0xab, 0x8f, 0x71, 0x9a, 0x08, 0x2f, 0x5e, 0xe3, 0xc6, 0xa1, 0x99,
	0x09, 0x2f, 0x82, 0x57, 0x67, 0x10, 0xbe, 0xf7, 0xee, 0x40, 0xb3, 0xab, 0xcc, 0xc5, 0x82, 0x32,
	0x0a, 0xca, 0x8a, 0xf2, 0x4d, 0x68, 0x1d, 0x64, 0x99, 0x45, 0xe7, 0xf8, 0x3a, 0x34, 0xe4, 0xa4,
	0x62, 0x6b, 0xc8, 0x09, 0x91, 0x4d, 0x8c, 0xf5, 0x81, 0x2c, 0x1e, 0x84, 0xef, 0xbd, 0x77, 0x22,
	0x68, 0x9d, 0xb8, 0x51, 0x57, 0x38, 0xe4, 0x5f, 0x82, 0x95, 0xdc, 0x8d, 0x5e, 0xf8, 0xe9, 0x64,
	0x56, 0x9a, 0x3b, 0x9f, 0x58, 0x9a, 0x13, 0x37, 0x3a, 0x9f, 0x4e, 0x70, 0xd0, 0xca, 0xcb, 0x0f,
	0xca, 0x24, 0x77, 0xa3, 0x7e, 0x52, 0x31, 0x97, 0x02, 0xbf, 0x03, 0x6d, 0x2f, 0x73, 0x74, 0x5e,
	0xe4, 0x93, 0xdd, 0xf8, 0x6e, 0x74, 0xaf, 0x39, 0x58, 0x00, 0xfc, 0x36, 0xac, 0x38, 0x53, 0xd8,
	0x14, 0xfb, 0xc9, 0x6e, 0x33, 0xb8, 0xcd, 0xe5, 0xbd, 0xb7, 0xa0, 0x7d, 0xe2, 0x46, 0x8f, 0x50,
	0x64, 0x68, 0xf9, 0x67, 0xa1, 0x79, 0x21, 0x5c, 0x99, 0xd1, 0xea, 0xeb, 0x33, 0xa2, 0x13, 0x0c,
	0x82, 0xe5, 0xde, 0xd7, 0xa0, 0x93, 0x9c, 0x3c, 0xf9, 0x3f, 0x18, 0x28, 0x75, 0x37, 0x16, 0x36,
	0x3b, 0x15, 0xf9, 0xac, 0x63, 0x0b, 0x60, 0xff, 0xd7, 0x4d, 0x68, 0xcf, 0xc7, 0x83, 0xaf, 0x42,
	0x6b, 0x58, 0xa4, 0x29, 0x3a, 0xc7, 0x96, 0xf8, 0x26, 0x6c, 0x3c, 0xd3, 0x78, 0x3d, 0xc1, 0xd4,
	0x63, 0x16, 0x6c, 0x58, 0xc4, 0x6f, 0xc2, 0x5a, 0xcf, 0x68, 0x8d, 0xa9, 0x3f, 0x12, 0x52, 0x61,
	0xc6, 0x1a, 0x7c, 0x0b, 0xd8, 0x19, 0xda, 0x5c, 0x3a, 0x27, 0x8d, 0x4e, 0x50, 0x4b, 0xcc, 0x58,
	0xcc, 0x6f, 0xc1, 0x66, 0xcf, 0x28, 0x85, 0xa9, 0x97, 0x46, 0x9f, 0x1a, 0x7f, 0x78, 0x2d, 0x9d,
	0x77, 0xac, 0x49, 0xb4, 0x7d, 0xa5, 0x70, 0x24, 0xd4, 0x81, 0x1d, 0x15, 0x39, 0x6a, 0xcf, 0x6e,
	0x10, 0x47, 0x05, 0x26, 0x32, 0x47, 0x4d, 0x4c, 0xac, 0x55, 0x43, 0xfb, 0x3a, 0xc3, 0x6b, 0xea,
	0x0f, 0x5b, 0xe1, 0x6f, 0xc0, 0x76, 0x85, 0xd6, 0x02, 0x88, 0x1c, 0x59, 0x9b, 0x6f, 0xc0, 0x6a,
	0xa5, 0x3a, 0x7f, 0x7a, 0xf6, 0x98, 0x41, 0x8d, 0x61, 0x60, 0x5e, 0x0d, 0x30, 0x35, 0x36, 0x63,
	0xab, 0xb5, 0x14, 0x9e, 0x63, 0xea, 0x8d, 0xed, 0x27, 0xac, 0x43, 0x09, 0x57, 0xe0, 0x10, 0x85,
	0x4d, 0xc7, 0x03, 0x74, 0x85, 0xf2, 0x6c, 0x8d, 0x33, 0xe8, 0x1c, 0x49, 0x85, 0xa7, 0xc6, 0x1f,
	0x99, 0x42, 0x67, 0x6c, 0x9d, 0xaf, 0x03, 0x9c, 0xa0, 0x17, 0x55, 0x05, 0x36, 0x28, 0x6c, 0x4f,
	0xa4, 0x63, 0xac, 0x00, 0xc6, 0x77, 0x80, 0xf7, 0x84, 0xd6, 0xc6, 0xf7, 0x2c, 0x0a, 0x8f, 0x47,
	0x46, 0x65, 0x68, 0xd9, 0x4d, 0x4a, 0xe7, 0x63, 0xb8, 0x54, 0xc8, 0xf8, 0xc2, 0x3a, 0x41, 0x85,
	0x73, 0xeb, 0xcd, 0x85, 0x75, 0x85, 0x93, 0xf5, 0x16, 0x25, 0xdf, 0x2d, 0xa4, 0xca, 0x42, 0x49,
	0xca, 0xb6, 0x6c, 0x53, 0x8e, 0x55, 0xf2, 0xa7, 0x4f, 0xfa, 0xc3, 0x73, 0xb6, 0xc3, 0xb7, 0xe1,
	0x66, 0x85, 0x9c, 0xa0, 0xb7, 0x32, 0x0d, 0xc5, 0xbb, 0x45, 0xa9, 0x3e, 0x2d, 0xfc, 0xd3, 0xcb,
	0x13, 0xcc, 0x8d, 0x9d, 0xb2, 0x5d, 0x6a, 0x68, 0x60, 0x9a, 0xb5, 0x88, 0xbd, 0x41, 0x11, 0x0e,
	0xf3, 0x89, 0x9f, 0x2e, 0xca, 0xcb, 0x6e, 0x73, 0x0e, 0x6b, 0x49, 0x32, 0xc0, 0x6f, 0x14, 0xe8,
	0xfc, 0x40, 0xa4, 0xc8, 0xfe, 0xde, 0xda, 0x7f, 0x1b, 0x20, 0xf8, 0xd2, 0x42, 0x42, 0xce, 0x61,
	0x7d, 0x21, 0x9d, 0x1a, 0x8d, 0x6c, 0x89, 0x77, 0x60, 0xe5, 0x99, 0x96, 0xce, 0x15, 0x98, 0xb1,
	0x88, 0xea, 0xd6, 0xd7, 0x67, 0xd6, 0x8c, 0xe8, 0x4a, 0xb3, 0x06, 0x69, 0x8f, 0xa4, 0x96, 0x6e,
	0x1c, 0x26, 0x06, 0x60, 0xb9, 0x2a, 0x60, 0x73, 0xdf, 0x41, 0x67, 0x88, 0x23, 0x1a, 0x8e, 0x92,
	0x7b, 0x0b, 0x58, 0x5d, 0x5e, 0xb0, 0xcf, 0xd3, 0x8e, 0x68, 0x78, 0x8f, 0xad, 0x79, 0x25, 0xf5,
	0x88, 0x35, 0x88, 0x6c, 0x88, 0x42, 0x05, 0xe2, 0x55, 0x68, 0x1d, 0xa9, 0x22, 0x44, 0x69, 0x86,
	0x98, 0x24, 0x90, 0xd9, 0x0d, 0x52, 0x25, 0xd6, 0x4c, 0x26, 0x98, 0xb1, 0xe5, 0xfd, 0xef, 0xb5,
	0xc3, 0xfe, 0x08, 0x6b, 0x60, 0x0d, 0xda, 0xcf, 0x74, 0x86, 0x97, 0x52, 0x63, 0xc6, 0x96, 0x42,
	0x2b, 0x42, 0xcb, 0x6a, 0x35, 0xc9, 0xe8, 0xc4, 0xe4, 0x5d, 0xc3, 0x90, 0xea, 0xf9, 0x48, 0xb8,
	0x1a, 0x74, 0x49, 0xfd, 0x4d, 0xd0, 0xa5, 0x56, 0x5e, 0xd4, 0xdd, 0x47, 0x54, 0xe7, 0xe1, 0xd8,
	0xbc, 0x5a, 0x60, 0x8e, 0x8d, 0x29, 0xd2, 0x31, 0xfa, 0xe1, 0xd4, 0x79, 0xcc, 0x7b, 0x46, 0x5f,
	0xca, 0x91, 0x63, 0x92, 0x22, 0x3d, 0x31, 0x22, 0xab, 0xb9, 0x7f, 0x9d, 0x3a, 0x3c, 0x40, 0x85,
	0xc2, 0xd5, 0x59, 0x5f, 0x86, 0x61, 0x0c, 0xa9, 0x1e, 0x28, 0x29, 0x1c, 0x53, 0x74, 0x14, 0xca,
	0xb2, 0x14, 0x73, 0x6a, 0xc2, 0x81, 0xf2, 0x68, 0x4b, 0x59, 0xf3, 0x2d, 0xd8, 0x28, 0xed, 0xcf,
	0x84, 0xf5, 0x32, 0x90, 0xfc, 0x26, 0x0a, 0xed, 0xb6, 0x66, 0xb2, 0xc0, 0x7e, 0x4b, 0x77, 0xbf,
	0xf3, 0x48, 0xb8, 0x05, 0xf4, 0xbb, 0x88, 0xef, 0xc0, 0xcd, 0xd9, 0xd1, 0x16, 0xf8, 0xef, 0x23,
	0xbe, 0x09, 0xeb, 0x74, 0xb4, 0x39, 0xe6, 0xd8, 0x1f, 0x02, 0x48, 0x87, 0xa8, 0x81, 0x7f, 0x0c,
	0x0c, 0xd5, 0x29, 0x6a, 0xf8, 0x9f, 0x42, 0x30, 0x62, 0xa8, 0xba, 0xee, 0xd8, 0xbb, 0x11, 0x65,
	0x3a, 0x0b, 0x56, 0xc1, 0xec, 0xbd, 0x60, 0x48, 0xac, 0x73, 0xc3, 0xf7, 0x83, 0x61, 0xc5, 0x39,
	0x47, 0x3f, 0x08, 0xe8, 0x23, 0xa1, 0x33, 0x73, 0x79, 0x39, 0x47, 0x3f, 0x8c, 0xf8, 0x2e, 0x6c,
	0x92, 0x7b, 0x57, 0x28, 0xa1, 0xd3, 0x85, 0xfd, 0x47, 0x11, 0x67, 0xb3, 0x42, 0x86, 0xa9, 0x66,
	0xdf, 0x6f, 0x84, 0xa2, 0x54, 0x09, 0x94, 0xd8, 0x0f, 0x1a, 0x7c, 0xbd, 0xac, 0x6e, 0x29, 0xff,
	0xb0, 0xc1, 0x57, 0x61, 0xb9, 0xaf, 0x1d, 0x5a, 0xcf, 0xbe, 0x43, 0x93, 0xb7, 0x5c, 0xde, 0x5d,
	0xf6, 0x5d, 0x9a, 0xef, 0x1b, 0x61, 0xf2, 0xd8, 0x3b, 0x41, 0x51, 0x6e, 0x19, 0xf6, 0x8f, 0x38,
	0x1c, 0xb5, 0xbe, 0x72, 0xfe, 0x19, 0x53, 0xa4, 0x63, 0xf4, 0x8b, 0xeb, 0xc4, 0xfe, 0x15, 0xf3,
	0xdb, 0xb0, 0x3d, 0xc3, 0xc2, 0x02, 0x98, 0x5f, 0xa4, 0x7f, 0xc7, 0xfc, 0x0e, 0xdc, 0x3a, 0x46,
	0xbf, 0x98, 0x03, 0x72, 0x92, 0xce, 0xcb, 0xd4, 0xb1, 0xff, 0xc4, 0xfc, 0x53, 0xb0, 0x73, 0x8c,
	0x7e, 0x5e, 0xdf, 0x9a, 0xf2, 0xbf, 0x31, 0x5f, 0x83, 0x95, 0x01, 0x6d, 0x08, 0xbc, 0x42, 0xf6,
	0x6e, 0x4c, 0x4d, 0x9a, 0x89, 0x55, 0x3a, 0xef, 0xc5, 0x54, 0xba, 0xaf, 0x08, 0x9f, 0x8e, 0x93,
	0xbc, 0x37, 0x16, 0x5a, 0xa3, 0x72, 0xec, 0xfd, 0x98, 0x6f, 0x03, 0x1b, 0x60, 0x6e, 0xae, 0xb0,
	0x06, 0x7f, 0x40, 0x9b, 0x9f, 0x07, 0xe3, 0x2f, 0x17, 0x68, 0xa7, 0x73, 0xc5, 0x87, 0x31, 0x95,
	0xba, 0xb4, 0xff, 0xb8, 0xe6, 0xa3, 0x98, 0x7f, 0x1a, 0x76, 0xcb, 0xdb, 0x3a, 0xab, 0x3f, 0x29,
	0x47, 0xd8, 0xd7, 0x97, 0x86, 0x7d, 0xab, 0x39, 0x67, 0x4c, 0x50, 0x79, 0x31, 0xf7, 0xfb, 0x76,
	0x93, 0x5a, 0x54, 0x79, 0x04, 0xd3, 0x3f, 0x37, 0xf9, 0x06, 0x40, 0x79, 0x77, 0x02, 0xf0, 0x97,
	0x26, 0xa5, 0x7e, 0x8c, 0x9e, 0x56, 0xff, 0x15, 0xda, 0x69, 0x40, 0xff, 0x3a, 0x43, 0xeb, 0x2b,
	0x85, 0xfd, 0xad, 0x49, 0xa5, 0x38, 0x97, 0x39, 0x9e, 0xcb, 0xf4, 0x25, 0xfb, 0x51, 0x9b, 0x4a,
	0x11, 0x32, 0x3d, 0x35, 0x19, 0x92, 0x8d, 0x63, 0x3f, 0x6e, 0x53, 0xbf, 0x69, 0x5e, 0xca, 0x7e,
	0xff, 0x24, 0xc8, 0xd5, 0x56, 0xec, 0x27, 0xec, 0xa7, 0xf4, 0x0b, 0x82, 0x4a, 0x3e, 0x1f, 0x3e,
	0x65, 0x3f, 0x6b, 0x53, 0xa8, 0x03, 0xa5, 0x4c, 0x2a, 0xfc, 0x7c, 0x6a, 0x7f, 0xde, 0xa6, 0xb1,
	0xaf, 0x45, 0xaf, 0xba, 0xf1, 0x8b, 0x36, 0xd5, 0xb4, 0xc2, 0xc3, 0xac, 0x24, 0xb4, 0xe8, 0x7e,
	0x19, 0x58, 0xe9, 0x65, 0x45, 0x99, 0x9c, 0x7b, 0xf6, 0xab, 0xf6, 0xfe, 0x1e, 0xb4, 0x12, 0xa7,
	0xc2, 0xaa, 0x6a, 0x41, 0x9c, 0x38, 0xc5, 0x96, 0xe8, 0x66, 0x77, 0x8d, 0x51, 0x87, 0xd7, 0x13,
	0xfb, 0xfc, 0x73, 0x2c, 0xda, 0xef, 0xc2, 0x46, 0xcf, 0xe4, 0x13, 0x31, 0x9f, 0x88, 0xb0, 0x9d,
	0xca, 0xb5, 0x86, 0x59, 0x79, 0xea, 0x25, 0x5a, 0x0f, 0x87, 0xd7, 0x98, 0x16, 0x9e, 0x36, 0x62,
	0x44, 0x22, 0x39, 0xd1, 0xd0, 0x66, 0xac, 0xb1, 0xff, 0x36, 0xb0, 0x9e, 0xd1, 0x4e, 0x3a, 0x8f,
	0x3a, 0x9d, 0x3e, 0xc1, 0x2b, 0x54, 0x61, 0xb7, 0x7a, 0x6b, 0xf4, 0x88, 0x2d, 0x85, 0x17, 0x03,
	0x86, 0x3f, 0x7f, 0xb9, 0x81, 0xbb, 0xf4, 0x8b, 0x0c, 0xcf, 0x82, 0x75, 0x80, 0xc3, 0x2b, 0xd4,
	0xbe, 0x10, 0x4a, 0x4d, 0x59, 0x4c, 0x72, 0xaf, 0x70, 0xde, 0xe4, 0xf2, 0x9b, 0xb4, 0x88, 0xbb,
	0x5f, 0xf8, 0xea, 0xc3, 0x91, 0xf4, 0xe3, 0xe2, 0x82, 0x9e, 0x2d, 0x0f, 0xca, 0x77, 0xcc, 0x9b,
	0xd2, 0x54, 0x5f, 0x0f, 0xa4, 0xf6, 0x68, 0xb5, 0x50, 0x0f, 0xc2, 0xd3, 0xe6, 0x41, 0xf9, 0xb4,
	0x99, 0x5c, 0x5c, 0x2c, 0x07, 0xf9, 0xe1, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x51, 0x37, 0xf3,
	0x7a, 0x2b, 0x0b, 0x00, 0x00,
}
