// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: api/proto/book.proto

package api

import (
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

type EmptyBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyBook) Reset() {
	*x = EmptyBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyBook) ProtoMessage() {}

func (x *EmptyBook) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyBook.ProtoReflect.Descriptor instead.
func (*EmptyBook) Descriptor() ([]byte, []int) {
	return file_api_proto_book_proto_rawDescGZIP(), []int{0}
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_api_proto_book_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type BookList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BookList) Reset() {
	*x = BookList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookList) ProtoMessage() {}

func (x *BookList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookList.ProtoReflect.Descriptor instead.
func (*BookList) Descriptor() ([]byte, []int) {
	return file_api_proto_book_proto_rawDescGZIP(), []int{2}
}

func (x *BookList) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type BookId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BookId) Reset() {
	*x = BookId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookId) ProtoMessage() {}

func (x *BookId) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookId.ProtoReflect.Descriptor instead.
func (*BookId) Descriptor() ([]byte, []int) {
	return file_api_proto_book_proto_rawDescGZIP(), []int{3}
}

func (x *BookId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_api_proto_book_proto protoreflect.FileDescriptor

var file_api_proto_book_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x0b, 0x0a, 0x09, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x42, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x08,
	0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x18, 0x0a, 0x06, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xbc, 0x01, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x56, 0x31, 0x12, 0x29,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x1a, 0x09, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x06, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a,
	0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x22,
	0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_book_proto_rawDescOnce sync.Once
	file_api_proto_book_proto_rawDescData = file_api_proto_book_proto_rawDesc
)

func file_api_proto_book_proto_rawDescGZIP() []byte {
	file_api_proto_book_proto_rawDescOnce.Do(func() {
		file_api_proto_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_book_proto_rawDescData)
	})
	return file_api_proto_book_proto_rawDescData
}

var file_api_proto_book_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_book_proto_goTypes = []interface{}{
	(*EmptyBook)(nil), // 0: api.EmptyBook
	(*Book)(nil),      // 1: api.Book
	(*BookList)(nil),  // 2: api.BookList
	(*BookId)(nil),    // 3: api.BookId
}
var file_api_proto_book_proto_depIdxs = []int32{
	1, // 0: api.BookList.books:type_name -> api.Book
	0, // 1: api.BookV1.GetAll:input_type -> api.EmptyBook
	3, // 2: api.BookV1.Get:input_type -> api.BookId
	1, // 3: api.BookV1.Insert:input_type -> api.Book
	1, // 4: api.BookV1.Update:input_type -> api.Book
	3, // 5: api.BookV1.Remove:input_type -> api.BookId
	2, // 6: api.BookV1.GetAll:output_type -> api.BookList
	1, // 7: api.BookV1.Get:output_type -> api.Book
	1, // 8: api.BookV1.Insert:output_type -> api.Book
	1, // 9: api.BookV1.Update:output_type -> api.Book
	1, // 10: api.BookV1.Remove:output_type -> api.Book
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_book_proto_init() }
func file_api_proto_book_proto_init() {
	if File_api_proto_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyBook); i {
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
		file_api_proto_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_api_proto_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookList); i {
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
		file_api_proto_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookId); i {
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
			RawDescriptor: file_api_proto_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_book_proto_goTypes,
		DependencyIndexes: file_api_proto_book_proto_depIdxs,
		MessageInfos:      file_api_proto_book_proto_msgTypes,
	}.Build()
	File_api_proto_book_proto = out.File
	file_api_proto_book_proto_rawDesc = nil
	file_api_proto_book_proto_goTypes = nil
	file_api_proto_book_proto_depIdxs = nil
}
