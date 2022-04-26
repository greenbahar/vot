/* eslint-disable */
import { Params } from "../vot/params";
import { NextVote } from "../vot/next_vote";
import { StoredVote } from "../vot/stored_vote";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "greenbahar.vot.vot";

/** GenesisState defines the vot module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  nextVote: NextVote | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  storedVoteList: StoredVote[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextVote !== undefined) {
      NextVote.encode(message.nextVote, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.storedVoteList) {
      StoredVote.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.storedVoteList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.nextVote = NextVote.decode(reader, reader.uint32());
          break;
        case 3:
          message.storedVoteList.push(
            StoredVote.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.storedVoteList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.nextVote !== undefined && object.nextVote !== null) {
      message.nextVote = NextVote.fromJSON(object.nextVote);
    } else {
      message.nextVote = undefined;
    }
    if (object.storedVoteList !== undefined && object.storedVoteList !== null) {
      for (const e of object.storedVoteList) {
        message.storedVoteList.push(StoredVote.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.nextVote !== undefined &&
      (obj.nextVote = message.nextVote
        ? NextVote.toJSON(message.nextVote)
        : undefined);
    if (message.storedVoteList) {
      obj.storedVoteList = message.storedVoteList.map((e) =>
        e ? StoredVote.toJSON(e) : undefined
      );
    } else {
      obj.storedVoteList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.storedVoteList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.nextVote !== undefined && object.nextVote !== null) {
      message.nextVote = NextVote.fromPartial(object.nextVote);
    } else {
      message.nextVote = undefined;
    }
    if (object.storedVoteList !== undefined && object.storedVoteList !== null) {
      for (const e of object.storedVoteList) {
        message.storedVoteList.push(StoredVote.fromPartial(e));
      }
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
