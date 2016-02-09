#!/bin/bash
if [ $# -eq 0 ]; then
        echo "Process a PDF document to make it look scanned."
        echo "Usage: ./$0 DOCUMENT"
        exit
fi
echo "Scanning..."
exiftool -all:all "$1"
convert -compress lzw  -rotate 0.5  -density 300x300 "$1" -colorspace gray -threshold 90% scanned.pdf
echo "Removing metadata..."
exiftool -all:all= scanned.pdf
