// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: CloudCostDTO.proto

package proto

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

// Historical demand data type.
type DemandType int32

const (
	// Consumption based data to be used
	DemandType_CONSUMPTION DemandType = 1
	// Allocation based data to be used
	DemandType_ALLOCATION DemandType = 2
)

// Enum value maps for DemandType.
var (
	DemandType_name = map[int32]string{
		1: "CONSUMPTION",
		2: "ALLOCATION",
	}
	DemandType_value = map[string]int32{
		"CONSUMPTION": 1,
		"ALLOCATION":  2,
	}
)

func (x DemandType) Enum() *DemandType {
	p := new(DemandType)
	*p = x
	return p
}

func (x DemandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DemandType) Descriptor() protoreflect.EnumDescriptor {
	return file_CloudCostDTO_proto_enumTypes[0].Descriptor()
}

func (DemandType) Type() protoreflect.EnumType {
	return &file_CloudCostDTO_proto_enumTypes[0]
}

func (x DemandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DemandType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DemandType(num)
	return nil
}

// Deprecated: Use DemandType.Descriptor instead.
func (DemandType) EnumDescriptor() ([]byte, []int) {
	return file_CloudCostDTO_proto_rawDescGZIP(), []int{0}
}

// The tenancy of an instance defines what hardware the instance is running on.
type Tenancy int32

const (
	// Instance runs on shared/default hardware.
	// This is typically the cheapest option.
	Tenancy_DEFAULT Tenancy = 1
	// Instance runs on single-tenant hardware.
	// That means your instance runs on a host that's separate from other customers,
	// but the host details are abstracted away, and you're not paying for the whole host.
	Tenancy_DEDICATED Tenancy = 2
	// Instance runs on a dedicated Host.
	// This means your instance runs on a specific host, and you are paying for the full host and
	// are responsible for managing it.
	Tenancy_HOST Tenancy = 3
)

// Enum value maps for Tenancy.
var (
	Tenancy_name = map[int32]string{
		1: "DEFAULT",
		2: "DEDICATED",
		3: "HOST",
	}
	Tenancy_value = map[string]int32{
		"DEFAULT":   1,
		"DEDICATED": 2,
		"HOST":      3,
	}
)

func (x Tenancy) Enum() *Tenancy {
	p := new(Tenancy)
	*p = x
	return p
}

func (x Tenancy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tenancy) Descriptor() protoreflect.EnumDescriptor {
	return file_CloudCostDTO_proto_enumTypes[1].Descriptor()
}

func (Tenancy) Type() protoreflect.EnumType {
	return &file_CloudCostDTO_proto_enumTypes[1]
}

func (x Tenancy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Tenancy) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Tenancy(num)
	return nil
}

// Deprecated: Use Tenancy.Descriptor instead.
func (Tenancy) EnumDescriptor() ([]byte, []int) {
	return file_CloudCostDTO_proto_rawDescGZIP(), []int{1}
}

// The supported operating systems. Keep in sync with com.vmturbo.mediation.hybrid.cloud.common.OsType.
type OSType int32

const (
	OSType_UNKNOWN_OS OSType = 0
	// Unix OS.
	OSType_LINUX                     OSType = 2
	OSType_SUSE                      OSType = 3
	OSType_RHEL                      OSType = 4
	OSType_LINUX_WITH_SQL_ENTERPRISE OSType = 5
	OSType_LINUX_WITH_SQL_STANDARD   OSType = 6
	OSType_LINUX_WITH_SQL_WEB        OSType = 7
	OSType_UBUNTU_PRO                OSType = 8
	// Windows OS.
	OSType_WINDOWS                     OSType = 20
	OSType_WINDOWS_WITH_SQL_STANDARD   OSType = 21
	OSType_WINDOWS_WITH_SQL_WEB        OSType = 22
	OSType_WINDOWS_WITH_SQL_ENTERPRISE OSType = 23
	OSType_WINDOWS_BYOL                OSType = 24
)

