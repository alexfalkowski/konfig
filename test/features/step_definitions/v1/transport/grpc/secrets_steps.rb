# frozen_string_literal: true

When('I request secrets with gRPC:') do |table|
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }.merge(Konfig.token)
  request = Konfig::V1::GetSecretsRequest.new(secrets: table.rows_hash)
  @response = Konfig::V1.server_grpc.get_secrets(request, { metadata: })
rescue StandardError => e
  @response = e
end

Then('I should receive valid secrets from gRPC:') do |table|
  expect(@response.meta.length).to be > 0
  expect(@response.secrets).to eq(table.rows_hash)
end
