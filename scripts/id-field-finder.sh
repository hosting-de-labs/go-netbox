#!/usr/bin/env bash

declare swagger_file="swagger.json"
if [ -n "${1}" ]; then
    swagger_file="${1}"
fi

declare mode="get"
if [ -n "${2}" ]; then
    mode="${2}"
fi

echo "Reading file ${swagger_file} in ${mode} mode"
echo

while read -e path_json; do
    declare operation_id=$(echo "${path_json}" | jq --raw-output ".${mode}.operationId")
    while read -e field_json; do
        declare name=$(echo "${field_json}" | jq --raw-output '.name')

        if [ "${name}" = "id" ] || echo "${name}" | grep '_id' | grep -q -v '_id__'; then
            declare type=$(echo "${field_json}" | jq --raw-output '.type')
            if [ "${name}" = "id" ] && [ "${type}" != "integer" ]; then
                echo "## ${operation_id}: field ${name} has type ${type}"
            fi

            if echo "${name}" | grep '_id' | grep -q -v '_id__'; then
                if [ "${type}" != "integer" ]; then
                    echo "## ${operation_id}: field ${name} has type ${type}"
                fi
            fi
        fi

    done < <(echo "${path_json}" | jq --compact-output ".${mode}.parameters[]")
done < <(cat "${swagger_file}" | jq --compact-output '.paths[]')
