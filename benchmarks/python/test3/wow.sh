echo "vu 1" > res.txt

/usr/bin/time -ao res.txt -f "%e" ./script.sh 1
/usr/bin/time -ao res.txt -f "%e" ./script.sh 1
/usr/bin/time -ao res.txt -f "%e" ./script.sh 1
/usr/bin/time -ao res.txt -f "%e" ./script.sh 1

echo "vu 50" >> res.txt
/usr/bin/time -ao res.txt -f "%e" ./script.sh 50
/usr/bin/time -ao res.txt -f "%e" ./script.sh 50
/usr/bin/time -ao res.txt -f "%e" ./script.sh 50
/usr/bin/time -ao res.txt -f "%e" ./script.sh 50

echo "vu 100" >> res.txt
/usr/bin/time -ao res.txt -f "%e" ./script.sh 100
/usr/bin/time -ao res.txt -f "%e" ./script.sh 100
/usr/bin/time -ao res.txt -f "%e" ./script.sh 100
/usr/bin/time -ao res.txt -f "%e" ./script.sh 100

echo "vu 400" >> res.txt
/usr/bin/time -ao res.txt -f "%e" ./script.sh 400
/usr/bin/time -ao res.txt -f "%e" ./script.sh 400
/usr/bin/time -ao res.txt -f "%e" ./script.sh 400
/usr/bin/time -ao res.txt -f "%e" ./script.sh 400