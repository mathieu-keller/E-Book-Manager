FROM alpine:3.15.4
LABEL maintainer=MathieuKeller@gmx.de
WORKDIR /app
COPY backend/e-book-manager .
COPY frontend ./bundles
COPY wait-for-db.sh .
RUN chmod 0777 e-book-manager
EXPOSE 8080
ENTRYPOINT ["/bin/sh","wait-for-db.sh", "--", "/app/e-book-manager"]
