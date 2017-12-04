run:
	export GOPATH="$$(readlink -f .)/.gopath" ; \
	export PATH="$$PATH:$$GOPATH/src/github.com/matlinuxer2/telegraf" ; \
	telegraf --config telegraf.conf

test:
	export GOPATH="$$(readlink -f .)/.gopath" ; \
	export PATH="$$PATH:$$GOPATH/src/github.com/matlinuxer2/telegraf" ; \
	telegraf --config telegraf.conf --test

build:
	export GOPATH="$$(readlink -f .)/.gopath" ; \
	export PATH="$$PATH:$$GOPATH/bin" ; \
	export PATH="$$PATH:$$GOPATH/src/github.com/matlinuxer2/telegraf" ; \
	go get -d github.com/influxdata/telegraf; \
	go get -d github.com/matlinuxer2/telegraf; \
	( cd $$GOPATH/src/github.com/matlinuxer2/telegraf; make );
