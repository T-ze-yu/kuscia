# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1/interconn/common.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n*kuscia/proto/api/v1/interconn/common.proto\x12\x1dkuscia.proto.api.v1.interconn\x1a\x1cgoogle/protobuf/struct.proto\"R\n\x0e\x43ommonResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12%\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"(\n\x0b\x43omponentIO\x12\x0c\n\x04type\x18\x01 \x01(\t\x12\x0b\n\x03key\x18\x02 \x01(\t\"\x98\x02\n\x06\x43onfig\x12\x37\n\x04role\x18\x01 \x01(\x0b\x32).kuscia.proto.api.v1.interconn.ConfigRole\x12\x41\n\tinitiator\x18\x02 \x01(\x0b\x32..kuscia.proto.api.v1.interconn.ConfigInitiator\x12?\n\njob_params\x18\x03 \x01(\x0b\x32+.kuscia.proto.api.v1.interconn.ConfigParams\x12@\n\x0btask_params\x18\x04 \x01(\x0b\x32+.kuscia.proto.api.v1.interconn.ConfigParams\x12\x0f\n\x07version\x18\x05 \x01(\t\"0\n\x0f\x43onfigInitiator\x12\x0c\n\x04role\x18\x01 \x01(\t\x12\x0f\n\x07node_id\x18\x02 \x01(\t\":\n\nConfigRole\x12\x0f\n\x07\x61rbiter\x18\x01 \x03(\t\x12\x0c\n\x04host\x18\x02 \x03(\t\x12\r\n\x05guest\x18\x03 \x03(\t\"\xb0\x01\n\x0c\x43onfigParams\x12%\n\x04host\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12(\n\x07\x61rbiter\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12&\n\x05guest\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\'\n\x06\x63ommon\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"<\n\x0f\x43onfigResources\x12\x0c\n\x04\x64isk\x18\x01 \x01(\x05\x12\x0e\n\x06memory\x18\x02 \x01(\x05\x12\x0b\n\x03\x63pu\x18\x03 \x01(\x05\x42\x35Z3github.com/secretflow/kuscia/proto/api/v1/interconnb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kuscia.proto.api.v1.interconn.common_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z3github.com/secretflow/kuscia/proto/api/v1/interconn'
  _globals['_COMMONRESPONSE']._serialized_start=107
  _globals['_COMMONRESPONSE']._serialized_end=189
  _globals['_COMPONENTIO']._serialized_start=191
  _globals['_COMPONENTIO']._serialized_end=231
  _globals['_CONFIG']._serialized_start=234
  _globals['_CONFIG']._serialized_end=514
  _globals['_CONFIGINITIATOR']._serialized_start=516
  _globals['_CONFIGINITIATOR']._serialized_end=564
  _globals['_CONFIGROLE']._serialized_start=566
  _globals['_CONFIGROLE']._serialized_end=624
  _globals['_CONFIGPARAMS']._serialized_start=627
  _globals['_CONFIGPARAMS']._serialized_end=803
  _globals['_CONFIGRESOURCES']._serialized_start=805
  _globals['_CONFIGRESOURCES']._serialized_end=865
# @@protoc_insertion_point(module_scope)
