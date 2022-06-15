# frozen_string_literal: true

Given('I have the following provider information:') do |table|
  table.hashes.each do |h|
    Konfig.send(h['provider']).write(h['key'], h['value'])
  end
end

Given('I do not have the following provider information:') do |table|
  table.hashes.each do |h|
    Konfig.send(h['provider']).delete(h['key'])
  end
end
