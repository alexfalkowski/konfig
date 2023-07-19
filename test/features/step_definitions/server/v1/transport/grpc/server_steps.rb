# frozen_string_literal: true

When('I request a config with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id, 'ua' => Konfig.server_config(rows['source'])['transport']['grpc']['user_agent'] }

  request = Konfig::V1::GetConfigRequest.new(application: rows['app'], version: rows['ver'], environment: rows['env'],
                                             continent: rows['continent'], country: rows['country'], command: rows['cmd'],
                                             kind: rows['kind'])
  @response = Konfig::V1.server_grpc.get_config(request, { metadata: })
rescue StandardError => e
  @response = e
end

Then('I should receive a valid config from gRPC:') do |table|
  rows = table.rows_hash
  data = Konfig.load_config(rows['kind'], @response.config.data)

  expect(@response.config.application).to eq(rows['app'])
  expect(@response.config.version).to eq(rows['ver'])
  expect(@response.config.environment).to eq(rows['env'])
  expect(@response.config.continent).to eq(rows['continent'])
  expect(@response.config.country).to eq(rows['country'])
  expect(@response.config.command).to eq(rows['cmd'])
  expect(@response.config.kind).to eq(rows['kind'])
  expect(data['transport']['http']['user_agent']).to eq('Konfig-server/1.0 http/1.0')
  expect(data['transport']['grpc']['user_agent']).to eq('Konfig-server/1.0 grpc/1.0')
  expect(data['source']['git']['url']).to eq(ENV.fetch('GITHUB_URL', nil))
end

Then('I should receive a valid config with missing provider data from gRPC:') do |table|
  rows = table.rows_hash
  data = Konfig.load_config(rows['kind'], @response.config.data)

  expect(@response.config.application).to eq(rows['app'])
  expect(@response.config.version).to eq(rows['ver'])
  expect(@response.config.environment).to eq(rows['env'])
  expect(@response.config.continent).to eq(rows['continent'])
  expect(@response.config.country).to eq(rows['country'])
  expect(@response.config.command).to eq(rows['cmd'])
  expect(@response.config.kind).to eq(rows['kind'])
  expect(data['transport']['http']['user_agent']).to eq('/secret/data/transport/http/user_agent')
  expect(data['transport']['grpc']['user_agent']).to eq('/secret/data/transport/grpc/user_agent')
  expect(data['source']['git']['url']).to eq(ENV.fetch('GITHUB_URL', nil))
end

Then('I should receive a missing config from gRPC') do
  expect(@response).to be_a(GRPC::NotFound)
end

Then('I should receive a invalid config from gRPC') do
  expect(@response).to be_a(GRPC::InvalidArgument)
end

Then('I should receive an internal error from gRPC') do
  expect(@response).to be_a(GRPC::Internal)
end
