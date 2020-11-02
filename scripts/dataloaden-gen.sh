#!/bin/bash
read -p "Model name (User, Episode, etc): " MODEL_NAME
read -p "Primary key type (string, int64, etc): " PRIMARY_KEY_TYPE
DEFAULT_MODEL_PATH="internal/database/entities"
read -p "Path to package the model's struct is in ($DEFAULT_MODEL_PATH): " MODEL_PATH

if [ "$MODEL_PATH" == "" ]; then
  MODEL_PATH="$DEFAULT_MODEL_PATH"
fi
LOADER_NAME="${MODEL_NAME}Loader"
COMMAND="go run github.com/vektah/dataloaden $LOADER_NAME $PRIMARY_KEY_TYPE $MODEL_PATH.$MODEL_NAME"

echo ""
echo "$COMMAND"
read -p "Press Enter to continue"
echo ""

pushd internal/dataloaders
eval "$COMMAND"
popd

echo ""
