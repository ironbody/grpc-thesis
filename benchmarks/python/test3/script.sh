../ghz \
--insecure \
--skipTLS \
--proto ../experiment.proto \
--call Experiment.Test3 \
--total=20000 \
--concurrency=$1 \
--connections=$1 \
--disable-template-functions \
--disable-template-data \
-O csv \
0.0.0.0:9090 > grpcThroughputsTest2VU$1.csv