FROM debian:stretch

RUN apt-get update && apt-get -y install fortune-mod fortunes-bofh-excuses fortunes-mario fortunes-off  && apt-get clean

COPY hello-http /

ENTRYPOINT "/hello-http"
