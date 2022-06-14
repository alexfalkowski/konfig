# frozen_string_literal: true

module Konfig
  class S3
    def initialize
      credentials = Aws::Credentials.new('access', 'secret')
      @client = Aws::S3::Client.new(endpoint: 'http://localhost:4566', credentials: credentials, region: 'eu-west-1', force_path_style: true)
    end

    def write(key, value)
      create
      client.put_object(bucket: 'bucket', key: key, body: value)
    end

    def delete(key)
      client.delete_object(bucket: 'bucket', key: key)
    end

    private

    attr_reader :client

    def create
      client.head_bucket(bucket: 'bucket')
    rescue StandardError
      client.create_bucket(bucket: 'bucket')
    end
  end
end
