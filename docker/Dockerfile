FROM golang:1.13-alpine


ENV GOSUMDB=off

RUN mkdir /app
RUN mkdir /app/warungpintar

# Install curl 
RUN apk --no-cache add bash \
	curl \
	git \
	gcc \
	g++ \
	inotify-tools

WORKDIR /app/warungpintar
COPY . /app/warungpintar

RUN mkdir -p /go/_tools

RUN ls -l

ADD docker/config/build.sh /go/_tools
ADD docker/config/reflex.conf /go/_tools

RUN chmod +x /go/_tools/build.sh

# Add file watcher and swagger generator
RUN go get -v github.com/cespare/reflex

CMD ["reflex","-c","/go/_tools/reflex.conf"]
