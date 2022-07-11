/**
 * @fileoverview gRPC-Web generated client stub for protobuf
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.protobuf = require('./news_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.protobuf.NewServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.protobuf.NewServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.protobuf.Request,
 *   !proto.protobuf.Response>}
 */
const methodDescriptor_NewService_GetHotTopNews = new grpc.web.MethodDescriptor(
  '/protobuf.NewService/GetHotTopNews',
  grpc.web.MethodType.UNARY,
  proto.protobuf.Request,
  proto.protobuf.Response,
  /**
   * @param {!proto.protobuf.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protobuf.Response.deserializeBinary
);


/**
 * @param {!proto.protobuf.Request} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.protobuf.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.protobuf.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.protobuf.NewServiceClient.prototype.getHotTopNews =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/protobuf.NewService/GetHotTopNews',
      request,
      metadata || {},
      methodDescriptor_NewService_GetHotTopNews,
      callback);
};


/**
 * @param {!proto.protobuf.Request} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.protobuf.Response>}
 *     Promise that resolves to the response
 */
proto.protobuf.NewServicePromiseClient.prototype.getHotTopNews =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/protobuf.NewService/GetHotTopNews',
      request,
      metadata || {},
      methodDescriptor_NewService_GetHotTopNews);
};


module.exports = proto.protobuf;

