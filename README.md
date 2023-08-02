
[![Build Status](https://github.com/10sr/tflint-ruleset-domainname/workflows/build/badge.svg?branch=main)](https://github.com/10sr/tflint-ruleset-domainname/actions)

tflint-ruleset-domainname
=========================

Check if domain name for DNS record is valid.

**CAUTION: I developed this plugin mainly for my personal usage.
So there should be many corner cases where this plugin does not cover,
and currently this plugin only supports `aws_route53_record` resources.
(But, contributions are of course welcome :raised_hands:)**


```hcl
# Fail
resource "aws_route53_record" "invalid" {
    type = "A"
    name = "invalid domain.example.com"
}

# Pass
resource "aws_route53_record" "valid" {
    type = "A"
    name = "valid-domain.example.com"
}
```


Requirements
------------

- TFLint v0.42+



Installation
------------

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:



```hcl
plugin "domainname" {
  enabled = true

  version = "0.0.3"
  source  = "github.com/10sr/tflint-ruleset-domainname"

  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----

  mQINBGTHqDgBEADIxKlgONJ3IREBc5P5nr+pmHBnNwanXtR2nNnUFUj4Ro3Q5og5
  G+evy7n3nShuNbgY64vO3glUPs1vOqgPllRuxRepBoDrplqOHoOFwCvNUQjp8IpM
  LjhvHvwfgX2kOkTdBkTQwf6fLs67xVsXE1pBj8tQq4j5TfOJ/+tofn6N2kokDxXD
  ...
  KEY
}
```

`signing_key` is available from [signing_key.pub](signing_key.pub).


Rules
-----

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|route53_domain_name|Check letters in route53 domain name (Supports A, AAAA, CNAME type)|ERROR|✔|[AWS Doc](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html#domain-name-format-hosted-zones)|



Building the plugin
-------------------

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "domainname" {
  enabled = true
}
EOS
$ tflint
```
