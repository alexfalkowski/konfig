# frozen_string_literal: true

def generate_service_token(cfg, kind, audience, token)
  auth = Nonnative::Header.auth_bearer(token)
  metadata = { 'request-id' => SecureRandom.uuid, 'user-agent' => 'Konfig-ruby-client/1.0 gRPC/1.0' }.merge(auth)
  request = Auth::V1::GenerateServiceTokenRequest.new(kind:, audience:)
  stub = Auth::V1::Service::Stub.new(cfg.auth.client.v1.host, :this_channel_is_insecure, channel_args: Konfig.user_agent)

  stub.generate_service_token(request, { metadata: })
end

def service_token(cfg)
  token = generate_service_token(cfg, 'jwt', 'konfig', cfg.auth.client.v1.access)
  bearer = token.token.bearer

  Nonnative::Header.auth_bearer(bearer)
end
