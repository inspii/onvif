.PHONY: specs
specs:
	rm -rf specs
	git clone git@github.com:onvif/specs.git
	cd specs && git checkout tags/22.12 -b 22.12
	rm -rf specs/.git
	rm -rf specs/doc
	rm specs/guidelines.md specs/LICENSE.md specs/README.md

gen:
	go run tools/wsdlgen/cmd/main.go -o devicemgmt.go -p specs specs/wsdl/ver10/device/wsdl/devicemgmt.wsdl