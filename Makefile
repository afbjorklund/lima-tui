
GO = go

PROGRAM = lima-tui

.PHONY: bin/$(PROGRAM)
bin/$(PROGRAM): cmd
	$(GO) build -o $@ ./$^/...
