# frozen_string_literal: true

When('I request a config with HTTP:') do |table|
  rows = table.rows_hash
  auth = service_token(Nonnative.configurations('.config/existing.client.yaml'))
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Konfig-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }.merge(auth),
    read_timeout: 10, open_timeout: 10
  }
  params = {
    app: rows['app'], ver: rows['ver'], env: rows['env'], continent: rows['continent'],
    country: rows['country'], cmd: rows['cmd'], kind: rows['kind']
  }

  @response = Konfig::V1.server_http.get_config(params, opts)
end

Then('I should receive a valid config from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  config = resp['config']
  rows = table.rows_hash
  data = Konfig.load_config(rows['kind'], Base64.decode64(config['data']))

  expect(resp['meta'].length).to be > 0
  expect(config['application']).to eq(rows['app'])
  expect(config['version']).to eq(rows['ver'])
  expect(config['environment']).to eq(rows['env'])
  expect(config['continent']).to eq(rows['continent'])
  expect(config['country']).to eq(rows['country'])
  expect(config['command']).to eq(rows['cmd'])
  expect(config['kind']).to eq(rows['kind'])
  expect(data['transport']['http']['user_agent']).to eq('Konfig-server/1.0 http/1.0')
  expect(data['transport']['grpc']['user_agent']).to eq('Konfig-server/1.0 grpc/1.0')
  expect(data['source']['git']['url']).to eq(ENV.fetch('GITHUB_URL', nil))
end

Then('I should receive a valid config with missing provider data from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  config = resp['config']
  rows = table.rows_hash
  data = Konfig.load_config(rows['kind'], Base64.decode64(config['data']))

  expect(resp['meta'].length).to be > 0
  expect(config['application']).to eq(rows['app'])
  expect(config['version']).to eq(rows['ver'])
  expect(config['environment']).to eq(rows['env'])
  expect(config['continent']).to eq(rows['continent'])
  expect(config['continent']).to eq(rows['continent'])
  expect(config['country']).to eq(rows['country'])
  expect(config['kind']).to eq(rows['kind'])
  expect(data['transport']['http']['user_agent']).to eq('/secret/data/transport/http/user_agent')
  expect(data['transport']['grpc']['user_agent']).to eq('/secret/data/transport/grpc/user_agent')
  expect(data['source']['git']['url']).to eq(ENV.fetch('GITHUB_URL', nil))
end

Then('I should receive a missing config from HTTP') do
  expect(@response.code).to eq(404)
end

Then('I should receive a invalid config from HTTP') do
  expect(@response.code).to eq(400)
end

Then('I should receive an internal error from HTTP') do
  expect(@response.code).to eq(500)
end
