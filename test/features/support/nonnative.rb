# frozen_string_literal: true

Nonnative.configure do |config|
  config.load_file('nonnative.yml')
end

Given('I start nonnative') do
  Nonnative.start
end

Before('@startup') do
  Nonnative.start
end

After('@startup') do
  Nonnative.stop
end

After('@manual') do
  Nonnative.stop
end
