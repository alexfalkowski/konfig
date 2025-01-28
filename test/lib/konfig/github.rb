# frozen_string_literal: true

module Konfig
  class GitHubServer < Nonnative::HTTPProxyServer
    def initialize(service)
      super('api.github.com', service)
    end
  end
end
