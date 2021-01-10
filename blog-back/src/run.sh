#!/bin/bash

curDir=`dirname $0`
cd $curDir/../
prjHome=`pwd`

if [ ! -d $prjHome/logs ]
then
    mkdir -p $prjHome/logs
fi

if [ ! -d $prjHome/tmp ]
then
    mkdir -p $prjHome/tmp
fi

if [ ! -f $prjHome/tmp/api.pid ]
then
    touch $prjHome/tmp/api.pid
fi

cd $prjHome/src/main

go run main.go --prj-home=$prjHome