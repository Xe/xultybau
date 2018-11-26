#!/bin/sh

rm parsed-en.js
rm en-jbo.json

wget https://raw.githubusercontent.com/La-Lojban/sutysisku/master/data/parsed-en.js
node ./sutysisku-json.js
