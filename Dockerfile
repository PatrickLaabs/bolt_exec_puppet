# Dockerfile
FROM alpine
COPY golang-pipeline /usr/bin/golang-pipeline
ENTRYPOINT ["/usr/bin/golang-pipeline"]