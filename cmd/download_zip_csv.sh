#!/bin/sh
wget http://zipcloud.ibsnet.co.jp/zipcodedata/download?di=1506672625545 -O temp.zip;
unzip -o temp.zip
iconv -f sjis -t utf8 KEN_ALL.CSV > KEN_ALL.utf8.CSV
rm temp.zip KEN_ALL.CSV
