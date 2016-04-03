
install:
	@godep go install

build:
	@godep go build -o random .

deps:
	@godep get .

clean:
	@rm $(binary)

.PHONY: build deps clean run
