# credenv

> Add your cloud credentials to your environment in a single command

## Installation

```sh
go get -u github.com/charliekenney23/credenv
```

## Usage

```sh
credenv SOURCE
```

### View your credential environment variables for a given source

```sh
credenv aws
# sources credentials from ~/.aws/credentials and outputs:
# export AWS_ACCESS_KEY_ID='XXXXXXXXXXXXXXXX'
# export AWS_SECRET_ACCESS_KEY='XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX'
```

### Set your credential environment variables for a given source

```sh
eval `credenv aws`
```

## Supported Sources

- [AWS](https://aws.com)
- [Pulumi](https://pulumi.io)

<br>

---

&copy; 2019 [Charles Kenney](https://github.com/charliekenney23)
