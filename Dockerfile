# Dockerfile
FROM alpine
COPY bolt_exec_puppet /usr/bin/bolt_exec_puppet
ENTRYPOINT ["/usr/bin/bolt_exec_puppet"]