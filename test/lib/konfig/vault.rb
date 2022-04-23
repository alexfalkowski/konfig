# frozen_string_literal: true

module Konfig
  class Vault
    def initialize
      @client = ::Vault.logical
    end

    def delete(key)
      client.delete(key)
    end

    def read(key)
      client.read(key).data[:data][:value]
    end

    def write(key, value)
      client.write(key, JSON.parse(value))
    end

    private

    attr_reader :client
  end
end
