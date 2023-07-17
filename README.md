# `gcr.io/garethjevans/ytt`

The Paketo Buildpack for Ytt is a Cloud Native Buildpack that provides the Ytt binary tool.

## Behavior

This buildpack will participate all the following conditions are met

* Another buildpack requires `ytt`

The buildpack will do the following:

* Contributes Ytt to a layer marked `launch` with command on `$PATH`

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
