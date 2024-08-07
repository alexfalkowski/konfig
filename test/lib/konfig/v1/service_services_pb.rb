# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: konfig/v1/service.proto for package 'Konfig.V1'

require 'grpc'
require 'konfig/v1/service_pb'

module Konfig
  module V1
    module Service
      # Service allows to manage all application configurations.
      class Service

        include ::GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'konfig.v1.Service'

        # GetConfig for a specific application.
        rpc :GetConfig, ::Konfig::V1::GetConfigRequest, ::Konfig::V1::GetConfigResponse
        # GetSecrets that are configured.
        rpc :GetSecrets, ::Konfig::V1::GetSecretsRequest, ::Konfig::V1::GetSecretsResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end
