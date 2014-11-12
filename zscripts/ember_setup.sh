#!/bin/bash

set -e

mkdir -p public/js/utils

wget -O ~/Golang/src/github.com/adred/wiki-player/static/js/utils/jquery.js http://code.jquery.com/jquery-2.1.1.js
wget -O ~/Golang/src/github.com/adred/wiki-player/static/js/utils/handlebars.js http://builds.handlebarsjs.com.s3.amazonaws.com/handlebars-1.0.0.js
wget -O ~/Golang/src/github.com/adred/wiki-player/static/js/utils/ember.js http://builds.emberjs.com.s3.amazonaws.com/ember-latest.js
wget -O ~/Golang/src/github.com/adred/wiki-player/static/js/utils/ember-data.js http://builds.emberjs.com.s3.amazonaws.com/ember-data-latest.js
