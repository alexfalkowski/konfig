# frozen_string_literal: true

Given('I should see {string} as unhealthy') do |service|
  wait_for do
    @response = Konfig.observability.health
    @response.code
  end.to eq(503)

  expect(@response.body).to include(service)
end

When('the system requests the {string} with HTTP') do |name|
  @response = Konfig.observability.send(name)
end

Then('the system should respond with a healthy status with HTTP') do
  expect(@response.code).to eq(200)
  expect(JSON.parse(@response.body)).to eq('status' => 'SERVING')
end

Then('the system should respond with metrics') do
  expect(@response.code).to eq(200)
  expect(@response.body).to include('go_info')
end

Then('I should see {string} as healthy') do |service|
  wait_for do
    @response = Konfig.observability.health
    @response.code
  end.to eq(200)

  expect(@response.body).to_not include(service)
end
