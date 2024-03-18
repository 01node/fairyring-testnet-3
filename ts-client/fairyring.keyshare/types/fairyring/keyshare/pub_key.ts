/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "fairyring.keyshare";

export interface ActivePubKey {
  publicKey: string;
  creator: string;
  expiry: number;
  numberOfValidators: number;
}

export interface QueuedPubKey {
  publicKey: string;
  creator: string;
  expiry: number;
  numberOfValidators: number;
}

function createBaseActivePubKey(): ActivePubKey {
  return { publicKey: "", creator: "", expiry: 0, numberOfValidators: 0 };
}

export const ActivePubKey = {
  encode(message: ActivePubKey, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.publicKey !== "") {
      writer.uint32(10).string(message.publicKey);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.expiry !== 0) {
      writer.uint32(24).uint64(message.expiry);
    }
    if (message.numberOfValidators !== 0) {
      writer.uint32(32).uint64(message.numberOfValidators);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ActivePubKey {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseActivePubKey();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.publicKey = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.expiry = longToNumber(reader.uint64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.numberOfValidators = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ActivePubKey {
    return {
      publicKey: isSet(object.publicKey) ? String(object.publicKey) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      expiry: isSet(object.expiry) ? Number(object.expiry) : 0,
      numberOfValidators: isSet(object.numberOfValidators) ? Number(object.numberOfValidators) : 0,
    };
  },

  toJSON(message: ActivePubKey): unknown {
    const obj: any = {};
    if (message.publicKey !== "") {
      obj.publicKey = message.publicKey;
    }
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.expiry !== 0) {
      obj.expiry = Math.round(message.expiry);
    }
    if (message.numberOfValidators !== 0) {
      obj.numberOfValidators = Math.round(message.numberOfValidators);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ActivePubKey>, I>>(base?: I): ActivePubKey {
    return ActivePubKey.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ActivePubKey>, I>>(object: I): ActivePubKey {
    const message = createBaseActivePubKey();
    message.publicKey = object.publicKey ?? "";
    message.creator = object.creator ?? "";
    message.expiry = object.expiry ?? 0;
    message.numberOfValidators = object.numberOfValidators ?? 0;
    return message;
  },
};

function createBaseQueuedPubKey(): QueuedPubKey {
  return { publicKey: "", creator: "", expiry: 0, numberOfValidators: 0 };
}

export const QueuedPubKey = {
  encode(message: QueuedPubKey, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.publicKey !== "") {
      writer.uint32(10).string(message.publicKey);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.expiry !== 0) {
      writer.uint32(24).uint64(message.expiry);
    }
    if (message.numberOfValidators !== 0) {
      writer.uint32(32).uint64(message.numberOfValidators);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueuedPubKey {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueuedPubKey();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.publicKey = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.expiry = longToNumber(reader.uint64() as Long);
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.numberOfValidators = longToNumber(reader.uint64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueuedPubKey {
    return {
      publicKey: isSet(object.publicKey) ? String(object.publicKey) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      expiry: isSet(object.expiry) ? Number(object.expiry) : 0,
      numberOfValidators: isSet(object.numberOfValidators) ? Number(object.numberOfValidators) : 0,
    };
  },

  toJSON(message: QueuedPubKey): unknown {
    const obj: any = {};
    if (message.publicKey !== "") {
      obj.publicKey = message.publicKey;
    }
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.expiry !== 0) {
      obj.expiry = Math.round(message.expiry);
    }
    if (message.numberOfValidators !== 0) {
      obj.numberOfValidators = Math.round(message.numberOfValidators);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueuedPubKey>, I>>(base?: I): QueuedPubKey {
    return QueuedPubKey.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueuedPubKey>, I>>(object: I): QueuedPubKey {
    const message = createBaseQueuedPubKey();
    message.publicKey = object.publicKey ?? "";
    message.creator = object.creator ?? "";
    message.expiry = object.expiry ?? 0;
    message.numberOfValidators = object.numberOfValidators ?? 0;
    return message;
  },
};

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
