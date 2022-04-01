# frozen_string_literal: true

Nonnative.configure do |config|
  config.load_file('nonnative.yml')
end

Given('I start nonnative') do
  Nonnative.start
end

Given('I wait for {int} seconds for the changes to apply') do |seconds|
  sleep seconds # HACK: Need to give the server sometime to adjust to the change.
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
