#!/bin/sh -l

BASE=$1
ACCESS_TOKEN=$2
APP_ID=$3

FLAG_BASE=""
if [ ! -z $BASE ]; then
    FLAG_BASE="-base $BASE"
fi

FLAG_ACCESS_TOKEN=""
if [ ! -z $ACCESS_TOKEN ]; then
    FLAG_ACCESS_TOKEN="-access $ACCESS_TOKEN"
fi

FLAG_APP_ID=""
if [ ! -z $APP_ID ]; then
    FLAG_APP_ID="-app $APP_ID"
fi

/app $FLAG_BASE $FLAG_ACCESS_TOKEN $FLAG_APP_ID
