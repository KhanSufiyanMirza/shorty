cd ./vegeta/

cat results.bin | vegeta plot > plot.html

open -a "Google Chrome" plot.html