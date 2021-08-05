FROM golang

WORKDIR /app

COPY authsvc /app/authsvc
COPY catsvc /app/catsvc
COPY apisvc /app/apisvc

EXPOSE 9000
EXPOSE 9001
EXPOSE 9002

CMD ["/app/authsvc", "/app/catsvc", "/app/apisvc"]