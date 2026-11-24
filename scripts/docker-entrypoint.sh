#!/usr/bin/env bash

# Copyright (c) 2026 Huawei Technologies Co., Ltd.
# opengauss-prometheus-exporter is licensed under Mulan PSL v2.
# You can use this software according to the terms and conditions of the Mulan PSL v2.
# You may obtain a copy of Mulan PSL v2 at:
#          http://license.coscl.org.cn/MulanPSL2
# THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
# EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
# MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
# See the Mulan PSL v2 for more details.

set -Eeo pipefail 



# check to see if this file is being run or sourced from another script
_is_sourced() {
	[ "${#FUNCNAME[@]}" -ge 2 ] \
		&& [ "${FUNCNAME[0]}" = '_is_sourced' ] \
		&& [ "${FUNCNAME[1]}" = 'source' ]
}

_mtk_want_help() {
	local arg
	for arg; do
		case "$arg" in
			--help|-h|-v|--version)
				return 0
				;;
		esac
	done
	return 1
}

_main() {
	# if first arg looks like a flag, assume we want to run openGauss server
	if [ "${1:0:1}" = '-' ]; then
		set -- opengauss_exporter "$@"
	fi

	exec "$@"
}

if ! _is_sourced; then
	_main "$@"
fi
