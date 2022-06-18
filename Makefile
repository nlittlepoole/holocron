#
# Holocron
#
HOLOCRON_NAME?=
HOLOCRON_GATEKEEPER?=
HOLOCRON_TREASURE?=
HOLOCRON_ASCERTAINMENT?=
HOLOCRON_SALT?=foobar
HOLOCRON_OUTDIR?=.build


ifeq ($(OS),Windows_NT)
	uname_S := Windows
else
	uname_S := $(shell uname -s)
endif

all: forge

forge: $(HOLOCRON_OUTDIR)/forge
	$(HOLOCRON_OUTDIR)/forge -n "$(HOLOCRON_NAME)" -g "$(HOLOCRON_GATEKEEPER)" -a "$(HOLOCRON_ASCERTAINMENT)" -t "$(HOLOCRON_TREASURE)" -s "$(HOLOCRON_SALT)" -o "$(HOLOCRON_OUTDIR)"

$(HOLOCRON_OUTDIR)/forge:
	go build -o $(HOLOCRON_OUTDIR)/forge cmd/forge/main.go 