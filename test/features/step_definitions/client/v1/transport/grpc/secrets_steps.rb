# frozen_string_literal: true

When('I write secrets for {string} application') do |app|
  env = {
    'KONFIG_CONFIG_FILE' => ".config/#{app}.client.yaml"
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../konfig', 'secrets')
  pid = spawn(env, cmd, %i[out err] => ["reports/#{app}.client.log", 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have secrets for {string} application') do |_app|
  expect(File).to exist('reports/vault.secret')
  expect(File).to exist('reports/ssm.secret')
end
