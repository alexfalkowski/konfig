# frozen_string_literal: true

Given('I have a {string} valid setup') do |source|
  Nonnative.configuration.processes[0].environment['KONFIG_GIT_TOKEN'] = ENV['GITHUB_TOKEN'] if source == 'git'
end

Given('I have a {string} invalid setup') do |source|
  case source
  when 'git'
    Nonnative.configuration.processes[0].environment['KONFIG_GIT_TOKEN'] = 'not_a_valid_token'
  when 'folder'
    Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = ".config/invalid.#{source}.server.config.yml"
    @config_set = true
  end
end

Given('I have key {string} with {string} value in vault') do |key, value|
  Konfig.vault.write(key, value)
end

Given('I have {string} as the config file') do |source|
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = ".config/#{source}.server.config.yml" unless @config_set
end

When('I request a config with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'ua' => Konfig.server_config(rows['source'])['transport']['grpc']['user_agent']
  }

  request = Konfig::V1::GetConfigRequest.new(application: rows['app'], version: rows['ver'], environment: rows['env'], command: rows['cmd'])
  @response = Konfig.server_grpc.get_config(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

Then('I should receive a valid config from gRPC:') do |table|
  data = YAML.safe_load(@response.data)
  rows = table.rows_hash

  expect(@response.application).to eq(rows['app'])
  expect(@response.version).to eq(rows['ver'])
  expect(@response.environment).to eq(rows['env'])
  expect(@response.command).to eq(rows['cmd'])
  expect(@response.content_type).to eq('application/yaml')
  expect(data['transport']['http']['user_agent']).to eq('Konfig-server/1.0 http/1.0')
  expect(data['server']['vcs']['git']['url']).to eq(ENV['GITHUB_URL'])
end

Then('I should receive a missing config from gRPC:') do |_|
  expect(@response).to be_a(GRPC::NotFound)
end

Then('I should receive a invalid config from gRPC:') do |_|
  expect(@response).to be_a(GRPC::InvalidArgument)
end

Then('I should receive an internal error from gRPC:') do |_|
  expect(@response).to be_a(GRPC::Internal)
end
