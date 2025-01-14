# OpenStack

```
$ make
$ source .venv/bin/activate
$
```

```
$ openstack server create --net local-net --image cirros --flavor 1v-512M-1G testvm
```

```
$ openstack application credential create service --debug --expiration 2026-01-01T00:00:00 -f value -c secret
```
