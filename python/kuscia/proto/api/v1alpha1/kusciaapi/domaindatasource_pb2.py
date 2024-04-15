# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/kusciaapi/domaindatasource.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1 import common_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n:kuscia/proto/api/v1alpha1/kusciaapi/domaindatasource.proto\x12#kuscia.proto.api.v1alpha1.kusciaapi\x1a&kuscia/proto/api/v1alpha1/common.proto\"\xd4\x02\n\x1d\x43reateDomainDataSourceRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x15\n\rdatasource_id\x18\x03 \x01(\t\x12\x0c\n\x04type\x18\x04 \x01(\t\x12\x11\n\x04name\x18\x05 \x01(\tH\x00\x88\x01\x01\x12\x46\n\x04info\x18\x06 \x01(\x0b\x32\x33.kuscia.proto.api.v1alpha1.kusciaapi.DataSourceInfoH\x01\x88\x01\x01\x12\x15\n\x08info_key\x18\x07 \x01(\tH\x02\x88\x01\x01\x12\x1c\n\x0f\x61\x63\x63\x65ss_directly\x18\x08 \x01(\x08H\x03\x88\x01\x01\x42\x07\n\x05_nameB\x07\n\x05_infoB\x0b\n\t_info_keyB\x12\n\x10_access_directly\"\xaa\x01\n\x1e\x43reateDomainDataSourceResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12U\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32G.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainDataSourceResponseData\";\n\"CreateDomainDataSourceResponseData\x12\x15\n\rdatasource_id\x18\x01 \x01(\t\"\xd4\x02\n\x1dUpdateDomainDataSourceRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x15\n\rdatasource_id\x18\x03 \x01(\t\x12\x0c\n\x04type\x18\x04 \x01(\t\x12\x11\n\x04name\x18\x05 \x01(\tH\x00\x88\x01\x01\x12\x46\n\x04info\x18\x06 \x01(\x0b\x32\x33.kuscia.proto.api.v1alpha1.kusciaapi.DataSourceInfoH\x01\x88\x01\x01\x12\x15\n\x08info_key\x18\x07 \x01(\tH\x02\x88\x01\x01\x12\x1c\n\x0f\x61\x63\x63\x65ss_directly\x18\x08 \x01(\x08H\x03\x88\x01\x01\x42\x07\n\x05_nameB\x07\n\x05_infoB\x0b\n\t_info_keyB\x12\n\x10_access_directly\"S\n\x1eUpdateDomainDataSourceResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\"\x83\x01\n\x1d\x44\x65leteDomainDataSourceRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x15\n\rdatasource_id\x18\x03 \x01(\t\"S\n\x1e\x44\x65leteDomainDataSourceResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\"\x82\x01\n\x1cQueryDomainDataSourceRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x15\n\rdatasource_id\x18\x03 \x01(\t\"\x97\x01\n\x1dQueryDomainDataSourceResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12\x43\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x35.kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSource\"L\n QueryDomainDataSourceRequestData\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x15\n\rdatasource_id\x18\x02 \x01(\t\"\xb2\x01\n!BatchQueryDomainDataSourceRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\x12S\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x45.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainDataSourceRequestData\"\xa0\x01\n\"BatchQueryDomainDataSourceResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12G\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x39.kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSourceList\"f\n\x14\x44omainDataSourceList\x12N\n\x0f\x64\x61tasource_list\x18\x01 \x03(\x0b\x32\x35.kuscia.proto.api.v1alpha1.kusciaapi.DomainDataSource\"\xd6\x01\n\x10\x44omainDataSource\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x15\n\rdatasource_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x0c\n\x04type\x18\x04 \x01(\t\x12\x0e\n\x06status\x18\x05 \x01(\t\x12\x41\n\x04info\x18\x06 \x01(\x0b\x32\x33.kuscia.proto.api.v1alpha1.kusciaapi.DataSourceInfo\x12\x10\n\x08info_key\x18\x07 \x01(\t\x12\x17\n\x0f\x61\x63\x63\x65ss_directly\x18\x08 \x01(\x08\"\xef\x01\n\x0e\x44\x61taSourceInfo\x12I\n\x07localfs\x18\x01 \x01(\x0b\x32\x38.kuscia.proto.api.v1alpha1.kusciaapi.LocalDataSourceInfo\x12\x43\n\x03oss\x18\x02 \x01(\x0b\x32\x36.kuscia.proto.api.v1alpha1.kusciaapi.OssDataSourceInfo\x12M\n\x08\x64\x61tabase\x18\x03 \x01(\x0b\x32;.kuscia.proto.api.v1alpha1.kusciaapi.DatabaseDataSourceInfo\"#\n\x13LocalDataSourceInfo\x12\x0c\n\x04path\x18\x01 \x01(\t\"\xb3\x01\n\x11OssDataSourceInfo\x12\x10\n\x08\x65ndpoint\x18\x01 \x01(\t\x12\x0e\n\x06\x62ucket\x18\x02 \x01(\t\x12\x0e\n\x06prefix\x18\x03 \x01(\t\x12\x15\n\raccess_key_id\x18\x04 \x01(\t\x12\x19\n\x11\x61\x63\x63\x65ss_key_secret\x18\x05 \x01(\t\x12\x13\n\x0bvirtualhost\x18\x06 \x01(\x08\x12\x0f\n\x07version\x18\x07 \x01(\t\x12\x14\n\x0cstorage_type\x18\x08 \x01(\t\"\\\n\x16\x44\x61tabaseDataSourceInfo\x12\x10\n\x08\x65ndpoint\x18\x01 \x01(\t\x12\x0c\n\x04user\x18\x02 \x01(\t\x12\x10\n\x08password\x18\x03 \x01(\t\x12\x10\n\x08\x64\x61tabase\x18\x04 \x01(\t2\xd6\x06\n\x17\x44omainDataSourceService\x12\xa1\x01\n\x16\x43reateDomainDataSource\x12\x42.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainDataSourceRequest\x1a\x43.kuscia.proto.api.v1alpha1.kusciaapi.CreateDomainDataSourceResponse\x12\x9e\x01\n\x15QueryDomainDataSource\x12\x41.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainDataSourceRequest\x1a\x42.kuscia.proto.api.v1alpha1.kusciaapi.QueryDomainDataSourceResponse\x12\xa1\x01\n\x16UpdateDomainDataSource\x12\x42.kuscia.proto.api.v1alpha1.kusciaapi.UpdateDomainDataSourceRequest\x1a\x43.kuscia.proto.api.v1alpha1.kusciaapi.UpdateDomainDataSourceResponse\x12\xa1\x01\n\x16\x44\x65leteDomainDataSource\x12\x42.kuscia.proto.api.v1alpha1.kusciaapi.DeleteDomainDataSourceRequest\x1a\x43.kuscia.proto.api.v1alpha1.kusciaapi.DeleteDomainDataSourceResponse\x12\xad\x01\n\x1a\x42\x61tchQueryDomainDataSource\x12\x46.kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainDataSourceRequest\x1aG.kuscia.proto.api.v1alpha1.kusciaapi.BatchQueryDomainDataSourceResponseB^\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapib\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kuscia.proto.api.v1alpha1.kusciaapi.domaindatasource_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapi'
  _globals['_CREATEDOMAINDATASOURCEREQUEST']._serialized_start=140
  _globals['_CREATEDOMAINDATASOURCEREQUEST']._serialized_end=480
  _globals['_CREATEDOMAINDATASOURCERESPONSE']._serialized_start=483
  _globals['_CREATEDOMAINDATASOURCERESPONSE']._serialized_end=653
  _globals['_CREATEDOMAINDATASOURCERESPONSEDATA']._serialized_start=655
  _globals['_CREATEDOMAINDATASOURCERESPONSEDATA']._serialized_end=714
  _globals['_UPDATEDOMAINDATASOURCEREQUEST']._serialized_start=717
  _globals['_UPDATEDOMAINDATASOURCEREQUEST']._serialized_end=1057
  _globals['_UPDATEDOMAINDATASOURCERESPONSE']._serialized_start=1059
  _globals['_UPDATEDOMAINDATASOURCERESPONSE']._serialized_end=1142
  _globals['_DELETEDOMAINDATASOURCEREQUEST']._serialized_start=1145
  _globals['_DELETEDOMAINDATASOURCEREQUEST']._serialized_end=1276
  _globals['_DELETEDOMAINDATASOURCERESPONSE']._serialized_start=1278
  _globals['_DELETEDOMAINDATASOURCERESPONSE']._serialized_end=1361
  _globals['_QUERYDOMAINDATASOURCEREQUEST']._serialized_start=1364
  _globals['_QUERYDOMAINDATASOURCEREQUEST']._serialized_end=1494
  _globals['_QUERYDOMAINDATASOURCERESPONSE']._serialized_start=1497
  _globals['_QUERYDOMAINDATASOURCERESPONSE']._serialized_end=1648
  _globals['_QUERYDOMAINDATASOURCEREQUESTDATA']._serialized_start=1650
  _globals['_QUERYDOMAINDATASOURCEREQUESTDATA']._serialized_end=1726
  _globals['_BATCHQUERYDOMAINDATASOURCEREQUEST']._serialized_start=1729
  _globals['_BATCHQUERYDOMAINDATASOURCEREQUEST']._serialized_end=1907
  _globals['_BATCHQUERYDOMAINDATASOURCERESPONSE']._serialized_start=1910
  _globals['_BATCHQUERYDOMAINDATASOURCERESPONSE']._serialized_end=2070
  _globals['_DOMAINDATASOURCELIST']._serialized_start=2072
  _globals['_DOMAINDATASOURCELIST']._serialized_end=2174
  _globals['_DOMAINDATASOURCE']._serialized_start=2177
  _globals['_DOMAINDATASOURCE']._serialized_end=2391
  _globals['_DATASOURCEINFO']._serialized_start=2394
  _globals['_DATASOURCEINFO']._serialized_end=2633
  _globals['_LOCALDATASOURCEINFO']._serialized_start=2635
  _globals['_LOCALDATASOURCEINFO']._serialized_end=2670
  _globals['_OSSDATASOURCEINFO']._serialized_start=2673
  _globals['_OSSDATASOURCEINFO']._serialized_end=2852
  _globals['_DATABASEDATASOURCEINFO']._serialized_start=2854
  _globals['_DATABASEDATASOURCEINFO']._serialized_end=2946
  _globals['_DOMAINDATASOURCESERVICE']._serialized_start=2949
  _globals['_DOMAINDATASOURCESERVICE']._serialized_end=3803
# @@protoc_insertion_point(module_scope)
