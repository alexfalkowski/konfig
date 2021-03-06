# frozen_string_literal: true

Given('I have a {string} valid setup') do |source|
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = ".config/#{source}.server.config.yml"

  case source
  when 'git'
    Nonnative.configuration.processes[0].environment['KONFIG_GIT_TOKEN'] = ENV.fetch('GITHUB_TOKEN', nil)
  when 's3'
    files = [
      ['test/v1.8.0/staging/server.config.yml', '.config/test/v1.8.0/staging/server.config.yml'],
      ['test/v1.8.0/staging/eu/server.config.yml', '.config/test/v1.8.0/staging/eu/server.config.yml'],
      ['test/v1.8.0/staging/eu/de/server.config.yml', '.config/test/v1.8.0/staging/eu/de/server.config.yml']
    ]
    files.each { |f| Konfig.s3.write(f[0], File.read(f[1])) }
  end
end

Given('I have a {string} invalid setup') do |source|
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = ".config/invalid.#{source}.server.config.yml"

  case source
  when 's3'
    files = [
      'test/v1.8.0/staging/server.config.yml',
      'test/v1.8.0/staging/eu/server.config.yml',
      'test/v1.8.0/staging/eu/de/server.config.yml'
    ]
    files.each { |f| Konfig.s3.delete(f) }
  end
end
