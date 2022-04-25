# frozen_string_literal: true

When('I download the configuration for {string} application') do |app|
  @config_file = ".config/#{app}.server.config.yml"

  File.delete(@config_file) if File.exist?(@config_file)

  env = {
    'CONFIG_FILE' => ".config/#{app}.client.config.yml",
    'APP_CONFIG_FILE' => @config_file
  }
  cmd = Nonnative.go_executable('reports', '../konfig', 'client')
  pid = spawn(env, cmd, %i[out err] => ["reports/#{app}.client.log", 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have a configuration for {string} application') do |_|
  expect(File.file?(@config_file)).to be true
  expect(@status.exitstatus).to eq(0)
end

Then('I should not have a configuration for {string} application') do |_|
  expect(File.file?(@config_file)).to be false
  expect(@status.exitstatus).to eq(1)
end
