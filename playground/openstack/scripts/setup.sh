#!/bin/bash -xe

IAM_TOKEN=$(openstack application credential create service --expiration 2026-01-01T00:00:00 -f value -c secret)

sudo modprobe openvswitch

sudo docker exec openstack-allinone /opt/openstack/scripts/setup-db.sh
sudo docker exec openstack-allinone /opt/openstack/scripts/setup-rabbitmq.sh
sudo docker exec openstack-allinone /opt/openstack/scripts/setup-common.sh

sudo docker exec -e IAM_TOKEN="${IAM_TOKEN}" openstack-allinone /opt/openstack/scripts/setup-glance.sh
sudo docker exec -e IAM_TOKEN="${IAM_TOKEN}" openstack-allinone /opt/openstack/scripts/setup-neutron.sh
sudo docker exec -e IAM_TOKEN="${IAM_TOKEN}" openstack-allinone /opt/openstack/scripts/setup-placement.sh
sudo docker exec -e IAM_TOKEN="${IAM_TOKEN}" openstack-allinone /opt/openstack/scripts/setup-nova.sh
sudo docker exec -e IAM_TOKEN="${IAM_TOKEN}" openstack-allinone /opt/openstack/scripts/setup-compute.sh