// Enum value maps for OSType.
var (
	OSType_name = map[int32]string{
		0:  "UNKNOWN_OS",
		2:  "LINUX",
		3:  "SUSE",
		4:  "RHEL",
		5:  "LINUX_WITH_SQL_ENTERPRISE",
		6:  "LINUX_WITH_SQL_STANDARD",
		7:  "LINUX_WITH_SQL_WEB",
		8:  "UBUNTU_PRO",
		20: "WINDOWS",
		21: "WINDOWS_WITH_SQL_STANDARD",
		22: "WINDOWS_WITH_SQL_WEB",
		23: "WINDOWS_WITH_SQL_ENTERPRISE",
		24: "WINDOWS_BYOL",
	}
	OSType_value = map[string]int32{
		"UNKNOWN_OS":                  0,
		"LINUX":                       2,
		"SUSE":                        3,
		"RHEL":                        4,
		"LINUX_WITH_SQL_ENTERPRISE":   5,
		"LINUX_WITH_SQL_STANDARD":     6,
		"LINUX_WITH_SQL_WEB":          7,
		"UBUNTU_PRO":                  8,
		"WINDOWS":                     20,
		"WINDOWS_WITH_SQL_STANDARD":   21,
		"WINDOWS_WITH_SQL_WEB":        22,
		"WINDOWS_WITH_SQL_ENTERPRISE": 23,
		"WINDOWS_BYOL":                24,
	}
)

func (x OSType) Enum() *OSType {
	p := new(OSType)
	*p = x
	return p
}

func (x OSType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OSType) Descriptor() protoreflect.EnumDescriptor {
	return file_CloudCostDTO_proto_enumTypes[2].Descriptor()
}

func (OSType) Type() protoreflect.EnumType {
	return &file_CloudCostDTO_proto_enumTypes[2]
}

func (x OSType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *OSType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = OSType(num)
	return nil
}

// Deprecated: Use OSType.Descriptor instead.
func (OSType) EnumDescriptor() ([]byte, []int) {
	return file_CloudCostDTO_proto_rawDescGZIP(), []int{2}
}

// The engine for a database tier.
// This is an enum to save on space - and also because
// the list of supported engines across cloud providers is pretty small.
type DatabaseEngine int32

const (
	DatabaseEngine_UNKNOWN    DatabaseEngine = 0
	DatabaseEngine_MYSQL      DatabaseEngine = 1
	DatabaseEngine_MARIADB    DatabaseEngine = 2
	DatabaseEngine_POSTGRESQL DatabaseEngine = 3
	DatabaseEngine_ORACLE     DatabaseEngine = 4
	DatabaseEngine_SQLSERVER  DatabaseEngine = 5
	// Deprecated: Do not use.
	DatabaseEngine_AURORA           DatabaseEngine = 6
	DatabaseEngine_AURORAMYSQL      DatabaseEngine = 7
	DatabaseEngine_AURORAPOSTGRESQL DatabaseEngine = 8
	DatabaseEngine_MONGO            DatabaseEngine = 9
	DatabaseEngine_SYBASE           DatabaseEngine = 10
)

// Enum value maps for DatabaseEngine.
var (
	DatabaseEngine_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "MYSQL",
		2:  "MARIADB",
		3:  "POSTGRESQL",
		4:  "ORACLE",
		5:  "SQLSERVER",
		6:  "AURORA",
		7:  "AURORAMYSQL",
		8:  "AURORAPOSTGRESQL",
		9:  "MONGO",
		10: "SYBASE",
	}
	DatabaseEngine_value = map[string]int32{
		"UNKNOWN":          0,
		"MYSQL":            1,
		"MARIADB":          2,
		"POSTGRESQL":       3,
		"ORACLE":           4,
		"SQLSERVER":        5,
		"AURORA":           6,
		"AURORAMYSQL":      7,
		"AURORAPOSTGRESQL": 8,
		"MONGO":            9,
		"SYBASE":           10,
	}
)

func (x DatabaseEngine) Enum() *DatabaseEngine {
	p := new(DatabaseEngine)
	*p = x
	return p
}

func (x DatabaseEngine) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatabaseEngine) Descriptor() protoreflect.EnumDescriptor {
	return file_CloudCostDTO_proto_enumTypes[3].Descriptor()
}

