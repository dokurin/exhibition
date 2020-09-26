FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

ENV DOCKER_VERSION=5:19.03.12~3-0~ubuntu-focal \
    DOCKER_COMPOSE_VERSION=1.26.2 \
    GO_VERSION=1.15.2 \
    NODE_VERSION=14.11.0

# Install Docker
RUN apt-get update \
 && apt-get install -y \
      apt-transport-https \
      bash-completion \
      ca-certificates \
      curl \
      gnupg-agent \
      software-properties-common \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* \
 && curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - \
 && apt-key fingerprint 0EBFCD88 \
 && add-apt-repository \
      "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
      $(lsb_release -cs) \
      stable" \
 && apt-get update \
 && apt-get install -y \
      docker-ce-cli=${DOCKER_VERSION} \
      containerd.io \
 && apt-get clean

## Install Docker Compose
RUN curl -sL "https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
 && chmod +x /usr/local/bin/docker-compose

# Install tools(with apt-get)
RUN apt-get update \
 && apt-get install -y \
      build-essential \
      gettext \
      git \
      hub \
      jq \
      python3-pip \
      tree \
      vim \
      wget \
 && apt-get clean

# Install Golang
RUN wget -q https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
 && echo 'export PATH=${PATH}:/usr/local/go/bin:/root/go/bin' >> .bashrc \
 && rm -f go${GO_VERSION}.linux-amd64.tar.gz

# Install Nodejs
WORKDIR /tmp/nodejs
RUN curl -L https://raw.githubusercontent.com/tj/n/master/bin/n -o n \
 && bash n ${NODE_VERSION}

# Setup prompt
ADD https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh /root/.git-prompt.sh
COPY bashrc /root/.bashrc

WORKDIR /root
