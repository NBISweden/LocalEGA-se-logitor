ARG elasticsearch_version=7.8.0
FROM docker.elastic.co/elasticsearch/elasticsearch:${elasticsearch_version}
RUN mkdir -p /usr/share/elasticsearch/config/s3-keys
RUN bin/elasticsearch-plugin install --batch repository-s3
COPY conf/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

