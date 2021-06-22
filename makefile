.PHONY: gen-proto 

setup_dev_env: # Mac dev dev setup
	curl -L -o opa https://openpolicyagent.org/downloads/v0.29.4/opa_darwin_amd64
	chmod 755 ./opa

gen_proto:
	# Golang 
	buf generate

opa_test: # test all .rego files
	./opa test ./**/ -v