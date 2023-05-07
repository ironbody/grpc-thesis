../ghz \
--insecure \
--skipTLS \
--proto ../experiment.proto \
--call Experiment.Test1 \
--total=500000 \
--concurrency=1 \
--connections=1 \
--disable-template-functions \
--disable-template-data \
-O csv \
0.0.0.0:9090 > grpcThroughputsTest1VU1.csv