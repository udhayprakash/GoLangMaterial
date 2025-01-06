#!/bin/bash

# Directory containing log files
LOG_DIR="./logs"

# Maximum number of log files to retain
MAX_LOG_FILES=3

# Delete old log files
ls -t $LOG_DIR/* 2>/dev/null | tail -n +$(($MAX_LOG_FILES + 1)) | xargs rm -f 2>/dev/null