func (DatabaseEngine) Type() protoreflect.EnumType {
	return &file_CloudCostDTO_proto_enumTypes[3]
}

func (x DatabaseEngine) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DatabaseEngine) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DatabaseEngine(num)
	return nil
}

// Deprecated: Use DatabaseEngine.Descriptor instead.
func (DatabaseEngine) EnumDescriptor() ([]byte, []int) {
	return file_CloudCostDTO_proto_rawDescGZIP(), []int{3}
}

// LicenseModel describes all supported license models in cloud projects.
type LicenseModel int32

const (
	LicenseModel_BRING_YOUR_OWN_LICENSE LicenseModel = 0
	LicenseModel_LICENSE_INCLUDED       LicenseModel = 1
	LicenseModel_NO_LICENSE_REQUIRED    LicenseModel = 2
)

// Enum value maps for LicenseModel.
var (
	LicenseModel_name = map[int32]string{
		0: "BRING_YOUR_OWN_LICENSE",
		1: "LICENSE_INCLUDED",
		2: "NO_LICENSE_REQUIRED",
	}
	LicenseModel_value = map[string]int32{
		"BRING_YOUR_OWN_LICENSE": 0,
		"LICENSE_INCLUDED":       1,
		"NO_LICENSE_REQUIRED":    2,
	}
)

func (x LicenseModel) Enum() *LicenseModel {
	p := new(LicenseModel)
	*p = x
	return p
}

func (x LicenseModel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LicenseModel) Descriptor() protoreflect.EnumDescriptor {
	return file_CloudCostDTO_proto_enumTypes[4].Descriptor()
}

func (LicenseModel) Type() protoreflect.EnumType {
	return &file_CloudCostDTO_proto_enumTypes[4]
}

func (x LicenseModel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LicenseModel) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LicenseModel(num)
	return nil
}

// Deprecated: Use LicenseModel.Descriptor instead.
func (LicenseModel) EnumDescriptor() ([]byte, []int) {
	return file_CloudCostDTO_proto_rawDescGZIP(), []int{4}
}

// DeploymentType describes all supported deployment types by cloud probes.
type DeploymentType int32

const (
	DeploymentType_SINGLE_AZ DeploymentType = 0
	DeploymentType_MULTI_AZ  DeploymentType = 1
)

// Enum value maps for DeploymentType.
var (
	DeploymentType_name = map[int32]string{
		0: "SINGLE_AZ",
		1: "MULTI_AZ",
	}
	DeploymentType_value = map[string]int32{
		"SINGLE_AZ": 0,
		"MULTI_AZ":  1,
	}
)

func (x DeploymentType) Enum() *DeploymentType {
	p := new(DeploymentType)
	*p = x
	return p
}

func (x DeploymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_CloudCostDTO_proto_enumTypes[5].Descriptor()
}

func (DeploymentType) Type() protoreflect.EnumType {
	return &file_CloudCostDTO_proto_enumTypes[5]
}

func (x DeploymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DeploymentType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DeploymentType(num)
	return nil
}

// Deprecated: Use DeploymentType.Descriptor instead.
func (DeploymentType) EnumDescriptor() ([]byte, []int) {
	return file_CloudCostDTO_proto_rawDescGZIP(), []int{5}
}

// The edition of a database engine.
// The edition enum is closely related to the DatabaseEngine enum, and in the future it may be
// worth it to have a separate "database identifier" message that forbids illegal
// engine-edition combinations. For now, there are only two database engines with editions,
// so this seems manageable.
type DatabaseEdition int32

const (
	// Possible list of edition values.
	DatabaseEdition_NONE        DatabaseEdition = 0
	DatabaseEdition_ENTERPRISE  DatabaseEdition = 1
	DatabaseEdition_STANDARD    DatabaseEdition = 2
	DatabaseEdition_STANDARDONE DatabaseEdition = 3
	DatabaseEdition_STANDARDTWO DatabaseEdition = 4
	DatabaseEdition_WEB         DatabaseEdition = 12
	DatabaseEdition_EXPRESS     DatabaseEdition = 13
)

