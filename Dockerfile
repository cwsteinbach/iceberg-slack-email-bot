# Copied from https://joe-bot.net/recipes/docker/

# We use busybox as base image to create very small Docker images.
# Before you use this, you should check if there is a newer version you want to use.
FROM alpine

# Add config file
ADD build/pinot-bot /bin/pinot-bot

# Compile
ENTRYPOINT ["pinot-bot"]