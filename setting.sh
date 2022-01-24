#!/bin/bash

cur_path=$(pwd)
path=$1

config="$cur_path/$path"
if [ -e $config ] 
then
    read -p "Type a generator host value: " generator_host
    read -p "Type a portal address: " portal_address
    read -p "Type a portal rest key: " portal_rest_key
else
    echo "$config isn't exists"
fi