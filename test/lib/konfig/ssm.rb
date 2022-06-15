# frozen_string_literal: true

module Konfig
  class SSM
    def initialize
      credentials = Aws::Credentials.new('access', 'secret')
      @client = Aws::SSM::Client.new(endpoint: 'http://localhost:4566', credentials: credentials, region: 'eu-west-1')
    end

    def write(name, value)
      client.put_parameter(name: name, value: value, overwrite: true)
      true
    end

    def delete(name)
      client.delete_parameter(name: name)
      true
    rescue Aws::SSM::Errors::ParameterNotFound
      false
    end

    private

    attr_reader :client
  end
end
