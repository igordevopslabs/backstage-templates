FROM golang:${{ values.goVersion | lower }} as builder

#avoid root
ENV USER=appuser 
ENV UID=10001

#avoid root
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /app
COPY /src/ /app

RUN CGO_ENABLED=1 GO111MODULES=on go build -ldflags="-s -w" .

FROM golang:${{ values.goVersion | lower }}

#avoid rootless
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /app
COPY --from=builder /app/${{ values.apiName | lower }} /app

#avoid rootless
USER appuser:appuser

ENTRYPOINT ["/app/${{ values.apiName | lower }}"]