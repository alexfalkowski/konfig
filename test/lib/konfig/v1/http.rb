# frozen_string_literal: true

module Konfig
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_config(app, ver, env, cmd, headers = {})
        default_headers = {
          content_type: :json,
          accept: :json
        }

        default_headers.merge!(headers)

        get("v1/config/#{app}/#{ver}/#{env}/#{cmd}", headers, 10)
      end
    end
  end
end
