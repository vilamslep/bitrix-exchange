FROM ubuntu

COPY ./downloads /tmp/downloads

RUN apt-get update && apt-get install -y openjdk-8-jre-headless net-tools telnet iputils-ping \
    && tar -xvf /tmp/downloads/kafka.tgz -C /opt \
    && mv -f /opt/kafka_2.13-3.0.0 /opt/kafka

ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64

EXPOSE 9092

WORKDIR /opt/kafka

VOLUME ["/opt/kafka/config", "/tmp/kafka-logs"]

ENTRYPOINT ["/opt/kafka/bin/kafka-server-start.sh"]

CMD [ "/opt/kafka/config/server.properties"]
