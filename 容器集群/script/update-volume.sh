#! /bin/bash

rm -rf temp/
mkdir temp/

node update-volume.js

ls temp/

kubectl apply  -f temp/