// Enum value maps for DatabaseEdition.
var (
	DatabaseEdition_name = map[int32]string{
		0:  "NONE",
		1:  "ENTERPRISE",
		2:  "STANDARD",
		3:  "STANDARDONE",
		4:  "STANDARDTWO",
		12: "WEB",
		13: "EXPRESS",
	}
	DatabaseEdition_value = map[string]int32{
		"NONE":        0,
		"ENTERPRISE":  1,
		"STANDARD":    2,
		"STANDARDONE": 3,
		"STANDARDTWO": 4,
		"WEB":         12,
		"EXPRESS":     13,
	}
)

func (x DatabaseEdition) Enum() *DatabaseEdition {
	p := new(DatabaseEdition)
	*p = x
	return p
}

func (x DatabaseEdition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatabaseEdition) Descriptor() protoreflect.EnumDescriptor {
	return file_CloudCostDTO_proto_enumTypes[6].Descriptor()
}

func (DatabaseEdition) Type() protoreflect.EnumType {
	return &file_CloudCostDTO_proto_enumTypes[6]
}

func (x DatabaseEdition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DatabaseEdition) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DatabaseEdition(num)
	return nil
}

// Deprecated: Use DatabaseEdition.Descriptor instead.
func (DatabaseEdition) EnumDescriptor() ([]byte, []int) {
	return file_CloudCostDTO_proto_rawDescGZIP(), []int{6}
}

type ReservedInstanceType_OfferingClass int32

const (
	// Most of the attributes of a standard reserved instance are "fixed" at the time it's
	// bought.
	ReservedInstanceType_STANDARD ReservedInstanceType_OfferingClass = 1
	// A convertible reserved instance can be exchanged for a different
	// instance type. Not all service providers offer convertible reserved
	// instances.
	ReservedInstanceType_CONVERTIBLE ReservedInstanceType_OfferingClass = 2
)

// Enum value maps for ReservedInstanceType_OfferingClass.
var (
	ReservedInstanceType_OfferingClass_name = map[int32]string{
		1: "STANDARD",
		2: "CONVERTIBLE",
	}
	ReservedInstanceType_OfferingClass_value = map[string]int32{
		"STANDARD":    1,
		"CONVERTIBLE": 2,
	}
)

func (x ReservedInstanceType_OfferingClass) Enum() *ReservedInstanceType_OfferingClass {
	p := new(ReservedInstanceType_OfferingClass)
	*p = x
	return p
}

func (x ReservedInstanceType_OfferingClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReservedInstanceType_OfferingClass) Descriptor() protoreflect.EnumDescriptor {
	return file_CloudCostDTO_proto_enumTypes[7].Descriptor()
}

func (ReservedInstanceType_OfferingClass) Type() protoreflect.EnumType {
	return &file_CloudCostDTO_proto_enumTypes[7]
}

func (x ReservedInstanceType_OfferingClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ReservedInstanceType_OfferingClass) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ReservedInstanceType_OfferingClass(num)
	return nil
}

// Deprecated: Use ReservedInstanceType_OfferingClass.Descriptor instead.
func (ReservedInstanceType_OfferingClass) EnumDescriptor() ([]byte, []int) {
	return file_CloudCostDTO_proto_rawDescGZIP(), []int{0, 0}
}

// Identifies the type of "reservation" for the instance, and the
// payment conditions.
type ReservedInstanceType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of offering.
	OfferingClass *ReservedInstanceType_OfferingClass `protobuf:"varint,1,opt,name=offering_class,json=offeringClass,enum=common_dto.ReservedInstanceType_OfferingClass" json:"offering_class,omitempty"`
	// The payment option for this reserved instance.
	PaymentOption *PaymentOption `protobuf:"varint,2,opt,name=payment_option,json=paymentOption,enum=common_dto.PaymentOption" json:"payment_option,omitempty"`
	// The term, in years.
	TermYears *uint32 `protobuf:"varint,3,opt,name=term_years,json=termYears" json:"term_years,omitempty"`
}

