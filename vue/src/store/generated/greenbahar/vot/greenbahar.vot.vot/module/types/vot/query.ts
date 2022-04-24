/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../vot/params";
import { NextVote } from "../vot/next_vote";

export const protobufPackage = "greenbahar.vot.vot";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetNextVoteRequest {}

export interface QueryGetNextVoteResponse {
  NextVote: NextVote | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetNextVoteRequest: object = {};

export const QueryGetNextVoteRequest = {
  encode(_: QueryGetNextVoteRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNextVoteRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNextVoteRequest,
    } as QueryGetNextVoteRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetNextVoteRequest {
    const message = {
      ...baseQueryGetNextVoteRequest,
    } as QueryGetNextVoteRequest;
    return message;
  },

  toJSON(_: QueryGetNextVoteRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetNextVoteRequest>
  ): QueryGetNextVoteRequest {
    const message = {
      ...baseQueryGetNextVoteRequest,
    } as QueryGetNextVoteRequest;
    return message;
  },
};

const baseQueryGetNextVoteResponse: object = {};

export const QueryGetNextVoteResponse = {
  encode(
    message: QueryGetNextVoteResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.NextVote !== undefined) {
      NextVote.encode(message.NextVote, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNextVoteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNextVoteResponse,
    } as QueryGetNextVoteResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.NextVote = NextVote.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNextVoteResponse {
    const message = {
      ...baseQueryGetNextVoteResponse,
    } as QueryGetNextVoteResponse;
    if (object.NextVote !== undefined && object.NextVote !== null) {
      message.NextVote = NextVote.fromJSON(object.NextVote);
    } else {
      message.NextVote = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNextVoteResponse): unknown {
    const obj: any = {};
    message.NextVote !== undefined &&
      (obj.NextVote = message.NextVote
        ? NextVote.toJSON(message.NextVote)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNextVoteResponse>
  ): QueryGetNextVoteResponse {
    const message = {
      ...baseQueryGetNextVoteResponse,
    } as QueryGetNextVoteResponse;
    if (object.NextVote !== undefined && object.NextVote !== null) {
      message.NextVote = NextVote.fromPartial(object.NextVote);
    } else {
      message.NextVote = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a NextVote by index. */
  NextVote(request: QueryGetNextVoteRequest): Promise<QueryGetNextVoteResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "greenbahar.vot.vot.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  NextVote(
    request: QueryGetNextVoteRequest
  ): Promise<QueryGetNextVoteResponse> {
    const data = QueryGetNextVoteRequest.encode(request).finish();
    const promise = this.rpc.request(
      "greenbahar.vot.vot.Query",
      "NextVote",
      data
    );
    return promise.then((data) =>
      QueryGetNextVoteResponse.decode(new Reader(data))
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
