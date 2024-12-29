#!/bin/bash

# See https://gdal.org/en/stable/download.html:
#
# - brew install gdal

TMPDIR=/tmp/ne_110m
ZIPFILE=ne_110m_admin_0_countries.zip
mkdir -p "$TMPDIR"
curl -o "$TMPDIR/$ZIPFILE" https://www.naturalearthdata.com/download/110m/cultural/ne_110m_admin_0_countries.zip
unzip $ZIPFILE -d $TMPDIR
