FROM ubuntu
WORKDIR /build
ADD termPaint .
ADD config.yaml .
CMD ["./termPaint"]