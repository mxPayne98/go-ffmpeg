# Video Compressor using FFmpeg and Go

This script allows you to compress all the `.mp4` videos under a given directory using `ffmpeg`. The compressed videos will be saved to an output directory. The script compresses videos concurrently to speed up the process.

## Prerequisites

1. Ensure you have Go installed. You can download and install Go from [the official website](https://golang.org/).
2. Ensure `ffmpeg` is installed and available in your `PATH`. If you haven't installed it yet, refer to the [FFmpeg download page](https://ffmpeg.org/download.html) for instructions.

## Installation

1. Clone this repository or download the Go script.
2. Navigate to the directory containing the script.

## Usage

To compress videos:

```bash
go run compress.go /path/to/input/directory /path/to/output/directory
```

Replace `/path/to/input/directory` with the path of the directory containing the `.mp4` videos you want to compress and `/path/to/output/directory` with the path where you want the compressed videos to be saved.

## Configuration

The script by default uses 4 concurrent workers (goroutines) to compress videos. If you wish to modify the number of workers, change the `concurrency` variable value in the script.

## Details

The script compresses videos using the following `ffmpeg` command:

```
ffmpeg -i video.mp4 -vf "scale=1280:-1" -c:v libx264 -preset veryslow -crf 28 -c:a aac -b:a 128k video_cmp.mp4
```

This command scales the video to a width of 1280 pixels (keeping the aspect ratio intact), uses the `libx264` codec with the `veryslow` preset and a crf (constant rate factor) value of 28, and compresses the audio using the `aac` codec with a bitrate of 128k.