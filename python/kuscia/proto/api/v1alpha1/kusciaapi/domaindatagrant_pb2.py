# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/kusciaapi/domaindatagrant.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1 import common_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n9kuscia/proto/api/v1alpha1/kusciaapi/domaindatagrant.proto\x12#kuscia.proto.api.v1alpha1.kusciaapi\x1a&kuscia/proto/api/v1alpha1/common.proto\"\xa4\x03\n\x1c\x43reateDomainDataGrantRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x1a\n\x12\x64omaindatagrant_id\x18\x02 \x01(\t\x12\x15\n\rdomaindata_id\x18\x03 \x01(\t\x12\x14\n\x0cgrant_domain\x18\x04 \x01(\t\x12>\n\x05limit\x18\x05 \x01(\x0b\x32/.kuscia.proto.api.v1alpha1.kusciaapi.GrantLimit\x12g\n\x0b\x64\x65scription\x18\x06 \x03(\x0b\x32R.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainDataGrantRequest.DescriptionEntry\x12\x11\n\tsignature\x18\x07 \x01(\t\x12\x11\n\tdomain_id\x18\x08 \x01(\t\x1a\x32\n\x10\x44\x65scriptionEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xa8\x01\n\x1d\x43reateDomainDataGrantResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12T\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x46.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainDataGrantResponseData\"?\n!CreateDomainDataGrantResponseData\x12\x1a\n\x12\x64omaindatagrant_id\x18\x01 \x01(\t\"\xa5\x01\n\x0f\x44omainDataGrant\x12\x46\n\x04\x64\x61ta\x18\x01 \x01(\x0b\x32\x38.kuscia.proto.api.v1alpha1.kusciaapi.DomainDataGrantData\x12J\n\x06status\x18\x02 \x01(\x0b\x32:.kuscia.proto.api.v1alpha1.kusciaapi.DomainDataGrantStatus\"x\n\x15\x44omainDataGrantStatus\x12\r\n\x05phase\x18\x01 \x01(\t\x12\x0f\n\x07message\x18\x02 \x01(\t\x12?\n\x07records\x18\x03 \x03(\x0b\x32..kuscia.proto.api.v1alpha1.kusciaapi.UseRecord\"V\n\tUseRecord\x12\x10\n\x08use_time\x18\x01 \x01(\x03\x12\x14\n\x0cgrant_domain\x18\x02 \x01(\t\x12\x11\n\tcomponent\x18\x03 \x01(\t\x12\x0e\n\x06output\x18\x04 \x01(\t\"\xe8\x02\n\x13\x44omainDataGrantData\x12\x1a\n\x12\x64omaindatagrant_id\x18\x01 \x01(\t\x12\x0e\n\x06\x61uthor\x18\x02 \x01(\t\x12\x15\n\rdomaindata_id\x18\x03 \x01(\t\x12\x14\n\x0cgrant_domain\x18\x04 \x01(\t\x12>\n\x05limit\x18\x05 \x01(\x0b\x32/.kuscia.proto.api.v1alpha1.kusciaapi.GrantLimit\x12^\n\x0b\x64\x65scription\x18\x06 \x03(\x0b\x32I.kuscia.proto.api.v1alpha1.kusciaapi.DomainDataGrantData.DescriptionEntry\x12\x11\n\tsignature\x18\x07 \x01(\t\x12\x11\n\tdomain_id\x18\x08 \x01(\t\x1a\x32\n\x10\x44\x65scriptionEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x86\x01\n\nGrantLimit\x12\x17\n\x0f\x65xpiration_time\x18\x01 \x01(\x03\x12\x11\n\tuse_count\x18\x02 \x01(\x05\x12\x0f\n\x07\x66low_id\x18\x03 \x01(\t\x12\x12\n\ncomponents\x18\x04 \x03(\t\x12\x11\n\tinitiator\x18\x05 \x01(\t\x12\x14\n\x0cinput_config\x18\x06 \x01(\t\"\xa4\x03\n\x1cUpdateDomainDataGrantRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x1a\n\x12\x64omaindatagrant_id\x18\x02 \x01(\t\x12\x15\n\rdomaindata_id\x18\x03 \x01(\t\x12\x14\n\x0cgrant_domain\x18\x04 \x01(\t\x12>\n\x05limit\x18\x05 \x01(\x0b\x32/.kuscia.proto.api.v1alpha1.kusciaapi.GrantLimit\x12g\n\x0b\x64\x65scription\x18\x06 \x03(\x0b\x32R.kuscia.proto.api.v1alpha1.kusciaapi.UpdateDomainDataGrantRequest.DescriptionEntry\x12\x11\n\tsignature\x18\x07 \x01(\t\x12\x11\n\tdomain_id\x18\x08 \x01(\t\x1a\x32\n\x10\x44\x65scriptionEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"R\n\x1dUpdateDomainDataGrantResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\"\x87\x01\n\x1c\x44\x65leteDomainDataGrantRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x1a\n\x12\x64omaindatagrant_id\x18\x03 \x01(\t\"R\n\x1d\x44\x65leteDomainDataGrantResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\"P\n\x1fQueryDomainDataGrantRequestData\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x1a\n\x12\x64omaindatagrant_id\x18\x02 \x01(\t\"\x86\x01\n\x1bQueryDomainDataGrantRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x1a\n\x12\x64omaindatagrant_id\x18\x03 \x01(\t\"\x95\x01\n\x1cQueryDomainDataGrantResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12\x42\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x34.kuscia.proto.api.v1alpha1.kusciaapi.DomainDataGrant\"\xb0\x01\n BatchQueryDomainDataGrantRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12R\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x44.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainDataGrantRequestData\"\x9a\x01\n!BatchQueryDomainDataGrantResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12\x42\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x34.kuscia.proto.api.v1alpha1.kusciaapi.DomainDataGrant\"\xa9\x01\n\x1aListDomainDataGrantRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12Q\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x43.kuscia.proto.api.v1alpha1.kusciaapi.ListDomainDataGrantRequestData\"d\n\x1eListDomainDataGrantRequestData\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x14\n\x0cgrant_domain\x18\x02 \x01(\t\x12\x19\n\x11\x64omaindata_vendor\x18\x03 \x01(\t\"\x98\x01\n\x1bListDomainDataGrantResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12\x46\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x38.kuscia.proto.api.v1alpha1.kusciaapi.DomainDataGrantList\"i\n\x13\x44omainDataGrantList\x12R\n\x14\x64omaindatagrant_list\x18\x01 \x03(\x0b\x32\x34.kuscia.proto.api.v1alpha1.kusciaapi.DomainDataGrant2\xe1\x07\n\x16\x44omainDataGrantService\x12\x9e\x01\n\x15\x43reateDomainDataGrant\x12\x41.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainDataGrantRequest\x1a\x42.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainDataGrantResponse\x12\x9e\x01\n\x15UpdateDomainDataGrant\x12\x41.kuscia.proto.api.v1alpha1.kusciaapi.UpdateDomainDataGrantRequest\x1a\x42.kuscia.proto.api.v1alpha1.kusciaapi.UpdateDomainDataGrantResponse\x12\x9e\x01\n\x15\x44\x65leteDomainDataGrant\x12\x41.kuscia.proto.api.v1alpha1.kusciaapi.DeleteDomainDataGrantRequest\x1a\x42.kuscia.proto.api.v1alpha1.kusciaapi.DeleteDomainDataGrantResponse\x12\x9b\x01\n\x14QueryDomainDataGrant\x12@.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainDataGrantRequest\x1a\x41.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainDataGrantResponse\x12\xaa\x01\n\x19\x42\x61tchQueryDomainDataGrant\x12\x45.kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainDataGrantRequest\x1a\x46.kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainDataGrantResponse\x12\x98\x01\n\x13ListDomainDataGrant\x12?.kuscia.proto.api.v1alpha1.kusciaapi.ListDomainDataGrantRequest\x1a@.kuscia.proto.api.v1alpha1.kusciaapi.ListDomainDataGrantResponseB^\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapib\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kuscia.proto.api.v1alpha1.kusciaapi.domaindatagrant_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapi'
  _globals['_CREATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._options = None
  _globals['_CREATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_options = b'8\001'
  _globals['_DOMAINDATAGRANTDATA_DESCRIPTIONENTRY']._options = None
  _globals['_DOMAINDATAGRANTDATA_DESCRIPTIONENTRY']._serialized_options = b'8\001'
  _globals['_UPDATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._options = None
  _globals['_UPDATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_options = b'8\001'
  _globals['_CREATEDOMAINDATAGRANTREQUEST']._serialized_start=139
  _globals['_CREATEDOMAINDATAGRANTREQUEST']._serialized_end=559
  _globals['_CREATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_start=509
  _globals['_CREATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_end=559
  _globals['_CREATEDOMAINDATAGRANTRESPONSE']._serialized_start=562
  _globals['_CREATEDOMAINDATAGRANTRESPONSE']._serialized_end=730
  _globals['_CREATEDOMAINDATAGRANTRESPONSEDATA']._serialized_start=732
  _globals['_CREATEDOMAINDATAGRANTRESPONSEDATA']._serialized_end=795
  _globals['_DOMAINDATAGRANT']._serialized_start=798
  _globals['_DOMAINDATAGRANT']._serialized_end=963
  _globals['_DOMAINDATAGRANTSTATUS']._serialized_start=965
  _globals['_DOMAINDATAGRANTSTATUS']._serialized_end=1085
  _globals['_USERECORD']._serialized_start=1087
  _globals['_USERECORD']._serialized_end=1173
  _globals['_DOMAINDATAGRANTDATA']._serialized_start=1176
  _globals['_DOMAINDATAGRANTDATA']._serialized_end=1536
  _globals['_DOMAINDATAGRANTDATA_DESCRIPTIONENTRY']._serialized_start=509
  _globals['_DOMAINDATAGRANTDATA_DESCRIPTIONENTRY']._serialized_end=559
  _globals['_GRANTLIMIT']._serialized_start=1539
  _globals['_GRANTLIMIT']._serialized_end=1673
  _globals['_UPDATEDOMAINDATAGRANTREQUEST']._serialized_start=1676
  _globals['_UPDATEDOMAINDATAGRANTREQUEST']._serialized_end=2096
  _globals['_UPDATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_start=509
  _globals['_UPDATEDOMAINDATAGRANTREQUEST_DESCRIPTIONENTRY']._serialized_end=559
  _globals['_UPDATEDOMAINDATAGRANTRESPONSE']._serialized_start=2098
  _globals['_UPDATEDOMAINDATAGRANTRESPONSE']._serialized_end=2180
  _globals['_DELETEDOMAINDATAGRANTREQUEST']._serialized_start=2183
  _globals['_DELETEDOMAINDATAGRANTREQUEST']._serialized_end=2318
  _globals['_DELETEDOMAINDATAGRANTRESPONSE']._serialized_start=2320
  _globals['_DELETEDOMAINDATAGRANTRESPONSE']._serialized_end=2402
  _globals['_QUERYDOMAINDATAGRANTREQUESTDATA']._serialized_start=2404
  _globals['_QUERYDOMAINDATAGRANTREQUESTDATA']._serialized_end=2484
  _globals['_QUERYDOMAINDATAGRANTREQUEST']._serialized_start=2487
  _globals['_QUERYDOMAINDATAGRANTREQUEST']._serialized_end=2621
  _globals['_QUERYDOMAINDATAGRANTRESPONSE']._serialized_start=2624
  _globals['_QUERYDOMAINDATAGRANTRESPONSE']._serialized_end=2773
  _globals['_BATCHQUERYDOMAINDATAGRANTREQUEST']._serialized_start=2776
  _globals['_BATCHQUERYDOMAINDATAGRANTREQUEST']._serialized_end=2952
  _globals['_BATCHQUERYDOMAINDATAGRANTRESPONSE']._serialized_start=2955
  _globals['_BATCHQUERYDOMAINDATAGRANTRESPONSE']._serialized_end=3109
  _globals['_LISTDOMAINDATAGRANTREQUEST']._serialized_start=3112
  _globals['_LISTDOMAINDATAGRANTREQUEST']._serialized_end=3281
  _globals['_LISTDOMAINDATAGRANTREQUESTDATA']._serialized_start=3283
  _globals['_LISTDOMAINDATAGRANTREQUESTDATA']._serialized_end=3383
  _globals['_LISTDOMAINDATAGRANTRESPONSE']._serialized_start=3386
  _globals['_LISTDOMAINDATAGRANTRESPONSE']._serialized_end=3538
  _globals['_DOMAINDATAGRANTLIST']._serialized_start=3540
  _globals['_DOMAINDATAGRANTLIST']._serialized_end=3645
  _globals['_DOMAINDATAGRANTSERVICE']._serialized_start=3648
  _globals['_DOMAINDATAGRANTSERVICE']._serialized_end=4641
# @@protoc_insertion_point(module_scope)
