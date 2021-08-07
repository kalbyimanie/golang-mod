FROM alpine:latest
RUN apk --no-cache update  && \
    apk --no-cache upgrade && \
    apk add mysql-client \
            bash \
            vim \
            net-tools

COPY scripts/connect.sh /bin/connect.sh

RUN chmod +x /bin/connect.sh

CMD [ "connect.sh" ]

# test 



    



