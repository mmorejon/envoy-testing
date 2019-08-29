#!/bin/bash

myArray=("$@")

for i in "${myArray[@]}"
do
  if [ 'eds' == $i ];
  then
  cp /tmp-files/eds.conf /eds-tmp
  mv /eds-tmp /etc/envoy/eds.conf
  echo 'File EDS moved.'
  fi

  if [ 'cds' == $i ];
  then
  cp /tmp-files/cds.conf /cds-tmp
  mv /cds-tmp /etc/envoy/cds.conf
  echo 'File CDS moved.'
  fi

  if [ 'lds' == $i ];
  then
  cp /tmp-files/lds.conf /lds-tmp
  mv /lds-tmp /etc/envoy/lds.conf
  echo 'File LDS moved.'
  fi
done

exit 0