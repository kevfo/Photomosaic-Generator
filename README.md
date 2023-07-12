# Photomosaic Generator

This is a minimalistic CLI application written using Go (i.e., Golang) that - given a PNG image and a directory containing PNG photos - attempts to build a photomosaoic of the said PNG image.

## How to Use 

Once you have cloned the repository onto your machine, you will need to build the application's entry point `build.go` using the command `go build build.go`.  Doing so will generate a binary executable `build.exe` that you can then run with various flags (i.e., options):

1.  `img`
    This flag specifies the image that you want `build.exe` to construct a photomosaic of.  The default value of `img` is `"image.png"`.

1.  `out`
    This flag specifies the name of the photomosaic that you want `build.exe` to name.  The default value of `img` is `"result.png"`.

1.  `lib`
    This flag specifies the name of the photo library that you want to use to generate the photomosaic.  The default value of `lib` is `"pictures"`.

1.  `tileLength` and `tileHeight`
    These two flags represent the length and the height of each tile in the photomosaic.  The default value of both flags are `20` and `40` respectively.

## Limitations

Because the application is simple, there are a few limitaions that may influence results:

1.  The photos that you use in your photo library may influence the final result.