#!/bin/bash

set -e

mkdir -p public/js/lib

wget -O ~/Golang/src/github.com/adred/wiki-player/static/js/lib/jquery.js http://code.jquery.com/jquery-2.1.1.js
wget -O ~/Golang/src/github.com/adred/wiki-player/static/js/lib/handlebars.js http://builds.handlebarsjs.com.s3.amazonaws.com/handlebars-1.0.0.js
wget -O ~/Golang/src/github.com/adred/wiki-player/static/js/lib/ember.js http://builds.emberjs.com.s3.amazonaws.com/ember-latest.js
wget -O ~/Golang/src/github.com/adred/wiki-player/static/js/lib/ember-data.js http://builds.emberjs.com.s3.amazonaws.com/ember-data-latest.js
