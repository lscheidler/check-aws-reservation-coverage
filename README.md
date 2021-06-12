# check-aws-reservation-coverage

check aws reservation coverage

## Usage

```
Usage of build/linux_amd64/check-aws-reservation-coverage (0.1):
  -c float
    	set critical threshold (default 25)
  -critical float
    	set critical threshold (default 25)
  -d	set debug mode
  -ec2
    	check ec2 (default true)
  -elasticache
    	check elasticache
  -elasticsearch
    	check elasticsearch
  -rds
    	check rds
  -redshift
    	check redshift
  -w float
    	set warning threshold (default 50)
  -warning float
    	set warning threshold (default 50)
```

## Icinga2 example config

You can find an example config for icinga2 [here](doc/icinga2.conf)
