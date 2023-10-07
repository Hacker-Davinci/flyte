// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/plugins/kubeflow/tensorflow.proto

#ifndef PROTOBUF_INCLUDED_flyteidl_2fplugins_2fkubeflow_2ftensorflow_2eproto
#define PROTOBUF_INCLUDED_flyteidl_2fplugins_2fkubeflow_2ftensorflow_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
#include "flyteidl/core/tasks.pb.h"
#include "flyteidl/plugins/kubeflow/common.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_flyteidl_2fplugins_2fkubeflow_2ftensorflow_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_flyteidl_2fplugins_2fkubeflow_2ftensorflow_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[2]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_flyteidl_2fplugins_2fkubeflow_2ftensorflow_2eproto();
namespace flyteidl {
namespace plugins {
namespace kubeflow {
class DistributedTensorflowTrainingReplicaSpec;
class DistributedTensorflowTrainingReplicaSpecDefaultTypeInternal;
extern DistributedTensorflowTrainingReplicaSpecDefaultTypeInternal _DistributedTensorflowTrainingReplicaSpec_default_instance_;
class DistributedTensorflowTrainingTask;
class DistributedTensorflowTrainingTaskDefaultTypeInternal;
extern DistributedTensorflowTrainingTaskDefaultTypeInternal _DistributedTensorflowTrainingTask_default_instance_;
}  // namespace kubeflow
}  // namespace plugins
}  // namespace flyteidl
namespace google {
namespace protobuf {
template<> ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* Arena::CreateMaybeMessage<::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec>(Arena*);
template<> ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingTask* Arena::CreateMaybeMessage<::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingTask>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace flyteidl {
namespace plugins {
namespace kubeflow {

// ===================================================================

class DistributedTensorflowTrainingTask final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask) */ {
 public:
  DistributedTensorflowTrainingTask();
  virtual ~DistributedTensorflowTrainingTask();

  DistributedTensorflowTrainingTask(const DistributedTensorflowTrainingTask& from);

  inline DistributedTensorflowTrainingTask& operator=(const DistributedTensorflowTrainingTask& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  DistributedTensorflowTrainingTask(DistributedTensorflowTrainingTask&& from) noexcept
    : DistributedTensorflowTrainingTask() {
    *this = ::std::move(from);
  }

  inline DistributedTensorflowTrainingTask& operator=(DistributedTensorflowTrainingTask&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const DistributedTensorflowTrainingTask& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const DistributedTensorflowTrainingTask* internal_default_instance() {
    return reinterpret_cast<const DistributedTensorflowTrainingTask*>(
               &_DistributedTensorflowTrainingTask_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(DistributedTensorflowTrainingTask* other);
  friend void swap(DistributedTensorflowTrainingTask& a, DistributedTensorflowTrainingTask& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline DistributedTensorflowTrainingTask* New() const final {
    return CreateMaybeMessage<DistributedTensorflowTrainingTask>(nullptr);
  }

  DistributedTensorflowTrainingTask* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<DistributedTensorflowTrainingTask>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const DistributedTensorflowTrainingTask& from);
  void MergeFrom(const DistributedTensorflowTrainingTask& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(DistributedTensorflowTrainingTask* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // .flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec worker_replicas = 1;
  bool has_worker_replicas() const;
  void clear_worker_replicas();
  static const int kWorkerReplicasFieldNumber = 1;
  const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec& worker_replicas() const;
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* release_worker_replicas();
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* mutable_worker_replicas();
  void set_allocated_worker_replicas(::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* worker_replicas);

  // .flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec ps_replicas = 2;
  bool has_ps_replicas() const;
  void clear_ps_replicas();
  static const int kPsReplicasFieldNumber = 2;
  const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec& ps_replicas() const;
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* release_ps_replicas();
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* mutable_ps_replicas();
  void set_allocated_ps_replicas(::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* ps_replicas);

  // .flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec chief_replicas = 3;
  bool has_chief_replicas() const;
  void clear_chief_replicas();
  static const int kChiefReplicasFieldNumber = 3;
  const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec& chief_replicas() const;
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* release_chief_replicas();
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* mutable_chief_replicas();
  void set_allocated_chief_replicas(::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* chief_replicas);

  // .flyteidl.plugins.kubeflow.RunPolicy run_policy = 4;
  bool has_run_policy() const;
  void clear_run_policy();
  static const int kRunPolicyFieldNumber = 4;
  const ::flyteidl::plugins::kubeflow::RunPolicy& run_policy() const;
  ::flyteidl::plugins::kubeflow::RunPolicy* release_run_policy();
  ::flyteidl::plugins::kubeflow::RunPolicy* mutable_run_policy();
  void set_allocated_run_policy(::flyteidl::plugins::kubeflow::RunPolicy* run_policy);

  // .flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec evaluator_replicas = 5;
  bool has_evaluator_replicas() const;
  void clear_evaluator_replicas();
  static const int kEvaluatorReplicasFieldNumber = 5;
  const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec& evaluator_replicas() const;
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* release_evaluator_replicas();
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* mutable_evaluator_replicas();
  void set_allocated_evaluator_replicas(::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* evaluator_replicas);

  // @@protoc_insertion_point(class_scope:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* worker_replicas_;
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* ps_replicas_;
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* chief_replicas_;
  ::flyteidl::plugins::kubeflow::RunPolicy* run_policy_;
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* evaluator_replicas_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fplugins_2fkubeflow_2ftensorflow_2eproto;
};
// -------------------------------------------------------------------

class DistributedTensorflowTrainingReplicaSpec final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec) */ {
 public:
  DistributedTensorflowTrainingReplicaSpec();
  virtual ~DistributedTensorflowTrainingReplicaSpec();

  DistributedTensorflowTrainingReplicaSpec(const DistributedTensorflowTrainingReplicaSpec& from);

  inline DistributedTensorflowTrainingReplicaSpec& operator=(const DistributedTensorflowTrainingReplicaSpec& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  DistributedTensorflowTrainingReplicaSpec(DistributedTensorflowTrainingReplicaSpec&& from) noexcept
    : DistributedTensorflowTrainingReplicaSpec() {
    *this = ::std::move(from);
  }

  inline DistributedTensorflowTrainingReplicaSpec& operator=(DistributedTensorflowTrainingReplicaSpec&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const DistributedTensorflowTrainingReplicaSpec& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const DistributedTensorflowTrainingReplicaSpec* internal_default_instance() {
    return reinterpret_cast<const DistributedTensorflowTrainingReplicaSpec*>(
               &_DistributedTensorflowTrainingReplicaSpec_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  void Swap(DistributedTensorflowTrainingReplicaSpec* other);
  friend void swap(DistributedTensorflowTrainingReplicaSpec& a, DistributedTensorflowTrainingReplicaSpec& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline DistributedTensorflowTrainingReplicaSpec* New() const final {
    return CreateMaybeMessage<DistributedTensorflowTrainingReplicaSpec>(nullptr);
  }

  DistributedTensorflowTrainingReplicaSpec* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<DistributedTensorflowTrainingReplicaSpec>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const DistributedTensorflowTrainingReplicaSpec& from);
  void MergeFrom(const DistributedTensorflowTrainingReplicaSpec& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(DistributedTensorflowTrainingReplicaSpec* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string image = 2;
  void clear_image();
  static const int kImageFieldNumber = 2;
  const ::std::string& image() const;
  void set_image(const ::std::string& value);
  #if LANG_CXX11
  void set_image(::std::string&& value);
  #endif
  void set_image(const char* value);
  void set_image(const char* value, size_t size);
  ::std::string* mutable_image();
  ::std::string* release_image();
  void set_allocated_image(::std::string* image);

  // .flyteidl.core.Resources resources = 3;
  bool has_resources() const;
  void clear_resources();
  static const int kResourcesFieldNumber = 3;
  const ::flyteidl::core::Resources& resources() const;
  ::flyteidl::core::Resources* release_resources();
  ::flyteidl::core::Resources* mutable_resources();
  void set_allocated_resources(::flyteidl::core::Resources* resources);

  // int32 replicas = 1;
  void clear_replicas();
  static const int kReplicasFieldNumber = 1;
  ::google::protobuf::int32 replicas() const;
  void set_replicas(::google::protobuf::int32 value);

  // .flyteidl.plugins.kubeflow.RestartPolicy restart_policy = 4;
  void clear_restart_policy();
  static const int kRestartPolicyFieldNumber = 4;
  ::flyteidl::plugins::kubeflow::RestartPolicy restart_policy() const;
  void set_restart_policy(::flyteidl::plugins::kubeflow::RestartPolicy value);

  // @@protoc_insertion_point(class_scope:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr image_;
  ::flyteidl::core::Resources* resources_;
  ::google::protobuf::int32 replicas_;
  int restart_policy_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fplugins_2fkubeflow_2ftensorflow_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// DistributedTensorflowTrainingTask

// .flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec worker_replicas = 1;
inline bool DistributedTensorflowTrainingTask::has_worker_replicas() const {
  return this != internal_default_instance() && worker_replicas_ != nullptr;
}
inline void DistributedTensorflowTrainingTask::clear_worker_replicas() {
  if (GetArenaNoVirtual() == nullptr && worker_replicas_ != nullptr) {
    delete worker_replicas_;
  }
  worker_replicas_ = nullptr;
}
inline const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec& DistributedTensorflowTrainingTask::worker_replicas() const {
  const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* p = worker_replicas_;
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.worker_replicas)
  return p != nullptr ? *p : *reinterpret_cast<const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec*>(
      &::flyteidl::plugins::kubeflow::_DistributedTensorflowTrainingReplicaSpec_default_instance_);
}
inline ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* DistributedTensorflowTrainingTask::release_worker_replicas() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.worker_replicas)
  
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* temp = worker_replicas_;
  worker_replicas_ = nullptr;
  return temp;
}
inline ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* DistributedTensorflowTrainingTask::mutable_worker_replicas() {
  
  if (worker_replicas_ == nullptr) {
    auto* p = CreateMaybeMessage<::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec>(GetArenaNoVirtual());
    worker_replicas_ = p;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.worker_replicas)
  return worker_replicas_;
}
inline void DistributedTensorflowTrainingTask::set_allocated_worker_replicas(::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* worker_replicas) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete worker_replicas_;
  }
  if (worker_replicas) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      worker_replicas = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, worker_replicas, submessage_arena);
    }
    
  } else {
    
  }
  worker_replicas_ = worker_replicas;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.worker_replicas)
}

// .flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec ps_replicas = 2;
inline bool DistributedTensorflowTrainingTask::has_ps_replicas() const {
  return this != internal_default_instance() && ps_replicas_ != nullptr;
}
inline void DistributedTensorflowTrainingTask::clear_ps_replicas() {
  if (GetArenaNoVirtual() == nullptr && ps_replicas_ != nullptr) {
    delete ps_replicas_;
  }
  ps_replicas_ = nullptr;
}
inline const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec& DistributedTensorflowTrainingTask::ps_replicas() const {
  const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* p = ps_replicas_;
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.ps_replicas)
  return p != nullptr ? *p : *reinterpret_cast<const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec*>(
      &::flyteidl::plugins::kubeflow::_DistributedTensorflowTrainingReplicaSpec_default_instance_);
}
inline ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* DistributedTensorflowTrainingTask::release_ps_replicas() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.ps_replicas)
  
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* temp = ps_replicas_;
  ps_replicas_ = nullptr;
  return temp;
}
inline ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* DistributedTensorflowTrainingTask::mutable_ps_replicas() {
  
  if (ps_replicas_ == nullptr) {
    auto* p = CreateMaybeMessage<::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec>(GetArenaNoVirtual());
    ps_replicas_ = p;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.ps_replicas)
  return ps_replicas_;
}
inline void DistributedTensorflowTrainingTask::set_allocated_ps_replicas(::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* ps_replicas) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete ps_replicas_;
  }
  if (ps_replicas) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      ps_replicas = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, ps_replicas, submessage_arena);
    }
    
  } else {
    
  }
  ps_replicas_ = ps_replicas;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.ps_replicas)
}

