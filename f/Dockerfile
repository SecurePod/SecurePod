# syntax=docker/dockerfile:1

ARG NODE_VERSION=20.10.0

FROM node:${NODE_VERSION}-alpine

# Use production node environment by default.
ENV NODE_ENV production

WORKDIR /usr/src/app

# Install pnpm
RUN npm install -g pnpm

# Download dependencies as a separate step to take advantage of Docker's caching.
# Leverage a cache mount to /root/.pnpm-store to speed up subsequent builds.
# Leverage a bind mounts to pnpm-lock.yaml to avoid having to copy it into this layer.
RUN --mount=type=bind,source=pnpm-lock.yaml,target=pnpm-lock.yaml \
    --mount=type=bind,source=package.json,target=package.json \
    --mount=type=cache,target=/root/.pnpm-store \
    pnpm install --frozen-lockfile

# Copy the rest of the source files into the image.
COPY . .

RUN pnpm run build

# Expose the port that the application listens on.
EXPOSE 3000

# Run the application.
CMD ["pnpm", "start"]
