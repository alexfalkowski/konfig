# frozen_string_literal: true

require 'rspec/expectations'
require 'rspec/wait'

World(RSpec::Matchers)
World(RSpec::Wait)

RSpec.configure do |config|
  config.wait_timeout = 3
end
