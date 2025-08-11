#!/bin/bash

# Script untuk generate boilerplate code template
echo "Generating boilerplate code templates..."

find . -name "*.go" -type f | while read file; do
    if [ ! -s "$file" ]; then  # If file is empty
        cat > "$file" << 'GOTEMPLATE'
package main

import (
	"fmt"
)

// TODO: Implement the functionality for this concept
// Refer to the README.md in this directory for what to learn

func main() {
	fmt.Println("Learning:", "REPLACE_WITH_CONCEPT_NAME")
	
	// TODO: Add your implementation here
}

// Example function - replace with actual implementation
func exampleFunction() {
	// Implementation goes here
}
GOTEMPLATE
    fi
done

echo "Boilerplate code generated!"
