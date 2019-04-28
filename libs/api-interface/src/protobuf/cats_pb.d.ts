// package: cats
// file: libs/api-interface/src/protobuf/cats.proto

import * as jspb from "google-protobuf";

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export class Cat extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getTitle(): string;
  setTitle(value: string): void;

  getAuthor(): string;
  setAuthor(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Cat.AsObject;
  static toObject(includeInstance: boolean, msg: Cat): Cat.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Cat, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Cat;
  static deserializeBinaryFromReader(message: Cat, reader: jspb.BinaryReader): Cat;
}

export namespace Cat {
  export type AsObject = {
    id: number,
    title: string,
    author: string,
  }
}

export class CatList extends jspb.Message {
  clearCatsList(): void;
  getCatsList(): Array<Cat>;
  setCatsList(value: Array<Cat>): void;
  addCats(value?: Cat, index?: number): Cat;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CatList.AsObject;
  static toObject(includeInstance: boolean, msg: CatList): CatList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CatList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CatList;
  static deserializeBinaryFromReader(message: CatList, reader: jspb.BinaryReader): CatList;
}

export namespace CatList {
  export type AsObject = {
    catsList: Array<Cat.AsObject>,
  }
}

export class CatIdRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CatIdRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CatIdRequest): CatIdRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CatIdRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CatIdRequest;
  static deserializeBinaryFromReader(message: CatIdRequest, reader: jspb.BinaryReader): CatIdRequest;
}

export namespace CatIdRequest {
  export type AsObject = {
    id: number,
  }
}

