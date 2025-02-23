#!/bin/bash

svcName=${1}
make gen-server-suyiiyii svc=${svcName}
make gen-client svc=${svcName}