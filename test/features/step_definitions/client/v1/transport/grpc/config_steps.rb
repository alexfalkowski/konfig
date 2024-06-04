# frozen_string_literal: true

When('I download the configuration for {string} application') do |app|
  env = {
    'KONFIG_CONFIG_FILE' => ".config/#{app}.client.yaml",
    'KONFIG_APP_CONFIG_FILE' => "reports/#{app}.server.yaml"
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../konfig', 'config')
  pid = spawn(env, cmd, %i[out err] => ["reports/#{app}.client.log", 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have a configuration for {string} application') do |app|
  expect(File.file?("reports/#{app}.server.yaml")).to be true
  expect(@status.exitstatus).to eq(0)
end

Then('I should not have a configuration for {string} application') do |app|
  expect(File.file?("reports/#{app}.server.yaml")).to be false
  expect(@status.exitstatus).to eq(1)
end