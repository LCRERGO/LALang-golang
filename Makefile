VERSION = 1.0.0

GOCC = go
CC = gcc
BUILDDIR = build
EXECNAME ?= LALexer

PACKAGEPREFIX = pkg
TESTPREFIX = tests
TESTSDIR = test-results

AUTOTESTER=compiladores-corretor-automatico-1.0-SNAPSHOT-jar-with-dependencies.jar

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
		-w /la-build ${DOCKER_IMAGE_NAME} make

test: build
	@echo "Testing executables"
	java -jar ${TESTPREFIX}/${AUTOTESTER} \
		build/${EXECNAME}-${VERSION} \
		${CC} \
		./${TESTSDIR} \
		${TESTPREFIX}/test-cases \
		"754757" \
		"t2"

test/docker: build/docker
	@echo "Testing executables in docker"
	docker build -t ${DOCKER_IMAGE_NAME} ./docker
	docker-compose -f ./docker/docker-compose.yml run \
		-w /la-build ${DOCKER_IMAGE_NAME} \
		java -jar ${TESTPREFIX}/${AUTOTESTER} \
		build/${EXECNAME}-${VERSION} \
		${CC} \
		/tmp/la-build/${TESTSDIR} \
		${TESTPREFIX}/test-cases \
		"754757" \
		"t2"

grammar: ${ANTLRPREFIX}/LA.g4
	@echo "Generating grammar files"
	${ANTLR} -Dlanguage=Go \
		-o ${PACKAGEPREFIX}/$@ \
		-package $@ \
		${ANTLRPREFIX}/LA.g4


.PHONY: main clean install uninstall build test grammar
