#
#
#
FROM scratch

LABEL maintainer="darryl.west" version="18.01.20"

ADD list-service /

# where the data lives
VOLUME /data

# rest/web
EXPOSE 80

CMD [ "/list-service", "--port", "80", "--db-filename", "/data/list-service.db" ]

