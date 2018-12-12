FROM golang

COPY ${PWD} /go/src/app

WORKDIR /go/src/app

RUN /go/src/app/compile.sh

CMD /go/src/app/bin/frontend
