# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: konfig.v1.proto for package 'Konfig.V1'

require 'grpc'

module Konfig
  module V1
    module Configurator
      # Configurator allows to manage all application configurations.
      class Service

        include ::GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'konfig.v1.Configurator'

        # GetConfig for a specific application, version, environment and command.
        rpc :GetConfig, ::Konfig::V1::GetConfigRequest, ::Konfig::V1::GetConfigResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end
