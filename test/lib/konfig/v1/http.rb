# frozen_string_literal: true

module Konfig
  module V1
    class HTTP < Nonnative::HTTPClient
      def initialize(host)
        @tries = 3
        @wait = 0.5

        super
      end

      def get_config(params, opts = {})
        with_retry(tries, wait) do
          post('/konfig.v1.Service/GetConfig', params.to_json, opts)
        end
      end

      def get_secrets(params, opts = {})
        post('/konfig.v1.Service/GetSecrets', { secrets: params }.to_json, opts)
      end

      private

      attr_reader :tries, :wait
    end
  end
end
