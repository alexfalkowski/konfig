# frozen_string_literal: true

# Adapted from https://gist.github.com/RaVbaker/d9ead3c92b915f997dab25c7f0c0ab65
module Konfig
  class GitHub < Sinatra::Application
    configure do
      enable :logging, :dump_errors, :raise_errors
    end

    def retrieve_headers(request)
      headers = request.env.map do |header, value|
        [header[5..].split('_').map(&:capitalize).join('-'), value] if header.start_with?('HTTP_')
      end
      headers = headers.compact.to_h

      headers.except('Host', 'Accept-Encoding', 'Version')
    end

    def build_url(request)
      URI::HTTPS.build(host: 'api.github.com', path: request.path_info, query: request.query_string).to_s
    end

    def make_request(verb, uri, body, opts)
      client = RestClient::Resource.new(uri, opts)

      if body.empty?
        client.send(verb).body
      else
        client.send(verb, JSON.parse(body).body)
      end
    rescue RestClient::ExceptionWithResponse => e
      e.response.body
    end

    %w[get post put patch delete].each do |verb|
      send(verb, /.*/) do
        uri = build_url(request)
        opts = {
          headers: retrieve_headers(request),
          read_timeout: 10, open_timeout: 10
        }

        make_request(verb, uri, request.body.read, opts)
      end
    end
  end

  class GitHubServer < Nonnative::HTTPServer
    def app
      GitHub.new
    end
  end
end
