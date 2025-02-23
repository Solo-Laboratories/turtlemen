// package: mockauth
// file: mock-user-pass.proto

import * as jspb from "google-protobuf";

export class AuthRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuthRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AuthRequest): AuthRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AuthRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuthRequest;
  static deserializeBinaryFromReader(message: AuthRequest, reader: jspb.BinaryReader): AuthRequest;
}

export namespace AuthRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class RandomStringResponse extends jspb.Message {
  getRandomString(): string;
  setRandomString(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RandomStringResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RandomStringResponse): RandomStringResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RandomStringResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RandomStringResponse;
  static deserializeBinaryFromReader(message: RandomStringResponse, reader: jspb.BinaryReader): RandomStringResponse;
}

export namespace RandomStringResponse {
  export type AsObject = {
    randomString: string,
  }
}

