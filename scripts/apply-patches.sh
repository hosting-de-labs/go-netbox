#!/usr/bin/env bash

set -e

patch_files_directory="."
if [ -n "${1}" ]; then
    patch_files_directory="${1}"
fi

while read -e patch_file; do
    echo "Applying file ${patch_file}}"
    patch --batch --forward --strip=1 --silent < "${patch_file}"
done < <(find "${patch_files_directory}" -name '*.patch')
