# frozen_string_literal: true

# Adapted from https://gist.github.com/RaVbaker/d9ead3c92b915f997dab25c7f0c0ab65
module Konfig
  class GitHub < Sinatra::Application
    configure do
      set :logging, false
      set :host, 'api.github.com'
    end

    def retrieve_headers(request)
      headers = request.env.map do |header, value|
        [header[5..].split('_').map(&:capitalize).join('-'), value] if header.start_with?('HTTP_')
      end
      headers = headers.compact.to_h

      headers.except('Host', 'Accept-Encoding', 'Version')
    end

    def build_url(request, settings)
      URI::HTTPS.build(host: settings.host, path: request.path_info, query: request.query_string).to_s
    end

    def api_response(verb, uri, opts)
      client = RestClient::Resource.new(uri, opts)

      client.send(verb)
    rescue RestClient::ExceptionWithResponse => e
      e.response
    end

    %w[get post put patch delete].each do |verb|
      send(verb, /.*/) do
        uri = build_url(request, settings)
        opts = { headers: retrieve_headers(request) }
        res = api_response(verb, uri, opts)

        status res.code
        res.body
      end
    end
  end

  class GitHubServer < Nonnative::HTTPServer
    def app
      Konfig::GitHub
    end
  end
end
