# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: konfig.v1.proto

require 'google/protobuf'

require 'google/api/annotations_pb'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("konfig.v1.proto", :syntax => :proto3) do
    add_message "konfig.v1.GetConfigRequest" do
      optional :application, :string, 1
      optional :version, :string, 2
      optional :environment, :string, 3
      optional :command, :string, 4
    end
    add_message "konfig.v1.GetConfigResponse" do
      optional :application, :string, 1
      optional :version, :string, 2
      optional :environment, :string, 3
      optional :command, :string, 4
      optional :content_type, :string, 5
      optional :data, :bytes, 6
    end
  end
end

module Konfig
  module V1
    GetConfigRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("konfig.v1.GetConfigRequest").msgclass
    GetConfigResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("konfig.v1.GetConfigResponse").msgclass
  end
end
