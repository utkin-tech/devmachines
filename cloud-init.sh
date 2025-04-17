sudo apt install cloud-image-utils genisoimage

genisoimage -output seed.iso -volid cidata -joliet -rock user-data meta-data

# genisoimage -output seed.iso -volid cidata -joliet -rock user-data meta-data network-config
