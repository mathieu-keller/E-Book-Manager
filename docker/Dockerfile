FROM alpine:3.15.4
LABEL maintainer=MathieuKeller@gmx.de
WORKDIR /app
COPY backend .
COPY frontend ./dist
RUN chmod 0777 backend
EXPOSE 8080
ENTRYPOINT ["/app/backend"]
