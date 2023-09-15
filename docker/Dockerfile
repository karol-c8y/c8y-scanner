FROM golang:1.19-alpine
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /c8y-scanner
COPY rest/ ./rest/
COPY scanner/ ./scanner/
COPY cumulocity/ ./cumulocity/
COPY go.* *.go ./
RUN go mod download
RUN go build -o c8y-scanner


FROM clamav/clamav:latest_base

RUN mkdir /c8y-scanner
COPY --from=0 /c8y-scanner/c8y-scanner /c8y-scanner/c8y-scanner

EXPOSE 80

COPY docker/entrypoint.sh /
ENTRYPOINT /entrypoint.sh
