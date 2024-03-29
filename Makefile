include bin/build/make/service.mak
include bin/build/make/git.mak

features-grpc:
	$(MAKE) feature=features tags=@grpc features

features-http:
	$(MAKE) feature=features tags=@http features