// .flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec chief_replicas = 3;
inline bool DistributedTensorflowTrainingTask::has_chief_replicas() const {
  return this != internal_default_instance() && chief_replicas_ != nullptr;
}
inline void DistributedTensorflowTrainingTask::clear_chief_replicas() {
  if (GetArenaNoVirtual() == nullptr && chief_replicas_ != nullptr) {
    delete chief_replicas_;
  }
  chief_replicas_ = nullptr;
}
inline const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec& DistributedTensorflowTrainingTask::chief_replicas() const {
  const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* p = chief_replicas_;
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.chief_replicas)
  return p != nullptr ? *p : *reinterpret_cast<const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec*>(
      &::flyteidl::plugins::kubeflow::_DistributedTensorflowTrainingReplicaSpec_default_instance_);
}
inline ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* DistributedTensorflowTrainingTask::release_chief_replicas() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.chief_replicas)
  
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* temp = chief_replicas_;
  chief_replicas_ = nullptr;
  return temp;
}
inline ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* DistributedTensorflowTrainingTask::mutable_chief_replicas() {
  
  if (chief_replicas_ == nullptr) {
    auto* p = CreateMaybeMessage<::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec>(GetArenaNoVirtual());
    chief_replicas_ = p;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.chief_replicas)
  return chief_replicas_;
}
inline void DistributedTensorflowTrainingTask::set_allocated_chief_replicas(::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* chief_replicas) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete chief_replicas_;
  }
  if (chief_replicas) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      chief_replicas = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, chief_replicas, submessage_arena);
    }
    
  } else {
    
  }
  chief_replicas_ = chief_replicas;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.chief_replicas)
}

