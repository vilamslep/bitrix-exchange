FROM ubuntu

COPY ./downloads /tmp/downloads

RUN apt-get update && apt-get install -y openjdk-8-jre-headless \
    && tar -xvf /tmp/downloads/zookeeper.tar.gz -C /opt \
    && mv -f /opt/apache-zookeeper-3.7.0-bin /opt/zookeeper \
    && cp /opt/zookeeper/conf/zoo_sample.cfg /opt/zookeeper/conf/zoo.cfg

ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64

EXPOSE 2181

WORKDIR /opt/zookeeper

VOLUME ["/opt/zookeeper/conf", "/tmp/zookeeper"]

ENTRYPOINT ["/opt/zookeeper/bin/zkServer.sh"]

CMD ["start-foreground"]
