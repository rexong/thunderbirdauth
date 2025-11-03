#!/bin/bash

IMAGE_NAME="thunderbird/thunderbirdauth"
IMAGE_TAG=$(date +%Y%m%d%H%M%S)
FULL_IMAGE_NAME="$IMAGE_NAME:$IMAGE_TAG"

DOCKERFILE_PATH="./Dockerfile"
BUILD_CONTEXT="."

echo "🚀 Starting Docker image build..."
echo "Image Name: $IMAGE_NAME"
echo "Image Tag: $IMAGE_TAG"

if [ ! -f "$DOCKERFILE_PATH" ]; then
    echo "❌ Error: Dockerfile not found at $DOCKERFILE_PATH"
    exit 1
fi

docker build -t "$FULL_IMAGE_NAME" -f "$DOCKERFILE_PATH" "$BUILD_CONTEXT"

BUILD_STATUS=$?

if [ $BUILD_STATUS -eq 0 ]; then
    echo ""
    echo "✅ Success! Docker image built and tagged as: $FULL_IMAGE_NAME"
    echo "You can run it with: docker run -d $FULL_IMAGE_NAME"
    
    # Optional: Tag it as 'latest' as well
    docker tag "$FULL_IMAGE_NAME" "$IMAGE_NAME:latest"
    echo "Also tagged as: $IMAGE_NAME:latest"
else
    echo ""
    echo "❌ Build Failed. Please check the error messages above."
    exit $BUILD_STATUS
fi
