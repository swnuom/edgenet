#!/bin/bash

# Define the file containing image names (one per line)
IMAGE_LIST_FILE="image_names"

# Define the container registry where you want to push the images
CONTAINER_REGISTRY="swnuom"

# Function to build and push a Docker image
build_and_push_image() {
  local image_name="$1"
  local image_tag="$2"
  local full_image_name="$CONTAINER_REGISTRY/$image_name:$image_tag"
  
  # Build the Docker image
  docker build -t $full_image_name -f ./images/$image_name/Dockerfile ../
  # Push the Docker image to the registry
  docker push $full_image_name
  
  # Clean up: Remove the locally built image
  docker image rm $full_image_name
}

# Check if the image list file exists
if [ ! -f "$IMAGE_LIST_FILE" ]; then
  echo "Image list file $IMAGE_LIST_FILE not found."
  exit 1
fi

# Iterate through each line in the image list file
while IFS= read -r line; do
  # Split the line into image name and tag (assuming format: "image-name:tag")
  image_name=$(echo "$line" | cut -d ':' -f 1)
  image_tag=$(echo "$line" | cut -d ':' -f 2)
  
  # Call the build_and_push_image function with image name and tag
  build_and_push_image "$image_name" "$image_tag"
done < "$IMAGE_LIST_FILE"

echo "Image build and push completed."

