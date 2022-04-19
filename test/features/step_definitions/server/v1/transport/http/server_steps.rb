# frozen_string_literal: true

When('I request a config with HTTP:') do |table|
  rows = table.rows_hash
  headers = {
    request_id: SecureRandom.uuid,
    user_agent: Konfig.server_config(rows['source'])['transport']['grpc']['user_agent']
  }

  rows['cluster'] ||= '*'

  params = { app: rows['app'], ver: rows['ver'], env: rows['env'], cluster: rows['cluster'], cmd: rows['cmd'] }
  @response = Konfig.server_http.get_config(params, headers)
end

Then('I should receive a valid config from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  data = YAML.safe_load(Base64.decode64(resp['data']))
  rows = table.rows_hash

  rows['cluster'] ||= '*'

  expect(resp['application']).to eq(rows['app'])
  expect(resp['version']).to eq(rows['ver'])
  expect(resp['environment']).to eq(rows['env'])
  expect(resp['cluster']).to eq(rows['cluster'])
  expect(resp['command']).to eq(rows['cmd'])
  expect(resp['contentType']).to eq('application/yaml')
  expect(data['transport']['http']['user_agent']).to eq('Konfig-server/1.0 http/1.0')
  expect(data['server']['vcs']['git']['url']).to eq(ENV['GITHUB_URL'])
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