// .flyteidl.plugins.kubeflow.RunPolicy run_policy = 4;
inline bool DistributedTensorflowTrainingTask::has_run_policy() const {
  return this != internal_default_instance() && run_policy_ != nullptr;
}
inline const ::flyteidl::plugins::kubeflow::RunPolicy& DistributedTensorflowTrainingTask::run_policy() const {
  const ::flyteidl::plugins::kubeflow::RunPolicy* p = run_policy_;
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.run_policy)
  return p != nullptr ? *p : *reinterpret_cast<const ::flyteidl::plugins::kubeflow::RunPolicy*>(
      &::flyteidl::plugins::kubeflow::_RunPolicy_default_instance_);
}
inline ::flyteidl::plugins::kubeflow::RunPolicy* DistributedTensorflowTrainingTask::release_run_policy() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.run_policy)
  
  ::flyteidl::plugins::kubeflow::RunPolicy* temp = run_policy_;
  run_policy_ = nullptr;
  return temp;
}
inline ::flyteidl::plugins::kubeflow::RunPolicy* DistributedTensorflowTrainingTask::mutable_run_policy() {
  
  if (run_policy_ == nullptr) {
    auto* p = CreateMaybeMessage<::flyteidl::plugins::kubeflow::RunPolicy>(GetArenaNoVirtual());
    run_policy_ = p;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.run_policy)
  return run_policy_;
}
inline void DistributedTensorflowTrainingTask::set_allocated_run_policy(::flyteidl::plugins::kubeflow::RunPolicy* run_policy) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::google::protobuf::MessageLite*>(run_policy_);
  }
  if (run_policy) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      run_policy = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, run_policy, submessage_arena);
    }
    
  } else {
    
  }
  run_policy_ = run_policy;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.run_policy)
}

