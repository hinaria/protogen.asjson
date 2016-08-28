# protogen.asjson: protobuf to json

Takes a [FileDescriptorSet](https://github.com/google/protobuf/blob/3d9d1a1255583bac550f7bf94f3016e8c238fa5e/src/google/protobuf/descriptor.proto) from `stdin` and outputs its equivalent JSON representation to `stdout`. Errors are logged to `stderr`.

We need this because some language libraries (namely Ruby and .NET) don't support reading proto2 profobufs, and thus cannot read `FileDescriptorSet`s which are exclusively serialized as proto2 only.