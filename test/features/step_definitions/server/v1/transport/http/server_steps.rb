# frozen_string_literal: true

When('I request a config with HTTP:') do |table|
  @response = request_with_http(table)
end

When('I request a config with HTTP {int} times:') do |times, table|
  times.times do
    @response = request_with_http(table)
  end
end

Then('I should receive a valid config from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  config = resp['config']
  data = YAML.safe_load(Base64.decode64(config['data']))
  rows = table.rows_hash

  expect(config['application']).to eq(rows['app'])
  expect(config['version']).to eq(rows['ver'])
  expect(config['environment']).to eq(rows['env'])
  expect(config['cluster']).to eq(rows['cluster'])
  expect(config['command']).to eq(rows['cmd'])
  expect(config['contentType']).to eq('application/yaml')
  expect(data['transport']['http']['user_agent']).to eq('Konfig-server/1.0 http/1.0')
  expect(data['server']['v1']['source']['git']['url']).to eq(ENV.fetch('GITHUB_URL', nil))
end

Then('I should receive a valid config with missing vault value from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  config = resp['config']
  data = YAML.safe_load(Base64.decode64(config['data']))
  rows = table.rows_hash

  expect(config['application']).to eq(rows['app'])
  expect(config['version']).to eq(rows['ver'])
  expect(config['environment']).to eq(rows['env'])
  expect(config['cluster']).to eq(rows['cluster'])
  expect(config['command']).to eq(rows['cmd'])
  expect(config['contentType']).to eq('application/yaml')
  expect(data['transport']['http']['user_agent']).to eq('secret/data/transport/http/user_agent')
  expect(data['server']['v1']['source']['git']['url']).to eq(ENV.fetch('GITHUB_URL', nil))
end

Then('I should receive a missing config from HTTP:') do |_|
  expect(@response.code).to eq(404)
end

Then('I should receive a invalid config from HTTP:') do |_|
  expect(@response.code).to eq(400)
end

Then('I should receive an internal error from HTTP:') do |_|
  expect(@response.code).to eq(500)
end

def request_with_http(table)
  rows = table.rows_hash
  headers = { request_id: SecureRandom.uuid, user_agent: Konfig.server_config(rows['source'])['transport']['grpc']['user_agent'] }

  params = { app: rows['app'], ver: rows['ver'], env: rows['env'], cluster: rows['cluster'], cmd: rows['cmd'] }
  Konfig::V1.server_http.get_config(params, headers)
end
