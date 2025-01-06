#!/bin/bash -ex

PUBLIC_KEY_DIR=/etc/compute/token_keys/public
PRIVATE_KEY_DIR=/etc/compute/token_keys/private
MAX_FILES=4

sudo mkdir -p "${PUBLIC_KEY_DIR}"
sudo mkdir -p "${PRIVATE_KEY_DIR}"

key_name=$(date "+%Y%m%d%H%M%S").pem
# sudo openssl genrsa -out "${PRIVATE_KEY_DIR}/${key_name}.pem" 4096
# sudo openssl rsa -in "${PRIVATE_KEY_DIR}/${key_name}.pem" -pubout -out "${PUBLIC_KEY_DIR}/${key_name}.pem"
#

sudo openssl genpkey -algorithm ed25519 -out "${PRIVATE_KEY_DIR}/${key_name}"
sudo openssl pkey -in "${PRIVATE_KEY_DIR}/${key_name}" -pubout -out "${PUBLIC_KEY_DIR}/${key_name}"

dirs=($PRIVATE_KEY_DIR $PUBLIC_KEY_DIR)
for dir in ${dirs[@]}; do
	sudo rm -rf $(ls -t $dir/* | tail -n+${MAX_FILES})
	ls $dir
	sudo chown -R "$USER:$USER" $dir
done
