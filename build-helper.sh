#!/usr/bin/env bash

# clean cache
go clean -cache

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}
	
platforms=("windows/amd64" "windows/386" "linux/amd64" "linux/386")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	out_dir=jkt48show'-'$GOOS'-'$GOARCH
	output_name=jkt48sr
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o ./releases/$out_dir/$output_name $package.go
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done

# Get the directory where the script is located
cd releases

# Zip each subdirectory separately
for dir in */; do
    zip -r "${dir%/}.zip" "$dir"
done
