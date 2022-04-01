# frozen_string_literal: true

module Konfig
  class Vault
    def initialize
      @client = ::Vault.logical
    end

    def read(key)
      client.read("secret/data/#{key}").data[:data][:value]
    end

    def write(key, value)
      client.write("secret/data/#{key}", data: { value: value })
    end

    private

    attr_reader :client
  end
end
