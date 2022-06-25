VERSION = 1.0.0

GOCC = go
BUILDDIR = build
EXECNAME ?= LALexer

PACKAGEPREFIX = pkg

ANTLRPREFIX = antlr
ANTLR ?= antlr4
DOCKER_IMAGE_NAME = la-builder

all: build

clean:
	@echo "Cleaning up build files"
	rm -rf ${PACKAGEPREFIX}/grammar ${BUILDDIR}

build: grammar
	@echo "Building executables"
	${GOCC} build -o ${BUILDDIR}/${EXECNAME}-${VERSION}

build/docker:
	@echo "Building executables in docker"
	docker build -t ${DOCKER_IMAGE_NAME} ./docker
	docker-compose -f ./docker/docker-compose.yml run \
		-w /la-build la-builder make


grammar: ${ANTLRPREFIX}/LA.g4
	@echo "Generating grammar files"
	${ANTLR} -Dlanguage=Go \
		-o ${PACKAGEPREFIX}/$@ \
		-package $@ \
		${ANTLRPREFIX}/LA.g4


.PHONY: main clean install uninstall build grammar
