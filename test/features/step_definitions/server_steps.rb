# frozen_string_literal: true

When('I request {string} app from {string} env with HTTP') do |app, env|
  headers = {
    request_id: SecureRandom.uuid,
    user_agent: Konfig.server_config['transport']['http']['user_agent']
  }

  @response = Konfig.server_http.get_config(app, env, headers)
end

Then('I should receive a valid config from {string} app with {string} env from HTTP') do |app, env|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  resp['data'] = YAML.safe_load(Base64.decode64(resp['data']))

  expect(resp).to eq('application' => app, 'environment' => env, 'contentType' => 'application/yaml',
                     'data' => Konfig.server_config)
end

When('I request {string} app from {string} env with gRPC') do |app, env|
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'ua' => Konfig.server_config['transport']['grpc']['user_agent']
  }

  request = Konfig::V1::GetConfigRequest.new(application: app, environment: env)
  @response = Konfig.server_grpc.get_config(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

Then('I should receive a valid config from {string} app with {string} env from gRPC') do |app, env|
  expect(@response.application).to eq(app)
  expect(@response.environment).to eq(env)
  expect(@response.content_type).to eq('application/yaml')
  expect(YAML.safe_load(@response.data)).to eq(Konfig.server_config)
end
