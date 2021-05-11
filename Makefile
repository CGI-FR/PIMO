# Copyright (C) 2021 CGI France
#
# This file is part of PIMO.
#
# PIMO is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# PIMO is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with PIMO.  If not, see <http://www.gnu.org/licenses/>.

##
## This project uses neon as a build tool, all of the make rules are mapped to the neon build in build.yml
##

SHELL := /bin/zsh # Use zsh syntax

.PHONY: init
init:
	@touch ~/.dockerhub.yml
	@chmod 0600 ~/.dockerhub.yml
	@touch ~/.github.yml
	@chmod 0600 ~/.github.yml
	@git config --local core.hooksPath githooks

.PHONY: warning
warning:
	@echo "This project uses neon as a build tool, all of the make rules are mapped to the neon build in build.yml"
	@echo "You should use the 'neon' command directly, it works the same way as make"
	@echo "example: neon build"

.PHONY: help
.DEFAULT_GOAL := help
help: warning
	@neon help

.PHONY: info
info: warning
	@neon info

.PHONY: promote
promote: warning
	@neon promote

.PHONY: refresh
refresh: warning
	@neon refresh

# tidy is an alias of refresh
.PHONY: tidy
tidy: refresh

# build is an alias of compile
.PHONY: build-%
build-%: warning
	@neon -props "{buildpaths: ["cmd/$*"]}" compile

# build is an alias of compile
.PHONY: build
build: warning
	@neon compile

.PHONY: compile-%
compile-%: warning
	@neon -props "{buildpaths: ["cmd/$*"]}" compile

.PHONY: compile
compile: warning
	@neon compile

.PHONY: run-%
run-%: build-%
	bin/$*

.PHONY: run
run: $(patsubst cmd/%,run-%,$(wildcard cmd/*)) ## Build and execute a binary

.PHONY: test
test: warning
	@neon test

.PHONY: lint
lint: warning
	@neon lint

.PHONY: release-%
release-%: warning
	@neon -props "{buildpaths: ["cmd/$*"]}" release

.PHONY: release
release: warning
	@neon release

.PHONY: test-int
test-int: warning
	@neon test-int

.PHONY: publish-%
publish-%: warning
	@neon -props "{buildpaths: ["cmd/$*"]}" publish

.PHONY: publish
publish: warning
	@neon publish

.PHONY: snapshot
snapshot:  ## Create snapshot binaries
	@neon -props '{snapshot: true}' publish


.PHONY: docker
docker: warning
	@neon docker

.PHONY: docker-tag
docker-tag: warning
	@neon docker-tag

.PHONY: docker-push
docker-push: warning
	@neon docker-push

.PHONY: license-%
license-%: warning
	@neon -props "{buildpaths: ["cmd/$*"]}" license

.PHONY: license
license: warning
	@neon license
