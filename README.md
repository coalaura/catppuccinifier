# Catppuccinifier

A blazingly fast tool to recolor any image (PNG, JPEG, WEBP, GIF) using the [Catppuccin](https://github.com/catppuccin/catppuccin) palette. It uses CIELAB perceptual color matching to ensure the output looks natural to the human eye.

## Features

- **Perceptual Matching**: Uses Lab color space (Delta-E) for high-quality results.
- **High Performance**: Multi-threaded processing with row-based concurrency and result caching.
- **Transparency Support**: Correct handling of alpha channels for icons and transparent images.
- **Format Support**: Handles PNG, JPEG, WEBP, and GIF inputs.

## Installation

Download the latest binary for your platform from the [Releases](https://github.com/coalaura/catppuccinifier/releases) page.

## Usage

Recolor an image and save it as `input_catppuccin.png`:
```bash
catppuccin input.png
```

Specify a custom output path:
```bash
catppuccin input.png output.png
```

## How it works

1. **Decoding**: Reads the input image (png/jpeg/gif/webp) and converts it into a standardized color format.
2. **Matching**: For every unique color, it calculates the "distance" to every color in the Catppuccin palette using the CIELAB space.
3. **Caching**: Results are cached in memory to avoid redundant calculations for identical pixels.
4. **Encoding**: Saves the finalized result as a PNG.
