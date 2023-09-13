FROM clamav/clamav

COPY --from=golang:1.18-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /c8y-scanner
COPY testfiles/ ./testfiles/
COPY rest/ ./rest/
COPY scanner/ ./scanner/
COPY go.* *.go ./
RUN go mod download
RUN go build -o c8y-scanner

EXPOSE 8888

COPY docker/entrypoint.sh /
ENTRYPOINT /entrypoint.sh
