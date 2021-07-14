FROM alpine:edge
COPY ./main /opt/bin/main
CMD [ "/bin/sh", "-c", "/opt/bin/main" ]