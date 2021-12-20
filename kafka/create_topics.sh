#!/bin/bash -e

cd /opt/kafka/bin

./kafka-topics.sh --bootstrap-server kafka:9092 --create --topic changes --partitions 1 --replication-factor 1
./kafka-topics.sh --bootstrap-server kafka:9092 --create --topic tocrm --partitions 1 --replication-factor 1