# Installing libplctag

## Linux

Follow instructions on core lib repo:
https://github.com/libplctag/libplctag/blob/release/BUILD.md

Example for a system wide install on debian:

```
git clone https://github.com/libplctag/libplctag --depth 1 --branch v2.5.1
cd libplctag
mkdir build
cd build
cmake .. -DCMAKE_BUILD_TYPE=Release
make
sudo make install
```

## Windows libplctag download

For Windows you can download release files or compile if you'd like to follow instructions here:

- https://github.com/libplctag/libplctag/releases
- https://github.com/libplctag/libplctag/blob/release/BUILD.md

## Windows Prequisites

There are probably a few ways to work with golang and cgo on windows. Documentation is scarce online for windows. The following is what worked for me. If you find other instructions or steps that work better, please submit PR.

- chocolatey https://chocolatey.org/install (run pwsh in administrator mode)
- mingw64
- pkgconfiglite

Once chocolatey is installed (run pwsh in administrator mode):

```
choco install mingw pkgconfiglite
```

## Create libplctag lib dir

The packageconfig pc file on releases uses `Program Files` as the library install directory, the compiler doesn't like spaces in paths you can try adding single quotes. My preferred style is to install c/c++ libs on paths without spaces. i.e. C:\libs\x64 or D:\libs\x86 adjust paths accordingly in the libplctag.pc which comes in the zipped release.

Under your preferred libs directory create the following hierarchy:

- libplctag
  - bin
  - lib
  - include
  - pc

Then from the unzipped release:

- copy libplctag.h to include
- copy libplctag.pc to pc
- copy plctag.lib and plctag_static.lib to lib
- copy plctag.dll into bin

## Copy files to mingw

Go to the root directory of mingw64:

- C:\ProgramData\chocolatey\lib\mingw\tools\install\mingw64
- copy libplctag.h to include
- copy plctag.lib and plctag_static.lib to lib
- copy plctag.dll into bin
