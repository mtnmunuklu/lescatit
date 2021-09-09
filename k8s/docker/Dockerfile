FROM golang

WORKDIR /app

COPY authsvc /app/authsvc
COPY crawlsvc /app/crawlsvc
COPY catsvc /app/catsvc
COPY apisvc /app/apisvc

EXPOSE 9001
EXPOSE 9003
EXPOSE 9002
EXPOSE 9000

CMD ["/app/authsvc", "/app/crawlsvc", "/app/catsvc", "/app/apisvc"]