# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'
require 'vault'

require 'grpc/health/v1/health_services_pb'

require 'konfig/vault'
require 'konfig/v1/http'
require 'konfig/v1/service_pb'
require 'konfig/v1/service_services_pb'

module Konfig
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('http://localhost:8080')
    end

    def vault
      @vault ||= Konfig::Vault.new
    end

    def server_config
      @server_config ||= YAML.load_file('.config/server.config.yml')
    end

    def worker_config
      @worker_config ||= YAML.load_file('.config/worker.config.yml')
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:9090', :this_channel_is_insecure)
    end

    def server_http
      @server_http ||= Konfig::V1::HTTP.new('http://localhost:8080')
    end

    def server_grpc
      @server_grpc ||= Konfig::V1::ConfiguratorService::Stub.new('localhost:9090', :this_channel_is_insecure)
    end
  end
end
