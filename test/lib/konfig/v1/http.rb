# frozen_string_literal: true

module Konfig
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_config(params, headers = {})
        default_headers = {
          content_type: :json,
          accept: :json
        }

        default_headers.merge!(headers)

        get("v1/config/#{params[:app]}/#{params[:ver]}/#{params[:env]}/#{params[:continent]}/#{params[:country]}/#{params[:cmd]}", headers, 10)
      end
    end
  end
end
