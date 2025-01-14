#!/bin/bash -xe

source /opt/openstack/envrc

SERVICE_NAME="placement"
VENV_DIR=${VENV_DIR:-"/opt/${SERVICE_NAME}"}

useradd "${SERVICE_NAME}" || echo "ignore because the user already exists"

python3 -m venv --system-site-packages "${VENV_DIR}"
"${VENV_DIR}/bin/pip" install -U pip

mkdir -p "$VENV_DIR/src"
mkdir -p "/etc/${SERVICE_NAME}"
mkdir -p "/var/log/${SERVICE_NAME}"

cd "$VENV_DIR/src"
if test -e placement; then
	cd placement
	git pull
else
	git clone -b "${OS_VERSION}" https://github.com/openstack/placement.git
	cd placement
fi

"${VENV_DIR}/bin/pip" install . \
	-c "/opt/oscommon/src/requirements/upper-constraints.txt" \
	-r requirements.txt \
	-r /opt/openstack/placement/requirements.txt

cp /opt/openstack/placement/auth.py /opt/placement/lib/python3.9/site-packages/placement/auth.py
cp /opt/openstack/placement/placement.conf /etc/placement/placement.conf

/opt/placement/bin/placement-manage db sync

systemctl reset-failed placement-api || echo 'ignored'
systemctl status placement-api || systemd-run \
	--unit placement-api -- \
	/opt/placement/bin/placement-api --host 127.0.0.1 --port 8778
systemctl restart placement-api
