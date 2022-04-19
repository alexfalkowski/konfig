lint:
	buf lint

format:
	buf format -w

generate:
	buf generate

update-all:
	buf mod update

breaking:
	buf breaking --against 'https://github.com/alexfalkowski/konfig.git#branch=master,subdir=api' || true
