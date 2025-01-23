#!/bin/bash

svcName=${1}
make gen-server svc=${svcName}
make gen-client svc=${svcName}