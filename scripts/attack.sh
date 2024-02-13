cd ./vegeta/
cat targets.http | vegeta attack -rate 0 -max-workers 100 -duration 10s | tee results.bin | vegeta report 