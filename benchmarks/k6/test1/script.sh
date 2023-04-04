../ghz \
--insecure \
--skipTLS \
--proto ../experiment.proto \
--call Experiment.Test1 \
--total=0 \
--duration=10s \
--concurrency=1000 \
--rps=500 \
--connections=10 \
--disable-template-functions \
-O csv \
--disable-template-data \
0.0.0.0:9090 >> output.csv