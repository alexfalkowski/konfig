# frozen_string_literal: true

Given('I have a {string} valid setup') do |source|
  case source
  when 'git'
    Nonnative.configuration.processes[0].environment['KONFIG_GIT_TOKEN'] = ENV.fetch('GITHUB_TOKEN', nil)
  when 's3'
    files = [
      ['test/v1.7.0/staging/server.config.yml', '.config/test/v1.7.0/staging/server.config.yml'],
      ['test/v1.7.0/staging/eu/server.config.yml', '.config/test/v1.7.0/staging/eu/server.config.yml']
    ]
    files.each { |f| Konfig.s3.write(f[0], File.read(f[1])) }
  end
end

Given('I have a {string} invalid setup') do |source|
  case source
  when 'git'
    Nonnative.configuration.processes[0].environment['KONFIG_GIT_TOKEN'] = 'not_a_valid_token'
  when 'folder'
    Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = ".config/invalid.#{source}.server.config.yml"
    @config_set = true
  when 's3'
    files = [
      'test/v1.7.0/staging/server.config.yml',
      'test/v1.7.0/staging/eu/server.config.yml'
    ]
    files.each { |f| Konfig.s3.delete(f) }
  end
end

Given('I have {string} as the config file') do |source|
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = ".config/#{source}.server.config.yml" unless @config_set
end