// .flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec evaluator_replicas = 5;
inline bool DistributedTensorflowTrainingTask::has_evaluator_replicas() const {
  return this != internal_default_instance() && evaluator_replicas_ != nullptr;
}
inline void DistributedTensorflowTrainingTask::clear_evaluator_replicas() {
  if (GetArenaNoVirtual() == nullptr && evaluator_replicas_ != nullptr) {
    delete evaluator_replicas_;
  }
  evaluator_replicas_ = nullptr;
}
inline const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec& DistributedTensorflowTrainingTask::evaluator_replicas() const {
  const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* p = evaluator_replicas_;
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.evaluator_replicas)
  return p != nullptr ? *p : *reinterpret_cast<const ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec*>(
      &::flyteidl::plugins::kubeflow::_DistributedTensorflowTrainingReplicaSpec_default_instance_);
}
inline ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* DistributedTensorflowTrainingTask::release_evaluator_replicas() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.evaluator_replicas)
  
  ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* temp = evaluator_replicas_;
  evaluator_replicas_ = nullptr;
  return temp;
}
inline ::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* DistributedTensorflowTrainingTask::mutable_evaluator_replicas() {
  
  if (evaluator_replicas_ == nullptr) {
    auto* p = CreateMaybeMessage<::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec>(GetArenaNoVirtual());
    evaluator_replicas_ = p;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.evaluator_replicas)
  return evaluator_replicas_;
}
inline void DistributedTensorflowTrainingTask::set_allocated_evaluator_replicas(::flyteidl::plugins::kubeflow::DistributedTensorflowTrainingReplicaSpec* evaluator_replicas) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete evaluator_replicas_;
  }
  if (evaluator_replicas) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      evaluator_replicas = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, evaluator_replicas, submessage_arena);
    }
    
  } else {
    
  }
  evaluator_replicas_ = evaluator_replicas;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask.evaluator_replicas)
}

