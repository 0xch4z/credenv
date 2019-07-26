VERSION := 0.0.1
REPO    := github.com/charliekenney23/credenv

HACK_DIR     := ./hack
BUILD_SCRIPT := $(HACK_DIR)/build.sh

build:
	CREDENV_VERSION=$(VERSION) bash $(BUILD_SCRIPT)

dev:
	CREDENV_DEV=1 bash $(BUILD_SCRIPT)
