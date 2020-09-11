.PHONY: generate clean integration patches

generate:
	swagger generate client --target=./netbox --spec=./swagger.json --copyright-file=./copyright_header.txt

clean:
	rm -rf netbox/client netbox/models

integration:
	go test ./... -tags=integration

patches:
	patches/apply-patches.sh patches
