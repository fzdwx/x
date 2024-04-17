#!/usr/bin/env just --justfile

hello:
  echo "hello world"

new name:
    mkdir {{name}}
    cd {{name}} && go mod init github.com/fzdwx/x/{{name}}
