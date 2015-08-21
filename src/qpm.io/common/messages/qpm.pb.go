// Code generated by protoc-gen-go.
// source: qpm.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	qpm.proto

It has these top-level messages:
	Package
	Dependency
	VersionInfo
	SearchResult
	PingRequest
	PingResponse
	PublishRequest
	PublishResponse
	DependencyRequest
	DependencyResponse
	SearchRequest
	SearchResponse
	LoginRequest
	LoginResponse
	InfoRequest
	InfoResponse
*/
package messages

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type RepoType int32

const (
	RepoType_AUTO   RepoType = 0
	RepoType_GITHUB RepoType = 1
)

var RepoType_name = map[int32]string{
	0: "AUTO",
	1: "GITHUB",
}
var RepoType_value = map[string]int32{
	"AUTO":   0,
	"GITHUB": 1,
}

func (x RepoType) String() string {
	return proto.EnumName(RepoType_name, int32(x))
}

// The values in this enum should correspond to an SPDX identifier
type LicenseType int32

const (
	LicenseType_NONE LicenseType = 0
	LicenseType_MIT  LicenseType = 1
)

var LicenseType_name = map[int32]string{
	0: "NONE",
	1: "MIT",
}
var LicenseType_value = map[string]int32{
	"NONE": 0,
	"MIT":  1,
}

func (x LicenseType) String() string {
	return proto.EnumName(LicenseType_name, int32(x))
}

type Package struct {
	Name         string              `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description  string              `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Author       *Package_Author     `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	Repository   *Package_Repository `protobuf:"bytes,4,opt,name=repository" json:"repository,omitempty"`
	Version      *Package_Version    `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	Dependencies []string            `protobuf:"bytes,6,rep,name=dependencies" json:"dependencies,omitempty"`
	License      LicenseType         `protobuf:"varint,7,opt,name=license,enum=messages.LicenseType" json:"license,omitempty"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}

func (m *Package) GetAuthor() *Package_Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Package) GetRepository() *Package_Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *Package) GetVersion() *Package_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type Package_Repository struct {
	Type RepoType `protobuf:"varint,1,opt,name=type,enum=messages.RepoType" json:"type,omitempty"`
	Url  string   `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *Package_Repository) Reset()         { *m = Package_Repository{} }
func (m *Package_Repository) String() string { return proto.CompactTextString(m) }
func (*Package_Repository) ProtoMessage()    {}

type Package_Version struct {
	Label    string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	Revision string `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
}

func (m *Package_Version) Reset()         { *m = Package_Version{} }
func (m *Package_Version) String() string { return proto.CompactTextString(m) }
func (*Package_Version) ProtoMessage()    {}

type Package_Author struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *Package_Author) Reset()         { *m = Package_Author{} }
func (m *Package_Author) String() string { return proto.CompactTextString(m) }
func (*Package_Author) ProtoMessage()    {}

type Dependency struct {
	Name       string              `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Repository *Package_Repository `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
	Version    *Package_Version    `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *Dependency) Reset()         { *m = Dependency{} }
func (m *Dependency) String() string { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()    {}

func (m *Dependency) GetRepository() *Package_Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *Dependency) GetVersion() *Package_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type VersionInfo struct {
	Version       *Package_Version `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	DatePublished string           `protobuf:"bytes,2,opt,name=date_published" json:"date_published,omitempty"`
}

func (m *VersionInfo) Reset()         { *m = VersionInfo{} }
func (m *VersionInfo) String() string { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()    {}

func (m *VersionInfo) GetVersion() *Package_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type SearchResult struct {
	Name    string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version string          `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Author  *Package_Author `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
}

func (m *SearchResult) Reset()         { *m = SearchResult{} }
func (m *SearchResult) String() string { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()    {}

func (m *SearchResult) GetAuthor() *Package_Author {
	if m != nil {
		return m.Author
	}
	return nil
}

type PingRequest struct {
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}

type PingResponse struct {
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}

type PublishRequest struct {
	PackageDescription *Package `protobuf:"bytes,1,opt,name=package_description" json:"package_description,omitempty"`
	Token              string   `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}

