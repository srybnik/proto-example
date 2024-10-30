APP_NAME=example
PROJECT_API=api
PROJECT_CLIENTS=pkg/api
PROJECT_IMPLEMENTATIONS=internal/app/api
SERVICES_ROOT=$(shell echo $(PROJECT_CLIENTS) | perl -F/ -lane 'print "../"x scalar(@F)')

SERVICES=$(shell ls -1 $(PROJECT_API) | grep \.proto | sed s/\.proto//)

IMPLEMENTATION_TYPE_NAME="Implementation"

LOCAL_BIN:=$(CURDIR)/bin
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

PKG:=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,$\
	Mgoogle/protobuf/api.proto=github.com/gogo/protobuf/types,$\
	Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/types,$\
	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,$\
	Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,$\
	Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,$\
	Mgoogle/protobuf/source_context.proto=github.com/gogo/protobuf/types,$\
	Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,$\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,$\
	Mgoogle/protobuf/type.proto=github.com/gogo/protobuf/types,$\
	Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types


.PHONY: bin-deps
bin-deps: ; $(info $(M) install bin deps)
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/gogo/protobuf/protoc-gen-gofast
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/utrack/clay/v2/cmd/protoc-gen-goclay


.PHONY: gen
gen: bin-deps ; $(info $(M) protoc gen)
	$(Q) for srv in $(SERVICES); do \
	    echo "Generate $(CURDIR)/$(PROJECT_CLIENTS)/$$srv" && \
	    echo "Implementation $(SERVICES_ROOT)../../$(PROJECT_IMPLEMENTATIONS)/$$srv" && \
		mkdir -p $(CURDIR)/$(PROJECT_CLIENTS)/$$srv && \
		cd $(CURDIR)/$(PROJECT_CLIENTS)/$$srv && \
		protoc --plugin=protoc-gen-goclay=$(LOCAL_BIN)/protoc-gen-goclay \
			--plugin=protoc-gen-gofast=$(LOCAL_BIN)/protoc-gen-gofast \
			-I$(SERVICES_ROOT)../api/:$(CURDIR)/vendor.pb \
			--gofast_out=$(PKG),plugins=grpc:. \
			--goclay_out=$(PKG),impl=true,impl_service_sub_dir=false,impl_path=$(SERVICES_ROOT)../$(PROJECT_IMPLEMENTATIONS)/$$srv,impl_type_name_tmpl=$(IMPLEMENTATION_TYPE_NAME):. \
			$(SERVICES_ROOT)../$(PROJECT_API)/$$srv.proto ; \
	done


.PHONY: run
run:
	go run cmd/example/main.go
