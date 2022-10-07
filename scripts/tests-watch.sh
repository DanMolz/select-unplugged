#!/bin/bash

main() {
	set -euo pipefail
	cd "$(dirname "$0")/.."

	while : ; do
		go test ./... || true
		date
		read -r _ 
	done < <(fswatch -o .)
}

main "$@"