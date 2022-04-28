/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "greenbahar.vot.vot";

export interface MsgCreateVote {
  creator: string;
  question: string;
  options: string;
}

export interface MsgCreateVoteResponse {
  idValue: string;
}

const baseMsgCreateVote: object = { creator: "", question: "", options: "" };

export const MsgCreateVote = {
  encode(message: MsgCreateVote, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.question !== "") {
      writer.uint32(18).string(message.question);
    }
    if (message.options !== "") {
      writer.uint32(26).string(message.options);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVote {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateVote } as MsgCreateVote;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.question = reader.string();
          break;
        case 3:
          message.options = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateVote {
    const message = { ...baseMsgCreateVote } as MsgCreateVote;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.question !== undefined && object.question !== null) {
      message.question = String(object.question);
    } else {
      message.question = "";
    }
    if (object.options !== undefined && object.options !== null) {
      message.options = String(object.options);
    } else {
      message.options = "";
    }
    return message;
  },

  toJSON(message: MsgCreateVote): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.question !== undefined && (obj.question = message.question);
    message.options !== undefined && (obj.options = message.options);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateVote>): MsgCreateVote {
    const message = { ...baseMsgCreateVote } as MsgCreateVote;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.question !== undefined && object.question !== null) {
      message.question = object.question;
    } else {
      message.question = "";
    }
    if (object.options !== undefined && object.options !== null) {
      message.options = object.options;
    } else {
      message.options = "";
    }
    return message;
  },
};

const baseMsgCreateVoteResponse: object = { idValue: "" };

export const MsgCreateVoteResponse = {
  encode(
    message: MsgCreateVoteResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.idValue !== "") {
      writer.uint32(10).string(message.idValue);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVoteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateVoteResponse } as MsgCreateVoteResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.idValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateVoteResponse {
    const message = { ...baseMsgCreateVoteResponse } as MsgCreateVoteResponse;
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue);
    } else {
      message.idValue = "";
    }
    return message;
  },

  toJSON(message: MsgCreateVoteResponse): unknown {
    const obj: any = {};
    message.idValue !== undefined && (obj.idValue = message.idValue);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateVoteResponse>
  ): MsgCreateVoteResponse {
    const message = { ...baseMsgCreateVoteResponse } as MsgCreateVoteResponse;
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue;
    } else {
      message.idValue = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateVote(request: MsgCreateVote): Promise<MsgCreateVoteResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateVote(request: MsgCreateVote): Promise<MsgCreateVoteResponse> {
    const data = MsgCreateVote.encode(request).finish();
    const promise = this.rpc.request(
      "greenbahar.vot.vot.Msg",
      "CreateVote",
      data
    );
    return promise.then((data) =>
      MsgCreateVoteResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
