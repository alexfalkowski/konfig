# frozen_string_literal: true

Given('I have already downloaded the configuration for {string} application') do |app|
  status = download(app)
  expect(status.exitstatus).to eq(0)

  @mtime = File.stat("reports/#{app}.server.config.yml").mtime
end

When('I download the configuration for {string} application') do |app|
  @status = download(app)
end

Then('I should have a configuration for {string} application') do |app|
  expect(File.file?("reports/#{app}.server.config.yml")).to be true
  expect(@status.exitstatus).to eq(0)
end

Then('I should not have a configuration for {string} application') do |app|
  expect(File.file?("reports/#{app}.server.config.yml")).to be false
  expect(@status.exitstatus).to eq(1)
end

Then('I should not have written a config for {string} application') do |app|
  file = "reports/#{app}.server.config.yml"

  expect(File.file?(file)).to be true
  expect(@status.exitstatus).to eq(0)
  expect(File.stat(file).mtime).to eq @mtime
end

def download(app)
  env = {
    'CONFIG_FILE' => ".config/#{app}.client.config.yml",
    'APP_CONFIG_FILE' => "reports/#{app}.server.config.yml"
  }
  cmd = Nonnative.go_executable('reports', '../konfig', 'client')
  pid = spawn(env, cmd, %i[out err] => ["reports/#{app}.client.log", 'a'])

  _, status = Process.waitpid2(pid)
  status
end
