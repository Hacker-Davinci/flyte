# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/types.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19\x66lyteidl/core/types.proto\x12\rflyteidl.core\x1a\x1cgoogle/protobuf/struct.proto\"\xa1\x02\n\nSchemaType\x12@\n\x07\x63olumns\x18\x03 \x03(\x0b\x32&.flyteidl.core.SchemaType.SchemaColumnR\x07\x63olumns\x1a\xd0\x01\n\x0cSchemaColumn\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12K\n\x04type\x18\x02 \x01(\x0e\x32\x37.flyteidl.core.SchemaType.SchemaColumn.SchemaColumnTypeR\x04type\"_\n\x10SchemaColumnType\x12\x0b\n\x07INTEGER\x10\x00\x12\t\n\x05\x46LOAT\x10\x01\x12\n\n\x06STRING\x10\x02\x12\x0b\n\x07\x42OOLEAN\x10\x03\x12\x0c\n\x08\x44\x41TETIME\x10\x04\x12\x0c\n\x08\x44URATION\x10\x05\"\xc7\x02\n\x15StructuredDatasetType\x12L\n\x07\x63olumns\x18\x01 \x03(\x0b\x32\x32.flyteidl.core.StructuredDatasetType.DatasetColumnR\x07\x63olumns\x12\x16\n\x06\x66ormat\x18\x02 \x01(\tR\x06\x66ormat\x12\x30\n\x14\x65xternal_schema_type\x18\x03 \x01(\tR\x12\x65xternalSchemaType\x12\x32\n\x15\x65xternal_schema_bytes\x18\x04 \x01(\x0cR\x13\x65xternalSchemaBytes\x1a\x62\n\rDatasetColumn\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12=\n\x0cliteral_type\x18\x02 \x01(\x0b\x32\x1a.flyteidl.core.LiteralTypeR\x0bliteralType\"\xa7\x01\n\x08\x42lobType\x12\x16\n\x06\x66ormat\x18\x01 \x01(\tR\x06\x66ormat\x12R\n\x0e\x64imensionality\x18\x02 \x01(\x0e\x32*.flyteidl.core.BlobType.BlobDimensionalityR\x0e\x64imensionality\"/\n\x12\x42lobDimensionality\x12\n\n\x06SINGLE\x10\x00\x12\r\n\tMULTIPART\x10\x01\"\"\n\x08\x45numType\x12\x16\n\x06values\x18\x01 \x03(\tR\x06values\"C\n\tUnionType\x12\x36\n\x08variants\x18\x01 \x03(\x0b\x32\x1a.flyteidl.core.LiteralTypeR\x08variants\"!\n\rTypeStructure\x12\x10\n\x03tag\x18\x01 \x01(\tR\x03tag\"K\n\x0eTypeAnnotation\x12\x39\n\x0b\x61nnotations\x18\x01 \x01(\x0b\x32\x17.google.protobuf.StructR\x0b\x61nnotations\"\xbc\x05\n\x0bLiteralType\x12\x33\n\x06simple\x18\x01 \x01(\x0e\x32\x19.flyteidl.core.SimpleTypeH\x00R\x06simple\x12\x33\n\x06schema\x18\x02 \x01(\x0b\x32\x19.flyteidl.core.SchemaTypeH\x00R\x06schema\x12\x45\n\x0f\x63ollection_type\x18\x03 \x01(\x0b\x32\x1a.flyteidl.core.LiteralTypeH\x00R\x0e\x63ollectionType\x12\x42\n\x0emap_value_type\x18\x04 \x01(\x0b\x32\x1a.flyteidl.core.LiteralTypeH\x00R\x0cmapValueType\x12-\n\x04\x62lob\x18\x05 \x01(\x0b\x32\x17.flyteidl.core.BlobTypeH\x00R\x04\x62lob\x12\x36\n\tenum_type\x18\x07 \x01(\x0b\x32\x17.flyteidl.core.EnumTypeH\x00R\x08\x65numType\x12^\n\x17structured_dataset_type\x18\x08 \x01(\x0b\x32$.flyteidl.core.StructuredDatasetTypeH\x00R\x15structuredDatasetType\x12\x39\n\nunion_type\x18\n \x01(\x0b\x32\x18.flyteidl.core.UnionTypeH\x00R\tunionType\x12\x33\n\x08metadata\x18\x06 \x01(\x0b\x32\x17.google.protobuf.StructR\x08metadata\x12=\n\nannotation\x18\t \x01(\x0b\x32\x1d.flyteidl.core.TypeAnnotationR\nannotation\x12:\n\tstructure\x18\x0b \x01(\x0b\x32\x1c.flyteidl.core.TypeStructureR\tstructureB\x06\n\x04type\"z\n\x0fOutputReference\x12\x17\n\x07node_id\x18\x01 \x01(\tR\x06nodeId\x12\x10\n\x03var\x18\x02 \x01(\tR\x03var\x12<\n\tattr_path\x18\x03 \x03(\x0b\x32\x1f.flyteidl.core.PromiseAttributeR\x08\x61ttrPath\"_\n\x10PromiseAttribute\x12#\n\x0cstring_value\x18\x01 \x01(\tH\x00R\x0bstringValue\x12\x1d\n\tint_value\x18\x02 \x01(\x05H\x00R\x08intValueB\x07\n\x05value\"G\n\x05\x45rror\x12$\n\x0e\x66\x61iled_node_id\x18\x01 \x01(\tR\x0c\x66\x61iledNodeId\x12\x18\n\x07message\x18\x02 \x01(\tR\x07message*\x86\x01\n\nSimpleType\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07INTEGER\x10\x01\x12\t\n\x05\x46LOAT\x10\x02\x12\n\n\x06STRING\x10\x03\x12\x0b\n\x07\x42OOLEAN\x10\x04\x12\x0c\n\x08\x44\x41TETIME\x10\x05\x12\x0c\n\x08\x44URATION\x10\x06\x12\n\n\x06\x42INARY\x10\x07\x12\t\n\x05\x45RROR\x10\x08\x12\n\n\x06STRUCT\x10\tB\xb0\x01\n\x11\x63om.flyteidl.coreB\nTypesProtoP\x01Z:github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core\xa2\x02\x03\x46\x43X\xaa\x02\rFlyteidl.Core\xca\x02\rFlyteidl\\Core\xe2\x02\x19\x46lyteidl\\Core\\GPBMetadata\xea\x02\x0e\x46lyteidl::Coreb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.core.types_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\021com.flyteidl.coreB\nTypesProtoP\001Z:github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core\242\002\003FCX\252\002\rFlyteidl.Core\312\002\rFlyteidl\\Core\342\002\031Flyteidl\\Core\\GPBMetadata\352\002\016Flyteidl::Core'
  _globals['_SIMPLETYPE']._serialized_start=2081
  _globals['_SIMPLETYPE']._serialized_end=2215
  _globals['_SCHEMATYPE']._serialized_start=75
  _globals['_SCHEMATYPE']._serialized_end=364
  _globals['_SCHEMATYPE_SCHEMACOLUMN']._serialized_start=156
  _globals['_SCHEMATYPE_SCHEMACOLUMN']._serialized_end=364
  _globals['_SCHEMATYPE_SCHEMACOLUMN_SCHEMACOLUMNTYPE']._serialized_start=269
  _globals['_SCHEMATYPE_SCHEMACOLUMN_SCHEMACOLUMNTYPE']._serialized_end=364
  _globals['_STRUCTUREDDATASETTYPE']._serialized_start=367
  _globals['_STRUCTUREDDATASETTYPE']._serialized_end=694
  _globals['_STRUCTUREDDATASETTYPE_DATASETCOLUMN']._serialized_start=596
  _globals['_STRUCTUREDDATASETTYPE_DATASETCOLUMN']._serialized_end=694
  _globals['_BLOBTYPE']._serialized_start=697
  _globals['_BLOBTYPE']._serialized_end=864
  _globals['_BLOBTYPE_BLOBDIMENSIONALITY']._serialized_start=817
  _globals['_BLOBTYPE_BLOBDIMENSIONALITY']._serialized_end=864
  _globals['_ENUMTYPE']._serialized_start=866
  _globals['_ENUMTYPE']._serialized_end=900
  _globals['_UNIONTYPE']._serialized_start=902
  _globals['_UNIONTYPE']._serialized_end=969
  _globals['_TYPESTRUCTURE']._serialized_start=971
  _globals['_TYPESTRUCTURE']._serialized_end=1004
  _globals['_TYPEANNOTATION']._serialized_start=1006
  _globals['_TYPEANNOTATION']._serialized_end=1081
  _globals['_LITERALTYPE']._serialized_start=1084
  _globals['_LITERALTYPE']._serialized_end=1784
  _globals['_OUTPUTREFERENCE']._serialized_start=1786
  _globals['_OUTPUTREFERENCE']._serialized_end=1908
  _globals['_PROMISEATTRIBUTE']._serialized_start=1910
  _globals['_PROMISEATTRIBUTE']._serialized_end=2005
  _globals['_ERROR']._serialized_start=2007
  _globals['_ERROR']._serialized_end=2078
# @@protoc_insertion_point(module_scope)
