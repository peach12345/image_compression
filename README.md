# ImageCompressor

A Go-based image compression tool that efficiently reduces image file sizes while maintaining quality.

## Overview
This project provides a simple yet powerful solution for image compression, built in Go. It handles various image formats and offers customizable compression settings.

## Features
- Supports multiple image formats (JPEG, PNG, WebP)
- Batch processing capabilities
- Customizable compression quality
- Maintains EXIF data (optional)
- Command-line interface

## Requirements
- Go 1.16 or higher
- ImageMagick (optional for advanced features)

## Installation
```bash
go get github.com/yourusername/imagecompressor
```

## Usage
Basic compression:
```bash
imagecompressor -input image.jpg -output compressed.jpg
```

## Configuration
Configure compression settings in `config.yaml`:
```yaml
quality: 85
preserve_metadata: true
output_format: "jpg"
```

## Contributing
1. Fork the repository
2. Create feature branch
3. Submit pull request

## License
MIT License
