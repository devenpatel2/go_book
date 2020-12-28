#!/bin/bash
run_cmd="go run $1"
build_cmd="go build $1"
echo $run_cmd
eval $run_cmd
echo $build_cmd
eval $build_cmd

