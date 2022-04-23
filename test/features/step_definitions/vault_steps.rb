# frozen_string_literal: true

Given('I have key {string} with {string} value in vault') do |key, value|
  Konfig.vault.write(key, value)
end

Given('I have no key {string} in vault') do |key|
  Konfig.vault.delete(key)
end
