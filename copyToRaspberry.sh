declare -r BINARY_NAME="bremen_trash-arm"
declare -r RASPI_SCP_FOLDER="pi@pi4-rack-0.local:/home/pi/test"

echo "Copy arm binary '$BINARY_NAME' to $RASPI_SCP_FOLDER"
scp $BINARY_NAME $RASPI_SCP_FOLDER