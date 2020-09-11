#!/usr/bin/env bash

set -e

patch_files_directory="."
if [ -n "${1}" ]; then
    patch_files_directory="${1}"
fi

while read -e patch_file; do
    patch --batch --strip=1 < "${patch_file}"
done < <(find "${patch_files_directory}" -name '*.patch')
