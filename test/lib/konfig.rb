# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'aws-sdk-s3'
require 'aws-sdk-ssm'
require 'vault'
require 'grpc/health/v1/health_services_pb'

require 'konfig/s3'
require 'konfig/ssm'
require 'konfig/vault'
require 'konfig/v1/http'
require 'konfig/v1/service_services_pb'

module Konfig
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('http://localhost:8080')
    end

    def vault
      @vault ||= Konfig::Vault.new
    end

    def s3
      @s3 ||= Konfig::S3.new
    end

    def ssm
      @ssm ||= Konfig::SSM.new
    end

    def server_config(source)
      YAML.load_file(".config/#{source}.server.yaml")
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:8080', :this_channel_is_insecure)
    end
  end

  module V1
    class << self
      def server_http
        @server_http ||= Konfig::V1::HTTP.new('http://localhost:8080')
      end

      def server_grpc
        @server_grpc ||= Konfig::V1::Service::Stub.new('localhost:8080', :this_channel_is_insecure)
      end
    end
  end
end
