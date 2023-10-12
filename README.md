## Setup Serverless Framework

First install [AWS-CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html):
```bash
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

**You have to configure your AWS credentials into `~/.aws/credentials` file**

Second, if you don't have the node in your machine, [install it](https://www.freecodecamp.org/news/node-version-manager-nvm-install-guide/):
```
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
source ~/.zshrc
```

Now, install the serverless cli:
```bash
npm install -g serverless serverless-offline
```

## Setup the project

```
git clone https://github.com/tenlisboa/go-passwordless-otp-auth.git
cd go-passwordless-otp-auth
go mod tidy
```

## Deploying the project

```
make deploy-dev
```