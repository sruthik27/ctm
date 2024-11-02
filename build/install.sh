#!/bin/bash

# Installation script for macOS/Linux
set -e

# Check if `ctm` binary exists
if [[ ! -f "ctm" ]]; then
  echo "Error: 'ctm' binary not found. Make sure to build the project first."
  exit 1
fi

# Move binary to /usr/local/bin for global access
sudo cp ctm /usr/local/bin/
sudo chmod +x /usr/local/bin/ctm

echo "Installation complete. You can now use 'ctm' from anywhere in the terminal."

#To install the CLI, run the following command:
#$ ./build/install.sh

#To test the CLI, run the following command:
#$ ctm -l

#To add auto-complete of ctm commands append the content in ctm-completion-unix.sh to your .~/zshrc or ~/.bashrc file
