# frozen_string_literal: true

module Konfig
  class SSM
    def initialize
      credentials = Aws::Credentials.new('access', 'secret')
      @client = Aws::SSM::Client.new(endpoint: 'http://localhost:4566', credentials:, region: 'eu-west-1')
    end

    def write(name, value)
      client.put_parameter(name:, value:, type: 'String', overwrite: true)
      true
    end

    def delete(name)
      client.delete_parameter(name:)
      true
    rescue Aws::SSM::Errors::ParameterNotFound
      false
    end

    private

    attr_reader :client
  end
end
