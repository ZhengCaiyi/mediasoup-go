#!/bin/sh
cur_dir=$(cd "$(dirname "$0")"; pwd)
chmod +x ${cur_dir}/flatc

 ${cur_dir}/flatc --gen-object-api --go-module-name github.com/jiyeyuran/mediasoup-go -g -o ${cur_dir}/../ ${cur_dir}/proto/*.fbs