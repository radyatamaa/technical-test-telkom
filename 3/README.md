# technical-test-telkom

Apakah ada kesalahan dari script di bawah ini? Jika ada tolong jelaskan dimana letak
kesalahannya dan bagaimana anda memperbaikinya. Jika tidak ada, tolong jelaskan untuk apa
script di bawah ini.
FROM golang
ADD . /go/src/github.com/telkomdev/indihome/backend
WORKDIR /go/src/github.com/telkomdev/indihome
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
LISTEN 80

jawaban : 
FROM golang
RUN  apk update && apk add --no-cache git
WORKDIR /go/src/github.com/telkomdev/indihome
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
LISTEN 80

script diatas adalah script untuk dockerfile yang mana untuk membuild suatu aplikasi menjadi docker image