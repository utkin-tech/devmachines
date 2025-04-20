FROM scratch

WORKDIR /image

RUN wget -O ubuntu.img \
    https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64-disk-kvm.img 
