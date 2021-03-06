/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "greenbahar.vot.vot";

export interface StoredVote {
  index: string;
  vote: string;
  creator: string;
  voters: StoredVote_Voter[];
}

export interface StoredVote_Voter {
  voter: string;
}

const baseStoredVote: object = { index: "", vote: "", creator: "" };

export const StoredVote = {
  encode(message: StoredVote, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.vote !== "") {
      writer.uint32(18).string(message.vote);
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    for (const v of message.voters) {
      StoredVote_Voter.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredVote {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredVote } as StoredVote;
    message.voters = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.vote = reader.string();
          break;
        case 3:
          message.creator = reader.string();
          break;
        case 4:
          message.voters.push(StoredVote_Voter.decode(reader, reader.uint32()));
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
    message.voters = [];
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.voters !== undefined && object.voters !== null) {
      for (const e of object.voters) {
        message.voters.push(StoredVote_Voter.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: StoredVote): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.vote !== undefined && (obj.vote = message.vote);
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.voters) {
      obj.voters = message.voters.map((e) =>
        e ? StoredVote_Voter.toJSON(e) : undefined
      );
    } else {
      obj.voters = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<StoredVote>): StoredVote {
    const message = { ...baseStoredVote } as StoredVote;
    message.voters = [];
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.voters !== undefined && object.voters !== null) {
      for (const e of object.voters) {
        message.voters.push(StoredVote_Voter.fromPartial(e));
      }
    }
    return message;
  },
};

const baseStoredVote_Voter: object = { voter: "" };

export const StoredVote_Voter = {
  encode(message: StoredVote_Voter, writer: Writer = Writer.create()): Writer {
    if (message.voter !== "") {
      writer.uint32(10).string(message.voter);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredVote_Voter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredVote_Voter } as StoredVote_Voter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.voter = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredVote_Voter {
    const message = { ...baseStoredVote_Voter } as StoredVote_Voter;
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = String(object.voter);
    } else {
      message.voter = "";
    }
    return message;
  },

  toJSON(message: StoredVote_Voter): unknown {
    const obj: any = {};
    message.voter !== undefined && (obj.voter = message.voter);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredVote_Voter>): StoredVote_Voter {
    const message = { ...baseStoredVote_Voter } as StoredVote_Voter;
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = object.voter;
    } else {
      message.voter = "";
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