func (x *ReservedInstanceType) Reset() {
	*x = ReservedInstanceType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CloudCostDTO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservedInstanceType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservedInstanceType) ProtoMessage() {}

func (x *ReservedInstanceType) ProtoReflect() protoreflect.Message {
	mi := &file_CloudCostDTO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservedInstanceType.ProtoReflect.Descriptor instead.
func (*ReservedInstanceType) Descriptor() ([]byte, []int) {
	return file_CloudCostDTO_proto_rawDescGZIP(), []int{0}
}

func (x *ReservedInstanceType) GetOfferingClass() ReservedInstanceType_OfferingClass {
	if x != nil && x.OfferingClass != nil {
		return *x.OfferingClass
	}
	return ReservedInstanceType_STANDARD
}

func (x *ReservedInstanceType) GetPaymentOption() PaymentOption {
	if x != nil && x.PaymentOption != nil {
		return *x.PaymentOption
	}
	return PaymentOption_ALL_UPFRONT
}

func (x *ReservedInstanceType) GetTermYears() uint32 {
	if x != nil && x.TermYears != nil {
		return *x.TermYears
	}
	return 0
}

// The ReservedInstanceSpec describes a solution offered by the cloud provider that allows the customer
// to buy in advance a number of compute instances, for a discounted price. Usually those solutions
// have long terms like 1 or 3 years.
type ReservedInstanceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the reserved instance.
	Type *ReservedInstanceType `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// The tenancy of the reserved instance.
	Tenancy *Tenancy `protobuf:"varint,2,opt,name=tenancy,enum=common_dto.Tenancy" json:"tenancy,omitempty"`
	// The operating system of the reserved instance.
	Os *OSType `protobuf:"varint,3,opt,name=os,enum=common_dto.OSType" json:"os,omitempty"`
	// The entity profile of the reserved instance is using, such as t2.large.
	Tier *EntityDTO `protobuf:"bytes,4,opt,name=tier" json:"tier,omitempty"`
	// The region of the reserved instance.
	Region       *EntityDTO `protobuf:"bytes,5,opt,name=region" json:"region,omitempty"`
	SizeFlexible *bool      `protobuf:"varint,6,opt,name=size_flexible,json=sizeFlexible" json:"size_flexible,omitempty"`
	// Whether the reserved instance can be applied to any operating system. If true,
	// the os attribute of the RI spec will be ignored.
	PlatformFlexible *bool `protobuf:"varint,7,opt,name=platform_flexible,json=platformFlexible" json:"platform_flexible,omitempty"`
}

func (x *ReservedInstanceSpec) Reset() {
	*x = ReservedInstanceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CloudCostDTO_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservedInstanceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservedInstanceSpec) ProtoMessage() {}

func (x *ReservedInstanceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_CloudCostDTO_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservedInstanceSpec.ProtoReflect.Descriptor instead.
func (*ReservedInstanceSpec) Descriptor() ([]byte, []int) {
	return file_CloudCostDTO_proto_rawDescGZIP(), []int{1}
}

func (x *ReservedInstanceSpec) GetType() *ReservedInstanceType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ReservedInstanceSpec) GetTenancy() Tenancy {
	if x != nil && x.Tenancy != nil {
		return *x.Tenancy
	}
	return Tenancy_DEFAULT
}

func (x *ReservedInstanceSpec) GetOs() OSType {
	if x != nil && x.Os != nil {
		return *x.Os
	}
	return OSType_UNKNOWN_OS
}

func (x *ReservedInstanceSpec) GetTier() *EntityDTO {
	if x != nil {
		return x.Tier
	}
	return nil
}

func (x *ReservedInstanceSpec) GetRegion() *EntityDTO {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *ReservedInstanceSpec) GetSizeFlexible() bool {
	if x != nil && x.SizeFlexible != nil {
		return *x.SizeFlexible
	}
	return false
}

func (x *ReservedInstanceSpec) GetPlatformFlexible() bool {
	if x != nil && x.PlatformFlexible != nil {
		return *x.PlatformFlexible
	}
	return false
}

var File_CloudCostDTO_proto protoreflect.FileDescriptor

var file_CloudCostDTO_proto_rawDesc = []byte{
	0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x44, 0x54, 0x4f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f,
	0x1a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x54, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x55, 0x0a, 0x0e,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x0d, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x40, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x5f, 0x79, 0x65,
	0x61, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x65, 0x72, 0x6d, 0x59,
	0x65, 0x61, 0x72, 0x73, 0x22, 0x2e, 0x0a, 0x0d, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52,
	0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x54, 0x49, 0x42,
	0x4c, 0x45, 0x10, 0x02, 0x22, 0xcb, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x34, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74,
	0x6f, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x79, 0x12, 0x22, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x4f, 0x53, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74,
	0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x54, 0x4f, 0x52, 0x04, 0x74, 0x69, 0x65,
	0x72, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x54, 0x4f, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x66, 0x6c, 0x65, 0x78, 0x69, 0x62, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x69, 0x7a, 0x65, 0x46, 0x6c, 0x65,
	0x78, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x66, 0x6c, 0x65, 0x78, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x10, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x6c, 0x65, 0x78, 0x69, 0x62,
	0x6c, 0x65, 0x2a, 0x2d, 0x0a, 0x0a, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4c, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x02, 0x2a, 0x2f, 0x0a, 0x07, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x44,
	0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x53, 0x54,
	0x10, 0x03, 0x2a, 0x94, 0x02, 0x0a, 0x06, 0x4f, 0x53, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x0a, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4f, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x55, 0x53, 0x45,
	0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x48, 0x45, 0x4c, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19,
	0x4c, 0x49, 0x4e, 0x55, 0x58, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x51, 0x4c, 0x5f, 0x45,
	0x4e, 0x54, 0x45, 0x52, 0x50, 0x52, 0x49, 0x53, 0x45, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x4c,
	0x49, 0x4e, 0x55, 0x58, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x51, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x49, 0x4e, 0x55,
	0x58, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x51, 0x4c, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x07,
	0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x42, 0x55, 0x4e, 0x54, 0x55, 0x5f, 0x50, 0x52, 0x4f, 0x10, 0x08,
	0x12, 0x0b, 0x0a, 0x07, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x14, 0x12, 0x1d, 0x0a,
	0x19, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x51,
	0x4c, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x15, 0x12, 0x18, 0x0a, 0x14,
	0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x51, 0x4c,
	0x5f, 0x57, 0x45, 0x42, 0x10, 0x16, 0x12, 0x1f, 0x0a, 0x1b, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57,
	0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x51, 0x4c, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52,
	0x50, 0x52, 0x49, 0x53, 0x45, 0x10, 0x17, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x49, 0x4e, 0x44, 0x4f,
	0x57, 0x53, 0x5f, 0x42, 0x59, 0x4f, 0x4c, 0x10, 0x18, 0x2a, 0xae, 0x01, 0x0a, 0x0e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x59, 0x53,
	0x51, 0x4c, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x52, 0x49, 0x41, 0x44, 0x42, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x51, 0x4c, 0x10,
	0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x51, 0x4c, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x06,
	0x41, 0x55, 0x52, 0x4f, 0x52, 0x41, 0x10, 0x06, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x41, 0x55, 0x52, 0x4f, 0x52, 0x41, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x07, 0x12, 0x14, 0x0a,
	0x10, 0x41, 0x55, 0x52, 0x4f, 0x52, 0x41, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x51,
	0x4c, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x4e, 0x47, 0x4f, 0x10, 0x09, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x59, 0x42, 0x41, 0x53, 0x45, 0x10, 0x0a, 0x2a, 0x59, 0x0a, 0x0c, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x16, 0x42, 0x52,
	0x49, 0x4e, 0x47, 0x5f, 0x59, 0x4f, 0x55, 0x52, 0x5f, 0x4f, 0x57, 0x4e, 0x5f, 0x4c, 0x49, 0x43,
	0x45, 0x4e, 0x53, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53,
	0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x4e, 0x4f, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49,
	0x52, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x2d, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x49, 0x4e, 0x47, 0x4c,
	0x45, 0x5f, 0x41, 0x5a, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f,
	0x41, 0x5a, 0x10, 0x01, 0x2a, 0x71, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x45, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x50, 0x52, 0x49, 0x53, 0x45, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x03,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x54, 0x57, 0x4f, 0x10,
	0x04, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x0c, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58,
	0x50, 0x52, 0x45, 0x53, 0x53, 0x10, 0x0d, 0x42, 0x4f, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x6d, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x6e, 0x6f, 0x6d, 0x69,
	0x63, 0x2f, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_CloudCostDTO_proto_rawDescOnce sync.Once
	file_CloudCostDTO_proto_rawDescData = file_CloudCostDTO_proto_rawDesc
)

func file_CloudCostDTO_proto_rawDescGZIP() []byte {
	file_CloudCostDTO_proto_rawDescOnce.Do(func() {
		file_CloudCostDTO_proto_rawDescData = protoimpl.X.CompressGZIP(file_CloudCostDTO_proto_rawDescData)
	})
	return file_CloudCostDTO_proto_rawDescData
}

var file_CloudCostDTO_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_CloudCostDTO_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_CloudCostDTO_proto_goTypes = []interface{}{
	(DemandType)(0),                         // 0: common_dto.DemandType
	(Tenancy)(0),                            // 1: common_dto.Tenancy
	(OSType)(0),                             // 2: common_dto.OSType
	(DatabaseEngine)(0),                     // 3: common_dto.DatabaseEngine
	(LicenseModel)(0),                       // 4: common_dto.LicenseModel
	(DeploymentType)(0),                     // 5: common_dto.DeploymentType
	(DatabaseEdition)(0),                    // 6: common_dto.DatabaseEdition
	(ReservedInstanceType_OfferingClass)(0), // 7: common_dto.ReservedInstanceType.OfferingClass
	(*ReservedInstanceType)(nil),            // 8: common_dto.ReservedInstanceType
	(*ReservedInstanceSpec)(nil),            // 9: common_dto.ReservedInstanceSpec
	(PaymentOption)(0),                      // 10: common_dto.PaymentOption
	(*EntityDTO)(nil),                       // 11: common_dto.EntityDTO
}
var file_CloudCostDTO_proto_depIdxs = []int32{
	7,  // 0: common_dto.ReservedInstanceType.offering_class:type_name -> common_dto.ReservedInstanceType.OfferingClass
	10, // 1: common_dto.ReservedInstanceType.payment_option:type_name -> common_dto.PaymentOption
	8,  // 2: common_dto.ReservedInstanceSpec.type:type_name -> common_dto.ReservedInstanceType
	1,  // 3: common_dto.ReservedInstanceSpec.tenancy:type_name -> common_dto.Tenancy
	2,  // 4: common_dto.ReservedInstanceSpec.os:type_name -> common_dto.OSType
	11, // 5: common_dto.ReservedInstanceSpec.tier:type_name -> common_dto.EntityDTO
	11, // 6: common_dto.ReservedInstanceSpec.region:type_name -> common_dto.EntityDTO
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_CloudCostDTO_proto_init() }
func file_CloudCostDTO_proto_init() {
	if File_CloudCostDTO_proto != nil {
		return
	}
	file_CommonDTO_proto_init()
	file_CommonCost_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CloudCostDTO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservedInstanceType); i {
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
		file_CloudCostDTO_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservedInstanceSpec); i {
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
			RawDescriptor: file_CloudCostDTO_proto_rawDesc,
			NumEnums:      8,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CloudCostDTO_proto_goTypes,
		DependencyIndexes: file_CloudCostDTO_proto_depIdxs,
		EnumInfos:         file_CloudCostDTO_proto_enumTypes,
		MessageInfos:      file_CloudCostDTO_proto_msgTypes,
	}.Build()
	File_CloudCostDTO_proto = out.File
	file_CloudCostDTO_proto_rawDesc = nil
	file_CloudCostDTO_proto_goTypes = nil
	file_CloudCostDTO_proto_depIdxs = nil
}
