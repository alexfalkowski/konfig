# frozen_string_literal: true

module Konfig
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_config(params, opts = {})
        post('/v1/config', params.to_json, opts)
      end

      def get_secrets(params, opts = {})
        post('/v1/secrets', { secrets: params }.to_json, opts)
      end
    end
  end
end
