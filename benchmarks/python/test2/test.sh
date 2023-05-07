../ghz \
--insecure \
--skipTLS \
--proto ../experiment.proto \
--call Experiment.Test2 \
--duration=$1s \
--concurrency=1 \
--connections=1 \
--disable-template-functions \
--disable-template-data \
0.0.0.0:9090