#
# Horcrux
#
HORCRUX_PROMPT?=
HORCRUX_TREASURE?=
HORCRUX_ENCRYPTED?=


ifeq ($(OS),Windows_NT)
	uname_S := Windows
else
	uname_S := $(shell uname -s)
endif

all: build serve

serve:
	docker build --target horcrux_server -t horcrux_server .
	docker run -e HORCRUX_PROMPT="$(HORCRUX_PROMPT)" -e HORCRUX_TREASURE="$(HORCRUX_TREASURE)" -e HORCRUX_ENCRYPTED="$(HORCRUX_ENCRYPTED)" -itp "8090:8090" horcrux_server

build:
	mkdir -p .build
	docker build --target horcrux_build -t horcrux_build .
	docker run -itv "$(PWD)/.build:/build/" -e HORCRUX_PROMPT="$(HORCRUX_PROMPT)" -e HORCRUX_TREASURE="$(HORCRUX_TREASURE)" -e HORCRUX_ENCRYPTED="$(HORCRUX_ENCRYPTED)" horcrux_build


