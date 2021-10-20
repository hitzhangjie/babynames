bindata:
	go get -u github.com/hitzhangjie/codeblocks/bindata
	bindata -input dicts -gopkg assets -output assets/dicts.go

.PHONY: bindata
