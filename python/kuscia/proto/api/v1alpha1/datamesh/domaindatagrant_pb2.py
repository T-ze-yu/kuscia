# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/datamesh/domaindatagrant.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1 import common_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n8kuscia/proto/api/v1alpha1/datamesh/domaindatagrant.proto\x12\"kuscia.proto.api.v1alpha1.datamesh\x1a&kuscia/proto/api/v1alpha1/common.proto\"\x86\x01\n\nGrantLimit\x12\x17\n\x0f\x65xpiration_time\x18\x01 \x01(\x03\x12\x11\n\tuse_count\x18\x02 \x01(\x05\x12\x0f\n\x07\x66low_id\x18\x03 \x01(\t\x12\x12\n\ncomponents\x18\x04 \x03(\t\x12\x11\n\tinitiator\x18\x05 \x01(\t\x12\x14\n\x0cinput_config\x18\x06 \x01(\t\"\xd3\x02\n\x13\x44omainDataGrantData\x12\x1a\n\x12\x64omaindatagrant_id\x18\x01 \x01(\t\x12\x0e\n\x06\x61uthor\x18\x02 \x01(\t\x12\x15\n\rdomaindata_id\x18\x03 \x01(\t\x12\x14\n\x0cgrant_domain\x18\x04 \x01(\t\x12=\n\x05limit\x18\x05 \x01(\x0b\x32..kuscia.proto.api.v1alpha1.datamesh.GrantLimit\x12]\n\x0b\x64\x65scription\x18\x06 \x03(\x0b\x32H.kuscia.proto.api.v1alpha1.datamesh.DomainDataGrantData.DescriptionEntry\x12\x11\n\tsignature\x18\x07 \x01(\t\x1a\x32\n\x10\x44\x65scriptionEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xfc\x02\n\x1c\x43reateDomainDataGrantRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x1a\n\x12\x64omaindatagrant_id\x18\x02 \x01(\t\x12\x15\n\rdomaindata_id\x18\x03 \x01(\t\x12\x14\n\x0cgrant_domain\x18\x04 \x01(\t\x12=\n\x05limit\x18\x05 \x01(\x0b\x32..kuscia.proto.api.v1alpha1.datamesh.GrantLimit\x12\x66\n\x0b\x64\x65scription\x18\x06 \x03(\x0b\x32Q.kuscia.proto.api.v1alpha1.datamesh.CreateDomainDataGrantRequest.DescriptionEntry\x1a\x32\n\x10\x44\x65scriptionEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xa7\x01\n\x1d\x43reateDomainDataGrantResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12S\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x45.kuscia.proto.api.v1alpha1.datamesh.CreateDomainDataGrantResponseData\"?\n!CreateDomainDataGrantResponseData\x12\x1a\n\x12\x64omaindatagrant_id\x18\x01 \x01(\t\"\xfc\x02\n\x1cUpdateDomainDataGrantRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x1a\n\x12\x64omaindatagrant_id\x18\x02 \x01(\t\x12\x15\n\rdomaindata_id\x18\x03 \x01(\t\x12\x14\n\x0cgrant_domain\x18\x04 \x01(\t\x12=\n\x05limit\x18\x05 \x01(\x0b\x32..kuscia.proto.api.v1alpha1.datamesh.GrantLimit\x12\x66\n\x0b\x64\x65scription\x18\x06 \x03(\x0b\x32Q.kuscia.proto.api.v1alpha1.datamesh.UpdateDomainDataGrantRequest.DescriptionEntry\x1a\x32\n\x10\x44\x65scriptionEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"R\n\x1dUpdateDomainDataGrantResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\"t\n\x1c\x44\x65leteDomainDataGrantRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x1a\n\x12\x64omaindatagrant_id\x18\x02 \x01(\t\"R\n\x1d\x44\x65leteDomainDataGrantResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\"s\n\x1bQueryDomainDataGrantRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x1a\n\x12\x64omaindatagrant_id\x18\x02 \x01(\t\"\x98\x01\n\x1cQueryDomainDataGrantResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12\x45\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x37.kuscia.proto.api.v1alpha1.datamesh.DomainDataGrantData2\x91\x05\n\x16\x44omainDataGrantService\x12\x9c\x01\n\x15\x43reateDomainDataGrant\x12@.kuscia.proto.api.v1alpha1.datamesh.CreateDomainDataGrantRequest\x1a\x41.kuscia.proto.api.v1alpha1.datamesh.CreateDomainDataGrantResponse\x12\x99\x01\n\x14QueryDomainDataGrant\x12?.kuscia.proto.api.v1alpha1.datamesh.QueryDomainDataGrantRequest\x1a@.kuscia.proto.api.v1alpha1.datamesh.QueryDomainDataGrantResponse\x12\x9c\x01\n\x15UpdateDomainDataGrant\x12@.kuscia.proto.api.v1alpha1.datamesh.UpdateDomainDataGrantRequest\x1a\x41.kuscia.proto.api.v1alpha1.datamesh.UpdateDomainDataGrantResponse\x12\x9c\x01\n\x15\x44\x65leteDomainDataGrant\x12@.kuscia.proto.api.v1alpha1.datamesh.DeleteDomainDataGrantRequest\x1a\x41.kuscia.proto.api.v1alpha1.datamesh.DeleteDomainDataGrantResponseB\\\n org.secretflow.v1alpha1.datameshZ8github.com/secretflow/kuscia/proto/api/v1alpha1/datameshb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kuscia.proto.api.v1alpha1.datamesh.domaindatagrant_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n org.secretflow.v1alpha1.datameshZ8github.com/secretflow/kuscia/proto/api/v1alpha1/datamesh'
  _globals['_DOMAINDATAGRANTDATA_DESCRIPTIONENTRY']._options = None
  _globals['_DOMAINDATAGRANTDATA_DESCRIPTIONENTRY']._serialized_options = b'8\001'
  _globals['_CREATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._options = None
  _globals['_CREATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_options = b'8\001'
  _globals['_UPDATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._options = None
  _globals['_UPDATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_options = b'8\001'
  _globals['_GRANTLIMIT']._serialized_start=137
  _globals['_GRANTLIMIT']._serialized_end=271
  _globals['_DOMAINDATAGRANTDATA']._serialized_start=274
  _globals['_DOMAINDATAGRANTDATA']._serialized_end=613
  _globals['_DOMAINDATAGRANTDATA_DESCRIPTIONENTRY']._serialized_start=563
  _globals['_DOMAINDATAGRANTDATA_DESCRIPTIONENTRY']._serialized_end=613
  _globals['_CREATEDOMAINDATAGRANTREQUEST']._serialized_start=616
  _globals['_CREATEDOMAINDATAGRANTREQUEST']._serialized_end=996
  _globals['_CREATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_start=563
  _globals['_CREATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_end=613
  _globals['_CREATEDOMAINDATAGRANTRESPONSE']._serialized_start=999
  _globals['_CREATEDOMAINDATAGRANTRESPONSE']._serialized_end=1166
  _globals['_CREATEDOMAINDATAGRANTRESPONSEDATA']._serialized_start=1168
  _globals['_CREATEDOMAINDATAGRANTRESPONSEDATA']._serialized_end=1231
  _globals['_UPDATEDOMAINDATAGRANTREQUEST']._serialized_start=1234
  _globals['_UPDATEDOMAINDATAGRANTREQUEST']._serialized_end=1614
  _globals['_UPDATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_start=563
  _globals['_UPDATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_end=613
  _globals['_UPDATEDOMAINDATAGRANTRESPONSE']._serialized_start=1616
  _globals['_UPDATEDOMAINDATAGRANTRESPONSE']._serialized_end=1698
  _globals['_DELETEDOMAINDATAGRANTREQUEST']._serialized_start=1700
  _globals['_DELETEDOMAINDATAGRANTREQUEST']._serialized_end=1816
  _globals['_DELETEDOMAINDATAGRANTRESPONSE']._serialized_start=1818
  _globals['_DELETEDOMAINDATAGRANTRESPONSE']._serialized_end=1900
  _globals['_QUERYDOMAINDATAGRANTREQUEST']._serialized_start=1902
  _globals['_QUERYDOMAINDATAGRANTREQUEST']._serialized_end=2017
  _globals['_QUERYDOMAINDATAGRANTRESPONSE']._serialized_start=2020
  _globals['_QUERYDOMAINDATAGRANTRESPONSE']._serialized_end=2172
  _globals['_DOMAINDATAGRANTSERVICE']._serialized_start=2175
  _globals['_DOMAINDATAGRANTSERVICE']._serialized_end=2832
# @@protoc_insertion_point(module_scope)
