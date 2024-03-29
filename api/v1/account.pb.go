// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: api/v1/account.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type AccountListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique 不支持重复消息。
	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// 必须大于 0
	MaxHealth           float64 `protobuf:"fixed64,2,opt,name=max_health,json=maxHealth,proto3" json:"max_health,omitempty"`
	MinBorrowValueInEth float64 `protobuf:"fixed64,3,opt,name=min_borrow_value_in_eth,json=minBorrowValueInEth,proto3" json:"min_borrow_value_in_eth,omitempty"`
	// 必须大于 0
	PageNumber uint32 `protobuf:"varint,4,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// 大于0小于100
	PageSize uint32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *AccountListRequest) Reset() {
	*x = AccountListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountListRequest) ProtoMessage() {}

func (x *AccountListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountListRequest.ProtoReflect.Descriptor instead.
func (*AccountListRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *AccountListRequest) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *AccountListRequest) GetMaxHealth() float64 {
	if x != nil {
		return x.MaxHealth
	}
	return 0
}

func (x *AccountListRequest) GetMinBorrowValueInEth() float64 {
	if x != nil {
		return x.MinBorrowValueInEth
	}
	return 0
}

func (x *AccountListRequest) GetPageNumber() uint32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *AccountListRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// The response message containing the greetings
type AccountListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*AccountListReply_Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	//    int32 error = 2;
	PaginationSummary *PaginationSummary  `protobuf:"bytes,3,opt,name=pagination_summary,json=paginationSummary,proto3" json:"pagination_summary,omitempty"`
	Request           *AccountListRequest `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *AccountListReply) Reset() {
	*x = AccountListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountListReply) ProtoMessage() {}

func (x *AccountListReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountListReply.ProtoReflect.Descriptor instead.
func (*AccountListReply) Descriptor() ([]byte, []int) {
	return file_api_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountListReply) GetAccounts() []*AccountListReply_Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *AccountListReply) GetPaginationSummary() *PaginationSummary {
	if x != nil {
		return x.PaginationSummary
	}
	return nil
}

func (x *AccountListReply) GetRequest() *AccountListRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type PaginationSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber   uint32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize     uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TotalEntries uint32 `protobuf:"varint,3,opt,name=total_entries,json=totalEntries,proto3" json:"total_entries,omitempty"`
	TotalPages   uint32 `protobuf:"varint,4,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
}

func (x *PaginationSummary) Reset() {
	*x = PaginationSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationSummary) ProtoMessage() {}

func (x *PaginationSummary) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationSummary.ProtoReflect.Descriptor instead.
func (*PaginationSummary) Descriptor() ([]byte, []int) {
	return file_api_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationSummary) GetPageNumber() uint32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PaginationSummary) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaginationSummary) GetTotalEntries() uint32 {
	if x != nil {
		return x.TotalEntries
	}
	return 0
}

func (x *PaginationSummary) GetTotalPages() uint32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address                       string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Symbol                        string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	BorrowBalanceUnderlying       string `protobuf:"bytes,3,opt,name=borrow_balance_underlying,json=borrowBalanceUnderlying,proto3" json:"borrow_balance_underlying,omitempty"`
	SupplyBalanceUnderlying       string `protobuf:"bytes,4,opt,name=supply_balance_underlying,json=supplyBalanceUnderlying,proto3" json:"supply_balance_underlying,omitempty"`
	LifetimeBorrowInterestAccrued string `protobuf:"bytes,5,opt,name=lifetime_borrow_interest_accrued,json=lifetimeBorrowInterestAccrued,proto3" json:"lifetime_borrow_interest_accrued,omitempty"`
	LifetimeSupplyInterestAccrued string `protobuf:"bytes,6,opt,name=lifetime_supply_interest_accrued,json=lifetimeSupplyInterestAccrued,proto3" json:"lifetime_supply_interest_accrued,omitempty"`
	SafeWithdrawAmountUnderlying  string `protobuf:"bytes,7,opt,name=safe_withdraw_amount_underlying,json=safeWithdrawAmountUnderlying,proto3" json:"safe_withdraw_amount_underlying,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_api_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Token) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Token) GetBorrowBalanceUnderlying() string {
	if x != nil {
		return x.BorrowBalanceUnderlying
	}
	return ""
}

func (x *Token) GetSupplyBalanceUnderlying() string {
	if x != nil {
		return x.SupplyBalanceUnderlying
	}
	return ""
}

func (x *Token) GetLifetimeBorrowInterestAccrued() string {
	if x != nil {
		return x.LifetimeBorrowInterestAccrued
	}
	return ""
}

func (x *Token) GetLifetimeSupplyInterestAccrued() string {
	if x != nil {
		return x.LifetimeSupplyInterestAccrued
	}
	return ""
}

func (x *Token) GetSafeWithdrawAmountUnderlying() string {
	if x != nil {
		return x.SafeWithdrawAmountUnderlying
	}
	return ""
}

type AccountListReply_Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address                   string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Health                    string   `protobuf:"bytes,2,opt,name=health,proto3" json:"health,omitempty"`
	Tokens                    []*Token `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens,omitempty"`
	TotalBorrowValueInEth     string   `protobuf:"bytes,4,opt,name=total_borrow_value_in_eth,json=totalBorrowValueInEth,proto3" json:"total_borrow_value_in_eth,omitempty"`
	TotalCollateralValueInEth string   `protobuf:"bytes,5,opt,name=total_collateral_value_in_eth,json=totalCollateralValueInEth,proto3" json:"total_collateral_value_in_eth,omitempty"`
}

