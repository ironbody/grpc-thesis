ghz \
--insecure \
--skipTLS \
--proto ../experiment.proto \
--call Experiment.Test1 \
--total=500000 \
--concurrency=400 \
--connections=400 \
--disable-template-functions \
--disable-template-data \
-O csv \
--cpus=4 \
0.0.0.0:9091 > grpcThroughputsTest1VU400.csv