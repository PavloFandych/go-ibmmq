1. Download mq-client: 9.2.3.0-IBM-MQC-Redist-LinuxX64
2. Export env variables:
- export MQ_INSTALLATION_PATH=/home/total/Documents/software/9.2.3.0-IBM-MQC-Redist-LinuxX64
  export CGO_CFLAGS="-I$MQ_INSTALLATION_PATH/inc"
  export CGO_LDFLAGS="-L$MQ_INSTALLATION_PATH/lib64 -Wl,-rpath,$MQ_INSTALLATION_PATH/lib64"
