# frozen_string_literal: true

require 'securerandom'
require 'yaml'

require 'grpc/health/v1/health_services_pb'

module Konfig
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('http://localhost:8080')
    end

    def serve_config
      @serve_config ||= YAML.load_file('.config/serve.config.yml')
    end

    def worker_config
      @worker_config ||= YAML.load_file('.config/worker.config.yml')
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:9090', :this_channel_is_insecure)
    end
  end
end
