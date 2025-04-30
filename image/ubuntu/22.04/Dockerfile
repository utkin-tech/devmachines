FROM scratch

ADD --link https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64-disk-kvm.img /image/ubuntu.img

COPY --link --from=devmachines/exit0 /exit0 /exit0

CMD ["/exit0"]
