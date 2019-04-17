// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/transaction.proto

package com_thoughtworks_zpay

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AccountType int32

const (
	AccountType_ACCOUNT_TYPE_UNKNOWN AccountType = 0
	AccountType_PAYTM_WALLET         AccountType = 1
	AccountType_UPI_ID               AccountType = 2
	AccountType_CREDIT_CARD          AccountType = 3
	AccountType_INTERNAL_ACCOUNT     AccountType = 4
)

var AccountType_name = map[int32]string{
	0: "ACCOUNT_TYPE_UNKNOWN",
	1: "PAYTM_WALLET",
	2: "UPI_ID",
	3: "CREDIT_CARD",
	4: "INTERNAL_ACCOUNT",
}

var AccountType_value = map[string]int32{
	"ACCOUNT_TYPE_UNKNOWN": 0,
	"PAYTM_WALLET":         1,
	"UPI_ID":               2,
	"CREDIT_CARD":          3,
	"INTERNAL_ACCOUNT":     4,
}

func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}

func (AccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_80e80868f3cd63d0, []int{0}
}

type Currency int32

const (
	Currency_CURRENCY_UNKNOWN Currency = 0
	Currency_CURRENCY_USD     Currency = 1
	Currency_CURRENCY_INR     Currency = 2
)

var Currency_name = map[int32]string{
	0: "CURRENCY_UNKNOWN",
	1: "CURRENCY_USD",
	2: "CURRENCY_INR",
}

var Currency_value = map[string]int32{
	"CURRENCY_UNKNOWN": 0,
	"CURRENCY_USD":     1,
	"CURRENCY_INR":     2,
}

func (x Currency) String() string {
	return proto.EnumName(Currency_name, int32(x))
}

func (Currency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_80e80868f3cd63d0, []int{1}
}

type TransactionStatus int32

const (
	TransactionStatus_STATUS_UNKNOWN TransactionStatus = 0
	TransactionStatus_INITIATED      TransactionStatus = 1
	TransactionStatus_BANK_FAILURE   TransactionStatus = 2
	TransactionStatus_SUCCESS        TransactionStatus = 3
	TransactionStatus_SUSPECT        TransactionStatus = 4
)

var TransactionStatus_name = map[int32]string{
	0: "STATUS_UNKNOWN",
	1: "INITIATED",
	2: "BANK_FAILURE",
	3: "SUCCESS",
	4: "SUSPECT",
}

var TransactionStatus_value = map[string]int32{
	"STATUS_UNKNOWN": 0,
	"INITIATED":      1,
	"BANK_FAILURE":   2,
	"SUCCESS":        3,
	"SUSPECT":        4,
}

func (x TransactionStatus) String() string {
	return proto.EnumName(TransactionStatus_name, int32(x))
}

func (TransactionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_80e80868f3cd63d0, []int{2}
}

type CompletedTransaction struct {
	TransactionReference string            `protobuf:"bytes,1,opt,name=transaction_reference,json=transactionReference,proto3" json:"transaction_reference,omitempty"`
	Amount               *Money            `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	DebitAccount         *Account          `protobuf:"bytes,3,opt,name=debit_account,json=debitAccount,proto3" json:"debit_account,omitempty"`
	CreditAccount        *Account          `protobuf:"bytes,4,opt,name=credit_account,json=creditAccount,proto3" json:"credit_account,omitempty"`
	Description          string            `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Status               TransactionStatus `protobuf:"varint,6,opt,name=status,proto3,enum=com.thoughtworks.zpay.TransactionStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CompletedTransaction) Reset()         { *m = CompletedTransaction{} }
func (m *CompletedTransaction) String() string { return proto.CompactTextString(m) }
func (*CompletedTransaction) ProtoMessage()    {}
func (*CompletedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_80e80868f3cd63d0, []int{0}
}

func (m *CompletedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompletedTransaction.Unmarshal(m, b)
}
func (m *CompletedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompletedTransaction.Marshal(b, m, deterministic)
}
func (m *CompletedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompletedTransaction.Merge(m, src)
}
func (m *CompletedTransaction) XXX_Size() int {
	return xxx_messageInfo_CompletedTransaction.Size(m)
}
func (m *CompletedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_CompletedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_CompletedTransaction proto.InternalMessageInfo

func (m *CompletedTransaction) GetTransactionReference() string {
	if m != nil {
		return m.TransactionReference
	}
	return ""
}

func (m *CompletedTransaction) GetAmount() *Money {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *CompletedTransaction) GetDebitAccount() *Account {
	if m != nil {
		return m.DebitAccount
	}
	return nil
}

func (m *CompletedTransaction) GetCreditAccount() *Account {
	if m != nil {
		return m.CreditAccount
	}
	return nil
}

func (m *CompletedTransaction) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CompletedTransaction) GetStatus() TransactionStatus {
	if m != nil {
		return m.Status
	}
	return TransactionStatus_STATUS_UNKNOWN
}

type Money struct {
	CurrencyCode         Currency `protobuf:"varint,1,opt,name=currency_code,json=currencyCode,proto3,enum=com.thoughtworks.zpay.Currency" json:"currency_code,omitempty"`
	Value                int64    `protobuf:"zigzag64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Money) Reset()         { *m = Money{} }
func (m *Money) String() string { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()    {}
func (*Money) Descriptor() ([]byte, []int) {
	return fileDescriptor_80e80868f3cd63d0, []int{1}
}

func (m *Money) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Money.Unmarshal(m, b)
}
func (m *Money) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Money.Marshal(b, m, deterministic)
}
func (m *Money) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Money.Merge(m, src)
}
func (m *Money) XXX_Size() int {
	return xxx_messageInfo_Money.Size(m)
}
func (m *Money) XXX_DiscardUnknown() {
	xxx_messageInfo_Money.DiscardUnknown(m)
}

