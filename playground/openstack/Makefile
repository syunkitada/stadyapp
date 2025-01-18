all:
	@make env
	@make setup

env:
	test -e .venv || python3 -m venv .venv
	.venv/bin/pip install -r opt/common/requirements.txt
	sudo docker-compose up -d

setup:
	./scripts/setup.sh
	./scripts/setup-openstack-data.sh

clean:
	rm -rf .venv
	sudo docker exec -i openstack-allinone /opt/openstack/scripts/clean-db.sh
	sudo docker-compose down
