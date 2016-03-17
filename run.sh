#!/bin/bash
cd word_filter && go build
mv word_filter ../bin
cd ..
bin/word_filter