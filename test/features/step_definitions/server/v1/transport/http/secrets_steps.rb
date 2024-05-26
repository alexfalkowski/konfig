# frozen_string_literal: true

When('I request secrets with HTTP:') do |table|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Konfig-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }.merge(Konfig.token),
    read_timeout: 10, open_timeout: 10
  }

  @response = Konfig::V1.server_http.get_secrets(table.rows_hash, opts)
end

Then('I should receive valid secrets from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  secrets = resp['secrets'].transform_values { |v| Base64.decode64(v) }

  expect(secrets).to eq(table.rows_hash)
end