var xxx_messageInfo_Money proto.InternalMessageInfo

func (m *Money) GetCurrencyCode() Currency {
	if m != nil {
		return m.CurrencyCode
	}
	return Currency_CURRENCY_UNKNOWN
}

func (m *Money) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Account struct {
	AccountType          AccountType `protobuf:"varint,1,opt,name=accountType,proto3,enum=com.thoughtworks.zpay.AccountType" json:"accountType,omitempty"`
	AccountIdentifier    string      `protobuf:"bytes,2,opt,name=account_identifier,json=accountIdentifier,proto3" json:"account_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_80e80868f3cd63d0, []int{2}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAccountType() AccountType {
	if m != nil {
		return m.AccountType
	}
	return AccountType_ACCOUNT_TYPE_UNKNOWN
}

func (m *Account) GetAccountIdentifier() string {
	if m != nil {
		return m.AccountIdentifier
	}
	return ""
}

func init() {
	proto.RegisterEnum("com.thoughtworks.zpay.AccountType", AccountType_name, AccountType_value)
	proto.RegisterEnum("com.thoughtworks.zpay.Currency", Currency_name, Currency_value)
	proto.RegisterEnum("com.thoughtworks.zpay.TransactionStatus", TransactionStatus_name, TransactionStatus_value)
	proto.RegisterType((*CompletedTransaction)(nil), "com.thoughtworks.zpay.CompletedTransaction")
	proto.RegisterType((*Money)(nil), "com.thoughtworks.zpay.Money")
	proto.RegisterType((*Account)(nil), "com.thoughtworks.zpay.Account")
}

func init() { proto.RegisterFile("proto/transaction.proto", fileDescriptor_80e80868f3cd63d0) }

var fileDescriptor_80e80868f3cd63d0 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x5f, 0x6f, 0xda, 0x3c,
	0x14, 0xc6, 0xdf, 0xf0, 0xaf, 0x2f, 0x27, 0xc0, 0x5c, 0x8b, 0xaa, 0x5c, 0x4c, 0x1b, 0xe2, 0x0a,
	0x21, 0x8d, 0x49, 0x74, 0x1f, 0x60, 0xa9, 0x93, 0x49, 0x51, 0x69, 0x8a, 0x1c, 0x47, 0x15, 0x57,
	0x56, 0x48, 0xdc, 0x35, 0x1a, 0xc4, 0x28, 0x98, 0x4d, 0xec, 0x62, 0xdf, 0x79, 0xdf, 0x60, 0xc2,
	0x49, 0xd6, 0x54, 0x2b, 0xd3, 0x2e, 0xcf, 0x39, 0xcf, 0xf3, 0x3b, 0x3a, 0x8f, 0x65, 0xb8, 0xdc,
	0x66, 0x52, 0xc9, 0xf7, 0x2a, 0x0b, 0xd3, 0x5d, 0x18, 0xa9, 0x44, 0xa6, 0x53, 0xdd, 0xc1, 0x17,
	0x91, 0xdc, 0x4c, 0xd5, 0xa3, 0xdc, 0x7f, 0x7e, 0x54, 0xdf, 0x64, 0xf6, 0x65, 0x37, 0xfd, 0xbe,
	0x0d, 0x0f, 0xa3, 0x9f, 0x35, 0xe8, 0x13, 0xb9, 0xd9, 0xae, 0x85, 0x12, 0x31, 0x7b, 0x72, 0xe1,
	0x2b, 0xb8, 0xa8, 0x40, 0x78, 0x26, 0x1e, 0x44, 0x26, 0xd2, 0x48, 0x0c, 0x8c, 0xa1, 0x31, 0x6e,
	0xd3, 0x7e, 0x65, 0x48, 0xcb, 0x19, 0xfe, 0x00, 0xad, 0x70, 0x23, 0xf7, 0xa9, 0x1a, 0xd4, 0x86,
	0xc6, 0xd8, 0x9c, 0xbd, 0x9e, 0xbe, 0xb8, 0x75, 0x7a, 0x2b, 0x53, 0x71, 0xa0, 0x85, 0x16, 0x13,
	0xe8, 0xc6, 0x62, 0x95, 0x28, 0x1e, 0x46, 0x91, 0x36, 0xd7, 0xb5, 0xf9, 0xcd, 0x09, 0xb3, 0x95,
	0xab, 0x68, 0x47, 0x9b, 0x8a, 0x0a, 0x3b, 0xd0, 0x8b, 0x32, 0x11, 0x57, 0x28, 0x8d, 0x7f, 0xa2,
	0x74, 0x73, 0x57, 0x89, 0x19, 0x82, 0x19, 0x8b, 0x5d, 0x94, 0x25, 0xdb, 0xe3, 0x65, 0x83, 0xa6,
	0x3e, 0xb6, 0xda, 0xc2, 0x1f, 0xa1, 0xb5, 0x53, 0xa1, 0xda, 0xef, 0x06, 0xad, 0xa1, 0x31, 0xee,
	0xcd, 0xc6, 0x27, 0x16, 0x54, 0xc2, 0xf4, 0xb5, 0x9e, 0x16, 0xbe, 0x51, 0x04, 0x4d, 0x1d, 0x00,
	0xb6, 0xa1, 0x1b, 0xed, 0xb3, 0x63, 0x74, 0x07, 0x1e, 0xc9, 0x38, 0xcf, 0xb6, 0x37, 0x7b, 0x7b,
	0x82, 0x48, 0x0a, 0x2d, 0xed, 0x94, 0x2e, 0x22, 0x63, 0x81, 0xfb, 0xd0, 0xfc, 0x1a, 0xae, 0xf7,
	0x42, 0x67, 0x8e, 0x69, 0x5e, 0x8c, 0x7e, 0xc0, 0x59, 0x79, 0x93, 0x0d, 0x66, 0x91, 0x09, 0x3b,
	0x6c, 0xcb, 0x25, 0xa3, 0xbf, 0xe7, 0x72, 0x54, 0xd2, 0xaa, 0x0d, 0xbf, 0x03, 0x5c, 0x94, 0x3c,
	0x89, 0x45, 0xaa, 0x92, 0x87, 0x44, 0x64, 0x7a, 0x67, 0x9b, 0x9e, 0x17, 0x13, 0xf7, 0xf7, 0x60,
	0xb2, 0x06, 0xb3, 0x82, 0xc2, 0x03, 0xe8, 0x5b, 0x84, 0xdc, 0x05, 0x1e, 0xe3, 0x6c, 0xb9, 0x70,
	0x78, 0xe0, 0xdd, 0x78, 0x77, 0xf7, 0x1e, 0xfa, 0x0f, 0x23, 0xe8, 0x2c, 0xac, 0x25, 0xbb, 0xe5,
	0xf7, 0xd6, 0x7c, 0xee, 0x30, 0x64, 0x60, 0x80, 0x56, 0xb0, 0x70, 0xb9, 0x6b, 0xa3, 0x1a, 0x7e,
	0x05, 0x26, 0xa1, 0x8e, 0xed, 0x32, 0x4e, 0x2c, 0x6a, 0xa3, 0x3a, 0xee, 0x03, 0x72, 0x3d, 0xe6,
	0x50, 0xcf, 0x9a, 0xf3, 0x82, 0x88, 0x1a, 0x13, 0x1b, 0xfe, 0x2f, 0xd3, 0x39, 0x2a, 0x48, 0x40,
	0xa9, 0xe3, 0x91, 0xe5, 0xf3, 0x35, 0x4f, 0x5d, 0xdf, 0x46, 0xc6, 0xb3, 0x8e, 0xeb, 0x51, 0x54,
	0x9b, 0xac, 0xe0, 0xfc, 0x8f, 0x57, 0xc3, 0x18, 0x7a, 0x3e, 0xb3, 0x58, 0xe0, 0x57, 0x60, 0x5d,
	0x68, 0xbb, 0x9e, 0xcb, 0x5c, 0x8b, 0x39, 0x05, 0xe9, 0xda, 0xf2, 0x6e, 0xf8, 0x27, 0xcb, 0x9d,
	0x07, 0xd4, 0x41, 0x35, 0x6c, 0xc2, 0x99, 0x1f, 0x10, 0xe2, 0xf8, 0x3e, 0xaa, 0xe7, 0x85, 0xbf,
	0x70, 0x08, 0x43, 0x8d, 0xeb, 0x4b, 0x78, 0xf9, 0x27, 0xae, 0x5a, 0xfa, 0x9f, 0x5e, 0xfd, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0x17, 0x22, 0x68, 0xc2, 0x03, 0x00, 0x00,
}
