#!/bin/bash
cd word_filter && go build
mv word_filter ../bin
cd ..
nohup bin/word_filter > log/access.log 2>&1 &
