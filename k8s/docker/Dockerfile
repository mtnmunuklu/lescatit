FROM golang

WORKDIR /app

COPY authsvc /app/authsvc
COPY catsvc /app/catsvc
COPY crawlsvc /app/crawlsvc
COPY apisvc /app/apisvc

EXPOSE 9001
EXPOSE 9002
EXPOSE 9003
EXPOSE 9000

CMD ["/app/authsvc", "/app/catsvc", "/app/crawlsvc", "/app/apisvc"]