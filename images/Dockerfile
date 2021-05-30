FROM ubuntu:latest
MAINTAINER steve.g.szabo@gmail.com

ENV WEBAPP_VERSION=1.0.5

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y curl python3 python3-pip gunicorn && \
    apt-get clean && \
    pip3 install flask

COPY root/app /app
RUN chown -R www-data:www-data /app
EXPOSE 8080
USER www-data
ENTRYPOINT ["/app/entrypoint.sh"]
