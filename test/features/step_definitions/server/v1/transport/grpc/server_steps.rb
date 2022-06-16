# frozen_string_literal: true

When('I request a config with gRPC:') do |table|
  @response = request_with_grpc(table)
end

When('I request a config with gRPC {int} times:') do |times, table|
  times.times do
    @response = request_with_grpc(table)
  end
end

Then('I should receive a valid config from gRPC:') do |table|
  data = YAML.safe_load(@response.config.data)
  rows = table.rows_hash

  expect(@response.config.application).to eq(rows['app'])
  expect(@response.config.version).to eq(rows['ver'])
  expect(@response.config.environment).to eq(rows['env'])
  expect(@response.config.continent).to eq(rows['continent'])
  expect(@response.config.country).to eq(rows['country'])
  expect(@response.config.command).to eq(rows['cmd'])
  expect(@response.config.content_type).to eq('application/yaml')
  expect(data['transport']['http']['user_agent']).to eq('Konfig-server/1.0 http/1.0')
  expect(data['transport']['grpc']['user_agent']).to eq('Konfig-server/1.0 grpc/1.0')
  expect(data['server']['v1']['source']['git']['url']).to eq(ENV.fetch('GITHUB_URL', nil))
end

Then('I should receive a valid config with missing provider data from gRPC:') do |table|
  data = YAML.safe_load(@response.config.data)
  rows = table.rows_hash

  expect(@response.config.application).to eq(rows['app'])
  expect(@response.config.version).to eq(rows['ver'])
  expect(@response.config.environment).to eq(rows['env'])
  expect(@response.config.continent).to eq(rows['continent'])
  expect(@response.config.country).to eq(rows['country'])
  expect(@response.config.command).to eq(rows['cmd'])
  expect(@response.config.content_type).to eq('application/yaml')
  expect(data['transport']['http']['user_agent']).to eq('/secret/data/transport/http/user_agent')
  expect(data['transport']['grpc']['user_agent']).to eq('/secret/data/transport/grpc/user_agent')
  expect(data['server']['v1']['source']['git']['url']).to eq(ENV.fetch('GITHUB_URL', nil))
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

def request_with_grpc(table)
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id, 'ua' => Konfig.server_config(rows['source'])['transport']['grpc']['user_agent'] }

  request = Konfig::V1::GetConfigRequest.new(application: rows['app'], version: rows['ver'], environment: rows['env'],
                                             continent: rows['continent'], country: rows['country'], command: rows['cmd'])
  Konfig::V1.server_grpc.get_config(request, { metadata: metadata })
rescue StandardError => e
  e
end
