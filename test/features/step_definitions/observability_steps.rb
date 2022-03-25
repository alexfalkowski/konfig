# frozen_string_literal: true

When('the system requests the {string} with HTTP') do |name|
  @response = Konfig.observability.send(name)
end

When('the system requests the health status with gRPC') do
  request = Grpc::Health::V1::HealthCheckRequest.new
  @response = Konfig.health_grpc.check(request)
end

Then('the system should respond with a healthy status with HTTP') do
  expect(@response.code).to eq(200)
  expect(JSON.parse(@response.body)).to eq('status' => 'SERVING')
end

Then('the system should respond with a healthy status with gRPC') do
  expect(@response.status).to eq(:SERVING)
end

Then('the system should respond with metrics') do
  expect(@response.code).to eq(200)
  expect(@response.body).to include('go_info')
end
