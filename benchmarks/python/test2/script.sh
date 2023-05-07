../ghz \
--insecure \
--skipTLS \
--proto ../experiment.proto \
--call Experiment.Test2 \
--total=200000 \
--concurrency=$1 \
--connections=$1 \
--disable-template-functions \
--disable-template-data \
-O csv \
0.0.0.0:9090 > grpcThroughputsTest2VU$1.csv