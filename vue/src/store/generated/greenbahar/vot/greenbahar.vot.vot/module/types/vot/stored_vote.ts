/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "greenbahar.vot.vot";

export interface StoredVote {
  index: string;
  vote: string;
}

const baseStoredVote: object = { index: "", vote: "" };

export const StoredVote = {
  encode(message: StoredVote, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.vote !== "") {
      writer.uint32(18).string(message.vote);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredVote {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredVote } as StoredVote;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.vote = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredVote {
    const message = { ...baseStoredVote } as StoredVote;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.vote !== undefined && object.vote !== null) {
      message.vote = String(object.vote);
    } else {
      message.vote = "";
    }
    return message;
  },

  toJSON(message: StoredVote): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.vote !== undefined && (obj.vote = message.vote);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredVote>): StoredVote {
    const message = { ...baseStoredVote } as StoredVote;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.vote !== undefined && object.vote !== null) {
      message.vote = object.vote;
    } else {
      message.vote = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
