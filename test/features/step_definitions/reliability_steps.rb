# frozen_string_literal: true

Given('the github api is down') do
  Nonnative.configuration.processes[0].environment['GITHUB_API_URL'] = 'http://not_valid:4567/'
end

Then('the github api is back up') do
  Nonnative.configuration.processes[0].environment['GITHUB_API_URL'] = 'http://localhost:4567'
end
