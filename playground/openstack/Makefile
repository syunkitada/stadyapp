all:
	@make env
	@make setup

env:
	test -e .venv || python3 -m venv .venv
	.venv/bin/pip install -r opt/common/requirements.txt
	sudo docker-compose up -d

setup:
	# TODO load admintoken
	sudo modprobe openvswitch
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-db.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-rabbitmq.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-common.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-glance.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-neutron.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-placement.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-nova.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-compute.sh
	# sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-openstack-data.sh

clean:
	rm -rf .venv
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/clean-db.sh
	sudo docker-compose down
