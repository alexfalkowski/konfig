# frozen_string_literal: true

Given('I have a valid vcs token') do
  Nonnative.configuration.processes[0].environment['KONFIG_GIT_TOKEN'] = ENV['GITHUB_TOKEN']
end

Given('I have a misconfigured vcs token') do
  Nonnative.configuration.processes[0].environment['KONFIG_GIT_TOKEN'] = 'not_a_valid_token'
end

When('I request {string} app with {string} ver from {string} env and {string} cmd with HTTP') do |app, ver, env, cmd|
  headers = {
    request_id: SecureRandom.uuid,
    user_agent: Konfig.server_config['transport']['http']['user_agent']
  }

  @response = Konfig.server_http.get_config(app, ver, env, cmd, headers)
end

When('I request {string} app with {string} ver from {string} env and {string} cmd with gRPC') do |app, ver, env, cmd|
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'ua' => Konfig.server_config['transport']['grpc']['user_agent']
  }

  request = Konfig::V1::GetConfigRequest.new(application: app, version: ver, environment: env, command: cmd)
  @response = Konfig.server_grpc.get_config(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

Then('I should receive a valid cfg from {string} app with {string} ver and {string} env and {string} cmd from HTTP') do |app, ver, env, cmd|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  resp['data'] = YAML.safe_load(Base64.decode64(resp['data']))

  expect(resp).to eq('application' => app, 'version' => ver, 'environment' => env, 'command' => cmd,
                     'contentType' => 'application/yaml', 'data' => Konfig.server_config)
end

Then('I should receive a missing cfg from {string} app with {string} ver and {string} env and {string} cmd from HTTP') do |_, _, _, _|
  expect(@response.code).to eq(404)
end

Then('I should receive an invalid cfg from {string} app with {string} ver and {string} env and {string} cmd from HTTP') do |_, _, _, _|
  expect(@response.code).to eq(400)
end

Then('I should receive a valid cfg from {string} app with {string} ver and {string} env and {string} cmd from gRPC') do |app, ver, env, cmd|
  expect(@response.application).to eq(app)
  expect(@response.version).to eq(ver)
  expect(@response.environment).to eq(env)
  expect(@response.command).to eq(cmd)
  expect(@response.content_type).to eq('application/yaml')
  expect(YAML.safe_load(@response.data)).to eq(Konfig.server_config)
end

Then('I should receive a missing cfg from {string} app with {string} ver and {string} env and {string} cmd from gRPC') do |_, _, _, _|
  expect(@response).to be_a(GRPC::NotFound)
end

Then('I should receive an invalid cfg from {string} app with {string} ver and {string} env and {string} cmd from gRPC') do |_, _, _, _|
  expect(@response).to be_a(GRPC::InvalidArgument)
end