// -------------------------------------------------------------------

// DistributedTensorflowTrainingReplicaSpec

// int32 replicas = 1;
inline void DistributedTensorflowTrainingReplicaSpec::clear_replicas() {
  replicas_ = 0;
}
inline ::google::protobuf::int32 DistributedTensorflowTrainingReplicaSpec::replicas() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.replicas)
  return replicas_;
}
inline void DistributedTensorflowTrainingReplicaSpec::set_replicas(::google::protobuf::int32 value) {
  
  replicas_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.replicas)
}

// string image = 2;
inline void DistributedTensorflowTrainingReplicaSpec::clear_image() {
  image_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& DistributedTensorflowTrainingReplicaSpec::image() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.image)
  return image_.GetNoArena();
}
inline void DistributedTensorflowTrainingReplicaSpec::set_image(const ::std::string& value) {
  
  image_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.image)
}
#if LANG_CXX11
inline void DistributedTensorflowTrainingReplicaSpec::set_image(::std::string&& value) {
  
  image_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.image)
}
#endif
inline void DistributedTensorflowTrainingReplicaSpec::set_image(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  image_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.image)
}
inline void DistributedTensorflowTrainingReplicaSpec::set_image(const char* value, size_t size) {
  
  image_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.image)
}
inline ::std::string* DistributedTensorflowTrainingReplicaSpec::mutable_image() {
  
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.image)
  return image_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* DistributedTensorflowTrainingReplicaSpec::release_image() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.image)
  
  return image_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void DistributedTensorflowTrainingReplicaSpec::set_allocated_image(::std::string* image) {
  if (image != nullptr) {
    
  } else {
    
  }
  image_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), image);
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.image)
}

// .flyteidl.core.Resources resources = 3;
inline bool DistributedTensorflowTrainingReplicaSpec::has_resources() const {
  return this != internal_default_instance() && resources_ != nullptr;
}
inline const ::flyteidl::core::Resources& DistributedTensorflowTrainingReplicaSpec::resources() const {
  const ::flyteidl::core::Resources* p = resources_;
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.resources)
  return p != nullptr ? *p : *reinterpret_cast<const ::flyteidl::core::Resources*>(
      &::flyteidl::core::_Resources_default_instance_);
}
inline ::flyteidl::core::Resources* DistributedTensorflowTrainingReplicaSpec::release_resources() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.resources)
  
  ::flyteidl::core::Resources* temp = resources_;
  resources_ = nullptr;
  return temp;
}
inline ::flyteidl::core::Resources* DistributedTensorflowTrainingReplicaSpec::mutable_resources() {
  
  if (resources_ == nullptr) {
    auto* p = CreateMaybeMessage<::flyteidl::core::Resources>(GetArenaNoVirtual());
    resources_ = p;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.resources)
  return resources_;
}
inline void DistributedTensorflowTrainingReplicaSpec::set_allocated_resources(::flyteidl::core::Resources* resources) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::google::protobuf::MessageLite*>(resources_);
  }
  if (resources) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      resources = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, resources, submessage_arena);
    }
    
  } else {
    
  }
  resources_ = resources;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.resources)
}

// .flyteidl.plugins.kubeflow.RestartPolicy restart_policy = 4;
inline void DistributedTensorflowTrainingReplicaSpec::clear_restart_policy() {
  restart_policy_ = 0;
}
inline ::flyteidl::plugins::kubeflow::RestartPolicy DistributedTensorflowTrainingReplicaSpec::restart_policy() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.restart_policy)
  return static_cast< ::flyteidl::plugins::kubeflow::RestartPolicy >(restart_policy_);
}
inline void DistributedTensorflowTrainingReplicaSpec::set_restart_policy(::flyteidl::plugins::kubeflow::RestartPolicy value) {
  
  restart_policy_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec.restart_policy)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace kubeflow
}  // namespace plugins
}  // namespace flyteidl

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_flyteidl_2fplugins_2fkubeflow_2ftensorflow_2eproto
