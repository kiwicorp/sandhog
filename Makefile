# sandhog makefile

# PROTOS=registry sandhog
# internal/api/$(PROTOS)/*.go: $(SRC_PROTOS)

SRC=main.go	\
	internal/api/selftechio/*/*.go \
	internal/links/*.go \
	internal/tunnel/*.go
BIN=bin/sandhog

all: $(BIN)

$(BIN): $(SRC)
	./scripts/build.sh $(BIN)

clean:
	rm -r bin
.PHONY: clean