func (x *AccountListReply_Account) Reset() {
	*x = AccountListReply_Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountListReply_Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountListReply_Account) ProtoMessage() {}

func (x *AccountListReply_Account) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountListReply_Account.ProtoReflect.Descriptor instead.
func (*AccountListReply_Account) Descriptor() ([]byte, []int) {
	return file_api_v1_account_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AccountListReply_Account) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AccountListReply_Account) GetHealth() string {
	if x != nil {
		return x.Health
	}
	return ""
}

func (x *AccountListReply_Account) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *AccountListReply_Account) GetTotalBorrowValueInEth() string {
	if x != nil {
		return x.TotalBorrowValueInEth
	}
	return ""
}

func (x *AccountListReply_Account) GetTotalCollateralValueInEth() string {
	if x != nil {
		return x.TotalCollateralValueInEth
	}
	return ""
}

var File_api_v1_account_proto protoreflect.FileDescriptor

var file_api_v1_account_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x31, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x18, 0x01, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x28,
	0x01, 0xfa, 0x42, 0x1e, 0x92, 0x01, 0x1b, 0x22, 0x19, 0x72, 0x17, 0x28, 0x2a, 0x32, 0x13, 0x5e,
	0x30, 0x78, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x34, 0x30,
	0x7d, 0x24, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x34, 0x0a,
	0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x42, 0x15, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0xfa, 0x42, 0x04, 0x12, 0x02, 0x40, 0x01, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x12, 0x4b, 0x0a, 0x17, 0x6d, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x65, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x42, 0x15, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x21, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xfa, 0x42, 0x04, 0x12, 0x02, 0x40, 0x01, 0x52, 0x13, 0x6d, 0x69, 0x6e,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x45, 0x74, 0x68,
	0x12, 0x2f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0e, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0xfa, 0x42,
	0x04, 0x2a, 0x02, 0x40, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x2d, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x10, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x64, 0x20, 0x00, 0xfa,
	0x42, 0x04, 0x2a, 0x02, 0x40, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0xb7, 0x03, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x12, 0x48, 0x0a, 0x12, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x11, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x34, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0xde, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x25, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x19, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69,
	0x6e, 0x5f, 0x65, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x45,
	0x74, 0x68, 0x12, 0x40, 0x0a, 0x1d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6c, 0x6c,
	0x61, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x5f,
	0x65, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49,
	0x6e, 0x45, 0x74, 0x68, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x8a, 0x03, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x3a, 0x0a, 0x19, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x17, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x19,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x75,
	0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x17, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x6e,
	0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x12, 0x47, 0x0a, 0x20, 0x6c, 0x69, 0x66, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x72, 0x75, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x1d, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x72, 0x75, 0x65,
	0x64, 0x12, 0x47, 0x0a, 0x20, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x63,
	0x63, 0x72, 0x75, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x6c, 0x69, 0x66,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x72, 0x75, 0x65, 0x64, 0x12, 0x45, 0x0a, 0x1f, 0x73, 0x61,
	0x66, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x1c, 0x73, 0x61, 0x66, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e,
	0x67, 0x32, 0x60, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x55, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x1c, 0x5a, 0x1a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_account_proto_rawDescOnce sync.Once
	file_api_v1_account_proto_rawDescData = file_api_v1_account_proto_rawDesc
)

func file_api_v1_account_proto_rawDescGZIP() []byte {
	file_api_v1_account_proto_rawDescOnce.Do(func() {
		file_api_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_account_proto_rawDescData)
	})
	return file_api_v1_account_proto_rawDescData
}

var file_api_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_account_proto_goTypes = []interface{}{
	(*AccountListRequest)(nil),       // 0: api.v1.AccountListRequest
	(*AccountListReply)(nil),         // 1: api.v1.AccountListReply
	(*PaginationSummary)(nil),        // 2: api.v1.PaginationSummary
	(*Token)(nil),                    // 3: api.v1.Token
	(*AccountListReply_Account)(nil), // 4: api.v1.AccountListReply.Account
}
var file_api_v1_account_proto_depIdxs = []int32{
	4, // 0: api.v1.AccountListReply.accounts:type_name -> api.v1.AccountListReply.Account
	2, // 1: api.v1.AccountListReply.pagination_summary:type_name -> api.v1.PaginationSummary
	0, // 2: api.v1.AccountListReply.request:type_name -> api.v1.AccountListRequest
	3, // 3: api.v1.AccountListReply.Account.tokens:type_name -> api.v1.Token
	0, // 4: api.v1.Account.List:input_type -> api.v1.AccountListRequest
	1, // 5: api.v1.Account.List:output_type -> api.v1.AccountListReply
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_account_proto_init() }
func file_api_v1_account_proto_init() {
	if File_api_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountListRequest); i {
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
		file_api_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountListReply); i {
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
		file_api_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationSummary); i {
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
		file_api_v1_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_api_v1_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountListReply_Account); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_account_proto_goTypes,
		DependencyIndexes: file_api_v1_account_proto_depIdxs,
		MessageInfos:      file_api_v1_account_proto_msgTypes,
	}.Build()
	File_api_v1_account_proto = out.File
	file_api_v1_account_proto_rawDesc = nil
	file_api_v1_account_proto_goTypes = nil
	file_api_v1_account_proto_depIdxs = nil
}
