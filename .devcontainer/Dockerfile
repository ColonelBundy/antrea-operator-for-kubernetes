FROM mcr.microsoft.com/devcontainers/go:1.23

# Set environment variable to allow pip to break system packages
ENV PIP_BREAK_SYSTEM_PACKAGES=1

# Define build argument for Operator SDK version
ARG OPERATOR_SDK_VERSION=v1.39.1

# Update package list and install dependencies
RUN apt-get update && apt-get install -y \
    unzip \
    curl \
    python3 \
    python3-pip \
    && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /tmp

# Install kubectl with architecture detection
RUN export ARCH=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n $(uname -m) ;; esac) && \
    curl -sSL -o /usr/local/bin/kubectl "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/${ARCH}/kubectl" && \
    chmod +x /usr/local/bin/kubectl

# Install Operator SDK with architecture detection
RUN export ARCH=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n $(uname -m) ;; esac) && \
    export OS=$(uname | awk '{print tolower($0)}') && \
    curl -LO "https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_SDK_VERSION}/operator-sdk_${OS}_${ARCH}" && \
    chmod +x "operator-sdk_${OS}_${ARCH}" && \
    mv "operator-sdk_${OS}_${ARCH}" /usr/local/bin/operator-sdk

USER vscode

# Set PATH for go
ENV PATH="/go/bin:$PATH"

WORKDIR /workspace