sudo apt install genisoimage

genisoimage -output blobs/seed.iso -volid cidata -joliet -rock \
  ci-config/user-data \
  ci-config/meta-data \
  ci-config/network-config
