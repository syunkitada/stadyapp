all:
	test -e .venv || python3 -m venv .venv
	.venv/bin/pip install -r requirements.txt
	sudo docker-compose up -d
	@make setup

setup:
	sudo modprobe openvswitch
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-db.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-memcached.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-rabbitmq.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-common.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-keystone.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-glance.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-neutron.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-placement.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-nova.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-openstack-data.sh
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/setup-compute.sh

clean:
	rm -rf .venv
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/clean-db.sh
	sudo docker-compose down
