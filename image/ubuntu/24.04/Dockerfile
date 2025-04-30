FROM scratch

ADD --link https://cloud-images.ubuntu.com/noble/current/noble-server-cloudimg-amd64.img /image/ubuntu.img

COPY --link --from=devmachines/exit0 /exit0 /exit0

CMD ["/exit0"]
