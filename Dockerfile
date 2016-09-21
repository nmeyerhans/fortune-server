FROM debian

RUN apt-get update && apt-get -y install fortune-mod && apt-get clean

COPY hello-http /

ENTRYPOINT "/hello-http"
