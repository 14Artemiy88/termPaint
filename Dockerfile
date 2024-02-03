FROM busybox
WORKDIR /build
ADD termPaint .
ADD config.yaml .
CMD ["./termPaint"]