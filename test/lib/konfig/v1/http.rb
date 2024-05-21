# frozen_string_literal: true

module Konfig
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_config(params, opts = {})
        url = "v1/config/#{params[:app]}/#{params[:ver]}/#{params[:env]}/#{params[:continent]}/#{params[:country]}/#{params[:cmd]}/#{params[:kind]}"

        get(url, opts)
      end

      def get_secrets(params, opts = {})
        post('/v1/secrets', params.to_json, opts)
      end
    end
  end
end
