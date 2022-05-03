/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "greenbahar.vot.vot";

export interface MsgCreateVote {
  creator: string;
  question: string;
  options: string;
  days: number;
}

export interface MsgCreateVoteResponse {
  idValue: string;
}

export interface MsgSelectVotingOption {
  creator: string;
  idValue: string;
  selectedOption: number;
}

export interface MsgSelectVotingOptionResponse {
  idValue: string;
  results: MsgSelectVotingOptionResponse_Result[];
}

export interface MsgSelectVotingOptionResponse_Result {
  option: string;
  counter: number;
}

const baseMsgCreateVote: object = {
  creator: "",
  question: "",
  options: "",
  days: 0,
};

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
    if (message.days !== 0) {
      writer.uint32(32).uint32(message.days);
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
        case 4:
          message.days = reader.uint32();
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
    if (object.days !== undefined && object.days !== null) {
      message.days = Number(object.days);
    } else {
      message.days = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateVote): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.question !== undefined && (obj.question = message.question);
    message.options !== undefined && (obj.options = message.options);
    message.days !== undefined && (obj.days = message.days);
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
    if (object.days !== undefined && object.days !== null) {
      message.days = object.days;
    } else {
      message.days = 0;
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

const baseMsgSelectVotingOption: object = {
  creator: "",
  idValue: "",
  selectedOption: 0,
};

export const MsgSelectVotingOption = {
  encode(
    message: MsgSelectVotingOption,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.idValue !== "") {
      writer.uint32(18).string(message.idValue);
    }
    if (message.selectedOption !== 0) {
      writer.uint32(24).uint64(message.selectedOption);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSelectVotingOption {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSelectVotingOption } as MsgSelectVotingOption;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.idValue = reader.string();
          break;
        case 3:
          message.selectedOption = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSelectVotingOption {
    const message = { ...baseMsgSelectVotingOption } as MsgSelectVotingOption;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue);
    } else {
      message.idValue = "";
    }
    if (object.selectedOption !== undefined && object.selectedOption !== null) {
      message.selectedOption = Number(object.selectedOption);
    } else {
      message.selectedOption = 0;
    }
    return message;
  },

  toJSON(message: MsgSelectVotingOption): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.idValue !== undefined && (obj.idValue = message.idValue);
    message.selectedOption !== undefined &&
      (obj.selectedOption = message.selectedOption);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSelectVotingOption>
  ): MsgSelectVotingOption {
    const message = { ...baseMsgSelectVotingOption } as MsgSelectVotingOption;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue;
    } else {
      message.idValue = "";
    }
    if (object.selectedOption !== undefined && object.selectedOption !== null) {
      message.selectedOption = object.selectedOption;
    } else {
      message.selectedOption = 0;
    }
    return message;
  },
};

const baseMsgSelectVotingOptionResponse: object = { idValue: "" };

export const MsgSelectVotingOptionResponse = {
  encode(
    message: MsgSelectVotingOptionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.idValue !== "") {
      writer.uint32(10).string(message.idValue);
    }
    for (const v of message.results) {
      MsgSelectVotingOptionResponse_Result.encode(
        v!,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSelectVotingOptionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSelectVotingOptionResponse,
    } as MsgSelectVotingOptionResponse;
    message.results = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.idValue = reader.string();
          break;
        case 2:
          message.results.push(
            MsgSelectVotingOptionResponse_Result.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSelectVotingOptionResponse {
    const message = {
      ...baseMsgSelectVotingOptionResponse,
    } as MsgSelectVotingOptionResponse;
    message.results = [];
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue);
    } else {
      message.idValue = "";
    }
    if (object.results !== undefined && object.results !== null) {
      for (const e of object.results) {
        message.results.push(MsgSelectVotingOptionResponse_Result.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgSelectVotingOptionResponse): unknown {
    const obj: any = {};
    message.idValue !== undefined && (obj.idValue = message.idValue);
    if (message.results) {
      obj.results = message.results.map((e) =>
        e ? MsgSelectVotingOptionResponse_Result.toJSON(e) : undefined
      );
    } else {
      obj.results = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSelectVotingOptionResponse>
  ): MsgSelectVotingOptionResponse {
    const message = {
      ...baseMsgSelectVotingOptionResponse,
    } as MsgSelectVotingOptionResponse;
    message.results = [];
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue;
    } else {
      message.idValue = "";
    }
    if (object.results !== undefined && object.results !== null) {
      for (const e of object.results) {
        message.results.push(
          MsgSelectVotingOptionResponse_Result.fromPartial(e)
        );
      }
    }
    return message;
  },
};

const baseMsgSelectVotingOptionResponse_Result: object = {
  option: "",
  counter: 0,
};

export const MsgSelectVotingOptionResponse_Result = {
  encode(
    message: MsgSelectVotingOptionResponse_Result,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.option !== "") {
      writer.uint32(10).string(message.option);
    }
    if (message.counter !== 0) {
      writer.uint32(16).int64(message.counter);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSelectVotingOptionResponse_Result {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSelectVotingOptionResponse_Result,
    } as MsgSelectVotingOptionResponse_Result;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.option = reader.string();
          break;
        case 2:
          message.counter = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSelectVotingOptionResponse_Result {
    const message = {
      ...baseMsgSelectVotingOptionResponse_Result,
    } as MsgSelectVotingOptionResponse_Result;
    if (object.option !== undefined && object.option !== null) {
      message.option = String(object.option);
    } else {
      message.option = "";
    }
    if (object.counter !== undefined && object.counter !== null) {
      message.counter = Number(object.counter);
    } else {
      message.counter = 0;
    }
    return message;
  },

  toJSON(message: MsgSelectVotingOptionResponse_Result): unknown {
    const obj: any = {};
    message.option !== undefined && (obj.option = message.option);
    message.counter !== undefined && (obj.counter = message.counter);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSelectVotingOptionResponse_Result>
  ): MsgSelectVotingOptionResponse_Result {
    const message = {
      ...baseMsgSelectVotingOptionResponse_Result,
    } as MsgSelectVotingOptionResponse_Result;
    if (object.option !== undefined && object.option !== null) {
      message.option = object.option;
    } else {
      message.option = "";
    }
    if (object.counter !== undefined && object.counter !== null) {
      message.counter = object.counter;
    } else {
      message.counter = 0;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateVote(request: MsgCreateVote): Promise<MsgCreateVoteResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SelectVotingOption(
    request: MsgSelectVotingOption
  ): Promise<MsgSelectVotingOptionResponse>;
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

  SelectVotingOption(
    request: MsgSelectVotingOption
  ): Promise<MsgSelectVotingOptionResponse> {
    const data = MsgSelectVotingOption.encode(request).finish();
    const promise = this.rpc.request(
      "greenbahar.vot.vot.Msg",
      "SelectVotingOption",
      data
    );
    return promise.then((data) =>
      MsgSelectVotingOptionResponse.decode(new Reader(data))
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

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
