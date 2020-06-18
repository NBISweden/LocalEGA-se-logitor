ARG elasticsearch_version=7.7.1
FROM docker.elastic.co/elasticsearch/elasticsearch:${elasticsearch_version}
USER root
RUN mkdir -p /usr/share/elasticsearch/config/s3-keys
RUN chown elasticsearch:elasticsearch config/elasticsearch.yml \
    && chown elasticsearch:elasticsearch /usr/share/elasticsearch/config/s3-keys \
    && bin/elasticsearch-plugin install --batch repository-s3
COPY conf/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
USER elasticsearch
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

