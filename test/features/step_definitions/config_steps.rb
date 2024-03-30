# frozen_string_literal: true

Given('I have a {string} valid setup') do |source|
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = ".config/#{source}.server.yaml"

  case source
  when 'git'
    Nonnative.configuration.processes[0].environment['KONFIG_GIT_TOKEN'] = ENV.fetch('GITHUB_TOKEN', nil)
  when 's3'
    Nonnative.configuration.processes[0].environment['AWS_URL'] = 'http://localhost:4566'

    files = [
      ['test/v1.11.0/staging/server.yaml', '.config/test/v1.11.0/staging/server.yaml'],
      ['test/v1.11.0/staging/eu/server.yaml', '.config/test/v1.11.0/staging/eu/server.yaml'],
      ['test/v1.11.0/staging/eu/de/server.yaml', '.config/test/v1.11.0/staging/eu/de/server.yaml'],
      ['test/v1.11.0/staging/server.toml', '.config/test/v1.11.0/staging/server.toml'],
      ['test/v1.11.0/staging/eu/server.toml', '.config/test/v1.11.0/staging/eu/server.toml'],
      ['test/v1.11.0/staging/eu/de/server.toml', '.config/test/v1.11.0/staging/eu/de/server.toml']
    ]
    files.each { |f| Konfig.s3.write(f[0], File.read(f[1])) }
  end
end

Given('I have a {string} invalid setup') do |source|
  Nonnative.configuration.processes[0].environment['CONFIG_FILE'] = ".config/invalid.#{source}.server.yaml"

  case source
  when 'git'
    Nonnative.configuration.processes[0].environment['KONFIG_GIT_TOKEN'] = 'invalid_token'
  when 's3'
    Nonnative.configuration.processes[0].environment['AWS_URL'] = 'does_not_exist'

    files = [
      'test/v1.11.0/staging/server.yaml', 'test/v1.11.0/staging/eu/server.yaml', 'test/v1.11.0/staging/eu/de/server.yaml',
      'test/v1.11.0/staging/server.toml', 'test/v1.11.0/staging/eu/server.toml', 'test/v1.11.0/staging/eu/de/server.toml'
    ]
    files.each { |f| Konfig.s3.delete(f) }
  end
end