func (m *PublishRequest) GetPackageDescription() *Package {
	if m != nil {
		return m.PackageDescription
	}
	return nil
}

type PublishResponse struct {
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}

type DependencyRequest struct {
	PackageNames []string `protobuf:"bytes,1,rep,name=package_names" json:"package_names,omitempty"`
}

func (m *DependencyRequest) Reset()         { *m = DependencyRequest{} }
func (m *DependencyRequest) String() string { return proto.CompactTextString(m) }
func (*DependencyRequest) ProtoMessage()    {}

type DependencyResponse struct {
	Dependencies []*Dependency `protobuf:"bytes,1,rep,name=dependencies" json:"dependencies,omitempty"`
}

func (m *DependencyResponse) Reset()         { *m = DependencyResponse{} }
func (m *DependencyResponse) String() string { return proto.CompactTextString(m) }
func (*DependencyResponse) ProtoMessage()    {}

func (m *DependencyResponse) GetDependencies() []*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

type SearchRequest struct {
	PackageName string `protobuf:"bytes,1,opt,name=package_name" json:"package_name,omitempty"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}

type SearchResponse struct {
	Results []*SearchResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}

func (m *SearchResponse) GetResults() []*SearchResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type LoginRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Create   bool   `protobuf:"varint,3,opt,name=create" json:"create,omitempty"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}

type InfoRequest struct {
	PackageName string `protobuf:"bytes,1,opt,name=package_name" json:"package_name,omitempty"`
}

func (m *InfoRequest) Reset()         { *m = InfoRequest{} }
func (m *InfoRequest) String() string { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()    {}

type InfoResponse struct {
	Package      *Package       `protobuf:"bytes,1,opt,name=package" json:"package,omitempty"`
	Versions     []*VersionInfo `protobuf:"bytes,2,rep,name=versions" json:"versions,omitempty"`
	Dependencies []*Dependency  `protobuf:"bytes,3,rep,name=dependencies" json:"dependencies,omitempty"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}

func (m *InfoResponse) GetPackage() *Package {
	if m != nil {
		return m.Package
	}
	return nil
}

func (m *InfoResponse) GetVersions() []*VersionInfo {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *InfoResponse) GetDependencies() []*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func init() {
	proto.RegisterEnum("messages.RepoType", RepoType_name, RepoType_value)
	proto.RegisterEnum("messages.LicenseType", LicenseType_name, LicenseType_value)
}

// Client API for Qpm service

type QpmClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	GetDependencies(ctx context.Context, in *DependencyRequest, opts ...grpc.CallOption) (*DependencyResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type qpmClient struct {
	cc *grpc.ClientConn
}

func NewQpmClient(cc *grpc.ClientConn) QpmClient {
	return &qpmClient{cc}
}

func (c *qpmClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) GetDependencies(ctx context.Context, in *DependencyRequest, opts ...grpc.CallOption) (*DependencyResponse, error) {
	out := new(DependencyResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/GetDependencies", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qpmClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/messages.Qpm/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Qpm service

type QpmServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	GetDependencies(context.Context, *DependencyRequest) (*DependencyResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
}

func RegisterQpmServer(s *grpc.Server, srv QpmServer) {
	s.RegisterService(&_Qpm_serviceDesc, srv)
}

func _Qpm_Ping_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(PingRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(QpmServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Qpm_Publish_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(PublishRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(QpmServer).Publish(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Qpm_GetDependencies_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(DependencyRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(QpmServer).GetDependencies(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Qpm_Search_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(SearchRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(QpmServer).Search(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Qpm_Login_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(LoginRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(QpmServer).Login(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Qpm_Info_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(InfoRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(QpmServer).Info(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Qpm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.Qpm",
	HandlerType: (*QpmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Qpm_Ping_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Qpm_Publish_Handler,
		},
		{
			MethodName: "GetDependencies",
			Handler:    _Qpm_GetDependencies_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Qpm_Search_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Qpm_Login_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Qpm_Info_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}