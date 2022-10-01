#!/bin/bash

main() {
	set -euo pipefail
	cd "$(dirname "$0")/.."

	while : ; do
		go test ./... || true
		read -r _ 
	done < <(fswatch -o .)
}

main "$@"