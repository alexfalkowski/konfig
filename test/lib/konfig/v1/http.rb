# frozen_string_literal: true

module Konfig
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_config(application, environment, headers = {})
        default_headers = {
          content_type: :json,
          accept: :json
        }

        default_headers.merge!(headers)

        get("v1/konfig/#{application}/#{environment}", headers, 1)
      end
    end
  end
end
